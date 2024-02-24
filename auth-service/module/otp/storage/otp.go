package storage

import (
	"account-service/module/otp/dto"
	"account-service/module/otp/model"
	"context"
	"gorm.io/gorm"
)

type OtpStorage interface {
	Create(ctx context.Context, otp *model.Otp) error
	Verify(ctx context.Context, message dto.OtpVerifyRequest) (*model.Otp, error)
	FindUserIdByPhoneNumber(ctx context.Context, phoneNumber string) (*dto.GetUserIdByPhoneNumber, error)
}

type otpStorage struct {
	db *gorm.DB
}

func New(db *gorm.DB) OtpStorage {
	return &otpStorage{
		db: db,
	}
}

func (os *otpStorage) Create(ctx context.Context, otp *model.Otp) error {
	if err := os.db.Table(model.Otp{}.TableName()).Create(otp).Error; err != nil {
		return err
	}
	return nil
}

func (os *otpStorage) Verify(ctx context.Context, message dto.OtpVerifyRequest) (*model.Otp, error) {
	var otp model.Otp
	if err := os.db.Table(model.Otp{}.TableName()).
		Where("otp_code = ?", message.Code).
		Where("phone_number = ?", message.PhoneNumber).
		Where("expires_at > now()").
		Where("is_valid = true").
		First(&otp).Error; err != nil {
		return nil, err
	}
	otp.IsValid = false
	if err := os.db.Table(model.Otp{}.TableName()).Save(&otp).Error; err != nil {
		return nil, err
	}
	return &otp, nil
}

func (os *otpStorage) FindUserIdByPhoneNumber(ctx context.Context, phoneNumber string) (*dto.GetUserIdByPhoneNumber, error) {
	var resp dto.GetUserIdByPhoneNumber
	if err := os.db.Table(model.Otp{}.TableName()).
		Where("phone_number = ?", phoneNumber).
		Select([]string{
			"user_id",
		}).
		First(&resp).Error; err != nil {
		return nil, err
	}
	return &resp, nil
}

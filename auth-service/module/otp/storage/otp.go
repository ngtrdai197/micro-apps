package storage

import (
	"account-service/module/otp/model"
	"context"
	"gorm.io/gorm"
)

type OtpStorage interface {
	Create(ctx context.Context, otp *model.Otp) error
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

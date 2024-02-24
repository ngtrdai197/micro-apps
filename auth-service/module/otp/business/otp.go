package business

import (
	"account-service/module/otp/dto"
	"account-service/module/otp/model"
	"account-service/module/otp/storage"
	"context"
)

type Business interface {
	CreateOtp(ctx context.Context, otp *model.Otp) error
	VerifyOtp(ctx context.Context, message dto.OtpVerifyRequest) (*model.Otp, error)
	FindUserIdByPhoneNumber(ctx context.Context, phoneNumber string) (*dto.GetUserIdByPhoneNumber, error)
}

type business struct {
	otpStorage storage.OtpStorage
}

func New(otpStorage storage.OtpStorage) Business {
	return &business{
		otpStorage: otpStorage,
	}
}

func (b *business) CreateOtp(ctx context.Context, otp *model.Otp) error {
	if err := b.otpStorage.Create(ctx, otp); err != nil {
		return err
	}
	return nil
}

func (b *business) VerifyOtp(ctx context.Context, message dto.OtpVerifyRequest) (*model.Otp, error) {
	otp, err := b.otpStorage.Verify(ctx, message)
	if err != nil {
		return nil, err
	}
	return otp, nil
}

func (b *business) FindUserIdByPhoneNumber(ctx context.Context, phoneNumber string) (*dto.GetUserIdByPhoneNumber, error) {
	resp, err := b.otpStorage.FindUserIdByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

package business

import (
	"account-service/module/otp/model"
	"account-service/module/otp/storage"
	"context"
)

type Business interface {
	CreateOtp(ctx context.Context, otp *model.Otp) error
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

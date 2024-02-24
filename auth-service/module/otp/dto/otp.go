package dto

import (
	"account-service/module/otp/model"
	"github.com/google/uuid"
	"time"
)

type OtpRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
}
type OtpResponse struct {
	Code string `json:"code"`
}

func (req *OtpRequest) TransformToModel(code string) *model.Otp {
	return &model.Otp{
		UserId:      uuid.NewString(),
		PhoneNumber: req.PhoneNumber,
		ExpiresAt:   time.Now().Add(time.Minute * 5),
		OtpCode:     code,
		IsValid:     true,
	}
}

type GetUserIdByPhoneNumber struct {
	PhoneNumber string `json:"phone_number"`
	UserId      string `json:"user_id"`
}

type OtpVerifyRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
	Code        string `json:"code" validate:"required"`
}

type OtpVerifyResponse struct {
	UserId      string `json:"user_id"`
	PhoneNumber string `json:"phone_number"`
}

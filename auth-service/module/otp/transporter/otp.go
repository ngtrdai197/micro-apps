package otp_trpt

import (
	"account-service/config"
	"account-service/module/otp/business"
	"account-service/module/otp/dto"
	"account-service/module/otp/storage"
	"account-service/pkg/helper"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type OtpTransporter interface {
	SendOtp(c *fiber.Ctx) error
}
type otpTransporter struct {
	db        *gorm.DB
	validator *validator.Validate
}

func New(db *gorm.DB, validator *validator.Validate) OtpTransporter {
	return &otpTransporter{
		db:        db,
		validator: validator,
	}
}

func (ot *otpTransporter) SendOtp(c *fiber.Ctx) error {
	message := new(dto.OtpRequest)
	if err := c.BodyParser(message); err != nil {
		helper.JsonResponse(c, "Invalid request")
		return nil
	}
	if err := ot.validator.Struct(message); err != nil {
		helper.JsonResponse(c, "Invalid request")
		return nil
	}
	otp := helper.GenerateOTP()
	msg := message.TransformToModel(otp)
	otpBusiness := business.New(storage.New(ot.db))

	if err := otpBusiness.CreateOtp(c.Context(), msg); err != nil {
		helper.JsonResponse(c, "Failed to send OTP")
		return nil
	}
	log.Info().Str("user_id", msg.UserId).Str("otp", otp).Msg("OTP sent")
	log.Info().Str("pg_dsn", config.Config.PostgresDSN).Msg("Postgres DSN")

	helper.JsonResponse(c, dto.OtpResponse{
		Code: otp,
	})
	return nil
}

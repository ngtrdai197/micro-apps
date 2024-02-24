package otp_trpt

import (
	"account-service/module/otp/business"
	"account-service/module/otp/dto"
	"account-service/module/otp/storage"
	"account-service/pkg/helper"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type OtpTransporter interface {
	SendOtp(c *fiber.Ctx) error
	VerifyOtp(c *fiber.Ctx) error
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
	otpBusiness := business.New(storage.New(ot.db))

	resp, err := otpBusiness.FindUserIdByPhoneNumber(c.Context(), message.PhoneNumber)
	if err != nil {
		return err
	}

	otp := helper.GenerateOTP()
	msg := message.TransformToModel(otp)

	if resp != nil || resp.UserId != "" {
		msg.UserId = resp.UserId
	}

	if err := otpBusiness.CreateOtp(c.Context(), msg); err != nil {
		helper.JsonResponse(c, "Failed to send OTP")
		return nil
	}

	helper.JsonResponse(c, dto.OtpResponse{
		Code: otp,
	})
	return nil
}

func (ot *otpTransporter) VerifyOtp(c *fiber.Ctx) error {
	var message dto.OtpVerifyRequest
	if err := c.BodyParser(&message); err != nil {
		helper.JsonResponse(c, "Invalid request")
		return nil
	}
	if err := ot.validator.Struct(message); err != nil {
		helper.JsonResponse(c, "Invalid request")
		return nil
	}
	otpBusiness := business.New(storage.New(ot.db))

	otp, err := otpBusiness.VerifyOtp(c.Context(), message)
	if err != nil {
		helper.JsonResponse(c, "Invalid OTP")
		return nil
	}
	helper.JsonResponse(c, dto.OtpVerifyResponse{UserId: otp.UserId, PhoneNumber: otp.PhoneNumber})
	return nil
}

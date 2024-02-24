package main

import (
	otp_trpt "account-service/module/otp/transporter"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func OtpRoutesRegister(app *fiber.App, db *gorm.DB) {
	// OtpRegister registers the routes of the Otp domain
	api := app.Group("/api")
	v1 := api.Group("/v1/otp")

	trpt := otp_trpt.New(db, validator.New())
	v1.Post("/", trpt.SendOtp)
	v1.Post("/verify", trpt.VerifyOtp)
}

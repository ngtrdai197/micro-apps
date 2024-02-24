package model

import (
	"gorm.io/gorm"
	"time"
)

type Otp struct {
	*gorm.Model
	UserId      string    `json:"user_id" gorm:"user_id"`
	PhoneNumber string    `json:"phone_number" gorm:"phone_number"`
	OtpCode     string    `json:"otp_code" gorm:"otp_code"`
	IsValid     bool      `json:"is_validate" gorm:"is_valid"`
	ExpiresAt   time.Time `json:"expires_at" gorm:"expires_at"`
}

func (o Otp) TableName() string {
	return "otp"
}

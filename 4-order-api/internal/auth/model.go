package auth

import "gorm.io/gorm"

type Auth struct {
	gorm.Model
	Phone     string `gorm:"uniqueIndex"`
	Code      int
	SessionId string
}

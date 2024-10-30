package repo

import (
	"fmt"
	"time"

	"ecommerce/global"
)

type IUserAuthRepository interface {
	AddOTP(email string, otp int, expirationTime int64) error
}

type userAuthRepository struct{}

func NewUserAuthRepository() IUserAuthRepository {
	return &userAuthRepository{}
}

// AddOTP implements IUserAuthRepository.e
func (u *userAuthRepository) AddOTP(email string, otp int, expirationTime int64) error {
	// panic("unimplemented")
	key := fmt.Sprintf("usr:%s:otp", email)
	return global.Rdb.SetEx(ctx, key, otp, time.Duration(expirationTime)).Err()
}
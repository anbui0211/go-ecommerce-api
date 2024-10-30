package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"ecommerce/global"
	"ecommerce/internal/repo"
	"ecommerce/internal/utils/crypto"
	"ecommerce/internal/utils/random"
	"ecommerce/pkg/response"

	"github.com/segmentio/kafka-go"
)

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo     repo.IUserRepository
	userAuthRepo repo.IUserAuthRepository
}

func NewUserService(userRepo repo.IUserRepository, userAuthRep repo.IUserAuthRepository) IUserService {
	return &userService{userRepo: userRepo, userAuthRepo: userAuthRep}

}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	// 0. Hash Email
	hashEmail := crypto.GetHash(email)

	// 5. check OTP is available
	// 6. user spam

	// 1. Check email exist in email
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrorCodeUserHasExists
	}

	// 2. New OTP
	otp := random.GenerateSixDigitOtp()
	if purpose == "TEST_USER" {
		otp = 123456
	}
	fmt.Println("OTP:: ", otp)

	// 3. Save OTP in redis with expiration time
	err := us.userAuthRepo.AddOTP(hashEmail, otp, int64(10*time.Minute))
	if err != nil {
		fmt.Println("err:: ", err)
		return response.ErrorInvalidOTP
	}

	// 4. Send mail OTP
	// senderEmail := "shopdev@gmail.com"
	// nameTemplate := "otp-auth.html"
	// optCode := "OTP_CODE"
	// err = sendto.SendTemplateEmailOTP([]string{email}, senderEmail, nameTemplate, map[string]interface{}{
	// 	optCode: strconv.Itoa(otp),
	// })
	// if err != nil {
	// 	fmt.Printf("error send otp::%d\n ", err)
	// 	return response.ErrorSendEmailOTP
	// }

	// Send email OTP by Java
	// err = sendto.SenEmailToJavaAPI(strconv.Itoa(otp), email, "otp-auth.html")
	// if err != nil {
	// 	fmt.Printf("error send to JAVA::%d\n ", err)
	// 	return response.ErrorSendEmailOTP
	// }

	// Send OTP via Kafka Java
	body := make(map[string]interface{})
	body["otp"] = otp
	body["email"] = email

	bodyRequest, _ := json.Marshal(body)
	message := kafka.Message{
		Key:   []byte("otp-auth"),
		Value: []byte(bodyRequest),
		Time:  time.Now(),
	}

	err = global.KafkaProducer.WriteMessages(context.Background(), message)
	if err != nil {
		return response.ErrorSendEmailOTP
	}

	return response.ErrorCodeSuccess
}

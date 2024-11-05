package impl

import (
	"context"
	"database/sql"
	"ecommerce/global"
	"ecommerce/internal/consts"
	"ecommerce/internal/database"
	"ecommerce/internal/model"
	"ecommerce/internal/utils"
	"ecommerce/internal/utils/auth"
	"ecommerce/internal/utils/crypto"
	"ecommerce/internal/utils/random"
	"ecommerce/internal/utils/sendto"
	"ecommerce/pkg/response"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type sUserLogin struct {
	// Implement the IUserLogin interface here
	r *database.Queries
}

func NewUserLoginImpl(r *database.Queries) *sUserLogin {
	return &sUserLogin{r: r}
}

func (s *sUserLogin) Login(ctx context.Context, in *model.LoginInput) (codeResult int, out model.LoginOutput, err error) {
	// logic login
	userBase, err := s.r.GetOneUserInfo(ctx, in.UserAccount)
	if err != nil {
		return response.ErrorCodeAuthFailed, out, err
	}

	// 2. check password
	if !crypto.MatchingPassword(userBase.UserPassword, in.UserPassword, userBase.UserSalt) {
		return response.ErrorCodeAuthFailed, out, fmt.Errorf("does not match password")
	}

	// 3. check two-factor authentication
	// 4.update password time
	go s.r.LoginUserBase(ctx, database.LoginUserBaseParams{
		UserLoginIp:  sql.NullString{String: "127.0.0.1", Valid: true}, // IP temp
		UserAccount:  in.UserAccount,
		UserPassword: in.UserPassword,
	})

	// 5. Create UUID user
	subToken := utils.GenerateUUID(int(userBase.UserID))

	// 6. Get user_info table
	infoUser, err := s.r.GetUser(ctx, uint64(userBase.UserID))
	if err != nil {
		return response.ErrorCodeAuthFailed, out, err
	}

	// convert to json to save in redis
	infoUserJson, err := json.Marshal(infoUser)
	if err != nil {
		return response.ErrorCodeAuthFailed, out, fmt.Errorf("convert to json failed: %v", err)
	}

	// 7. Give infoUserJson to redis with key = subToken
	if err = global.Rdb.Set(ctx, subToken, infoUserJson, time.Duration(consts.TIME_OTP_REGISTER)*time.Minute).Err(); err != nil {
		return response.ErrorCodeAuthFailed, out, err
	}

	// 8. Create token
	out.Token, err = auth.CrateToken(subToken)
	if err != nil {
		return
	}

	return 200, out, nil
}

func (s *sUserLogin) Register(ctx context.Context, in *model.RegisterInput) (codeResult int, err error) {
	// 1. hash email
	hashKey := crypto.GetHash(strings.ToLower(in.VerifyKey))

	// 2. check user exists in user base
	userFound, err := s.r.CheckUserBaseExists(ctx, in.VerifyKey)
	if err != nil {
		return response.ErrorCodeUserHasExists, err
	}

	if userFound > 0 {
		return response.ErrorCodeUserHasExists, fmt.Errorf("user has already registered")
	}

	// 3. create OTP
	userKey := utils.GetUserKey(hashKey) //fmt.Sprintf("u:%s:otp", hashKey)
	otpFound, err := global.Rdb.Get(ctx, userKey).Result()

	// util..
	switch {
	case errors.Is(err, redis.Nil):
		fmt.Println("Key does not exist")
	case err != nil:
		fmt.Println("Get failed::", err)
		return response.ErrorInvalidOTP, err
	case otpFound != "":
		return response.ErrorCodeOtpNotExists, fmt.Errorf("")
	}

	// 4. Generate OTP
	otpNew := random.GenerateSixDigitOtp()
	if in.VerifyPurpose == "TEST_USER" {
		otpNew = 123456
	}

	// 5. Save OTP to Redis with expiration time
	err = global.Rdb.SetEx(ctx, userKey, strconv.Itoa(otpNew), time.Duration(consts.TIME_OTP_REGISTER)*time.Minute).Err()
	if err != nil {
		return response.ErrorInvalidOTP, err
	}

	// 6. Send OTP
	switch in.VerifyType {
	case consts.EMAIL:
		if err = sendto.SendTextEmailOTP([]string{in.VerifyKey}, consts.HOST_EMAIL, strconv.Itoa(otpNew)); err != nil {
			return response.ErrorSendEmailOTP, err
		}

		// 7. Save OTP in MySQL
		result, err := s.r.InsertOTPVerify(ctx, database.InsertOTPVerifyParams{
			VerifyOtp: strconv.Itoa(otpNew),
			VerifyType: sql.NullInt32{
				Int32: 1, Valid: true,
			},
			VerifyKey:     in.VerifyKey,
			VerifyKeyHash: hashKey,
		})
		if err != nil {
			return response.ErrorSendEmailOTP, err
		}

		// 8.GetLastId (muốn lấy id để làm việc tiếp)
		lastIdVerifyUser, err := result.LastInsertId()
		if err != nil {
			return response.ErrorSendEmailOTP, err
		}
		fmt.Println("lastIdVerifyUser:: ", lastIdVerifyUser)
		return response.ErrorCodeSuccess, nil

	case consts.MOBILE:
		return response.ErrorCodeSuccess, nil
	}

	return response.ErrorCodeSuccess, nil
}

func (s *sUserLogin) VerifyOTP(ctx context.Context, in *model.VerifyInput) (out model.VerifyOTPOutput, err error) {
	// login
	hashKey := crypto.GetHash(strings.ToLower(in.VerifyKey))

	// get OTP
	otpFound, err := global.Rdb.Get(ctx, utils.GetUserKey(hashKey)).Result()
	if err != nil {
		return out, err
	}

	if otpFound != in.VerifyCode {
		// Nếu như sai 3 lần trong vòng 1 phút
		return out, fmt.Errorf("OTP not match")
	}

	infoOTP, err := s.r.GetInfoOTP(ctx, hashKey)
	if err != nil {
		return out, err
	}

	// success -> update status verify
	err = s.r.UpdateUserVerificationStatus(ctx, hashKey)
	if err != nil {
		return out, err
	}

	// output
	out.Token = infoOTP.VerifyKeyHash
	out.Message = "success"

	return out, err
}

func (s *sUserLogin) UpdatePasswordRegister(ctx context.Context, token string, password string) (userId int, err error) {
	// 1. token is already verified. : user_verify table
	infoOTP, err := s.r.GetInfoOTP(ctx, token)
	if err != nil {
		return response.ErrCodeUserOtpNotExists, err
	}

	// check isVerified OK
	if infoOTP.IsVerified.Int32 == 0 {
		return response.ErrCodeUserOtpNotExists, fmt.Errorf("user OTP not verified")
	}

	// update user_base table
	userBase := database.AddUserBaseParams{}
	userBase.UserAccount = infoOTP.VerifyKey
	userSalt, err := crypto.GenerateSalt(16)
	if err != nil {
		return response.ErrCodeUserOtpNotExists, err
	}
	userBase.UserSalt = userSalt
	userBase.UserPassword = crypto.HashPassword(password, userSalt)

	// add userBase to user_base table
	newUserBase, err := s.r.AddUserBase(ctx, userBase)
	if err != nil {
		return response.ErrCodeUserOtpNotExists, err
	}
	user_id, err := newUserBase.LastInsertId()
	if err != nil {
		return response.ErrCodeUserOtpNotExists, err
	}

	// add user_id to user_info table
	newUserInfo, err := s.r.AddUserHaveUserId(ctx, database.AddUserHaveUserIdParams{
		UserID:               uint64(user_id),
		UserAccount:          infoOTP.VerifyKey,
		UserNickname:         sql.NullString{String: infoOTP.VerifyKey, Valid: true},
		UserAvatar:           sql.NullString{String: "", Valid: true},
		UserState:            1,
		UserMobile:           sql.NullString{String: "", Valid: true},
		UserGender:           sql.NullInt16{Int16: 0, Valid: true},
		UserBirthday:         sql.NullTime{Time: time.Time{}, Valid: false},
		UserEmail:            sql.NullString{String: infoOTP.VerifyKey, Valid: true},
		UserIsAuthentication: 1,
	})
	if err != nil {
		return response.ErrCodeUserOtpNotExists, err
	}

	user_id, err = newUserInfo.LastInsertId()
	if err != nil {
		return response.ErrCodeUserOtpNotExists, err
	}

	return int(user_id), nil
}

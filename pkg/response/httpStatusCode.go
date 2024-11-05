package response

const (
	ErrorCodeSuccess      = 20001 // Success
	ErrorCodeParamInvalid = 20003 // Email invalid

	ErrorCodeInvalidToken = 30001
	ErrorInvalidOTP       = 30002
	ErrorSendEmailOTP     = 30003

	// USer Authentication
	ErrorCodeAuthFailed = 40005

	// Register code
	ErrorCodeUserHasExists = 50001 // User has already already registered

	// Error login
	ErrorCodeOtpNotExists   = 60009
	ErrCodeUserOtpNotExists = 60008
)

// Message
var msg = map[int]string{
	ErrorCodeSuccess:        "Success",
	ErrorCodeParamInvalid:   "Email is invalid",
	ErrorCodeInvalidToken:   "Token is invalid",
	ErrorInvalidOTP:         "Otp error",
	ErrorSendEmailOTP:       "Failed to send email OTP",
	ErrorCodeUserHasExists:  "User has already been registered",
	ErrorCodeOtpNotExists:   "OTP exists but not registered",
	ErrCodeUserOtpNotExists: "User OTP not exists",
	ErrorCodeAuthFailed:     "Authentication failed",
}

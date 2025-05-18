package response

const (
	ErrCodeSuccess      = 20001 // success
	ErrCodeParamInvalid = 10003 // email invalid
	ErrInvalidToken     = 30001 // token invalid
	ErrInvalidOTP       = 30002
	ErrSendEmailOtp     = 30003
	ErrUserNotFound     = 30004

	// register code
	ErrCodeUserHasExists     = 50001
	ErrCreateUserError       = 50002
	ErrInvalidUserOrPassword = 50003
)

var msg = map[int]string{
	ErrCodeSuccess:       "success",
	ErrCodeParamInvalid:  "<UNK>",
	ErrInvalidToken:      "invalid token",
	ErrInvalidOTP:        "invalid OTP",
	ErrSendEmailOtp:      "failed to send email otp",
	ErrCodeUserHasExists: "User has exists",
	ErrCreateUserError:   "failed to create user",
}

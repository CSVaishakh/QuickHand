package src

import (
	"github.com/golang-jwt/jwt/v5"
)

type UserRole string
type HandymanType string

const (
	ClientRole   UserRole = "client"
	HandymanRole UserRole = "handyman"
)

const (
	Plumber        HandymanType = "plumber"
	Electrician    HandymanType = "electrician"
	Carpenter      HandymanType = "carpenter"
	Mason          HandymanType = "mason"
	Mechanic       HandymanType = "mechanic"
	HVACTechnician HandymanType = "hvac_technician"
	Landscaper     HandymanType = "landscaper"
	DeepCleaner    HandymanType = "deep_cleaner"
)

type HandymanSignUpReq struct {
	FirstName   string
	LastName    string
	Email       string
	Password    string
	Role        UserRole
	Type        HandymanType
	PhoneNumber string
	Img         *string
}

type HandymanSignUpRes struct {
	UserID string
	FirstName string
	Token string
	Role UserRole
	Type HandymanType
}

type Claims struct {
	UserID string
	FirstName string
	Role   UserRole
	jwt.RegisteredClaims
}

type VerifySessionReq struct {
	Token string
}

type ClientSignUpReq struct {
	FirstName   string
	LastName    string
	Email       string
	Password    string
	Role        UserRole
	Type        HandymanType
	PhoneNumber string
	Img         *string
}

type ClientSignUpRes struct {
	UserID string
	FirstName string
	Token string
	Role UserRole
}

type SignInReq struct {
	Email    string
	Password string
}

type HandymanSignInRes struct {
	UserID string
	FirstName string
	Token string
	Role UserRole
	Type HandymanType
}

type ClientSignInRes struct {
	UserID string
	FirstName string
	Token string
	Role UserRole
}

type ForgotPasswordReq struct {
	Email string
}
type OtpJWT_Claims struct{
	OtpHash string
	Email string
	jwt.RegisteredClaims
}

type OtpVerificationReq struct {
	Email string
	Otp string
	Token string
}

type ResetJWT_Claims struct{
	Email string
	jwt.RegisteredClaims
}

type ResetPasswordReq struct {
	Token string
	NewPassword string
}

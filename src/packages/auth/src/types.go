package src

import (
	"github.com/CSVaishakh/QuickHand/src/packages/db/models"
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
	Session *models.Session
	User    *models.Handyman
}

type JWTService struct {
	secret []byte
}

type Claims struct {
	UserID string
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

type SignInReq struct {
	Email    string
	Password string
}

type HandymanSignInRes struct {
	Token string
}

type ClientSignInRes struct {
	Token string
}

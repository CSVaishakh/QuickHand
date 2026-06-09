package src

import (
	"gorm.io/gorm"

	repo "github.com/CSVaishakh/QuickHand/src/packages/db/repositories"
)

type AuthService struct {
	handymenRepo      *repo.HandymenRepository
	clientRepo *repo.ClientRepository
	sessionRepo *repo.SessionRepository
	jwtService        *JWTService
	db                *gorm.DB
}

func NewAuthService(
	handymenRepo *repo.HandymenRepository,
	clientRepo *repo.ClientRepository,
	sessionRepo *repo.SessionRepository,
	jwtService *JWTService,
	db *gorm.DB,
) *AuthService {
	return &AuthService{
		handymenRepo:      handymenRepo,
		clientRepo:  clientRepo,
		sessionRepo: sessionRepo,
		jwtService:        jwtService,
		db:                db,
	}
}
package src

import (
	"gorm.io/gorm"

	repo "github.com/CSVaishakh/QuickHand/src/packages/db/repositories"
)

type AuthService struct {
	handymenRepo      *repo.HandymenRepository
	clientRepository  *repo.ClientRepository
	sessionRepository *repo.SessionRepository
	jwtService        *JWTService
	db                *gorm.DB
}

func NewAuthService(
	handymenRepo *repo.HandymenRepository,
	clientRepository *repo.ClientRepository,
	sessionRepository *repo.SessionRepository,
	jwtService *JWTService,
	db *gorm.DB,
) *AuthService {
	return &AuthService{
		handymenRepo:      handymenRepo,
		clientRepository:  clientRepository,
		sessionRepository: sessionRepository,
		jwtService:        jwtService,
		db:                db,
	}
}
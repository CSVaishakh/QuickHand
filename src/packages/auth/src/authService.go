package src

import (
	"errors"

	"github.com/CSVaishakh/QuickHand/src/packages/db/models"
	repo "github.com/CSVaishakh/QuickHand/src/packages/db/repositories"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
		ErrInvalidToken = errors.New("invalid token")
		ErrSignOutFailed = errors.New("sign out failed")
		ErrSessionNotFound = errors.New("session not found")
		ErrInvalidSession = errors.New("invalid session")
		ErrSessionExpired = errors.New("session expired")
		ErrDBFailed = errors.New("session expired but databse failed")
)

type AuthService struct {
	handymenRepo *repo.HandymenRepository 
	clientRepository *repo.ClientRepository
	sessionRepository *repo.SessionRepository
	jwtService *JWTService
	db *gorm.DB
}

func NewAuthService(
	handymenRepo *repo.HandymenRepository, 
	clientRepository *repo.ClientRepository,
	sessionRepository *repo.SessionRepository,
	jwtService *JWTService,
	db *gorm.DB,
) *AuthService {
	return &AuthService{
		handymenRepo: handymenRepo,
		clientRepository: clientRepository,
		sessionRepository: sessionRepository,
		jwtService: jwtService,
		db: db,
	}
}

func (s *AuthService) HandymanSignUp (req HandymanSignUpReq) error {
	return s.db.Transaction( func (tx *gorm.DB) error {
		// Checking for exisitng user
		userExists, err := s.handymenRepo.GetByEmail(req.Email, tx)

		if err != nil {
			return err
		}

		if userExists {
			return errors.New("User with email exists")
		}

		// Password hashing
		hashedPass, err := bcrypt.GenerateFromPassword(
			[]byte(req.Password),
			bcrypt.DefaultCost,
		)

		if err != nil {
			return err
		}

		// Creating the user object
		user := &models.Handyman{
			User: models.User{
				FirstName:    req.FirstName,
				LastName:     req.LastName,
				Email:        req.Email,
				PasswordHash: string(hashedPass),
				Role:         models.HandymanRole,
				PhoneNumber:  req.PhoneNumber,
				Img:          req.Img,
			},
			Type: models.HandymanType(req.Type),
		}

		//Creating the handyman user
		err = s.handymenRepo.CreateUser(user, tx)
		if err != nil {
			return errors.New("user not created: " + err.Error())
		}

		//Adding the handymen type
		err = s.handymenRepo.AddHandymenType(user, tx)
		if err != nil {
			return errors.New("user not created: " + err.Error())
		}

		// Create the session object
		token, err:= s.jwtService.GenerateJWT(user.UserID.String(), UserRole(user.Role))	//Generate JWT Token
		if err != nil {
			return errors.New("Unable to genrate JWT token: " + err.Error())
		}

		session:= &models.Session{
			UserID: user.UserID,
			TokenHash: token,
			Revoked: false,
		}

		// creating the session
		err = s.sessionRepository.CreateSession(session, tx)
		if err != nil {
			return errors.New("Session not created: " + err.Error())
		}

		return nil
	})
}


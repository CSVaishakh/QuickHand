package src

import (
	"errors"
	"crypto/sha256"
	"encoding/hex"

	"github.com/CSVaishakh/QuickHand/src/packages/db/models"
	repo "github.com/CSVaishakh/QuickHand/src/packages/db/repositories"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ErrEmailAldreadyExists = errors.New("email ardeady registered")
	ErrInvalidToken    = errors.New("invalid token")
	ErrSignOutFailed   = errors.New("sign out failed")
	ErrSessionNotFound = errors.New("session not found")
	ErrInvalidSession  = errors.New("invalid session")
	ErrSessionExpired  = errors.New("session expired")
	ErrDBFailed        = errors.New("session expired but databse failed")
	ErrInvalidCredentials = errors.New("invalid email or password")
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

func (s *AuthService) HandymanSignUp(req HandymanSignUpReq) (string, error) {
	var token string

	err := s.db.Transaction(func(tx *gorm.DB) error {

		userExists, err := s.handymenRepo.CheckByEmail(req.Email, tx)

		if err != nil {
			return err
		}

		if userExists {
			return ErrEmailAldreadyExists
		}

		hashedPass, err := bcrypt.GenerateFromPassword(
			[]byte(req.Password),
			bcrypt.DefaultCost,
		)

		if err != nil {
			return err
		}

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

		err = s.handymenRepo.CreateUser(user, tx)
		if err != nil {
			return err
		}

		err = s.handymenRepo.AddHandymenType(user, tx)
		if err != nil {
			return err
		}
		//generate JWT
		token, err = s.jwtService.GenerateJWT(user.UserID.String(), UserRole(user.Role))
		if err != nil {
			return err
		}

		hash := sha256.Sum256([]byte(token))
		tokenHash := hex.EncodeToString(hash[:])

		session := &models.Session{
			UserID:    user.UserID,
			TokenHash: tokenHash,
			Revoked:   false,
		}
		
		//add session to db
		err = s.sessionRepository.CreateSession(session, tx)
		if err != nil {
			return err
		}

		return  nil
	})

	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) HandymanSignIn(req SignInReq) (HandymanSignInRes, error) {
	//get User details
	user, err := s.handymenRepo.GetUser(req.Email, s.db)
	if err != nil {
		return HandymanSignInRes{}, err
	}

	if user == nil {
		return HandymanSignInRes{}, ErrInvalidCredentials
	}

	//verifry password
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(req.Password),
	) 
	if err != nil {
		return HandymanSignInRes{}, ErrInvalidCredentials
	}

	//generate JWT
	token, err := s.jwtService.GenerateJWT(user.UserID.String(), UserRole(user.Role))
	if err != nil {
		return HandymanSignInRes{}, err
	}
	hash := sha256.Sum256([]byte(token))
	tokenHash := hex.EncodeToString(hash[:])

	//genrate session object
	session := models.Session{
		UserID:    user.UserID,
		TokenHash: tokenHash,
		Revoked:   false,
	}

	//add session to db
	err = s.sessionRepository.CreateSession(&session, s.db)
	if err != nil {
		return HandymanSignInRes{}, err
	}

	return HandymanSignInRes{
		Token: token,
	}, nil
}

func (s *AuthService) ClientSignUp(req ClientSignUpReq) (string, error) {
	var token string
	err := s.db.Transaction(func(tx *gorm.DB) error {
		// Checking for exisitng user
		userExists, err := s.clientRepository.GetByEmail(req.Email, tx)

		if err != nil {
			return err
		}

		if userExists {
			return ErrEmailAldreadyExists
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
		user := &models.Client{
			User: models.User{
				FirstName:    req.FirstName,
				LastName:     req.LastName,
				Email:        req.Email,
				PasswordHash: string(hashedPass),
				Role:         models.ClientRole,
				PhoneNumber:  req.PhoneNumber,
				Img:          req.Img,
			},
		}

		//Creating the client user
		err = s.clientRepository.CreateUser(user, tx)
		if err != nil {
			return err
		}

		// Create the session object
		token, err := s.jwtService.GenerateJWT(user.UserID.String(), UserRole(user.Role)) //Generate JWT Token
		if err != nil {
			return err
		}

		hash := sha256.Sum256([]byte(token))
		tokenHash := hex.EncodeToString(hash[:])

		session := &models.Session{
			UserID:    user.UserID,
			TokenHash: tokenHash,
			Revoked:   false,
		}

		// creating the session
		err = s.sessionRepository.CreateSession(session, tx)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return token, nil 
}

func (s *AuthService)ClientSignIn(req SignInReq) (ClientSignInRes, error) {
	//get User details
	user, err := s.handymenRepo.GetUser(req.Email, s.db)
	if err != nil {
		return ClientSignInRes{}, err
	}

	if user == nil {
		return ClientSignInRes{}, ErrInvalidCredentials
	}

	//verifry password
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(req.Password),
	) 
	if err != nil {
		return ClientSignInRes{}, ErrInvalidCredentials
	}

	//generate JWT
	token, err := s.jwtService.GenerateJWT(user.UserID.String(), UserRole(user.Role))
	if err != nil {
		return ClientSignInRes{}, err
	}
	hash := sha256.Sum256([]byte(token))
	tokenHash := hex.EncodeToString(hash[:])

	//genrate session object
	session := models.Session{
		UserID:    user.UserID,
		TokenHash: tokenHash,
		Revoked:   false,
	}

	//add session to db
	err = s.sessionRepository.CreateSession(&session, s.db)
	if err != nil {
		return ClientSignInRes{}, err
	}

	return ClientSignInRes{
		Token: token,
	}, nil
}

func (s *AuthService) SignOut(token string) error {
	session, err := s.sessionRepository.RevokeSession(token)
	if err != nil {
		return err
	}

	if !session.Revoked {
		return ErrSignOutFailed
	}

	return nil
}

func (s *AuthService) VerifySession(req VerifySessionReq) (session *models.Session, err error) {
	// validate jwt
	claims, err := s.jwtService.VerifyJWT(req.Token)
	if err != nil {
		return nil, ErrInvalidToken
	}
	if claims == nil {
		return nil, ErrInvalidToken
	}

	// get session
	hash := sha256.Sum256([]byte(req.Token))
	tokenHash := hex.EncodeToString(hash[:])

	session, err = s.sessionRepository.GetSession(tokenHash)
	if err != nil {
		return nil, err
	}

	if session == nil {
		return nil, ErrSessionNotFound
	}

	// check revoked
	if session.Revoked {
		return nil, ErrInvalidSession
	}
	return session, nil
}
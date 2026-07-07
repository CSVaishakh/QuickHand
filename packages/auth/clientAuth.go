package auth

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/CSVaishakh/QuickHand/packages/db/models"
)

func (s *AuthService) ClientSignUp(req ClientSignUpReq) (ClientSignUpRes, error) {
	var token string
	var user models.Client
	err := s.db.Transaction(func(tx *gorm.DB) error {
		// Checking for existing user
		userExists, err := s.userRepo.WithTx(tx).CheckByEmail(req.Email)

		if err != nil {
			return err
		}

		if userExists {
			return ErrEmailAlreadyExists
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
		user = models.Client{
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
		err = s.userRepo.WithTx(tx).CreateUser(&user.User)
		if err != nil {
			return err
		}

		// Create the session object
		token, err = s.jwtService.GenerateJWT(user.UserID, UserRole(user.Role)) //Generate JWT Token
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
		err = s.sessionRepo.WithTx(tx).CreateSession(session)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return ClientSignUpRes{}, err
	}

	return ClientSignUpRes{
		UserID: user.UserID,
		FirstName: user.FirstName,
		Token: token,
		Role: UserRole(user.Role),
	}, nil
}

func (s *AuthService) ClientSignIn(req SignInReq) (ClientSignInRes, error) {
	//get User details
	user, err := s.clientRepo.GetByEmail(req.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ClientSignInRes{}, ErrInvalidCredentials
	}

	if err != nil {
		return ClientSignInRes{}, err
	}

	//verify password
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(req.Password),
	)
	if err != nil {
		return ClientSignInRes{}, ErrInvalidCredentials
	}

	//generate JWT
	token, err := s.jwtService.GenerateJWT(user.UserID, UserRole(user.Role))
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
	err = s.sessionRepo.CreateSession(&session)
	if err != nil {
		return ClientSignInRes{}, err
	}

	return ClientSignInRes{
		UserID: user.UserID,
		FirstName: user.FirstName,
		Token: token,
		Role: UserRole(user.Role),
	}, nil
}
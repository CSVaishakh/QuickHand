package src

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"crypto/sha256"
	"encoding/hex"

	"github.com/CSVaishakh/QuickHand/src/packages/db/models"
)

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

		return nil
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

	//verify password
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

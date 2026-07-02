	package src

	import (
		"golang.org/x/crypto/bcrypt"
		"gorm.io/gorm"

		"crypto/sha256"
		"encoding/hex"
		"errors"

		"github.com/CSVaishakh/QuickHand/src/packages/db/models"
	)

	func (s *AuthService) HandymanSignUp(req HandymanSignUpReq) (HandymanSignUpRes, error) {
		var token string
		var user models.Handyman

		err := s.db.Transaction(func(tx *gorm.DB) error {

			userExists, err := s.userRepo.CheckByEmail(req.Email, tx)

			if err != nil {
				return err
			}

			if userExists {
				return ErrEmailAlreadyExists
			}

			hashedPass, err := bcrypt.GenerateFromPassword(
				[]byte(req.Password),
				bcrypt.DefaultCost,
			)

			if err != nil {
				return err
			}

			user = models.Handyman{
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

			err = s.userRepo.CreateUser(&user.User, tx)
			if err != nil {
				return err
			}

			err = s.handymenRepo.AddHandymenType(&user, tx)
			if err != nil {
				return err
			}
			//generate JWT
			token, err = s.jwtService.GenerateJWT(user.UserID, UserRole(user.Role))
			if err != nil {
				return err
			}

			hash := sha256.Sum256([]byte(token))
			tokenHash := hex.EncodeToString(hash[:])

			session := models.Session{
				UserID:    user.UserID,
				TokenHash: tokenHash,
				Revoked:   false,
			}

			//add session to db
			err = s.sessionRepo.CreateSession(&session, tx)
			if err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			return HandymanSignUpRes{}, err
		}

		return HandymanSignUpRes{
			UserID: user.UserID,
			FirstName: user.FirstName,
			Token: token,
			Role: UserRole(user.Role),
			Type: HandymanType(user.Type),
		}, nil
	}

	func (s *AuthService) HandymanSignIn(req SignInReq) (HandymanSignInRes, error) {
		//get User details
		user, err := s.handymenRepo.GetByEmail(req.Email, s.db)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return HandymanSignInRes{}, ErrInvalidCredentials
		}

		if err != nil {
			return HandymanSignInRes{}, err
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
		token, err := s.jwtService.GenerateJWT(user.UserID, UserRole(user.Role))
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
		err = s.sessionRepo.CreateSession(&session, s.db)
		if err != nil {
			return HandymanSignInRes{}, err
		}

		return HandymanSignInRes{
			UserID: user.UserID,
			FirstName: user.FirstName,
			Token: token,
			Role: UserRole(user.Role),
			Type: HandymanType(user.Type),
		}, nil
	}

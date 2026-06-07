package src

import(
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"

	"crypto/sha256"
    "encoding/hex"

	"github.com/CSVaishakh/QuickHand/src/packages/db/models"
)

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



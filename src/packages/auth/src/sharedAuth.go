package src

import(
	"crypto/sha256"
    "encoding/hex"

	"github.com/CSVaishakh/QuickHand/src/packages/db/models"
)

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


func (s *AuthService) VerifySession(req VerifySessionReq) (session *models.Session, claims *Claims, err error) {
	// validate jwt
	claims, err = s.jwtService.VerifyJWT(req.Token)
	if err != nil {
		return nil, nil, ErrInvalidToken
	}
	if claims == nil {
		return nil, nil, ErrInvalidToken
	}

	// get session
	hash := sha256.Sum256([]byte(req.Token))
	tokenHash := hex.EncodeToString(hash[:])

	session, err = s.sessionRepository.GetSession(tokenHash)
	if err != nil {
		return nil, nil, err
	}

	if session == nil {
		return nil, nil, ErrSessionNotFound
	}

	// check revoked
	if session.Revoked {
		return nil, nil, ErrInvalidSession
	}
	return session, claims, nil
}

// func (s *AuthService) ForgotPassowrd(req ForgotPasswordReq) error{
// 	//verify user exisits
// 	userExists, err := s.handymenRepo.CheckByEmail(req.Email, s.db)
// 	if err != nil {
// 		return err
// 	}

// 	if userExists {
// 		return ErrUserDoesNotExist
// 	}

	

// 	return nil
// }
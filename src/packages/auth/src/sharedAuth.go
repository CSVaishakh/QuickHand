package src

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/CSVaishakh/QuickHand/src/packages/db/models"
	"golang.org/x/crypto/bcrypt"
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
	if errors.Is(err, ErrInvalidToken) {
		return nil, nil, ErrInvalidToken
	}
	if err != nil {
		return nil, nil, err
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

func (s *AuthService) ForgotPassowrd(req ForgotPasswordReq) (string, error) {
	//verify user exisits
	userExists, err := s.handymenRepo.CheckByEmail(req.Email, s.db)
	if err != nil {
		return "", err
	}

	if userExists {
		return "", ErrUserDoesNotExist
	}

	//genrate otp and token
	otp, err := GenerateOTP(12)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256([]byte(otp))
	OtpHash := hex.EncodeToString(hash[:])

	token, err := s.jwtService.GenerateOTP_JWT(OtpHash, req.Email)
	if err != nil {
		return "", err
	}

	//send email with otp to the email id
	err = SendEmail(req.Email, otp)
	if err != nil {
		return "", err
	}

	return token, err
}

func (s *AuthService) VerifyOTP (req OtpVerificationReq) (string,error) {
	claims, err := s.jwtService.VerifyOTP_JWT(req.Token)
	if errors.Is(err, ErrInvalidToken) {
		return "",ErrInvalidToken
	}
	if err != nil {
		return "",err
	}

	hash := sha256.Sum256([]byte(req.Otp))
	submittedHash := hex.EncodeToString(hash[:])

	if submittedHash != claims.OtpHash {
		return "",ErrInvalidOTP
	}

	token, err := s.jwtService.GenerateResetJWT(claims.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) ResetPassword(req ResetPasswordReq) error {
	claims, err := s.jwtService.VerifyResetJWT(req.Token)
	if errors.Is(err, ErrInvalidToken) {
		return ErrInvalidToken
	}
	if err != nil {
		return err
	}

	hashedPass, err := bcrypt.GenerateFromPassword(
		[]byte(req.NewPassword),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	result := s.db.
		Model(&models.User{}).
		Where("email = ?", claims.Email).
		Update("password_hash", string(hashedPass))

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrUserDoesNotExist
	}

	return nil
}
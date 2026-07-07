package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"github.com/CSVaishakh/QuickHand/packages/db/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *AuthService) SignOut(token string) error {
	hash := sha256.Sum256([]byte(token))
	tokenHash := hex.EncodeToString(hash[:])

	session, err := s.sessionRepo.RevokeSession(tokenHash)

	if err != nil {
		if err.Error() == "session not found" {
			return ErrSessionNotFound
		}
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

	session, err = s.sessionRepo.GetSession(tokenHash)
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

	if claims.ExpiresAt.Before(time.Now()) {
		return nil, nil, ErrSessionExpired
	}
	
	return session, claims, nil
}

func (s *AuthService) ForgotPassword(req ForgotPasswordReq) (string, error) {
	//verify user exisits
	userExists, err := s.userRepo.CheckByEmail(req.Email)
	if err != nil {
		return "", err
	}

	if !userExists {
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

	err = s.userRepo.ResetPassword(hashedPass, claims.Email)

	if errors.Is(err, gorm.ErrRecordNotFound){
		return ErrUserDoesNotExist
	}

	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) GetSession(req GetSessionReq) (GetSessionRes, error) {
	claims, err := s.jwtService.VerifyJWT(req.Token)
	if err != nil {
		return GetSessionRes{}, err
	}

	hash := sha256.Sum256([]byte(req.Token))
	tokenHash := hex.EncodeToString(hash[:])

	session, err := s.sessionRepo.GetSession(tokenHash)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return GetSessionRes{}, ErrSessionNotFound
	}

	if err != nil {
		return GetSessionRes{}, err
	}

	if session.Revoked {
		return GetSessionRes{}, ErrSessionNotFound
	}

	res := GetSessionRes{
		SessionId: session.SessionID,
		Revoked:   session.Revoked,
		CreatedAt: session.CreatedAt,
	}

	switch claims.Role {
		case ClientRole:
			user, err := s.clientRepo.GetByUserID(
				session.UserID.String(),
			)
	
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return GetSessionRes{}, ErrInvalidCredentials
			}
	
			if err != nil {
				return GetSessionRes{}, err
			}
	
			res.UserId = user.UserID
			res.FirstName = user.FirstName
			res.Email = user.Email
			res.Role = UserRole(user.Role)
	
		case HandymanRole:
			user, err := s.handymenRepo.GetByUserID(
				session.UserID.String(),
			)
	
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return GetSessionRes{}, ErrInvalidCredentials
			}
	
			if err != nil {
				return GetSessionRes{}, err
			}
	
			res.UserId = user.UserID
			res.FirstName = user.FirstName
			res.Email = user.Email
			res.Role = UserRole(user.Role)
	
			ht := HandymanType(user.Type)
			res.Type = &ht
	
		default:
			return GetSessionRes{}, ErrInvalidCredentials
	}

	return res, nil
}
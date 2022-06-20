package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/forever-eight/sport-point-backend.git/internal/app/ds"
)

const (
	salt       = "skdjnbfgdhuw4rbdjskjsn89"
	tokenTTL   = 12 * time.Hour
	signingKey = "jdjdkdoen3432nskd9"
)

func (s *Service) CreateUser(user *ds.UserInput) (*ds.UserOutput, error) {
	err := user.Validate()
	if err != nil {
		return nil, fmt.Errorf("service err: CreateUser: %s", err)
	}
	u := user.ToDB()
	u.Password = generatePasswordHash(u.Password)

	id, err := s.repos.CreateUser(u)
	if err != nil {
		return nil, fmt.Errorf("service err: CreateUser: %s", err)
	}
	output := u.ToOutput(id)

	return output, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

type tokenClaims struct {
	jwt.StandardClaims
	UserID uint32 `json:"user_id"`
}

func (s *Service) GenerateToken(email, password string) (string, error) {
	user, err := s.repos.GetUser(email, generatePasswordHash(password))
	if err != nil {
		return "", fmt.Errorf("service err: GenerateToken: %s", err)
	}
	tokenSigner := jwt.NewWithClaims(jwt.SigningMethodHS512, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(), // время на 12 часов больше, чем сейчас
			IssuedAt:  time.Now().Unix(),
		},
		UserID: user.ID,
	})

	token, err := tokenSigner.SignedString([]byte(signingKey))
	if err != nil {
		return "", fmt.Errorf("sign error: %w", err)
	}

	return token, nil
}

func (s *Service) ParseToken(accessToken string) (uint32, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("accessToken claims arent of right type")
	}
	return claims.UserID, nil
}

package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/forever-eight/sport-point-backend.git/internal/app/ds"
)

const (
	salt = "skdjnbfgdhuw4rbdjskjsn89"
)

func (s *Service) CreateUser(user *ds.UserInput) (*ds.UserOutput, error) {
	err := user.Validate()
	if err != nil {
		return nil, fmt.Errorf("service err: CreateUser: %s", err)
	}
	u := user.ToDB()
	u.Password = s.generatePasswordHash(u.Password)

	id, err := s.repos.CreateUser(u)
	if err != nil {
		return nil, fmt.Errorf("service err: CreateUser: %s", err)
	}
	output := u.ToOutput(id)

	return output, nil
}

func (s *Service) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

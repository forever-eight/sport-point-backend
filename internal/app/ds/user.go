package ds

import (
	"fmt"

	"github.com/go-email-validator/go-email-validator/pkg/ev"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evmail"
)

type UserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (i *UserInput) Validate() error {
	if i == nil {
		return fmt.Errorf("empty input")

	}
	var v = ev.NewSyntaxValidator().Validate(ev.NewInput(evmail.FromString(i.Email)))
	if !v.IsValid() {
		return fmt.Errorf(" email is invalid")
	}
	if i.Name == "" {
		return fmt.Errorf("name is invalid")
	}
	if i.Password == "" {
		return fmt.Errorf("password is invalid")
	}

	return nil
}

func (i *UserInput) ToDB() *User {
	return &User{
		Name:     i.Name,
		Email:    i.Email,
		Password: i.Password,
	}
}

type User struct {
	ID       uint32 `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

func (u *User) ToOutput(id uint32) *UserOutput {
	return &UserOutput{
		ID:    id,
		Name:  u.Name,
		Email: u.Email,
	}
}

type UserOutput struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

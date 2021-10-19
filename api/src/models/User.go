package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID        uint64    `json:"id,omitempty`
	Name      string    `json:name,omitempty`
	Nick      string    `json:nick,omitempty`
	Email     string    `json:email,omitempty`
	Password  string    `json:password,omitempty`
	CreatedOn time.Time `json:createdOn,omitempty`
}

func (user *User) Prepare(opetarion string) error {
	if erro := user.validation(opetarion); erro != nil {
		return erro
	}
	user.format()
	return nil
}

func (user *User) validation(opetarion string) error {
	if user.Name == "" {
		return errors.New("name is required and cannot be emptys")
	}

	if user.Nick == "" {
		return errors.New("nick is required and cannot be empty")
	}

	if user.Email == "" {
		return errors.New("email is required and cannot be empty")
	}

	if opetarion == "create" && user.Password == "" {
		return errors.New("password is required and cannot be empty")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}

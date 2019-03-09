package main

import (
	"errors"
	"io"

	"github.com/emersion/go-smtp"
	"github.com/dusankasan/parsemail"
)

type Backend struct{
	Username string
	Password string
	Output OutputPlugin
}

func (bkd *Backend) Login(username, password string) (smtp.User, error) {
	if username != bkd.Username || password != bkd.Password {
		return nil, errors.New("Invalid username or password")
	}
	return &User{Output: bkd.Output}, nil
}

func (bkd *Backend) AnonymousLogin() (smtp.User, error) {
	return nil, smtp.ErrAuthRequired
}

type User struct{
	Output OutputPlugin
}

func (u *User) Send(from string, to []string, r io.Reader) error {
	email, err := parsemail.Parse(r)
	if err != nil {
		return err
	}

	u.Output.Send(email.TextBody)
	return nil
}

func (u *User) Logout() error {
	return nil
}

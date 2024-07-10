package service

import "errors"

type User struct {
    Id int
    Username string
    Email string
    Password string
}

func Authenticate(email, password string) (User, bool) {
    return User{}, false
}

type Register struct {
    Username string
    Email    string
    Password string
    Name string
}

func (r *Register) Register() error {
    return errors.New("not implemented")
}


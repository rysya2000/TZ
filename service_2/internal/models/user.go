package models

import "regexp"

type User struct {
	Email    string `json:"email" example:"abcd@gmail.com"`
	Salt     string `json:"salt,omitempty" example:"Ac3428x5L3xq"`
	Password string `json:"password" example:"qwerty"`
}

func (u *User) ValidateEmail() bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	return re.MatchString(u.Email)
}

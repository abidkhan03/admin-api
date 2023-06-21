package request

import "github.com/spongeling/admin-api/internal/errors"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) Validate() error {
	if u.Username == "" {
		return errors.BadRequest("username is required")
	}
	if u.Password == "" {
		return errors.BadRequest("password is required")
	}
	return nil
}

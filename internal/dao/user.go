package dao

import "github.com/spongeling/admin-api/internal/errors"

type User struct {
	Id       uint64 `gorm:"column:id"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}

func (u *User) GetID() uint64 {
	return u.Id
}

func (*User) TableName() string {
	return "user"
}

func (u *User) Validate() error {
	if u.Email == "" {
		return errors.BadRequest("email is required")
	}
	if u.Password == "" {
		return errors.BadRequest("password is required")
	}
	return nil
}

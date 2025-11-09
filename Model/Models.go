package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//validation
type UserInfo struct {
	Username string `binding:"required,min=6"`
	Email    string `gorm:"unique" binding:"required,email"`
	Password string `binding:"required,min=6,max=16"`
}

//database struct
type User struct {
	ID       uint `gorm:"primaryKey"`
	Username string
	Email    string `gorm:"unique"`
	Password string
}


//beforehook for hashing password
func (u *User) BeforeCreate(db *gorm.DB) (err error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashed)
	return nil
}

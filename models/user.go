package models

import (
	"final-project/utils/token"
	"html"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Username  string    `json:"username" gorm:"not null;unique"`
	Email     string    `json:"email" gorm:"not null;unique"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Review	  []Review	`json:"-"`
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username, password string, db *gorm.DB) (string, error) {
	var err error

	usr := User{}

	err = db.Model(User{}).Where("username = ?", username).Take(&usr).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, usr.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(usr.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (usr *User) SaveUser(db *gorm.DB) (*User, error) {
	hashedPassword, errPassword := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
	if errPassword != nil {
		return &User{}, errPassword
	}

	usr.Password = string(hashedPassword)
	usr.Username = html.EscapeString(strings.TrimSpace(usr.Username))

	var err error = db.Create(&usr).Error
	if err != nil {
		return &User{}, err
	}

	return usr, nil
}


func (usr *User) SaveNewPassword(db *gorm.DB) (*User, error) {
	hashedPassword, errPassword := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
	if errPassword != nil {
		return &User{}, errPassword
	}

	usr.Password = string(hashedPassword)

	var err error = db.Save(&usr).Error
	if err != nil {
		return &User{}, err
	}

	return usr, nil
}

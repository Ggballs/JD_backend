package mdDef

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type UserBasic struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null; index; column:name"`
	PassWord string `json:"password" gorm:"column:password"`
}

func (u *UserBasic) Compare(password string) error {
	if err := Compare(u.PassWord, password); err != nil {
		log.Println("user password err " + err.Error())
		return err
	}
	return nil
}

func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *UserBasic) Encrypt() error {
	password, err := Encrypt(u.PassWord)
	if err != nil {
		log.Println("encrypt error: " + err.Error())
		return err
	}
	u.PassWord = password
	return nil
}

func Encrypt(source string) (string, error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashBytes), nil
}

func (u *UserBasic) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

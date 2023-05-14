package mdDef

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type ViewedJob struct {
	JobId      string `json:"job_id"`
	ViewedTime string `json:"viewed_time"`
}

type CollectedJob struct {
	JobId         string `json:"job_id"`
	CollectedTime string `json:"collected_time"`
}
type UserBasic struct {
	gorm.Model
	UserId        string `json:"id" gorm:"not null; index; column:user_id"`
	Name          string `json:"name" gorm:"column:name"`
	PassWord      string `json:"password" gorm:"column:password"`
	ViewedJobs    []byte `gorm:"type:json"`
	CollectedJobs []byte `gorm:"type:json"`
}

func (UserBasic) TableName() string {
	return "user_basics"
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

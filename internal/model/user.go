package model

import (
	"errors"
	"os"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName          string        `json:"user_name"`
	Email             string        `gorm:"uniq;uniqueIndex;not null" json:"email"`
	Password          string        `json:"password,omitempty" gorm:"-"`
	EncryptedPassword string        `json:"-"` // stored hash
	ContactNo         string        `json:"contact_no"`
	Groups            []Group       `json:"groups" gorm:"many2many:user_groups"`
	PaidTransactions  []Transaction `json:"paid_transactions" gorm:"foreignKey:PaidById"`
	Transactions      []Transaction `json:"transactions" gorm:"many2many:user_transactions"`
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

var emailRegex = regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,}$`)

var (
	upperRegex   = regexp.MustCompile(`[A-Z]`)
	lowerRegex   = regexp.MustCompile(`[a-z]`)
	specialRegex = regexp.MustCompile(`[^A-Za-z0-9]`)
)

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.Email == "" || !emailRegex.MatchString(u.Email) {
		return errors.New("invalid email")
	}
	if u.Password == "" {
		return errors.New("password is required")
	}
	if len(u.Password) < 8 || !upperRegex.MatchString(u.Password) || !lowerRegex.MatchString(u.Password) || !specialRegex.MatchString(u.Password) {
		return errors.New("password must be at least 8 characters and include upper, lower, and special character")
	}
	hash, err := hashPassword(u.Password)
	if err != nil {
		return err
	}
	u.EncryptedPassword = hash
	u.Password = ""
	return nil
}

func (u *User) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password))
	return err == nil
}

func (u *User) CreateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    u.ID,
		"email": u.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	var secretKey = []byte(os.Getenv("SECRET_KEY"))
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

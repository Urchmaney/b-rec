package authenticator

import (
	"golang.org/x/crypto/bcrypt"
  "fmt"
)

type AuthenticationService struct {}

func(_ AuthenticationService) HashPassword(password string) (string, error) {
  bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
  return string(bytes), err
}

func (_ AuthenticationService) CheckPasswordHash(password, hash string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
  return err == nil
}


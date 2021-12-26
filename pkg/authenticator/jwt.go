package authenticator

import (
  "github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
  "fmt"
  "time"
)

type AuthenticationService struct {}

func(_ AuthenticationService) HashPassword(password string) (string, error) {
  bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
  return string(bytes), err
}

func (_ AuthenticationService) CheckPasswordHash(password, hash string) bool {
  fmt.Println(password)
  fmt.Println(hash)
  err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
  fmt.Println(err)
  return err == nil
}

func (_ AuthenticationService) GenerateJWT(account_id int64) (string, error) {
  var signingKey = []byte("5ToDQbRoJQ6UVclAgOLK17NXrUGlkKFtfn3ulwHkYxIkjZ4JrvSfTHN5X9Pd")
  token := jwt.New(jwt.SigningMethodHS256)
  claims := token.Claims.(jwt.MapClaims)
  claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
  claims["id"] = account_id

  tokenString, err := token.SignedString(signingKey)
  if err != nil {
		fmt.Errorf("Something Went Wrong Generating Token: %s", err.Error())
		return "", err
	}
  return tokenString, nil
}


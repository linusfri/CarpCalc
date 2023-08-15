package auth

import (
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	userModel "github.com/linusfri/calc-api/models/user"
	"golang.org/x/crypto/bcrypt"
)

type AuthResponse struct {
	AccessToken string
	TokeType    string
}

func PasswordValid(hashedUserPassword []byte, userPassword []byte) (success bool) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedUserPassword), []byte(userPassword)); err != nil {
		return false
	}

	return true
}

func GenerateJWT(user *userModel.User) (string, error) {
	var (
		key []byte
		t   *jwt.Token
		s   string
		err error
	)

	key = []byte(os.Getenv("AUTH_SECRET_KEY"))
	t = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Iss": "calc-api",
		"Sub": user.UserName,
		"Id":  user.ID,
		"Exp": time.Now().Add(10 * time.Minute),
	})
	s, err = t.SignedString(key)

	if err != nil {
		return "", err
	}

	return s, err
}

func VerifyJWT(authHeader string) bool {
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("AUTH_SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		return false
	}

	return true
}

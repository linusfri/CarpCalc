package authController

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/linusfri/calc-api/db"
	"github.com/linusfri/calc-api/helper"
	authModel "github.com/linusfri/calc-api/models/auth"
	userModel "github.com/linusfri/calc-api/models/user"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var requestUser *userModel.User
	var user *userModel.User

	c, err := db.Connect()
	if err != nil {
		helper.HandleErr(err)
	}

	if err := json.NewDecoder(r.Body).Decode(&requestUser); err != nil {
		helper.HandleErr(err)
		return
	}
	if err := c.Gorm.Where("Email = ?", requestUser.Email).First(&user).Error; err != nil {
		helper.HandleErr(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !authModel.PasswordValid([]byte(user.Password), []byte(requestUser.Password)) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, err := authModel.GenerateJWT(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong during authentication"))
	}

	cookie := http.Cookie{
		Name:     "jwt",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
	w.Write([]byte("Logged in"))
}

package userController

import (
	"encoding/json"
	"net/http"

	"github.com/linusfri/calc-api/db"
	"github.com/linusfri/calc-api/helper"
	userModel "github.com/linusfri/calc-api/models/user"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/clause"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user *userModel.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helper.HandleErr(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		helper.HandleErr(err)
		return
	}

	user.Password = string(pass)

	c, err := db.Connect()
	if err != nil {
		helper.HandleErr(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if UserExists(user) {
		w.WriteHeader(http.StatusConflict)
		return
	}

	c.Gorm.Clauses(clause.OnConflict{DoNothing: true}).Create(&user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []userModel.User

	c, err := db.Connect()

	if err != nil {
		helper.HandleErr(err)
		return
	}

	c.Gorm.Find(&users)

	res, err := json.Marshal(&users)
	if err != nil {
		helper.HandleErr(err)
		return
	}
	w.Write(res)
}

func UserExists(user *userModel.User) bool {
	var foundUser userModel.User

	c, err := db.Connect()

	if err != nil {
		helper.HandleErr(err)
		panic(err)
	}

	if err := c.Gorm.Where("user_name = ?", user.UserName).First(&foundUser).Error; err != nil {
		return false
	}
	return true
}

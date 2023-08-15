package user

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"not null;unique"`
	Email    string `gorm:"type:varchar(100);unique_index"`
	Password string `gorm:"not null"`
}

func (user *User) SayFuck() {
	fmt.Println("Fuck")
}

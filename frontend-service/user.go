package main

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}

// func createUser(username, password string) error {
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return err
// 	}

// 	user := User{
// 		Username: username,
// 		Password: string(hashedPassword),
// 	}

// 	return db.Create(&user).Error
// }

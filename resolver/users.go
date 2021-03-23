package resolver

import (
	"log"
	"mvreview-heroku/config"
	logger "mvreview-heroku/log"
	"mvreview-heroku/models"
	"time"

	"github.com/graphql-go/graphql"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// func GetAllUser
func GetAllUser(params graphql.ResolveParams) (interface{}, error) {
	dbPG := config.Connection()
	// Users struct
	type Users struct {
		ID        uint           `gorm:"primarykey" json:"id"`
		Email     string         `gorm:"primarykey" json:"email"`
		FullName  string         `json:"fullName"`
		Role      string         `json:"role"`
		CreatedAt time.Time      `json:"created_at"`
		UpdatedAt time.Time      `json:"updated_at"`
		DeletedAt gorm.DeletedAt `gorm:"index"`
	}
	var (
		users []Users
	)
	dbPG.Find(&users)
	return users, nil //return users data response
}

// func GetUser
func GetUser(params graphql.ResolveParams) (interface{}, error) {
	email, ok := params.Args["email"].(string) // Get email from params arguments
	dbPG := config.Connection()
	// Users struct
	type Users struct {
		ID        uint           `gorm:"primarykey" json:"id"`
		Email     string         `gorm:"primarykey" json:"email"`
		FullName  string         `json:"fullName"`
		Role      string         `json:"role"`
		CreatedAt time.Time      `json:"created_at"`
		UpdatedAt time.Time      `json:"updated_at"`
		DeletedAt gorm.DeletedAt `gorm:"index"`
	}
	var (
		users Users
	)

	if ok {
		dbPG.Where("email = ?", email).First(&users)
		return users, nil //return users data response
	}
	return nil, nil
}

// func Register
func Register(params graphql.ResolveParams) (interface{}, error) {
	email, checkEmail := params.Args["email"].(string)          // Get email from params arguments
	password, checkPassword := params.Args["password"].(string) // Get password from params arguments
	fullName, checkfullName := params.Args["fullName"].(string) // Get fullName from params arguments
	role, checkRole := params.Args["role"].(string)             // Get role from params arguments
	dbPG := config.Connection()
	var (
		users models.Users
	)

	encrypt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) // encrypt password
	// chec if erros while encrypting
	if err != nil {
		log.Println(err)
		logger.Sentry(err)
	}
	users.Email = email
	users.Password = string(encrypt)
	users.FullName = fullName
	users.Role = role

	if checkEmail && checkPassword && checkfullName && checkRole {
		dbPG.Create(&users)
		return users, nil //return Users data response
	}
	return nil, nil
}

package user

import (
	conn "gobanks/config"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var db = conn.DBConn()

type User struct {
	gorm.Model
	Fullname      string
	Email         string
	Password      string
	AccountNumber uuid.UUID
	Balance       int
}

type Transfer struct {
	gorm.Model
	UserID   uint
	Receiver string
	Amount   int
	User     User
}

//func GetUser(email string) User {
//	return nil
//}

func DeleteUser(email string) bool {
	return true
}

func GetAllUser() []User {
	var user []User

	db.Find(&user)

	return user
}

func NewUser(fullname, email, password string) User {
	var user User

	user.Fullname = fullname
	user.Email = email
	user.Password = password
	user.AccountNumber = uuid.New()
	user.Balance = 0

	db.Create(&user)

	return user
}

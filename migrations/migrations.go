package migrations

import (
	"github.com/mytrix-technology/go-banking/helpers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Username 	string
	Email 		string
	Password 	string
}

type Account struct {
	gorm.Model
	Type 		string
	Name 		string
	Balance 	uint
	UserID 		uint
}

func connDB() *gorm.DB  {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=go-banking password=postgres sslmode=disable")
	helper.HandleErr(err)
	return db
}

func createAccount()  {
	db := connDB()

	users := [2]User{
		{
			Username: "Yudhis",
			Email: "yudhis@yopmail.com",
		},
		{
			Username: "Aronggo",
			Email: "aronggo@yopmail.com",
		},
	}

	for i := 0; i < len(users); i++ {
		generatePassword := helper.HashAndSalt([]byte(users[i].Username))
		user := User{Username:users[i].Username, Email:users[i].Email, Password:generatePassword}
		db.Create(&user)

		account := Account{
			Type: "Daily Account",
			Name: string(users[i].Username + "'s" + " account"),
			Balance: uint(10000 * int(i+1)),
			UserID: user.ID,
		}
		db.Create(&account)

	}
	defer db.Close()
}

func Migrate()  {
	db := connDB()
	db.AutoMigrate(&User{}, &Account{})
	defer db.Close()

	createAccount()
}


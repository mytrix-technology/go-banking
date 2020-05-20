package vulnerableDB

import (
	"database/sql"
	"fmt"
	helper "github.com/mytrix-technology/go-banking/helpers"
	_ "github.com/lib/pq"
)

type User struct {
	ID int
	Username string
	Email string
	Accounts []Account
}

type Account struct {
	ID int
	Name string
	Balance int
}

func connDB() *sql.DB  {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=go-banking password=postgres sslmode=disable")
	helper.HandleErr(err)
	return db
}

func dbCall(query string) *sql.Rows  {
	db := connDB()

	call, err := db.Query(query)
	helper.HandleErr(err)

	return call
}

func VulnerableLogin(username string, pass string) []User  {
	password := helper.HashOnlyVulnerable([]byte(pass))
	result := dbCall("SELECT id, username, email FROM users x WHERE username='" + username + "' AND password='" + password + "'")
	var users []User

	for result.Next() {
		var user User
		err := result.Scan(&user.ID, &user.Username, &user.Email)
		helper.HandleErr(err)
		accounts := dbCall("SELECT id, name, balance FROM accounts x WHERE user_id=" + fmt.Sprint(user.ID) + "")
		var userAccounts []Account

		for accounts.Next() {
			var account Account
			err := accounts.Scan(&account.ID, &account.Name, &account.Balance)
			helper.HandleErr(err)
			userAccounts = append(userAccounts, account)
		}

		user.Accounts = userAccounts
		users = append(users, user)
	}
	defer result.Close()

	return users
}

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID        string `json:"id"`
	Citizenid string `json:"citizenid"`
	Title     string `json:"title"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
}

func main() {

	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	db, err := sql.Open("mysql", "root:P@ssw0rd@tcp(127.0.0.1:3306)/go")

	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT * FROM userinfo")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var user User

		err = results.Scan(&user.ID, &user.Citizenid, &user.Title, &user.Name, &user.Gender)
		if err != nil {
			panic(err.Error())
		}

		fmt.Printf(user.Citizenid)
		fmt.Printf(" : ")
		fmt.Printf(user.Name)
	}

}

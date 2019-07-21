package models

import (
	"database/sql"
	entities "rest-api/rest-api/entities"

	_ "github.com/go-sql-driver/mysql"
)

type PersonModel struct {
	Db *sql.DB
}

func (personModel PersonModel) FindAll() (person []entities.Person, err error) {
	rows, err := personModel.Db.Query("SELECT * FROM userinfo")

	if err != nil {
		return nil, err
	} else {
		var persons []entities.Person
		for rows.Next() {
			var id int
			var citizenid string
			var title string
			var name string
			var gender string

			err2 := rows.Scan(&id, &citizenid, &title, &name, &gender)

			if err2 != nil {
				return nil, err
			} else {
				person := entities.Person{
					ID:        id,
					CitizenID: citizenid,
					Title:     title,
					Name:      name,
					Gender:    gender,
				}
				persons = append(persons, person)
			}
		}
		return persons, nil
	}
}

package entities

import "fmt"

type Person struct {
	ID        int    `json:"id"`
	CitizenID string `json:"citizenid"`
	Title     string `json:"title"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
}

func (person Person) ToString() string {
	return fmt.Sprintf("ID: %d\nCitizen Id: %s\nTitle: %s\n Name: %s\nGender: %s",
		person.ID, person.CitizenID, person.Title, person.Name, person.Gender)
}

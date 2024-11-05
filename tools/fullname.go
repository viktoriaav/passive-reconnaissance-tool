package tools

import (
	"database/sql"
	"fmt"
)

// Define the User structure to hold the database result
type User struct {
	FirstName   string
	LastName    string
	Address     string
	PostalCode  string
	City        string
	PhoneNumber string
}

// Function to Process full name and fetch user data from the database
func ProcessFullName(db *sql.DB, firstName, lastName string) (string, error) {
	var user User
	query := `SELECT first_name, last_name, address, postal_code, city, phone_number 
	          FROM user_info WHERE first_name = ? AND last_name = ?`
	err := db.QueryRow(query, firstName, lastName).Scan(&user.FirstName, &user.LastName, &user.Address, &user.PostalCode, &user.City, &user.PhoneNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("no user found with the given full name")
		}
		return "", err
	}
	result := fmt.Sprintf("First name: %s\nLast name: %s\nAddress: %s\n%s %s\nNumber: %s", user.FirstName, user.LastName, user.Address, user.PostalCode, user.City, user.PhoneNumber)
	return result, nil
}

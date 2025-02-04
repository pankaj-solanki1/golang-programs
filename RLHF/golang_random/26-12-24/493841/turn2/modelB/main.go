package main  

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/pkg/errors"
)

// Custom structs to represent the data
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

// Custom validation function for User struct
func (u *User) Validate() error {
	if u.Username == "" {
		return errors.New("username is required")
	}
	if !strings.Contains(u.Email, "@") {
		return errors.New("invalid email format")
	}
	if u.Age <= 0 {
		return errors.New("age must be a positive integer")
	}
	return nil
}

func migrateDataFromCSVToJSON(csvFilePath, jsonFilePath string) error {
	// Read CSV data from file
	file, err := os.Open(csvFilePath)
	if err != nil {
		return errors.Wrap(err, "failed to open CSV file")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return errors.Wrap(err, "failed to read CSV data")
	}

	// Validate CSV header and row count
	expectedHeader := []string{"id", "username", "email", "age"}
	if len(records) == 0 {
		return errors.New("CSV file is empty")
	}
	if len(records[0]) != len(expectedHeader) {
		return errors.Errorf("CSV file header doesn't match the expected format. Expected header: %s", expectedHeader)
	}
	for i, h := range records[0] {
		if h != expectedHeader[i] {
			return errors.Errorf("CSV file header field %d doesn't match the expected format. Expected: %s, Got: %s", i+1, expectedHeader[i], h)
		}
	}

	// Parse CSV data into slice of User structs and validate each row
	var users []*User
	for _, record := range records[1:] {
		user := &User{}
		if err := user.parseAndValidateCSVRecord(record); err != nil {
			log.Printf("Warning: Ignoring invalid record: %v\n", record)
			continue
		}
		users = append(users, user)
	}

	// Write JSON data to file
	jsonData, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return errors.Wrap(err, "failed to marshal JSON data")
	}

	err = ioutil.WriteFile(jsonFilePath, jsonData, 0644)
	if err != nil {
		return errors.Wrap(err, "failed to write JSON file")
	}

	return nil
}

func (u *User) parseAndValidateCSVRecord(record []string) error {
	// We expect a record with exactly 4 fields: id, username, email, age
	if len(record) != 4 {
		return errors.Errorf("expected 4 fields in record, got: %d", len(record))
	}
	var err error
	u.ID, err = parseInt(record[0])
	if err != nil {
		return errors.Wrapf(err, "failed to parse field 'id'")
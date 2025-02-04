package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func GetUserData() (user *User, age int, err error) {
	// Some complex logic that might fail
	// ...

	// If everything goes well
	user = &User{Name: "Alice", Age: 25}
	age = 30
	return

	// If an error occurs
	// return nil, 0, errors.New("Failed to fetch user data")
}

func main() {
	user, age, err := GetUserData()
	if err != nil {
		// Handle error
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("User:", user.Name)
	fmt.Println("Age:", user.Age)
	fmt.Println("Additional Age:", age)
}

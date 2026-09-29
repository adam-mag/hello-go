package main

import "fmt"

type User struct {
	ID int
	Username string
	IsActive bool
}

func FormatUser(u User) (string, error) {
	if !u.IsActive {
		return "", fmt.Errorf("user %s is not active", u.Username)
	}

	result := fmt.Sprintf("User ID: %d, Username: %s", u.ID, u.Username)
	return result, nil
}

func main() {
	age := 27
	name := "John Doe"
	var role string = "Admin"

	fmt.Printf("Name: %s, Age: %d, Role: %s\n", name, age, role)
	fmt.Println("--------------------------------")

	user1 := User{
		ID : 1,
		Username : "johndoe",
		IsActive : true,
	}

	info, err := FormatUser(user1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Hasil Format : " + info)

    fmt.Println("Hello World dari Go!")
}
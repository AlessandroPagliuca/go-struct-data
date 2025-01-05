package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	//FirstName is a public field
	FirstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// struct embedded in another struct
type Admin struct {
	email    string
	password string
	User
}

func (u User) OutputUserDetails() {
	fmt.Println(u.FirstName, u.lastName, u.birthdate)
}

func (u *User) ClearUserName() {
	u.FirstName = ""
	u.lastName = ""
}

func NewAdmin(email, password string) Admin {
	// create a new Admin struct
	return Admin{
		email:    email,
		password: password,
		User: User{
			FirstName: "ADMIN",
			lastName:  "ADMIN",
			birthdate: "-----",
			createdAt: time.Now(),
		},
	}

}

// naming connvention for constructor New or NewUser (in this case) is a constructor function that returns a new User struct
func New(firstName, lastName, birthdate string) (*User, error) {

	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name, last name, and birthdate are required")

	}

	// return a pointer to a new User struct
	return &User{
		FirstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

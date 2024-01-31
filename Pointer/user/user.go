package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName string
	birthdate string
	created time.Time
}

type Admin struct{
	email string
	password string
	User
}

func NewAdmin (email, password string) Admin{
	return Admin{
		email: email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName: "Admin",
			birthdate: "___",
			created: time.Now(),
		},
	}

}


func (u *User) OutputUserDetails()  {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
	
}

func (u *User) ClearUsername()  {
	u.firstName = ""
	u.lastName = ""
	u.birthdate = ""
	
}
func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == ""{
		return nil, errors.New("Firstname, Lastname and Brithdate are required.")
	}
	return &User{
		firstName: firstName,
		lastName: lastName,
		birthdate: birthdate,
		created: time.Now(),
	},nil
	
}

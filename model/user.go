package model

import (
	"github.com/google/uuid"
)

type Role int

const (
	Admin Role = iota + 1
	RestorantManager
	Customer
	Driver
	Undefined
	MaxRole
)

var roleType = [MaxRole]string{
	Admin:            "admin",
	RestorantManager: "restorant manager",
	Customer:         "customer",
	Driver:           "driver",
	Undefined:        "undefined",
}

func (r Role) String() string {
	return roleType[r]
}

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	RoleType  Role   `json:"role_type"`
}

func NewUser(firstName, lastName string, age int) *User {
	return &User{
		ID:        uuid.NewString(),
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
		RoleType:  Customer,
	}
}

var UserList = map[string]*User{}

func (u *User) SayHi() string {
	return "Hello " + u.FirstName
}

package model

import (
	"encoding/json"
	"log"
)

type Role int

const (
	Admin Role = iota + 1
	RestaurantManager
	Customer
	Driver
	Undefined
	MaxRole
)

var roleType = [MaxRole]string{
	Admin:             "admin",
	RestaurantManager: "restaurant manager",
	Customer:          "customer",
	Driver:            "driver",
	Undefined:         "undefined",
}

func (r Role) String() string {
	return roleType[r]
}

func ToRoleType(s string) Role {
	switch s {
	case "admin":
		return Admin
	case "restaurant manager":
		return RestaurantManager
	case "customer":
		return Customer
	case "driver":
		return Driver
	default:
		return Undefined
	}
}

func (r *Role) UnmarshalJSON(text []byte) error {
	log.Println("UnmarshalJSON recived Role:", string(text))

	var s string
	if err := json.Unmarshal(text, &s); err != nil {
		return err
	}
	log.Println("UnmarshalJSON unmarshal Role:", s)

	*r = ToRoleType(s)

	return nil
}

func (r Role) MarshalJSON() ([]byte, error) {
	log.Println("MarshalJSON Role:", r.String())
	return []byte(`"` + r.String() + `"`), nil
}

// User is the model for a user.
type User struct {
	// DBData are the fields that are stored in the database.
	DBData
	// Email is the user's email.
	Email string `json:"email"`
	// Password is the user's password.
	Password *string `json:"password,omitempty"`
	// RoleType is the user's role.
	RoleType Role `json:"role_type"`
	// Address is the user's address.
	Address *Address `json:"address,omitempty"`
	// ConfigUser is the user's optional fields.
	ConfigUser
}

// Login is the model for a login.
// It helps to validate the user's email and password.
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ConfigUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Age       int    `json:"age"`
}

func (u *User) SayHi() string {
	return "Hello " + u.FirstName
}

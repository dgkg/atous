package model

import (
	"encoding/json"
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
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

type User struct {
	DBData
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleType Role   `json:"role_type"`
	ConfigUser
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DBData struct {
	ID string `json:"id"`
	// DB dates.
	CreateAt time.Time  `json:"create_at"`
	UpdateAt time.Time  `json:"update_at"`
	DeleteAt *time.Time `json:"delete_at"`
}

type ConfigUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Age       int    `json:"age"`
}

func NewUser(email, password string, config *ConfigUser) *User {
	var u User
	u.ID = uuid.NewString()
	u.Email = email
	u.Password = password
	u.RoleType = Customer
	if config != nil {
		u.Age = config.Age
		u.FirstName = config.FirstName
		u.LastName = config.LastName
		u.Phone = config.Phone
	}
	u.CreateAt = time.Now()

	spew.Printf("Model : NewUser: %#v\n", u)
	return &u
}

func (u *User) SayHi() string {
	return "Hello " + u.FirstName
}

package model

import (
	"time"

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
	DBData
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleType Role   `json:"role_type"`
	ConfigUser
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
	return &u
}

var UserList = map[string]*User{}

func (u *User) SayHi() string {
	return "Hello " + u.FirstName
}

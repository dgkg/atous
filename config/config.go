package config

import "os"

func Exec() string {
	return "Hello World"
}

var (
	Login = os.Getenv("TEST_LOGIN")
	Pass  = os.Getenv("TEST_PASS")
)

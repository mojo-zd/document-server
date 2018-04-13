package models

type User struct {
	Id       int64
	Name     string
	Password string
	Salt     string
}

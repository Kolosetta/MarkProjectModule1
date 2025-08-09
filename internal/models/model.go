package models

type User struct {
	Id       int64
	Username string
	Email    string
}

type Post struct {
	Id     int64
	Author string
	Text   string
}

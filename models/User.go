package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User //pointers to User objects
	nextID = 1
)

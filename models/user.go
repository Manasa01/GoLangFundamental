package models

import {
	errors
}
type User struct {
	ID int
	FirstName string
    LastName string
}

var (
	users []*User //slice of pointers to User, can update throughtout the application without copying around
	nextID = 1 //initialization at the package level, need not use :=
)

func GetUsers() []*User {
    return users
}

func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New User must not include id or it must be set to zero")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u,nil
}
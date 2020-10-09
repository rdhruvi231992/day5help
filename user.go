package main

// UserStore is database()
var UserStore []*User

// User ()
type User struct {
	ID    string
	Email string
}

func createUser(u *User) *User {
	UserStore = append(UserStore, u)
	return u
}

func listUsers() []*User {
	return UserStore

}

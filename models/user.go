package models

// User model
type User struct {
	Name        string
	Email       string
	PhoneNumber string
}

// NewUser - simple User constructor
func NewUser(name, email, phoneNumber string) *User {
	user := new(User)
	user.Name = name
	user.Email = email
	user.PhoneNumber = phoneNumber

	return user
}

package moviestore

import "fmt"

// Age is just an unsigned 8-bit int, we do not expect older users.
type Age uint8

// UserID is just an unsigned 16-bit int
type UserID uint16

// User with name, age and id
type User struct {
	Name   string
	Age    Age
	UserID UserID
}

// Returns a string representing the use like that:
// "   8. Helga Meier, 28"
// The userid field should be 4 digits wide.
func (u *User) String() string {
	return fmt.Sprintf("%4d. %s,%3d", u.UserID, u.Name, u.Age)
}

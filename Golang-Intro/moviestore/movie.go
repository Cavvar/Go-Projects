package moviestore

import "fmt"

// FSK is an unsigned 8-bit int
type FSK uint8

// Serial is an unsigned int
type Serial uint

const (
	// FSK0 age 0
	FSK0 FSK = 0
	// FSK6 age 6
	FSK6 FSK = 6
	// FSK12 age 12
	FSK12 FSK = 12
	// FSK16 age 16
	FSK16 FSK = 16
	// FSK18 age 18
	FSK18 FSK = 18
)

// A Movie consists of a title, the fsk minimum age and a serial
type Movie struct {
	Title  string
	Fsk    FSK
	Serial Serial
}

// AllowedAtAge checks whether the movie is allowed at a given age or not.
func AllowedAtAge(m *Movie, age Age) bool {
	return uint8(age) >= uint8(m.Fsk)
}

// Returns a string representing the movie like that:
// "  23. Texas Chainsaw Massacre (Ab 18)"
// The serial field should be 4 digits wide.
func (m *Movie) String() string {
	return fmt.Sprintf("%4d. %s (Ab %d)", m.Serial, m.Title, m.Fsk)
}

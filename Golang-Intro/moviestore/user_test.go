package moviestore

import "testing"

func TestUser(t *testing.T) {
	cases := []struct {
		in   User
		want string
	}{
		{User{"Helga Meier", Age(28), UserID(8)}, "   8. Helga Meier, 28"},
		{User{}, "   0. ,  0"},
		{User{"Denis Angeletta", Age(25), UserID(60)}, "  60. Denis Angeletta, 25"},
	}
	for _, c := range cases {
		got := c.in.String()
		if got != c.want {
			t.Errorf("%q.String() == %q, want %q", c.in, got, c.want)
		}
	}
}

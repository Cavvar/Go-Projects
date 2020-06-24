package moviestore

import "testing"

func TestMovie(t *testing.T) {
	cases := []struct {
		in   Movie
		want string
	}{
		{Movie{"Am Limit", FSK0, 12}, "  12. Am Limit (Ab 0)"},
		{Movie{"Texas Chainsaw Massacre", FSK18, 8}, "   8. Texas Chainsaw Massacre (Ab 18)"},
		{Movie{"Inglourious Basterds", FSK16, 13}, "  13. Inglourious Basterds (Ab 16)"},
	}
	for _, c := range cases {
		got := c.in.String()
		if got != c.want {
			t.Errorf("%q.String() == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestAllowedAtAge(t *testing.T) {
	cases := []struct {
		age       Age
		fsk       FSK
		isAllowed bool
	}{
		{Age(0), FSK0, true},
		{Age(0), FSK6, false},
		{Age(0), FSK12, false},
		{Age(0), FSK16, false},
		{Age(0), FSK18, false},
		{Age(5), FSK6, false},
		{Age(11), FSK12, false},
		{Age(15), FSK16, false},
		{Age(17), FSK18, false},
		{Age(6), FSK6, true},
		{Age(12), FSK12, true},
		{Age(16), FSK16, true},
		{Age(18), FSK18, true},
	}
	for _, c := range cases {
		got := AllowedAtAge(&Movie{Fsk: c.fsk}, c.age)
		if got != c.isAllowed {
			t.Errorf("got: %v, wanted: %v", got, c.isAllowed)
		}
	}
}

package moviestore

import (
	"testing"
)

func newMoviestoreImpl() *moviestoreImpl {
	ms := new(moviestoreImpl)
	ms.available = make(map[Serial]Movie)
	ms.users = make(map[UserID]User)
	ms.rented = make(map[UserID][]Movie)

	return ms
}

func TestMoviestore_AddMovie(t *testing.T) {
	cases := []struct {
		title string
		fsk   FSK
	}{
		{"Am Limit", FSK0},
		{"Texas Chainsaw Massacre", FSK18},
		{"Inglourious Basterds", FSK16},
	}
	moviestore := newMoviestoreImpl()

	for _, c := range cases {
		gotSerial := moviestore.AddMovie(c.title, c.fsk)
		gotMovie, present := moviestore.available[gotSerial]

		if !present || gotMovie.Title != c.title || gotMovie.Fsk != c.fsk {
			t.Errorf("Added %v, but got %v", c, gotMovie)
		}
	}
}

func TestMoviestore_AddUser(t *testing.T) {
	cases := []struct {
		name string
		age  Age
	}{
		{"Hugo", 12},
		{"Helga", 19},
		{"Ronja", 17},
	}
	moviestore := newMoviestoreImpl()

	for _, c := range cases {
		gotUserID := moviestore.AddUser(c.name, c.age)
		gotUser, present := moviestore.users[gotUserID]

		if !present || gotUser.Name != c.name || gotUser.Age != c.age {
			t.Errorf("Added %v, but got %v", c, gotUser)
		}
	}
}

func TestMoviestore_Rent(t *testing.T) {
	cases := []struct {
		serial  Serial
		userID  UserID
		isError error
	}{
		{Serial(0), UserID(0), nil},
		{Serial(1), UserID(6), userMissingError(UserID(6))},
		{Serial(4), UserID(1), missingMovieError(Serial(1))},
		{Serial(1), UserID(1), userNotOfAgeError(UserID(1))},
	}
	moviestore := newMoviestoreImpl()
	moviestore.AddUser("Denis Angeletta", Age(25))
	moviestore.AddUser("Max Mustermann", Age(15))
	moviestore.AddMovie("IT", FSK16)
	moviestore.AddMovie("IT-2", FSK16)

	for _, c := range cases {
		_, _, err := moviestore.Rent(c.serial, c.userID)
		if err != nil {
			t.Logf("Error return: %v", err)
		}
	}
}

func TestMoviestore_RentedByUser(t *testing.T) {
	cases := []struct {
		userID      UserID
		countMovies int
	}{
		{UserID(0), 2},
		{UserID(1), 0},
		{UserID(5), 0},
	}
	moviestore := newMoviestoreImpl()
	userIDone := moviestore.AddUser("Denis Angeletta", Age(25))
	moviestore.AddUser("Max Mustermann", Age(15))
	movieSerialOne := moviestore.AddMovie("IT", FSK16)
	movieSerialTwo := moviestore.AddMovie("IT-2", FSK16)
	_, _, errRentOne := moviestore.Rent(movieSerialOne, userIDone)
	_, _, errRentTwo := moviestore.Rent(movieSerialTwo, userIDone)

	if errRentOne == nil && errRentTwo == nil {
		for _, c := range cases {
			movies, err := moviestore.RentedByUser(c.userID)
			if err != nil {
				switch err.(type) {
				case userMissingError:
					t.Log("Everything is fine")
				default:
					t.Errorf("UserMissingError expected but not thrown!")
				}
			} else if len(movies) != c.countMovies {
				t.Errorf("Number of movies expected: %d, Number of movies actually in map: %d", c.countMovies, len(movies))
			}
		}
	}
}

func TestMoviestore_Return(t *testing.T) {
	cases := []struct {
		serial Serial
	}{
		{Serial(0)},
		{Serial(1)},
		{Serial(2)},
	}
	moviestore := newMoviestoreImpl()
	userIDone := moviestore.AddUser("Denis Angeletta", Age(25))
	moviestore.AddUser("Max Mustermann", Age(15))
	movieSerialOne := moviestore.AddMovie("IT", FSK16)
	movieSerialTwo := moviestore.AddMovie("IT-2", FSK16)
	_, _, errRentOne := moviestore.Rent(movieSerialOne, userIDone)
	_, _, errRentTwo := moviestore.Rent(movieSerialTwo, userIDone)

	if errRentOne == nil && errRentTwo == nil {
		for _, c := range cases {
			_, _, err := moviestore.Return(c.serial)
			if err != nil {
				switch err.(type) {
				case movieNotRentedError:
					t.Log("Everything is fine!")
				default:
					t.Errorf("Not expected error type")
				}
			}
		}
	}
}

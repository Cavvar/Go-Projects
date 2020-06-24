package moviestore

import "fmt"

// A Moviestore interface
type Moviestore interface {
	AddMovie(string, FSK) Serial
	AddUser(string, Age) UserID
	Rent(Serial, UserID) (User, Movie, error)
	Return(Serial) (User, Movie, error)
	RentedByUser(UserID) ([]Movie, error)
}

// NewMoviestore generates a Moviestore which contains available movies,
// users and a mapping between users and currently rented movies.
func NewMoviestore() Moviestore {
	return nil
}

type moviestoreImpl struct {
	available  map[Serial]Movie
	rented     map[UserID][]Movie
	users      map[UserID]User
	nextSerial Serial
	nextUserID UserID
}

type missingMovieError Serial
type userMissingError UserID
type userNotOfAgeError UserID
type movieNotRentedError Serial

func (err missingMovieError) Error() string {
	return fmt.Sprintf("movie not available for rent")
}

func (err userMissingError) Error() string {
	return fmt.Sprintf("user id not found")
}

func (err userNotOfAgeError) Error() string {
	return fmt.Sprintf("user is too young")
}

func (err movieNotRentedError) Error() string {
	return fmt.Sprintf("movie not found in rented movies")
}

// AddMovie adds a movie to the available movies map which is part of themoviestoreImpl struct:
// available   map[Serial]Movie
// The serial will be generated and returned.
func (ms *moviestoreImpl) AddMovie(title string, fsk FSK) Serial {
	resultSerial := ms.nextSerial // Do we have one movie with several serials or one movie with one serial id
	ms.nextSerial++
	ms.available[resultSerial] = Movie{title, fsk, resultSerial}

	return resultSerial
}

// AddUser adds an user to the users map which is part of themoviestoreImpl struct:
// users   map[UserID]User
// The userid will be generated and returned.
func (ms *moviestoreImpl) AddUser(name string, age Age) UserID {
	resultUserID := ms.nextUserID
	ms.nextUserID++
	ms.users[resultUserID] = User{name, age, resultUserID}

	return resultUserID
}

// Rent a movie. If the user is in users and the movie is in available,
// the movie will be removed from available and appended to the slice of rented
// movies by this user.
// rented	map[UserID] []Movie
// The following error cases are handled and will be returned as error containing the following texts:
//	- user not found
//	- movie not available for rent
//	- user ist too young
func (ms *moviestoreImpl) Rent(serial Serial, userID UserID) (User, Movie, error) {
	user, userFound := ms.users[userID]
	movie, movieInAvailable := ms.available[serial]

	if !userFound {
		return User{UserID: userID}, Movie{Serial: serial}, userMissingError(userID)
	}

	if !movieInAvailable {
		return user, Movie{Serial: serial}, missingMovieError(serial)
	}

	if uint8(user.Age) < uint8(movie.Fsk) {
		return user, movie, userNotOfAgeError(userID)
	}

	delete(ms.available, serial)
	newRentedForUser := append(ms.rented[userID], movie)
	ms.rented[userID] = newRentedForUser

	return user, movie, nil
}

// RentedByUser returns a slice of movies currently rented by the user.
// Error case is "userID unknown"
// Be aware that slices are returned byreference.
func (ms *moviestoreImpl) RentedByUser(userID UserID) ([]Movie, error) {
	_, userFound := ms.users[userID]
	if !userFound {
		return []Movie{}, userMissingError(userID)
	}

	moviesRentedByUser := ms.rented[userID]
	resultSlice := make([]Movie, len(moviesRentedByUser))
	copy(resultSlice, moviesRentedByUser)

	return resultSlice, nil
}

// Return a movie. Searches in all slices of the rented map,
// does some "housekeeping", and returns the user and movie if found.
// The error case is "movie not found in rented movies".
func (ms *moviestoreImpl) Return(serial Serial) (User, Movie, error) {
	for userID, movies := range ms.rented {
		for movieIndex, movie := range movies {
			if movie.Serial == serial {
				movies = append(movies[:movieIndex], movies[movieIndex+1:]...)
				ms.rented[userID] = movies
				ms.available[movie.Serial] = movie

				return ms.users[userID], movie, nil
			}
		}
	}

	// errors.New("")
	return User{}, Movie{}, movieNotRentedError(serial)
}

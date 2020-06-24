package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro"
	protoMovie "github.com/ob-vss-ws19/blatt-4-cavvar/movie/proto"
	protoReservation "github.com/ob-vss-ws19/blatt-4-cavvar/reservation/proto"
	protoRoom "github.com/ob-vss-ws19/blatt-4-cavvar/room/proto"
	protoShowing "github.com/ob-vss-ws19/blatt-4-cavvar/showing/proto"
	protoUser "github.com/ob-vss-ws19/blatt-4-cavvar/user/proto"
)

func main() {
	// Create a new service
	service := micro.NewService(
		micro.Name("cinema.client"),
	)

	// Initialise client and parse command line flags if available
	service.Init()

	// Create clients of services
	movie := protoMovie.NewMovieService("movie", service.Client())
	showing := protoShowing.NewShowingService("showing", service.Client())
	room := protoRoom.NewRoomService("room", service.Client())
	user := protoUser.NewUserService("user", service.Client())
	reservation := protoReservation.NewReservationService("reservation", service.Client())

	// Create four movies!
	movieOne, _ := movie.CreateMovie(context.TODO(), &protoMovie.CreateMovieRequest{
		Name: "Avengers Endgame",
		Fsk:  12,
	})
	fmt.Printf(movieOne.Message)
	movieTwo, _ := movie.CreateMovie(context.TODO(), &protoMovie.CreateMovieRequest{
		Name: "Joker",
		Fsk:  16,
	})
	fmt.Printf(movieTwo.Message)
	movieThree, _ := movie.CreateMovie(context.TODO(), &protoMovie.CreateMovieRequest{
		Name: "Star Wars The Force Awakens",
		Fsk:  12,
	})
	fmt.Printf(movieThree.Message)
	movieFour, _ := movie.CreateMovie(context.TODO(), &protoMovie.CreateMovieRequest{
		Name: "Interstellar",
		Fsk:  12,
	})
	fmt.Printf(movieFour.Message)

	// Create two rooms
	roomOne, _ := room.AddRoom(context.TODO(), &protoRoom.AddRoomRequest{
		RoomName:        "Room 1",
		IsRoomAvailable: true,
		AvailableSeats:  150,
	})
	fmt.Printf("%v\n", roomOne.Message)
	roomTwo, _ := room.AddRoom(context.TODO(), &protoRoom.AddRoomRequest{
		RoomName:        "Room 2",
		IsRoomAvailable: true,
		AvailableSeats:  154,
	})
	fmt.Printf("%v\n", roomTwo.Message)

	// Create 4 users
	userOne, _ := user.CreateUser(context.TODO(), &protoUser.CreateUserRequest{
		FirstName: "Angela",
		LastName:  "Merkel",
		Age:       68,
	})
	fmt.Printf("%v\n", userOne.Message)
	userTwo, _ := user.CreateUser(context.TODO(), &protoUser.CreateUserRequest{
		FirstName: "Donald",
		LastName:  "Trump",
		Age:       73,
	})
	fmt.Printf("%v\n", userTwo.Message)
	userThree, _ := user.CreateUser(context.TODO(), &protoUser.CreateUserRequest{
		FirstName: "Emmanuel",
		LastName:  "Macro",
		Age:       40,
	})
	fmt.Printf("%v\n", userThree.Message)
	userFour, _ := user.CreateUser(context.TODO(), &protoUser.CreateUserRequest{
		FirstName: "Justin",
		LastName:  "Trudeau",
		Age:       42,
	})
	fmt.Printf("%v\n", userFour.Message)

	// Create four showings
	showingOne, _ := showing.CreateShowing(context.TODO(), &protoShowing.CreateShowingRequest{
		MovieID:  movieOne.NewMovie.Id,
		RoomName: roomOne.Room.RoomName,
	})
	fmt.Printf("%v\n", showingOne.Message)
	showingTwo, _ := showing.CreateShowing(context.TODO(), &protoShowing.CreateShowingRequest{
		MovieID:  movieTwo.NewMovie.Id,
		RoomName: roomTwo.Room.RoomName,
	})
	fmt.Printf("%v\n", showingTwo.Message)
	showingThree, _ := showing.CreateShowing(context.TODO(), &protoShowing.CreateShowingRequest{
		MovieID:  movieThree.NewMovie.Id,
		RoomName: roomOne.Room.RoomName,
	})
	fmt.Printf("%v\n", showingThree.Message)
	showingFour, _ := showing.CreateShowing(context.TODO(), &protoShowing.CreateShowingRequest{
		MovieID:  movieFour.NewMovie.Id,
		RoomName: roomTwo.Room.RoomName,
	})
	fmt.Printf("%v\n", showingFour.Message)

	// Create four reservations
	initReservationOne, _ := reservation.InitReservation(context.TODO(), &protoReservation.InitReservationRequest{
		UserID:    userOne.NewUser.UserID,
		ShowingID: showingOne.NewShowing.ShowingID,
		SeatCount: 5,
	})
	fmt.Printf("%v\n", initReservationOne.Message)
	initReservationTwo, _ := reservation.InitReservation(context.TODO(), &protoReservation.InitReservationRequest{
		UserID:    userTwo.NewUser.UserID,
		ShowingID: showingTwo.NewShowing.ShowingID,
		SeatCount: 2,
	})
	fmt.Printf("%v\n", initReservationTwo.Message)
	initReservationThree, _ := reservation.InitReservation(context.TODO(), &protoReservation.InitReservationRequest{
		UserID:    userThree.NewUser.UserID,
		ShowingID: showingThree.NewShowing.ShowingID,
		SeatCount: 1,
	})
	fmt.Printf("%v\n", initReservationThree.Message)
	initReservationFour, _ := reservation.InitReservation(context.TODO(), &protoReservation.InitReservationRequest{
		UserID:    userFour.NewUser.UserID,
		ShowingID: showingFour.NewShowing.ShowingID,
		SeatCount: 3,
	})
	fmt.Printf("%v\n", initReservationFour.Message)
	reservationOne, _ := reservation.CreateReservation(context.TODO(), &protoReservation.CreateReservationRequest{
		UserID:        userOne.NewUser.UserID,
		ShowingID:     showingOne.NewShowing.ShowingID,
		SeatCount:     5,
		DoReservation: true,
	})
	fmt.Printf("%v\n", reservationOne.Message)
	reservationTwo, _ := reservation.CreateReservation(context.TODO(), &protoReservation.CreateReservationRequest{
		UserID:        userTwo.NewUser.UserID,
		ShowingID:     showingTwo.NewShowing.ShowingID,
		SeatCount:     2,
		DoReservation: true,
	})
	fmt.Printf("%v\n", reservationTwo.Message)
	reservationThree, _ := reservation.CreateReservation(context.TODO(), &protoReservation.CreateReservationRequest{
		UserID:        userThree.NewUser.UserID,
		ShowingID:     showingThree.NewShowing.ShowingID,
		SeatCount:     1,
		DoReservation: true,
	})
	fmt.Printf("%v\n", reservationThree.Message)
	reservationFour, _ := reservation.CreateReservation(context.TODO(), &protoReservation.CreateReservationRequest{
		UserID:        userFour.NewUser.UserID,
		ShowingID:     showingFour.NewShowing.ShowingID,
		SeatCount:     3,
		DoReservation: true,
	})
	fmt.Printf("%v\n", reservationFour.Message)
	// ---------------------------------------------------
	// Scenario 1
	readRoomResponse, _ := room.FindRoomByID(context.TODO(), &protoRoom.FindRoomByIDRequest{
		RoomName: roomTwo.Room.RoomName,
	})
	fmt.Printf("Before delete: %v\n", readRoomResponse.Message)
	removeResponse, _ := room.RemoveRoomByID(context.TODO(), &protoRoom.RemoveRoomByIDRequest{
		RoomName: roomTwo.Room.RoomName,
	})
	fmt.Printf("%v\n", removeResponse.Message)
	_, err := showing.ReadShowing(context.TODO(), &protoShowing.ReadShowingRequest{
		ShowingID: showingTwo.NewShowing.ShowingID,
	})
	if err != nil {
		fmt.Printf("Showing 2 in room 2 should not exist: %v\n", err)
	}
	_, err = showing.ReadShowing(context.TODO(), &protoShowing.ReadShowingRequest{
		ShowingID: showingFour.NewShowing.ShowingID,
	})
	if err != nil {
		fmt.Printf("Showing 4 in room 2 should not exist: %v\n", err)
	}
	_, errAfterDelete := room.FindRoomByID(context.TODO(), &protoRoom.FindRoomByIDRequest{
		RoomName: roomTwo.Room.RoomName,
	})
	if errAfterDelete != nil {
		fmt.Printf("%v\n", errAfterDelete)
	}

	// Scenario 2
	initReservationScenarioOne, _ := reservation.InitReservation(context.TODO(), &protoReservation.InitReservationRequest{
		UserID:    userOne.NewUser.UserID,
		ShowingID: showingOne.NewShowing.ShowingID,
		SeatCount: 100,
	})
	fmt.Printf("%v\n", initReservationScenarioOne.Message)
	// This should already fail
	_, err = reservation.InitReservation(context.TODO(), &protoReservation.InitReservationRequest{
		UserID:    userTwo.NewUser.UserID,
		ShowingID: showingOne.NewShowing.ShowingID,
		SeatCount: 80,
	})
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

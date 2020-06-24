package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/micro/go-micro"
	protoReservation "github.com/ob-vss-ws19/blatt-4-cavvar/reservation/proto"
	proto "github.com/ob-vss-ws19/blatt-4-cavvar/user/proto"
)

type user struct {
	firstName, lastName string
	age                 int
	userID              string
}

type UserHandler struct {
	userDatabase       *sync.Map // "Database". Non Blocking Data Structure or mutex needed to avoid deadlocks
	reservationService protoReservation.ReservationService
}

func (handler UserHandler) CreateUser(ctx context.Context, req *proto.CreateUserRequest, res *proto.CreateUserResponse) error {
	// Create new userID
	newID := fmt.Sprintf("%v %v %v", req.FirstName, req.LastName, req.Age)
	// Add to database
	newUser := user{
		firstName: req.FirstName,
		lastName:  req.LastName,
		age:       int(req.Age),
		userID:    newID,
	}
	handler.userDatabase.Store(newID, newUser)
	log.Printf("User %v was added to the database!\n", newUser)
	// Respond with a message
	res.Message = fmt.Sprintf("User %v was successfully added to the database", newUser.userID)
	res.NewUser = &proto.UserStruct{
		FirstName: newUser.firstName,
		LastName:  newUser.lastName,
		Age:       int32(newUser.age),
		UserID:    newUser.userID,
	}
	return nil
}

func (handler UserHandler) DeleteUser(ctx context.Context, req *proto.DeleteUserRequest, res *proto.DeleteUserResponse) error {
	// See if the user actually exists
	userTemp, ok := handler.userDatabase.Load(req.UserID)
	if !ok {
		// User does NOT exist
		// Set Error Return Value
		log.Printf("The requested user was NOT found in the database")
		return fmt.Errorf("user %v does NOT exist! Nothing will be deleted", req.UserID)
	}
	actualUser := userTemp.(user)
	// User does exist
	// Remove that user
	// User can only be removed if he has no more reservations!
	allReservations, err := handler.reservationService.GetReservationsForUser(context.TODO(), &protoReservation.GetReservationForUserRequest{
		UserID: actualUser.userID,
	})
	if err != nil {
		// Communication with reservation service failed
		return err
	}
	if len(allReservations.AllReservations) != 0 {
		// Users still has reservations available!
		return fmt.Errorf("user %v still has reservations available. Cannot delete user", req.UserID)
	}
	log.Printf("The requested user does exist in the database. The user %v will be deleted", actualUser.userID)
	res.Message = fmt.Sprintf("User %v was deleted from the user database", actualUser.userID)
	return nil
}

func (handler UserHandler) ShowAllUser(ctx context.Context, req *proto.ShowAllUserRequest, res *proto.ShowAllUserResponse) error {
	// Create an array of user struct
	// Iterate through database
	// Add users to the array of user struct
	allUsers := make([]*proto.UserStruct, 0)
	handler.userDatabase.Range(func(key, value interface{}) bool {
		switch m := value.(type) {
		case *proto.UserStruct:
			allUsers = append(allUsers, m)
			return true
		default:
			return false
		}
	})
	res.AllUsers = allUsers
	return nil
}

func (handler UserHandler) ReadUser(ctx context.Context, req *proto.ReadUserRequest, res *proto.ReadUserResponse) error {
	// Check if that user ID exists in the database
	userTemp, ok := handler.userDatabase.Load(req.UserID)
	if !ok {
		// User was not found!
		log.Printf("User was NOT found!\n")
		res.Message = "User was NOT found"
		res.User = nil
		return fmt.Errorf("user %v was NOT found", req.UserID)
	}
	actualUser := userTemp.(user)
	// User was found!
	log.Printf("user %v was found!\n", actualUser)
	res.User = &proto.UserStruct{
		FirstName: actualUser.firstName,
		LastName:  actualUser.lastName,
		Age:       int32(actualUser.age),
		UserID:    actualUser.userID,
	}
	res.Message = fmt.Sprintf("User %v exists in the current database!", actualUser.userID)
	return nil
}

func NewUserHandler(reservationService protoReservation.ReservationService) *UserHandler {
	// Init Method needed to avoid initialization errors
	userHandler := &UserHandler{
		userDatabase:       &sync.Map{},
		reservationService: reservationService,
	}
	return userHandler
}

func main() {
	// Create a new user service
	userService := micro.NewService(
		micro.Name("user"),
	)

	// Initialize service and parse command line flags if available
	userService.Init()

	// Init Reservation service
	reservationService := protoReservation.NewReservationService("reservation", userService.Client())

	// Register handlers
	err := proto.RegisterUserHandler(userService.Server(), NewUserHandler(reservationService))
	if err != nil {
		fmt.Printf("Could not register user handler!")
	}

	// Run the server or print the error
	if err := userService.Run(); err != nil {
		fmt.Printf("Error is %v\n", err)
	}
}

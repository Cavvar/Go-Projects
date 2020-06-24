package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/micro/go-micro"
	proto "github.com/ob-vss-ws19/blatt-4-cavvar/room/proto"
	protoShowing "github.com/ob-vss-ws19/blatt-4-cavvar/showing/proto"
)

type room struct {
	roomName        string
	availableSeats  int
	isRoomAvailable bool
}

type RoomHandler struct {
	roomDatabase   *sync.Map // "Database". Non Blocking Data Structure or mutex needed to avoid deadlocks
	showingService protoShowing.ShowingService
}

func (handler RoomHandler) AddRoom(ctx context.Context, req *proto.AddRoomRequest, res *proto.AddRoomResponse) error {
	// Add to database
	newRoom := room{
		roomName:        req.RoomName,
		availableSeats:  int(req.AvailableSeats),
		isRoomAvailable: true,
	}
	handler.roomDatabase.Store(newRoom.roomName, newRoom)
	log.Printf("Room %v was added to the database!\n", newRoom)
	// Respond with a message
	res.Message = fmt.Sprintf("%v with %v seats was successfully added to the database", newRoom.roomName, newRoom.availableSeats)
	res.Room = &proto.RoomStruct{
		RoomName:        newRoom.roomName,
		IsRoomAvailable: newRoom.isRoomAvailable,
		AvailableSeats:  int32(newRoom.availableSeats),
	}
	return nil
}

func (handler RoomHandler) RemoveRoomByID(ctx context.Context, req *proto.RemoveRoomByIDRequest, res *proto.RemoveRoomByIDResponse) error {
	// See if the room actually exists
	roomTemp, ok := handler.roomDatabase.Load(req.RoomName)
	if !ok {
		// room does not exist
		// Set Error Return Value
		log.Printf("The requested room %v for deletion does NOT exist in the database!\n", req.RoomName)
		return fmt.Errorf("room %v does NOT exist! Nothing will be deleted", req.RoomName)
	}
	actualRoom := roomTemp.(room)
	// room does exist
	// Remove that room
	// Don't forget that if a room gets deleted it causes a 'chain reaction'
	_, err := handler.showingService.DeleteShowingsForRoom(context.TODO(), &protoShowing.DeleteShowingsForRoomRequest{RoomName: req.RoomName})
	if err != nil {
		return err
	}
	// Send Success Message back
	log.Printf("The requested room for deletion does exist in the database! room %v will be deleted!\n", actualRoom.roomName)
	handler.roomDatabase.Delete(req.RoomName)
	res.Message = fmt.Sprintf("room %v was deleted from the room database!\n", actualRoom.roomName)
	return nil
}

func (handler RoomHandler) ShowAllRooms(ctx context.Context, req *proto.ShowAllRoomsRequest, res *proto.ShowAllRoomsResponse) error {
	// Create an array of room struct
	// Iterate through database
	// Add rooms to the array of room struct
	allRooms := make([]*proto.RoomStruct, 0)
	handler.roomDatabase.Range(func(key, value interface{}) bool {
		switch m := value.(type) {
		case *proto.RoomStruct:
			allRooms = append(allRooms, m)
			return true
		default:
			return false
		}
	})
	res.Rooms = allRooms
	return nil
}

func (handler RoomHandler) FindRoomByID(ctx context.Context, req *proto.FindRoomByIDRequest, res *proto.FindRoomByIDResponse) error {
	// Check if that room exists in the database
	roomTemp, ok := handler.roomDatabase.Load(req.RoomName)
	if !ok {
		// room was not found!
		log.Printf("Room was NOT found!\n")
		res.Message = "Room was NOT found"
		res.Room = nil
		return fmt.Errorf("room with %v was NOT found", req.RoomName)
	}
	actualRoom := roomTemp.(room)
	// Room was found!
	log.Printf("room %v was found!\n", actualRoom)
	res.Room = &proto.RoomStruct{
		RoomName:        actualRoom.roomName,
		IsRoomAvailable: actualRoom.isRoomAvailable,
		AvailableSeats:  int32(actualRoom.availableSeats),
	}
	res.Message = fmt.Sprintf("Room %v exists in the current database!", actualRoom.roomName)
	return nil
}

func (handler RoomHandler) SetRoomAvailability(ctx context.Context, req *proto.SetRoomAvailabilityRequest, res *proto.SetRoomAvailabilityResponse) error {
	// Check if the room exists
	roomTemp, ok := handler.roomDatabase.Load(req.RoomName)
	if !ok {
		// room was not found!
		log.Printf("Room was NOT found\n")
		res.Message = "Room was NOT found"
		res.Room = nil
		return fmt.Errorf("room %v was NOT found. Room availability will not be set", req.RoomName)
	}
	actualRoom := roomTemp.(room)
	// Room was found
	log.Printf("room %v was found\n", actualRoom)
	updatedRoom := &proto.RoomStruct{
		RoomName:        actualRoom.roomName,
		IsRoomAvailable: req.IsRoomAvailable, // Updated Room Availability
		AvailableSeats:  int32(actualRoom.availableSeats),
	}
	// Update room in database
	handler.roomDatabase.Delete(req.RoomName)
	handler.roomDatabase.Store(req.RoomName, updatedRoom)
	// Set Response
	res.Room = updatedRoom
	return nil
}

func NewRoomHandler(showingService protoShowing.ShowingService) *RoomHandler {
	// Init Method needed to avoid initialization errors
	userHandler := &RoomHandler{
		roomDatabase:   &sync.Map{},
		showingService: showingService,
	}
	return userHandler
}

func main() {
	// Create a new room service
	roomService := micro.NewService(
		micro.Name("room"),
	)

	// Initialize service and parse command line flags if available
	roomService.Init()

	// Init Services
	showingService := protoShowing.NewShowingService("showing", roomService.Client())

	// Register handlers
	err := proto.RegisterRoomHandler(roomService.Server(), NewRoomHandler(showingService))
	if err != nil {
		fmt.Printf("Could not register user handler!")
	}

	// Run the server or print the error
	if err := roomService.Run(); err != nil {
		fmt.Printf("Error is %v\n", err)
	}
}

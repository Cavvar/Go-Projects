package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/micro/go-micro"
	protoMovie "github.com/ob-vss-ws19/blatt-4-cavvar/movie/proto"
	protoReservation "github.com/ob-vss-ws19/blatt-4-cavvar/reservation/proto"
	protoRoom "github.com/ob-vss-ws19/blatt-4-cavvar/room/proto"
	proto "github.com/ob-vss-ws19/blatt-4-cavvar/showing/proto"
)

type ShowingHandler struct {
	showingDatabase    *sync.Map // "Database". Non Blocking Data Structure or mutex needed to avoid deadlocks
	roomService        protoRoom.RoomService
	reservationService protoReservation.ReservationService
	movieService       protoMovie.MovieService
}

func (handler *ShowingHandler) CreateShowing(ctx context.Context, req *proto.CreateShowingRequest, res *proto.CreateShowingResponse) error {
	// Create showingID
	showingID := fmt.Sprintf("%v %v", req.MovieID, req.RoomName)
	// Find out how much seats are in the given room and use that value to initialize
	roomResponse, err := handler.roomService.FindRoomByID(context.TODO(), &protoRoom.FindRoomByIDRequest{
		RoomName: req.RoomName,
	})
	if err != nil {
		// Communication with room service failed
		return err
	}
	// Create new showing
	// First check if room and movie exist
	_, movieErr := handler.movieService.ReadMovie(context.TODO(), &protoMovie.ReadMovieRequest{MovieID: req.MovieID})
	_, roomErr := handler.roomService.FindRoomByID(context.TODO(), &protoRoom.FindRoomByIDRequest{RoomName: req.RoomName})
	if movieErr != nil {
		// movie not found
		res.Message = "Given movie was not found. Showing won't be created\n"
		return fmt.Errorf("movie %v was not found. showing won't be created", req.MovieID)
	}
	if roomErr != nil {
		// room not found
		res.Message = "Given room was not found. Showing won't be created\n"
		return fmt.Errorf("room %v was not found. showing won't be created", req.RoomName)
	}
	newShowing := &proto.ShowingStruct{
		ShowingID:      showingID,
		MovieID:        req.MovieID,
		RoomName:       req.RoomName,
		AvailableSeats: roomResponse.Room.AvailableSeats,
	}
	// Check if showing already exists
	_, showingExists := handler.showingDatabase.Load(showingID)
	if showingExists {
		// Showing already exists
		log.Printf("The showing %v already exists\n", showingID)
		return errors.New("showing already exists. showing won't be created")
	}
	// Showing doesn't exist. Add it to the "database"
	handler.showingDatabase.Store(showingID, newShowing)
	log.Printf("Showing %v was added to database\n", showingID)
	res.Message = fmt.Sprintf("Showing %v was successfully created", showingID)
	res.NewShowing = newShowing
	return nil
}

func (handler *ShowingHandler) ReadShowing(ctx context.Context, req *proto.ReadShowingRequest, res *proto.ReadShowingResponse) error {
	// See if showing exists
	showingTemp, showingExists := handler.showingDatabase.Load(req.ShowingID)
	if !showingExists {
		// Showing does not exists!
		log.Printf("Showing %v was not found!", req.ShowingID)
		return fmt.Errorf("showing %v was not found", req.ShowingID)
	}
	// showing does exist!
	actualShowing := showingTemp.(*proto.ShowingStruct)
	log.Printf("Showing %v was found\n", req.ShowingID)
	res.Response = actualShowing
	res.Message = fmt.Sprintf("Showing %v was found in the database!", req.ShowingID)
	return nil
}

func (handler *ShowingHandler) UpdateShowing(ctx context.Context, req *proto.UpdateShowingRequest, res *proto.UpdateShowingResponse) error {
	// See if showing exists
	showingTemp, showingExists := handler.showingDatabase.Load(req.ShowingID)
	if !showingExists {
		// Showing does not exists!
		log.Printf("Showing %v was not found!", req.ShowingID)
		return fmt.Errorf("showing %v was not found", req.ShowingID)
	}
	actualShowing := showingTemp.(*proto.ShowingStruct)
	// When updating a showing, check that the seat count doesn't get bigger than the actual seat count in the room
	roomResponse, err := handler.roomService.FindRoomByID(context.TODO(), &protoRoom.FindRoomByIDRequest{
		RoomName: actualShowing.RoomName,
	})
	if err != nil {
		// Communication with room service failed
		return err
	}
	if req.UpdatedShowing.AvailableSeats > roomResponse.Room.AvailableSeats {
		// Can't update because the available seat count is bigger than the room actually has available seats!
		return fmt.Errorf("update not possible. New seat count: %v, actual available seat count: %v", req.UpdatedShowing.AvailableSeats, roomResponse.Room.AvailableSeats)
	}
	// Showing exists. Update it!
	// Add checks for greater zero
	newShowing := req.UpdatedShowing
	handler.showingDatabase.Store(req.ShowingID, newShowing)
	res.UpdatedShowing = newShowing
	res.Message = fmt.Sprintf("Showing %v was successfully updated", req.ShowingID)
	return nil
}

func (handler *ShowingHandler) DeleteShowing(ctx context.Context, req *proto.DeleteShowingRequest, res *proto.DeleteShowingResponse) error {
	// See if showing exists
	_, showingExists := handler.showingDatabase.Load(req.ShowingID)
	if !showingExists {
		// Showing does not exists!
		log.Printf("Showing %v was not found!", req.ShowingID)
		return fmt.Errorf("showing %v was not found. Nothing will be deleted", req.ShowingID)
	}
	// Deleting a showing causes deletions in reservations
	_, err := handler.reservationService.DeleteReservationsForShowing(context.TODO(), &protoReservation.DeleteReservationsForShowingRequest{
		ShowingID: req.ShowingID,
	})
	if err != nil {
		// Something went wrong deleting the reservations for specified showing
		return err
	}
	handler.showingDatabase.Delete(req.ShowingID)
	res.Message = fmt.Sprintf("Showing %v was successfully deleted from the database", req.ShowingID)
	return nil
}

func (handler *ShowingHandler) GetAllShowings(ctx context.Context, req *proto.GetAllShowingsRequest, res *proto.GetAllShowingsResponse) error {
	// Create an array of showing struct
	// Iterate through database
	// Add showings to the array of movie struct
	allShowings := make([]*proto.ShowingStruct, 0)
	handler.showingDatabase.Range(func(key, value interface{}) bool {
		switch m := value.(type) {
		case *proto.ShowingStruct:
			allShowings = append(allShowings, m)
			return true
		default:
			return false
		}
	})
	res.AllShowings = allShowings
	return nil
}

func (handler *ShowingHandler) DeleteShowingsForMovie(ctx context.Context, req *proto.DeleteShowingsForMovieRequest, res *proto.DeleteShowingsForMovieResponse) error {
	handler.showingDatabase.Range(func(key, value interface{}) bool {
		switch m := value.(type) {
		case *proto.ShowingStruct:
			if m.MovieID == req.MovieID {
				// For every movie that matches, delete it from database
				err := handler.DeleteShowing(context.TODO(), &proto.DeleteShowingRequest{ShowingID: m.ShowingID}, &proto.DeleteShowingResponse{Message: "Please work!"})
				if err != nil {
					return false
				}
			}
			return true
		default:
			return false
		}
	})
	return nil
}

func (handler *ShowingHandler) DeleteShowingsForRoom(ctx context.Context, req *proto.DeleteShowingsForRoomRequest, res *proto.DeleteShowingsForRoomResponse) error {
	handler.showingDatabase.Range(func(key, value interface{}) bool {
		switch m := value.(type) {
		case *proto.ShowingStruct:
			if m.RoomName == req.RoomName {
				// For every room that matches, delete it from database
				err := handler.DeleteShowing(context.TODO(), &proto.DeleteShowingRequest{ShowingID: m.ShowingID}, &proto.DeleteShowingResponse{Message: "Please work!"})
				if err != nil {
					return false
				}
			}
			return true
		default:
			return false
		}
	})
	return nil
}

func NewShowingHandler(movieService protoMovie.MovieService, roomService protoRoom.RoomService, reservationService protoReservation.ReservationService) *ShowingHandler {
	// Init Method needed to avoid initialization errors
	showingHandler := &ShowingHandler{
		showingDatabase:    &sync.Map{},
		roomService:        roomService,
		reservationService: reservationService,
		movieService:       movieService,
	}
	return showingHandler
}

func main() {
	// Create a new showing service
	showingService := micro.NewService(
		micro.Name("showing"),
	)

	// Initialize service and parse command line flags if available
	showingService.Init()

	// Init service clients
	roomService := protoRoom.NewRoomService("room", showingService.Client())
	reservationService := protoReservation.NewReservationService("reservation", showingService.Client())
	movieService := protoMovie.NewMovieService("movie", showingService.Client())

	// Register handlers
	err := proto.RegisterShowingHandler(showingService.Server(), NewShowingHandler(movieService, roomService, reservationService))
	if err != nil {
		fmt.Printf("Could not register showing handler!")
	}

	// Run the server or print the error
	if err := showingService.Run(); err != nil {
		fmt.Printf("Error is %v\n", err)
	}
}

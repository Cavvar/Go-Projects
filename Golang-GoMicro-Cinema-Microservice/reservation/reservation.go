package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/micro/go-micro"
	proto "github.com/ob-vss-ws19/blatt-4-cavvar/reservation/proto"
	protoShowing "github.com/ob-vss-ws19/blatt-4-cavvar/showing/proto"
)

type reservation struct {
	userID, showingID, reservationID string
	seatCount                        int
}

type ReservationHandler struct {
	reservationDatabase     *sync.Map // "Database". Non Blocking Data Structure or mutex needed to avoid deadlocks
	initReservationDatabase *sync.Map
	showingService          protoShowing.ShowingService
}

func (handler ReservationHandler) InitReservation(ctx context.Context, req *proto.InitReservationRequest, res *proto.InitReservationResponse) error {
	// Create reservation id
	newReservationID := fmt.Sprintf("%v %v %v", req.ShowingID, req.UserID, req.SeatCount)
	// Check if for the showing the amount of seats are available
	showingResponse, err := handler.showingService.ReadShowing(context.TODO(), &protoShowing.ReadShowingRequest{
		ShowingID: req.ShowingID,
	})
	if err != nil {
		// Communication with showing service failed!
		return err
	}
	if showingResponse.Response.AvailableSeats < req.SeatCount {
		// The amount of seats we want is not available!
		return fmt.Errorf("the amount of seats we want %v is not available. Only %v seats are left", req.SeatCount, showingResponse.Response.AvailableSeats)
	}
	// Add this reservationID to the init Map
	handler.initReservationDatabase.Store(newReservationID, "")
	// Decrease the available seat count on showing side
	updatedShowing := &protoShowing.ShowingStruct{
		ShowingID:      showingResponse.Response.ShowingID,
		MovieID:        showingResponse.Response.MovieID,
		RoomName:       showingResponse.Response.RoomName,
		AvailableSeats: showingResponse.Response.AvailableSeats - req.SeatCount,
	}
	_, updateErr := handler.showingService.UpdateShowing(context.TODO(), &protoShowing.UpdateShowingRequest{
		ShowingID:      showingResponse.Response.ShowingID,
		UpdatedShowing: updatedShowing,
	})
	if updateErr != nil {
		return updateErr
	}
	log.Printf("Reservation %v was added to a temporary database. Please confirm your reservation!\n", newReservationID)
	res.Message = "Reservation was added to a temporary database. Please confirm your reservation!"
	return nil
}

func (handler ReservationHandler) CreateReservation(ctx context.Context, req *proto.CreateReservationRequest, res *proto.CreateReservationResponse) error {
	// Check if user agrees to the reservation
	reservationKey := fmt.Sprintf("%v %v %v", req.ShowingID, req.UserID, req.SeatCount)
	if !req.DoReservation {
		// User does not confirm the reservation!
		handler.initReservationDatabase.Delete(reservationKey)
		return fmt.Errorf("user did not confirm the reservation")
	}
	// User confirms the reservation!
	// Check if reservation key matches
	_, ok := handler.initReservationDatabase.Load(reservationKey)
	if !ok {
		// Reservation key does not match
		return fmt.Errorf("reservation key does not match")
	}
	// Create the new reservation
	res.CreatedReservation = &proto.ReservationStruct{
		UserID:        req.UserID,
		ShowingID:     req.ShowingID,
		ReservationID: reservationKey,
		SeatCount:     req.SeatCount,
	}
	// Add to database
	handler.reservationDatabase.Store(reservationKey, res.CreatedReservation)
	handler.initReservationDatabase.Delete(reservationKey)
	res.Message = fmt.Sprintf("Reservation for user %v and the showing %v was created", req.UserID, req.ShowingID)
	return nil
}

func (handler ReservationHandler) ReadReservation(ctx context.Context, req *proto.ReadReservationRequest, res *proto.ReadReservationResponse) error {
	// See if reservation exists
	reservationTemp, reservationExists := handler.reservationDatabase.Load(req.ReservationID)
	if !reservationExists {
		// Reservation does not exist!
		log.Printf("Reservation %v was not found!\n", req.ReservationID)
		return fmt.Errorf("reservation %v was not found", req.ReservationID)
	}
	actualReservation := reservationTemp.(reservation)
	// Reservation was found
	log.Printf("Reservation %v was found!\n", req.ReservationID)
	res.Reservation = &proto.ReservationStruct{
		UserID:        actualReservation.userID,
		ShowingID:     actualReservation.showingID,
		ReservationID: actualReservation.reservationID,
		SeatCount:     int32(actualReservation.seatCount),
	}
	res.Message = fmt.Sprintf("Reservation %v was found in the database!", req.ReservationID)
	return nil
}

func (handler ReservationHandler) UpdateReservation(ctx context.Context, req *proto.UpdateReservationRequest, res *proto.UpdateReservationResponse) error {
	// Reservations can only update its seat count. That means we have to update the seat count in showings accordingly
	// See if reservation exists
	reservationTemp, reservationExists := handler.reservationDatabase.Load(req.ReservationID)
	if !reservationExists {
		// reservation does not exists!
		log.Printf("reservation %v was not found!", req.ReservationID)
		return fmt.Errorf("reservation %v was not found. nothing will be updated", req.ReservationID)
	}
	// See if the corresponding userID fits
	actualReservation := reservationTemp.(reservation)
	if actualReservation.userID != req.UserID {
		// UserID does not match!
		log.Printf("Given userID does not match with the reservation found!")
		return fmt.Errorf("given userID does not match with the reservation found")
	}
	// Reservation exists. Update it!
	actualReservation.seatCount = int(req.SeatCount)
	handler.reservationDatabase.Store(req.ReservationID, actualReservation)
	// Update showing
	res.UpdatedReservation = &proto.ReservationStruct{
		UserID:        actualReservation.userID,
		ShowingID:     actualReservation.showingID,
		ReservationID: actualReservation.reservationID,
		SeatCount:     int32(actualReservation.seatCount),
	}
	res.Message = fmt.Sprintf("Reservation %v was successfully updated", req.ReservationID)
	return nil
}

func (handler ReservationHandler) DeleteReservation(ctx context.Context, req *proto.DeleteReservationRequest, res *proto.DeleteReservationResponse) error {
	// See if reservation exists
	reservationTemp, reservationExists := handler.reservationDatabase.Load(req.ReservationID)
	if !reservationExists {
		// Reservation does not exists!
		log.Printf("Reservation %v was not found!", req.ReservationID)
		return fmt.Errorf("reservation %v was not found. Nothing will be deleted", req.ReservationID)
	}
	actualReservation := reservationTemp.(reservation)
	// Deleting a reservation has to make sure that the available seat count of a showing is updated
	showingResponse, err := handler.showingService.ReadShowing(context.TODO(), &protoShowing.ReadShowingRequest{
		ShowingID: actualReservation.showingID,
	})
	if err != nil {
		// Communication with showing service failed
		return err
	}
	updatedShowing := &protoShowing.ShowingStruct{
		ShowingID:      showingResponse.Response.ShowingID,
		MovieID:        showingResponse.Response.MovieID,
		RoomName:       showingResponse.Response.RoomName,
		AvailableSeats: int32(actualReservation.seatCount) + showingResponse.Response.AvailableSeats,
	}
	_, err = handler.showingService.UpdateShowing(context.TODO(), &protoShowing.UpdateShowingRequest{
		ShowingID:      updatedShowing.ShowingID,
		UpdatedShowing: updatedShowing,
	})
	if err != nil {
		return err
	}
	// Reservation exists. Delete it!
	handler.reservationDatabase.Delete(req.ReservationID)
	res.Message = fmt.Sprintf("Showing %v was successfully deleted from the database", req.ReservationID)
	res.DeletedReservation = &proto.ReservationStruct{
		UserID:        actualReservation.userID,
		ShowingID:     actualReservation.showingID,
		ReservationID: actualReservation.reservationID,
		SeatCount:     int32(actualReservation.seatCount),
	}
	// Returns an error
	return nil
}

func (handler ReservationHandler) GetReservationsForUser(ctx context.Context, req *proto.GetReservationForUserRequest, res *proto.GetReservationForUserResponse) error {
	allReservations := make([]*proto.ReservationStruct, 0)
	handler.reservationDatabase.Range(func(key, value interface{}) bool {
		switch m := value.(type) {
		case *proto.ReservationStruct:
			if m.UserID == req.UserID {
				allReservations = append(allReservations, m)
				return true
			}
			return false
		default:
			return false
		}
	})
	res.AllReservations = allReservations
	return nil
}

func (handler ReservationHandler) DeleteReservationsForShowing(ctx context.Context, req *proto.DeleteReservationsForShowingRequest, res *proto.DeleteReservationsForShowingResponse) error {
	allReservationsForShowing := make([]*proto.ReservationStruct, 0)
	handler.reservationDatabase.Range(func(key, value interface{}) bool {
		switch m := value.(type) {
		case *proto.ReservationStruct:
			if m.ShowingID == req.ShowingID {
				allReservationsForShowing = append(allReservationsForShowing, m)
				return true
			}
			return false
		default:
			return false
		}
	})
	for _, v := range allReservationsForShowing {
		handler.reservationDatabase.Delete(v.ReservationID)
	}
	return nil
}

func NewReservationHandler(showingService protoShowing.ShowingService) *ReservationHandler {
	// Init Method needed to avoid initialization errors
	reservationHandler := &ReservationHandler{
		reservationDatabase:     &sync.Map{},
		initReservationDatabase: &sync.Map{},
		showingService:          showingService,
	}
	return reservationHandler
}

func main() {
	// Create a new reservation service
	reservationService := micro.NewService(
		micro.Name("reservation"),
	)

	// Init showingService
	showingService := protoShowing.NewShowingService("showing", reservationService.Client())
	// Initialize service and parse command line flags if available
	reservationService.Init()

	// Register handlers
	err := proto.RegisterReservationHandler(reservationService.Server(), NewReservationHandler(showingService))
	if err != nil {
		fmt.Printf("Could not register reservation handler!")
	}

	// Run the server or print the error
	if err := reservationService.Run(); err != nil {
		fmt.Printf("Error is %v\n", err)
	}
}

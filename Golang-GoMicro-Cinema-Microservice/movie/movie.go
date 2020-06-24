package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/micro/go-micro"
	proto "github.com/ob-vss-ws19/blatt-4-cavvar/movie/proto"
	protoShowing "github.com/ob-vss-ws19/blatt-4-cavvar/showing/proto"
)

type MovieHandler struct {
	movieDatabase  *sync.Map // "Database". Non Blocking Data Structure or mutex needed to avoid deadlocks
	showingService protoShowing.ShowingService
}

type movie struct {
	movieID string
	name    string
	fsk     int
}

func NewMovieHandler(showingService protoShowing.ShowingService) *MovieHandler {
	// Init Method needed to avoid initialization errors
	movieHandler := &MovieHandler{
		movieDatabase:  &sync.Map{},
		showingService: showingService,
	}
	return movieHandler
}

func (handler *MovieHandler) ReadMovie(ctx context.Context, req *proto.ReadMovieRequest, res *proto.ReadMovieResponse) error {
	// See if the movie actually exists
	movieTemp, ok := handler.movieDatabase.Load(req.MovieID)
	if !ok {
		// Movie was not found!
		log.Printf("Movie was NOT found!\n")
		res.Message = "Movie was NOT found!"
		res.Movie = nil
		return fmt.Errorf("movie %v was NOT found", req.MovieID)
	}
	actualMovie := movieTemp.(movie)
	// Movie was found!
	log.Printf("Movie %v was found!\n", actualMovie)
	res.Movie = &proto.MovieStruct{
		Id:   actualMovie.movieID,
		Name: actualMovie.name,
		Fsk:  int32(actualMovie.fsk),
	}
	res.Message = fmt.Sprintf("Movie %v exists in the current database!", actualMovie.movieID)
	return nil
}

func (handler *MovieHandler) CreateMovie(ctx context.Context, req *proto.CreateMovieRequest, res *proto.CreateMovieResponse) error {
	// Create new ID
	newID := fmt.Sprintf("%v %v", req.Name, req.Fsk)
	// Add to database
	newMovie := movie{
		movieID: newID,
		name:    req.Name,
		fsk:     int(req.Fsk),
	}
	_, doesMovieAlreadyExist := handler.movieDatabase.Load(newID)
	if doesMovieAlreadyExist {
		log.Printf("Movie %v does already exist in the database!\n", newID)
		res.Message = fmt.Sprintf("Movie %v does already exist in the database. Nothing will be added!", newID)
		return errors.New("nothing will be added because the movie already exists")
	}
	handler.movieDatabase.Store(newID, newMovie)
	log.Printf("Movie %v was added to the database!\n", newID)
	// Respond with a message
	res.Message = fmt.Sprintf("Movie %v was successfully added to the database\n", newMovie.name)
	res.NewMovie = &proto.MovieStruct{
		Id:   newMovie.movieID,
		Name: newMovie.name,
		Fsk:  int32(newMovie.fsk),
	}
	return nil
}

func (handler *MovieHandler) DeleteMovie(ctx context.Context, req *proto.DeleteMovieRequest, res *proto.DeleteMovieResponse) error {
	// See if the movie actually exists
	movieTemp, ok := handler.movieDatabase.Load(req.MovieID)
	if !ok {
		// Movie does not exist
		// Set Error Return Value
		log.Printf("The requested movie %v for deletion does NOT exist in the database!\n", req.MovieID)
		return fmt.Errorf("movie with %v does NOT exist! Nothing will be deleted", req.MovieID)
	}
	actualMovie := movieTemp.(movie)
	// Movie does exist
	// Remove that movie
	// Don't forget that if a movie gets deleted it causes a 'chain reaction'
	_, err := handler.showingService.DeleteShowingsForMovie(context.TODO(), &protoShowing.DeleteShowingsForMovieRequest{MovieID: req.MovieID})
	if err != nil {
		return err
	}
	// Send Success Message back
	log.Printf("The requested movie for deletion does exist in the database! Movie %v will be deleted!\n", actualMovie.name)
	handler.movieDatabase.Delete(req.MovieID)
	res.Message = fmt.Sprintf("Movie %v was deleted from the movie database!\n", actualMovie.movieID)

	return nil
}

func (handler *MovieHandler) GetAllMovies(ctx context.Context, req *proto.GetAllMoviesRequest, res *proto.GetAllMoviesResponse) error {
	// Create an array of movie struct
	// Iterate through database
	// Add movies to the array of movie struct
	allMovies := make([]*proto.MovieStruct, 0)
	handler.movieDatabase.Range(func(key, value interface{}) bool {
		switch m := value.(type) {
		case *proto.MovieStruct:
			allMovies = append(allMovies, m)
			return true
		default:
			return false
		}
	})
	res.AllMovies = allMovies
	return nil
}

func main() {
	// Create a new movie service
	movieService := micro.NewService(
		micro.Name("movie"),
	)

	// Init showing service
	showingService := protoShowing.NewShowingService("showing", movieService.Client())

	// Initialize service and parse command line flags if available
	movieService.Init()

	// Register handlers
	err := proto.RegisterMovieHandler(movieService.Server(), NewMovieHandler(showingService))
	if err != nil {
		fmt.Printf("Could not register movie handler!")
	}

	// Run the server or print the error
	if err := movieService.Run(); err != nil {
		fmt.Printf("Error is %v\n", err)
	}
}

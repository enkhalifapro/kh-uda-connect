package internal

import (
	"context"
	"database/sql"
	"encoding/binary"
	"encoding/hex"
	"enkhalifapro/connections/pb"
	"fmt"
	"github.com/paulmach/go.geo"
	"github.com/segmentio/kafka-go"
	"time"
)

type Service struct {
	dbConnector     DBConnector
	kafkaConnector  KafkaConnector
	locationService LocationService
}

// KafkaConnector interface for kafka message consuming operations
type KafkaConnector interface {
	ReadMessage(ctx context.Context) (kafka.Message, error)
}

type LocationService interface {
	GetLocationsByDate(date time.Time) (*pb.LocationsList, error)
}

func NewService(db DBConnector, kafkaConn KafkaConnector, locationSvc LocationService) *Service {
	return &Service{
		dbConnector:     db,
		kafkaConnector:  kafkaConn,
		locationService: locationSvc,
	}
}

// DBConnector interface for database operations
type DBConnector interface {
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
	MustExec(query string, args ...interface{}) sql.Result
}

// AddPersonConnections reads all available locations by day
// and add them as connections to person
func (s *Service) AddPersonConnections(locationEvent *LocationAddedEvent) error {
	// get locations by date using grpc
	currentTime := time.Now()
	locations, err := s.locationService.GetLocationsByDate(currentTime)
	if err != nil {
		return err
	}

	// add connections
	for _, location := range locations.Locations {
		personLocation, _ := wkbStringToPoint(locationEvent.Coordinate)
		connectionLocation, _ := wkbStringToPoint(location.Coordinate)
		distance := personLocation.DistanceFrom(&connectionLocation)
		query := fmt.Sprintf(`INSERT INTO public.connections(person_id, person_location, connection_id, connection_location, distance,creation_time) values ('%s','%s','%s', '%s',%v)`, locationEvent.PersonID, locationEvent.Coordinate, location.PersonId, location.Coordinate, distance, currentTime)
		res := s.dbConnector.MustExec(query)
		affRows, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if int(affRows) == 0 {
			return fmt.Errorf("person saving failure")
		}
	}

	return nil
}

func wkbStringToPoint(wkb string) (geo.Point, error) {
	// Decode the hex string to bytes
	wkbBytes, err := hex.DecodeString(wkb)
	if err != nil {
		return geo.Point{}, err
	}

	// Check the byte order (little-endian)
	if wkbBytes[0] != 0x01 {
		return geo.Point{}, fmt.Errorf("unsupported WKB type")
	}

	// Extract the coordinates (8 bytes for each coordinate)
	x := binary.LittleEndian.Uint64(wkbBytes[5:13])
	y := binary.LittleEndian.Uint64(wkbBytes[13:21])

	// Convert to float64
	point := geo.Point{}

	point.SetX(float64(x))
	point.SetY(float64(y))

	return point, nil
}

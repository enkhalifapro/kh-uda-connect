package internal

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/segmentio/kafka-go"
)

type Service struct {
	dbConnector    DBConnector
	kafkaConnector KafkaConnector
}

// KafkaConnector interface for kafka message consuming operations
type KafkaConnector interface {
	ReadMessage(ctx context.Context) (kafka.Message, error)
}

func NewService(db DBConnector, kafkaConn KafkaConnector) *Service {
	return &Service{
		dbConnector:    db,
		kafkaConnector: kafkaConn,
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
	var conns []Connection
	// todo: get locations by time using grpc

	// todo: add connections
	for _, conn := range conns {
		query := fmt.Sprintf(`INSERT INTO public.connections(person_id, person_location, connection_id, connection_location,creation_time) values ('%s','%s','%s', '%s',%v)`, conn.PersonID, conn.PersonLocation, conn.ConnectionID, conn.ConnectionLocation, conn.CreationTime)
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

package internal

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type Service struct {
	dbConnector    DBConnector
	kafkaConnector KafkaConnector
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

// KafkaConnector interface for kafka message producing operations
type KafkaConnector interface {
	WriteMessages(msgs ...kafka.Message) (int, error)
}

// Add a new person
func (s *Service) Add(location *CreatePayload) error {
	query := fmt.Sprintf(`insert into public.location (person_id, coordinate) values (%v, '%s');`, location.PersonID, location.Coordinate)
	res := s.dbConnector.MustExec(query)
	affRows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if int(affRows) == 0 {
		return fmt.Errorf("location saving failure")
	}
	
	err = s.PublishLocationAddedEvent(location)
	return err
}

func (s *Service) PublishLocationAddedEvent(location *CreatePayload) error {
	id, _ := uuid.NewUUID()
	msg, err := json.Marshal(location)
	if err != nil {
		return err
	}
	_, err = s.kafkaConnector.WriteMessages(
		kafka.Message{
			Key:   []byte(id.String()),
			Value: msg,
		},
	)
	return err
}

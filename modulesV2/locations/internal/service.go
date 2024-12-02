package internal

import (
	"database/sql"
	"fmt"
)

type Service struct {
	dbConnector DBConnector
}

func NewService(db DBConnector) *Service {
	return &Service{
		dbConnector: db,
	}
}

// DBConnector interface for database operations
type DBConnector interface {
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
	MustExec(query string, args ...interface{}) sql.Result
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
	return nil
}

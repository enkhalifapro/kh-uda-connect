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
func (s *Service) Add(person *CreatePayload) error {
	query := fmt.Sprintf(`INSERT INTO public.person(first_name, last_name,company_name) values ('%s','%s','%s')`, person.FirstName, person.LastName, person.CompanyName)
	res := s.dbConnector.MustExec(query)
	affRows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if int(affRows) == 0 {
		return fmt.Errorf("person saving failure")
	}
	return nil
}

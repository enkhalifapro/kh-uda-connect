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

// GetAll persons
func (s *Service) GetAll() ([]Person, error) {
	var res []Person
	err := s.dbConnector.Select(&res, `SELECT * FROM public.person;`)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetByID persons
func (s *Service) GetByID(id int) (*Person, error) {
	var res []Person
	err := s.dbConnector.Select(&res, fmt.Sprintf(`SELECT * FROM public.person WHERE ID=%v limit 1;`, id))
	if err != nil {
		return nil, err
	}
	if len(res) > 0 {
		return &res[0], nil
	}
	return nil, fmt.Errorf("person not found")
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

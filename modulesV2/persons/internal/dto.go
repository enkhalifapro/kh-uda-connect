package internal

import "time"

type Person struct {
	ID          int    `json:"id"`
	FirstName   string `json:"firstName" db:"first_name"`
	LastName    string `json:"lastName" db:"last_name"`
	CompanyName string `json:"companyName" db:"company_name"`
}

type CreatePayload struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	CompanyName string `json:"companyName"`
}

type HealthResponse struct {
	Message string `json:"message"`
}

type Connection struct {
	ConnectionID       int       `json:"connectionId" db:"connection_id"`
	ConnectionLocation string    `json:"connection_location" db:"connection_location"`
	Distance           float64   `json:"distance" db:"distance"`
	CreationTime       time.Time `json:"creationTime" db:"creation_time"`
}

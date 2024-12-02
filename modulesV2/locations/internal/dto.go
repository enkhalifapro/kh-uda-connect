package internal

type Person struct {
	ID          int    `json:"id"`
	FirstName   string `json:"firstName" db:"first_name"`
	LastName    string `json:"lastName" db:"last_name"`
	CompanyName string `json:"companyName" db:"company_name"`
}

type Location struct {
	ID         int    `json:"id"`
	PersonID   int    `json:"personId" db:"person_id"`
	Coordinate string `json:"coordinate" db:"coordinate"`
}

type CreatePayload struct {
	PersonID   int    `json:"personId" db:"person_id"`
	Coordinate string `json:"coordinate" db:"coordinate"`
}

type HealthResponse struct {
	Message string `json:"message"`
}

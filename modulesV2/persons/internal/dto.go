package internal

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

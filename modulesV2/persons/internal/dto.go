package internal

type CreatePayload struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	CompanyName string `json:"companyName"`
}

type HealthResponse struct {
	Message string `json:"message"`
}

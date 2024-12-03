package internal

import "time"

type HealthResponse struct {
	Message string `json:"message"`
}

type LocationAddedEvent struct {
	PersonID   int
	Coordinate string
}

type Connection struct {
	PersonID           int
	PersonLocation     string
	ConnectionID       int
	ConnectionLocation string
	CreationTime       time.Time
}

type Location struct {
	ID         int
	PersonID   int
	Coordinate string
}

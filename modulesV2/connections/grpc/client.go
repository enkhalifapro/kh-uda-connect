package grpc

import (
	"context"
	"enkhalifapro/connections/pb"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"log"
	"time"
)

type LocationsClient struct {
	client *grpc.ClientConn
}

func NewLocationsClient(client *grpc.ClientConn) *LocationsClient {
	return &LocationsClient{
		client: client,
	}
}

func (l *LocationsClient) GetLocationsByDate(date time.Time) (*pb.LocationsList, error) {
	// Set up a connection to the server.
	conn, err := grpc.NewClient("localhost:5051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewLocationsServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Second) // Increase to 5 seconds
	defer cancel()
	t := &timestamp.Timestamp{
		Seconds: date.Unix(),
		Nanos:   int32(date.Nanosecond()),
	}
	req := &pb.LocationsRequest{CreatedAt: t}

	return client.GetLocations(ctx, req)
}

package grpc

import (
	"context"
	"enkhalifapro/locations/internal"
	pb "enkhalifapro/locations/pb" // Update this to the correct import path
)

type LocationServer struct {
	pb.UnimplementedLocationsServiceServer
	service Service
}

type Service interface {
	GetLocationsByDate(day string) ([]internal.Location, error)
}

func NewLocationServer() *LocationServer {
	return &LocationServer{}
}

func (s *LocationServer) GetLocations(ctx context.Context, req *pb.LocationsRequest) (*pb.LocationsList, error) {
	locations, err := s.service.GetLocationsByDate(req.CreatedAt.AsTime().Format("2006-01-02"))
	if err != nil {
		return nil, err
	}
	res := &pb.LocationsList{}
	locs := make([]*pb.Location, len(locations))
	for _, l := range locations {
		locs = append(locs, &pb.Location{
			Id:         int32(l.ID),
			PersonId:   int32(l.PersonID),
			Coordinate: l.Coordinate,
		})
	}
	res.Locations = locs

	return res, nil
}

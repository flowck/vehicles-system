package grpc_service

import (
	"context"
	"time"
	pb "vehicles-system/api"
	"vehicles-system/internal/app"
	"vehicles-system/internal/domain/vehicle"
)

type vehicleService struct {
	application *app.App
	pb.UnimplementedVehicleServiceServer
}

func (s vehicleService) GetVehicles(ctx context.Context, req *pb.GetVehiclesRequest) (*pb.GetVehiclesResponse, error) {
	vehicles, err := s.application.VehicleService.GetVehicles(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.GetVehiclesResponse{Vehicles: mapToRpcVehicles(vehicles)}, nil
}

func (s vehicleService) GetVehicle(ctx context.Context, req *pb.GetVehicleRequest) (*pb.GetVehicleResponse, error) {
	result, err := s.application.VehicleService.GetVehicle(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.GetVehicleResponse{Vehicle: mapToRpcVehicle(result)}, nil
}

func (s vehicleService) StreamVehicles(req *pb.GetVehiclesRequest, stream pb.VehicleService_StreamVehiclesServer) error {
	vehicles, err := s.application.VehicleService.GetVehicles(stream.Context())
	if err != nil {
		return nil
	}

	for _, item := range vehicles {
		stream.Send(&pb.StreamVehiclesResponse{Vehicle: mapToRpcVehicle(&item)})
		time.Sleep(time.Millisecond * 500) // Not so fast dude
	}

	return nil
}

// Helpers
func mapToRpcModel(model string) *pb.Model {
	return &pb.Model{Name: model}
}

func mapToRpcModels(models []string) []*pb.Model {
	result := make([]*pb.Model, len(models))

	for idx, model := range models {
		result[idx] = mapToRpcModel(model)
	}

	return result
}

func mapToRpcVehicle(item *vehicle.Vehicle) *pb.Vehicle {
	return &pb.Vehicle{
		Id:     item.Id(),
		Name:   item.Name(),
		Models: mapToRpcModels(item.Models()),
	}
}

func mapToRpcVehicles(vehicles []vehicle.Vehicle) []*pb.Vehicle {
	result := make([]*pb.Vehicle, len(vehicles))

	for idx, item := range vehicles {
		result[idx] = mapToRpcVehicle(&item)
	}

	return result
}

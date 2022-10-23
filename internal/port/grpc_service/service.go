package grpc_service

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "vehicles-system/api"
	"vehicles-system/internal/app"
)

type grpcService struct {
	server         *grpc.Server
	vehicleService vehicleService
}

func NewService(application *app.App) *grpcService {
	return &grpcService{vehicleService: vehicleService{application: application}}
}

func (s *grpcService) Start(port int) {
	// Create grpc server and register it to proto's grpcService
	s.server = grpc.NewServer(grpc.EmptyServerOption{})
	pb.RegisterVehicleServiceServer(s.server, s.vehicleService)

	// Create a tcp listener
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("could not create a tpc listener: %v", err)
	}

	// Server
	if err = s.server.Serve(listener); err != nil {
		log.Fatalf("could not server: %v", err)
	}
}

func (s *grpcService) Stop() {
	fmt.Println("Stopping grpcService...")
	s.server.GracefulStop() // Does it stop really? Shouldn't I pass a context with timeout?
	fmt.Println("grpcService stopped gracefully")
}

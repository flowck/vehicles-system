package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
	pb "vehicles-system/api"
)

func printGetVehiclesResponse(data *pb.GetVehiclesResponse) {
	for idx, item := range data.Vehicles {
		fmt.Printf("%d - id: %s | name: %s\n", idx, item.Id, item.Name)
	}
}

func main() {
	fmt.Println("Simple client for testing stuff")

	// Connect to the server
	conn, err := grpc.Dial("localhost:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to the server: %v", err)
	}
	defer conn.Close()
	service := pb.NewVehicleServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var data *pb.GetVehiclesResponse
	if data, err = service.GetVehicles(ctx, &pb.GetVehiclesRequest{}); err != nil {
		log.Fatalf("Could not get vehicles: %v", err)
	}

	printGetVehiclesResponse(data)
}

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"vehicles-system/internal/app"
	"vehicles-system/internal/ports/grpc_service"
)

func main() {
	application := app.NewApp()
	service := grpc_service.NewService(application)
	defer service.Stop()
	// Listen for termination signals from the os
	done := make(chan os.Signal)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	go func() {
		service.Start(3000)
	}()

	<-done
	fmt.Println("\nprocess has been terminated")
}

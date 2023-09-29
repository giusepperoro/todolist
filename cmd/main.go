package main

import (
	"context"
	config "github.com/giusepperoro/todolist.git/intenal/config"
	"github.com/giusepperoro/todolist.git/intenal/database"
	"github.com/giusepperoro/todolist.git/intenal/repository"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	config, err := config.GetConfigFromFile("config.yaml")
	if err != nil {
		return err
	}

	dbConn, err := database.New(ctx, config)

	repository.NewAssociationRepository(dbConn)

	return nil
}

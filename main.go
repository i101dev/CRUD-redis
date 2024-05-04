package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/i101dev/microservices-NN/application"
)

func main() {

	cfg := application.LoadConfig()
	app := application.New(cfg)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	defer cancel()

	if err := app.Start(ctx); err != nil {
		fmt.Println("failed to start app: ", err)
	}
}

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/i101dev/microservices-NN/application"
)

func main() {
	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	defer cancel()

	if err := app.Start(ctx); err != nil {
		fmt.Println("failed to start app: ", err)
	}
}

package main

import (
	"context"
	"fmt"

	"github.com/i101dev/microservices-NN/application"
)

func main() {
	app := application.New()
	if err := app.Start(context.TODO()); err != nil {
		fmt.Println("failed to start app: ", err)
	}
}

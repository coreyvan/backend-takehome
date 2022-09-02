package main

import (
	"go.uber.org/zap"
	
	"github.com/coreyvan/backend-takehome/internal/app"
)

func main() {
	log, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	if err := app.Run(log); err != nil {
		log.Sugar().Fatalf("running app: %v", err)
	}
}

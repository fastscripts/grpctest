package main

import (
	"fmt"
	"os"

	"github.com/fastscripts/grpctest/internal/config"
	"github.com/fastscripts/grpctest/internal/database"
	"github.com/fastscripts/grpctest/internal/database/seeder"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run cmd/cli/main.go seed")
		os.Exit(1)
	}
	cmd := os.Args[1]

	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	switch cmd {
	case "seed":
		if len(os.Args) > 2 {
			cfg.Database.DSN = os.Args[2]
		}

		dbService, err := database.NewDatabaseService(cfg)
		if err != nil {
			fmt.Printf("Error initializing database: %v\n", err)
			os.Exit(1)
		}

		defer dbService.Close()

		err = seeder.RunAllSeeders(seeder.Opts{
			DB: dbService.DB(),
		})
		if err != nil {
			fmt.Printf("Error running seeders: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Printf("Unknown command: %s\n", cmd)
		os.Exit(1)
	}

}

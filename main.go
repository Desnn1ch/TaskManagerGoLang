package main

import (
	"TaskManagerGoLang/config"
	"TaskManagerGoLang/database"
	"TaskManagerGoLang/handlers"
	"log"
)

func main() {
	cfg, err := config.LoadDataBaseConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := database.NewPostgresDb(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	router := handlers.SetupRouter(db)

	router.Run(":8080")
}

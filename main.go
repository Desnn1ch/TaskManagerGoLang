package main

import (
	"TaskManagerGoLang/config"
	"TaskManagerGoLang/database"
	"TaskManagerGoLang/handlers"

	"fmt"
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

	redisCfg, err := config.LoadRedisConfig()
	if err != nil {
		log.Fatalf("Failed to load Redis config: %v", err)
	}
  
	rdb, err := database.NewRedisClient(redisCfg)

	if err != nil {
		log.Fatalf("failed to connect to redis: %s", err)
	}
	router := handlers.SetupRouter(db, rdb)

	router.Run(":8080")
}
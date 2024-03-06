package main

import (
	"log"
	"os"
	"untitled13/internal/config"
	"untitled13/internal/database"
)

func main() {
	cfg := config.Load()

	db, err := database.InitDb(cfg.Database)
	if err != nil {
		log.Print("failed to init db ", err)
		os.Exit(1)
	}
	_ = db
}

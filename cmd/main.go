package main

import (
	"goticket/internal/config"
	// "goticket/internal/event"
	"goticket/internal/server"
	// "log"
)

func main() {
	// Load environment variables
	cfg := config.LoadEnv()

	// Connect database
	db := config.ConnectDatabase(cfg)

		// Automatically create/update the events table
	// if err := db.AutoMigrate(&event.Event{}); err != nil {
	// 	log.Fatal("Database migration failed:", err)
	// }

	// log.Println("Database migration completed successfully")

	// Start server
	server.Start(db, cfg)
}
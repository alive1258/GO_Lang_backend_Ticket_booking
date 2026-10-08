package main

import (
	"goticket/internal/config"
	"goticket/internal/server"
)

func main() {
	// Load environment variables
	cfg := config.LoadEnv()

	// Connect database
	db := config.ConnectDatabase(cfg)

	// Start server
	server.Start(db, cfg)
}
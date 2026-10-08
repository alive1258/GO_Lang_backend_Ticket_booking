package main

import (
	"goticket/internal/config"
	"goticket/internal/server"
)


func main() {
    // load env
   cfg :=  config.LoadEnv()
   // connect db
   db := config.ConnectDatabase(cfg)
  // start server
    server.Start( db, cfg)
  
}
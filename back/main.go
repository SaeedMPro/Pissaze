package main

import (
	"github.com/pissaze/internal/server"
	_ "github.com/pissaze/internal/storage"
)

// @title Pisaz
// @version 1.0
// @description API for managing products, and users in a hardware shopping site.
// @contact.name Saeed, Danny
// @contact.email dan
// @host localhost:8082
// // @BasePath /api
func main(){
	server.Start()
}
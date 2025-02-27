package main

import (
	"github.com/pissaze/internal/server"
	_ "github.com/pissaze/internal/storage"
)

// @title Pissaze
// @version 1.0
// @description API for managing products, and users in a hardware shopping site.
// @contact.name Saeed, Danny
// @contact.email mzahry36@gmail.com, dankeshavarz1075@gmail.com
// @host localhost:8082
// @BasePath /api
func main(){
	server.Start()
}
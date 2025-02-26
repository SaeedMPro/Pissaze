package main

import (
	"github.com/pissaze/internal/server"
	_ "github.com/pissaze/internal/storage"
)


func main(){
	server.Start()
}
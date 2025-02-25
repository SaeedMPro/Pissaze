package main

import "github.com/pissaze/internal/storage"

func main(){
	err := storage.InitDB()
	if err != nil {
		panic(err)
	}
}
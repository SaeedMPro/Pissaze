package main

import "pisaz/internal/storage"

func main(){
	err := storage.InitDB()
	if err != nil {
		panic(err)
	}
}
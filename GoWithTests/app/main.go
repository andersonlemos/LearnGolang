package main

import (
	"app/server"
	"fmt"
	"log"
	"net/http"
	"os"
)

const PORT string = "5000"
const DbFilename = "game.db.json"

func main() {
	db, err := os.OpenFile(DbFilename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("Cannot open DB file %s %v", DbFilename, err)
	}

	/*storage := &server.PlayerStorageFileSystem{db}*/
	/*server := server.NewPlayerServer(storage)*/
	storage, err := server.NewPlayerStorageFileSystem(db)
	if err != nil {
		log.Fatalf("Cannot create DB file %s %v", DbFilename, err)
	}
	server := server.NewPlayerServer(&storage)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", PORT), db; err != nil {
		log.Fatalf("was not able to listen on port 5000 %v", err)
	}
}

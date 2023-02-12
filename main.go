package main

import (
	"go-livestream/server"
	"log"
	"net/http"
)

func main() {
	server.AllRooms.Init()

	http.Handle("/create", http.HandlerFunc(server.CreateRoomRequestHandler))
	http.Handle("/join", http.HandlerFunc(server.JoinRoomRequestHandler))

	log.Println("Starting server on port 8000")

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatal("Server startup error: ", err)
	}
}

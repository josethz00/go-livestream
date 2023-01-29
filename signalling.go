package server

import (
	"log"
	"net/http"
)

func CreateRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateRoomRequestHandler")
}

func JoinRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("JoinRoomRequestHandler")
}

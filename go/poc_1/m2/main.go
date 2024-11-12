package main

import (
	"log"
	"net/http"

	user "m2/ogen/user"
)

func main() {
	// Create service instance.
	service := &usersService{
		users: map[int]user.User{},
	}
	// Create generated server.
	srv, err := user.NewServer(service)
	if err != nil {
		log.Fatal(err)
	}
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}

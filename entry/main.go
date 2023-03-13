package main

import (
	"OnlineDeck/entry/server"
	"OnlineDeck/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {

	// gin engine
	r := gin.Default()

	rv1 := r.Group("/api/v1/")
	// injecting an instance of deck map which serves as an inmemory database
	// can be replaced by a real database instance
	decks := make(map[uuid.UUID]*models.Deck)
	s := server.NewServer(rv1, decks)

	//register all endpoints
	s.RegisterAll()

	//run the server on port 3000
	err := r.Run(":3000")
	if err != nil {
		return
	}

}

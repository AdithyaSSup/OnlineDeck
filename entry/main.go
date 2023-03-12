package main

import (
	"OnlineDeck/entry/server"
	"github.com/gin-gonic/gin"
)

func main() {

	// gin engine
	r := gin.Default()

	rv1 := r.Group("/api/v1/")
	s := server.NewServer(rv1)

	//register all endpoints
	s.RegisterAll()

	//run the server on port 3000
	err := r.Run(":3000")
	if err != nil {
		return
	}

}

package server

import (
	"OnlineDeck/entry/inject"
	"github.com/gin-gonic/gin"
)

type Server struct {
	g *gin.RouterGroup
	//further fields can be added here to manage configuration dependencies
}

func NewServer(g *gin.RouterGroup) *Server {
	return &Server{
		g: g,
	}
}

func (s *Server) RegisterAll() {
	s.registerDeckEndpoints()
}

func (s *Server) registerDeckEndpoints() {
	//rg := s.g.Group("/casino")
	// all required dependencies for controller are injected here
	dc := inject.DeckController()

	s.g.POST("/decks", dc.CreateDeck)
	s.g.GET("/decks/:id", dc.Open)
	s.g.GET("/decks/:id/draw", dc.DrawCards)
}

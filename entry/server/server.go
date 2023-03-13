package server

import (
	"OnlineDeck/entry/inject"
	"OnlineDeck/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Server struct {
	g       *gin.RouterGroup
	deckMap map[uuid.UUID]*models.Deck
	//further fields can be added here to manage configuration dependencies
	// or replace memory based deck map to real database instance
}

func NewServer(g *gin.RouterGroup, dm map[uuid.UUID]*models.Deck) *Server {
	return &Server{
		g:       g,
		deckMap: dm,
	}
}

func (s *Server) RegisterAll() {
	s.registerDeckEndpoints()
}

func (s *Server) registerDeckEndpoints() {
	// all required dependencies for controller are injected here
	dc := inject.DeckController(s.deckMap)
	// all deck endpoints with their controller methods are mapped here
	s.g.POST("/decks", dc.CreateDeck)
	s.g.GET("/decks/:id", dc.OpenDeck)
	s.g.GET("/decks/:id/draw", dc.DrawCards)
}

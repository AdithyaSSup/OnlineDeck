// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package inject

import (
	"OnlineDeck/entry/controller/decks"
	"OnlineDeck/pkg/dao"
	"OnlineDeck/pkg/models"
	"OnlineDeck/pkg/services/deck"
	"github.com/google/uuid"
)

// Injectors from wire.go:

func DeckController(arg map[uuid.UUID]*models.Deck) *decks.Controller {
	deckDao := dao.NewDeckDao(arg)
	service := deck.NewService(deckDao)
	controller := decks.NewDeckController(service)
	return controller
}

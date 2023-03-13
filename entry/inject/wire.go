//go:build wireinject
// +build wireinject

package inject

import (
	"OnlineDeck/entry/controller/decks"
	"OnlineDeck/pkg/dao"
	"OnlineDeck/pkg/models"
	"OnlineDeck/pkg/services/deck"
	"github.com/google/uuid"
	"github.com/google/wire"
)

func DeckController(map[uuid.UUID]*models.Deck) *decks.Controller {
	panic(wire.Build(
		wire.Bind(new(decks.DeckService), new(*deck.Service)), deck.NewService,
		wire.Bind(new(deck.DeckDao), new(*dao.DeckDao)), dao.NewDeckDao,
		decks.NewDeckController))
}

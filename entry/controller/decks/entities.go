package decks

import (
	"OnlineDeck/pkg/models"
	"github.com/google/uuid"
)

type CreateDeckRequest struct {
	Shuffled bool `json:"shuffle,omitempty"`
}

type DrawCardRequest struct {
	Number int `form:"count" binding:"gte=1"`
}

type CreateDeckResponse struct {
	ID        uuid.UUID `json:"id"`
	Shuffle   bool      `json:"shuffled"`
	Remaining int       `json:"remaining"`
}

type OpenDeckResponse struct {
	CreateDeckResponse
	Cards []models.Card `json:"cards"`
}

type DrawCardsResponse struct {
	Cards []models.Card `json:"cards"`
}

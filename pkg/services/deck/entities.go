package deck

import "OnlineDeck/pkg/models"

type omit *struct{}

// RequestDTO's
type CreateDeckRequestDTO struct {
	Shuffled  bool
	CardNames []string
}

type DrawCardRequestDTO struct {
	DeckID string
	Number int
}

type OpenDeckRequestDTO struct {
	DeckID string
}

// Response DTO's
type DeckResponseDTO struct {
	models.Deck
	RemainingCards int `json:"remaining"`
}

type CreateDeckResponseDTO struct {
	DeckResponseDTO
	Cards omit `json:"cards,omitempty"`
}

type OpenDeckResponseDTO struct {
	Deck models.Deck
}

type DrawCardResponseDTO struct {
	Cards []models.Card `json:"cards"`
}

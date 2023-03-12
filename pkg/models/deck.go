package models

import "github.com/google/uuid"

type Deck struct {
	ID       uuid.UUID `json:"id"`
	Shuffled bool      `json:"shuffled"`
	Cards    []Card    `json:"cards"`
}

package models

import "github.com/google/uuid"

type Deck struct {
	ID       uuid.UUID
	Shuffled bool
	Cards    []Card
}

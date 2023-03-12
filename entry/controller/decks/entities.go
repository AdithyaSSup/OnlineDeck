package decks

type CreateDeckRequest struct {
	Shuffled bool `json:"shuffle,omitempty"`
}

type DrawCardRequest struct {
	Number int `form:"count" binding:"gte=1"`
}

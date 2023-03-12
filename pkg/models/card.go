package models

type Card struct {
	Suit  CardSuit `json:"suit"`
	Value CardRank `json:"value"`
	Code  string   `json:"code"`
}

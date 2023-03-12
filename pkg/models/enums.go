package models

type CardSuit string

const (
	SuitDiamonds CardSuit = "DIAMONDS"
	SuitHearts   CardSuit = "HEARTS"
	SuitClubs    CardSuit = "CLUBS"
	SuitSpades   CardSuit = "SPADES"
)

type CardRank string

const (
	RankAce   CardRank = "ACE"
	RankOne   CardRank = "1"
	RankTwo   CardRank = "2"
	RankThree CardRank = "3"
	RankFour  CardRank = "4"
	RankFive  CardRank = "5"
	RankSix   CardRank = "6"
	RankSeven CardRank = "7"
	RankEight CardRank = "8"
	RankNine  CardRank = "9"
	RankTen   CardRank = "10"
	RankJack  CardRank = "JACK"
	RankQueen CardRank = "QUEEN"
	RankKing  CardRank = "KING"
)

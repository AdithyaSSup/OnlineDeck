package deck

import (
	"OnlineDeck/pkg/models"
	"context"
	"errors"
	"fmt"
)

var ErrInvalidCardName = errors.New("invalid Card Name")
var ErrInvalidCardValue = errors.New("invalid Card Value")
var ErrInvalidCardSuit = errors.New("invalid Card Suit")

type DeckDao interface {
	Create(ctx context.Context, cards []models.Card, shuffle bool) (*models.Deck, error)
	Draw(ctx context.Context, deck *models.Deck, count int) ([]models.Card, error)
	Get(ctx context.Context, deckID string) (*models.Deck, error)
	Shuffle(ctx context.Context, deck *models.Deck) (*models.Deck, error)
}

type Service struct {
	DeckDao DeckDao
}

func NewService(dao DeckDao) *Service {
	return &Service{
		DeckDao: dao,
	}
}

func (ds *Service) Create(ctx context.Context, req CreateDeckRequestDTO) (*CreateDeckResponseDTO, error) {

	cards, err := ds.GetCards(ctx, req.CardNames)
	if err != nil {
		return nil, err
	}
	res, err := ds.DeckDao.Create(ctx, cards, req.Shuffled)

	if err != nil {
		return nil, err
	}

	return &CreateDeckResponseDTO{
		DeckResponseDTO: DeckResponseDTO{
			Deck:           *res,
			RemainingCards: len(res.Cards),
		},
	}, err
}

func (ds *Service) Open(ctx context.Context, deckID string) (*DeckResponseDTO, error) {
	resp, err := ds.DeckDao.Get(ctx, deckID)
	if err != nil {
		return nil, err
	}
	return &DeckResponseDTO{
		Deck:           *resp,
		RemainingCards: len(resp.Cards),
	}, err
}

func (ds *Service) DrawCard(ctx context.Context, req DrawCardRequestDTO) (*DrawCardResponseDTO, error) {
	deck, err := ds.DeckDao.Get(ctx, req.DeckID)
	if err != nil {
		return nil, err
	}
	cards, err := ds.DeckDao.Draw(ctx, deck, req.Number)
	if err != nil {
		return nil, err
	}

	return &DrawCardResponseDTO{Cards: cards}, err
}

func (ds *Service) GetCards(ctx context.Context, cardNames []string) ([]models.Card, error) {
	var cards []models.Card

	// If cardNames are provided, create a partial deck
	if len(cardNames) > 0 {
		cardRankMap, cardSuitMap := ds.GetCardSuitRankMap()
		// Generate a partial deck of cards based on the given card names
		for _, name := range cardNames {
			card := models.Card{}
			if len(name) != 2 {
				// Invalid card name
				fmt.Println("Invalid card name passed", name)
				return nil, ErrInvalidCardName
			}
			value, valueExists := cardRankMap[name[0]]
			suit, suitExists := cardSuitMap[name[1]]
			if !valueExists {
				// Invalid card value
				return nil, ErrInvalidCardValue
			}
			if !suitExists {
				// Invalid card suit
				return nil, ErrInvalidCardSuit
			}
			card.Value = value
			card.Suit = suit
			card.Code = name
			cards = append(cards, card)
		}
	} else {
		// Generate a full deck of cards as card names aren't provided
		cards = GetAllCards()
	}
	return cards, nil
}

func (ds *Service) GetCardSuitRankMap() (map[byte]models.CardRank, map[byte]models.CardSuit) {
	cardRankMap := map[byte]models.CardRank{
		'A': models.RankAce,
		'2': models.RankTwo,
		'3': models.RankThree,
		'4': models.RankFour,
		'5': models.RankFive,
		'6': models.RankSix,
		'7': models.RankSeven,
		'8': models.RankEight,
		'9': models.RankNine,
		'0': models.RankTen,
		'J': models.RankJack,
		'Q': models.RankQueen,
		'K': models.RankKing,
	}
	cardSuitMap := map[byte]models.CardSuit{
		'H': models.SuitHearts,
		'C': models.SuitClubs,
		'D': models.SuitDiamonds,
		'S': models.SuitSpades,
	}
	return cardRankMap, cardSuitMap
}

func getAllSuits() []models.CardSuit {
	return []models.CardSuit{models.SuitClubs, models.SuitDiamonds, models.SuitHearts, models.SuitSpades}
}

func getAllRanks() []models.CardRank {
	return []models.CardRank{models.RankAce, models.RankTwo, models.RankThree, models.RankFour, models.RankFive,
		models.RankSix, models.RankSeven, models.RankEight, models.RankNine, models.RankTen, models.RankJack,
		models.RankQueen, models.RankKing}
}

func GetAllCards() []models.Card {
	var cards []models.Card
	for _, suit := range getAllSuits() {
		for _, rank := range getAllRanks() {
			card := models.Card{
				Value: rank,
				Suit:  suit,
				Code:  fmt.Sprintf("%c%c", rank[0], suit[0]),
			}
			cards = append(cards, card)
		}
	}
	return cards
}

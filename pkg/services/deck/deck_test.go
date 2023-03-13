package deck

import (
	"OnlineDeck/pkg/dao"
	"OnlineDeck/pkg/models"
	"context"
	"github.com/google/uuid"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	createFunctionTestCase1 = "test - create method with valid valid card names for creation of partial deck "
	createFunctionTestCase2 = "test - create method with no input cards names for complete deck creation"
	createFunctionTestCase3 = "test - create method with invalid card name"
	createFunctionTestCase4 = "test - create method with invalid card suit name"
	createFunctionTestCase5 = "test - create method with invalid card rank name"
	OpenFunctionTestCase1   = "test - open method with valid uuid"
	OpenFunctionTestCase2   = "test - open method with invalid uuid"
	OpenFunctionTestCase3   = "test - open method with valid uuid but no deck against it"
)

var (
	uuidOpenFunctionTestCase1 = uuid.New()
	uuidOpenFunctionTestCase2 = "not-a-valid-uuid"
	uuidOpenFunctionTestCase3 = uuid.New()
)

// Mock implementation of the DeckDao interface
type MockDeckDao struct {
	mock.Mock
}

func (m *MockDeckDao) Create(ctx context.Context, cards []models.Card, shuffle bool) (*models.Deck, error) {
	args := m.Called(ctx, cards, shuffle)
	return args.Get(0).(*models.Deck), args.Error(1)
}

func (m *MockDeckDao) Draw(ctx context.Context, deck *models.Deck, count int) ([]models.Card, error) {
	args := m.Called(ctx, deck, count)
	return args.Get(0).([]models.Card), args.Error(1)
}

func (m *MockDeckDao) Get(ctx context.Context, deckID string) (*models.Deck, error) {
	args := m.Called(ctx, deckID)
	return args.Get(0).(*models.Deck), args.Error(1)
}

func (m *MockDeckDao) Shuffle(ctx context.Context, deck *models.Deck) (*models.Deck, error) {
	args := m.Called(ctx, deck)
	return args.Get(0).(*models.Deck), args.Error(1)
}

func TestCreate(t *testing.T) {
	testCases := []struct {
		name        string
		cardNames   []string
		shuffled    bool
		expectedErr error
	}{
		{
			name:      createFunctionTestCase1,
			cardNames: []string{"AH", "KH", "1D", "JS"},
			shuffled:  true,
		},
		{
			name:      createFunctionTestCase2,
			cardNames: []string{},
			shuffled:  false,
		},
		{
			name:        createFunctionTestCase3,
			cardNames:   []string{"Aced", "King", "Queen", "Jack"},
			shuffled:    true,
			expectedErr: ErrInvalidCardName,
		}, {
			name:        createFunctionTestCase4,
			cardNames:   []string{"AM"},
			shuffled:    true,
			expectedErr: ErrInvalidCardSuit,
		}, {
			name:        createFunctionTestCase5,
			cardNames:   []string{"ZS"},
			shuffled:    true,
			expectedErr: ErrInvalidCardValue,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockDeckDao := new(MockDeckDao)
			service := NewService(mockDeckDao)

			ctx := context.Background()
			req := CreateDeckRequestDTO{
				CardNames: tc.cardNames,
				Shuffled:  tc.shuffled,
			}
			var expectedDeck *models.Deck
			// Mock the DeckDao's Create method
			if tc.expectedErr == nil {

				if tc.name == createFunctionTestCase1 {
					expectedDeck = &models.Deck{
						ID: uuid.New(),
						Cards: []models.Card{
							{Suit: models.SuitHearts,
								Value: models.RankAce,
								Code:  "AH"},
							{Suit: models.SuitHearts,
								Value: models.RankKing,
								Code:  "KH"},
							{Suit: models.SuitDiamonds,
								Value: models.RankTen,
								Code:  "1D"},
							{Suit: models.SuitSpades,
								Value: models.RankJack,
								Code:  "JS"},
						},
					}
				} else {
					expectedDeck = &models.Deck{
						ID:    uuid.New(),
						Cards: GetAllCards(),
					}
				}
				mockDeckDao.On("Create", ctx, expectedDeck.Cards, req.Shuffled).Return(expectedDeck, nil)
			} else {
				mockDeckDao.On("Create", ctx, mock.Anything, true).Return(nil, tc.expectedErr)
			}

			// Test
			actualRes, actualErr := service.Create(ctx, req)

			// Assert
			if tc.expectedErr != nil {
				assert.EqualError(t, actualErr, tc.expectedErr.Error())
			} else {
				assert.NoError(t, actualErr)
				assert.Equal(t, len(expectedDeck.Cards), actualRes.RemainingCards)
				assert.Equal(t, expectedDeck.ID, actualRes.ID)
				mockDeckDao.AssertExpectations(t)
			}
		})
	}
}

func TestOpen(t *testing.T) {
	testCases := []struct {
		name        string
		deckID      string
		expectedErr error
	}{
		{
			name:   OpenFunctionTestCase1,
			deckID: uuidOpenFunctionTestCase1.String(),
		},
		{
			name:        OpenFunctionTestCase2,
			deckID:      uuidOpenFunctionTestCase2,
			expectedErr: dao.ErrInvalidUUID,
		},
		{
			name:        OpenFunctionTestCase3,
			deckID:      uuidOpenFunctionTestCase3.String(),
			expectedErr: dao.ErrInvalidUUID,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockDeckDao := new(MockDeckDao)
			service := NewService(mockDeckDao)

			ctx := context.Background()
			req := OpenDeckRequestDTO{
				Id: tc.deckID,
			}
			var expectedDeck *models.Deck
			// Mock the DeckDao's Create method
			if tc.expectedErr == nil {

				if tc.name == OpenFunctionTestCase1 {
					expectedDeck = &models.Deck{
						ID: uuidOpenFunctionTestCase1,
						Cards: []models.Card{
							{Suit: models.SuitHearts,
								Value: models.RankAce,
								Code:  "AH"},
							{Suit: models.SuitHearts,
								Value: models.RankKing,
								Code:  "KH"},
							{Suit: models.SuitDiamonds,
								Value: models.RankTen,
								Code:  "1D"},
							{Suit: models.SuitSpades,
								Value: models.RankJack,
								Code:  "JS"},
						},
					}
				}
				mockDeckDao.On("Get", ctx, expectedDeck.ID.String()).Return(expectedDeck, nil)
			} else {
				mockDeckDao.On("Get", ctx, mock.Anything).Return(expectedDeck, tc.expectedErr)
			}

			// Test
			actualRes, actualErr := service.Open(ctx, req)

			// Assert
			if tc.expectedErr != nil {
				assert.EqualError(t, actualErr, tc.expectedErr.Error())
			} else {
				assert.NoError(t, actualErr)
				assert.Equal(t, len(expectedDeck.Cards), actualRes.RemainingCards)
				assert.Equal(t, expectedDeck.ID, actualRes.ID)
				mockDeckDao.AssertExpectations(t)
			}
		})
	}
}

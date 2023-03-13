package decks

import (
	"OnlineDeck/pkg/dao"
	error2 "OnlineDeck/pkg/errors"
	"OnlineDeck/pkg/models"
	"OnlineDeck/pkg/services/deck"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

var ErrInvalidReqParam = errors.New("invalid request parameter")

type DeckService interface {
	Create(ctx context.Context, req deck.CreateDeckRequestDTO) (*deck.CreateDeckResponseDTO, error)
	Open(ctx context.Context, req deck.OpenDeckRequestDTO) (*deck.DeckResponseDTO, error)
	DrawCard(ctx context.Context, req deck.DrawCardRequestDTO) (*deck.DrawCardResponseDTO, error)
}

type Controller struct {
	DeckService DeckService
}

func NewDeckController(deckService DeckService) *Controller {
	return &Controller{DeckService: deckService}
}

func (d *Controller) CreateDeck(c *gin.Context) {

	var (
		cards             []string
		createDeckRequest CreateDeckRequest
	)

	ctx := c.Request.Context()

	//check if anything is passed in the Query Param, if passed then
	//those will contain cards names separated by "," to include in partial deck

	cardsStr := c.Query("cards")
	if cardsStr != "" {
		cards = strings.Split(cardsStr, ",")
	}

	// Check and assign body parameter available ,
	//as of now its shuffle status only , can be extended to include other params for future use cases
	if !models.BindRequestBody(c, &createDeckRequest) {
		return
	}

	resp, err := d.DeckService.Create(ctx, deck.CreateDeckRequestDTO{
		Shuffled:  createDeckRequest.Shuffled,
		CardNames: cards,
	})
	if err != nil {
		d.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, CreateDeckResponse{
		ID:        resp.ID,
		Shuffle:   resp.Shuffled,
		Remaining: resp.RemainingCards,
	})
}

func (d *Controller) OpenDeck(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")

	res, err := d.DeckService.Open(ctx, deck.OpenDeckRequestDTO{Id: id})

	if err != nil {
		d.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, OpenDeckResponse{
		CreateDeckResponse: CreateDeckResponse{
			ID:        res.ID,
			Shuffle:   res.Shuffled,
			Remaining: res.RemainingCards,
		},
		Cards: res.Cards,
	})

}

func (d *Controller) DrawCards(c *gin.Context) {
	var req DrawCardRequest
	ctx := c.Request.Context()

	id := c.Param("id")

	if err := c.Bind(&req); err != nil {
		d.handleError(c, ErrInvalidReqParam)
		return
	}

	res, err := d.DeckService.DrawCard(ctx, deck.DrawCardRequestDTO{
		DeckID: id,
		Number: req.Number,
	})

	if err != nil {
		d.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, DrawCardsResponse{Cards: res.Cards})

}

// handleError: Method to handle all the errors returned from service layer
// it checks the errors defined in service layer and then sends the response to the client
// Rest format
func (d *Controller) handleError(c *gin.Context, err error) {

	switch errors.Cause(err) {
	case dao.ErrInvalidDraw:
		c.JSON(http.StatusBadRequest, error2.HttpError{
			Type:   "INVALID_COUNT",
			Title:  "invalid draw count",
			Detail: err.Error(),
		})
	case dao.ErrUUIDGeneration:
		c.JSON(http.StatusInternalServerError, error2.HttpError{
			Type:   "INTERNAL_SERVER_ERROR",
			Title:  "unique Id generation failure",
			Detail: err.Error(),
		})
	case dao.ErrDeckNotFound:
		c.JSON(http.StatusNotFound, error2.HttpError{
			Type:   "INVALID_RESOURCE_ID",
			Title:  "resource not found",
			Detail: err.Error(),
		})
	case deck.ErrInvalidCardSuit:
		c.JSON(http.StatusBadRequest, error2.HttpError{
			Type:   "INVALID_SUIT_ID",
			Title:  "invalid suit",
			Detail: err.Error(),
		})
	case deck.ErrInvalidCardValue:
		c.JSON(http.StatusBadRequest, error2.HttpError{
			Type:   "INVALID_CARD_VALUE",
			Title:  "invalid card value",
			Detail: err.Error(),
		})
	case dao.ErrInvalidUUID:
		c.JSON(http.StatusBadRequest, error2.HttpError{
			Type:   "INVALID_RESOURCE_ID",
			Title:  "invalid uuid",
			Detail: err.Error(),
		})
	case deck.ErrInvalidCardName:
		c.JSON(http.StatusBadRequest, error2.HttpError{
			Type:   "INVALID_CARD_NAME",
			Title:  "invalid card name",
			Detail: err.Error(),
		})
	case ErrInvalidReqParam:
		c.JSON(http.StatusBadRequest, error2.HttpError{
			Type:   "INVALID_REQUEST",
			Title:  "invalid request",
			Detail: err.Error(),
		})
	default:
		c.JSON(http.StatusInternalServerError, error2.HttpError{
			Type:   "INTERNAL_SERVER_ERROR",
			Title:  "internal server error",
			Detail: err.Error(),
		})
	}
}

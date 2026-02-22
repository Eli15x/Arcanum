package handler

import (
	"Arcanum/internal/core/ports"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ...existing code...
type DeckHandler struct {
	service ports.DeckService
}

func NewDeckHandler(s ports.DeckService) *DeckHandler {
	return &DeckHandler{service: s}
}

// GetRandomCards -> GET /draw/:deck_type?count=#
func (h *DeckHandler) GetRandomCards(c *gin.Context) {
	deckType := c.Param("deck_type")
	count, _ := strconv.Atoi(c.DefaultQuery("count", "1"))

	cards, err := h.service.DrawCards(deckType, count)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cards)
}

// GetComplexSpread -> GET /spread/:spread_type?deck_type=slug
func (h *DeckHandler) GetComplexSpread(c *gin.Context) {
	spreadType := c.Param("spread_type")
	deckType := c.DefaultQuery("deck_type", "tarot")

	spread, err := h.service.GetComplexSpread(spreadType, deckType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, spread)
}

// GetComplexSpred -> POST /spread/:spread_type with body { "deck_type": "slug" }
func (h *DeckHandler) GetComplexSpred(c *gin.Context) {
	spreadType := c.Param("spread_type")
	var payload struct {
		DeckType string `json:"deck_type"`
	}
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}
	if payload.DeckType == "" {
		payload.DeckType = "tarot"
	}

	spread, err := h.service.GetComplexSpread(spreadType, payload.DeckType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, spread)
}

func (h *DeckHandler) handleDeckBySlug(c *gin.Context, slug string) {
	cards, err := h.service.DrawCards(slug, 1)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cards)
}

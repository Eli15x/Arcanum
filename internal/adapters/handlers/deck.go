package handlers

import (
    "arcanum/internal/core/ports"
    "github.com/gin-gonic/gin"
    "strconv"
)

type HTTPHandler struct {
    service ports.DeckService
}

func (h *HTTPHandler) Draw(c *gin.Context) {
    deckType := c.Param("deck_type")
    count, _ := strconv.Atoi(c.DefaultQuery("count", "1"))

    cards, err := h.service.DrawCards(deckType, count)
    if err != nil {
        c.JSON(404, gin.H{"error": "Baralho não encontrado"})
        return
    }
    c.JSON(200, cards)
}
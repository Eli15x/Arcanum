package service

import (
	"Arcanum/internal/core/domain"
	"Arcanum/internal/core/ports"
	"math/rand"
	"time"
)

// ...existing code...
type deckService struct {
	repo ports.DeckRepository
}

func NewDeckService(r ports.DeckRepository) ports.DeckService {
	return &deckService{repo: r}
}

func (s *deckService) DrawCards(deckType string, count int) ([]domain.Card, error) {
	deck, err := s.repo.GetBySlug(deckType)
	if err != nil {
		return nil, err
	}

	if count <= 0 {
		return nil, nil
	}
	if count > len(deck.Cards) {
		count = len(deck.Cards)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck.Cards), func(i, j int) {
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	})

	return deck.Cards[:count], nil
}

// GetComplexSpread implements [ports.DeckService].
func (s *deckService) GetComplexSpread(spreadType string, deckType string) (map[string]domain.Card, error) {
	return nil, nil
}

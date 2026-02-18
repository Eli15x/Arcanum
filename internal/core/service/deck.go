package services

import (
    "arcanum/internal/core/domain"
    "arcanum/internal/core/ports"
    "math/rand"
    "time"
)

type deckService struct {
    repo ports.DeckRepository
}

func NewDeckService(r ports.DeckRepository) ports.DeckService {
    return &deckService{repo: r}
}

func (s *deckService) DrawCards(deckType string, count int) ([]domain.Card, error) {
    deck, err := s.repo.GetDeckByType(deckType)
    if err != nil { return nil, err }

    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(deck.Cards), func(i, j int) {
        deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
    })

    return deck.Cards[:count], nil
}
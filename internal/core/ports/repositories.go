package ports

import "arcanum/internal/core/domain"

type DeckRepository interface {
    GetDeckByType(deckType string) (domain.Deck, error)
}


type DeckService interface {
    DrawCards(deckType string, count int) ([]domain.Card, error)
    GetComplexSpread(spreadType string, deckType string) (map[string]domain.Card, error)
}
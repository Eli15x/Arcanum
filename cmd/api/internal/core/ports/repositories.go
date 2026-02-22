package ports

import "Arcanum/internal/core/domain"


type DeckRepository interface {
    Save(deck domain.Deck) error
    GetBySlug(slug string) (domain.Deck, error)
}


type DeckService interface {
    DrawCards(deckType string, count int) ([]domain.Card, error)
    GetComplexSpread(spreadType string, deckType string) (map[string]domain.Card, error)
}
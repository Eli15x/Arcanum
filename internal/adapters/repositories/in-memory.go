// internal/adapters/repositories/memory_repository.go
package repositories

import (
    "sync"
    "arcanum/internal/core/domain"
    "errors"
)

type MemoryDeckRepository struct {
    // Usamos RWMutex para garantir que a leitura seja segura se houver concorrência
    mu    sync.RWMutex
    decks map[string]domain.Deck
}

func NewMemoryDeckRepository() *MemoryDeckRepository {
    return &MemoryDeckRepository{
        decks: make(map[string]domain.Deck),
    }
}

func (r *MemoryDeckRepository) Save(deck domain.Deck) {
    r.mu.Lock()
    defer r.mu.Unlock()
    r.decks[deck.Slug] = deck
}

func (r *MemoryDeckRepository) GetDeckByType(deckType string) (domain.Deck, error) {
    r.mu.RLock()
    defer r.mu.RUnlock()
    
    deck, ok := r.decks[deckType]
    if !ok {
        return domain.Deck{}, errors.New("baralho não encontrado")
    }
    return deck, nil
}
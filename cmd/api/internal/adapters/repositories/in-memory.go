package repositories

import (
	"Arcanum/internal/core/domain"
	"Arcanum/internal/core/ports"
	"errors"
	"sync"
)

var ErrDeckNotFound = errors.New("baralho não encontrado")

type MemoryDeckRepository struct {
	mu    sync.RWMutex
	decks map[string]domain.Deck
}

func NewMemoryDeckRepository() ports.DeckRepository {
	return &MemoryDeckRepository{
		decks: make(map[string]domain.Deck),
	}
}

func (r *MemoryDeckRepository) Save(deck domain.Deck) error {
	if deck.Slug == "" {
		return errors.New("deck slug vazio")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	r.decks[deck.Slug] = deck
	return nil
}

func (r *MemoryDeckRepository) GetBySlug(slug string) (domain.Deck, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	deck, ok := r.decks[slug]
	if !ok {
		return domain.Deck{}, ErrDeckNotFound
	}
	return deck, nil
}

// Asserção de compilação: garante que MemoryDeckRepository implementa ports.DeckRepository
var _ ports.DeckRepository = (*MemoryDeckRepository)(nil)

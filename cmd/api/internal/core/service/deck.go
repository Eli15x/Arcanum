package service

import (
	"Arcanum/internal/core/domain"
	"Arcanum/internal/core/ports"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// ...existing code...
type deckService struct {
	repo ports.DeckRepository
}

func NewDeckService(r ports.DeckRepository) ports.DeckService {
	return &deckService{repo: r}
}

// package-level random source (seed once)
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

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

	cards := make([]domain.Card, len(deck.Cards))
	copy(cards, deck.Cards)

	rnd.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})

	isCigano := strings.EqualFold(deck.Slug, "cigano")
	if isCigano {
		return s.drawCIgano(cards, count), nil
	}

	return cards[:count], nil
}

func (s *deckService) drawCIgano(cards []domain.Card, count int) []domain.Card {
	drawn := cards
	if count < len(cards) {
		drawn = cards[:count]
	}

	result := make([]domain.Card, 0, (len(drawn)+1)/2)
	counter := 1

	for i := 0; i < len(drawn); i += 2 {
		first := drawn[i]

		if i+1 < len(drawn) {
			second := drawn[i+1]
			key := strconv.Itoa(second.ID)

			combined := domain.Card{
				ID:           counter,
				Name:         first.Name + " + " + second.Name,
				Description:  "",
				Deck:         first.Deck,
				ImageURL:     first.ImageURL,
				Combinations: nil,
			}

			combined.CombinedIDs = strconv.Itoa(first.ID) + " + " + strconv.Itoa(second.ID)

			if comb, ok := first.Combinations[key]; ok && comb != "" {
				combined.Description = comb
			} else {
				if first.Description != "" && second.Description != "" {
					combined.Description = first.Description + " / " + second.Description
				} else if first.Description != "" {
					combined.Description = first.Description
				} else {
					combined.Description = second.Description
				}
			}

			result = append(result, combined)
			counter++
		} else {
			single := domain.Card{
				ID:           first.ID,
				Name:         first.Name,
				Description:  first.Description,
				Deck:         first.Deck,
				ImageURL:     first.ImageURL,
				Combinations: nil,
			}
			result = append(result, single)
		}
	}

	return result
}

func (s *deckService) GetComplexSpread(spreadType string, deckType string) (map[string]domain.Card, error) {
	return nil, nil
}

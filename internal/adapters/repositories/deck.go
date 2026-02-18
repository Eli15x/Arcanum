package repositories

import (
    "encoding/json"
    "os"
    "arcanum/internal/core/domain"
)

type JSONDeckRepository struct {
    basePath string
}

func (r *JSONDeckRepository) GetDeckByType(deckType string) (domain.Deck, error) {
    file, err := os.ReadFile(r.basePath + "/" + deckType + ".json")
    if err != nil { return domain.Deck{}, err }

    var deck domain.Deck
    if err := json.Unmarshal(file, &deck); err != nil { return domain.Deck{}, err }
    
    return deck, nil
}
// internal/core/domain/card.go
package domain

type Card struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
	Deck 		string `json:"deck"`
    ImageURL    string `json:"image_url"`
}

type Deck struct {
    Name  string `json:"name"`
    Cards []Card `json:"cards"`
}
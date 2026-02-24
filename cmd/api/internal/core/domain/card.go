// internal/core/domain/card.go
package domain

type Card struct {
	ID           int               `json:"id"`
	Name         string            `json:"name"`
	Description  string            `json:"description"`
	Deck         string            `json:"deck"`
	ImageURL     string            `json:"image_url"`
	Combinations map[string]string `json:"combinations,omitempty"`
	CombinedIDs  string            `json:"combined_ids,omitempty"`
}

type SpecialSequence struct {
	IDs            []int  `json:"ids"`
	Interpretation string `json:"interpretation"`
}

type Deck struct {
	Slug             string            `json:"slug"`
	Name             string            `json:"deck_name"`
	Cards            []Card            `json:"cards"`
	SpecialSequences []SpecialSequence `json:"special_sequences"`
}

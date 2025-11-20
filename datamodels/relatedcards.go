package datamodels

// RelatedCards contains info about cards related to the current card (e.g. other faces, meld parts, and cards in a card's spellbook)
type RelatedCards struct {
	ReverseRelated []string `json:"reverseRelated,omitempty"`
	Spellbook      []string `json:"spellbook,omitempty"`
}

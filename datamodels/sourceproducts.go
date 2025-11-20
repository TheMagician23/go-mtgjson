package datamodels

// SourceProducts contains the UUIDs for the Sealed Products that a card is found in
type SourceProducts struct {
	Etched  []string `json:"etched,omitempty"`
	Foil    []string `json:"foil,omitempty"`
	NonFoil []string `json:"nonfoil,omitempty"`
}

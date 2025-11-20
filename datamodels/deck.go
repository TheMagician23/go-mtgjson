package datamodels

type CardSetDeck struct {
	Count  int    `json:"count"`
	IsFoil bool   `json:"isFoil,omitempty"`
	Uuid   string `json:"uuid"`
}

type Deck struct {
	Code               string        `json:"code"`
	Commander          []CardSetDeck `json:"commander,omitempty"`
	MainBoard          []CardSetDeck `json:"mainBoard"`
	Name               string        `json:"name"`
	ReleaseDate        string        `json:"releaseDate,omitempty"`
	SealedProductUuids []string      `json:"sealedProductUuids,omitempty"`
	SideBoard          []CardSetDeck `json:"sideBoard"`
	Type               string        `json:"type"`
}

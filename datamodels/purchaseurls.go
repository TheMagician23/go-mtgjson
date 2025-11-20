package datamodels

// PurchaseUrls describes a set of links to purchase a product on certain marketplaces.
type PurchaseUrls struct {
	CardKingdom       string `json:"cardKingdom,omitempty"`
	CardKingdomEtched string `json:"cardKingdomEtched,omitempty"`
	CardKingdomFoil   string `json:"cardKingdomFoil,omitempty"`
	Cardmarket        string `json:"cardmarket,omitempty"`
	TcgPlayer         string `json:"tcgplayer,omitempty"`
	TcgPlayerEtched   string `json:"tcgplayerEtched,omitempty"`
}

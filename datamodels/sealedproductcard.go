package datamodels

// SealedProductCard describes the card product configuration in SealedProductContents.
type SealedProductCard struct {
	Foil   bool   `json:"foil,omitempty"`
	Name   string `json:"name"`
	Number string `json:"number"`
	Set    string `json:"set"`
	Uuid   string `json:"uuid"`
}

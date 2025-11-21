package datamodels

// SealedProductContents describes the properties for the purchaseable product of a Set.
type SealedProductContents struct {
	Card     []SealedProductCard                  `json:"card,omitempty"`
	Deck     []SealedProductDeck                  `json:"deck,omitempty"`
	Other    []SealedProductOther                 `json:"other,omitempty"`
	Pack     []SealedProductPack                  `json:"pack,omitempty"`
	Sealed   []SealedProductSealed                `json:"sealed,omitempty"`
	Variable []map[string][]SealedProductContents `json:"variable,omitempty"`
}

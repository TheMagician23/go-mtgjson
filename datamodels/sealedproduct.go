package datamodels

// SealedProduct describes the properties of a purchaseable product of a Set.
type SealedProduct struct {
	CardCount    int                   `json:"cardCount"`
	Category     string                `json:"category,omitempty"`
	Contents     SealedProductContents `json:"contents"`
	Identifiers  Identifiers           `json:"identifiers"`
	Name         string                `json:"name"`
	ProductSize  int                   `json:"productSize,omitempty"`
	PurchaseUrls PurchaseUrls          `json:"purchaseUrls"`
	ReleaseDate  string                `json:"releaseDate,omitempty"`
	Subtype      string                `json:"subtype,omitempty"`
	Uuid         string                `json:"uuid"`
}

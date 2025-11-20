package mtgjson

// Based on MTGJSON v5.5.2+20251117

type SetListMeta struct {
	Date    string `json:"date"`
	Version string `json:"version"`
}

// Set has information about the card set
type Set struct {
	BaseSetSize        int               `json:"baseSetSize"`
	Block              string            `json:"block,omitempty"`
	Booster            []interface{}     `json:"booster"`
	Cards              []CardSet         `json:"cards,omitempty"`
	CardsphereSetId    int               `json:"cardsphereSetId,omitempty"`
	Code               string            `json:"code"`
	CodeV3             string            `json:"codeV3,omitempty"`
	Decks              []Deck            `json:"decks"`
	IsForeignOnly      bool              `json:"isForeignOnly,omitempty"`
	IsFoilOnly         bool              `json:"isFoilOnly,omitempty"`
	IsNonFoilOnly      bool              `json:"isNonFoilOnly,omitempty"`
	IsOnlineOnly       bool              `json:"isOnlineOnly,omitempty"`
	IsPaperOnly        bool              `json:"isPaperOnly,omitempty"`
	IsPartialPreview   bool              `json:"isPartialPreview,omitempty"`
	KeyruneCode        string            `json:"keyruneCode,omitempty"`
	Languages          []string          `json:"languages"`
	CardmarketId       int               `json:"mcmId,omitempty"`
	CardmarketIdExtras int               `json:"mcmIdExtras,omitempty"`
	CardmarketName     string            `json:"mcmName,omitempty"`
	MTGOCode           string            `json:"mtgoCode,omitempty"`
	Name               string            `json:"name"`
	ParentCode         string            `json:"parentCode,omitempty"`
	ReleaseDate        string            `json:"releaseDate"`
	SealedProduct      []SealedProduct   `json:"sealedProduct,omitempty"`
	TCGPlayerGroupID   int               `json:"tcgplayerGroupId,omitempty"`
	Tokens             []Token           `json:"tokens,omitempty"`
	TokenSetCode       string            `json:"tokenSetCode,omitempty"`
	TotalSetSize       int               `json:"totalSetSize"`
	Translations       map[string]string `json:"translations,omitempty"`
	Type               string            `json:"type"`
}

type SetList struct {
	Meta SetListMeta    `json:"meta"`
	Sets map[string]Set `json:"data"`
}

package mtgjson

type SealedProductCard struct {
	Foil   bool   `json:"foil,omitempty"`
	Name   string `json:"name"`
	Number string `json:"number"`
	Set    string `json:"set"`
	Uuid   string `json:"uuid"`
}

type SealedProductDeck struct {
	Name string `json:"name"`
	Set  string `json:"set"`
}

type SealedProductOther struct {
	Name string `json:"name"`
}

type SealedProductPack struct {
	Code string `json:"code"`
	Set  string `json:"set"`
}

type SealedProductSealed struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Set   string `json:"set"`
	Uuid  string `json:"uuid"`
}

type SealedProductContents struct {
	Card     []SealedProductCard              `json:"card,omitempty"`
	Deck     []SealedProductDeck              `json:"deck,omitempty"`
	Other    []SealedProductOther             `json:"other,omitempty"`
	Pack     []SealedProductPack              `json:"pack,omitempty"`
	Sealed   []SealedProductSealed            `json:"sealed,omitempty"`
	Variable map[string]SealedProductContents `json:"variable,omitempty"`
}

type Identifiers struct {
	AbuGamesId             string `json:"abuId,omitempty"`
	CardKingdomEtchedId    string `json:"cardKingdomEtchedId,omitempty"`
	CardKingdomFoilId      string `json:"cardKingdomFoilId,omitempty"`
	CardKingdomId          string `json:"cardKingdomId,omitempty"`
	CardsphereId           string `json:"cardsphereId,omitempty"`
	CardsphereFoilId       string `json:"cardsphereFoilId,omitempty"`
	CardtraderId           string `json:"cardtraderId,omitempty"`
	CoolStuffIncId         string `json:"csiId,omitempty"`
	CardmarketId           string `json:"mcmId,omitempty"`
	CardmarketMetaId       string `json:"mcmMetaId,omitempty"`
	MiniatureMarketId      string `json:"miniaturemarketId,omitempty"`
	MtgArenaId             string `json:"mtgArenaId,omitempty"`
	MtgJsonFoilId          string `json:"mtgjsonFoilVersionId,omitempty"`
	MtgJsonNonFoilId       string `json:"mtgjsonNonFoilVersionId,omitempty"`
	MtgJsonV4Id            string `json:"mtgjsonV4Id,omitempty"`
	MtgoFoilId             string `json:"mtgoFoilId,omitempty"`
	MtgoId                 string `json:"mtgoId,omitempty"`
	MultiverseId           string `json:"multiverseId,omitempty"`
	StarCityGamesId        string `json:"scgId,omitempty"`
	ScryfallId             string `json:"scryfallId,omitempty"`
	ScryfallCardBackId     string `json:"scryfallCardBackId,omitempty"`
	ScryfallOracleId       string `json:"scryfallOracleId,omitempty"`
	ScryfallIllustrationId string `json:"scryfallIllustrationId,omitempty"`
	TcgPlayerId            string `json:"tcgplayerProductId,omitempty"`
	TcgPlayerEtchedId      string `json:"tcgplayerEtchedProductId,omitempty"`
	TrollAndToadId         string `json:"tntId,omitempty"`
}

type PurchaseUrls struct {
	CardKingdom       string `json:"cardKingdom,omitempty"`
	CardKingdomEtched string `json:"cardKingdomEtched,omitempty"`
	CardKingdomFoil   string `json:"cardKingdomFoil,omitempty"`
	Cardmarket        string `json:"cardmarket,omitempty"`
	TcgPlayer         string `json:"tcgplayer,omitempty"`
	TcgPlayerEtched   string `json:"tcgplayerEtched,omitempty"`
}

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

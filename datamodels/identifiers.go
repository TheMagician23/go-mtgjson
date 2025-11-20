package datamodels

// Identifiers describes the IDs associated to a specific card or product on various marketplaces.
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

package datamodels

// BoosterConfig describes the properties of how a Set's booster data may be configured.
type BoosterConfig struct {
	Boosters            []BoosterPack           `json:"boosters"`
	BoostersTotalWeight float32                 `json:"boostersTotalWeight"`
	Name                string                  `json:"name,omitempty"`
	Sheets              map[string]BoosterSheet `json:"sheets"`
	SourceSetCodes      []string                `json:"sourceSetCodes"`
}

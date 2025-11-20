package datamodels

// BoosterPack describes the properties of how a booster pack can be configured.
type BoosterPack struct {
	Contents map[string]float32 `json:"contents,omitempty"`
	Weight   float32            `json:"weight"`
}

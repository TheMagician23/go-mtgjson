package datamodels

// BoosterSheet describes the properties of how a sheet of printed cards can be configured.
type BoosterSheet struct {
	AllowDuplicates bool               `json:"allowDuplicates,omitempty"`
	BalanceColors   bool               `json:"balanceColors,omitempty"`
	Cards           map[string]float32 `json:"cards"`
	Foil            bool               `json:"foil"`
	Fixed           bool               `json:"fixed,omitempty"`
	TotalWeight     float32            `json:"totalWeight"`
}

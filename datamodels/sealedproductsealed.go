package datamodels

// SealedProductSealed descrubes the sealed product configuration in SealedProductContents.
type SealedProductSealed struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Set   string `json:"set"`
	Uuid  string `json:"uuid"`
}

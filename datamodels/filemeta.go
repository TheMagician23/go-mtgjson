package datamodels

// Based on MTGJSON v5.5.2+20251117
// FileMeta has info about when a file was generated
type FileMeta struct {
	Date    string `json:"date"`
	Version string `json:"version"`
}

package datamodels

// ForeignData describes the properties for a card in alternate languages
type ForeignData struct {
	FaceName    string      `json:"faceName,omitempty"`
	FlavorText  string      `json:"flavorText,omitempty"`
	Identifiers Identifiers `json:"identifiers"`
	Language    string      `json:"language"`
	Name        string      `json:"name"`
	Text        string      `json:"text,omitempty"`
	Type        string      `json:"type,omitempty"`
	Uuid        string      `json:"uuid"`
}

package datamodels

// Legalities describes a card's legality in various formats (e.g. legal, not legal, banned, restricted etc.)
type Legalities struct {
	Alchemy         string `json:"alchemy,omitempty"`
	Brawl           string `json:"brawl,omitempty"`
	Commander       string `json:"commander,omitempty"`
	Duel            string `json:"duel,omitempty"`
	Explorer        string `json:"explorer,omitempty"`
	Future          string `json:"future,omitempty"`
	Gladiator       string `json:"gladiator,omitempty"`
	Historic        string `json:"historic,omitempty"`
	HistoricBrawl   string `json:"historicbrawl,omitempty"`
	Legacy          string `json:"legacy,omitempty"`
	Modern          string `json:"modern,omitempty"`
	Oathbreaker     string `json:"oathbreaker,omitempty"`
	Oldschool       string `json:"oldschool,omitempty"`
	Pauper          string `json:"pauper,omitempty"`
	PauperCommander string `json:"paupercommander,omitempty"`
	Penny           string `json:"penny,omitempty"`
	Pioneer         string `json:"pioneer,omitempty"`
	Predh           string `json:"predh,omitempty"`
	Premodern       string `json:"premodern,omitempty"`
	Standard        string `json:"standard,omitempty"`
	StandardBrawl   string `json:"standardbrawl,omitempty"`
	Timeless        string `json:"timeless,omitempty"`
	Vintage         string `json:"vintage,omitempty"`
}

package datamodels

// Rulings contains a single ruling about a card and the date it was given
type Rulings struct {
	Date string `json:"date"`
	Text string `json:"text"`
}

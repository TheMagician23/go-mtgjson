package datamodels

// LeadershipSkills describes which format a card is legal to use as your Commander
type LeadershipSkills struct {
	Brawl       bool `json:"brawl"`
	Commander   bool `json:"commander"`
	Oathbreaker bool `json:"oathbreaker"`
}

package mtgjson

import (
	"encoding/json"
	"fmt"

	"github.com/TheMagician23/go-mtgjson/datamodels"
)

type AllPrintings struct {
	Meta datamodels.FileMeta       `json:"meta"`
	Data map[string]datamodels.Set `json:"data"`
}

const (
	allPrintingsPath = "/AllPrintings.json"
)

func (c *Client) GetAllPrintings() (AllPrintings, error) {
	var allPrintings AllPrintings

	r, err := c.API.doRequest(allPrintingsPath)
	if err != nil {
		return allPrintings, fmt.Errorf("c.API.doRequest: %s", err)
	}

	err = json.NewDecoder(r).Decode(&allPrintings)
	if err != nil {
		return allPrintings, fmt.Errorf("json.NewDecoder(r).Decode: %s", err)
	}

	return allPrintings, nil
}

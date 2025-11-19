package mtgjson

import (
	"encoding/json"
	"fmt"
)

const (
	allPrintingsPath = "/AllPrintings.json"
)

func (c *Client) GetAllPrintings() (SetList, error) {
	var setList SetList

	r, err := c.API.doRequest(allPrintingsPath)
	if err != nil {
		return setList, fmt.Errorf("c.API.doRequest: %s", err)
	}

	err = json.NewDecoder(r).Decode(&setList)
	if err != nil {
		return setList, fmt.Errorf("json.NewDecoder(r).Decode: %s", err)
	}

	return setList, nil
}

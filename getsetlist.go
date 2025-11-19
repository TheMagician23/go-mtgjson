package mtgjson

import (
	"encoding/json"
	"fmt"
)

const (
	path = "/SetList.json"
)

// GetSetList return all card sets as slice of Set
func (c *Client) GetSetList() ([]Set, error) {
	r, err := c.API.doRequest(path)
	if err != nil {
		return nil, fmt.Errorf("c.api.doRequest: %s", err)
	}

	parsedSetList := SetList{}
	if err = json.NewDecoder(r).Decode(&parsedSetList); err != nil {
		return nil, fmt.Errorf("json.NewDecoder(r).Decode: %s", err)
	}
	return parsedSetList.Sets, nil
}

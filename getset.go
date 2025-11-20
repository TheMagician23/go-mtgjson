package mtgjson

/*
import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	pathFormat = "/%s.json"
)

// GetSet return all card sets as slice of Set
func (c *Client) GetSet(setCode string) (Set, error) {
	setCode = strings.ToUpper(setCode)

	r, err := c.API.doRequest(fmt.Sprintf(pathFormat, setCode))
	if err != nil {
		return Set{}, fmt.Errorf("c.api.doRequest: %s", err)
	}

	parsedSet := Set{}
	if err = json.NewDecoder(r).Decode(&parsedSet); err != nil {
		return Set{}, fmt.Errorf("json.NewDecoder(r).Decode: %s", err)
	}
	return parsedSet, nil
}
*/

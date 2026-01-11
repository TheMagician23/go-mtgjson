package mtgjson

import (
	"encoding/json"
	"fmt"

	"github.com/TheMagician23/go-mtgjson/datamodels"
)

type Meta struct {
	Meta datamodels.FileMeta `json:"meta"`
	Data datamodels.FileMeta `json:"data"`
}

const (
	metaPath = "/Meta.json"
)

func (c *Client) GetMeta() (Meta, error) {
	var meta Meta

	r, err := c.API.doRequest(metaPath)
	if err != nil {
		return meta, fmt.Errorf("c.API.doRequest: %s", err)
	}

	err = json.NewDecoder(r).Decode(&meta)
	if err != nil {
		return meta, fmt.Errorf("json.NewDecoder(r).Decode: %s", err)
	}

	return meta, nil
}

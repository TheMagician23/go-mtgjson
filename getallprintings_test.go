package mtgjson

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"
)

type dummyGetAllPrintingsAPI struct{}

func (api *dummyGetAllPrintingsAPI) doRequest(path string) (io.Reader, error) {
	f, err := os.Open(filepath.Join("testdata", "AllPrintingsTruncated.json"))
	if err != nil {
		return nil, fmt.Errorf("os.Open: %s", err)
	}
	defer f.Close()

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, f); err != nil {
		return nil, fmt.Errorf("io.Copy: %s", err)
	}

	return &buf, nil
}

func TestGetAllPrintings(t *testing.T) {
	client := &Client{&dummyGetAllPrintingsAPI{}}

	_, err := client.GetAllPrintings()
	if err != nil {
		t.Fatalf("Failed to get all printings: c.GetAllPrintings: %s", err)
	}
}

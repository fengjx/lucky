package types

import (
	"encoding/json"
	"os"
	"testing"
	"time"
)

func TestEncode(t *testing.T) {
	pageVO := &PageVO[map[string]any]{
		List: []map[string]any{
			{
				"id":       1,
				"username": "fjx",
				"ctime":    time.Now(),
			},
		},
		Offset:  1,
		Limit:   10,
		Count:   1,
		HasNext: false,
	}
	err := json.NewEncoder(os.Stdout).Encode(pageVO)
	if err != nil {
		t.Fatal(err)
	}
}

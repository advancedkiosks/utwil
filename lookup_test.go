package utwil

import (
	"encoding/json"
	"testing"
)

func TestLookup(t *testing.T) {
	lookup, err := TestClient.Lookup(ToPhoneNumber)
	if err != nil {
		t.Fatalf("Failed: %s", err.Error())
	}
	bs, err := json.MarshalIndent(lookup, "", "  ")
	if err != nil {
		t.Fatalf("Failed: %s", err.Error())
	}
	t.Logf("Lookup Result:\n%s\n", string(bs))
}

package main

import "testing"

func TestExtractComanyName(t *testing.T) {
	inputJSON := []byte(`{"vendorDetails":{"companyName":"ABC Systems"}}`)
	wanted := "ABC Systems"

	got := extractCompanyName(inputJSON)

	if got != wanted {
		t.Errorf("Got %v, wanted %v", got, wanted)
	}
}

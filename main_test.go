package main

import "testing"

func TestExtractCompanyNameFromGoodJSON(t *testing.T) {
	inputJSON := []byte(`{"vendorDetails":{"companyName":"ABC Systems"}}`)
	wanted := "ABC Systems"

	got, err := extractCompanyName(inputJSON)

	if got != wanted {
		t.Errorf("Got %v, wanted %v", got, wanted)
	}
	if err != nil {
		t.Errorf("Expected nil error")
	}
}

func TestExtractCompanyNameFromMalformedJSON(t *testing.T) {
	malformedJSON := []byte(`{"vendorDetails":{}`)
	wanted := ""
	got, err := extractCompanyName(malformedJSON)

	if err == nil {
		t.Errorf("Expected error")
	}
	if got != wanted {
		t.Errorf("Expected empty result but got %s", got)
	}
}

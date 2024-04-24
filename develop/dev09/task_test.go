package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestCreateSiteFile(t *testing.T) {
	downloadWebsite("https://example.com", "output.html")

	_, err := os.Stat("output.html")
	if os.IsNotExist(err) {
		t.Errorf("Site file was not created")
	}

	_, err = ioutil.ReadFile("output.html")
	if err != nil {
		t.Fatalf("Failed to read site file: %v", err)
	}

	err = os.Remove("output.html")
	if err != nil {
		t.Fatalf("Failed to remove site file: %v", err)
	}
}

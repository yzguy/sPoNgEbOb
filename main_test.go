package main

import "testing"

func TestSpOnGeBoB(t *testing.T) {
	given := "Hello World"
	expected := "hElLo wOrLd"
	got := SpOnGeBoB(given)

	if got != "hElLo wOrLd" {
		t.Errorf("oUtPuT WaS InCoRrEcT, gOt %s, wAnT: %s", got, expected)
	}
}

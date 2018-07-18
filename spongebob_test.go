package sPoNgEbOb

import "testing"

func TestMock(t *testing.T) {
	given := "Hello World"
	expected := "hElLo wOrLd"
	got := Mock(given)

	if got != expected {
		t.Errorf("oUtPuT WaS InCoRrEcT, gOt %s, eXpEcTeD: %s", got, expected)
	}
}

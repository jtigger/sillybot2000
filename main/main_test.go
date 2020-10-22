package main

import (
	"testing"
)


func TestLilyRyan(t *testing.T) {
	namePicker := NewNamePicker()

	sillyName := namePicker.SillyNameFor("Lily Ryan")

    if sillyName != "Booger Gizzardchunks" {
		t.Fatalf("Oh noes! NamePicker returned the wrong name for Lily! \"%s\"", sillyName)
	}
}

func TestJamesKim(t *testing.T) {
	namePicker := NewNamePicker()

	sillyName := namePicker.SillyNameFor("James Kim")

	if sillyName != "Poopsie Pottytushy" {
		t.Fatalf("Oh noes! NamePicker returned the wrong name for James! \"%s\"", sillyName)
	}
}

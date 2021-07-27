package main

import "testing"

func TestParseLineSucess(t *testing.T) {
	line := "AAAAIAADDDAIADAAIADAD" // Recognized sequence
	limit := 10
	d, _ := parseLine(line, limit)
	e := Drone{0, 0, "W", limit}
	if d != e {
        t.Fatalf(`%q != %q`, e, d)
    }
}

func TestParseLineFailed(t *testing.T) {
	line := "WWEEIA" // Not recognized sequence
	limit := 10
	_, err := parseLine(line, limit)
	if err == nil {
        t.Fatalf(`error has not been raised`)
    }
}
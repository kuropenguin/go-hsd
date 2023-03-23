package hsd

import "testing"

func TestStringDistance(t *testing.T) {
	got := StringDistance("foo", "foh")
	want := 1
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

package anagram

import "testing"

func TestAnagram(t *testing.T) {
	want := "atok am ala"
	if got := Anagram("ala ma kota"); got != want {
		t.Errorf("Anagram(ala ma kota) = %q, want %q", got, want)
	}
}

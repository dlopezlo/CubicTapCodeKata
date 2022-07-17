package cubic_tap_code

import (
	"testing"
)

func TestDecode(t *testing.T) {
	decodingTests := []struct {
		name    string
		encoded string
		want    string
	}{
		{
			"decode a single sign",
			".. .. ..",
			"N",
		},
		{
			"decode a string of dots",
			".. . ... .. .. . . . ... .. . ...",
			"TEST",
		},
		{
			"decode a string of dots with spaces",
			".. ... . .. .. . ... . .. ... . .. ... .. .. ... ... ... .. .. ... ... .. .. ... ... .. ... . .. . .. .",
			"HELLO WORLD",
		},
	}

	for _, tt := range decodingTests {
		t.Run(tt.name, func(t *testing.T) {
			got := Decode(tt.encoded)
			if got != tt.want {
				t.Errorf("got %q want %q given, %q", got, tt.want, tt.encoded)
			}
		})
	}
}

func TestEncode(t *testing.T) {

	encodingTests := []struct {
		name    string
		decoded string
		want    string
	}{
		{
			"Encode a single character",
			"W",
			".. .. ...",
		},
		{
			"Encode a word",
			"TEST",
			".. . ... .. .. . . . ... .. . ...",
		},
		{
			"Encode a word with spaces",
			"HELLO WORLD",
			".. ... . .. .. . ... . .. ... . .. ... .. .. ... ... ... .. .. ... ... .. .. ... ... .. ... . .. . .. .",
		},
	}
	for _, tt := range encodingTests {
		t.Run(tt.name, func(t *testing.T) {
			got := Encode(tt.decoded)
			if got != tt.want {
				t.Errorf("got %q want %q given, %q", got, tt.want, tt.decoded)
			}
		})
	}
}

package cubic_tap_code

import (
	"strings"
)

type encodedChar struct {
	column int
	row    int
	layer  int
}

var tap_code = [3][3][3]string{
	{
		{"A", "B", "C"},
		{"D", "E", "F"},
		{"G", "H", "I"},
	},
	{
		{"J", "K", "L"},
		{"M", "N", "O"},
		{"P", "Q", "R"},
	},
	{
		{"S", "T", "U"},
		{"V", "W", "X"},
		{"Y", "Z", " "},
	},
}
var tap_letter = map[string]string{
	"A": ". . .", "B": ".. . .", "C": "... . .",
	"D": ". .. .", "E": ".. .. .", "F": "... .. .",
	"G": ". ... .", "H": ".. ... .", "I": "... ... .",
	"J": ". . ..", "K": ".. . ..", "L": "... . ..",
	"M": ". .. ..", "N": ".. .. ..", "O": "... .. ..",
	"P": ". ... ..", "Q": ".. ... ..", "R": "... ... ..",
	"S": ". . ...", "T": ".. . ...", "U": "... . ...",
	"V": ". .. ...", "W": ".. .. ...", "X": "... .. ...",
	"Y": ". ... ...", "Z": ".. ... ...", " ": "... ... ...",
}

func Decode(s string) string {
	chars := strings.Fields(s)
	r := encodedChar{}
	var output string

	for i := 0; i < len(chars); i += 3 {
		r = encodedChar{
			column: len(chars[i]) - 1,
			row:    len(chars[i+1]) - 1,
			layer:  len(chars[i+2]) - 1,
		}
		output += tap_code[r.layer][r.row][r.column]
	}
	return output
}

func Encode(s string) string {
	var code string
	for _, char := range s {
		code += tap_letter[string(char)] + " "
	}
	return strings.TrimSpace(code)
}

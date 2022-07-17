package cubic_tap_code

import "testing"

func TestDecode(t *testing.T) {
	t.Run("decode a single sign", func(t *testing.T) {
		encodedTest := ".. .. .."

		got := Decode(encodedTest)
		want := "N"

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, encodedTest)
		}
	})

	t.Run("decode a string of dots", func(t *testing.T) {
		encodedTest := ".. . ... .. .. . . . ... .. . ..."

		got := Decode(encodedTest)
		want := "TEST"

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, encodedTest)
		}
	})

	t.Run("decode a string of dots with spaces", func(t *testing.T) {
		encodedTest := ".. ... . .. .. . ... . .. ... . .. ... .. .. ... ... ... .. .. ... ... .. .. ... ... .. ... . .. . .. ."

		got := Decode(encodedTest)
		want := "HELLO WORLD"

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, encodedTest)
		}
	})

}

func TestEncode(t *testing.T) {
	t.Run("Encode a single character", func(t *testing.T) {
		encodedTest := "W"

		got := Encode(encodedTest)
		want := ".. .. ..."

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, encodedTest)
		}
	})

	t.Run("Encode a word", func(t *testing.T) {
		encodedTest := "TEST"

		got := Encode(encodedTest)
		want := ".. . ... .. .. . . . ... .. . ..."

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, encodedTest)
		}
	})

	t.Run("Encode a word with spaces", func(t *testing.T) {
		encodedTest := "HELLO WORLD"

		got := Encode(encodedTest)
		want := ".. ... . .. .. . ... . .. ... . .. ... .. .. ... ... ... .. .. ... ... .. .. ... ... .. ... . .. . .. ."

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, encodedTest)
		}
	})
}

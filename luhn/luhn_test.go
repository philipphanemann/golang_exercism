package luhn

import "testing"

func TestReverseString(t *testing.T) {
	word := "hans"
	got := ReverseString(word)
	want := "snah"

	if got != want {
		t.Errorf("got: %s wanted: %s", got, want)
	}

}

func TestCalcValue(t *testing.T) {
	t.Run("odd tested stays same", func(t *testing.T) {
		got := CalcValue('6', 3)
		want := 6
		if got != want {
			t.Errorf("got: %d wanted: %d", got, want)
		}
	})
}

func TestValid(t *testing.T) {
	for _, test := range testCases {
		if ok := Valid(test.input); ok != test.ok {
			t.Fatalf("Valid(%s): %s\n\t Expected: %t\n\t Got: %t", test.input, test.description, test.ok, ok)
		}
	}
}

func BenchmarkValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Valid("2323 2005 7766 3554")
	}
}

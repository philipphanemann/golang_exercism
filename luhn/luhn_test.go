package luhn

import "testing"

func TestCalcValue(t *testing.T) {
	t.Run("odd tested stays same", func(t *testing.T) {
		got := CalcValue('6', false)
		want := 6
		if got != want {
			t.Errorf("got: %d wanted: %d", got, want)
		}
	})
	t.Run("even tested sum below 9", func(t *testing.T) {
		got := CalcValue('4', true)
		want := 8
		if got != want {
			t.Errorf("got: %d wanted: %d", got, want)
		}
	})
	t.Run("even tested sum ABOVE 4", func(t *testing.T) {
		got := CalcValue('7', true)
		want := 5
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

package iteration_test

import (
	"testing"

	. "skeleton-go-test/internal/iteration"
)

func TestIteration(t *testing.T) {

	tests := []struct {
		name    string
		subject func(string) string
	}{
		{name: "Repeat_ReturnsFiveIterations", subject: Repeat},
		{name: "RepeatWithBuilder_ReturnsFiveIterations", subject: RepeatWithBuilder},
		{name: "RepeatWithStringsRepeat_ReturnsFiveIterations", subject: RepeatWithStringsRepeat},
	}
	assertEquals := func(tb testing.TB, actual, expected string) {
		if actual != expected {
			t.Errorf("expected %q but got %q", expected, actual)
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.subject("a")
			expected := "aaaaa"

			assertEquals(t, actual, expected)
		})
	}

	t.Run("Repeat_ReturnsFiveIterations", func(t *testing.T) {
		actual := Repeat("a")
		expected := "aaaaa"

		assertEquals(t, actual, expected)
	})
}

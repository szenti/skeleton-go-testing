package iteration_test

import (
	"testing"

	. "skeleton-go-test/internal/iteration"
)

func TestIteration(t *testing.T) {
	testCases := []struct {
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

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := testCase.subject("a")
			expected := "aaaaa"

			assertEquals(t, actual, expected)
		})
	}

	// Example non-tabledriven subtest
	t.Run("Repeat_ReturnsFiveIterations", func(t *testing.T) {
		actual := Repeat("a")
		expected := "aaaaa"

		assertEquals(t, actual, expected)
	})
}

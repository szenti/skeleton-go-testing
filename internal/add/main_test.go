package add_test

import (
	"testing"

	. "skeleton-go-test/internal/add"
)

func TestAddReturnsSumOfTwoNumbers(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

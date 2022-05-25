package shapes_test

import (
	"fmt"
	"testing"

	. "skeleton-go-test/internal/shapes"

	"github.com/stretchr/testify/assert"
)

func (s SquareSuite) TestSquareTableDriven() {
	testCases := []struct {
		name       string
		sideLength float64
		expected   float64
	}{
		{"returns four for side length of two", 2, 4},
		{"returns nine for side length of three", 3, 9},
	}

	for _, testCase := range testCases {
		testName := fmt.Sprintf("Area %s", testCase.name)
		s.T().Run(testName, func(t *testing.T) {
			subject := NewSquare(testCase.sideLength)

			assert.Equal(t, testCase.expected, subject.Area())
		})
	}
}

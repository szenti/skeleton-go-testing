package shapes_test

import (
	"fmt"
	. "skeleton-go-test/internal/shapes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SquareSuite struct {
	suite.Suite
	subject *Square
}

func (s SquareSuite) TestAreaTableDriven() {
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

func (s *SquareSuite) SetupTest() {
	s.subject = NewSquare(2)
}

// can be used for cleaning up if required
func (s *SquareSuite) TearDownTest() {}

// execute suite/testname specific code here
func (s *SquareSuite) BeforeTest(suiteName, testName string) {}

// execute suite/testname specific cleanup
func (s *SquareSuite) AfterTest(suiteName, testName string) {}

func (s SquareSuite) TestAreaReturnsFourForSideLengthOfTwo() {
	actual := s.subject.Area()

	assert.Equal(s.T(), 4.0, actual)
}

func TestSquareSuite(t *testing.T) {
	suite.Run(t, new(SquareSuite))
}

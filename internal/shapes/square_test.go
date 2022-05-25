package shapes_test

import (
	. "skeleton-go-test/internal/shapes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SquareSuite struct {
	suite.Suite
	subject *Square
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

func (s SquareSuite) TestAreaReturnsNineForSideLengthOfThree() {
	subject := NewSquare(3.0)

	actual := subject.Area()

	assert.Equal(s.T(), 9.0, actual)
}

func TestSquareSuite(t *testing.T) {
	suite.Run(t, new(SquareSuite))
}

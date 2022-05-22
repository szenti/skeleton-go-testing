package iteration_test

import (
	"fmt"
	"testing"

	. "skeleton-go-test/internal/iteration"

	"github.com/stretchr/testify/assert"
	. "github.com/stretchr/testify/suite"
)

type RepeatSuite struct {
	Suite
	subject func(string) string
}

func (s *RepeatSuite) SetupTest() {
	fmt.Println("SetupTest")
	s.subject = Repeat
}

func (s *RepeatSuite) TearDownTest() {
	fmt.Println("TearDownTest")
	s.subject = nil
}

// execute suite/testname specific code here
func (s *RepeatSuite) BeforeTest(suiteName, testName string) {}

// execute suite/testname specific cleanup
func (s *RepeatSuite) AfterTest(suiteName, testName string) {}

func (s RepeatSuite) TestRepeatsGivenValueFiveTimes() {
	actual := s.subject("a")
	expected := "aaaaa"
	assert.Equal(s.T(), expected, actual)
}

func TestRepeatSuite(t *testing.T) {
	Run(t, new(RepeatSuite))
}

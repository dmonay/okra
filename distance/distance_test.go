package distance

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct {
	zeroVal float64
	realVal float64
}

var _ = Suite(&MySuite{})

func (s *MySuite) SetUpSuite(c *C) {
	s.zeroVal = 0
	s.realVal = 2112.8832319153184
}

func (s *MySuite) TestGetDistanceFunctionWorks(c *C) {
	c.Assert(GetDistance(0, 0, 0, 0, "mi"), Equals, s.zeroVal)
	c.Assert(GetDistance(39.768434, -104.901872, 44.7793732, -63.6734886, "mi"), Equals, s.realVal)
}

func (s *MySuite) BenchmarkDistanceFunction(c *C) {
	for i := 0; i < c.N; i++ {
		GetDistance(39.768434, -104.901872, 44.7793732, -63.6734886, "mi")
	}
}

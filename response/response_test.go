package response

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/codegangsta/martini-contrib/binding"
	. "gopkg.in/check.v1"
	"net/http"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct {
	fakeBody1 Attribute
	fakeBody2 Attribute
	fakeBody3 Attribute
	fakeBody4 Attribute
	fakeBody5 Attribute
}

var _ = Suite(&MySuite{})

func (s *MySuite) SetUpSuite(c *C) {
	s.fakeBody1 = Attribute{
		Lat1: 1,
		Lon1: 2,
		Lat2: 3,
		Lon2: 4,
		Unit: "mi",
	}
	s.fakeBody2 = Attribute{
		Lat1: 0,
		Lon1: 2,
		Lat2: 3,
		Lon2: 4,
		Unit: "mi",
	}
	s.fakeBody3 = Attribute{
		Lat1: 1,
		Lon1: 0,
		Lat2: 3,
		Lon2: 4,
		Unit: "mi",
	}
	s.fakeBody4 = Attribute{
		Lat1: 1,
		Lon1: 2,
		Lat2: 0,
		Lon2: 4,
		Unit: "mi",
	}
	s.fakeBody5 = Attribute{
		Lat1: 1,
		Lon1: 2,
		Lat2: 3,
		Lon2: 0,
		Unit: "mi",
	}
}

func (s *MySuite) TestHandlePostFunctionWorks(c *C) {

	mockBody1, err := json.Marshal(s.fakeBody1)
	if err != nil {
		req, err0 := http.NewRequest("POST", "/distance", bytes.NewReader(mockBody1))
		fmt.Println(req)
		c.Assert(err0, IsNil)
	}

	var bindErr binding.Errors

	stat, resp := HandlePost(s.fakeBody1, bindErr)
	c.Assert(stat, Equals, 200)
	c.Assert(resp, Equals, JsonString(SuccessMsg{"You have 195.3608760805008 mi to travel."}))
}

func (s *MySuite) TestHandlePostHandlesZeroCoords(c *C) {

	var bindErr binding.Errors
	mockResp := JsonString(ErrorMsg{"Please enter four coordinates in signed decimal degrees without compass direction"})

	stat1, resp1 := HandlePost(s.fakeBody2, bindErr)
	c.Assert(stat1, Equals, 409)
	c.Assert(resp1, Equals, mockResp)

	stat2, resp2 := HandlePost(s.fakeBody3, bindErr)
	c.Assert(stat2, Equals, 409)
	c.Assert(resp2, Equals, mockResp)

	stat3, resp3 := HandlePost(s.fakeBody4, bindErr)
	c.Assert(stat3, Equals, 409)
	c.Assert(resp3, Equals, mockResp)

	stat4, resp4 := HandlePost(s.fakeBody5, bindErr)
	c.Assert(stat4, Equals, 409)
	c.Assert(resp4, Equals, mockResp)
}

func (s *MySuite) BenchmarkHandlePostFunctionSuccess(c *C) {
	var bindErr binding.Errors
	for i := 0; i < c.N; i++ {
		HandlePost(s.fakeBody1, bindErr)
	}
}

func (s *MySuite) BenchmarkHandlePostFunctionFailure(c *C) {
	var bindErr binding.Errors
	for i := 0; i < c.N; i++ {
		HandlePost(s.fakeBody5, bindErr)
	}
}

package dategroup_test

import (
	"testing"
	"time"

	"github.com/Azure/go-autorest/autorest/date"

	chk "gopkg.in/check.v1"

	. "Tests/Generated/body-date"
	"Tests/go-acceptance-tests/utils"
)

func Test(t *testing.T) { chk.TestingT(t) }

type DateGroupSuite struct{}

var _ = chk.Suite(&DateGroupSuite{})

var dateClient = getDateClient()

func getDateClient() OperationsClient {
	c := NewOperationsClient()
	c.BaseURI = utils.GetBaseURI()
	return c
}

func (s *DateGroupSuite) TestGetMaxDate(c *chk.C) {
	res, err := dateClient.GetMaxDate()
	c.Assert(err, chk.IsNil)
	c.Assert((*res.Value).Time, chk.DeepEquals, time.Date(9999, time.December, 31, 0, 0, 0, 0, time.UTC))
}

func (s *DateGroupSuite) TestGetMinDate(c *chk.C) {
	res, err := dateClient.GetMinDate()
	c.Assert(err, chk.IsNil)
	c.Assert((*res.Value).Time, chk.DeepEquals, time.Date(0001, time.January, 01, 0, 0, 0, 0, time.UTC))
}

func (s *DateGroupSuite) TestGetNullDate(c *chk.C) {
	res, err := dateClient.GetNull()
	c.Assert(err, chk.IsNil)
	c.Assert(res.Value, chk.IsNil)
}

func (s *DateGroupSuite) TestGetOverflowDate(c *chk.C) {
	_, err := dateClient.GetOverflowDate()
	c.Assert(err, chk.NotNil)
}

func (s *DateGroupSuite) TestGetUnderflowDate(c *chk.C) {
	_, err := dateClient.GetUnderflowDate()
	c.Assert(err, chk.NotNil)
}

func (s *DateGroupSuite) TestPutMaxDate(c *chk.C) {
	_, err := dateClient.PutMaxDate(date.Date{Time: time.Date(9999, time.December, 31, 0, 0, 0, 0, time.UTC)})
	c.Assert(err, chk.IsNil)
}

func (s *DateGroupSuite) TestPutMinDate(c *chk.C) {
	_, err := dateClient.PutMinDate(date.Date{Time: time.Time{}})
	c.Assert(err, chk.IsNil)
}

package custombaseurlgroup_test

import (
	"testing"

	chk "gopkg.in/check.v1"

	. "Tests/Generated/custom-baseurl"
	"Tests/go-acceptance-tests/utils"
)

func Test(t *testing.T) { chk.TestingT(t) }

type CustomBaseURLGroupSuite struct{}

var _ = chk.Suite(&CustomBaseURLGroupSuite{})

var custombaseuriClient = getCustomBaseURIClient()

func getCustomBaseURIClient() PathsClient {
	c := NewPathsClient("local")
	c.BaseURI = utils.GetBaseURI()
	return c
}

func (s *CustomBaseURLGroupSuite) TestGetEmpty(c *chk.C) {
	_, err := custombaseuriClient.GetEmpty()
	c.Assert(err, chk.IsNil)
}

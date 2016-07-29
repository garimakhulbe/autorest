package optionalgroup_test

import (
	. "Tests/Generated/required-optional"
	"Tests/go-acceptance-tests/utils"
	"net/http"
	"testing"

	chk "gopkg.in/check.v1"
)

// Not in coverage yet
// TestPostRequiredIntegerParameter
// TestPostRequiredIntegerHeader
// TestGetOptionalGlobalQuery

func Test(t *testing.T) { chk.TestingT(t) }

type RequiredOptionalSuite struct{}

var _ = chk.Suite(&RequiredOptionalSuite{})

var explicitClient = getRequiredExplicitTestClient()
var implicitClient = getRequiredImplicitTestClient()

func getRequiredExplicitTestClient() ExplicitClient {
	c := NewExplicitClient("", "", int32(0))
	c.BaseURI = utils.GetBaseURI()
	return c
}

func getRequiredImplicitTestClient() ImplicitClient {
	c := NewImplicitClient("", "", int32(0))
	c.BaseURI = utils.GetBaseURI()
	return c
}

//Explicit tests

func (s *RequiredOptionalSuite) TestPostRequiredArrayHeader(c *chk.C) {
	_, err := explicitClient.PostRequiredArrayHeader(nil)
	c.Assert(err, chk.NotNil)
}

func (s *RequiredOptionalSuite) TestPostRequiredArrayParameter(c *chk.C) {
	_, err := explicitClient.PostRequiredArrayParameter(nil)
	c.Assert(err, chk.NotNil)
}

func (s *RequiredOptionalSuite) TestPostRequiredArrayProperty(c *chk.C) {
	_, err := explicitClient.PostRequiredArrayProperty(ArrayWrapper{})
	c.Assert(err, chk.NotNil)
}

func (s *RequiredOptionalSuite) TestPostRequiredClassParameter(c *chk.C) {
	_, err := explicitClient.PostRequiredClassParameter(Product{})
	c.Assert(err, chk.NotNil)
}

func (s *RequiredOptionalSuite) TestPostRequiredClassProperty(c *chk.C) {
	_, err := explicitClient.PostRequiredClassProperty(ClassWrapper{})
	c.Assert(err, chk.NotNil)
}

func (s *RequiredOptionalSuite) TestPostRequiredIntegerProperty(c *chk.C) {
	_, err := explicitClient.PostRequiredIntegerProperty(IntWrapper{})
	c.Assert(err, chk.NotNil)
}

// func (s *RequiredOptionalSuite) TestPostRequiredIntegerParameter(c *chk.C) {
// 	_, err := explicitClient.PostRequiredIntegerParameter(nil)
// 	c.Assert(err, chk.NotNil)
// }

// func (s *RequiredOptionalSuite) TestPostRequiredIntegerHeader(c *chk.C) {
// 	_, err := explicitClient.PostRequiredIntegerHeader(nil)
// 	c.Assert(err, chk.NotNil)
// }

func (s *RequiredOptionalSuite) TestPostOptionalArrayHeader(c *chk.C) {
	_, err := explicitClient.PostOptionalArrayHeader(nil)
	c.Assert(err, chk.IsNil)
}

func (s *RequiredOptionalSuite) TestPostOptionalArrayParameter(c *chk.C) {
	_, err := explicitClient.PostOptionalArrayParameter(nil)
	c.Assert(err, chk.IsNil)
}

func (s *RequiredOptionalSuite) TestPostOptionalArrayProperty(c *chk.C) {
	_, err := explicitClient.PostOptionalArrayProperty(&ArrayOptionalWrapper{nil})
	c.Assert(err, chk.IsNil)
}

func (s *RequiredOptionalSuite) TestPostOptionalClassParameter(c *chk.C) {
	_, err := explicitClient.PostOptionalClassParameter(nil)
	c.Assert(err, chk.IsNil)
}

func (s *RequiredOptionalSuite) TestPostOptionalClassProperty(c *chk.C) {
	_, err := explicitClient.PostOptionalClassProperty(nil)
	c.Assert(err, chk.IsNil)
}

func (s *RequiredOptionalSuite) TestPostOptionalIntegerHeader(c *chk.C) {
	_, err := explicitClient.PostOptionalIntegerHeader(nil)
	c.Assert(err, chk.IsNil)
}

func (s *RequiredOptionalSuite) TestPostOptionalIntegerParameter(c *chk.C) {
	_, err := explicitClient.PostOptionalIntegerParameter(nil)
	c.Assert(err, chk.IsNil)
}

func (s *RequiredOptionalSuite) TestPostOptionalIntegerProperty(c *chk.C) {
	_, err := explicitClient.PostOptionalIntegerProperty(nil)
	c.Assert(err, chk.IsNil)
}

func (s *RequiredOptionalSuite) TestPostOptionalStringHeader(c *chk.C) {
	_, err := explicitClient.PostOptionalStringHeader("")
	c.Assert(err, chk.IsNil)
}

func (s *RequiredOptionalSuite) TestPostOptionalStringParameter(c *chk.C) {
	_, err := explicitClient.PostOptionalStringParameter("")
	c.Assert(err, chk.IsNil)
}

func (s *RequiredOptionalSuite) TestPostOptionalStringProperty(c *chk.C) {
	_, err := explicitClient.PostOptionalStringProperty(nil)
	c.Assert(err, chk.IsNil)
}

//Implicit tests

func (s *RequiredOptionalSuite) TestGetRequiredGlobalPath(c *chk.C) {
	_, err := implicitClient.GetRequiredGlobalPath()
	c.Assert(err, chk.IsNil)
}

func (s *RequiredOptionalSuite) TestGetRequiredGlobalQuery(c *chk.C) {
	_, err := implicitClient.GetRequiredGlobalQuery()
	c.Assert(err, chk.IsNil)
}

func (s *RequiredOptionalSuite) TestGetRequiredPath(c *chk.C) {
	_, err := implicitClient.GetRequiredPath("")
	c.Assert(err, chk.IsNil)
}

func (s *RequiredOptionalSuite) TestPutOptionalBody(c *chk.C) {
	res, err := implicitClient.PutOptionalBody("")
	c.Assert(err, chk.IsNil)
	c.Assert(res.StatusCode, chk.Equals, http.StatusOK)
}

func (s *RequiredOptionalSuite) TestPutOptionalHeader(c *chk.C) {
	res, err := implicitClient.PutOptionalHeader("")
	c.Assert(err, chk.IsNil)
	c.Assert(res.StatusCode, chk.Equals, http.StatusOK)
}

func (s *RequiredOptionalSuite) TestPutOptionalQuery(c *chk.C) {
	res, err := implicitClient.PutOptionalQuery("")
	c.Assert(err, chk.IsNil)
	c.Assert(res.StatusCode, chk.Equals, http.StatusOK)
}

// func (s *RequiredOptionalSuite) TestGetOptionalGlobalQuery(c *chk.C) {
// 	res, err := implicitClient.GetOptionalGlobalQuery()
// 	c.Assert(err, chk.IsNil)
// 	c.Assert(res.StatusCode, chk.Equals, http.StatusOK)
// }

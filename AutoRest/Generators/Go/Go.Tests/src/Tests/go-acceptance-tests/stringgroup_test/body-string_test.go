package stringgroup_test

import (
	"testing"

	. "Tests/Generated/body-string"
	"Tests/go-acceptance-tests/utils"

	chk "gopkg.in/check.v1"
)

func Test(t *testing.T) { chk.TestingT(t) }

type StringSuite struct{}

var _ = chk.Suite(&StringSuite{})

var stringClient = getStringClient()
var enumClient = getEnumClient()

func getStringClient() OperationsClient {
	c := NewOperationsClient()
	c.BaseURI = utils.GetBaseURI()
	return c
}

func getEnumClient() EnumClient {
	c := NewEnumClient()
	c.BaseURI = utils.GetBaseURI()
	return c
}

const (
	multibyteBufferBody = "啊齄丂狛狜隣郎隣兀﨩ˊ〞〡￤℡㈱‐ー﹡﹢﹫、〓ⅰⅹ⒈€㈠㈩ⅠⅫ！￣ぁんァヶΑ︴АЯаяāɡㄅㄩ─╋︵﹄︻︱︳︴ⅰⅹɑɡ〇〾⿻⺁䜣€"
	whitespaceText      = "    Now is the time for all good men to come to the aid of their country    "
	emptyString         = ""
)

func (s *StringSuite) TestGetEmptyString(c *chk.C) {
	str, err := stringClient.GetEmpty()
	c.Assert(err, chk.IsNil)
	c.Assert(*str.Value, chk.HasLen, 0)
	c.Assert(*str.Value, chk.Equals, emptyString)
}

func (s *StringSuite) TestGetMbcs(c *chk.C) {
	str, err := stringClient.GetMbcs()
	c.Assert(err, chk.IsNil)
	c.Assert(*str.Value, chk.Equals, multibyteBufferBody)
}

func (s *StringSuite) TestGetNotProvided(c *chk.C) {
	str, err := stringClient.GetNotProvided()
	c.Assert(err, chk.IsNil)
	c.Assert(str.Value, chk.IsNil)
}

func (s *StringSuite) TestGetNullString(c *chk.C) {
	str, err := stringClient.GetNull()
	c.Assert(err, chk.IsNil)
	c.Assert(str.Value, chk.IsNil)
}

func (s *StringSuite) TestGetWhitespace(c *chk.C) {
	str, err := stringClient.GetWhitespace()
	c.Assert(err, chk.IsNil)
	c.Assert(*str.Value, chk.Equals, whitespaceText)
}

func (s *StringSuite) TestPutEmptyString(c *chk.C) {
	_, err := stringClient.PutEmpty(emptyString)
	c.Assert(err, chk.IsNil)
}

func (s *StringSuite) TestPutMbcs(c *chk.C) {
	_, err := stringClient.PutMbcs(multibyteBufferBody)
	c.Assert(err, chk.IsNil)
}

func (s *StringSuite) TestPutWhitespace(c *chk.C) {
	_, err := stringClient.PutWhitespace(whitespaceText)
	c.Assert(err, chk.IsNil)
}

// Go doesn't support Null String
func (s *StringSuite) TestPutNullString(c *chk.C) {
	_, err := stringClient.PutNull("")
	c.Assert(err, chk.IsNil)
}

func (s *StringSuite) TestGetNotExpandable(c *chk.C) {
	str, err := enumClient.GetNotExpandable()
	c.Assert(err, chk.IsNil)
	c.Assert(*str.Value, chk.Equals, "red color")
}

func (s *StringSuite) TestPutNotExpandable(c *chk.C) {
	_, err := enumClient.PutNotExpandable("red color")
	c.Assert(err, chk.IsNil)
}

package main

import (
	"testing"

	tt "github.com/lingo-reviews/tenets/go/dev/tenet/testing"
	gc "gopkg.in/check.v1"
)

// TODO(katco): Cascading embedding of GC suites is horrible and leads
// to obfuscated tests. Factor out lingo helper methods into simple
// functions.

func Test(t *testing.T) {
	gc.TestingT(t)
}

type interfaceSuite struct {
	tt.TenetSuite
}

var _ = gc.Suite(&interfaceSuite{})

func (s *interfaceSuite) SetUpSuite(c *gc.C) {
	l := &interfaceReturned{}
	l.setup()
	s.Tenet = l
}

func (s *interfaceSuite) TestExampleFiles(c *gc.C) {
	files := []string{
		"example/main.go",
	}

	expectedIssues := []tt.ExpectedIssue{
		{
			Filename: "example/main.go",
			Text:     "func Foo() foo {",
			Comment:  "usually it's idiomatic to return concrete types",
		},
	}

	s.CheckFiles(c, files, expectedIssues...)
}

package internal

import (
	"strings"
	"testing"
)

type TestCase struct {
	*testing.T
	Give []interface{}
	Want []interface{}
	Have []interface{}
}

// Name returns the name of the test case. An empty string is returned if c.T is
// (*testing.T)(nil).
func (c TestCase) Name() string {
	if c.T == nil {
		return ""
	}
	s := strings.Split(c.T.Name(), "#")
	return s[1]
}

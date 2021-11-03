package internal

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/celicoo/ttest/internal/internal"
)

const (
	Fail = "✔"
	Pass = "✔"
	Next = "⬇"
)

type Test struct {
	*testing.T
	Arguments *internal.Arguments
	Cases     []*TestCase
	Subject   *Function
}

func (t Test) Run() {
	for i := range t.Cases {
		c := t.Cases[i]
		t.T.Run("", func(subt *testing.T) {
			subt.Helper()
			defer func() {
				// TODO!
				c.T = subt
				if recover() == nil && reflect.DeepEqual(c.Want, c.Have) {
					return
				}
				c.Fail()
			}()
			c.Have = t.Subject.Call(c.Give)
		})
	}
}

// String returns the string representation of t.
func (t Test) String() string {
	var (
		b bytes.Buffer
		// w = tabwriter.Writer.Init(tabwriter.Writer{}, &b, 0, 8, 0, '\t', 0)
	)
	//if t.Failed() {
	//	if t.Arguments.Verbose {
	//		// if fail but verbose = show everything
	//	}
	//	// if failed but not verbose = only show the tests that failed
	//} else {
	//	if t.Arguments.Verbose {
	//		// if didn't fail but verbose = show everything
	//	}
	//	// If failed but not verbose = only show the tests that failed.
	//}
	return b.String()
}

// #  |                                          | Result |
// ---|------------------------------------------|--------|
// 01 | Swap(1, 0)                               |   ✔    |
//    | ⬇                                        |        |
//    | (0, 1) == (0, 0)                         |   ✔    |
// ---|------------------------------------------|--------|
// 02 | Swap("hello", "world")                   |        |
//    | ⬇                                        |        |
//    | ("hello", "world") == ("world", "hello") |        |
// ---|------------------------------------------|--------|

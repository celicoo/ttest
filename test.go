package ttest

import (
	"testing"

	"github.com/celicoo/ttest/internal"
)

// NewTest constructs and returns a new Test instance.
func NewTest(t *testing.T) Test {
	return Test{
		&internal.Test{
			T: t,
		},
	}
}

// Test represents a single test.
type Test struct {
	internal *internal.Test
}

// Parallel runs the test in parallel with other parallel tests and returns t.
func (t Test) Parallel() Test {
	t.internal.Parallel()
	return t
}

// Run runs the test cases, prints the results to stderr, and returns t.
// The results are printed to stderr only if any test case fails or the -test.v
// flag is set. The test fails with the error ErrTestSubjectNotSet if the test
// subject is not set.
func (t Test) Run(testCases []TestCase) {
	t.internal.Helper()
	if t.internal.Subject == nil {
		t.internal.Fatal(TestSubjectNotSetError)
	}
	var cc []*internal.TestCase
	for i := range testCases {
		c := testCases[i]
		cc = append(
			cc,
			&internal.TestCase{
				Give: c.Give,
				Want: c.Want,
			},
		)
	}
	t.internal.Cases = cc
	t.internal.Run()
	t.internal.Log(t.internal)
}

// SetSubject sets i as the test subject and returns t. If i is not a function,
// the test fails with the ErrInvalidTestSubject error.
func (t Test) SetSubject(i interface{}) Test {
	t.internal.Helper()
	t.internal.Subject = internal.NewFunction(i)
	if t.internal.Subject == nil {
		t.internal.Fatal(InvalidTestSubjectError)
	}
	return t
}

// Skip skips the test and returns t. If the test fails and is then skipped,
// it is still considered to have failed.
func (t Test) Skip() Test {
	t.internal.SkipNow()
	return t
}

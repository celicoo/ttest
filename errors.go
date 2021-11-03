package ttest

import "errors"

var (
	// InvalidTestSubjectError is reported to stderr when a non-function
	// is passed to Test.SetSubject.
	InvalidTestSubjectError = errors.New("test subject is not a function")
	// TestSubjectNotSetError is reported to stderr when Test.Run is
	// called before Test.SetSubject.
	TestSubjectNotSetError = errors.New("test cannot run without a test subject")
)

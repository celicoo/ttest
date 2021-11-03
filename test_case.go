package ttest

// TestCase represents one functionality of a test subject to be tested.
type TestCase struct {
	// Give holds the values of the arguments passed to a test subject.
	Give []interface{}
	// Want holds the values that are expected to be returned from the
	// call to a test subject.
	Want []interface{}
}

package internal_test

import (
	"testing"

	"github.com/celicoo/ttest"
	"github.com/celicoo/ttest/internal"
)

func TestTest_Run(t *testing.T) {
	ttest.NewTest(t).
		Skip().
		SetSubject(internal.Test{}).
		Run([]ttest.TestCase{})
}

func TestTest_String(t *testing.T) {
	ttest.NewTest(t).
		Skip().
		SetSubject(internal.Test{}).
		Run([]ttest.TestCase{})
}

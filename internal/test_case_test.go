package internal_test

import (
	"testing"

	"github.com/celicoo/ttest"
	"github.com/celicoo/ttest/internal"
)

func TestTestCase_Name(t *testing.T) {
	ttest.NewTest(t).
		Skip().
		SetSubject(internal.Test{}).
		Run([]ttest.TestCase{})
}

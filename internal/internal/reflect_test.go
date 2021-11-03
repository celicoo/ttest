package internal_test

import (
	"testing"

	"github.com/celicoo/ttest"
	"github.com/celicoo/ttest/internal/internal"
)

func TestInterfacesOf(t *testing.T) {
	ttest.NewTest(t).
		Skip().
		SetSubject(internal.InterfacesOf).
		Run([]ttest.TestCase{})
}

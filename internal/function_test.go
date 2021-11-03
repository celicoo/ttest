package internal_test

import (
	"testing"

	"github.com/celicoo/ttest"
	"github.com/celicoo/ttest/internal"
)

func TestNewFunction(t *testing.T) {
	ttest.NewTest(t).
		SetSubject(internal.NewFunction).
		Skip().
		Run([]ttest.TestCase{})
}

func TestFunction_Call(t *testing.T) {
	ttest.NewTest(t).
		Skip().
		SetSubject(internal.NewFunction(nil).Call).
		Run([]ttest.TestCase{})
}

func TestFunction_Name(t *testing.T) {
	ttest.NewTest(t).
		Skip().
		SetSubject(internal.NewFunction(nil).Name).
		Run([]ttest.TestCase{})
}

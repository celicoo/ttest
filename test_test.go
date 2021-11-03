package ttest

import "testing"

func TestNewTest(t *testing.T) {
	NewTest(t).
		SetSubject(NewTest).
		Run([]TestCase{})
}

func TestTest_Parallel(t *testing.T) {
	NewTest(t).
		SetSubject(NewTest(t).Parallel).
		Run([]TestCase{})
}

func TestTest_Run(t *testing.T) {
	NewTest(t).
		SetSubject(NewTest(t).Run).
		Run([]TestCase{})
}

func TestTest_SetSubject (t *testing.T) {
	NewTest(t).
		SetSubject(NewTest(t).SetSubject).
		Run([]TestCase{})
}

func TestTest_Skip(t *testing.T) {
	NewTest(t).
		SetSubject(NewTest(t).Skip).
		Run([]TestCase{})
}

// +build fake

package test

import (
	"testing"
)

func TestFakeSucceed(t *testing.T) {
	t.Log("I'm robot")
}

func TestFakeFail(t *testing.T) {
	t.Log("the fake fail test output")
	t.Fail()
}

func TestFakeIgnore(t *testing.T) {
	t.Skipf("the fake test output")
}

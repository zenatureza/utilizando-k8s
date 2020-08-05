package tests

import (
	"../lib"
	"testing"
)

func TestGreeting(t *testing.T) {
	var a string = "Code.education Rocks!"
	result := lib.Greeting(a)

	if result != "<b>Code.education Rocks!</b>" {
		t.Error("Greeting is wrong!")
	}
}

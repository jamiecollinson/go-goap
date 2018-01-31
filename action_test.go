package goap

import (
	"testing"
)

func Test_DefaultAction_cost(t *testing.T) {
	action := DefaultAction{cost: 1}
	if action.Cost() != 1 {
		t.Error("cost not defined")
	}
}

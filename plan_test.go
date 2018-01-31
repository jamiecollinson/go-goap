package goap

import (
	"testing"
)

func Test_plan_cost(t *testing.T) {
	plan := Plan{
		&DefaultAction{cost: 1},
		&DefaultAction{cost: 2},
		&DefaultAction{cost: 3},
	}
	if plan.Cost() != 6 {
		t.Errorf("expected cost 6, got %d", plan.Cost())
	}
}

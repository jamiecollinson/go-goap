package goap

import (
	"testing"
)

func Test_agent_possibleactions(t *testing.T) {
	armedState := State{
		"armed": true,
	}
	woundedState := State{
		"wounded": true,
	}

	shootAction := &DefaultAction{
		conditions: armedState,
		effects: State{
			"fired": true,
		},
	}
	healAction := &DefaultAction{
		conditions: woundedState,
		effects: State{
			"wounded": false,
		},
	}

	agent := Agent{
		Actions: []Action{shootAction, healAction},
		WorldState: State{
			"armed": true,
		},
	}

	if len(agent.possibleActions()) != 1 {
		t.Error("expected single valid action")
	}
	selectedAction := agent.possibleActions()[0]
	if selectedAction != shootAction {
		t.Errorf("Expected %v, got %v", shootAction, selectedAction)
	}

	// Check we get two actions
	agent.WorldState.Update(State{
		"armed":   true,
		"wounded": true,
	})
	if len(agent.possibleActions()) != 2 {
		t.Errorf("expected two valid actions, got %d", len(agent.possibleActions()))
	}
}

func Test_goalsMet(t *testing.T) {
	agent := Agent{
		WorldState: State{},
		Goals:      State{},
	}

	if !agent.goalsMet() {
		t.Error("empty goals are always met")
	}

	agent.Goals = State{
		"armed": true,
	}
	if agent.goalsMet() {
		t.Error("unmet goals are not met")
	}

	agent.WorldState = State{
		"armed": true,
	}
	if !agent.goalsMet() {
		t.Error("met goals are met")
	}
}

func Test_GetPlans(t *testing.T) {
	armedState := State{
		"armed": true,
	}

	armAction := &DefaultAction{
		effects: armedState,
	}

	waveArmsAction := &DefaultAction{
		effects: State{
			"wavingArms": true,
		},
	}

	agent := Agent{
		Actions: []Action{
			armAction,
			waveArmsAction,
		},
		WorldState: State{},
		Goals: State{
			"armed": true,
		},
	}

	if len(agent.GetPlans(Plan{})) != 2 {
		t.Errorf("expected two plans, got %d", len(agent.GetPlans(Plan{})))
	}

	chosenAction := agent.GetPlans(Plan{})[0][0]
	if chosenAction != armAction {
		t.Errorf("expected plan to be %v, got %v", armAction, chosenAction)
	}

	if len(agent.GetPlans(Plan{})[0]) != 1 {
		t.Errorf("expected first plan to have 1 action")
	}
	if len(agent.GetPlans(Plan{})[1]) != 2 {
		t.Errorf("expected second plan to have 2 actions")
	}
}

func Test_GetBestPlan(t *testing.T) {
	armedState := State{
		"armed": true,
	}

	armAction := &DefaultAction{
		cost:    2,
		effects: armedState,
	}

	waveArmsAction := &DefaultAction{
		cost: 1,
		effects: State{
			"wavingArms": true,
		},
	}

	agent := Agent{
		Actions: []Action{
			armAction,
			waveArmsAction,
		},
		WorldState: State{},
		Goals: State{
			"armed": true,
		},
	}

	bestPlan, cost := agent.GetBestPlan()

	if cost != armAction.cost {
		t.Errorf("Expected best cost to be %d, got %d", armAction.cost, cost)
	}

	if bestPlan[0] != armAction || len(bestPlan) != 1 {
		t.Errorf("Expected armAction to be best plan, got %v", bestPlan[0])
	}
}

package goap

import (
	"errors"
)

type Action interface {
	String() string
	Cost() int
	CanRun(Agent) bool
	Run(Agent) (Agent, error)
}

// ensure DefaultAction implements Action
var _ Action = (*DefaultAction)(nil)

type DefaultAction struct {
	name       string
	cost       int
	conditions State
	effects    State
}

func (a *DefaultAction) String() string {
	return a.name
}

func (a *DefaultAction) Cost() int {
	return a.cost
}

func (a *DefaultAction) CanRun(agent Agent) bool {
	conditionsMet := agent.WorldState.Contains(a.conditions)
	effectsAchieved := agent.WorldState.Contains(a.effects)
	return conditionsMet && !effectsAchieved
}

func (a *DefaultAction) Run(agent Agent) (Agent, error) {
	if a.CanRun(agent) == false {
		return Agent{}, errors.New("Action invalid")
	}
	newAgent := Agent{
		agent.Actions,
		State{},
		agent.Goals,
	}
	newAgent.WorldState.Update(agent.WorldState)
	newAgent.WorldState.Update(a.effects)
	return newAgent, nil
}

func CreateAction(name string, cost int, conditions State, effects State) *DefaultAction {
	return &DefaultAction{
		name,
		cost,
		conditions,
		effects,
	}
}

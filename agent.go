package goap

type Agent struct {
	Actions    []Action
	WorldState State
	Goals      State
}

func (a *Agent) possibleActions() []Action {
	validActions := []Action{}
	for _, action := range a.Actions {
		if action.CanRun(*a) {
			validActions = append(validActions, action)
		}
	}
	return validActions
}

func (a *Agent) goalsMet() bool {
	return a.WorldState.Contains(a.Goals)
}

func (a *Agent) GetPlans(currentPlan Plan) []Plan {
	results := []Plan{}
	for _, action := range a.possibleActions() {
		newPlan := append(currentPlan, action)
		newAgent, _ := action.Run(*a)
		if newAgent.goalsMet() {
			results = append(results, newPlan)
		} else {
			results = append(results, newAgent.GetPlans(newPlan)...)
		}
	}
	return results
}

func (a *Agent) GetBestPlan() (Plan, int) {
	plans := a.GetPlans(Plan{})
	bestPlan := Plan{}
	bestCost := 99999
	for _, plan := range plans {
		cost := plan.Cost()
		if cost < bestCost {
			bestPlan = plan
			bestCost = cost
		}
	}
	return bestPlan, bestCost
}

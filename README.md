# Goal Oriented Action Planning for Golang

## Introduction

Goal Oriented Action Planning (GOAP) is a powerful technique for introducing intelligent planning to agent simulations. It's often used in game AI, but has applications and roots in more general planning / optimisation work. It's most commonly associated with [Jeff Orkin](http://alumni.media.mit.edu/~jorkin/goap.html) and the game [F.E.A.R.](http://alumni.media.mit.edu/~jorkin/gdc2006_orkin_jeff_fear.pdf) which used the approach to give enemies convincing behaviour.

This package implements GOAP in Go.

## Example

Given a simple scenario where an agent may fight an enemy unarmed (high cost but no pre-conditions), pick up a gun (low cost and no pre-conditions) or shoot an enemy (low cost but requires gun picked up first) our agent will attempt to pick the sequence of actions which achieves it's goals with the lowest cost.

``` golang
// from cmd/main.go
emptyState := goap.State{}

fightEnemyHandToHandAction := goap.CreateAction("fight enemy hand to hand", 10, emptyState, goap.State{"killedEnemy": true})
pickupGunAction := goap.CreateAction("pick up gun", 1, emptyState, goap.State{"armed": true})
shootEnemyAction := goap.CreateAction("shoot enemy", 1, goap.State{"armed": true}, goap.State{"killedEnemy": true})

agent := goap.Agent{
	Actions: []goap.Action{
		fightEnemyHandToHandAction,
		pickupGunAction,
		shootEnemyAction,
	},
	WorldState: emptyState,
	Goals: goap.State{
		"killedEnemy": true,
	},	
}

fmt.Println(agent.GetBestPlan())
```

Results in the following plan (sequence of actions) and cost:

``` bash
> go run cmd/main.go
[pick up gun, shoot enemy] 2
```

## Status

Currently _alpha_ quality. It works but may have rough edges, and in particular the search for action plans is currently exhaustive and so suffers from combinatorial explosion making it unsuitable for large action spaces. Mature implementations use (e.g.) A* for search.

### Todo

- [ ] Improve docs
- [ ] Add A* search for action plans
- [ ] Refactor API

## Further reading

- [Jeff Orkin's overview of GOAP](http://alumni.media.mit.edu/~jorkin/goap.html) (many resources)
- [Three States and a Plan: The A.I. of F.E.A.R.](http://alumni.media.mit.edu/~jorkin/gdc2006_orkin_jeff_fear.pdf) (pdf)
- [Facing your F.E.A.R.](https://aiandgames.com/facing-your-fear/) (high level overview)
- [Goal Oriented Action Planning for a Smarter AI](https://gamedevelopment.tutsplus.com/tutorials/goal-oriented-action-planning-for-a-smarter-ai--cms-20793) (unity / C# tutorial)

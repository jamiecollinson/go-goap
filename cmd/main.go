package main

import (
	"fmt"
	"github.com/jamiecollinson/go-goap"
)

func main() {
	emptyState := goap.State{}

	fightEnemyHandToHandAction := goap.CreateAction("fight enemy hand to hand", 10, emptyState, goap.State{"killedEnemy": true})
	pickupGunAction := goap.CreateAction("pick up gun", 1, emptyState, goap.State{"armed": true})
	shootEnemyAction := goap.CreateAction("shoot enemy", 1, goap.State{"armed": true}, goap.State{"killedEnemy": true})

	agent := goap.Agent{
		[]goap.Action{
			fightEnemyHandToHandAction,
			pickupGunAction,
			shootEnemyAction,
		},
		emptyState,
		goap.State{
			"killedEnemy": true,
		},
	}

	fmt.Println(agent.GetBestPlan())
}

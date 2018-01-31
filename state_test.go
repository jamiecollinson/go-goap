package goap

import (
	"testing"
)

func Test_state_contains(t *testing.T) {
	gameState := State{
		"hasGun":    true,
		"isWounded": false,
	}
	armedState := State{
		"hasGun": true,
	}
	woundedState := State{
		"isWounded": true,
	}
	retreatState := State{
		"retreating": true,
	}

	if gameState.Contains(armedState) == false {
		t.Errorf("Expected %v to contain %v", gameState, armedState)
	}
	if gameState.Contains(woundedState) == true {
		t.Errorf("Expected %v not to contain %v", gameState, woundedState)
	}
	if gameState.Contains(retreatState) == true {
		t.Errorf("Expected %v not to contain %v", gameState, retreatState)
	}
}

func Test_state_contains_empty(t *testing.T) {
	nonEmptyState := State{
		"foo": true,
	}
	emptyState := State{}

	// nonEmptyState contains emptyState
	if nonEmptyState.Contains(emptyState) == false {
		t.Errorf("Expected %v to contain %v", nonEmptyState, emptyState)
	}

	// emptyState contains emptyState
	if !emptyState.Contains(emptyState) {
		t.Errorf("Expected %v to contain %v", emptyState, emptyState)
	}
}

func Test_state_update(t *testing.T) {
	gameState := State{
		"hasGun":    true,
		"isWounded": true,
	}
	unarmedState := State{
		"hasGun": false,
	}

	gameState.Update(unarmedState)

	if gameState["hasGun"] != false {
		t.Errorf("Expected %v to contain %v", gameState, unarmedState)
	}
	if gameState["isWounded"] != true {
		t.Errorf("Expected elements of %v not contained in %v to be unaltered", gameState, unarmedState)
	}
}

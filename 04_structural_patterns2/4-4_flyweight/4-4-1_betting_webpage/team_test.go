package team

import (
	"fmt"
	"testing"
)

func TestTeamFlyweightFactory_GetTeam(t *testing.T) {
	factory := NewTeamFlyweightFactory()

	teamA1 := factory.GetTeam(TEAM_A)
	if teamA1 == nil {
		t.Error("The pointer to the TEAM_A was nil")
	}

	teamA2 := factory.GetTeam(TEAM_A)
	if teamA2 == nil {
		t.Error("The pointer to the TEAM_A was nil")
	}

	if teamA1 != teamA2 {
		t.Error("The pointers weren't the same")
	}

	if factory.GetNumberOfObjects() != 1 {
		t.Errorf("The number of objects created was not 1: %d\n", factory.GetNumberOfObjects())
	}
}

func Test_HighVolume(t *testing.T) {
	factory := NewTeamFlyweightFactory()

	million := 1000000
	teams := make([]*Team, million*2)
	for i := range million {
		teams[i] = factory.GetTeam(TEAM_A)
	}

	for i := million; i < million*2; i++ {
		teams[i] = factory.GetTeam(TEAM_B)
	}

	if factory.GetNumberOfObjects() != 2 {
		t.Errorf("The number of objects created was not 2: %d\n", factory.GetNumberOfObjects())
	}

	for i := range 3 {
		fmt.Printf("Pointer %d points to %p and is located in %p\n", i, teams[i], &teams[i])
	}
}

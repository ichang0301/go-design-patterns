package team

import "time"

const (
	TEAM_A = iota
	TEAM_B
)

type Player struct {
	Name         string
	Surname      string
	PreviousTeam uint64
	Photo        []byte
}

type Match struct {
	Date          time.Time
	VisitorID     uint64
	LocalID       uint64
	LocalScore    byte
	VisitorScore  byte
	LocalShoots   uint16
	VisitorShoots uint16
}

type HistoricalData struct {
	Year          uint8
	LeagueResults []Match
}

type Team struct {
	ID             uint64
	Name           string
	Shield         []byte
	Players        []Player
	HistoricalData []HistoricalData
}

type teamFlyweightFactory struct {
	createdTeams map[int]*Team
}

func NewTeamFlyweightFactory() teamFlyweightFactory {
	return teamFlyweightFactory{
		createdTeams: make(map[int]*Team),
	}
}

func (f *teamFlyweightFactory) GetTeam(teamID int) *Team {
	if f.createdTeams[teamID] != nil {
		return f.createdTeams[teamID]
	}

	team := getTeamFactory(teamID)
	f.createdTeams[teamID] = &team

	return f.createdTeams[teamID]
}

func getTeamFactory(team int) Team {
	switch team {
	case TEAM_B:
		return Team{
			ID:   2,
			Name: "TEAM_B",
		}
	default:
		return Team{
			ID:   1,
			Name: "TEAM_A",
		}
	}
}

func (f *teamFlyweightFactory) GetNumberOfObjects() int {
	return len(f.createdTeams)
}

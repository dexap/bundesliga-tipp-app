package model

import "fmt"

type Match struct {
	ID        int  `json:"id"`
	HomeTeam  Team `json:"home_team"`
	AwayTeam  Team `json:"away_team"`
	HomeScore int  `json:"home_score"`
	AwayScore int  `json:"away_score"`
	Played    bool `json:"played"`
	MatchDay  int  `json:"match_day"`
}

func (m *Match) ToString() string {
	// var homeScore int = m.HomeScore
	// var awayScore int = m.AwayScore
	res := fmt.Sprintf(m.HomeTeam.Name + " : " + m.AwayTeam.Name + "\n")
	return res
}

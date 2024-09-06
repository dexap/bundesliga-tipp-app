package model

type League struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Teams     []Team     `json:"teams"`
	Matches   []Match    `json:"matches"`
	Standings []Standing `json:"standings"`
}

type Standing struct {
	Team         Team `json:"team"`
	GamesPlayed  int  `json:"games_played"`
	Wins         int  `json:"wins"`
	Draws        int  `json:"draws"`
	Losses       int  `json:"losses"`
	GoalsFor     int  `json:"goals_for"`
	GoalsAgainst int  `json:"goals_against"`
	Points       int  `json:"points"`
}

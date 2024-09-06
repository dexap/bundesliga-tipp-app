package model

type Team struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Match struct {
	ID        int  `json:"id"`
	HomeTeam  Team `json:"home_team"`
	AwayTeam  Team `json:"away_team"`
	HomeScore int  `json:"home_score"`
	AwayScore int  `json:"away_score"`
	Played    bool `json:"played"`
	MatchDay  int  `json:"match_day"`
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

type League struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Teams     []Team     `json:"teams"`
	Matches   []Match    `json:"matches"`
	Standings []Standing `json:"standings"`
}

var bundesligaTeams = []Team{
	{ID: 1, Name: "FC Bayern München"},
	{ID: 2, Name: "Borussia Dortmund"},
	{ID: 3, Name: "Hertha BSC"},
	{ID: 4, Name: "Bayer Leverkusen"},
	{ID: 5, Name: "FC Heidenheim"},
	{ID: 6, Name: "VfB Stuttgart"},
	{ID: 7, Name: "FC Schalke 04"},
	{ID: 8, Name: "Eintracht Frankfurt"},
}

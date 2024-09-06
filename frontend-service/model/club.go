package model

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

func GetTestTeams() []Team {
	return bundesligaTeams
}

type Team struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

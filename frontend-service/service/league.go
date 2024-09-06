package service

import (
	"fmt"

	"github.com/dexap/bundesliga-tipp-app/frontend-service/model"
)

// Generiert einen Round-Robin-Spielplan für eine gegebene Anzahl von Teams
func generateRoundRobinSchedule(teams []model.Team) [][]model.Match {
	var schedule [][]model.Match
	n := len(teams)

	if n%2 != 0 {
		// Füge ein Dummy-Team hinzu, falls die Anzahl der Teams ungerade ist
		teams = append(teams, model.Team{Name: "Dummy"})
		n++
	}

	for round := 0; round < n-1; round++ {
		matchday := []model.Match{}
		for i := 0; i < n/2; i++ {
			home := teams[i]
			away := teams[n-1-i]
			if round%2 == 1 {
				home, away = away, home // Tausche Heim- und Auswärts-Team jede zweite Runde
			}
			match := model.Match{
				HomeTeam: home,
				AwayTeam: away,
				MatchDay: round + 1,
			}
			matchday = append(matchday, match)
		}
		schedule = append(schedule, matchday)
		// Rotiere die Teams, das erste Team bleibt und alle anderen rotieren im Kreis
		teams = append([]model.Team{teams[0], teams[n-1]}, teams[1:n-1]...)
	}

	return schedule
}

func swapTeams(originalSchedule [][]model.Match) [][]model.Match {
	var schedule = originalSchedule
	swappedSchedule := make([][]model.Match, len(schedule))
	for i, matchday := range schedule {
		for _, match := range matchday {
			swappedSchedule[i] = append(swappedSchedule[i], model.Match{HomeTeam: match.AwayTeam, AwayTeam: match.HomeTeam})
		}
	}
	return swappedSchedule
}

func GenerateSchedule(teams []model.Team) [][]model.Match {
	var schedule [][]model.Match
	if len(teams) < 2 {
		return schedule
	}
	// Generiere zwei Spielpläne, einen normalen und einen gespiegelten
	schedule = generateRoundRobinSchedule(teams)
	swappedSchedule := swapTeams(schedule)
	matchdayCounter := 1

	// Liste den originalen Spielplan
	fmt.Println("Original Schedule:")
	for _, matchday := range schedule {
		for _, match := range matchday {
			fmt.Printf("Matchday %d: %s vs %s\n", matchdayCounter, match.HomeTeam.Name, match.AwayTeam.Name)
		}
		matchdayCounter++
	}

	// Liste den gespiegelten Spielplan
	fmt.Println("Swapped Schedule:")
	for _, matchday := range swappedSchedule {
		for _, match := range matchday {
			fmt.Printf("Matchday %d: %s vs %s\n", matchdayCounter, match.HomeTeam.Name, match.AwayTeam.Name)
		}
		matchdayCounter++
	}
	return schedule
}

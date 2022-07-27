package divider

import (
	"fmt"
	"math"
	"math/rand"
)

type Team struct {
	Title        string
	Participants []string
}

func CreateTeams(nTeams int, people []string) []Team {
	rand.Shuffle(len(people), func(i, j int) {
		people[i], people[j] = people[j], people[i]
	})

	teams := make([]Team, 0, nTeams)

	teamSize := int(math.Round(float64(len(people)) / float64(nTeams)))
	if teamSize == 0 {
		teamSize = 1 // min team size
	}

	for i, j := 0, 0; i < nTeams && j < len(people); i, j = i+1, j+teamSize {
		teams = append(teams, Team{})
		teams[i].Title = fmt.Sprintf("Team %v", i+1)
		if i == nTeams-1 {
			teams[i].Participants = people[j:]
			break
		}
		teams[i].Participants = people[j : j+teamSize]
	}

	return teams
}

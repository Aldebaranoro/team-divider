package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/Aldebaranoro/team-divider/config"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	cfg := config.MustLoad()
	teams := createTeams(cfg.People, cfg.TeamsCount)
	printTeams(teams)
}

func createTeams(people []string, count int) [][]string {
	rand.Shuffle(len(people), func(i, j int) {
		people[i], people[j] = people[j], people[i]
	})
	teams := make([][]string, count)
	teamSize := int(math.Round(float64(len(people)) / float64(count)))
	for i, j := 0, 0; i < count && j < len(people); i, j = i+1, j+teamSize {
		if i == count-1 {
			teams[i] = people[j:]
			break
		}
		teams[i] = people[j : j+teamSize]
	}
	return teams
}

func printTeams(teams [][]string) {
	for i, team := range teams {
		fmt.Printf("Team %v -> %v\n", i+1, strings.Join(team, ", "))
	}
}

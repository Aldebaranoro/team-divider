package indexer

import "math/rand"

type Player struct {
	Name  string
	Index int
}

func IndexTeam(people []string) []Player {
	teamSize := len(people)

	indexes := make([]int, 0, teamSize)
	for i := 0; i < teamSize; i++ {
		indexes = append(indexes, i+1)
	}
	rand.Shuffle(teamSize, func(i, j int) {
		indexes[i], indexes[j] = indexes[j], indexes[i]
	})

	players := make([]Player, 0, teamSize)
	for i, name := range people {
		players = append(players, Player{Name: name, Index: indexes[i]})
	}

	return players
}

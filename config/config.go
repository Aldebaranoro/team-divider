package config

import (
	"flag"
	"log"
	"strings"
)

type Config struct {
	People     []string
	TeamsCount int
}

func MustLoad() Config {
	teamsCount := flag.Int("teams-count", 2, "number of teams")
	peopleStr := flag.String("people", "", "names of participants")
	flag.Parse()

	if *peopleStr == "" {
		log.Fatal("people are not specified")
	}
	people := strings.Fields(*peopleStr)

	return Config{
		People:     people,
		TeamsCount: *teamsCount,
	}
}

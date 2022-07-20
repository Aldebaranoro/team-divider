package main

import (
	"github.com/Aldebaranoro/team-divider/cmd"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	cmd.Execute()
}

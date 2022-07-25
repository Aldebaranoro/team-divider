package main

import (
	"github.com/Aldebaranoro/team-divider/internal/command"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	cmdRoot := command.NewCmdRoot()
	if err := cmdRoot.Execute(); err != nil {
		os.Exit(1)
	}
}

package main

import (
	"github.com/Aldebaranoro/team-divider/internal/command"
	"github.com/Aldebaranoro/team-divider/internal/config"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	cfg := config.GetConfigInstance()

	cmdRoot := command.NewCmdRoot(cfg.Project.Build)
	if err := cmdRoot.Execute(); err != nil {
		os.Exit(1)
	}
}

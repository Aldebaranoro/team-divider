package main

import (
	"github.com/Aldebaranoro/team-divider/internal/command"
	"github.com/Aldebaranoro/team-divider/internal/config"
	"github.com/rs/zerolog/log"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}
	cfg := config.GetConfigInstance()

	cmdRoot := command.NewCmdRoot(cfg.Project.Build)
	if err := cmdRoot.Execute(); err != nil {
		os.Exit(1)
	}
}

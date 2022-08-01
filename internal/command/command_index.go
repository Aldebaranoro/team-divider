package command

import (
	"fmt"
	"github.com/Aldebaranoro/team-divider/internal/indexer"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func NewCmdIndex() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "index player_name...",
		Short:   "Randomly indexes team members",
		Example: `$ team-divider index Nobby Samuel Fred "Havelock Vetinari"`,
		Args:    cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			players := indexer.IndexTeam(args)
			for _, player := range players {
				c := color.New(color.FgYellow).Add(color.Bold).Add(color.Italic)
				fmt.Println(c.Sprint(player.Name), "->", player.Index)
			}
		},
	}

	return cmd
}

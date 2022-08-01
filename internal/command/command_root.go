package command

import (
	"fmt"
	"github.com/Aldebaranoro/team-divider/internal/config"
	"github.com/Aldebaranoro/team-divider/internal/divider"
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

func NewCmdRoot(build config.Build) *cobra.Command {
	cmd := &cobra.Command{
		Version: fmt.Sprintf("%v\ngit commit: %v\nbuilt: %v", build.Version, build.Commit, build.Date),
		Use:     "team-divider player_name...",
		Short:   "Randomly divides participants into different teams",
		Example: heredoc.Doc(`
			$ team-divider Nobby Samuel Fred "Havelock Vetinari"
			$ team-divider Nobby Samuel Fred "Havelock Vetinari" -n 3
			$ team-divider Nobby Samuel Fred "Havelock Vetinari" --teams 3`),
		Args: cobra.MinimumNArgs(2),
		Run:  rootRunFunc,
	}

	cmd.Flags().IntP("teams", "n", 2, "number of teams")

	cmd.AddCommand(NewCmdIndex())

	return cmd
}

func rootRunFunc(cmd *cobra.Command, args []string) {
	nTeams, _ := strconv.Atoi(cmd.Flag("teams").Value.String())
	if nTeams < 2 || nTeams > 100 {
		msg := color.RedString("team-divider: invalid number of teams: '%v'. " +
			"It should be in the range from 2 to 100")
		fmt.Printf(msg, nTeams)
		return
	}

	teams := divider.CreateTeams(nTeams, args)
	for _, team := range teams {
		c := color.New(color.FgCyan).Add(color.Bold).Add(color.Italic)
		fmt.Println(
			c.Sprint(team.Title),
			"->",
			strings.Join(team.Participants, ", "),
		)
	}
}

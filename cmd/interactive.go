package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ooizhenyi/GoLangCLI/ui"
	"github.com/spf13/cobra"
)

var interactiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "Start interactive file manager",
	Long:  `Starts the interactive TUI mode for exploring and managing files.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		dir, _ := cmd.Flags().GetString("dir")
		if dir == "." {
			var err error
			dir, err = os.Getwd()
			if err != nil {
				return err
			}
		}

		p := tea.NewProgram(ui.InitialModel(dir))
		if _, err := p.Run(); err != nil {
			return fmt.Errorf("error running program: %v", err)
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(interactiveCmd)
}

package cmd

import (
	"fmt"

	"github.com/diegiwg/devopness-cli/core"
	"github.com/diegiwg/devopness-cli/service"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(projectCmd)
	projectCmd.AddCommand(projectListCmd)
}

var project = service.Project{}

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Manage projects",
	Long:  "Provides commands to manage projects, including listing, creating, and modifying them.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var projectListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all projects",
	Long:  "Retrieves a list of all projects associated with the authenticated user.",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := core.NewContext()
		if err := ctx.LoadFromFile(); err != nil {
			fmt.Printf("Error loading context: %v\n", err)
			return
		}

		if !ctx.Authenticated {
			fmt.Println("Authentication required. Please log in first.")
			return
		}

		response := project.List(ctx)

		response.Dump()
	},
}

package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/diegiwg/devopness-cli/core"
	"github.com/diegiwg/devopness-cli/generated/services"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(projectCmd)
	projectCmd.AddCommand(projectListCmd)
}

var projects = service.Projects{}

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

		response := projects.ListProjects(ctx)

		if response == nil {
			fmt.Println("Empty response from API")
			return
		}

		prettyJSON, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			fmt.Printf("Error marshaling JSON: %v\n", err)
			return
		}

		fmt.Println(string(prettyJSON))
	},
}

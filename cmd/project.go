package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/diegiwg/devopness-cli/core"
	service "github.com/diegiwg/devopness-cli/generated"
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

		type ProjectListResponse []struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		}

		rawResponse := projects.ListProjects(ctx)

		var response ProjectListResponse
		err := json.Unmarshal([]byte(rawResponse), &response)
		if err != nil {
			panic(err)
		}

		prettyJSON, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			panic(err)
		}

		println(string(prettyJSON))
	},
}

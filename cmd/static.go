package cmd

import (
	"fmt"

	"github.com/diegiwg/devopness-cli/core"
	"github.com/diegiwg/devopness-cli/service"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(staticCmd)
	staticCmd.AddCommand(credentialOptionsCmd)
}

var static = service.Static{}

var staticCmd = &cobra.Command{
	Use:   "static",
	Short: "Retrieve static data from the API",
	Long:  "Provides access to static data such as predefined options and reference information.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var credentialOptionsCmd = &cobra.Command{
	Use:   "credential-options",
	Short: "List available credential options",
	Long:  "Fetches a list of supported credential types and their configurations from the API.",
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

		response := static.CredentialOptions(ctx)

		response.Dump()
	},
}

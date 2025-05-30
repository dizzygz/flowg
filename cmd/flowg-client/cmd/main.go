package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"link-society.com/flowg/internal/utils/client"
)

var ExitCode int = 0

func NewRootCommand() *cobra.Command {
	cobra.EnableTraverseRunHooks = true

	type globalOptions struct {
		apiUrl     string
		apiToken   string
		mgmtApiUrl string
	}

	opts := &globalOptions{}

	rootCmd := &cobra.Command{
		Use:   "flowg-client",
		Short: "API Client for FlowG",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			ctx = context.WithValue(ctx, ApiClient, client.NewClient(opts.apiUrl, opts.apiToken))
			ctx = context.WithValue(ctx, MgmtApiClient, client.NewClient(opts.mgmtApiUrl, ""))
			cmd.SetContext(ctx)
		},
	}

	rootCmd.Flags().StringVar(
		&opts.apiUrl,
		"api",
		defaultApiUrl,
		"URL to FlowG HTTP API",
	)

	rootCmd.Flags().StringVar(
		&opts.apiToken,
		"api-token",
		defaultApiToken,
		"Authentication token for FlowG HTTP API",
	)

	rootCmd.Flags().StringVar(
		&opts.mgmtApiUrl,
		"mgmt-api",
		defaultMgmtApiUrl,
		"URL to FlowG Management HTTP API",
	)

	rootCmd.AddCommand(
		NewLoginCommand(),
		NewStreamCommand(),
		NewPipelineCommand(),
		NewTransformerCommand(),
		NewForwarderCommand(),
		NewAclCommand(),
		NewTokenCommand(),
		NewAdminCommand(),
	)

	return rootCmd
}

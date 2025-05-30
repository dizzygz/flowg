package cmd

import (
	"fmt"
	"os"

	"net/http"

	"github.com/spf13/cobra"

	"link-society.com/flowg/internal/utils/client"
)

func NewPipelineDeleteCommand() *cobra.Command {
	type options struct {
		name string
	}

	opts := &options{}

	cmd := &cobra.Command{
		Use:   "rm",
		Short: "Delete a pipeline",
		Run: func(cmd *cobra.Command, args []string) {
			client := cmd.Context().Value(ApiClient).(*client.Client)
			url := fmt.Sprintf("/api/v1/pipelines/%s", opts.name)
			req, err := http.NewRequest(http.MethodDelete, url, nil)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ERROR: Could not prepare request: %v\n", err)
				ExitCode = 1
				return
			}

			resp, err := client.Do(req)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ERROR: Could not send request: %v\n", err)
				ExitCode = 1
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				fmt.Fprintf(os.Stderr, "ERROR: Received non-200 response: %s\n", resp.Status)
				ExitCode = 1
				return
			}
		},
	}

	cmd.Flags().StringVar(
		&opts.name,
		"name",
		"default",
		"Name of the pipeline to delete",
	)

	return cmd
}

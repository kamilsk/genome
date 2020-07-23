package cmd

import "github.com/spf13/cobra"

// New returns the new root command.
func New() *cobra.Command {
	command := cobra.Command{
		Use:   "genome",
		Short: "build your app by DNA",
		Long:  "Build your app by DNA.",

		SilenceErrors: false,
		SilenceUsage:  true,
	}
	/* configure instance */
	command.AddCommand( /* related commands */ )
	return &command
}

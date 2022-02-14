package app

import (
	"context"
	"flag"
	"os"

	"github.com/spf13/cobra"

	"github.com/zach593/karmada/cmd/aggregated-apiserver/app/options"
	"github.com/zach593/karmada/pkg/version/sharedcommand"
)

// NewAggregatedApiserverCommand creates a *cobra.Command object with default parameters
func NewAggregatedApiserverCommand(ctx context.Context) *cobra.Command {
	opts := options.NewOptions()

	cmd := &cobra.Command{
		Use:  "karmada-aggregated-apiserver",
		Long: `Launch the karmada-aggregated-apiserver`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.Complete(); err != nil {
				return err
			}
			if err := opts.Validate(); err != nil {
				return err
			}
			if err := opts.Run(ctx); err != nil {
				return err
			}
			return nil
		},
	}

	opts.AddFlags(cmd.Flags())
	cmd.AddCommand(sharedcommand.NewCmdVersion(os.Stdout, "karmada-aggregated-apiserver"))
	cmd.Flags().AddGoFlagSet(flag.CommandLine)
	return cmd
}

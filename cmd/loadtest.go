package cmd

import (
	"github.com/spf13/cobra"

	"github.com/mattermost/mattermost-load-test-ops/ops"
)

var loadTest = &cobra.Command{
	Use:   "loadtest -- [args...]",
	Short: "Runs a mattermost-load-test command againt the given cluster",
	RunE: func(cmd *cobra.Command, args []string) error {
		clusterName, _ := cmd.Flags().GetString("cluster-name")
		config, _ := cmd.Flags().GetString("config")
		return ops.LoadTest(clusterName, config, args)
	},
}

func init() {
	loadTest.Flags().String("cluster-name", "", "the name of the cluster to loadtest to (required)")
	loadTest.MarkFlagRequired("cluster-name")

	loadTest.Flags().String("config", "c", "a config file to use instead of the default (the ConnectionConfiguration section is mostly ignored)")

	rootCmd.AddCommand(loadTest)
}

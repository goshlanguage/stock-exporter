package cmd

import (
	"github.com/spf13/cobra"

	"github.com/goshlanguage/stock-exporter/pkg/exporter"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "stock-exporter",
	Short: "stock-exporter",
	Long:  `stock-exporter`,
	Run: func(cmd *cobra.Command, args []string) {
		watchList := []string{
			"AMZN",
			"JEPI",
			"JEPQ",
			"MSFT",
			"NVDA",
			"TQQQ",
			"TSLA",
		}

		exporter := exporter.NewStockPriceExporter(watchList)

		exporter.Serve("0.0.0.0", "9090")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err!= nil {
		// If an error occurs, it will be of type *cobra.CommandError
		// which is a wrapper around an error that contains the underlying error.
		// We can then use the underlying error to determine
		// what to do next.
		panic(err)
	}
}

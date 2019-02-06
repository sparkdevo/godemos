package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of cobrademo",
	Long:  `All software has versions. This is cobrademo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cobrademo version is v1.0")
	},
}
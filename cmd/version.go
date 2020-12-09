package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gocrm/config"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "example:gocrm version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gocrm-" + config.Version)
	},
}

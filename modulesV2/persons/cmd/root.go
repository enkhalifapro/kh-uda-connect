package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "persons-service",
		Short: "persons service manages persons data",
		Long:  `persons service manages persons data`,
		Run: func(cmd *cobra.Command, args []string) {
			
		},
	}
)

// Execute ...
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.AutomaticEnv()
}

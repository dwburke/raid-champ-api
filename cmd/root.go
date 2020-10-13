package cmd

import (
	"fmt"
	"os"

	"github.com/dwburke/vipertools"
	"github.com/spf13/cobra"

	"github.com/dwburke/raid-champ-api/api"
	"github.com/dwburke/raid-champ-api/logger"
)

var cfgFile string

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "./raid-champ-api.yml", "config file (default is ./raid-champ-api.yml)")
	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(logger.InitLogging)
}

var rootCmd = &cobra.Command{
	Use:   "raid-champ-api",
	Short: "raid-champ-api is a thing",
	Long:  `use with care`,
	Run: func(cmd *cobra.Command, args []string) {
		api.Run()
		<-api.ShutdownCh
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	if _, err := os.Stat(cfgFile); err == nil {
		if err := vipertools.MergeConfigs([]string{cfgFile}); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

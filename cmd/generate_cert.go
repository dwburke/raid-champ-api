package cmd

import (
	"fmt"
	"log"

	"github.com/kabukky/httpscerts"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(generateCertCmd)
}

var generateCertCmd = &cobra.Command{
	Use:   "generatecert",
	Short: "Generate self-signed certificate",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Checking:", viper.GetString("api.server.ssl-cert"))
		fmt.Println("Checking:", viper.GetString("api.server.ssl-key"))

		err := httpscerts.Check(
			viper.GetString("api.server.ssl-cert"),
			viper.GetString("api.server.ssl-key"),
		)
		// If they are not available, generate new ones.
		if err != nil {
			err = httpscerts.Generate(
				viper.GetString("api.server.ssl-cert"),
				viper.GetString("api.server.ssl-key"),
				"localhost:"+viper.GetString("api.server.port"),
			)
			if err != nil {
				log.Fatal("Error: Couldn't create https certs.")
			}
		} else {
			fmt.Println("Cert already exists.")
		}
	},
}

package cmd

import (
	"log"
	"os"

	"github.com/budimanjojo/talhelper/pkg/config"
	"github.com/budimanjojo/talhelper/pkg/secret"
	"github.com/spf13/cobra"
	"sigs.k8s.io/yaml"
)

var gensecretCmd = &cobra.Command{
	Use:   "gensecret",
	Short: "Generate Talos cluster secrets",
	Run: func(cmd *cobra.Command, args []string) {
		cf, err := os.ReadFile(configFile)
		if err != nil {
			log.Fatalf("failed to decrypt/read file: %s", err)
		}

		var m config.TalhelperConfig

		err = yaml.Unmarshal(cf, &m)
		if err != nil {
			log.Fatalf("failed to unmarshal config file: %s", err)
		}

		input, err := config.ParseTalosInput(m)
		if err != nil {
			log.Fatalf("failed to generate talos input: %s", err)
		}

		secret.PrintSortedSecrets(input)

		if patchConfig {
			secret.GenerateSecret(m, configFile)
		}
	},
}

func init() {
	rootCmd.AddCommand(gensecretCmd)

	gensecretCmd.Flags().StringVarP(&configFile, "config-file", "c", "talconfig.yaml", "File containing configurations for talhelper")
	gensecretCmd.Flags().BoolVarP(&patchConfig, "patch-configfile", "p", false, "Whether to generate inline patches into config file")
}

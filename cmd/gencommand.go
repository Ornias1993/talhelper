package cmd

import (
	"github.com/spf13/cobra"
)

var (
	gencommandCfgFile    string
	gencommandOutDir     string
	gencommandEnvFile    []string
	gencommandExtraFlags []string
	gencommandNode       string
)

var GencommandCmd = &cobra.Command{
	Use:   "gencommand",
	Short: "Generate commands for talosctl.",
}

func init() {
	RootCmd.AddCommand(GencommandCmd)
	GencommandCmd.PersistentFlags().StringVarP(&gencommandCfgFile, "config-file", "c", "talconfig.yaml", "File containing configurations for talhelper")
	GencommandCmd.PersistentFlags().StringVarP(&gencommandOutDir, "out-dir", "o", "./clusterconfig", "Directory that contains the generated config files to apply.")
	GencommandCmd.PersistentFlags().StringSliceVar(&gencommandEnvFile, "env-file", []string{"talenv.yaml", "talenv.sops.yaml", "talenv.yml", "talenv.sops.yml"}, "List of files containing env variables for config file")
	GencommandCmd.PersistentFlags().StringSliceVar(&gencommandExtraFlags, "extra-flags", []string{}, "List of additional flags that will be injected into the generated commands.")
	GencommandCmd.PersistentFlags().StringVarP(&gencommandNode, "node", "n", "", "A specific node to generate the command for. If not specified, will generate for all nodes.")
}

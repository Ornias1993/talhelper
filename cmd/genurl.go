package cmd

import (
	"github.com/budimanjojo/talhelper/v3/pkg/config"
	"github.com/spf13/cobra"
)

var (
	genurlCfgFile     string
	genurlEnvFile     []string
	genurlNode        string
	genurlRegistryURL string
	genurlVersion     string
	genurlExtensions  []string
	genurlKernelArgs  []string
	genurlOfflineMode bool
	genurlSecureboot  bool
)

var GenurlCmd = &cobra.Command{
	Use:   "genurl",
	Short: "Generate URL for Talos installer or ISO",
}

func init() {
	RootCmd.AddCommand(GenurlCmd)

	GenurlCmd.PersistentFlags().StringVarP(&genurlCfgFile, "config-file", "c", "talconfig.yaml", "File containing configurations for talhelper")
	GenurlCmd.PersistentFlags().StringSliceVar(&genurlEnvFile, "env-file", []string{"talenv.yaml", "talenv.sops.yaml", "talenv.yml", "talenv.sops.yml"}, "List of files containing env variables for config file")
	GenurlCmd.PersistentFlags().StringVarP(&genurlNode, "node", "n", "", "A specific node to generate command for. If not specified, will generate for all nodes (ignored when talconfig.yaml is not found)")
	GenurlCmd.PersistentFlags().StringVarP(&genurlRegistryURL, "registry-url", "r", "factory.talos.dev", "Registry url of the image")
	GenurlCmd.PersistentFlags().StringVarP(&genurlVersion, "version", "v", config.LatestTalosVersion, "Talos version to generate (defaults to latest Talos version)")
	GenurlCmd.PersistentFlags().StringSliceVarP(&genurlExtensions, "extension", "e", []string{}, "Official extension image to be included in the image (ignored when talconfig.yaml is found)")
	GenurlCmd.PersistentFlags().StringSliceVarP(&genurlKernelArgs, "kernel-arg", "k", []string{}, "Kernel arguments to be passed to the image kernel (ignored when talconfig.yaml is found)")
	GenurlCmd.PersistentFlags().BoolVar(&genurlOfflineMode, "offline-mode", false, "Generate schematic ID without doing POST request to image-factory")
	GenurlCmd.PersistentFlags().BoolVar(&genurlSecureboot, "secure-boot", false, "Whether to generate Secure Boot enabled URL")
}

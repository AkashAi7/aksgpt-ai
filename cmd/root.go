/*
Copyright 2023 The aksgpt Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/aksgpt-ai/aksgpt/cmd/analyze"
	"github.com/aksgpt-ai/aksgpt/cmd/auth"
	"github.com/aksgpt-ai/aksgpt/cmd/cache"
	"github.com/aksgpt-ai/aksgpt/cmd/filters"
	"github.com/aksgpt-ai/aksgpt/cmd/generate"
	"github.com/aksgpt-ai/aksgpt/cmd/integration"
	"github.com/aksgpt-ai/aksgpt/cmd/serve"
	"github.com/aksgpt-ai/aksgpt/pkg/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile     string
	kubecontext string
	kubeconfig  string
	Version     string
	Commit      string
	Date        string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aksgpt",
	Short: "Kubernetes debugging powered by AI",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(v string, c string, d string) {
	Version = v
	Commit = c
	Date = d
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	performConfigMigrationIfNeeded()

	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(auth.AuthCmd)
	rootCmd.AddCommand(analyze.AnalyzeCmd)
	rootCmd.AddCommand(filters.FiltersCmd)
	rootCmd.AddCommand(generate.GenerateCmd)
	rootCmd.AddCommand(integration.IntegrationCmd)
	rootCmd.AddCommand(serve.ServeCmd)
	rootCmd.AddCommand(cache.CacheCmd)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("Default config file (%s/aksgpt/aksgpt.yaml)", xdg.ConfigHome))
	rootCmd.PersistentFlags().StringVar(&kubecontext, "kubecontext", "", "Kubernetes context to use. Only required if out-of-cluster.")
	rootCmd.PersistentFlags().StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// the config will belocated under `~/.config/aksgpt/aksgpt.yaml` on linux
		configDir := filepath.Join(xdg.ConfigHome, "aksgpt")

		viper.AddConfigPath(configDir)
		viper.SetConfigType("yaml")
		viper.SetConfigName("aksgpt")

		_ = viper.SafeWriteConfig()
	}

	viper.Set("kubecontext", kubecontext)
	viper.Set("kubeconfig", kubeconfig)

	viper.SetEnvPrefix("aksgpt")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		_ = 1
		//	fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func performConfigMigrationIfNeeded() {
	oldConfig, err := getLegacyConfigFilePath()
	cobra.CheckErr(err)
	oldConfigExists, err := util.FileExists(oldConfig)
	cobra.CheckErr(err)

	newConfig := getConfigFilePath()
	newConfigExists, err := util.FileExists(newConfig)
	cobra.CheckErr(err)

	configDir := filepath.Dir(newConfig)
	err = util.EnsureDirExists(configDir)
	cobra.CheckErr(err)

	if oldConfigExists && !newConfigExists {
		err = os.Rename(oldConfig, newConfig)
		cobra.CheckErr(err)
	}
}

func getConfigFilePath() string {
	return filepath.Join(xdg.ConfigHome, "aksgpt", "aksgpt.yaml")
}

func getLegacyConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	return filepath.Join(home, ".aksgpt.yaml"), nil
}

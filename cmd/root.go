package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/intheclouddan/launchdarkly-code-generator/version"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "lcg",
	Short:   "Generate wrapper scripts for LaunchDarkly",
	Long:    `Generate wrapper scripts for LaunchDarkly`,
	Version: version.Version,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .launchdarkly/lcg.yaml)")
	rootCmd.SetVersionTemplate("launchdarkly-code-generation version: {{.Version}}\n")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in .launchdarkly directory with name "lcg.yaml".
		viper.AddConfigPath(".launchdarkly")
		viper.SetConfigName("lcg")
		viper.SetConfigType("yaml")
	}

	viper.SetEnvPrefix("LD")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

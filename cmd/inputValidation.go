package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func localOverridesValidation() {
	if (viper.GetStringSlice("localKey") != nil) && len(viper.GetStringSlice("localKey")) != 0 {
		if (viper.GetStringSlice("localDefault") != nil) && len(viper.GetStringSlice("localDefault")) == 0 {
			fmt.Println("Both a --localKey and --localDefault need to be passed in")
			os.Exit(1)
		}
	}
	if (viper.GetStringSlice("localDefault") != nil) && len(viper.GetStringSlice("localDefault")) != 0 {
		if (viper.GetStringSlice("localKey") != nil) && len(viper.GetStringSlice("localKey")) == 0 {
			fmt.Println("Both a --localKey and --localDefault need to be passed in")
			os.Exit(1)
		}
	}
}

func templateValidation() {
	if len(viper.GetString("language")) == 0 && len(viper.GetString("inputTemplate")) == 0 {
		fmt.Println("--language or --inputTemplate must be set")
		os.Exit(1)
	}
	if len(viper.GetString("language")) != 0 && len(viper.GetString("inputTemplate")) != 0 {
		fmt.Println("--language and --inputTemplate are mutually exclusive flags")
		os.Exit(1)
	}
}

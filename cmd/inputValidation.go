package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func localOverridesValidation() {
	if len(viper.GetString("localKey")) != 0 {
		if len(viper.GetString("localDefault")) == 0 {
			fmt.Println("Both a --localKey and --localDefault need to be passed in")
			os.Exit(1)
		}
	}
	if len(viper.GetString("localDefault")) != 0 {
		if len(viper.GetString("localKey")) == 0 {
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

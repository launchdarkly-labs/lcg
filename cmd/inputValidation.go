package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func localOverridesValidation(localFlag []string) {
	if len(localFlag) > 0 {
		for _, flag := range localFlag {
			flagParts := strings.Split(flag, ",")
			if len(flagParts) != 3 {
				fmt.Printf("Wrong number of argument passed in for flag: %s", flagParts[0])
				os.Exit(1)
			}
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

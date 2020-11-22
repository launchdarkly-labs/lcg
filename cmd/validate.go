/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		validateFlagFile()
	},
}

var flagFile string
var clean bool
var git bool

func init() {
	validateCmd.PersistentFlags().StringVar(&flagFile, "flagFile", "", "LaunchDarkly Project to query for flags")
	if err := viper.BindPFlag("flagFile", validateCmd.PersistentFlags().Lookup("flagFile")); err != nil {
		check(err)
	}
	validateCmd.PersistentFlags().BoolVar(&clean, "clean", false, "LaunchDarkly Project to query for flags")
	if err := viper.BindPFlag("clean", validateCmd.PersistentFlags().Lookup("clean")); err != nil {
		check(err)
	}
	validateCmd.PersistentFlags().BoolVar(&git, "git", false, "Check if file is staged for git commit")
	if err := viper.BindPFlag("git", validateCmd.PersistentFlags().Lookup("git")); err != nil {
		check(err)
	}
	rootCmd.AddCommand(validateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// validateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// validateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func validateFlagFile() {
	readFile, err := ioutil.ReadFile(viper.GetString("flagFile"))
	check(err)
	outFile := removeLocalFlags(string(readFile), "LOCAL_LCG_FLAGS_BEGIN", "LOCAL_LCG_FLAGS_END")
	if viper.GetBool("clean") {
		if err := ioutil.WriteFile(viper.GetString("flagFile"), []byte(outFile), 0644); err != nil {
			check(err)
		}
	}
}

func removeLocalFlags(str string, start string, end string) (result string) {
	lines := strings.Split(str, "\n")
	var startLine int
	var endLine int
	var outFile []string
	for idx, line := range lines {
		s := strings.Index(line, "LOCAL_LCG_FLAGS_BEGIN")
		if s > 0 {
			if viper.GetBool("git") && !(viper.GetBool("clean")) {
				cmd := exec.Command("git", "diff", "--cached", "--exit-code", "--", viper.GetString("flagFile"))
				err := cmd.Run()
				var (
					ee *exec.ExitError
					pe *os.PathError
				)
				if errors.As(err, &ee) {
					fmt.Println("Local flags found staged.") // ran, but non-zero exit code
					os.Exit(ee.ExitCode())

				} else if errors.As(err, &pe) {
					fmt.Printf("os.PathError: %v", pe)

				} else if err != nil {
					fmt.Printf("%v", err)
					os.Exit(ee.ExitCode())

				}
			} else if !(viper.GetBool("clean")) {
				fmt.Println("Local Flags found!")
				os.Exit(1)
			}
		}
		if s == -1 {
			continue
		}
		startLine = idx
		break
	}
	for idx, line := range lines {
		s := strings.Index(line, "LOCAL_LCG_FLAGS_END")
		if s == -1 {
			continue
		}
		endLine = idx + 1
		break
	}
	if startLine == -1 && endLine == -1 {
		return str
	}
	outFile = append(outFile, lines[0:startLine]...)
	outFile = append(outFile, lines[endLine:]...)
	return strings.Join(outFile, "\n")
}

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
	"fmt"
	"io/ioutil"
	"os"
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

var inFile string
var clean bool

func init() {
	validateCmd.PersistentFlags().StringVar(&inFile, "inFile", "", "LaunchDarkly Project to query for flags")
	viper.BindPFlag("inFile", validateCmd.PersistentFlags().Lookup("inFile"))
	validateCmd.PersistentFlags().BoolVar(&clean, "clean", false, "LaunchDarkly Project to query for flags")
	viper.BindPFlag("clean", validateCmd.PersistentFlags().Lookup("clean"))
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

	// inFile, err := os.Open()
	// check(err)
	readFile, err := ioutil.ReadFile(viper.GetString("inFile"))
	check(err)
	outFile := removeLocalFlags(string(readFile), "LOCAL_LCG_FLAGS_BEGIN", "LOCAL_LCG_FLAGS_END")

	fmt.Println(outFile)
	// lines := strings.Split(string(readFile), "\n")
	// outFile := []string{}
	// for idx, line := range lines {
	// 	// only print after %startHERE
	// 	equalStart := strings.Index(line, "LOCAL_LCG_FLAGS_BEGIN")
	// 	if equalStart > 0 {
	// 		if !(viper.GetBool("clean")) {
	// 			fmt.Println("Local Flags found!")
	// 			os.Exit(1)
	// 		}
	// 	}

	// 	equalEnd := strings.Index(line, "LOCAL_LCG_FLAGS_END")
	// 	fmt.Println(equalEnd)
	// 	if (equalStart == -1) || (idx < equalStart) {
	// 		outFile = append(outFile, line)
	// 	} else if equalEnd != -1 && idx > equalEnd {
	// 		fmt.Println(equalEnd)
	// 		outFile = append(outFile, line)
	// 	}
	//if startPrint {
	//fmt.Println(i, line)
	// instead of printing to the screen, this is where
	// you want to start writing to a new file a.k.a copying the file

	// how it is done at https://www.socketloop.com/tutorials/golang-simple-file-scaning-and-remove-virus-example
	//}
	//}
	//fmt.Println(outFile)

}

func removeLocalFlags(str string, start string, end string) (result string) {
	lines := strings.Split(str, "\n")
	var startLine int
	var endLine int
	var outFile []string
	for idx, line := range lines {
		s := strings.Index(line, "LOCAL_LCG_FLAGS_BEGIN")
		if s > 0 {
			if !(viper.GetBool("clean")) {
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

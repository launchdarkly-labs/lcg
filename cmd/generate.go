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
	"strings"

	"github.com/antihax/optional"
	"github.com/aymerick/raymond"
	"github.com/intheclouddan/launchdarkly-code-generator/launchdarkly"
	ldapi "github.com/launchdarkly/api-client-go"
	"github.com/markbates/pkger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Output a language wrapper for LaunchDarkly Feature Flags",
	Long:  `Output a language wrapper for LaunchDarkly Feature Flags.`,
	Run: func(cmd *cobra.Command, args []string) {
		generateTemplate()
	},
}

var language string
var apiToken string
var projectKey string
var outFile string
var baseUri string
var tags string

func init() {
	generateCmd.PersistentFlags().StringVarP(&language, "language", "l", "", "language files to output")
	viper.BindPFlag("language", generateCmd.PersistentFlags().Lookup("language"))
	viper.SetDefault("language", "node-typescript")
	generateCmd.PersistentFlags().StringVarP(&apiToken, "apiToken", "k", "", "LaunchDarkly API Token")
	viper.BindPFlag("apiToken", generateCmd.PersistentFlags().Lookup("apiToken"))
	generateCmd.PersistentFlags().StringVarP(&projectKey, "projectKey", "p", "", "LaunchDarkly Project to query for flags")
	viper.BindPFlag("projectKey", generateCmd.PersistentFlags().Lookup("projectKey"))
	generateCmd.PersistentFlags().StringVarP(&outFile, "outFile", "o", "", "Out file")
	viper.BindPFlag("outFile", generateCmd.PersistentFlags().Lookup("outFile"))
	generateCmd.PersistentFlags().StringVarP(&baseUri, "baseUri", "b", "", "LaunchDarkly Instance")
	viper.BindPFlag("baseUri", generateCmd.PersistentFlags().Lookup("baseUri"))
	viper.SetDefault("baseUri", "https://app.launchdarkly.com")
	generateCmd.PersistentFlags().StringVarP(&tags, "tags", "t", "", "Filter flags to specific tag")
	viper.BindPFlag("tags", generateCmd.PersistentFlags().Lookup("tags"))
	rootCmd.AddCommand(generateCmd)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func generateTemplate() {
	templateFile, err := pkger.Open(strings.Join([]string{"/templates/", viper.GetString("language"), ".hbs"}, ""))
	check(err)
	dat, err := ioutil.ReadAll(templateFile)
	check(err)
	flags, err := queryAPI()
	if err != nil {
		panic(err)
	}

	raymond.RegisterHelper("defaultValue", func(flag ldapi.FeatureFlag, quotes string) string {
		var quoteWrapper string
		if quotes == "single" {
			quoteWrapper = "'"
		} else {
			quoteWrapper = "\""
		}
		if flag.Defaults != nil {
			defaultVar := flag.Defaults.OffVariation
			tempVar := *flag.Variations[defaultVar].Value
			return parseReturnValues(tempVar, quoteWrapper)
		} else {
			offVar := flag.Variations[len(flag.Variations)-1]
			tempVar := *offVar.Value
			return parseReturnValues(tempVar, quoteWrapper)
		}
	})
	// Register raymond template helpers
	templateHelpers()
	//data, err := tpl.Exec(flags)
	data := raymond.MustRender(string(dat), flags)
	if err != nil {
		fmt.Printf("failed to parse flags: %s", err)
	}

	d1 := []byte(data)
	err = ioutil.WriteFile(outFile, d1, 0644)
	check(err)
	fmt.Printf("New LaunchDarkly flag generation written to %s\n", outFile)
}

func queryAPI() (ldapi.FeatureFlags, error) {
	client, err := launchdarkly.NewClient(&launchdarkly.LaunchdarklyConfig{viper.GetString("apiToken"), viper.GetString("baseUri")})
	if err != nil {
		fmt.Println(err)
		return ldapi.FeatureFlags{}, err
	}
	var tagFilter ldapi.GetFeatureFlagsOpts
	if viper.GetString("tags") != "" {
		tagFilter = ldapi.GetFeatureFlagsOpts{
			Tag: optional.NewString(viper.GetString("tags")),
		}
	}
	featureFlags, _, err := client.Ld.FeatureFlagsApi.GetFeatureFlags(client.Ctx, projectKey, &tagFilter)
	if err != nil {
		fmt.Println(err)
		return ldapi.FeatureFlags{}, err
	}

	return featureFlags, nil

}

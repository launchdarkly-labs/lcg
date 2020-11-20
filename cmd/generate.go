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
	"io"
	"io/ioutil"
	"os"
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

var (
	language        string
	apiToken        string
	projectKey      string
	outFile         string
	baseURI         string
	tags            string
	sdkAvailability string
	localKey        []string
	localDefault    []string
	localType       []string
	inputTemplate   string
)

func init() {
	generateCmd.PersistentFlags().StringVarP(&language, "language", "l", "", "language files to output")
	viper.BindPFlag("language", generateCmd.PersistentFlags().Lookup("language"))
	generateCmd.PersistentFlags().StringVarP(&apiToken, "apiToken", "k", "", "LaunchDarkly API Token")
	viper.BindPFlag("apiToken", generateCmd.PersistentFlags().Lookup("apiToken"))
	generateCmd.PersistentFlags().StringVarP(&projectKey, "projectKey", "p", "", "LaunchDarkly Project to query for flags")
	viper.BindPFlag("projectKey", generateCmd.PersistentFlags().Lookup("projectKey"))
	generateCmd.PersistentFlags().StringVarP(&outFile, "outFile", "o", "", "Out file")
	viper.BindPFlag("outFile", generateCmd.PersistentFlags().Lookup("outFile"))
	generateCmd.PersistentFlags().StringVarP(&baseURI, "baseUri", "b", "", "LaunchDarkly Instance")
	viper.BindPFlag("baseUri", generateCmd.PersistentFlags().Lookup("baseUri"))
	viper.SetDefault("baseUri", "https://app.launchdarkly.com")
	generateCmd.PersistentFlags().StringVarP(&tags, "tags", "t", "", "Filter flags to specific tag")
	viper.BindPFlag("tags", generateCmd.PersistentFlags().Lookup("tags"))
	generateCmd.PersistentFlags().StringVarP(&sdkAvailability, "sdkAvailability", "a", "", "Filter flags based on client side availability")
	viper.BindPFlag("sdkAvailability", generateCmd.PersistentFlags().Lookup("sdkAvailability"))
	generateCmd.PersistentFlags().StringSliceVar(&localKey, "localKey", nil, "Temporary local flags key appended to class")
	viper.BindPFlag("localKey", generateCmd.PersistentFlags().Lookup("localKey"))
	generateCmd.PersistentFlags().StringSliceVar(&localDefault, "localDefault", nil, "Temporary local default values appended to class")
	viper.BindPFlag("localDefault", generateCmd.PersistentFlags().Lookup("localDefault"))
	generateCmd.PersistentFlags().StringSliceVar(&localType, "localType", nil, "Temporary local return types appended to class")
	viper.BindPFlag("localType", generateCmd.PersistentFlags().Lookup("localType"))
	generateCmd.PersistentFlags().StringVarP(&inputTemplate, "inputTemplate", "", "", "Read in a template to be rendered")
	viper.BindPFlag("inputTemplate", generateCmd.PersistentFlags().Lookup("inputTemplate"))
	rootCmd.AddCommand(generateCmd)

}

func check(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

func generateTemplate() {
	localOverridesValidation()
	templateValidation()
	var templateFile io.Reader
	if len(viper.GetString("language")) != 0 {
		var err error
		templateFile, err = pkger.Open(strings.Join([]string{"/templates/", viper.GetString("language"), ".hbs"}, ""))
		check(err)
	} else {
		var err error
		templateFile, err = os.Open(viper.GetString("inputTemplate"))
		check(err)
	}
	dat, err := ioutil.ReadAll(templateFile)
	check(err)
	flags, err := queryAPI()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Register raymond template helpers
	templateHelpers()
	type localFlagss struct {
		Flag    string
		Default string
		Type    string
	}
	var localFlags []localFlagss
	//var localFlagMap map[string]interface{}
	if localKey != nil {
		for idx, flag := range localKey {
			localFlags = append(localFlags, localFlagss{
				Flag:    flag,
				Default: localDefault[idx],
				Type:    localType[idx],
			})
		}
	}
	//data, err := tpl.Exec(flags)
	ctx := make(map[string]interface{})
	ctx["flags"] = flags
	ctx["localFlags"] = localFlags

	data := raymond.MustRender(string(dat), ctx)
	if err != nil {
		fmt.Printf("failed to parse flags: %s", err)
	}

	d1 := []byte(data)
	localOutFile := viper.GetString("outFile")
	err = ioutil.WriteFile(localOutFile, d1, 0644)
	check(err)
	fmt.Printf("New LaunchDarkly flag generation written to %s\n", localOutFile)
}

func queryAPI() (ldapi.FeatureFlags, error) {
	client, err := launchdarkly.NewClient(&launchdarkly.LaunchdarklyConfig{AccessToken: viper.GetString("apiToken"), BaseUri: viper.GetString("baseUri")})
	if err != nil {
		return ldapi.FeatureFlags{}, err
	}

	// Build flag filters to determine which flags to return. Only supports tag and sdkAvailability
	var flagFilter ldapi.GetFeatureFlagsOpts
	if viper.GetString("tags") != "" {
		flagFilter.Tag = optional.NewString(viper.GetString("tags"))
	}
	if viper.GetString("sdkAvailability") != "" {
		flagFilter.Filter = optional.NewString(strings.Join([]string{"sdkAvailability", viper.GetString("sdkAvailability")}, ":"))
	}

	featureFlags, _, err := client.Ld.FeatureFlagsApi.GetFeatureFlags(client.Ctx, viper.GetString("projectKey"), &flagFilter)
	if err != nil {
		return ldapi.FeatureFlags{}, err
	}

	if len(featureFlags.Items) == 0 {
		fmt.Println("No flags found.")
		os.Exit(1)
	}
	return featureFlags, nil

}

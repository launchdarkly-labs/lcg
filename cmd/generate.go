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

type LocalFlagTemplate struct {
	Flag    string
	Default string
	Type    string
}

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
	baseURI         string
	tags            string
	sdkAvailability string
	localFlag       []string
	inputTemplate   string
)

func init() {
	generateCmd.PersistentFlags().StringVarP(&language, "language", "l", "", "language files to output")
	if err := viper.BindPFlag("language", generateCmd.PersistentFlags().Lookup("language")); err != nil {
		check(err)
	}
	generateCmd.PersistentFlags().StringVarP(&apiToken, "apiToken", "k", "", "LaunchDarkly API Token")
	if err := viper.BindPFlag("apiToken", generateCmd.PersistentFlags().Lookup("apiToken")); err != nil {
		check(err)
	}
	generateCmd.PersistentFlags().StringVarP(&projectKey, "projectKey", "p", "", "LaunchDarkly Project to query for flags")
	if err := viper.BindPFlag("projectKey", generateCmd.PersistentFlags().Lookup("projectKey")); err != nil {
		check(err)
	}
	generateCmd.PersistentFlags().StringVarP(&flagFile, "flagFile", "o", "", "Out file")
	if err := viper.BindPFlag("flagFile", generateCmd.PersistentFlags().Lookup("flagFile")); err != nil {
		check(err)
	}
	generateCmd.PersistentFlags().StringVarP(&baseURI, "baseUri", "b", "", "LaunchDarkly Instance")
	if err := viper.BindPFlag("baseUri", generateCmd.PersistentFlags().Lookup("baseUri")); err != nil {
		check(err)
	}
	viper.SetDefault("baseUri", "https://app.launchdarkly.com")
	generateCmd.PersistentFlags().StringVarP(&tags, "tags", "t", "", "Filter flags to specific tag")
	if err := viper.BindPFlag("tags", generateCmd.PersistentFlags().Lookup("tags")); err != nil {
		check(err)
	}
	generateCmd.PersistentFlags().StringVarP(&sdkAvailability, "sdkAvailability", "a", "", "Filter flags based on client side availability")
	if err := viper.BindPFlag("sdkAvailability", generateCmd.PersistentFlags().Lookup("sdkAvailability")); err != nil {
		check(err)
	}
	generateCmd.PersistentFlags().StringArrayVar(&localFlag, "localFlag", []string{}, "Temporary local flags appended to class")
	if err := viper.BindPFlag("localFlag", generateCmd.PersistentFlags().Lookup("localFlag")); err != nil {
		check(err)
	}
	generateCmd.PersistentFlags().StringVarP(&inputTemplate, "inputTemplate", "", "", "Read in a template to be rendered")
	if err := viper.BindPFlag("inputTemplate", generateCmd.PersistentFlags().Lookup("inputTemplate")); err != nil {
		check(err)
	}
	rootCmd.AddCommand(generateCmd)

}

func check(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

func generateTemplate() {
	localOverridesValidation(localFlag)
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

	var localFlags []LocalFlagTemplate
	for _, flag := range localFlag {
		flagParts := strings.Split(flag, ",")
		localFlags = append(localFlags, LocalFlagTemplate{
			Flag:    flagParts[0],
			Default: flagParts[1],
			Type:    flagParts[2],
		})
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
	localOutFile := viper.GetString("flagFile")
	err = ioutil.WriteFile(localOutFile, d1, 0644)
	check(err)
	fmt.Printf("New LaunchDarkly flag generation written to %s\n", localOutFile)
}

func queryAPI() (ldapi.FeatureFlags, error) {
	client, err := launchdarkly.NewClient(&launchdarkly.Config{AccessToken: viper.GetString("apiToken"), BaseUri: viper.GetString("baseUri")})
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
	fmt.Println(len(featureFlags.Items))
	if len(featureFlags.Items) == 0 {
		fmt.Println("No flags found.")
		os.Exit(1)
	}
	return featureFlags, nil

}

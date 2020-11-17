package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aymerick/raymond"
	"github.com/iancoleman/strcase"
	ldapi "github.com/launchdarkly/api-client-go"
	"github.com/spf13/viper"
)

func templateHelpers() {
	raymond.RegisterHelper("lowerCamelCase", func(name string) string {
		return strcase.ToLowerCamel(name)
	})

	raymond.RegisterHelper("snakeCase", func(name string) string {
		return strcase.ToSnake(name)
	})

	raymond.RegisterHelper("returnCheck", func(flag ldapi.FeatureFlag, options *raymond.Options) string {
		flagVar := *flag.Variations[0].Value
		switch s := flagVar.(type) {
		case float64:
			return options.DataStr("outNumber")
		case string:
			return options.DataStr("outString")
		case bool:
			return options.DataStr("outBool")
		case map[string]interface{}:
			return options.DataStr("outMap")
		case []interface{}:
			return options.DataStr("outMap")
		default:
			fmt.Printf("I don't know about type %T!\n", s)
			return ""
		}
	})
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

	raymond.RegisterHelper("outNumber", func(val1 string, options *raymond.Options) string {
		frame := options.DataFrame()
		frame.Set("outNumber", val1)
		return options.FnData(frame)
	})

	raymond.RegisterHelper("outBool", func(val1 string, options *raymond.Options) string {
		frame := options.DataFrame()
		frame.Set("outBool", val1)
		return options.FnData(frame)
	})

	raymond.RegisterHelper("outString", func(val1 string, options *raymond.Options) string {
		frame := options.DataFrame()
		frame.Set("outString", val1)
		return options.FnData(frame)
	})

	raymond.RegisterHelper("outMap", func(val1 string, options *raymond.Options) string {
		frame := options.DataFrame()
		frame.Set("outMap", val1)
		return options.FnData(frame)
	})

}

func parseReturnValues(tempVar interface{}, quoteWrapper string) string {
	varCheck := tempVar
	var returnVar string
	switch s := varCheck.(type) {
	case float64:
		returnVar = fmt.Sprintf("%v", tempVar)
	case string:
		returnVar = strings.Join([]string{quoteWrapper, fmt.Sprintf(`%s`, tempVar), quoteWrapper}, "")
	case bool:
		if viper.GetString("language") == "python" {
			returnVar = strings.Title(fmt.Sprintf("%v", tempVar))
		} else {
			returnVar = fmt.Sprintf("%v", tempVar)
		}
	case map[string]interface{}:
		jsonVal, err := json.Marshal(tempVar)
		if err != nil {
			panic(err)
		}
		returnVar = string(jsonVal)
	case []interface{}:
		jsonVal, err := json.Marshal(tempVar)
		if err != nil {
			panic(err)
		}
		returnVar = string(jsonVal)
	default:
		fmt.Printf("I don't know about type %T!\n", s)
		return ""
	}
	return returnVar
}

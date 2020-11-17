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
	raymond.RegisterHelper("camelCase", func(name string) string {
		return strcase.ToLowerCamel(name)
	})

	raymond.RegisterHelper("snakeCase", func(name string) string {
		return strcase.ToSnake(name)
	})

	raymond.RegisterHelper("returnCheck", func(flag ldapi.FeatureFlag) string {
		flagVar := *flag.Variations[0].Value
		switch s := flagVar.(type) {
		case float64:
			return "number"
		case string:
			return "string"
		case bool:
			return "boolean"
		case map[string]interface{}:
			return "Object"
		case []interface{}:
			return "Object"
		default:
			fmt.Printf("I don't know about type %T!\n", s)
			return ""
		}
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

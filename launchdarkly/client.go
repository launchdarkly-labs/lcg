package launchdarkly

import (
	"context"
	"fmt"

	ldapi "github.com/launchdarkly/api-client-go"
)

// The version string gets updated at build time using -ldflags
var version = "unreleased"

const (
	APIVersion = "20191212"
)

// Client is used by the provider to access the ld API.
type Client struct {
	ApiKey  string
	ApiHost string
	Ld      *ldapi.APIClient
	Ctx     context.Context
}

//Config is passed to NewClient
type Config struct {
	AccessToken string `json:"access_token"`
	BaseUri     string `json:"base_uri"`
}

//NewClient creates a newly configured LaunchDarkly API client
func NewClient(config *Config) (*Client, error) {
	basePath := fmt.Sprintf(`%s/api/v2`, config.BaseUri)

	cfg := &ldapi.Configuration{
		BasePath:      basePath,
		DefaultHeader: make(map[string]string),
		UserAgent:     fmt.Sprintf("launchdarkly-code-generator/0.0.1"),
	}

	cfg.AddDefaultHeader("LD-API-Version", APIVersion)

	ctx := context.WithValue(context.Background(), ldapi.ContextAPIKey, ldapi.APIKey{
		Key: config.AccessToken,
	})

	return &Client{
		ApiKey:  config.AccessToken,
		ApiHost: basePath,
		Ld:      ldapi.NewAPIClient(cfg),
		Ctx:     ctx,
	}, nil
}

package octoml

import (
	"errors"
	"net/http"
	"net/url"

	"glide/pkg/providers/clients"
	"glide/pkg/telemetry"
)

const (
	providerName = "octoml"
)

// ErrEmptyResponse is returned when the OctoML API returns an empty response.
var (
	ErrEmptyResponse = errors.New("empty response")
)

// Client is a client for accessing OctoML API
type Client struct {
	baseURL             string
	chatURL             string
	chatRequestTemplate *ChatRequest
	config              *Config
	httpClient          *http.Client
	telemetry           *telemetry.Telemetry
}

// NewClient creates a new OctoML client for the OctoML API.
func NewClient(providerConfig *Config, clientConfig *clients.ClientConfig, tel *telemetry.Telemetry) (*Client, error) {
	chatURL, err := url.JoinPath(providerConfig.BaseURL, providerConfig.ChatEndpoint)
	if err != nil {
		return nil, err
	}

	c := &Client{
		baseURL:             providerConfig.BaseURL,
		chatURL:             chatURL,
		config:              providerConfig,
		chatRequestTemplate: NewChatRequestFromConfig(providerConfig),
		httpClient: &http.Client{
			Timeout: *clientConfig.Timeout,
			// TODO: use values from the config
			Transport: &http.Transport{
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 2,
			},
		},
		telemetry: tel,
	}

	return c, nil
}

func (c *Client) Provider() string {
	return providerName
}

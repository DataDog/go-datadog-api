/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2013 by authors and contributors.
 */

package datadog

import "net/http"

// Client is the object that handles talking to the Datadog API. This maintains
// state information for a particular application connection.
type Client struct {
	apiKey, appKey, endpoint string

	//The Http Client that is used to make requests
	HttpClient *http.Client
}

// NewClient returns a new datadog.Client which can be used to access the API
// methods. The expected argument is the API key.
func NewClient(apiKey, appKey string) *Client {
	return &Client{
		apiKey:     apiKey,
		appKey:     appKey,
		HttpClient: http.DefaultClient,
	}
}

// SetEndpoint sets the API URL. Defaults to null, which will then fallback
// to the environment variable `"DATADOG_HOST"`.
func (c *Client) SetEndpoint(endpoint string) {
	c.endpoint = endpoint
}

// SetKeys changes the value of apiKey and appKey.
func (c *Client) SetKeys(apiKey, appKey string) {
	c.apiKey = apiKey
	c.appKey = appKey
}

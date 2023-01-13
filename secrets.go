// Copyright 2022 Paul Greenberg greenpau@outlook.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package secrets

import (
	"context"
	"fmt"
)

// Client provides interface to query AWS Secrets Manager service.
type Client interface {
	GetSecret(context.Context) (map[string]interface{}, error)
	GetSecretByKey(context.Context, string) (interface{}, error)
	GetConfig(context.Context) map[string]interface{}
}

type clientConfig struct {
	ID       string `json:"id,omitempty" xml:"id,omitempty" yaml:"id,omitempty"`
	Provider string `json:"provider,omitempty" xml:"provider,omitempty" yaml:"provider,omitempty"`
}

type client struct {
	config *clientConfig
	secret map[string]interface{}
}

// NewClient returns an instance of Client.
func NewClient(ctx context.Context, id string, secret map[string]interface{}) (Client, error) {
	if secret == nil {
		return nil, fmt.Errorf("secret is nil")
	}

	if len(secret) < 1 {
		return nil, fmt.Errorf("secret is empty")
	}

	if id == "" {
		return nil, fmt.Errorf("id is empty")
	}

	c := &client{
		config: &clientConfig{
			ID:       id,
			Provider: "static_secrets_manager",
		},
		secret: secret,
	}

	return c, nil
}

// GetSecret returns the key-value map of the stored secret.
func (c *client) GetSecret(ctx context.Context) (map[string]interface{}, error) {
	return c.secret, nil
}

// GetSecret returns the value of a key in a map of the stored secret.
func (c *client) GetSecretByKey(ctx context.Context, key string) (interface{}, error) {
	value, exists := c.secret[key]
	if !exists {
		return "", fmt.Errorf("key %q not found in %q secret", key, c.config.ID)
	}
	return value, nil
}

// GetConfig returns client configuration.
func (c *client) GetConfig(_ context.Context) map[string]interface{} {
	cfg := map[string]interface{}{
		"id":       c.config.ID,
		"provider": c.config.Provider,
	}
	return cfg
}

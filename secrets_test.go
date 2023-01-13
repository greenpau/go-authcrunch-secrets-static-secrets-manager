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
	"encoding/json"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// packMapToJSON converts a map to a JSON string.
func packMapToJSON(t *testing.T, m map[string]interface{}) string {
	b, err := json.Marshal(m)
	if err != nil {
		t.Fatalf("failed to marshal %v: %v", m, err)
	}
	return string(b)
}

func TestNewClient(t *testing.T) {
	testcases := []struct {
		name      string
		id        string
		secret    map[string]interface{}
		want      map[string]interface{}
		shouldErr bool
		err       error
	}{
		{
			name: "test new client with valid secret",
			id:   "foo",
			secret: map[string]interface{}{
				"username": "foo",
				"password": "bar",
			},
			want: map[string]interface{}{
				"id":       "foo",
				"provider": "static_secrets_manager",
			},
		},
		{
			name: "test new client without id",
			secret: map[string]interface{}{
				"username": "foo",
				"password": "bar",
			},
			shouldErr: true,
			err:       fmt.Errorf("id is empty"),
		},
		{
			name:      "test new client with nil secret",
			id:        "foo",
			shouldErr: true,
			err:       fmt.Errorf("secret is nil"),
		},
		{
			name:      "test new client with empty secret",
			id:        "foo",
			secret:    map[string]interface{}{},
			shouldErr: true,
			err:       fmt.Errorf("secret is empty"),
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c, err := NewClient(context.TODO(), tc.id, tc.secret)
			if err != nil {
				if !tc.shouldErr {
					t.Fatalf("expected success, got: %v", err)
				}
				if diff := cmp.Diff(tc.err.Error(), err.Error()); diff != "" {
					t.Logf("unexpected error: %v", err)
					t.Fatalf("NewClient() error mismatch (-want +got):\n%s", diff)
				}
				return
			}
			if tc.shouldErr {
				t.Fatalf("unexpected success, want: %v", tc.err)
			}

			got := c.GetConfig(context.TODO())
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Logf("JSON: %v", packMapToJSON(t, got))
				t.Errorf("NewClient() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestGetSecret(t *testing.T) {
	jsmith := map[string]interface{}{
		"api_key":  "bcrypt:10:$2a$10$TEQ7ZG9cAdWwhQK36orCGOlokqQA55ddE0WEsl00oLZh567okdcZ6",
		"email":    "jsmith@localhost.localdomain",
		"name":     "John Smith",
		"password": "bcrypt:10:$2a$10$iqq53VjdCwknBSBrnyLd9OH1Mfh6kqPezMMy6h6F41iLdVDkj13I6",
		"username": "jsmith",
	}

	testcases := []struct {
		name      string
		secret    map[string]interface{}
		want      map[string]interface{}
		shouldErr bool
		err       error
	}{
		{
			name:   "test valid user secret",
			secret: jsmith,
			want:   jsmith,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c, err := NewClient(context.TODO(), "foo", tc.secret)
			if err != nil {
				t.Fatalf("unxpected error during client initialization: %v", err)
			}

			got, err := c.GetSecret(context.TODO())

			if err != nil {
				if !tc.shouldErr {
					t.Fatalf("expected success, got: %v", err)
				}
				if diff := cmp.Diff(tc.err.Error(), err.Error()); diff != "" {
					t.Logf("unexpected error: %v", err)
					t.Fatalf("GetSecret() error mismatch (-want +got):\n%s", diff)
				}
				return
			}
			if tc.shouldErr {
				t.Fatalf("unexpected success, want: %v", tc.err)
			}

			if diff := cmp.Diff(packMapToJSON(t, tc.want), packMapToJSON(t, got)); diff != "" {
				t.Logf("JSON: %v", packMapToJSON(t, got))
				t.Errorf("GetSecret() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestGetSecretByKey(t *testing.T) {
	jsmith := map[string]interface{}{
		"api_key":  "bcrypt:10:$2a$10$TEQ7ZG9cAdWwhQK36orCGOlokqQA55ddE0WEsl00oLZh567okdcZ6",
		"email":    "jsmith@localhost.localdomain",
		"name":     "John Smith",
		"password": "bcrypt:10:$2a$10$iqq53VjdCwknBSBrnyLd9OH1Mfh6kqPezMMy6h6F41iLdVDkj13I6",
		"username": "jsmith",
	}

	testcases := []struct {
		name      string
		secret    map[string]interface{}
		key       string
		want      interface{}
		shouldErr bool
		err       error
	}{
		{
			name:   "test valid user secret key",
			secret: jsmith,
			key:    "username",
			want:   "jsmith",
		},
		{
			name:      "test user secret key not found",
			secret:    jsmith,
			key:       "bar",
			shouldErr: true,
			err:       fmt.Errorf("key %q not found in %q secret", "bar", "foo"),
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c, err := NewClient(context.TODO(), "foo", tc.secret)
			if err != nil {
				t.Fatalf("unxpected error during client initialization: %v", err)
			}

			got, err := c.GetSecretByKey(context.TODO(), tc.key)

			if err != nil {
				if !tc.shouldErr {
					t.Fatalf("expected success, got: %v", err)
				}
				if diff := cmp.Diff(tc.err.Error(), err.Error()); diff != "" {
					t.Logf("unexpected error: %v", err)
					t.Fatalf("GetSecretByKey() error mismatch (-want +got):\n%s", diff)
				}
				return
			}
			if tc.shouldErr {
				t.Fatalf("unexpected success, want: %v", tc.err)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("GetSecretByKey() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

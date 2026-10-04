// Copyright 2026 Google LLC
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

package config

import (
	"context"
	"testing"
)

func TestEnsureClusterCredentialsDoesNotUseGKEForEKS(t *testing.T) {
	t.Setenv("PATH", t.TempDir())
	cfg := &Config{
		Platform:       PlatformEKS,
		ProjectID:      "gke-project",
		ClusterName:    "gke-cluster",
		ClusterLocation: "us-east1",
	}
	if err := cfg.EnsureClusterCredentials(context.Background()); err != nil {
		t.Fatalf("EnsureClusterCredentials() = %v, want no GKE credential lookup", err)
	}
}

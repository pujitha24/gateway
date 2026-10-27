// Copyright Envoy Gateway Authors
// SPDX-License-Identifier: Apache-2.0
// The full text of the Apache license is available in the LICENSE file at
// the root of the repo.

package version

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEnvoyProxyTag(t *testing.T) {
	require.Equal(t, "distroless-dev", envoyProxyTag("docker.io/envoyproxy/envoy:distroless-dev"))
	require.Equal(t, "distroless-v1.39.5@sha256:abc123",
		envoyProxyTag("docker.io/envoyproxy/envoy:distroless-v1.39.5@sha256:abc123"))
}

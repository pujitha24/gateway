// Copyright Envoy Gateway Authors
// SPDX-License-Identifier: Apache-2.0
// The full text of the Apache license is available in the LICENSE file at
// the root of the repo.

package kubernetes

import (
	"testing"

	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func TestSetByObject(t *testing.T) {
	byObject := map[client.Object]cache.ByObject{
		&corev1.Secret{}:    {UnsafeDisableDeepCopy: new(true)},
		&corev1.ConfigMap{}: {UnsafeDisableDeepCopy: new(true)},
	}

	setByObject(byObject, &corev1.Secret{}, cache.ByObject{
		Namespaces: map[string]cache.Config{"envoy-gateway-system": {}},
	})

	require.Len(t, byObject, 2)
	for obj, opts := range byObject {
		if _, ok := obj.(*corev1.Secret); ok {
			require.Contains(t, opts.Namespaces, "envoy-gateway-system")
		}
	}
}

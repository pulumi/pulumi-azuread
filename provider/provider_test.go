// Copyright 2024, Pulumi Corporation.

package provider

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/pulumi/providertest/pulumitest"
	"github.com/pulumi/providertest/pulumitest/opttest"
	"github.com/pulumi/pulumi/sdk/v3/go/common/apitype"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

func TestSecretsAreEncrypted(t *testing.T) {
	cwd, err := os.Getwd()
	require.NoError(t, err)

	test := pulumitest.NewPulumiTest(t, filepath.Join("test-programs", "explicit-provider-with-config"),
		opttest.LocalProviderPath("azuread", filepath.Join(cwd, "..", "bin")),
		opttest.SkipInstall(),
	)

	res := test.Up(t)

	var d apitype.DeploymentV3
	require.NoError(t, json.Unmarshal(test.ExportStack(t).Deployment, &d))

	require.Len(t, d.Resources, 2) // [stack, provider]
	providerState := d.Resources[1]
	require.Equal(t, providerState.Type, tokens.Type("pulumi:providers:azuread"))

	// Secrets in state
	for _, property := range []string{"clientId", "clientSecret", "clientCertificatePassword"} {
		require.Contains(t, providerState.Inputs, property)
		input, ok := providerState.Inputs[property].(map[string]any)
		// ExportStack() exports with --show-secrets but we can still assert that it is a secret by
		// checking for the object shape and the expected plaintext 'property'.
		require.True(t, ok, "%s input should be a secret", property)
		require.Contains(t, input, "plaintext")
	}

	// Secrets in output
	for _, property := range []string{"clientId", "clientSecret", "clientCertificatePassword"} {
		out := property + "Out"
		require.Contains(t, res.Outputs, out)
		require.True(t, res.Outputs[out].Secret, "%s output should be marked as secret", property)
	}
}

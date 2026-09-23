// Copyright 2024, Pulumi Corporation.

package provider

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/pulumi/providertest/providers"
	"github.com/pulumi/providertest/pulumitest"
	"github.com/pulumi/providertest/pulumitest/opttest"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/common/apitype"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

//nolint:goconst // repeated "clientId" literal keeps the assertions self-contained
func TestSecretsAreEncrypted(t *testing.T) {
	// Configuring an explicit provider resource actually authenticates against Azure AD (the bridge
	// now correctly threads explicit provider config through to the underlying Terraform provider's
	// Configure call), so real credentials are required to create it successfully.
	clientID := os.Getenv("ARM_CLIENT_ID")
	clientSecret := os.Getenv("ARM_CLIENT_SECRET")
	if clientID == "" || clientSecret == "" {
		t.Skip("ARM_CLIENT_ID and ARM_CLIENT_SECRET must be set to exercise a real provider Configure call")
	}

	test := setupTest(t, "explicit-provider-with-config", opttest.SkipInstall())
	test.SetConfig(t, "clientId", clientID)
	require.NoError(t, test.CurrentStack().SetConfig(test.Context(), "clientSecret",
		auto.ConfigValue{Value: clientSecret, Secret: true}))

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

func setupTest(t *testing.T, testProgramDir string, opts ...opttest.Option) *pulumitest.PulumiTest {
	t.Helper()
	rpFactory := providers.ResourceProviderFactory(providerServer)
	dir := filepath.Join("test-programs", testProgramDir)
	opts = append(opts, opttest.AttachProvider("azuread", rpFactory))
	return pulumitest.NewPulumiTest(t, dir, opts...)
}

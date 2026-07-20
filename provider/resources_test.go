// Copyright 2026, Pulumi Corporation.

package provider

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

func TestConfigBoolValue(t *testing.T) {
	tests := []struct {
		name         string
		property     *bool
		environment  string
		defaultValue bool
		want         bool
	}{
		{
			name:         "uses default when unset",
			defaultValue: true,
			want:         true,
		},
		{
			name:         "uses false environment value",
			environment:  "false",
			defaultValue: true,
			want:         false,
		},
		{
			name:         "uses values accepted by strconv ParseBool",
			environment:  "1",
			defaultValue: false,
			want:         true,
		},
		{
			name:         "explicit false overrides environment",
			property:     boolPointer(false),
			environment:  "true",
			defaultValue: true,
			want:         false,
		},
		{
			name:         "explicit true overrides environment",
			property:     boolPointer(true),
			environment:  "false",
			defaultValue: false,
			want:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(armUseCLIEnv, tt.environment)
			vars := resource.PropertyMap{}
			if tt.property != nil {
				vars[useCLIConfigKey] = resource.NewBoolProperty(*tt.property)
			}

			got, err := configBoolValue(vars, useCLIConfigKey, []string{armUseCLIEnv}, tt.defaultValue)
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestConfigBoolValueRejectsInvalidEnvironmentValue(t *testing.T) {
	t.Setenv(armUseCLIEnv, "not-a-boolean")

	_, err := configBoolValue(resource.PropertyMap{}, useCLIConfigKey, []string{armUseCLIEnv}, true)
	require.ErrorContains(t, err, "parsing ARM_USE_CLI as a boolean")
}

func TestGetAuthConfig(t *testing.T) {
	t.Setenv(armUseCLIEnv, "true")
	t.Setenv("ARM_OIDC_TOKEN", "")
	t.Setenv("ARM_OIDC_TOKEN_FILE_PATH", "")

	environment, err := environments.FromName("public")
	require.NoError(t, err)
	tokenPath := writeTokenFile(t, "file-token\n")

	config, err := getAuthConfig(resource.PropertyMap{
		useCLIConfigKey:            resource.NewBoolProperty(false),
		"useOidc":                  resource.NewBoolProperty(true),
		oidcTokenFilePathConfigKey: resource.NewStringProperty(tokenPath),
	}, *environment)

	require.NoError(t, err)
	assert.False(t, config.EnableAuthenticatingUsingAzureCLI)
	assert.True(t, config.EnableAuthenticationUsingOIDC)
	assert.True(t, config.EnableAuthenticationUsingGitHubOIDC)
	assert.Equal(t, "file-token", config.OIDCAssertionToken)
}

func TestGetAuthConfigDefaultsToAzureCLI(t *testing.T) {
	t.Setenv(armUseCLIEnv, "")
	t.Setenv("ARM_OIDC_TOKEN", "")
	t.Setenv("ARM_OIDC_TOKEN_FILE_PATH", "")

	environment, err := environments.FromName("public")
	require.NoError(t, err)
	config, err := getAuthConfig(resource.PropertyMap{}, *environment)

	require.NoError(t, err)
	assert.True(t, config.EnableAuthenticatingUsingAzureCLI)
}

func TestPreConfigureCallbackAcceptsOIDCTokenFileWithCLIDisabled(t *testing.T) {
	t.Setenv("ARM_CLIENT_SECRET", "")
	t.Setenv("ARM_CLIENT_CERTIFICATE_PATH", "")
	t.Setenv("ARM_USE_MSI", "false")
	t.Setenv("ARM_OIDC_TOKEN", "")
	t.Setenv("ARM_OIDC_TOKEN_FILE_PATH", "")

	tokenPath := writeTokenFile(t, "file-token\n")
	err := preConfigureCallback(resource.PropertyMap{
		clientIDConfigKey:          resource.NewStringProperty("00000000-0000-0000-0000-000000000001"),
		"tenantId":                 resource.NewStringProperty("00000000-0000-0000-0000-000000000002"),
		useCLIConfigKey:            resource.NewBoolProperty(false),
		"useOidc":                  resource.NewBoolProperty(true),
		oidcTokenFilePathConfigKey: resource.NewStringProperty(tokenPath),
	}, nil)

	require.NoError(t, err)
}

//nolint:goconst // repeated literals keep each token precedence case self-contained
func TestGetOIDCToken(t *testing.T) {
	t.Setenv("ARM_OIDC_TOKEN", "")
	t.Setenv("ARM_OIDC_TOKEN_FILE_PATH", "")

	t.Run("uses direct token", func(t *testing.T) {
		token, err := getOIDCToken(resource.PropertyMap{
			"oidcToken": resource.NewStringProperty("direct-token"),
		})

		require.NoError(t, err)
		assert.Equal(t, "direct-token", token)
	})

	t.Run("reads and trims file token", func(t *testing.T) {
		path := writeTokenFile(t, "  file-token\n")
		token, err := getOIDCToken(resource.PropertyMap{
			"oidcTokenFilePath": resource.NewStringProperty(path),
		})

		require.NoError(t, err)
		assert.Equal(t, "file-token", token)
	})

	t.Run("accepts matching direct and file tokens", func(t *testing.T) {
		path := writeTokenFile(t, "same-token\n")
		token, err := getOIDCToken(resource.PropertyMap{
			"oidcToken":         resource.NewStringProperty("same-token"),
			"oidcTokenFilePath": resource.NewStringProperty(path),
		})

		require.NoError(t, err)
		assert.Equal(t, "same-token", token)
	})

	t.Run("rejects mismatched direct and file tokens", func(t *testing.T) {
		path := writeTokenFile(t, "file-token")
		_, err := getOIDCToken(resource.PropertyMap{
			"oidcToken":         resource.NewStringProperty("direct-token"),
			"oidcTokenFilePath": resource.NewStringProperty(path),
		})

		require.ErrorContains(t, err, "mismatch between supplied OIDC token")
	})

	t.Run("reports unreadable token file", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "missing-token")
		_, err := getOIDCToken(resource.PropertyMap{
			"oidcTokenFilePath": resource.NewStringProperty(path),
		})

		require.ErrorContains(t, err, "reading OIDC token from file")
	})

	t.Run("uses environment token file", func(t *testing.T) {
		path := writeTokenFile(t, "environment-token")
		t.Setenv("ARM_OIDC_TOKEN_FILE_PATH", path)
		token, err := getOIDCToken(resource.PropertyMap{})

		require.NoError(t, err)
		assert.Equal(t, "environment-token", token)
	})
}

func boolPointer(value bool) *bool {
	return &value
}

func writeTokenFile(t *testing.T, contents string) string {
	t.Helper()

	path := filepath.Join(t.TempDir(), "oidc-token")
	require.NoError(t, os.WriteFile(path, []byte(contents), 0o600))
	return path
}

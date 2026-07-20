// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.
//go:build nodejs || all
// +build nodejs all

package examples

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func getJSBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	environ := getEnviron(t)

	baseJS := base.With(integration.ProgramTestOptions{
		Config: map[string]string{
			"azure:environment": environ,
		},
		Dependencies: []string{
			"@pulumi/azuread",
		},
	})

	return baseJS
}

func oidcToken(t *testing.T) string {
	t.Helper()

	if token := os.Getenv("ARM_OIDC_TOKEN"); token != "" {
		return token
	}

	requestURL := os.Getenv("ACTIONS_ID_TOKEN_REQUEST_URL")
	requestToken := os.Getenv("ACTIONS_ID_TOKEN_REQUEST_TOKEN")
	require.NotEmpty(t, requestURL, "RUN_OIDC_TESTS requires ACTIONS_ID_TOKEN_REQUEST_URL")
	require.NotEmpty(t, requestToken, "RUN_OIDC_TESTS requires ACTIONS_ID_TOKEN_REQUEST_TOKEN")

	parsedURL, err := url.Parse(requestURL)
	require.NoError(t, err)
	query := parsedURL.Query()
	query.Set("audience", "api://AzureADTokenExchange")
	parsedURL.RawQuery = query.Encode()

	request, err := http.NewRequestWithContext(t.Context(), http.MethodGet, parsedURL.String(), nil)
	require.NoError(t, err)
	request.Header.Set("Authorization", "Bearer "+requestToken)

	client := &http.Client{Timeout: 30 * time.Second}
	response, err := client.Do(request)
	require.NoError(t, err)
	defer response.Body.Close()

	require.Equal(t, http.StatusOK, response.StatusCode, "GitHub OIDC token request failed: %s", response.Status)
	var result struct {
		Value string `json:"value"`
	}
	require.NoError(t, json.NewDecoder(response.Body).Decode(&result))
	require.NotEmpty(t, result.Value)
	return result.Value
}

func TestSimple(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:           path.Join(getCwd(t), "simple"),
			RunUpdateTest: true,
			Secrets: map[string]string{
				"password": "SecretP@sswd99!",
			},
		})

	integration.ProgramTest(t, &test)
}

// The same test than the above, but authenticating via OIDC.
func TestSimple_OIDC(t *testing.T) {
	if os.Getenv("RUN_OIDC_TESTS") != "true" {
		t.Skip("Skipping OIDC test without OIDC_ARM_CLIENT_ID")
	}

	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "simple"),
			Env: []string{
				"ARM_USE_OIDC=true",
				"ARM_USE_CLI=false",
				// not strictly necessary but making sure we test the OIDC path
				"ARM_CLIENT_SECRET=",
			},
			RunUpdateTest: true,
			Secrets: map[string]string{
				"password": "SecretP@sswd99!",
			},
		})

	integration.ProgramTest(t, &test)
}

func TestSimple_OIDC_TokenFile(t *testing.T) {
	if os.Getenv("RUN_OIDC_TESTS") != "true" {
		t.Skip("Skipping OIDC test without OIDC_ARM_CLIENT_ID")
	}

	tokenPath := filepath.Join(t.TempDir(), "oidc-token")
	require.NoError(t, os.WriteFile(tokenPath, []byte(oidcToken(t)), 0o600))

	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:                    path.Join(getCwd(t), "simple"),
			NoParallel:             true,
			SkipUpdate:             true,
			SkipRefresh:            true,
			SkipExportImport:       true,
			SkipEmptyPreviewUpdate: true,
			Env: []string{
				"ARM_USE_OIDC=true",
				"ARM_USE_CLI=false",
				"ARM_USE_MSI=false",
				"ARM_USE_AKS_WORKLOAD_IDENTITY=false",
				"ARM_OIDC_TOKEN_FILE_PATH=" + tokenPath,
				"ARM_OIDC_TOKEN=",
				"ARM_OIDC_REQUEST_TOKEN=",
				"ARM_OIDC_REQUEST_URL=",
				"ACTIONS_ID_TOKEN_REQUEST_TOKEN=",
				"ACTIONS_ID_TOKEN_REQUEST_URL=",
				"SYSTEM_ACCESSTOKEN=",
				"SYSTEM_OIDCREQUESTURI=",
				"ARM_ADO_PIPELINE_SERVICE_CONNECTION_ID=",
				"ARM_OIDC_AZURE_SERVICE_CONNECTION_ID=",
				"ARM_CLIENT_CERTIFICATE=",
				"ARM_CLIENT_CERTIFICATE_PATH=",
				"ARM_CLIENT_CERTIFICATE_PASSWORD=",
				"ARM_CLIENT_SECRET_FILE_PATH=",
				"ARM_MSI_ENDPOINT=",
				// not strictly necessary but making sure we test the OIDC path
				"ARM_CLIENT_SECRET=",
			},
			Secrets: map[string]string{
				"password": "SecretP@sswd99!",
			},
		})

	integration.ProgramTest(t, &test)
}

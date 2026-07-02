// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.
//go:build nodejs || all
// +build nodejs all

package examples

import (
	"os"
	"path"
	"testing"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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

func writeOIDCTokenToFile(path, audience string) error {
	reqURL := os.Getenv("ACTIONS_ID_TOKEN_REQUEST_URL")
	reqToken := os.Getenv("ACTIONS_ID_TOKEN_REQUEST_TOKEN")

	if reqURL == "" || reqToken == "" {
		return fmt.Errorf("GitHub OIDC environment variables not set")
	}

	req, err := http.NewRequest(http.MethodGet, reqURL+"&audience="+audience, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+reqToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("OIDC request failed: %s: %s", resp.Status, body)
	}

	var r struct {
		Value string `json:"value"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return err
	}

	return os.WriteFile(path, []byte(r.Value), 0600)
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
				// not strictly necessary but making sure we test the OIDC path
				"ARM_CLIENT_SECRET=",
				"ARM_USE_CLI=false",
			},
			RunUpdateTest: true,
			Secrets: map[string]string{
				"password": "SecretP@sswd99!",
			},
		})

	integration.ProgramTest(t, &test)
}

// Testing oidc login with token file, this is the same as the above test but using a token file instead of env vars
func TestSimple_OIDC_TOKEN_FILE(t *testing.T) {
	if os.Getenv("RUN_OIDC_TESTS") != "true" {
		t.Skip("Skipping OIDC test")
	}

	tokenFile := path.Join(getCwd(t), "simple", "token.jwt")

	err := writeOIDCTokenToFile(tokenFile, "api://AzureADTokenExchange")
	require.NoError(t, err)

	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "simple"),
			Env: []string{
				"ARM_USE_OIDC=true",
				"ARM_OIDC_TOKEN_FILE_PATH=token.jwt",
				// not strictly necessary but making sure we test the OIDC path
				"ARM_CLIENT_SECRET=",
				"ARM_USE_CLI=false",
				// explicatly clear to make sure we use the token from file
				"ACTIONS_ID_TOKEN_REQUEST_URL=",
				"ACTIONS_ID_TOKEN_REQUEST_TOKEN=",
			},
			RunUpdateTest: true,
			Secrets: map[string]string{
				"password": "SecretP@sswd99!",
			},
		})

	integration.ProgramTest(t, &test)
}

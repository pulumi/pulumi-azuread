// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.
//go:build nodejs || all
// +build nodejs all

package examples

import (
	"os"
	"path"
	"testing"

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
	oidcClientId := os.Getenv("OIDC_ARM_CLIENT_ID")
	if oidcClientId == "" {
		t.Skip("Skipping OIDC test without OIDC_ARM_CLIENT_ID")
	}

	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "simple"),
			Env: []string{
				// "ARM_USE_OIDC=true",
				"ARM_CLIENT_ID=" + oidcClientId,
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

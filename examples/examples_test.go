// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package examples

import (
	"os"
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

func TestSimple(t *testing.T) {
	test := getBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:           path.Join(getCwd(t), "simple"),
			RunUpdateTest: true,
			Secrets: map[string]string{
				"password": "SecretP@sswd99!",
			},
		})

	integration.ProgramTest(t, &test)
}

func getEnviron(t *testing.T) string {
	env := os.Getenv("ARM_ENVIRONMENT")
	if env == "" {
		t.Skipf("Skipping test due to missing ARM_ENVIRONMENT environment variable")
	}

	return env
}

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}

func getBaseOptions(t *testing.T) integration.ProgramTestOptions {
	environ := getEnviron(t)
	return integration.ProgramTestOptions{
		Config: map[string]string{
			"azure:environment": environ,
		},
		Dependencies: []string{
			"@pulumi/azuread",
		},
	}
}

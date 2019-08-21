// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package examples

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/pkg/testing/integration"
)

func TestExamples(t *testing.T) {
	// Ensure we have any required configuration points
	environ := os.Getenv("ARM_ENVIRONMENT")
	if environ == "" {
		t.Skipf("Skipping test due to missing ARM_ENVIRONMENT variable")
	}
	cwd, err := os.Getwd()
	if !assert.NoError(t, err, "expected a valid working directory: %v", err) {
		return
	}

	// base options shared amongst all tests.
	base := integration.ProgramTestOptions{
		Config: map[string]string{
			"azure:environment": environ,
		},
		Dependencies: []string{
			"@pulumi/azuread",
		},
	}

	tests := []integration.ProgramTestOptions{
		// List each test
		base.With(integration.ProgramTestOptions{
			Dir:           path.Join(cwd, "simple"),
			RunUpdateTest: true,
			Secrets: map[string]string{
				"password": "SecretP@sswd99!",
			},
		}),
	}

	for _, ex := range tests {
		example := ex
		t.Run(example.Dir, func(t *testing.T) {
			integration.ProgramTest(t, &example)
		})
	}
}

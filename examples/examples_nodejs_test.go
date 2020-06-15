// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package examples

import (
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
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

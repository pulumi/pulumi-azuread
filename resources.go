// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package azuread

import (
	"unicode"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/pulumi/pulumi-terraform/pkg/tfbridge"
	"github.com/pulumi/pulumi/pkg/resource"
	"github.com/pulumi/pulumi/pkg/tokens"
	"github.com/terraform-providers/terraform-provider-azuread/azuread"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "azuread"
	// modules:
	mainMod = "index" // the y module
)

// makeMember manufactures a type token for the package and the given module and type.
func makeMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(mainPkg + ":" + mod + ":" + mem)
}

// makeType manufactures a type token for the package and the given module and type.
func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

// makeDataSource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the data source's
// first character.
func makeDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeMember(mod+"/"+fn, res)
}

// makeResource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the resource's
// first character.
func makeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+fn, res)
}

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be.
func preConfigureCallback(vars resource.PropertyMap, c *terraform.ResourceConfig) error {
	return nil
}

// managedByPulumi is a default used for some managed resources, in the absence of something more meaningful.
// var managedByPulumi = &tfbridge.DefaultInfo{Value: "Managed by Pulumi"}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := azuread.Provider().(*schema.Provider)

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "azuread",
		Description: "A Pulumi package for creating and managing azuread cloud resources.",
		Keywords:    []string{"pulumi", "azuread"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/pulumi/pulumi-azuread",
		Config: map[string]*tfbridge.SchemaInfo{
			"client_id": {
				Default: &tfbridge.DefaultInfo{
					Value:   "",
					EnvVars: []string{"ARM_CLIENT_ID"},
				},
			},
			"environment": {
				Default: &tfbridge.DefaultInfo{
					Value:   "public",
					EnvVars: []string{"ARM_ENVIRONMENT"},
				},
			},
			"subscription_id": {
				Default: &tfbridge.DefaultInfo{
					Value:   "",
					EnvVars: []string{"ARM_SUBSCRIPTION_ID"},
				},
			},
			"tenant_id": {
				Default: &tfbridge.DefaultInfo{
					Value:   "",
					EnvVars: []string{"ARM_TENANT_ID"},
				},
			},
			"client_certificate_password": {
				Default: &tfbridge.DefaultInfo{
					Value:   "",
					EnvVars: []string{"ARM_CLIENT_CERTIFICATE_PASSWORD"},
				},
			},
			"client_certificate_path": {
				Default: &tfbridge.DefaultInfo{
					Value:   "",
					EnvVars: []string{"ARM_CLIENT_CERTIFICATE_PATH"},
				},
			},
			"client_secret": {
				Default: &tfbridge.DefaultInfo{
					Value:   "",
					EnvVars: []string{"ARM_CLIENT_SECRET"},
				},
			},
			"msi_endpoint": {
				Default: &tfbridge.DefaultInfo{
					Value:   "",
					EnvVars: []string{"ARM_MSI_ENDPOINT"},
				},
			},
			"use_msi": {
				Default: &tfbridge.DefaultInfo{
					Value:   false,
					EnvVars: []string{"ARM_USE_MSI"},
				},
			},
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"azuread_application":                {Tok: makeResource(mainMod, "Application")},
			"azuread_group":                      {Tok: makeResource(mainMod, "Group")},
			"azuread_service_principal":          {Tok: makeResource(mainMod, "ServicePrincipal")},
			"azuread_service_principal_password": {Tok: makeResource(mainMod, "ServicePrincipalPassword")},
			"azuread_user":                       {Tok: makeResource(mainMod, "User")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"azuread_application":       {Tok: makeDataSource(mainMod, "getApplication")},
			"azuread_domains":           {Tok: makeDataSource(mainMod, "getDomains")},
			"azuread_group":             {Tok: makeDataSource(mainMod, "getGroup")},
			"azuread_service_principal": {Tok: makeDataSource(mainMod, "getServicePrincipal")},
			"azuread_user":              {Tok: makeDataSource(mainMod, "getUser")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "latest",
			},
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=0.17.1,<0.18.0",
			},
		},
	}

	// For all resources with name properties, we will add an auto-name property.  Make sure to skip those that
	// already have a name mapping entry, since those may have custom overrides set above (e.g., for length).
	const nameProperty = "name"
	for resname, res := range prov.Resources {
		if schema := p.ResourcesMap[resname]; schema != nil {
			// Only apply auto-name to input properties (Optional || Required) named `name`
			if tfs, has := schema.Schema[nameProperty]; has && (tfs.Optional || tfs.Required) {
				if _, hasfield := res.Fields[nameProperty]; !hasfield {
					if res.Fields == nil {
						res.Fields = make(map[string]*tfbridge.SchemaInfo)
					}
					res.Fields[nameProperty] = tfbridge.AutoName(nameProperty, 255)
				}
			}
		}
	}

	return prov
}

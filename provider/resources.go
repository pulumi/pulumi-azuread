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

package provider

import (
	"unicode"

	"github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfbridge"
	shimv1 "github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfshim/sdk-v1"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
	"github.com/terraform-providers/terraform-provider-azuread/shim"
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

// managedByPulumi is a default used for some managed resources, in the absence of something more meaningful.
// var managedByPulumi = &tfbridge.DefaultInfo{Value: "Managed by Pulumi"}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv1.NewProvider(shim.NewProvider())

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
		Resources: map[string]*tfbridge.ResourceInfo{
			"azuread_application":                   {Tok: makeResource(mainMod, "Application")},
			"azuread_application_password":          {Tok: makeResource(mainMod, "ApplicationPassword")},
			"azuread_group":                         {Tok: makeResource(mainMod, "Group")},
			"azuread_service_principal":             {Tok: makeResource(mainMod, "ServicePrincipal")},
			"azuread_service_principal_password":    {Tok: makeResource(mainMod, "ServicePrincipalPassword")},
			"azuread_user":                          {Tok: makeResource(mainMod, "User")},
			"azuread_group_member":                  {Tok: makeResource(mainMod, "GroupMember")},
			"azuread_application_certificate":       {Tok: makeResource(mainMod, "ApplicationCertificate")},
			"azuread_application_app_role":          {Tok: makeResource(mainMod, "ApplicationAppRole")},
			"azuread_application_oauth2_permission": {Tok: makeResource(mainMod, "ApplicationOAuth2Permission")},
			"azuread_service_principal_certificate": {Tok: makeResource(mainMod, "ServicePrincipalCertificate")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"azuread_application":       {Tok: makeDataSource(mainMod, "getApplication")},
			"azuread_domains":           {Tok: makeDataSource(mainMod, "getDomains")},
			"azuread_group":             {Tok: makeDataSource(mainMod, "getGroup")},
			"azuread_service_principal": {Tok: makeDataSource(mainMod, "getServicePrincipal")},
			"azuread_user":              {Tok: makeDataSource(mainMod, "getUser")},
			"azuread_groups":            {Tok: makeDataSource(mainMod, "getGroups")},
			"azuread_users":             {Tok: makeDataSource(mainMod, "getUsers")},
			"azuread_client_config":     {Tok: makeDataSource(mainMod, "getClientConfig")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^2.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=2.9.0,<3.0.0",
			},
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi":                       "2.*",
				"System.Collections.Immutable": "1.6.0",
			},
			Namespaces: map[string]string{
				"azuread": "AzureAD",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}

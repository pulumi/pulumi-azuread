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
	"context"
	"fmt"
	"os"
	"path/filepath"
	"unicode"

	"github.com/manicminer/hamilton/auth"
	"github.com/manicminer/hamilton/environments"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"

	"github.com/hashicorp/terraform-provider-azuread/shim"
	"github.com/pulumi/pulumi-azuread/provider/v5/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	tfshim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
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

// stringValue gets a string value from a property map, then from environment vars;
// if neither are present, returns empty string ""
func stringValue(vars resource.PropertyMap, prop resource.PropertyKey, envs []string) string {
	val, ok := vars[prop]
	if ok && val.IsString() {
		return val.StringValue()
	}
	for _, env := range envs {
		val, ok := os.LookupEnv(env)
		if ok {
			return val
		}
	}
	return ""
}

// boolValue takes a bool value from a property map, then from environment vars; defaults to false
func boolValue(vars resource.PropertyMap, prop resource.PropertyKey, envs []string) bool {
	val, ok := vars[prop]
	if ok && val.IsBool() {
		return val.BoolValue()
	}
	for _, env := range envs {
		val, ok := os.LookupEnv(env)
		if ok && val == "true" {
			return true
		}
	}
	return false
}

// preConfigureCallback returns an error when cloud provider setup is misconfigured
//
// nolint: lll
func preConfigureCallback(vars resource.PropertyMap, c tfshim.ResourceConfig) error {

	envName := stringValue(vars, "environment", []string{"ARM_ENVIRONMENT"})
	if envName == "" {
		envName = "public"
	}

	env, err := environments.EnvironmentFromString(envName)
	if err != nil {
		return fmt.Errorf("failed to read Azure environment \"%s\": %v", envName, err)
	}

	authConfig := &auth.Config{
		Environment:            env,
		EnableClientSecretAuth: true,
		EnableAzureCliToken:    true,
		TenantID:               stringValue(vars, "tenantId", []string{"ARM_TENANT_ID"}),
		ClientID:               stringValue(vars, "clientId", []string{"ARM_CLIENT_ID"}),
		ClientSecret:           stringValue(vars, "clientSecret", []string{"ARM_CLIENT_SECRET"}),

		EnableClientCertAuth: true,
		// We don't handle ClientCertData yet, which is the actual base-64 encoded cert in config
		ClientCertPassword: stringValue(vars, "clientCertificatePassword", []string{"ARM_CLIENT_CERTIFICATE_PASSWORD"}),
		ClientCertPath:     stringValue(vars, "clientCertificatePath", []string{"ARM_CLIENT_CERTIFICATE_PATH"}),

		EnableMsiAuth: boolValue(vars, "msiEndpoint", []string{"ARM_USE_MSI"}),
		MsiEndpoint:   stringValue(vars, "msiEndpoint", []string{"ARM_MSI_ENDPOINT"}),

		// The configuration below would enable OIDC auth which we haven't tested and documented yet.
		//FederatedAssertion:        idToken,
		//IDTokenRequestURL:         d.Get("oidc_request_url").(string),
		//IDTokenRequestToken:       d.Get("oidc_request_token").(string),
		//EnableClientFederatedAuth: d.Get("use_oidc").(bool),
		//EnableGitHubOIDCAuth:      d.Get("use_oidc").(bool),
	}

	_, err = authConfig.NewAuthorizer(context.Background(), env.MsGraph)
	if err != nil {
		return fmt.Errorf("failed to load Azure credentials.\n"+
			"Details: %v\n\n"+
			"\tPlease make sure you have signed in via 'az login' or configured another authentication method.\n\n"+
			"\tSee https://www.pulumi.com/registry/packages/azuread/installation-configuration/ for more information.", err)
	}

	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(shim.NewProvider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "azuread",
		DisplayName: "Azure Active Directory (Azure AD)",
		Description: "A Pulumi package for creating and managing Azure Active Directory (Azure AD) cloud resources.",
		Keywords:    []string{"pulumi", "azuread"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		GitHubOrg:   "hashicorp",
		Repository:  "https://github.com/pulumi/pulumi-azuread",
		Config: map[string]*tfbridge.SchemaInfo{
			"environment": {
				Default: &tfbridge.DefaultInfo{
					Value:   "public",
					EnvVars: []string{"ARM_ENVIRONMENT"},
				},
			},
			"msi_endpoint": {
				Default: &tfbridge.DefaultInfo{
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
			"azuread_application_password":       {Tok: makeResource(mainMod, "ApplicationPassword")},
			"azuread_group":                      {Tok: makeResource(mainMod, "Group")},
			"azuread_service_principal":          {Tok: makeResource(mainMod, "ServicePrincipal")},
			"azuread_service_principal_password": {Tok: makeResource(mainMod, "ServicePrincipalPassword")},
			"azuread_service_principal_delegated_permission_grant": {
				Tok: makeResource(mainMod, "ServicePrincipalDelegatedPermissionGrant"),
			},
			"azuread_user":                          {Tok: makeResource(mainMod, "User")},
			"azuread_group_member":                  {Tok: makeResource(mainMod, "GroupMember")},
			"azuread_application_certificate":       {Tok: makeResource(mainMod, "ApplicationCertificate")},
			"azuread_service_principal_certificate": {Tok: makeResource(mainMod, "ServicePrincipalCertificate")},
			"azuread_application_pre_authorized":    {Tok: makeResource(mainMod, "ApplicationPreAuthorized")},
			"azuread_invitation":                    {Tok: makeResource(mainMod, "Invitation")},
			"azuread_conditional_access_policy":     {Tok: makeResource(mainMod, "ConditionalAccessPolicy")},
			"azuread_named_location":                {Tok: makeResource(mainMod, "NamedLocation")},
			"azuread_directory_role":                {Tok: makeResource(mainMod, "DirectoryRole")},
			"azuread_directory_role_member":         {Tok: makeResource(mainMod, "DirectoryRoleMember")},
			"azuread_app_role_assignment":           {Tok: makeResource(mainMod, "AppRoleAssignment")},
			"azuread_administrative_unit":           {Tok: makeResource(mainMod, "AdministrativeUnit")},
			"azuread_administrative_unit_member":    {Tok: makeResource(mainMod, "AdministrativeUnitMember")},
			"azuread_application_federated_identity_credential": {
				Tok: makeResource(mainMod, "ApplicationFederatedIdentityCredential"),
			},
			"azuread_custom_directory_role":     {Tok: makeResource(mainMod, "CustomDirectoryRole")},
			"azuread_claims_mapping_policy":     {Tok: makeResource(mainMod, "ClaimsMappingPolicy")},
			"azuread_directory_role_assignment": {Tok: makeResource(mainMod, "DirectoryRoleAssignment")},
			"azuread_service_principal_claims_mapping_policy_assignment": {
				Tok: makeResource(mainMod, "ServicePrincipalClaimsMappingPolicyAssignment"),
			},
			"azuread_synchronization_job":    {Tok: makeResource(mainMod, "SynchronizationJob")},
			"azuread_synchronization_secret": {Tok: makeResource(mainMod, "SynchronizationSecret")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"azuread_application":                   {Tok: makeDataSource(mainMod, "getApplication")},
			"azuread_domains":                       {Tok: makeDataSource(mainMod, "getDomains")},
			"azuread_group":                         {Tok: makeDataSource(mainMod, "getGroup")},
			"azuread_service_principal":             {Tok: makeDataSource(mainMod, "getServicePrincipal")},
			"azuread_user":                          {Tok: makeDataSource(mainMod, "getUser")},
			"azuread_groups":                        {Tok: makeDataSource(mainMod, "getGroups")},
			"azuread_users":                         {Tok: makeDataSource(mainMod, "getUsers")},
			"azuread_client_config":                 {Tok: makeDataSource(mainMod, "getClientConfig")},
			"azuread_application_published_app_ids": {Tok: makeDataSource(mainMod, "getApplicationPublishedAppIds")},
			"azuread_application_template":          {Tok: makeDataSource(mainMod, "getApplicationTemplate")},
			"azuread_service_principals":            {Tok: makeDataSource(mainMod, "getServicePrincipals")},
			"azuread_administrative_unit":           {Tok: makeDataSource(mainMod, "getAdministrativeUnit")},
			"azuread_directory_object":              {Tok: makeDataSource(mainMod, "getDirectoryObject")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			Namespaces: map[string]string{
				"azuread": "AzureAD",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}

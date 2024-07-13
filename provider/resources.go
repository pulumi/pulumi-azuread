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
	"path/filepath"
	"unicode"

	// embed is used to store bridge-metadata.json in the compiled binary
	_ "embed"

	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/hashicorp/terraform-provider-azuread/shim"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	tfbridgetokens "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	tfshim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/pulumi/pulumi-azuread/provider/v5/pkg/version"
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

// preConfigureCallback returns an error when cloud provider setup is misconfigured
//
// nolint: lll
func preConfigureCallback(vars resource.PropertyMap, _ tfshim.ResourceConfig) error {
	envName := tfbridge.ConfigStringValue(vars, "environment", []string{"ARM_ENVIRONMENT"})
	if envName == "" {
		envName = "public"
	}

	env, err := environments.FromName(envName)
	if err != nil {
		return fmt.Errorf("failed to read Azure environment \"%s\": %v", envName, err)
	}

	useOIDC := tfbridge.ConfigBoolValue(vars, "useOidc", []string{"ARM_USE_OIDC"})
	authConfig := auth.Credentials{
		Environment:        *env,
		ClientID:           tfbridge.ConfigStringValue(vars, "clientId", []string{"ARM_CLIENT_ID"}),
		ClientSecret:       tfbridge.ConfigStringValue(vars, "clientSecret", []string{"ARM_CLIENT_SECRET"}),
		TenantID:           tfbridge.ConfigStringValue(vars, "tenantId", []string{"ARM_TENANT_ID"}),
		AuxiliaryTenantIDs: tfbridge.ConfigArrayValue(vars, "auxiliaryTenantIDs", []string{"ARM_AUXILIARY_TENANT_IDS"}),

		// We don't handle ClientCertData yet, which is the actual base-64 encoded cert in config
		ClientCertificatePath:     tfbridge.ConfigStringValue(vars, "clientCertificatePath", []string{"ARM_CLIENT_CERTIFICATE_PATH"}),
		ClientCertificatePassword: tfbridge.ConfigStringValue(vars, "clientCertificatePassword", []string{"ARM_CLIENT_CERTIFICATE_PASSWORD"}),

		CustomManagedIdentityEndpoint: tfbridge.ConfigStringValue(vars, "msiEndpoint", []string{"ARM_MSI_ENDPOINT"}),

		// OIDC section. The ACTIONS_ variables are set by GitHub.
		GitHubOIDCTokenRequestToken: tfbridge.ConfigStringValue(vars, "oidcRequestToken", []string{"ARM_OIDC_REQUEST_TOKEN", "ACTIONS_ID_TOKEN_REQUEST_TOKEN"}),
		GitHubOIDCTokenRequestURL:   tfbridge.ConfigStringValue(vars, "oidcRequestUrl", []string{"ARM_OIDC_REQUEST_URL", "ACTIONS_ID_TOKEN_REQUEST_URL"}),
		OIDCAssertionToken:          tfbridge.ConfigStringValue(vars, "oidcToken", []string{"ARM_OIDC_TOKEN"}),

		// Feature Toggles
		EnableAuthenticatingUsingClientCertificate: true,
		EnableAuthenticatingUsingClientSecret:      true,
		EnableAuthenticatingUsingManagedIdentity:   tfbridge.ConfigBoolValue(vars, "useMsi", []string{"ARM_USE_MSI"}),
		EnableAuthenticatingUsingAzureCLI:          true,
		EnableAuthenticationUsingOIDC:              useOIDC,
		EnableAuthenticationUsingGitHubOIDC:        useOIDC,
	}

	_, err = auth.NewAuthorizerFromCredentials(context.Background(), authConfig, env.MicrosoftGraph)
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
			"client_certificate_password": {
				Secret: pulumi.BoolRef(true),
			},
			"client_id": {
				Secret: pulumi.BoolRef(true),
			},
			"client_secret": {
				Secret: pulumi.BoolRef(true),
			},
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
			"azuread_service_principal_token_signing_certificate": {
				Tok: makeResource(mainMod, "ServicePrincipalTokenSigningCertificate"),
			},
			"azuread_application_pre_authorized": {Tok: makeResource(mainMod, "ApplicationPreAuthorized")},
			"azuread_invitation":                 {Tok: makeResource(mainMod, "Invitation")},
			"azuread_conditional_access_policy":  {Tok: makeResource(mainMod, "ConditionalAccessPolicy")},
			"azuread_named_location":             {Tok: makeResource(mainMod, "NamedLocation")},
			"azuread_directory_role":             {Tok: makeResource(mainMod, "DirectoryRole")},
			"azuread_directory_role_member":      {Tok: makeResource(mainMod, "DirectoryRoleMember")},
			"azuread_app_role_assignment":        {Tok: makeResource(mainMod, "AppRoleAssignment")},
			"azuread_administrative_unit":        {Tok: makeResource(mainMod, "AdministrativeUnit")},
			"azuread_administrative_unit_member": {Tok: makeResource(mainMod, "AdministrativeUnitMember")},
			"azuread_application_federated_identity_credential": {
				Tok: makeResource(mainMod, "ApplicationFederatedIdentityCredential"),
			},
			"azuread_custom_directory_role":     {Tok: makeResource(mainMod, "CustomDirectoryRole")},
			"azuread_claims_mapping_policy":     {Tok: makeResource(mainMod, "ClaimsMappingPolicy")},
			"azuread_directory_role_assignment": {Tok: makeResource(mainMod, "DirectoryRoleAssignment")},
			"azuread_service_principal_claims_mapping_policy_assignment": {
				Tok: makeResource(mainMod, "ServicePrincipalClaimsMappingPolicyAssignment"),
			},
			"azuread_synchronization_job":              {Tok: makeResource(mainMod, "SynchronizationJob")},
			"azuread_synchronization_secret":           {Tok: makeResource(mainMod, "SynchronizationSecret")},
			"azuread_access_package":                   {Tok: makeResource(mainMod, "AccessPackage")},
			"azuread_access_package_assignment_policy": {Tok: makeResource(mainMod, "AccessPackageAssignmentPolicy")},
			"azuread_access_package_catalog":           {Tok: makeResource(mainMod, "AccessPackageCatalog")},
			"azuread_access_package_resource_catalog_association": {
				Tok: makeResource(mainMod, "AccessPackageResourceCatalogAssociation"),
			},
			"azuread_access_package_resource_package_association": {
				Tok: makeResource(mainMod, "AccessPackageResourcePackageAssociation"),
			},
			"azuread_administrative_unit_role_member": {Tok: makeResource(mainMod, "AdministrativeUnitRoleMember")},
			"azuread_user_flow_attribute":             {Tok: makeResource(mainMod, "UserFlowAttribute")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"azuread_client_config":                 {Tok: makeDataSource(mainMod, "getClientConfig")},
			"azuread_application_published_app_ids": {Tok: makeDataSource(mainMod, "getApplicationPublishedAppIds")},
			"azuread_application_template":          {Tok: makeDataSource(mainMod, "getApplicationTemplate")},
			"azuread_service_principals":            {Tok: makeDataSource(mainMod, "getServicePrincipals")},
			"azuread_administrative_unit":           {Tok: makeDataSource(mainMod, "getAdministrativeUnit")},
			"azuread_directory_object":              {Tok: makeDataSource(mainMod, "getDirectoryObject")},
			"azuread_directory_roles":               {Tok: makeDataSource(mainMod, "getDirectoryRoles")},
			"azuread_access_package":                {Tok: makeDataSource(mainMod, "getAccessPackage")},
			"azuread_access_package_catalog":        {Tok: makeDataSource(mainMod, "getAccessPackageCatalog")},
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
			RespectSchemaVersion: true,
		},
		Python: (func() *tfbridge.PythonInfo {
			i := &tfbridge.PythonInfo{
				RespectSchemaVersion: true,
				Requires: map[string]string{
					"pulumi": ">=3.0.0,<4.0.0",
				},
			}
			i.PyProject.Enabled = true
			i.InputTypes = tfbridge.PythonInputTypeClassesAndDicts
			return i
		})(),

		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
			RespectSchemaVersion:           true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RespectSchemaVersion: true,
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			Namespaces: map[string]string{
				"azuread": "AzureAD",
			},
		}, MetadataInfo: tfbridge.NewProviderMetadata(metadata),
	}

	prov.MustComputeTokens(tfbridgetokens.SingleModule("azuread_", mainMod,
		tfbridgetokens.MakeStandard(mainPkg)))
	prov.SetAutonaming(255, "-")
	prov.MustApplyAutoAliases()

	return prov
}

//go:embed cmd/pulumi-resource-azuread/bridge-metadata.json
var metadata []byte

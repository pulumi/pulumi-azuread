// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuread/sdk/v5/go/azuread/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets information about an existing service principal associated with an application within Azure Active Directory.
//
// ## API Permissions
//
// The following API permissions are required in order to use this data source.
//
// When authenticated with a service principal, this data source requires one of the following application roles: `Application.Read.All` or `Directory.Read.All`
//
// When authenticated with a user principal, this data source does not require any additional roles.
//
// ## Example Usage
//
// *Look up by application display name*
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-azuread/sdk/v5/go/azuread"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := azuread.LookupServicePrincipal(ctx, &azuread.LookupServicePrincipalArgs{
//				DisplayName: pulumi.StringRef("my-awesome-application"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// *Look up by client ID*
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-azuread/sdk/v5/go/azuread"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := azuread.LookupServicePrincipal(ctx, &azuread.LookupServicePrincipalArgs{
//				ClientId: pulumi.StringRef("00000000-0000-0000-0000-000000000000"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// *Look up by service principal object ID*
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-azuread/sdk/v5/go/azuread"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := azuread.LookupServicePrincipal(ctx, &azuread.LookupServicePrincipalArgs{
//				ObjectId: pulumi.StringRef("00000000-0000-0000-0000-000000000000"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupServicePrincipal(ctx *pulumi.Context, args *LookupServicePrincipalArgs, opts ...pulumi.InvokeOption) (*LookupServicePrincipalResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupServicePrincipalResult
	err := ctx.Invoke("azuread:index/getServicePrincipal:getServicePrincipal", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServicePrincipal.
type LookupServicePrincipalArgs struct {
	// Deprecated: The `application_id` property has been replaced with the `client_id` property and will be removed in version 3.0 of the AzureAD provider
	ApplicationId *string `pulumi:"applicationId"`
	// The client ID of the application associated with this service principal.
	ClientId *string `pulumi:"clientId"`
	// The display name of the application associated with this service principal.
	DisplayName *string `pulumi:"displayName"`
	// The object ID of the service principal.
	//
	// > One of `clientId`, `displayName` or `objectId` must be specified.
	ObjectId *string `pulumi:"objectId"`
}

// A collection of values returned by getServicePrincipal.
type LookupServicePrincipalResult struct {
	// Whether the service principal account is enabled.
	AccountEnabled bool `pulumi:"accountEnabled"`
	// A list of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities.
	AlternativeNames []string `pulumi:"alternativeNames"`
	// Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application.
	AppRoleAssignmentRequired bool `pulumi:"appRoleAssignmentRequired"`
	// A mapping of app role values to app role IDs, as published by the associated application, intended to be useful when referencing app roles in other resources in your configuration.
	AppRoleIds map[string]string `pulumi:"appRoleIds"`
	// A list of app roles published by the associated application, as documented below. For more information [official documentation](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).
	AppRoles []GetServicePrincipalAppRole `pulumi:"appRoles"`
	// Deprecated: The `application_id` property has been replaced with the `client_id` property and will be removed in version 3.0 of the AzureAD provider
	ApplicationId string `pulumi:"applicationId"`
	// The tenant ID where the associated application is registered.
	ApplicationTenantId string `pulumi:"applicationTenantId"`
	// The client ID of the application associated with this service principal.
	ClientId string `pulumi:"clientId"`
	// Permission help text that appears in the admin app assignment and consent experiences.
	Description string `pulumi:"description"`
	// Display name for the permission that appears in the admin consent and app assignment experiences.
	DisplayName string                          `pulumi:"displayName"`
	FeatureTags []GetServicePrincipalFeatureTag `pulumi:"featureTags"`
	// A `features` block as described below.
	//
	// Deprecated: This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider
	Features []GetServicePrincipalFeature `pulumi:"features"`
	// Home page or landing page of the associated application.
	HomepageUrl string `pulumi:"homepageUrl"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps.
	LoginUrl string `pulumi:"loginUrl"`
	// The URL that will be used by Microsoft's authorization service to logout an user using OpenId Connect front-channel, back-channel or SAML logout protocols, taken from the associated application.
	LogoutUrl string `pulumi:"logoutUrl"`
	// A free text field to capture information about the service principal, typically used for operational purposes.
	Notes string `pulumi:"notes"`
	// A list of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
	NotificationEmailAddresses []string `pulumi:"notificationEmailAddresses"`
	// A mapping of OAuth2.0 permission scope values to scope IDs, as exposed by the associated application, intended to be useful when referencing permission scopes in other resources in your configuration.
	Oauth2PermissionScopeIds map[string]string `pulumi:"oauth2PermissionScopeIds"`
	// A collection of OAuth 2.0 delegated permissions exposed by the associated application. Each permission is covered by an `oauth2PermissionScopes` block as documented below.
	Oauth2PermissionScopes []GetServicePrincipalOauth2PermissionScope `pulumi:"oauth2PermissionScopes"`
	// The object ID of the service principal.
	ObjectId string `pulumi:"objectId"`
	// The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps.
	PreferredSingleSignOnMode string `pulumi:"preferredSingleSignOnMode"`
	// A list of URLs where user tokens are sent for sign-in with the associated application, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent for the associated application.
	RedirectUris []string `pulumi:"redirectUris"`
	// The URL where the service exposes SAML metadata for federation.
	SamlMetadataUrl string `pulumi:"samlMetadataUrl"`
	// A `samlSingleSignOn` block as documented below.
	SamlSingleSignOns []GetServicePrincipalSamlSingleSignOn `pulumi:"samlSingleSignOns"`
	// A list of identifier URI(s), copied over from the associated application.
	ServicePrincipalNames []string `pulumi:"servicePrincipalNames"`
	// The Microsoft account types that are supported for the associated application. Possible values include `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`.
	SignInAudience string `pulumi:"signInAudience"`
	// A list of tags applied to the service principal.
	Tags []string `pulumi:"tags"`
	// Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Possible values are `User` or `Admin`.
	Type string `pulumi:"type"`
}

func LookupServicePrincipalOutput(ctx *pulumi.Context, args LookupServicePrincipalOutputArgs, opts ...pulumi.InvokeOption) LookupServicePrincipalResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServicePrincipalResult, error) {
			args := v.(LookupServicePrincipalArgs)
			r, err := LookupServicePrincipal(ctx, &args, opts...)
			var s LookupServicePrincipalResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServicePrincipalResultOutput)
}

// A collection of arguments for invoking getServicePrincipal.
type LookupServicePrincipalOutputArgs struct {
	// Deprecated: The `application_id` property has been replaced with the `client_id` property and will be removed in version 3.0 of the AzureAD provider
	ApplicationId pulumi.StringPtrInput `pulumi:"applicationId"`
	// The client ID of the application associated with this service principal.
	ClientId pulumi.StringPtrInput `pulumi:"clientId"`
	// The display name of the application associated with this service principal.
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
	// The object ID of the service principal.
	//
	// > One of `clientId`, `displayName` or `objectId` must be specified.
	ObjectId pulumi.StringPtrInput `pulumi:"objectId"`
}

func (LookupServicePrincipalOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServicePrincipalArgs)(nil)).Elem()
}

// A collection of values returned by getServicePrincipal.
type LookupServicePrincipalResultOutput struct{ *pulumi.OutputState }

func (LookupServicePrincipalResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServicePrincipalResult)(nil)).Elem()
}

func (o LookupServicePrincipalResultOutput) ToLookupServicePrincipalResultOutput() LookupServicePrincipalResultOutput {
	return o
}

func (o LookupServicePrincipalResultOutput) ToLookupServicePrincipalResultOutputWithContext(ctx context.Context) LookupServicePrincipalResultOutput {
	return o
}

func (o LookupServicePrincipalResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupServicePrincipalResult] {
	return pulumix.Output[LookupServicePrincipalResult]{
		OutputState: o.OutputState,
	}
}

// Whether the service principal account is enabled.
func (o LookupServicePrincipalResultOutput) AccountEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) bool { return v.AccountEnabled }).(pulumi.BoolOutput)
}

// A list of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities.
func (o LookupServicePrincipalResultOutput) AlternativeNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) []string { return v.AlternativeNames }).(pulumi.StringArrayOutput)
}

// Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application.
func (o LookupServicePrincipalResultOutput) AppRoleAssignmentRequired() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) bool { return v.AppRoleAssignmentRequired }).(pulumi.BoolOutput)
}

// A mapping of app role values to app role IDs, as published by the associated application, intended to be useful when referencing app roles in other resources in your configuration.
func (o LookupServicePrincipalResultOutput) AppRoleIds() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) map[string]string { return v.AppRoleIds }).(pulumi.StringMapOutput)
}

// A list of app roles published by the associated application, as documented below. For more information [official documentation](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).
func (o LookupServicePrincipalResultOutput) AppRoles() GetServicePrincipalAppRoleArrayOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) []GetServicePrincipalAppRole { return v.AppRoles }).(GetServicePrincipalAppRoleArrayOutput)
}

// Deprecated: The `application_id` property has been replaced with the `client_id` property and will be removed in version 3.0 of the AzureAD provider
func (o LookupServicePrincipalResultOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) string { return v.ApplicationId }).(pulumi.StringOutput)
}

// The tenant ID where the associated application is registered.
func (o LookupServicePrincipalResultOutput) ApplicationTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) string { return v.ApplicationTenantId }).(pulumi.StringOutput)
}

// The client ID of the application associated with this service principal.
func (o LookupServicePrincipalResultOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) string { return v.ClientId }).(pulumi.StringOutput)
}

// Permission help text that appears in the admin app assignment and consent experiences.
func (o LookupServicePrincipalResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) string { return v.Description }).(pulumi.StringOutput)
}

// Display name for the permission that appears in the admin consent and app assignment experiences.
func (o LookupServicePrincipalResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupServicePrincipalResultOutput) FeatureTags() GetServicePrincipalFeatureTagArrayOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) []GetServicePrincipalFeatureTag { return v.FeatureTags }).(GetServicePrincipalFeatureTagArrayOutput)
}

// A `features` block as described below.
//
// Deprecated: This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider
func (o LookupServicePrincipalResultOutput) Features() GetServicePrincipalFeatureArrayOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) []GetServicePrincipalFeature { return v.Features }).(GetServicePrincipalFeatureArrayOutput)
}

// Home page or landing page of the associated application.
func (o LookupServicePrincipalResultOutput) HomepageUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) string { return v.HomepageUrl }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupServicePrincipalResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) string { return v.Id }).(pulumi.StringOutput)
}

// The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps.
func (o LookupServicePrincipalResultOutput) LoginUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) string { return v.LoginUrl }).(pulumi.StringOutput)
}

// The URL that will be used by Microsoft's authorization service to logout an user using OpenId Connect front-channel, back-channel or SAML logout protocols, taken from the associated application.
func (o LookupServicePrincipalResultOutput) LogoutUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) string { return v.LogoutUrl }).(pulumi.StringOutput)
}

// A free text field to capture information about the service principal, typically used for operational purposes.
func (o LookupServicePrincipalResultOutput) Notes() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) string { return v.Notes }).(pulumi.StringOutput)
}

// A list of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
func (o LookupServicePrincipalResultOutput) NotificationEmailAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) []string { return v.NotificationEmailAddresses }).(pulumi.StringArrayOutput)
}

// A mapping of OAuth2.0 permission scope values to scope IDs, as exposed by the associated application, intended to be useful when referencing permission scopes in other resources in your configuration.
func (o LookupServicePrincipalResultOutput) Oauth2PermissionScopeIds() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) map[string]string { return v.Oauth2PermissionScopeIds }).(pulumi.StringMapOutput)
}

// A collection of OAuth 2.0 delegated permissions exposed by the associated application. Each permission is covered by an `oauth2PermissionScopes` block as documented below.
func (o LookupServicePrincipalResultOutput) Oauth2PermissionScopes() GetServicePrincipalOauth2PermissionScopeArrayOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) []GetServicePrincipalOauth2PermissionScope {
		return v.Oauth2PermissionScopes
	}).(GetServicePrincipalOauth2PermissionScopeArrayOutput)
}

// The object ID of the service principal.
func (o LookupServicePrincipalResultOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) string { return v.ObjectId }).(pulumi.StringOutput)
}

// The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps.
func (o LookupServicePrincipalResultOutput) PreferredSingleSignOnMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) string { return v.PreferredSingleSignOnMode }).(pulumi.StringOutput)
}

// A list of URLs where user tokens are sent for sign-in with the associated application, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent for the associated application.
func (o LookupServicePrincipalResultOutput) RedirectUris() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) []string { return v.RedirectUris }).(pulumi.StringArrayOutput)
}

// The URL where the service exposes SAML metadata for federation.
func (o LookupServicePrincipalResultOutput) SamlMetadataUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) string { return v.SamlMetadataUrl }).(pulumi.StringOutput)
}

// A `samlSingleSignOn` block as documented below.
func (o LookupServicePrincipalResultOutput) SamlSingleSignOns() GetServicePrincipalSamlSingleSignOnArrayOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) []GetServicePrincipalSamlSingleSignOn { return v.SamlSingleSignOns }).(GetServicePrincipalSamlSingleSignOnArrayOutput)
}

// A list of identifier URI(s), copied over from the associated application.
func (o LookupServicePrincipalResultOutput) ServicePrincipalNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) []string { return v.ServicePrincipalNames }).(pulumi.StringArrayOutput)
}

// The Microsoft account types that are supported for the associated application. Possible values include `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`.
func (o LookupServicePrincipalResultOutput) SignInAudience() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) string { return v.SignInAudience }).(pulumi.StringOutput)
}

// A list of tags applied to the service principal.
func (o LookupServicePrincipalResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Possible values are `User` or `Admin`.
func (o LookupServicePrincipalResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServicePrincipalResultOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuread/sdk/v5/go/azuread/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Application within Azure Active Directory.
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
// <!--Start PulumiCodeChooser -->
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
//			example, err := azuread.LookupApplication(ctx, &azuread.LookupApplicationArgs{
//				DisplayName: pulumi.StringRef("My First AzureAD Application"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("applicationObjectId", example.ObjectId)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func LookupApplication(ctx *pulumi.Context, args *LookupApplicationArgs, opts ...pulumi.InvokeOption) (*LookupApplicationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupApplicationResult
	err := ctx.Invoke("azuread:index/getApplication:getApplication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApplication.
type LookupApplicationArgs struct {
	// Deprecated: The `applicationId` property has been replaced with the `clientId` property and will be removed in version 3.0 of the AzureAD provider
	ApplicationId *string `pulumi:"applicationId"`
	// Specifies the Client ID of the application.
	ClientId *string `pulumi:"clientId"`
	// Specifies the display name of the application.
	DisplayName *string `pulumi:"displayName"`
	// Specifies the Object ID of the application.
	//
	// > One of `clientId`, `displayName`, or `objectId` must be specified.
	ObjectId *string `pulumi:"objectId"`
}

// A collection of values returned by getApplication.
type LookupApplicationResult struct {
	// An `api` block as documented below.
	Apis []GetApplicationApi `pulumi:"apis"`
	// A mapping of app role values to app role IDs, intended to be useful when referencing app roles in other resources in your configuration.
	AppRoleIds map[string]string `pulumi:"appRoleIds"`
	// A collection of `appRole` blocks as documented below. For more information see [official documentation on Application Roles](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).
	AppRoles []GetApplicationAppRoleType `pulumi:"appRoles"`
	// Deprecated: The `applicationId` property has been replaced with the `clientId` property and will be removed in version 3.0 of the AzureAD provider
	ApplicationId string `pulumi:"applicationId"`
	// The Client ID for the application.
	ClientId string `pulumi:"clientId"`
	// Description of the app role that appears when the role is being assigned and, if the role functions as an application permissions, during the consent experiences.
	Description string `pulumi:"description"`
	// Specifies whether this application supports device authentication without a user.
	DeviceOnlyAuthEnabled bool `pulumi:"deviceOnlyAuthEnabled"`
	// Whether Microsoft has disabled the registered application. If the application is disabled, this will be a string indicating the status/reason, e.g. `DisabledDueToViolationOfServicesAgreement`
	DisabledByMicrosoft string `pulumi:"disabledByMicrosoft"`
	// Display name for the app role that appears during app role assignment and in consent experiences.
	DisplayName string `pulumi:"displayName"`
	// The fallback application type as public client, such as an installed application running on a mobile device.
	FallbackPublicClientEnabled bool `pulumi:"fallbackPublicClientEnabled"`
	// A `features` block as described below.
	FeatureTags []GetApplicationFeatureTag `pulumi:"featureTags"`
	// The `groups` claim issued in a user or OAuth 2.0 access token that the app expects.
	GroupMembershipClaims []string `pulumi:"groupMembershipClaims"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
	IdentifierUris []string `pulumi:"identifierUris"`
	// CDN URL to the application's logo.
	LogoUrl string `pulumi:"logoUrl"`
	// URL of the application's marketing page.
	MarketingUrl string `pulumi:"marketingUrl"`
	// User-specified notes relevant for the management of the application.
	Notes string `pulumi:"notes"`
	// A mapping of OAuth2.0 permission scope values to scope IDs, intended to be useful when referencing permission scopes in other resources in your configuration.
	Oauth2PermissionScopeIds map[string]string `pulumi:"oauth2PermissionScopeIds"`
	// Specifies whether, as part of OAuth 2.0 token requests, Azure AD allows POST requests, as opposed to GET requests. When `false`, only GET requests are allowed.
	Oauth2PostResponseRequired bool `pulumi:"oauth2PostResponseRequired"`
	// The application's object ID.
	ObjectId string `pulumi:"objectId"`
	// An `optionalClaims` block as documented below.
	OptionalClaims []GetApplicationOptionalClaim `pulumi:"optionalClaims"`
	// A list of object IDs of principals that are assigned ownership of the application.
	Owners []string `pulumi:"owners"`
	// URL of the application's privacy statement.
	PrivacyStatementUrl string `pulumi:"privacyStatementUrl"`
	// A `publicClient` block as documented below.
	PublicClients []GetApplicationPublicClient `pulumi:"publicClients"`
	// The verified publisher domain for the application.
	PublisherDomain string `pulumi:"publisherDomain"`
	// A collection of `requiredResourceAccess` blocks as documented below.
	RequiredResourceAccesses []GetApplicationRequiredResourceAccess `pulumi:"requiredResourceAccesses"`
	// References application context information from a Service or Asset Management database.
	ServiceManagementReference string `pulumi:"serviceManagementReference"`
	// The Microsoft account types that are supported for the current application. One of `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`.
	SignInAudience string `pulumi:"signInAudience"`
	// A `singlePageApplication` block as documented below.
	SinglePageApplications []GetApplicationSinglePageApplication `pulumi:"singlePageApplications"`
	// URL of the application's support page.
	SupportUrl string `pulumi:"supportUrl"`
	// A list of tags applied to the application.
	Tags []string `pulumi:"tags"`
	// URL of the application's terms of service statement.
	TermsOfServiceUrl string `pulumi:"termsOfServiceUrl"`
	// A `web` block as documented below.
	Webs []GetApplicationWeb `pulumi:"webs"`
}

func LookupApplicationOutput(ctx *pulumi.Context, args LookupApplicationOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationResult, error) {
			args := v.(LookupApplicationArgs)
			r, err := LookupApplication(ctx, &args, opts...)
			var s LookupApplicationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplicationResultOutput)
}

// A collection of arguments for invoking getApplication.
type LookupApplicationOutputArgs struct {
	// Deprecated: The `applicationId` property has been replaced with the `clientId` property and will be removed in version 3.0 of the AzureAD provider
	ApplicationId pulumi.StringPtrInput `pulumi:"applicationId"`
	// Specifies the Client ID of the application.
	ClientId pulumi.StringPtrInput `pulumi:"clientId"`
	// Specifies the display name of the application.
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
	// Specifies the Object ID of the application.
	//
	// > One of `clientId`, `displayName`, or `objectId` must be specified.
	ObjectId pulumi.StringPtrInput `pulumi:"objectId"`
}

func (LookupApplicationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationArgs)(nil)).Elem()
}

// A collection of values returned by getApplication.
type LookupApplicationResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationResult)(nil)).Elem()
}

func (o LookupApplicationResultOutput) ToLookupApplicationResultOutput() LookupApplicationResultOutput {
	return o
}

func (o LookupApplicationResultOutput) ToLookupApplicationResultOutputWithContext(ctx context.Context) LookupApplicationResultOutput {
	return o
}

// An `api` block as documented below.
func (o LookupApplicationResultOutput) Apis() GetApplicationApiArrayOutput {
	return o.ApplyT(func(v LookupApplicationResult) []GetApplicationApi { return v.Apis }).(GetApplicationApiArrayOutput)
}

// A mapping of app role values to app role IDs, intended to be useful when referencing app roles in other resources in your configuration.
func (o LookupApplicationResultOutput) AppRoleIds() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApplicationResult) map[string]string { return v.AppRoleIds }).(pulumi.StringMapOutput)
}

// A collection of `appRole` blocks as documented below. For more information see [official documentation on Application Roles](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).
func (o LookupApplicationResultOutput) AppRoles() GetApplicationAppRoleTypeArrayOutput {
	return o.ApplyT(func(v LookupApplicationResult) []GetApplicationAppRoleType { return v.AppRoles }).(GetApplicationAppRoleTypeArrayOutput)
}

// Deprecated: The `applicationId` property has been replaced with the `clientId` property and will be removed in version 3.0 of the AzureAD provider
func (o LookupApplicationResultOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.ApplicationId }).(pulumi.StringOutput)
}

// The Client ID for the application.
func (o LookupApplicationResultOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.ClientId }).(pulumi.StringOutput)
}

// Description of the app role that appears when the role is being assigned and, if the role functions as an application permissions, during the consent experiences.
func (o LookupApplicationResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Description }).(pulumi.StringOutput)
}

// Specifies whether this application supports device authentication without a user.
func (o LookupApplicationResultOutput) DeviceOnlyAuthEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupApplicationResult) bool { return v.DeviceOnlyAuthEnabled }).(pulumi.BoolOutput)
}

// Whether Microsoft has disabled the registered application. If the application is disabled, this will be a string indicating the status/reason, e.g. `DisabledDueToViolationOfServicesAgreement`
func (o LookupApplicationResultOutput) DisabledByMicrosoft() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.DisabledByMicrosoft }).(pulumi.StringOutput)
}

// Display name for the app role that appears during app role assignment and in consent experiences.
func (o LookupApplicationResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// The fallback application type as public client, such as an installed application running on a mobile device.
func (o LookupApplicationResultOutput) FallbackPublicClientEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupApplicationResult) bool { return v.FallbackPublicClientEnabled }).(pulumi.BoolOutput)
}

// A `features` block as described below.
func (o LookupApplicationResultOutput) FeatureTags() GetApplicationFeatureTagArrayOutput {
	return o.ApplyT(func(v LookupApplicationResult) []GetApplicationFeatureTag { return v.FeatureTags }).(GetApplicationFeatureTagArrayOutput)
}

// The `groups` claim issued in a user or OAuth 2.0 access token that the app expects.
func (o LookupApplicationResultOutput) GroupMembershipClaims() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupApplicationResult) []string { return v.GroupMembershipClaims }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupApplicationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
func (o LookupApplicationResultOutput) IdentifierUris() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupApplicationResult) []string { return v.IdentifierUris }).(pulumi.StringArrayOutput)
}

// CDN URL to the application's logo.
func (o LookupApplicationResultOutput) LogoUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.LogoUrl }).(pulumi.StringOutput)
}

// URL of the application's marketing page.
func (o LookupApplicationResultOutput) MarketingUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.MarketingUrl }).(pulumi.StringOutput)
}

// User-specified notes relevant for the management of the application.
func (o LookupApplicationResultOutput) Notes() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Notes }).(pulumi.StringOutput)
}

// A mapping of OAuth2.0 permission scope values to scope IDs, intended to be useful when referencing permission scopes in other resources in your configuration.
func (o LookupApplicationResultOutput) Oauth2PermissionScopeIds() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApplicationResult) map[string]string { return v.Oauth2PermissionScopeIds }).(pulumi.StringMapOutput)
}

// Specifies whether, as part of OAuth 2.0 token requests, Azure AD allows POST requests, as opposed to GET requests. When `false`, only GET requests are allowed.
func (o LookupApplicationResultOutput) Oauth2PostResponseRequired() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupApplicationResult) bool { return v.Oauth2PostResponseRequired }).(pulumi.BoolOutput)
}

// The application's object ID.
func (o LookupApplicationResultOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.ObjectId }).(pulumi.StringOutput)
}

// An `optionalClaims` block as documented below.
func (o LookupApplicationResultOutput) OptionalClaims() GetApplicationOptionalClaimArrayOutput {
	return o.ApplyT(func(v LookupApplicationResult) []GetApplicationOptionalClaim { return v.OptionalClaims }).(GetApplicationOptionalClaimArrayOutput)
}

// A list of object IDs of principals that are assigned ownership of the application.
func (o LookupApplicationResultOutput) Owners() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupApplicationResult) []string { return v.Owners }).(pulumi.StringArrayOutput)
}

// URL of the application's privacy statement.
func (o LookupApplicationResultOutput) PrivacyStatementUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.PrivacyStatementUrl }).(pulumi.StringOutput)
}

// A `publicClient` block as documented below.
func (o LookupApplicationResultOutput) PublicClients() GetApplicationPublicClientArrayOutput {
	return o.ApplyT(func(v LookupApplicationResult) []GetApplicationPublicClient { return v.PublicClients }).(GetApplicationPublicClientArrayOutput)
}

// The verified publisher domain for the application.
func (o LookupApplicationResultOutput) PublisherDomain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.PublisherDomain }).(pulumi.StringOutput)
}

// A collection of `requiredResourceAccess` blocks as documented below.
func (o LookupApplicationResultOutput) RequiredResourceAccesses() GetApplicationRequiredResourceAccessArrayOutput {
	return o.ApplyT(func(v LookupApplicationResult) []GetApplicationRequiredResourceAccess {
		return v.RequiredResourceAccesses
	}).(GetApplicationRequiredResourceAccessArrayOutput)
}

// References application context information from a Service or Asset Management database.
func (o LookupApplicationResultOutput) ServiceManagementReference() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.ServiceManagementReference }).(pulumi.StringOutput)
}

// The Microsoft account types that are supported for the current application. One of `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`.
func (o LookupApplicationResultOutput) SignInAudience() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.SignInAudience }).(pulumi.StringOutput)
}

// A `singlePageApplication` block as documented below.
func (o LookupApplicationResultOutput) SinglePageApplications() GetApplicationSinglePageApplicationArrayOutput {
	return o.ApplyT(func(v LookupApplicationResult) []GetApplicationSinglePageApplication { return v.SinglePageApplications }).(GetApplicationSinglePageApplicationArrayOutput)
}

// URL of the application's support page.
func (o LookupApplicationResultOutput) SupportUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.SupportUrl }).(pulumi.StringOutput)
}

// A list of tags applied to the application.
func (o LookupApplicationResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupApplicationResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// URL of the application's terms of service statement.
func (o LookupApplicationResultOutput) TermsOfServiceUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.TermsOfServiceUrl }).(pulumi.StringOutput)
}

// A `web` block as documented below.
func (o LookupApplicationResultOutput) Webs() GetApplicationWebArrayOutput {
	return o.ApplyT(func(v LookupApplicationResult) []GetApplicationWeb { return v.Webs }).(GetApplicationWebArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationResultOutput{})
}

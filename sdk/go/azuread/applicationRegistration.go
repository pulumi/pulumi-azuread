// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azuread/sdk/v6/go/azuread/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an application registration within Azure Active Directory.
//
// For a more comprehensive alternative, please see the Application resource. Please note that this resource should not be used together with the `Application` resource when managing the same application.
//
// ## API Permissions
//
// The following API permissions are required in order to use this resource.
//
// When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.OwnedBy` or `Application.ReadWrite.All`
//
// When authenticated with a user principal, this resource may require one of the following directory roles: `Application Administrator` or `Global Administrator`
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-azuread/sdk/v6/go/azuread"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := azuread.NewApplicationRegistration(ctx, "example", &azuread.ApplicationRegistrationArgs{
//				DisplayName:         pulumi.String("Example Application"),
//				Description:         pulumi.String("My example application"),
//				SignInAudience:      pulumi.String("AzureADMyOrg"),
//				HomepageUrl:         pulumi.String("https://app.example.com/"),
//				LogoutUrl:           pulumi.String("https://app.example.com/logout"),
//				MarketingUrl:        pulumi.String("https://example.com/"),
//				PrivacyStatementUrl: pulumi.String("https://example.com/privacy"),
//				SupportUrl:          pulumi.String("https://support.example.com/"),
//				TermsOfServiceUrl:   pulumi.String("https://example.com/terms"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Application Registrations can be imported using the object ID of the application, in the following format.
//
// ```sh
// $ pulumi import azuread:index/applicationRegistration:ApplicationRegistration example /applications/00000000-0000-0000-0000-000000000000
// ```
type ApplicationRegistration struct {
	pulumi.CustomResourceState

	// The Client ID for the application, which is globally unique.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// A description of the application, as shown to end users.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether Microsoft has disabled the registered application. If the application is disabled, this will be a string indicating the status/reason, e.g. `DisabledDueToViolationOfServicesAgreement`
	DisabledByMicrosoft pulumi.StringOutput `pulumi:"disabledByMicrosoft"`
	// The display name for the application.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Configures the `groups` claim issued in a user or OAuth access token that the app expects. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
	GroupMembershipClaims pulumi.StringArrayOutput `pulumi:"groupMembershipClaims"`
	// Home page or landing page of the application.
	HomepageUrl pulumi.StringPtrOutput `pulumi:"homepageUrl"`
	// Whether this web application can request an access token using OAuth implicit flow.
	ImplicitAccessTokenIssuanceEnabled pulumi.BoolPtrOutput `pulumi:"implicitAccessTokenIssuanceEnabled"`
	// Whether this web application can request an ID token using OAuth implicit flow.
	ImplicitIdTokenIssuanceEnabled pulumi.BoolPtrOutput `pulumi:"implicitIdTokenIssuanceEnabled"`
	// The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols.
	LogoutUrl pulumi.StringPtrOutput `pulumi:"logoutUrl"`
	// URL of the marketing page for the application.
	MarketingUrl pulumi.StringPtrOutput `pulumi:"marketingUrl"`
	// User-specified notes relevant for the management of the application.
	Notes pulumi.StringPtrOutput `pulumi:"notes"`
	// The object ID of the application within the tenant.
	ObjectId pulumi.StringOutput `pulumi:"objectId"`
	// URL of the privacy statement for the application.
	PrivacyStatementUrl pulumi.StringPtrOutput `pulumi:"privacyStatementUrl"`
	// The verified publisher domain for the application.
	PublisherDomain pulumi.StringOutput `pulumi:"publisherDomain"`
	// The access token version expected by this resource. Must be one of `1` or `2`, and must be `2` when `signInAudience` is either `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount` Defaults to `2`.
	RequestedAccessTokenVersion pulumi.IntPtrOutput `pulumi:"requestedAccessTokenVersion"`
	// References application context information from a Service or Asset Management database.
	ServiceManagementReference pulumi.StringPtrOutput `pulumi:"serviceManagementReference"`
	// The Microsoft account types that are supported for the current application. Must be one of `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`. Defaults to `AzureADMyOrg`.
	SignInAudience pulumi.StringPtrOutput `pulumi:"signInAudience"`
	// URL of the support page for the application.
	SupportUrl pulumi.StringPtrOutput `pulumi:"supportUrl"`
	// URL of the terms of service statement for the application.
	TermsOfServiceUrl pulumi.StringPtrOutput `pulumi:"termsOfServiceUrl"`
}

// NewApplicationRegistration registers a new resource with the given unique name, arguments, and options.
func NewApplicationRegistration(ctx *pulumi.Context,
	name string, args *ApplicationRegistrationArgs, opts ...pulumi.ResourceOption) (*ApplicationRegistration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApplicationRegistration
	err := ctx.RegisterResource("azuread:index/applicationRegistration:ApplicationRegistration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationRegistration gets an existing ApplicationRegistration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationRegistration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationRegistrationState, opts ...pulumi.ResourceOption) (*ApplicationRegistration, error) {
	var resource ApplicationRegistration
	err := ctx.ReadResource("azuread:index/applicationRegistration:ApplicationRegistration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationRegistration resources.
type applicationRegistrationState struct {
	// The Client ID for the application, which is globally unique.
	ClientId *string `pulumi:"clientId"`
	// A description of the application, as shown to end users.
	Description *string `pulumi:"description"`
	// Whether Microsoft has disabled the registered application. If the application is disabled, this will be a string indicating the status/reason, e.g. `DisabledDueToViolationOfServicesAgreement`
	DisabledByMicrosoft *string `pulumi:"disabledByMicrosoft"`
	// The display name for the application.
	DisplayName *string `pulumi:"displayName"`
	// Configures the `groups` claim issued in a user or OAuth access token that the app expects. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
	GroupMembershipClaims []string `pulumi:"groupMembershipClaims"`
	// Home page or landing page of the application.
	HomepageUrl *string `pulumi:"homepageUrl"`
	// Whether this web application can request an access token using OAuth implicit flow.
	ImplicitAccessTokenIssuanceEnabled *bool `pulumi:"implicitAccessTokenIssuanceEnabled"`
	// Whether this web application can request an ID token using OAuth implicit flow.
	ImplicitIdTokenIssuanceEnabled *bool `pulumi:"implicitIdTokenIssuanceEnabled"`
	// The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols.
	LogoutUrl *string `pulumi:"logoutUrl"`
	// URL of the marketing page for the application.
	MarketingUrl *string `pulumi:"marketingUrl"`
	// User-specified notes relevant for the management of the application.
	Notes *string `pulumi:"notes"`
	// The object ID of the application within the tenant.
	ObjectId *string `pulumi:"objectId"`
	// URL of the privacy statement for the application.
	PrivacyStatementUrl *string `pulumi:"privacyStatementUrl"`
	// The verified publisher domain for the application.
	PublisherDomain *string `pulumi:"publisherDomain"`
	// The access token version expected by this resource. Must be one of `1` or `2`, and must be `2` when `signInAudience` is either `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount` Defaults to `2`.
	RequestedAccessTokenVersion *int `pulumi:"requestedAccessTokenVersion"`
	// References application context information from a Service or Asset Management database.
	ServiceManagementReference *string `pulumi:"serviceManagementReference"`
	// The Microsoft account types that are supported for the current application. Must be one of `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`. Defaults to `AzureADMyOrg`.
	SignInAudience *string `pulumi:"signInAudience"`
	// URL of the support page for the application.
	SupportUrl *string `pulumi:"supportUrl"`
	// URL of the terms of service statement for the application.
	TermsOfServiceUrl *string `pulumi:"termsOfServiceUrl"`
}

type ApplicationRegistrationState struct {
	// The Client ID for the application, which is globally unique.
	ClientId pulumi.StringPtrInput
	// A description of the application, as shown to end users.
	Description pulumi.StringPtrInput
	// Whether Microsoft has disabled the registered application. If the application is disabled, this will be a string indicating the status/reason, e.g. `DisabledDueToViolationOfServicesAgreement`
	DisabledByMicrosoft pulumi.StringPtrInput
	// The display name for the application.
	DisplayName pulumi.StringPtrInput
	// Configures the `groups` claim issued in a user or OAuth access token that the app expects. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
	GroupMembershipClaims pulumi.StringArrayInput
	// Home page or landing page of the application.
	HomepageUrl pulumi.StringPtrInput
	// Whether this web application can request an access token using OAuth implicit flow.
	ImplicitAccessTokenIssuanceEnabled pulumi.BoolPtrInput
	// Whether this web application can request an ID token using OAuth implicit flow.
	ImplicitIdTokenIssuanceEnabled pulumi.BoolPtrInput
	// The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols.
	LogoutUrl pulumi.StringPtrInput
	// URL of the marketing page for the application.
	MarketingUrl pulumi.StringPtrInput
	// User-specified notes relevant for the management of the application.
	Notes pulumi.StringPtrInput
	// The object ID of the application within the tenant.
	ObjectId pulumi.StringPtrInput
	// URL of the privacy statement for the application.
	PrivacyStatementUrl pulumi.StringPtrInput
	// The verified publisher domain for the application.
	PublisherDomain pulumi.StringPtrInput
	// The access token version expected by this resource. Must be one of `1` or `2`, and must be `2` when `signInAudience` is either `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount` Defaults to `2`.
	RequestedAccessTokenVersion pulumi.IntPtrInput
	// References application context information from a Service or Asset Management database.
	ServiceManagementReference pulumi.StringPtrInput
	// The Microsoft account types that are supported for the current application. Must be one of `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`. Defaults to `AzureADMyOrg`.
	SignInAudience pulumi.StringPtrInput
	// URL of the support page for the application.
	SupportUrl pulumi.StringPtrInput
	// URL of the terms of service statement for the application.
	TermsOfServiceUrl pulumi.StringPtrInput
}

func (ApplicationRegistrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationRegistrationState)(nil)).Elem()
}

type applicationRegistrationArgs struct {
	// A description of the application, as shown to end users.
	Description *string `pulumi:"description"`
	// The display name for the application.
	DisplayName string `pulumi:"displayName"`
	// Configures the `groups` claim issued in a user or OAuth access token that the app expects. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
	GroupMembershipClaims []string `pulumi:"groupMembershipClaims"`
	// Home page or landing page of the application.
	HomepageUrl *string `pulumi:"homepageUrl"`
	// Whether this web application can request an access token using OAuth implicit flow.
	ImplicitAccessTokenIssuanceEnabled *bool `pulumi:"implicitAccessTokenIssuanceEnabled"`
	// Whether this web application can request an ID token using OAuth implicit flow.
	ImplicitIdTokenIssuanceEnabled *bool `pulumi:"implicitIdTokenIssuanceEnabled"`
	// The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols.
	LogoutUrl *string `pulumi:"logoutUrl"`
	// URL of the marketing page for the application.
	MarketingUrl *string `pulumi:"marketingUrl"`
	// User-specified notes relevant for the management of the application.
	Notes *string `pulumi:"notes"`
	// URL of the privacy statement for the application.
	PrivacyStatementUrl *string `pulumi:"privacyStatementUrl"`
	// The access token version expected by this resource. Must be one of `1` or `2`, and must be `2` when `signInAudience` is either `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount` Defaults to `2`.
	RequestedAccessTokenVersion *int `pulumi:"requestedAccessTokenVersion"`
	// References application context information from a Service or Asset Management database.
	ServiceManagementReference *string `pulumi:"serviceManagementReference"`
	// The Microsoft account types that are supported for the current application. Must be one of `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`. Defaults to `AzureADMyOrg`.
	SignInAudience *string `pulumi:"signInAudience"`
	// URL of the support page for the application.
	SupportUrl *string `pulumi:"supportUrl"`
	// URL of the terms of service statement for the application.
	TermsOfServiceUrl *string `pulumi:"termsOfServiceUrl"`
}

// The set of arguments for constructing a ApplicationRegistration resource.
type ApplicationRegistrationArgs struct {
	// A description of the application, as shown to end users.
	Description pulumi.StringPtrInput
	// The display name for the application.
	DisplayName pulumi.StringInput
	// Configures the `groups` claim issued in a user or OAuth access token that the app expects. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
	GroupMembershipClaims pulumi.StringArrayInput
	// Home page or landing page of the application.
	HomepageUrl pulumi.StringPtrInput
	// Whether this web application can request an access token using OAuth implicit flow.
	ImplicitAccessTokenIssuanceEnabled pulumi.BoolPtrInput
	// Whether this web application can request an ID token using OAuth implicit flow.
	ImplicitIdTokenIssuanceEnabled pulumi.BoolPtrInput
	// The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols.
	LogoutUrl pulumi.StringPtrInput
	// URL of the marketing page for the application.
	MarketingUrl pulumi.StringPtrInput
	// User-specified notes relevant for the management of the application.
	Notes pulumi.StringPtrInput
	// URL of the privacy statement for the application.
	PrivacyStatementUrl pulumi.StringPtrInput
	// The access token version expected by this resource. Must be one of `1` or `2`, and must be `2` when `signInAudience` is either `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount` Defaults to `2`.
	RequestedAccessTokenVersion pulumi.IntPtrInput
	// References application context information from a Service or Asset Management database.
	ServiceManagementReference pulumi.StringPtrInput
	// The Microsoft account types that are supported for the current application. Must be one of `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`. Defaults to `AzureADMyOrg`.
	SignInAudience pulumi.StringPtrInput
	// URL of the support page for the application.
	SupportUrl pulumi.StringPtrInput
	// URL of the terms of service statement for the application.
	TermsOfServiceUrl pulumi.StringPtrInput
}

func (ApplicationRegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationRegistrationArgs)(nil)).Elem()
}

type ApplicationRegistrationInput interface {
	pulumi.Input

	ToApplicationRegistrationOutput() ApplicationRegistrationOutput
	ToApplicationRegistrationOutputWithContext(ctx context.Context) ApplicationRegistrationOutput
}

func (*ApplicationRegistration) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationRegistration)(nil)).Elem()
}

func (i *ApplicationRegistration) ToApplicationRegistrationOutput() ApplicationRegistrationOutput {
	return i.ToApplicationRegistrationOutputWithContext(context.Background())
}

func (i *ApplicationRegistration) ToApplicationRegistrationOutputWithContext(ctx context.Context) ApplicationRegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationRegistrationOutput)
}

// ApplicationRegistrationArrayInput is an input type that accepts ApplicationRegistrationArray and ApplicationRegistrationArrayOutput values.
// You can construct a concrete instance of `ApplicationRegistrationArrayInput` via:
//
//	ApplicationRegistrationArray{ ApplicationRegistrationArgs{...} }
type ApplicationRegistrationArrayInput interface {
	pulumi.Input

	ToApplicationRegistrationArrayOutput() ApplicationRegistrationArrayOutput
	ToApplicationRegistrationArrayOutputWithContext(context.Context) ApplicationRegistrationArrayOutput
}

type ApplicationRegistrationArray []ApplicationRegistrationInput

func (ApplicationRegistrationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationRegistration)(nil)).Elem()
}

func (i ApplicationRegistrationArray) ToApplicationRegistrationArrayOutput() ApplicationRegistrationArrayOutput {
	return i.ToApplicationRegistrationArrayOutputWithContext(context.Background())
}

func (i ApplicationRegistrationArray) ToApplicationRegistrationArrayOutputWithContext(ctx context.Context) ApplicationRegistrationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationRegistrationArrayOutput)
}

// ApplicationRegistrationMapInput is an input type that accepts ApplicationRegistrationMap and ApplicationRegistrationMapOutput values.
// You can construct a concrete instance of `ApplicationRegistrationMapInput` via:
//
//	ApplicationRegistrationMap{ "key": ApplicationRegistrationArgs{...} }
type ApplicationRegistrationMapInput interface {
	pulumi.Input

	ToApplicationRegistrationMapOutput() ApplicationRegistrationMapOutput
	ToApplicationRegistrationMapOutputWithContext(context.Context) ApplicationRegistrationMapOutput
}

type ApplicationRegistrationMap map[string]ApplicationRegistrationInput

func (ApplicationRegistrationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationRegistration)(nil)).Elem()
}

func (i ApplicationRegistrationMap) ToApplicationRegistrationMapOutput() ApplicationRegistrationMapOutput {
	return i.ToApplicationRegistrationMapOutputWithContext(context.Background())
}

func (i ApplicationRegistrationMap) ToApplicationRegistrationMapOutputWithContext(ctx context.Context) ApplicationRegistrationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationRegistrationMapOutput)
}

type ApplicationRegistrationOutput struct{ *pulumi.OutputState }

func (ApplicationRegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationRegistration)(nil)).Elem()
}

func (o ApplicationRegistrationOutput) ToApplicationRegistrationOutput() ApplicationRegistrationOutput {
	return o
}

func (o ApplicationRegistrationOutput) ToApplicationRegistrationOutputWithContext(ctx context.Context) ApplicationRegistrationOutput {
	return o
}

// The Client ID for the application, which is globally unique.
func (o ApplicationRegistrationOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationRegistration) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

// A description of the application, as shown to end users.
func (o ApplicationRegistrationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationRegistration) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Whether Microsoft has disabled the registered application. If the application is disabled, this will be a string indicating the status/reason, e.g. `DisabledDueToViolationOfServicesAgreement`
func (o ApplicationRegistrationOutput) DisabledByMicrosoft() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationRegistration) pulumi.StringOutput { return v.DisabledByMicrosoft }).(pulumi.StringOutput)
}

// The display name for the application.
func (o ApplicationRegistrationOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationRegistration) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Configures the `groups` claim issued in a user or OAuth access token that the app expects. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
func (o ApplicationRegistrationOutput) GroupMembershipClaims() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationRegistration) pulumi.StringArrayOutput { return v.GroupMembershipClaims }).(pulumi.StringArrayOutput)
}

// Home page or landing page of the application.
func (o ApplicationRegistrationOutput) HomepageUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationRegistration) pulumi.StringPtrOutput { return v.HomepageUrl }).(pulumi.StringPtrOutput)
}

// Whether this web application can request an access token using OAuth implicit flow.
func (o ApplicationRegistrationOutput) ImplicitAccessTokenIssuanceEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationRegistration) pulumi.BoolPtrOutput { return v.ImplicitAccessTokenIssuanceEnabled }).(pulumi.BoolPtrOutput)
}

// Whether this web application can request an ID token using OAuth implicit flow.
func (o ApplicationRegistrationOutput) ImplicitIdTokenIssuanceEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationRegistration) pulumi.BoolPtrOutput { return v.ImplicitIdTokenIssuanceEnabled }).(pulumi.BoolPtrOutput)
}

// The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols.
func (o ApplicationRegistrationOutput) LogoutUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationRegistration) pulumi.StringPtrOutput { return v.LogoutUrl }).(pulumi.StringPtrOutput)
}

// URL of the marketing page for the application.
func (o ApplicationRegistrationOutput) MarketingUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationRegistration) pulumi.StringPtrOutput { return v.MarketingUrl }).(pulumi.StringPtrOutput)
}

// User-specified notes relevant for the management of the application.
func (o ApplicationRegistrationOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationRegistration) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

// The object ID of the application within the tenant.
func (o ApplicationRegistrationOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationRegistration) pulumi.StringOutput { return v.ObjectId }).(pulumi.StringOutput)
}

// URL of the privacy statement for the application.
func (o ApplicationRegistrationOutput) PrivacyStatementUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationRegistration) pulumi.StringPtrOutput { return v.PrivacyStatementUrl }).(pulumi.StringPtrOutput)
}

// The verified publisher domain for the application.
func (o ApplicationRegistrationOutput) PublisherDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationRegistration) pulumi.StringOutput { return v.PublisherDomain }).(pulumi.StringOutput)
}

// The access token version expected by this resource. Must be one of `1` or `2`, and must be `2` when `signInAudience` is either `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount` Defaults to `2`.
func (o ApplicationRegistrationOutput) RequestedAccessTokenVersion() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ApplicationRegistration) pulumi.IntPtrOutput { return v.RequestedAccessTokenVersion }).(pulumi.IntPtrOutput)
}

// References application context information from a Service or Asset Management database.
func (o ApplicationRegistrationOutput) ServiceManagementReference() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationRegistration) pulumi.StringPtrOutput { return v.ServiceManagementReference }).(pulumi.StringPtrOutput)
}

// The Microsoft account types that are supported for the current application. Must be one of `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`. Defaults to `AzureADMyOrg`.
func (o ApplicationRegistrationOutput) SignInAudience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationRegistration) pulumi.StringPtrOutput { return v.SignInAudience }).(pulumi.StringPtrOutput)
}

// URL of the support page for the application.
func (o ApplicationRegistrationOutput) SupportUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationRegistration) pulumi.StringPtrOutput { return v.SupportUrl }).(pulumi.StringPtrOutput)
}

// URL of the terms of service statement for the application.
func (o ApplicationRegistrationOutput) TermsOfServiceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationRegistration) pulumi.StringPtrOutput { return v.TermsOfServiceUrl }).(pulumi.StringPtrOutput)
}

type ApplicationRegistrationArrayOutput struct{ *pulumi.OutputState }

func (ApplicationRegistrationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationRegistration)(nil)).Elem()
}

func (o ApplicationRegistrationArrayOutput) ToApplicationRegistrationArrayOutput() ApplicationRegistrationArrayOutput {
	return o
}

func (o ApplicationRegistrationArrayOutput) ToApplicationRegistrationArrayOutputWithContext(ctx context.Context) ApplicationRegistrationArrayOutput {
	return o
}

func (o ApplicationRegistrationArrayOutput) Index(i pulumi.IntInput) ApplicationRegistrationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplicationRegistration {
		return vs[0].([]*ApplicationRegistration)[vs[1].(int)]
	}).(ApplicationRegistrationOutput)
}

type ApplicationRegistrationMapOutput struct{ *pulumi.OutputState }

func (ApplicationRegistrationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationRegistration)(nil)).Elem()
}

func (o ApplicationRegistrationMapOutput) ToApplicationRegistrationMapOutput() ApplicationRegistrationMapOutput {
	return o
}

func (o ApplicationRegistrationMapOutput) ToApplicationRegistrationMapOutputWithContext(ctx context.Context) ApplicationRegistrationMapOutput {
	return o
}

func (o ApplicationRegistrationMapOutput) MapIndex(k pulumi.StringInput) ApplicationRegistrationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplicationRegistration {
		return vs[0].(map[string]*ApplicationRegistration)[vs[1].(string)]
	}).(ApplicationRegistrationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationRegistrationInput)(nil)).Elem(), &ApplicationRegistration{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationRegistrationArrayInput)(nil)).Elem(), ApplicationRegistrationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationRegistrationMapInput)(nil)).Elem(), ApplicationRegistrationMap{})
	pulumi.RegisterOutputType(ApplicationRegistrationOutput{})
	pulumi.RegisterOutputType(ApplicationRegistrationArrayOutput{})
	pulumi.RegisterOutputType(ApplicationRegistrationMapOutput{})
}

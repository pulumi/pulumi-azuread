// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServicePrincipal(ctx *pulumi.Context, args *LookupServicePrincipalArgs, opts ...pulumi.InvokeOption) (*LookupServicePrincipalResult, error) {
	var rv LookupServicePrincipalResult
	err := ctx.Invoke("azuread:index/getServicePrincipal:getServicePrincipal", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServicePrincipal.
type LookupServicePrincipalArgs struct {
	ApplicationId *string `pulumi:"applicationId"`
	DisplayName   *string `pulumi:"displayName"`
	ObjectId      *string `pulumi:"objectId"`
}

// A collection of values returned by getServicePrincipal.
type LookupServicePrincipalResult struct {
	AccountEnabled            bool                            `pulumi:"accountEnabled"`
	AlternativeNames          []string                        `pulumi:"alternativeNames"`
	AppRoleAssignmentRequired bool                            `pulumi:"appRoleAssignmentRequired"`
	AppRoleIds                map[string]string               `pulumi:"appRoleIds"`
	AppRoles                  []GetServicePrincipalAppRole    `pulumi:"appRoles"`
	ApplicationId             string                          `pulumi:"applicationId"`
	ApplicationTenantId       string                          `pulumi:"applicationTenantId"`
	Description               string                          `pulumi:"description"`
	DisplayName               string                          `pulumi:"displayName"`
	FeatureTags               []GetServicePrincipalFeatureTag `pulumi:"featureTags"`
	// Deprecated: This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider
	Features    []GetServicePrincipalFeature `pulumi:"features"`
	HomepageUrl string                       `pulumi:"homepageUrl"`
	// The provider-assigned unique ID for this managed resource.
	Id                         string                                     `pulumi:"id"`
	LoginUrl                   string                                     `pulumi:"loginUrl"`
	LogoutUrl                  string                                     `pulumi:"logoutUrl"`
	Notes                      string                                     `pulumi:"notes"`
	NotificationEmailAddresses []string                                   `pulumi:"notificationEmailAddresses"`
	Oauth2PermissionScopeIds   map[string]string                          `pulumi:"oauth2PermissionScopeIds"`
	Oauth2PermissionScopes     []GetServicePrincipalOauth2PermissionScope `pulumi:"oauth2PermissionScopes"`
	ObjectId                   string                                     `pulumi:"objectId"`
	PreferredSingleSignOnMode  string                                     `pulumi:"preferredSingleSignOnMode"`
	RedirectUris               []string                                   `pulumi:"redirectUris"`
	SamlMetadataUrl            string                                     `pulumi:"samlMetadataUrl"`
	SamlSingleSignOns          []GetServicePrincipalSamlSingleSignOn      `pulumi:"samlSingleSignOns"`
	ServicePrincipalNames      []string                                   `pulumi:"servicePrincipalNames"`
	SignInAudience             string                                     `pulumi:"signInAudience"`
	Tags                       []string                                   `pulumi:"tags"`
	Type                       string                                     `pulumi:"type"`
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
	ApplicationId pulumi.StringPtrInput `pulumi:"applicationId"`
	DisplayName   pulumi.StringPtrInput `pulumi:"displayName"`
	ObjectId      pulumi.StringPtrInput `pulumi:"objectId"`
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

func (o LookupServicePrincipalResultOutput) AccountEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) bool { return v.AccountEnabled }).(pulumi.BoolOutput)
}

func (o LookupServicePrincipalResultOutput) AlternativeNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) []string { return v.AlternativeNames }).(pulumi.StringArrayOutput)
}

func (o LookupServicePrincipalResultOutput) AppRoleAssignmentRequired() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) bool { return v.AppRoleAssignmentRequired }).(pulumi.BoolOutput)
}

func (o LookupServicePrincipalResultOutput) AppRoleIds() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) map[string]string { return v.AppRoleIds }).(pulumi.StringMapOutput)
}

func (o LookupServicePrincipalResultOutput) AppRoles() GetServicePrincipalAppRoleArrayOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) []GetServicePrincipalAppRole { return v.AppRoles }).(GetServicePrincipalAppRoleArrayOutput)
}

func (o LookupServicePrincipalResultOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) string { return v.ApplicationId }).(pulumi.StringOutput)
}

func (o LookupServicePrincipalResultOutput) ApplicationTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) string { return v.ApplicationTenantId }).(pulumi.StringOutput)
}

func (o LookupServicePrincipalResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupServicePrincipalResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupServicePrincipalResultOutput) FeatureTags() GetServicePrincipalFeatureTagArrayOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) []GetServicePrincipalFeatureTag { return v.FeatureTags }).(GetServicePrincipalFeatureTagArrayOutput)
}

// Deprecated: This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider
func (o LookupServicePrincipalResultOutput) Features() GetServicePrincipalFeatureArrayOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) []GetServicePrincipalFeature { return v.Features }).(GetServicePrincipalFeatureArrayOutput)
}

func (o LookupServicePrincipalResultOutput) HomepageUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) string { return v.HomepageUrl }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupServicePrincipalResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServicePrincipalResultOutput) LoginUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) string { return v.LoginUrl }).(pulumi.StringOutput)
}

func (o LookupServicePrincipalResultOutput) LogoutUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) string { return v.LogoutUrl }).(pulumi.StringOutput)
}

func (o LookupServicePrincipalResultOutput) Notes() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) string { return v.Notes }).(pulumi.StringOutput)
}

func (o LookupServicePrincipalResultOutput) NotificationEmailAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) []string { return v.NotificationEmailAddresses }).(pulumi.StringArrayOutput)
}

func (o LookupServicePrincipalResultOutput) Oauth2PermissionScopeIds() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) map[string]string { return v.Oauth2PermissionScopeIds }).(pulumi.StringMapOutput)
}

func (o LookupServicePrincipalResultOutput) Oauth2PermissionScopes() GetServicePrincipalOauth2PermissionScopeArrayOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) []GetServicePrincipalOauth2PermissionScope {
		return v.Oauth2PermissionScopes
	}).(GetServicePrincipalOauth2PermissionScopeArrayOutput)
}

func (o LookupServicePrincipalResultOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) string { return v.ObjectId }).(pulumi.StringOutput)
}

func (o LookupServicePrincipalResultOutput) PreferredSingleSignOnMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) string { return v.PreferredSingleSignOnMode }).(pulumi.StringOutput)
}

func (o LookupServicePrincipalResultOutput) RedirectUris() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) []string { return v.RedirectUris }).(pulumi.StringArrayOutput)
}

func (o LookupServicePrincipalResultOutput) SamlMetadataUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) string { return v.SamlMetadataUrl }).(pulumi.StringOutput)
}

func (o LookupServicePrincipalResultOutput) SamlSingleSignOns() GetServicePrincipalSamlSingleSignOnArrayOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) []GetServicePrincipalSamlSingleSignOn { return v.SamlSingleSignOns }).(GetServicePrincipalSamlSingleSignOnArrayOutput)
}

func (o LookupServicePrincipalResultOutput) ServicePrincipalNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) []string { return v.ServicePrincipalNames }).(pulumi.StringArrayOutput)
}

func (o LookupServicePrincipalResultOutput) SignInAudience() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) string { return v.SignInAudience }).(pulumi.StringOutput)
}

func (o LookupServicePrincipalResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupServicePrincipalResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServicePrincipalResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServicePrincipalResultOutput{})
}

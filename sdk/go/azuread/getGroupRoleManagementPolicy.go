// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuread/sdk/v6/go/azuread/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve a role policy for an Azure AD group.
//
// ## API Permissions
//
// The following API permissions are required in order to use this resource.
//
// When authenticated with a service principal, this resource requires the `RoleManagementPolicy.Read.AzureADGroup` Microsoft Graph API permissions.
//
// When authenticated with a user principal, this resource requires `Global Administrator` directory role, or the `Privileged Role Administrator` role in Identity Governance.
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
//			example, err := azuread.NewGroup(ctx, "example", &azuread.GroupArgs{
//				DisplayName:     pulumi.String("group-name"),
//				SecurityEnabled: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_ = azuread.LookupGroupRoleManagementPolicyOutput(ctx, azuread.GetGroupRoleManagementPolicyOutputArgs{
//				GroupId: example.ID(),
//				RoleId:  pulumi.String("owner"),
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func LookupGroupRoleManagementPolicy(ctx *pulumi.Context, args *LookupGroupRoleManagementPolicyArgs, opts ...pulumi.InvokeOption) (*LookupGroupRoleManagementPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGroupRoleManagementPolicyResult
	err := ctx.Invoke("azuread:index/getGroupRoleManagementPolicy:getGroupRoleManagementPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroupRoleManagementPolicy.
type LookupGroupRoleManagementPolicyArgs struct {
	// The ID of the Azure AD group for which the policy applies.
	GroupId string `pulumi:"groupId"`
	// The type of assignment this policy coveres. Can be either `member` or `owner`.
	RoleId string `pulumi:"roleId"`
}

// A collection of values returned by getGroupRoleManagementPolicy.
type LookupGroupRoleManagementPolicyResult struct {
	// (String) The description of this policy.
	Description string `pulumi:"description"`
	// (String) The display name of this policy.
	DisplayName string `pulumi:"displayName"`
	GroupId     string `pulumi:"groupId"`
	// The provider-assigned unique ID for this managed resource.
	Id     string `pulumi:"id"`
	RoleId string `pulumi:"roleId"`
}

func LookupGroupRoleManagementPolicyOutput(ctx *pulumi.Context, args LookupGroupRoleManagementPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupGroupRoleManagementPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupGroupRoleManagementPolicyResultOutput, error) {
			args := v.(LookupGroupRoleManagementPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azuread:index/getGroupRoleManagementPolicy:getGroupRoleManagementPolicy", args, LookupGroupRoleManagementPolicyResultOutput{}, options).(LookupGroupRoleManagementPolicyResultOutput), nil
		}).(LookupGroupRoleManagementPolicyResultOutput)
}

// A collection of arguments for invoking getGroupRoleManagementPolicy.
type LookupGroupRoleManagementPolicyOutputArgs struct {
	// The ID of the Azure AD group for which the policy applies.
	GroupId pulumi.StringInput `pulumi:"groupId"`
	// The type of assignment this policy coveres. Can be either `member` or `owner`.
	RoleId pulumi.StringInput `pulumi:"roleId"`
}

func (LookupGroupRoleManagementPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupRoleManagementPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getGroupRoleManagementPolicy.
type LookupGroupRoleManagementPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupGroupRoleManagementPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupRoleManagementPolicyResult)(nil)).Elem()
}

func (o LookupGroupRoleManagementPolicyResultOutput) ToLookupGroupRoleManagementPolicyResultOutput() LookupGroupRoleManagementPolicyResultOutput {
	return o
}

func (o LookupGroupRoleManagementPolicyResultOutput) ToLookupGroupRoleManagementPolicyResultOutputWithContext(ctx context.Context) LookupGroupRoleManagementPolicyResultOutput {
	return o
}

// (String) The description of this policy.
func (o LookupGroupRoleManagementPolicyResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupRoleManagementPolicyResult) string { return v.Description }).(pulumi.StringOutput)
}

// (String) The display name of this policy.
func (o LookupGroupRoleManagementPolicyResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupRoleManagementPolicyResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupGroupRoleManagementPolicyResultOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupRoleManagementPolicyResult) string { return v.GroupId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupGroupRoleManagementPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupRoleManagementPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGroupRoleManagementPolicyResultOutput) RoleId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupRoleManagementPolicyResult) string { return v.RoleId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGroupRoleManagementPolicyResultOutput{})
}

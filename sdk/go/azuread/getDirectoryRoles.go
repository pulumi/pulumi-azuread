// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuread/sdk/v6/go/azuread/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about activated directory roles within Azure Active Directory.
//
// ## API Permissions
//
// The following API permissions are required in order to use this resource.
//
// When authenticated with a service principal, this resource requires one of the following application roles: `RoleManagement.Read.Directory` or `Directory.Read.All`
//
// When authenticated with a user principal, this data source does not require any additional roles.
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
//			current, err := azuread.GetDirectoryRoles(ctx, map[string]interface{}{}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("roles", current.ObjectIds)
//			return nil
//		})
//	}
//
// ```
func GetDirectoryRoles(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetDirectoryRolesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDirectoryRolesResult
	err := ctx.Invoke("azuread:index/getDirectoryRoles:getDirectoryRoles", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getDirectoryRoles.
type GetDirectoryRolesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The object IDs of the roles.
	ObjectIds []string `pulumi:"objectIds"`
	// A list of users. Each `role` object provides the attributes documented below.
	Roles []GetDirectoryRolesRole `pulumi:"roles"`
	// The template IDs of the roles.
	TemplateIds []string `pulumi:"templateIds"`
}

func GetDirectoryRolesOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetDirectoryRolesResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetDirectoryRolesResultOutput, error) {
		opts = internal.PkgInvokeDefaultOpts(opts)
		var rv GetDirectoryRolesResult
		secret, err := ctx.InvokePackageRaw("azuread:index/getDirectoryRoles:getDirectoryRoles", nil, &rv, "", opts...)
		if err != nil {
			return GetDirectoryRolesResultOutput{}, err
		}

		output := pulumi.ToOutput(rv).(GetDirectoryRolesResultOutput)
		if secret {
			return pulumi.ToSecret(output).(GetDirectoryRolesResultOutput), nil
		}
		return output, nil
	}).(GetDirectoryRolesResultOutput)
}

// A collection of values returned by getDirectoryRoles.
type GetDirectoryRolesResultOutput struct{ *pulumi.OutputState }

func (GetDirectoryRolesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDirectoryRolesResult)(nil)).Elem()
}

func (o GetDirectoryRolesResultOutput) ToGetDirectoryRolesResultOutput() GetDirectoryRolesResultOutput {
	return o
}

func (o GetDirectoryRolesResultOutput) ToGetDirectoryRolesResultOutputWithContext(ctx context.Context) GetDirectoryRolesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetDirectoryRolesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDirectoryRolesResult) string { return v.Id }).(pulumi.StringOutput)
}

// The object IDs of the roles.
func (o GetDirectoryRolesResultOutput) ObjectIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDirectoryRolesResult) []string { return v.ObjectIds }).(pulumi.StringArrayOutput)
}

// A list of users. Each `role` object provides the attributes documented below.
func (o GetDirectoryRolesResultOutput) Roles() GetDirectoryRolesRoleArrayOutput {
	return o.ApplyT(func(v GetDirectoryRolesResult) []GetDirectoryRolesRole { return v.Roles }).(GetDirectoryRolesRoleArrayOutput)
}

// The template IDs of the roles.
func (o GetDirectoryRolesResultOutput) TemplateIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDirectoryRolesResult) []string { return v.TemplateIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDirectoryRolesResultOutput{})
}

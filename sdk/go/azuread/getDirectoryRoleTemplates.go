// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"github.com/pulumi/pulumi-azuread/sdk/v5/go/azuread/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about directory role templates within Azure Active Directory.
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
//	"github.com/pulumi/pulumi-azuread/sdk/v5/go/azuread"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			current, err := azuread.GetDirectoryRoleTemplates(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("roles", current.ObjectIds)
//			return nil
//		})
//	}
//
// ```
func GetDirectoryRoleTemplates(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetDirectoryRoleTemplatesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDirectoryRoleTemplatesResult
	err := ctx.Invoke("azuread:index/getDirectoryRoleTemplates:getDirectoryRoleTemplates", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getDirectoryRoleTemplates.
type GetDirectoryRoleTemplatesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The object IDs of the role templates.
	ObjectIds []string `pulumi:"objectIds"`
	// A list of role templates. Each `roleTemplate` object provides the attributes documented below.
	RoleTemplates []GetDirectoryRoleTemplatesRoleTemplate `pulumi:"roleTemplates"`
}
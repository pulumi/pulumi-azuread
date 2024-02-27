// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuread/sdk/v5/go/azuread/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about an access package catalog role.
//
// ## API Permissions
//
// The following API permissions are required in order to use this data source.
//
// When authenticated with a service principal, this data source requires one of the following application roles: `EntitlementManagement.Read.All` or `Directory.Read.All`
//
// When authenticated with a user principal, this data source does not require any additional roles.
//
// ## Example Usage
func GetAccessPackageCatalogRole(ctx *pulumi.Context, args *GetAccessPackageCatalogRoleArgs, opts ...pulumi.InvokeOption) (*GetAccessPackageCatalogRoleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAccessPackageCatalogRoleResult
	err := ctx.Invoke("azuread:index/getAccessPackageCatalogRole:getAccessPackageCatalogRole", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccessPackageCatalogRole.
type GetAccessPackageCatalogRoleArgs struct {
	// Specifies the display name of the role.
	DisplayName *string `pulumi:"displayName"`
	// Specifies the object ID of the role.
	//
	// > One of `displayName` or `objectId` must be specified.
	ObjectId *string `pulumi:"objectId"`
}

// A collection of values returned by getAccessPackageCatalogRole.
type GetAccessPackageCatalogRoleResult struct {
	// The description of the role.
	Description string `pulumi:"description"`
	// The display name of the role.
	DisplayName string `pulumi:"displayName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The object ID of the role.
	ObjectId string `pulumi:"objectId"`
	// The object ID of the role.
	TemplateId string `pulumi:"templateId"`
}

func GetAccessPackageCatalogRoleOutput(ctx *pulumi.Context, args GetAccessPackageCatalogRoleOutputArgs, opts ...pulumi.InvokeOption) GetAccessPackageCatalogRoleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAccessPackageCatalogRoleResult, error) {
			args := v.(GetAccessPackageCatalogRoleArgs)
			r, err := GetAccessPackageCatalogRole(ctx, &args, opts...)
			var s GetAccessPackageCatalogRoleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAccessPackageCatalogRoleResultOutput)
}

// A collection of arguments for invoking getAccessPackageCatalogRole.
type GetAccessPackageCatalogRoleOutputArgs struct {
	// Specifies the display name of the role.
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
	// Specifies the object ID of the role.
	//
	// > One of `displayName` or `objectId` must be specified.
	ObjectId pulumi.StringPtrInput `pulumi:"objectId"`
}

func (GetAccessPackageCatalogRoleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccessPackageCatalogRoleArgs)(nil)).Elem()
}

// A collection of values returned by getAccessPackageCatalogRole.
type GetAccessPackageCatalogRoleResultOutput struct{ *pulumi.OutputState }

func (GetAccessPackageCatalogRoleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccessPackageCatalogRoleResult)(nil)).Elem()
}

func (o GetAccessPackageCatalogRoleResultOutput) ToGetAccessPackageCatalogRoleResultOutput() GetAccessPackageCatalogRoleResultOutput {
	return o
}

func (o GetAccessPackageCatalogRoleResultOutput) ToGetAccessPackageCatalogRoleResultOutputWithContext(ctx context.Context) GetAccessPackageCatalogRoleResultOutput {
	return o
}

// The description of the role.
func (o GetAccessPackageCatalogRoleResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccessPackageCatalogRoleResult) string { return v.Description }).(pulumi.StringOutput)
}

// The display name of the role.
func (o GetAccessPackageCatalogRoleResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccessPackageCatalogRoleResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAccessPackageCatalogRoleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccessPackageCatalogRoleResult) string { return v.Id }).(pulumi.StringOutput)
}

// The object ID of the role.
func (o GetAccessPackageCatalogRoleResultOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccessPackageCatalogRoleResult) string { return v.ObjectId }).(pulumi.StringOutput)
}

// The object ID of the role.
func (o GetAccessPackageCatalogRoleResultOutput) TemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccessPackageCatalogRoleResult) string { return v.TemplateId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAccessPackageCatalogRoleResultOutput{})
}

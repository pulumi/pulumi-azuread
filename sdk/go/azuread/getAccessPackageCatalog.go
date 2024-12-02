// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuread/sdk/v6/go/azuread/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// i
// Use this resource to retrieve information for an existing access package catalog within Identity Governance in Azure Active Directory.
//
// ## API Permissions
//
// The following API permissions are required in order to use this data source.
//
// When authenticated with a service principal, this data source requires one of the following application roles: `EntitlementManagement.Read.All`, or `EntitlementManagement.ReadWrite.All`.
//
// When authenticated with a user principal, this data source requires one of the following directory roles: `Catalog owner`, `Catalog reader`, `Global Reader`, or `Global Administrator`.
//
// ## Example Usage
//
// *Look up by ID*
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
//			_, err := azuread.LookupAccessPackageCatalog(ctx, &azuread.LookupAccessPackageCatalogArgs{
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
//
// *Look up by DisplayName*
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
//			_, err := azuread.LookupAccessPackageCatalog(ctx, &azuread.LookupAccessPackageCatalogArgs{
//				DisplayName: pulumi.StringRef("My access package Catalog"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupAccessPackageCatalog(ctx *pulumi.Context, args *LookupAccessPackageCatalogArgs, opts ...pulumi.InvokeOption) (*LookupAccessPackageCatalogResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAccessPackageCatalogResult
	err := ctx.Invoke("azuread:index/getAccessPackageCatalog:getAccessPackageCatalog", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccessPackageCatalog.
type LookupAccessPackageCatalogArgs struct {
	// The display name of the access package catalog.
	DisplayName *string `pulumi:"displayName"`
	// The ID of this access package catalog.
	//
	// > One of `displayName` or `objectId` must be specified.
	ObjectId *string `pulumi:"objectId"`
}

// A collection of values returned by getAccessPackageCatalog.
type LookupAccessPackageCatalogResult struct {
	// The description of the access package catalog.
	Description string `pulumi:"description"`
	DisplayName string `pulumi:"displayName"`
	// Whether the access packages in this catalog can be requested by users outside the tenant.
	ExternallyVisible bool `pulumi:"externallyVisible"`
	// The provider-assigned unique ID for this managed resource.
	Id       string `pulumi:"id"`
	ObjectId string `pulumi:"objectId"`
	// Whether the access packages in this catalog are available for management.
	Published bool `pulumi:"published"`
}

func LookupAccessPackageCatalogOutput(ctx *pulumi.Context, args LookupAccessPackageCatalogOutputArgs, opts ...pulumi.InvokeOption) LookupAccessPackageCatalogResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupAccessPackageCatalogResultOutput, error) {
			args := v.(LookupAccessPackageCatalogArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupAccessPackageCatalogResult
			secret, deps, err := ctx.InvokePackageRawWithDeps("azuread:index/getAccessPackageCatalog:getAccessPackageCatalog", args, &rv, "", opts...)
			if err != nil {
				return LookupAccessPackageCatalogResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupAccessPackageCatalogResultOutput)
			output = pulumi.OutputWithDependencies(ctx.Context(), output, deps...).(LookupAccessPackageCatalogResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupAccessPackageCatalogResultOutput), nil
			}
			return output, nil
		}).(LookupAccessPackageCatalogResultOutput)
}

// A collection of arguments for invoking getAccessPackageCatalog.
type LookupAccessPackageCatalogOutputArgs struct {
	// The display name of the access package catalog.
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
	// The ID of this access package catalog.
	//
	// > One of `displayName` or `objectId` must be specified.
	ObjectId pulumi.StringPtrInput `pulumi:"objectId"`
}

func (LookupAccessPackageCatalogOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessPackageCatalogArgs)(nil)).Elem()
}

// A collection of values returned by getAccessPackageCatalog.
type LookupAccessPackageCatalogResultOutput struct{ *pulumi.OutputState }

func (LookupAccessPackageCatalogResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessPackageCatalogResult)(nil)).Elem()
}

func (o LookupAccessPackageCatalogResultOutput) ToLookupAccessPackageCatalogResultOutput() LookupAccessPackageCatalogResultOutput {
	return o
}

func (o LookupAccessPackageCatalogResultOutput) ToLookupAccessPackageCatalogResultOutputWithContext(ctx context.Context) LookupAccessPackageCatalogResultOutput {
	return o
}

// The description of the access package catalog.
func (o LookupAccessPackageCatalogResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessPackageCatalogResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupAccessPackageCatalogResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessPackageCatalogResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Whether the access packages in this catalog can be requested by users outside the tenant.
func (o LookupAccessPackageCatalogResultOutput) ExternallyVisible() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAccessPackageCatalogResult) bool { return v.ExternallyVisible }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAccessPackageCatalogResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessPackageCatalogResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAccessPackageCatalogResultOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessPackageCatalogResult) string { return v.ObjectId }).(pulumi.StringOutput)
}

// Whether the access packages in this catalog are available for management.
func (o LookupAccessPackageCatalogResultOutput) Published() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAccessPackageCatalogResult) bool { return v.Published }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccessPackageCatalogResultOutput{})
}

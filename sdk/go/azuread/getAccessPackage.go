// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information for an existing access package within Identity Governance in Azure Active Directory.
//
// ## API Permissions
//
// The following API permissions are required in order to use this data source.
//
// When authenticated with a service principal, this data source requires one of the following application roles: `EntitlementManagement.Read.All`, or `EntitlementManagement.ReadWrite.All`.
//
// When authenticated with a user principal, this data source requires one of the following directory roles: `Catalog owner`, `Catalog reader`, `Access package manager`, `Global Reader`, or `Global Administrator`.
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
//	"github.com/pulumi/pulumi-azuread/sdk/v5/go/azuread"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := azuread.LookupAccessPackage(ctx, &azuread.LookupAccessPackageArgs{
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
//	"github.com/pulumi/pulumi-azuread/sdk/v5/go/azuread"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := azuread.LookupAccessPackage(ctx, &azuread.LookupAccessPackageArgs{
//				CatalogId:   pulumi.StringRef("00000000-0000-0000-0000-000000000000"),
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
func LookupAccessPackage(ctx *pulumi.Context, args *LookupAccessPackageArgs, opts ...pulumi.InvokeOption) (*LookupAccessPackageResult, error) {
	var rv LookupAccessPackageResult
	err := ctx.Invoke("azuread:index/getAccessPackage:getAccessPackage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccessPackage.
type LookupAccessPackageArgs struct {
	// The ID of the Catalog this access package is in.
	CatalogId *string `pulumi:"catalogId"`
	// The display name of the access package.
	DisplayName *string `pulumi:"displayName"`
	// The ID of this access package.
	ObjectId *string `pulumi:"objectId"`
}

// A collection of values returned by getAccessPackage.
type LookupAccessPackageResult struct {
	CatalogId *string `pulumi:"catalogId"`
	// The description of the access package.
	Description string `pulumi:"description"`
	DisplayName string `pulumi:"displayName"`
	// Whether the access package is hidden from the requestor.
	Hidden bool `pulumi:"hidden"`
	// The provider-assigned unique ID for this managed resource.
	Id       string `pulumi:"id"`
	ObjectId string `pulumi:"objectId"`
}

func LookupAccessPackageOutput(ctx *pulumi.Context, args LookupAccessPackageOutputArgs, opts ...pulumi.InvokeOption) LookupAccessPackageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccessPackageResult, error) {
			args := v.(LookupAccessPackageArgs)
			r, err := LookupAccessPackage(ctx, &args, opts...)
			var s LookupAccessPackageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAccessPackageResultOutput)
}

// A collection of arguments for invoking getAccessPackage.
type LookupAccessPackageOutputArgs struct {
	// The ID of the Catalog this access package is in.
	CatalogId pulumi.StringPtrInput `pulumi:"catalogId"`
	// The display name of the access package.
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
	// The ID of this access package.
	ObjectId pulumi.StringPtrInput `pulumi:"objectId"`
}

func (LookupAccessPackageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessPackageArgs)(nil)).Elem()
}

// A collection of values returned by getAccessPackage.
type LookupAccessPackageResultOutput struct{ *pulumi.OutputState }

func (LookupAccessPackageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessPackageResult)(nil)).Elem()
}

func (o LookupAccessPackageResultOutput) ToLookupAccessPackageResultOutput() LookupAccessPackageResultOutput {
	return o
}

func (o LookupAccessPackageResultOutput) ToLookupAccessPackageResultOutputWithContext(ctx context.Context) LookupAccessPackageResultOutput {
	return o
}

func (o LookupAccessPackageResultOutput) CatalogId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessPackageResult) *string { return v.CatalogId }).(pulumi.StringPtrOutput)
}

// The description of the access package.
func (o LookupAccessPackageResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessPackageResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupAccessPackageResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessPackageResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Whether the access package is hidden from the requestor.
func (o LookupAccessPackageResultOutput) Hidden() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAccessPackageResult) bool { return v.Hidden }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAccessPackageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessPackageResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAccessPackageResultOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessPackageResult) string { return v.ObjectId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccessPackageResultOutput{})
}

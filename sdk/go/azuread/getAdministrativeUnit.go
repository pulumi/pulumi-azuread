// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuread/sdk/v6/go/azuread/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about an adminisrative unit in Azure Active Directory.
//
// ## API Permissions
//
// The following API permissions are required in order to use this data source.
//
// When authenticated with a service principal, this data source requires one of the following application roles: `AdministrativeUnit.Read.All` or `Directory.Read.All`
//
// When authenticated with a user principal, this data source does not require any additional roles.
//
// ## Example Usage
//
// ### By Group Display Name)
//
// *Look up by display name*
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
//			_, err := azuread.LookupAdministrativeUnit(ctx, &azuread.LookupAdministrativeUnitArgs{
//				DisplayName: pulumi.StringRef("Example-AU"),
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
// *Look up by object ID*
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
//			_, err := azuread.LookupAdministrativeUnit(ctx, &azuread.LookupAdministrativeUnitArgs{
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
func LookupAdministrativeUnit(ctx *pulumi.Context, args *LookupAdministrativeUnitArgs, opts ...pulumi.InvokeOption) (*LookupAdministrativeUnitResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAdministrativeUnitResult
	err := ctx.Invoke("azuread:index/getAdministrativeUnit:getAdministrativeUnit", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAdministrativeUnit.
type LookupAdministrativeUnitArgs struct {
	// Specifies the display name of the administrative unit.
	DisplayName *string `pulumi:"displayName"`
	// Specifies the object ID of the administrative unit.
	//
	// > One of `displayName` or `objectId` must be specified.
	ObjectId *string `pulumi:"objectId"`
}

// A collection of values returned by getAdministrativeUnit.
type LookupAdministrativeUnitResult struct {
	// The description of the administrative unit.
	Description string `pulumi:"description"`
	// The display name of the administrative unit.
	DisplayName string `pulumi:"displayName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of object IDs of members who are present in this administrative unit.
	Members []string `pulumi:"members"`
	// The object ID of the administrative unit.
	ObjectId string `pulumi:"objectId"`
	// Whether the administrative unit _and_ its members are hidden or publicly viewable in the directory. One of: `Hiddenmembership` or `Public`.
	Visibility string `pulumi:"visibility"`
}

func LookupAdministrativeUnitOutput(ctx *pulumi.Context, args LookupAdministrativeUnitOutputArgs, opts ...pulumi.InvokeOption) LookupAdministrativeUnitResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupAdministrativeUnitResultOutput, error) {
			args := v.(LookupAdministrativeUnitArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azuread:index/getAdministrativeUnit:getAdministrativeUnit", args, LookupAdministrativeUnitResultOutput{}, options).(LookupAdministrativeUnitResultOutput), nil
		}).(LookupAdministrativeUnitResultOutput)
}

// A collection of arguments for invoking getAdministrativeUnit.
type LookupAdministrativeUnitOutputArgs struct {
	// Specifies the display name of the administrative unit.
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
	// Specifies the object ID of the administrative unit.
	//
	// > One of `displayName` or `objectId` must be specified.
	ObjectId pulumi.StringPtrInput `pulumi:"objectId"`
}

func (LookupAdministrativeUnitOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAdministrativeUnitArgs)(nil)).Elem()
}

// A collection of values returned by getAdministrativeUnit.
type LookupAdministrativeUnitResultOutput struct{ *pulumi.OutputState }

func (LookupAdministrativeUnitResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAdministrativeUnitResult)(nil)).Elem()
}

func (o LookupAdministrativeUnitResultOutput) ToLookupAdministrativeUnitResultOutput() LookupAdministrativeUnitResultOutput {
	return o
}

func (o LookupAdministrativeUnitResultOutput) ToLookupAdministrativeUnitResultOutputWithContext(ctx context.Context) LookupAdministrativeUnitResultOutput {
	return o
}

// The description of the administrative unit.
func (o LookupAdministrativeUnitResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdministrativeUnitResult) string { return v.Description }).(pulumi.StringOutput)
}

// The display name of the administrative unit.
func (o LookupAdministrativeUnitResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdministrativeUnitResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAdministrativeUnitResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdministrativeUnitResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of object IDs of members who are present in this administrative unit.
func (o LookupAdministrativeUnitResultOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAdministrativeUnitResult) []string { return v.Members }).(pulumi.StringArrayOutput)
}

// The object ID of the administrative unit.
func (o LookupAdministrativeUnitResultOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdministrativeUnitResult) string { return v.ObjectId }).(pulumi.StringOutput)
}

// Whether the administrative unit _and_ its members are hidden or publicly viewable in the directory. One of: `Hiddenmembership` or `Public`.
func (o LookupAdministrativeUnitResultOutput) Visibility() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdministrativeUnitResult) string { return v.Visibility }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAdministrativeUnitResultOutput{})
}

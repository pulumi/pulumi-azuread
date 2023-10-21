// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuread/sdk/v5/go/azuread/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets basic information for multiple Azure Active Directory service principals.
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
// *Look up by application display names*
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
//			_, err := azuread.GetServicePrincipals(ctx, &azuread.GetServicePrincipalsArgs{
//				DisplayNames: []string{
//					"example-app",
//					"another-app",
//				},
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
// *Look up by application IDs (client IDs*
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
//			_, err := azuread.GetServicePrincipals(ctx, &azuread.GetServicePrincipalsArgs{
//				ClientIds: []string{
//					"11111111-0000-0000-0000-000000000000",
//					"22222222-0000-0000-0000-000000000000",
//					"33333333-0000-0000-0000-000000000000",
//				},
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
// *Look up by service principal object IDs*
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
//			_, err := azuread.GetServicePrincipals(ctx, &azuread.GetServicePrincipalsArgs{
//				ObjectIds: []string{
//					"00000000-0000-0000-0000-000000000000",
//					"00000000-0000-0000-0000-111111111111",
//					"00000000-0000-0000-0000-222222222222",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetServicePrincipals(ctx *pulumi.Context, args *GetServicePrincipalsArgs, opts ...pulumi.InvokeOption) (*GetServicePrincipalsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetServicePrincipalsResult
	err := ctx.Invoke("azuread:index/getServicePrincipals:getServicePrincipals", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServicePrincipals.
type GetServicePrincipalsArgs struct {
	// A list of client IDs of the applications associated with the service principals.
	//
	// Deprecated: The `application_ids` property has been replaced with the `client_ids` property and will be removed in version 3.0 of the AzureAD provider
	ApplicationIds []string `pulumi:"applicationIds"`
	// A list of client IDs of the applications associated with the service principals.
	ClientIds []string `pulumi:"clientIds"`
	// A list of display names of the applications associated with the service principals.
	DisplayNames []string `pulumi:"displayNames"`
	// Ignore missing service principals and return all service principals that are found. The data source will still fail if no service principals are found. Defaults to false.
	IgnoreMissing *bool `pulumi:"ignoreMissing"`
	// The object IDs of the service principals.
	ObjectIds []string `pulumi:"objectIds"`
	// When `true`, the data source will return all service principals. Cannot be used with `ignoreMissing`. Defaults to false.
	//
	// > Either `returnAll`, or one of `clientIds`, `displayNames` or `objectIds` must be specified. These _may_ be specified as an empty list, in which case no results will be returned.
	ReturnAll *bool `pulumi:"returnAll"`
}

// A collection of values returned by getServicePrincipals.
type GetServicePrincipalsResult struct {
	// A list of client IDs of the applications associated with the service principals.
	//
	// Deprecated: The `application_ids` property has been replaced with the `client_ids` property and will be removed in version 3.0 of the AzureAD provider
	ApplicationIds []string `pulumi:"applicationIds"`
	// The client ID of the application associated with this service principal.
	ClientIds []string `pulumi:"clientIds"`
	// A list of display names of the applications associated with the service principals.
	DisplayNames []string `pulumi:"displayNames"`
	// The provider-assigned unique ID for this managed resource.
	Id            string `pulumi:"id"`
	IgnoreMissing *bool  `pulumi:"ignoreMissing"`
	// The object IDs of the service principals.
	ObjectIds []string `pulumi:"objectIds"`
	ReturnAll *bool    `pulumi:"returnAll"`
	// A list of service principals. Each `servicePrincipal` object provides the attributes documented below.
	ServicePrincipals []GetServicePrincipalsServicePrincipal `pulumi:"servicePrincipals"`
}

func GetServicePrincipalsOutput(ctx *pulumi.Context, args GetServicePrincipalsOutputArgs, opts ...pulumi.InvokeOption) GetServicePrincipalsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetServicePrincipalsResult, error) {
			args := v.(GetServicePrincipalsArgs)
			r, err := GetServicePrincipals(ctx, &args, opts...)
			var s GetServicePrincipalsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetServicePrincipalsResultOutput)
}

// A collection of arguments for invoking getServicePrincipals.
type GetServicePrincipalsOutputArgs struct {
	// A list of client IDs of the applications associated with the service principals.
	//
	// Deprecated: The `application_ids` property has been replaced with the `client_ids` property and will be removed in version 3.0 of the AzureAD provider
	ApplicationIds pulumi.StringArrayInput `pulumi:"applicationIds"`
	// A list of client IDs of the applications associated with the service principals.
	ClientIds pulumi.StringArrayInput `pulumi:"clientIds"`
	// A list of display names of the applications associated with the service principals.
	DisplayNames pulumi.StringArrayInput `pulumi:"displayNames"`
	// Ignore missing service principals and return all service principals that are found. The data source will still fail if no service principals are found. Defaults to false.
	IgnoreMissing pulumi.BoolPtrInput `pulumi:"ignoreMissing"`
	// The object IDs of the service principals.
	ObjectIds pulumi.StringArrayInput `pulumi:"objectIds"`
	// When `true`, the data source will return all service principals. Cannot be used with `ignoreMissing`. Defaults to false.
	//
	// > Either `returnAll`, or one of `clientIds`, `displayNames` or `objectIds` must be specified. These _may_ be specified as an empty list, in which case no results will be returned.
	ReturnAll pulumi.BoolPtrInput `pulumi:"returnAll"`
}

func (GetServicePrincipalsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServicePrincipalsArgs)(nil)).Elem()
}

// A collection of values returned by getServicePrincipals.
type GetServicePrincipalsResultOutput struct{ *pulumi.OutputState }

func (GetServicePrincipalsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServicePrincipalsResult)(nil)).Elem()
}

func (o GetServicePrincipalsResultOutput) ToGetServicePrincipalsResultOutput() GetServicePrincipalsResultOutput {
	return o
}

func (o GetServicePrincipalsResultOutput) ToGetServicePrincipalsResultOutputWithContext(ctx context.Context) GetServicePrincipalsResultOutput {
	return o
}

func (o GetServicePrincipalsResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetServicePrincipalsResult] {
	return pulumix.Output[GetServicePrincipalsResult]{
		OutputState: o.OutputState,
	}
}

// A list of client IDs of the applications associated with the service principals.
//
// Deprecated: The `application_ids` property has been replaced with the `client_ids` property and will be removed in version 3.0 of the AzureAD provider
func (o GetServicePrincipalsResultOutput) ApplicationIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetServicePrincipalsResult) []string { return v.ApplicationIds }).(pulumi.StringArrayOutput)
}

// The client ID of the application associated with this service principal.
func (o GetServicePrincipalsResultOutput) ClientIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetServicePrincipalsResult) []string { return v.ClientIds }).(pulumi.StringArrayOutput)
}

// A list of display names of the applications associated with the service principals.
func (o GetServicePrincipalsResultOutput) DisplayNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetServicePrincipalsResult) []string { return v.DisplayNames }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetServicePrincipalsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetServicePrincipalsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetServicePrincipalsResultOutput) IgnoreMissing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetServicePrincipalsResult) *bool { return v.IgnoreMissing }).(pulumi.BoolPtrOutput)
}

// The object IDs of the service principals.
func (o GetServicePrincipalsResultOutput) ObjectIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetServicePrincipalsResult) []string { return v.ObjectIds }).(pulumi.StringArrayOutput)
}

func (o GetServicePrincipalsResultOutput) ReturnAll() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetServicePrincipalsResult) *bool { return v.ReturnAll }).(pulumi.BoolPtrOutput)
}

// A list of service principals. Each `servicePrincipal` object provides the attributes documented below.
func (o GetServicePrincipalsResultOutput) ServicePrincipals() GetServicePrincipalsServicePrincipalArrayOutput {
	return o.ApplyT(func(v GetServicePrincipalsResult) []GetServicePrincipalsServicePrincipal { return v.ServicePrincipals }).(GetServicePrincipalsServicePrincipalArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetServicePrincipalsResultOutput{})
}

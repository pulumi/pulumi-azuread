// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuread/sdk/v6/go/azuread/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the OData type for a generic directory object having the provided object ID.
//
// ## API Permissions
//
// The following API permissions are required in order to use this data source.
//
// When authenticated with a service principal, this data source requires either `User.Read.All`, `Group.Read.All` or `Directory.Read.All`, depending on the type of object being queried.
//
// When authenticated with a user principal, this data source does not require any additional roles.
//
// ## Example Usage
//
// *Look up and output type of object by ID*
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
//			example, err := azuread.GetDirectoryObject(ctx, &azuread.GetDirectoryObjectArgs{
//				ObjectId: "00000000-0000-0000-0000-000000000000",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("objectType", example.Type)
//			return nil
//		})
//	}
//
// ```
//
// ## Attributes Reference
//
// The following attributes are exported:
//
// *`objectId` - The object ID of the directory object.
// *`type` - The shortened OData type of the directory object. Possible values include: `Group`, `User` or `ServicePrincipal`.
func GetDirectoryObject(ctx *pulumi.Context, args *GetDirectoryObjectArgs, opts ...pulumi.InvokeOption) (*GetDirectoryObjectResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDirectoryObjectResult
	err := ctx.Invoke("azuread:index/getDirectoryObject:getDirectoryObject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDirectoryObject.
type GetDirectoryObjectArgs struct {
	// Specifies the Object ID of the directory object to look up.
	ObjectId string `pulumi:"objectId"`
}

// A collection of values returned by getDirectoryObject.
type GetDirectoryObjectResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id       string `pulumi:"id"`
	ObjectId string `pulumi:"objectId"`
	Type     string `pulumi:"type"`
}

func GetDirectoryObjectOutput(ctx *pulumi.Context, args GetDirectoryObjectOutputArgs, opts ...pulumi.InvokeOption) GetDirectoryObjectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDirectoryObjectResultOutput, error) {
			args := v.(GetDirectoryObjectArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetDirectoryObjectResult
			secret, err := ctx.InvokePackageRaw("azuread:index/getDirectoryObject:getDirectoryObject", args, &rv, "", opts...)
			if err != nil {
				return GetDirectoryObjectResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetDirectoryObjectResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetDirectoryObjectResultOutput), nil
			}
			return output, nil
		}).(GetDirectoryObjectResultOutput)
}

// A collection of arguments for invoking getDirectoryObject.
type GetDirectoryObjectOutputArgs struct {
	// Specifies the Object ID of the directory object to look up.
	ObjectId pulumi.StringInput `pulumi:"objectId"`
}

func (GetDirectoryObjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDirectoryObjectArgs)(nil)).Elem()
}

// A collection of values returned by getDirectoryObject.
type GetDirectoryObjectResultOutput struct{ *pulumi.OutputState }

func (GetDirectoryObjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDirectoryObjectResult)(nil)).Elem()
}

func (o GetDirectoryObjectResultOutput) ToGetDirectoryObjectResultOutput() GetDirectoryObjectResultOutput {
	return o
}

func (o GetDirectoryObjectResultOutput) ToGetDirectoryObjectResultOutputWithContext(ctx context.Context) GetDirectoryObjectResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetDirectoryObjectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDirectoryObjectResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDirectoryObjectResultOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDirectoryObjectResult) string { return v.ObjectId }).(pulumi.StringOutput)
}

func (o GetDirectoryObjectResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetDirectoryObjectResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDirectoryObjectResultOutput{})
}

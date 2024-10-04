// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuread/sdk/v6/go/azuread/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access the configuration of the AzureAD provider.
//
// ## API Permissions
//
// No additional roles are required to use this data source.
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
//			current, err := azuread.GetClientConfig(ctx, map[string]interface{}{}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("objectId", current.ObjectId)
//			return nil
//		})
//	}
//
// ```
func GetClientConfig(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetClientConfigResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetClientConfigResult
	err := ctx.Invoke("azuread:index/getClientConfig:getClientConfig", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getClientConfig.
type GetClientConfigResult struct {
	// The client ID (application ID) linked to the authenticated principal, or the application used for delegated authentication.
	ClientId string `pulumi:"clientId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The object ID of the authenticated principal.
	ObjectId string `pulumi:"objectId"`
	// The tenant ID of the authenticated principal.
	TenantId string `pulumi:"tenantId"`
}

func GetClientConfigOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetClientConfigResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetClientConfigResultOutput, error) {
		opts = internal.PkgInvokeDefaultOpts(opts)
		var rv GetClientConfigResult
		secret, err := ctx.InvokePackageRaw("azuread:index/getClientConfig:getClientConfig", nil, &rv, "", opts...)
		if err != nil {
			return GetClientConfigResultOutput{}, err
		}

		output := pulumi.ToOutput(rv).(GetClientConfigResultOutput)
		if secret {
			return pulumi.ToSecret(output).(GetClientConfigResultOutput), nil
		}
		return output, nil
	}).(GetClientConfigResultOutput)
}

// A collection of values returned by getClientConfig.
type GetClientConfigResultOutput struct{ *pulumi.OutputState }

func (GetClientConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClientConfigResult)(nil)).Elem()
}

func (o GetClientConfigResultOutput) ToGetClientConfigResultOutput() GetClientConfigResultOutput {
	return o
}

func (o GetClientConfigResultOutput) ToGetClientConfigResultOutputWithContext(ctx context.Context) GetClientConfigResultOutput {
	return o
}

// The client ID (application ID) linked to the authenticated principal, or the application used for delegated authentication.
func (o GetClientConfigResultOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientConfigResult) string { return v.ClientId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetClientConfigResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientConfigResult) string { return v.Id }).(pulumi.StringOutput)
}

// The object ID of the authenticated principal.
func (o GetClientConfigResultOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientConfigResult) string { return v.ObjectId }).(pulumi.StringOutput)
}

// The tenant ID of the authenticated principal.
func (o GetClientConfigResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientConfigResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetClientConfigResultOutput{})
}

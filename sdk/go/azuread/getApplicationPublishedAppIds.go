// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuread/sdk/v6/go/azuread/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// *Listing well-known application IDs*
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
//			wellKnown, err := azuread.GetApplicationPublishedAppIds(ctx, map[string]interface{}{}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("publishedAppIds", wellKnown.Result)
//			return nil
//		})
//	}
//
// ```
//
// *Granting access to an application*
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
//			wellKnown, err := azuread.GetApplicationPublishedAppIds(ctx, map[string]interface{}{}, nil)
//			if err != nil {
//				return err
//			}
//			msgraph, err := azuread.NewServicePrincipal(ctx, "msgraph", &azuread.ServicePrincipalArgs{
//				ClientId:    pulumi.String(wellKnown.Result.MicrosoftGraph),
//				UseExisting: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuread.NewApplication(ctx, "example", &azuread.ApplicationArgs{
//				DisplayName: pulumi.String("example"),
//				RequiredResourceAccesses: azuread.ApplicationRequiredResourceAccessArray{
//					&azuread.ApplicationRequiredResourceAccessArgs{
//						ResourceAppId: pulumi.String(wellKnown.Result.MicrosoftGraph),
//						ResourceAccesses: azuread.ApplicationRequiredResourceAccessResourceAccessArray{
//							&azuread.ApplicationRequiredResourceAccessResourceAccessArgs{
//								Id: msgraph.AppRoleIds.ApplyT(func(appRoleIds map[string]string) (string, error) {
//									return appRoleIds.User.Read.All, nil
//								}).(pulumi.StringOutput),
//								Type: pulumi.String("Role"),
//							},
//							&azuread.ApplicationRequiredResourceAccessResourceAccessArgs{
//								Id: msgraph.Oauth2PermissionScopeIds.ApplyT(func(oauth2PermissionScopeIds map[string]string) (string, error) {
//									return oauth2PermissionScopeIds.User.ReadWrite, nil
//								}).(pulumi.StringOutput),
//								Type: pulumi.String("Scope"),
//							},
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetApplicationPublishedAppIds(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetApplicationPublishedAppIdsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetApplicationPublishedAppIdsResult
	err := ctx.Invoke("azuread:index/getApplicationPublishedAppIds:getApplicationPublishedAppIds", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getApplicationPublishedAppIds.
type GetApplicationPublishedAppIdsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A map of application names to application IDs.
	Result map[string]string `pulumi:"result"`
}

func GetApplicationPublishedAppIdsOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetApplicationPublishedAppIdsResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetApplicationPublishedAppIdsResultOutput, error) {
		options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
		return ctx.InvokeOutput("azuread:index/getApplicationPublishedAppIds:getApplicationPublishedAppIds", nil, GetApplicationPublishedAppIdsResultOutput{}, options).(GetApplicationPublishedAppIdsResultOutput), nil
	}).(GetApplicationPublishedAppIdsResultOutput)
}

// A collection of values returned by getApplicationPublishedAppIds.
type GetApplicationPublishedAppIdsResultOutput struct{ *pulumi.OutputState }

func (GetApplicationPublishedAppIdsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApplicationPublishedAppIdsResult)(nil)).Elem()
}

func (o GetApplicationPublishedAppIdsResultOutput) ToGetApplicationPublishedAppIdsResultOutput() GetApplicationPublishedAppIdsResultOutput {
	return o
}

func (o GetApplicationPublishedAppIdsResultOutput) ToGetApplicationPublishedAppIdsResultOutputWithContext(ctx context.Context) GetApplicationPublishedAppIdsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetApplicationPublishedAppIdsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetApplicationPublishedAppIdsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A map of application names to application IDs.
func (o GetApplicationPublishedAppIdsResultOutput) Result() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetApplicationPublishedAppIdsResult) map[string]string { return v.Result }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetApplicationPublishedAppIdsResultOutput{})
}

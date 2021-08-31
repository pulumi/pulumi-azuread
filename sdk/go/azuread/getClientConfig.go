// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
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
// 	"github.com/pulumi/pulumi-azuread/sdk/v4/go/azuread"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		current, err := azuread.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("objectId", current.ObjectId)
// 		return nil
// 	})
// }
// ```
func GetClientConfig(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetClientConfigResult, error) {
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

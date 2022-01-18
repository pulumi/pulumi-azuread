// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets Object IDs or Display Names for multiple Azure Active Directory groups.
//
// ## API Permissions
//
// The following API permissions are required in order to use this data source.
//
// When authenticated with a service principal, this data source requires one of the following application roles: `Group.Read.All` or `Directory.Read.All`
//
// When authenticated with a user principal, this data source does not require any additional roles.
//
// ## Example Usage
//
// *Look up by group name*
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azuread/sdk/v5/go/azuread"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := azuread.GetGroups(ctx, &GetGroupsArgs{
// 			DisplayNames: []string{
// 				"group-a",
// 				"group-b",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// *Look up by display name prefix*
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azuread/sdk/v5/go/azuread"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "sales-"
// 		_, err := azuread.GetGroups(ctx, &GetGroupsArgs{
// 			DisplayNamePrefix: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// *Look up all groups*
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azuread/sdk/v5/go/azuread"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := true
// 		_, err := azuread.GetGroups(ctx, &GetGroupsArgs{
// 			ReturnAll: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// *Look up all mail-enabled groups*
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azuread/sdk/v5/go/azuread"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := true
// 		opt1 := true
// 		_, err := azuread.GetGroups(ctx, &GetGroupsArgs{
// 			MailEnabled: &opt0,
// 			ReturnAll:   &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// *Look up all security-enabled groups that are not mail-enabled*
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azuread/sdk/v5/go/azuread"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := false
// 		opt1 := true
// 		opt2 := true
// 		_, err := azuread.GetGroups(ctx, &GetGroupsArgs{
// 			MailEnabled:     &opt0,
// 			ReturnAll:       &opt1,
// 			SecurityEnabled: &opt2,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetGroups(ctx *pulumi.Context, args *GetGroupsArgs, opts ...pulumi.InvokeOption) (*GetGroupsResult, error) {
	var rv GetGroupsResult
	err := ctx.Invoke("azuread:index/getGroups:getGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroups.
type GetGroupsArgs struct {
	// A common display name prefix to match when returning groups.
	DisplayNamePrefix *string `pulumi:"displayNamePrefix"`
	// The display names of the groups.
	DisplayNames []string `pulumi:"displayNames"`
	// Whether the returned groups should be mail-enabled. By itself this does not exclude security-enabled groups. Setting this to `true` ensures all groups are mail-enabled, and setting to `false` ensures that all groups are _not_ mail-enabled. To ignore this filter, omit the property or set it to null. Cannot be specified together with `objectIds`.
	MailEnabled *bool `pulumi:"mailEnabled"`
	// The object IDs of the groups.
	ObjectIds []string `pulumi:"objectIds"`
	// A flag to denote if all groups should be fetched and returned.
	ReturnAll *bool `pulumi:"returnAll"`
	// Whether the returned groups should be security-enabled. By itself this does not exclude mail-enabled groups. Setting this to `true` ensures all groups are security-enabled, and setting to `false` ensures that all groups are _not_ security-enabled. To ignore this filter, omit the property or set it to null. Cannot be specified together with `objectIds`.
	SecurityEnabled *bool `pulumi:"securityEnabled"`
}

// A collection of values returned by getGroups.
type GetGroupsResult struct {
	DisplayNamePrefix string `pulumi:"displayNamePrefix"`
	// The display names of the groups.
	DisplayNames []string `pulumi:"displayNames"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	MailEnabled bool   `pulumi:"mailEnabled"`
	// The object IDs of the groups.
	ObjectIds       []string `pulumi:"objectIds"`
	ReturnAll       *bool    `pulumi:"returnAll"`
	SecurityEnabled bool     `pulumi:"securityEnabled"`
}

func GetGroupsOutput(ctx *pulumi.Context, args GetGroupsOutputArgs, opts ...pulumi.InvokeOption) GetGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetGroupsResult, error) {
			args := v.(GetGroupsArgs)
			r, err := GetGroups(ctx, &args, opts...)
			return *r, err
		}).(GetGroupsResultOutput)
}

// A collection of arguments for invoking getGroups.
type GetGroupsOutputArgs struct {
	// A common display name prefix to match when returning groups.
	DisplayNamePrefix pulumi.StringPtrInput `pulumi:"displayNamePrefix"`
	// The display names of the groups.
	DisplayNames pulumi.StringArrayInput `pulumi:"displayNames"`
	// Whether the returned groups should be mail-enabled. By itself this does not exclude security-enabled groups. Setting this to `true` ensures all groups are mail-enabled, and setting to `false` ensures that all groups are _not_ mail-enabled. To ignore this filter, omit the property or set it to null. Cannot be specified together with `objectIds`.
	MailEnabled pulumi.BoolPtrInput `pulumi:"mailEnabled"`
	// The object IDs of the groups.
	ObjectIds pulumi.StringArrayInput `pulumi:"objectIds"`
	// A flag to denote if all groups should be fetched and returned.
	ReturnAll pulumi.BoolPtrInput `pulumi:"returnAll"`
	// Whether the returned groups should be security-enabled. By itself this does not exclude mail-enabled groups. Setting this to `true` ensures all groups are security-enabled, and setting to `false` ensures that all groups are _not_ security-enabled. To ignore this filter, omit the property or set it to null. Cannot be specified together with `objectIds`.
	SecurityEnabled pulumi.BoolPtrInput `pulumi:"securityEnabled"`
}

func (GetGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupsArgs)(nil)).Elem()
}

// A collection of values returned by getGroups.
type GetGroupsResultOutput struct{ *pulumi.OutputState }

func (GetGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupsResult)(nil)).Elem()
}

func (o GetGroupsResultOutput) ToGetGroupsResultOutput() GetGroupsResultOutput {
	return o
}

func (o GetGroupsResultOutput) ToGetGroupsResultOutputWithContext(ctx context.Context) GetGroupsResultOutput {
	return o
}

func (o GetGroupsResultOutput) DisplayNamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupsResult) string { return v.DisplayNamePrefix }).(pulumi.StringOutput)
}

// The display names of the groups.
func (o GetGroupsResultOutput) DisplayNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetGroupsResult) []string { return v.DisplayNames }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetGroupsResultOutput) MailEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetGroupsResult) bool { return v.MailEnabled }).(pulumi.BoolOutput)
}

// The object IDs of the groups.
func (o GetGroupsResultOutput) ObjectIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetGroupsResult) []string { return v.ObjectIds }).(pulumi.StringArrayOutput)
}

func (o GetGroupsResultOutput) ReturnAll() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetGroupsResult) *bool { return v.ReturnAll }).(pulumi.BoolPtrOutput)
}

func (o GetGroupsResultOutput) SecurityEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetGroupsResult) bool { return v.SecurityEnabled }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGroupsResultOutput{})
}

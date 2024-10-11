// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azuread/sdk/v6/go/azuread/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Claims Mapping Policy within Azure Active Directory.
//
// ## API Permissions
//
// The following API permissions are required in order to use this resource.
//
// When authenticated with a service principal, this resource requires the following application roles: `Policy.ReadWrite.ApplicationConfiguration` and `Policy.Read.All`
//
// When authenticated with a user principal, this resource requires one of the following directory roles: `Application Administrator` or `Global Administrator`
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/pulumi/pulumi-azuread/sdk/v6/go/azuread"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"claimsMappingPolicy": map[string]interface{}{
//					"claimsSchema": []map[string]interface{}{
//						map[string]interface{}{
//							"ID":            "employeeid",
//							"jwtClaimType":  "name",
//							"samlClaimType": "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/name",
//							"source":        "user",
//						},
//						map[string]interface{}{
//							"ID":            "tenantcountry",
//							"jwtClaimType":  "country",
//							"samlClaimType": "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/country",
//							"source":        "company",
//						},
//					},
//					"includeBasicClaimSet": "true",
//					"version":              1,
//				},
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			_, err = azuread.NewClaimsMappingPolicy(ctx, "my_policy", &azuread.ClaimsMappingPolicyArgs{
//				Definitions: pulumi.StringArray{
//					pulumi.String(json0),
//				},
//				DisplayName: pulumi.String("My Policy"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Claims Mapping Policy can be imported using the `id`, e.g.
//
// ```sh
// $ pulumi import azuread:index/claimsMappingPolicy:ClaimsMappingPolicy my_policy 00000000-0000-0000-0000-000000000000
// ```
type ClaimsMappingPolicy struct {
	pulumi.CustomResourceState

	// The claims mapping policy. This is a JSON formatted string, for which the `jsonencode()` function can be used.
	Definitions pulumi.StringArrayOutput `pulumi:"definitions"`
	// The display name for this Claims Mapping Policy.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
}

// NewClaimsMappingPolicy registers a new resource with the given unique name, arguments, and options.
func NewClaimsMappingPolicy(ctx *pulumi.Context,
	name string, args *ClaimsMappingPolicyArgs, opts ...pulumi.ResourceOption) (*ClaimsMappingPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Definitions == nil {
		return nil, errors.New("invalid value for required argument 'Definitions'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ClaimsMappingPolicy
	err := ctx.RegisterResource("azuread:index/claimsMappingPolicy:ClaimsMappingPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClaimsMappingPolicy gets an existing ClaimsMappingPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClaimsMappingPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClaimsMappingPolicyState, opts ...pulumi.ResourceOption) (*ClaimsMappingPolicy, error) {
	var resource ClaimsMappingPolicy
	err := ctx.ReadResource("azuread:index/claimsMappingPolicy:ClaimsMappingPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClaimsMappingPolicy resources.
type claimsMappingPolicyState struct {
	// The claims mapping policy. This is a JSON formatted string, for which the `jsonencode()` function can be used.
	Definitions []string `pulumi:"definitions"`
	// The display name for this Claims Mapping Policy.
	DisplayName *string `pulumi:"displayName"`
}

type ClaimsMappingPolicyState struct {
	// The claims mapping policy. This is a JSON formatted string, for which the `jsonencode()` function can be used.
	Definitions pulumi.StringArrayInput
	// The display name for this Claims Mapping Policy.
	DisplayName pulumi.StringPtrInput
}

func (ClaimsMappingPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*claimsMappingPolicyState)(nil)).Elem()
}

type claimsMappingPolicyArgs struct {
	// The claims mapping policy. This is a JSON formatted string, for which the `jsonencode()` function can be used.
	Definitions []string `pulumi:"definitions"`
	// The display name for this Claims Mapping Policy.
	DisplayName string `pulumi:"displayName"`
}

// The set of arguments for constructing a ClaimsMappingPolicy resource.
type ClaimsMappingPolicyArgs struct {
	// The claims mapping policy. This is a JSON formatted string, for which the `jsonencode()` function can be used.
	Definitions pulumi.StringArrayInput
	// The display name for this Claims Mapping Policy.
	DisplayName pulumi.StringInput
}

func (ClaimsMappingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*claimsMappingPolicyArgs)(nil)).Elem()
}

type ClaimsMappingPolicyInput interface {
	pulumi.Input

	ToClaimsMappingPolicyOutput() ClaimsMappingPolicyOutput
	ToClaimsMappingPolicyOutputWithContext(ctx context.Context) ClaimsMappingPolicyOutput
}

func (*ClaimsMappingPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ClaimsMappingPolicy)(nil)).Elem()
}

func (i *ClaimsMappingPolicy) ToClaimsMappingPolicyOutput() ClaimsMappingPolicyOutput {
	return i.ToClaimsMappingPolicyOutputWithContext(context.Background())
}

func (i *ClaimsMappingPolicy) ToClaimsMappingPolicyOutputWithContext(ctx context.Context) ClaimsMappingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClaimsMappingPolicyOutput)
}

// ClaimsMappingPolicyArrayInput is an input type that accepts ClaimsMappingPolicyArray and ClaimsMappingPolicyArrayOutput values.
// You can construct a concrete instance of `ClaimsMappingPolicyArrayInput` via:
//
//	ClaimsMappingPolicyArray{ ClaimsMappingPolicyArgs{...} }
type ClaimsMappingPolicyArrayInput interface {
	pulumi.Input

	ToClaimsMappingPolicyArrayOutput() ClaimsMappingPolicyArrayOutput
	ToClaimsMappingPolicyArrayOutputWithContext(context.Context) ClaimsMappingPolicyArrayOutput
}

type ClaimsMappingPolicyArray []ClaimsMappingPolicyInput

func (ClaimsMappingPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClaimsMappingPolicy)(nil)).Elem()
}

func (i ClaimsMappingPolicyArray) ToClaimsMappingPolicyArrayOutput() ClaimsMappingPolicyArrayOutput {
	return i.ToClaimsMappingPolicyArrayOutputWithContext(context.Background())
}

func (i ClaimsMappingPolicyArray) ToClaimsMappingPolicyArrayOutputWithContext(ctx context.Context) ClaimsMappingPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClaimsMappingPolicyArrayOutput)
}

// ClaimsMappingPolicyMapInput is an input type that accepts ClaimsMappingPolicyMap and ClaimsMappingPolicyMapOutput values.
// You can construct a concrete instance of `ClaimsMappingPolicyMapInput` via:
//
//	ClaimsMappingPolicyMap{ "key": ClaimsMappingPolicyArgs{...} }
type ClaimsMappingPolicyMapInput interface {
	pulumi.Input

	ToClaimsMappingPolicyMapOutput() ClaimsMappingPolicyMapOutput
	ToClaimsMappingPolicyMapOutputWithContext(context.Context) ClaimsMappingPolicyMapOutput
}

type ClaimsMappingPolicyMap map[string]ClaimsMappingPolicyInput

func (ClaimsMappingPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClaimsMappingPolicy)(nil)).Elem()
}

func (i ClaimsMappingPolicyMap) ToClaimsMappingPolicyMapOutput() ClaimsMappingPolicyMapOutput {
	return i.ToClaimsMappingPolicyMapOutputWithContext(context.Background())
}

func (i ClaimsMappingPolicyMap) ToClaimsMappingPolicyMapOutputWithContext(ctx context.Context) ClaimsMappingPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClaimsMappingPolicyMapOutput)
}

type ClaimsMappingPolicyOutput struct{ *pulumi.OutputState }

func (ClaimsMappingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClaimsMappingPolicy)(nil)).Elem()
}

func (o ClaimsMappingPolicyOutput) ToClaimsMappingPolicyOutput() ClaimsMappingPolicyOutput {
	return o
}

func (o ClaimsMappingPolicyOutput) ToClaimsMappingPolicyOutputWithContext(ctx context.Context) ClaimsMappingPolicyOutput {
	return o
}

// The claims mapping policy. This is a JSON formatted string, for which the `jsonencode()` function can be used.
func (o ClaimsMappingPolicyOutput) Definitions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ClaimsMappingPolicy) pulumi.StringArrayOutput { return v.Definitions }).(pulumi.StringArrayOutput)
}

// The display name for this Claims Mapping Policy.
func (o ClaimsMappingPolicyOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *ClaimsMappingPolicy) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

type ClaimsMappingPolicyArrayOutput struct{ *pulumi.OutputState }

func (ClaimsMappingPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClaimsMappingPolicy)(nil)).Elem()
}

func (o ClaimsMappingPolicyArrayOutput) ToClaimsMappingPolicyArrayOutput() ClaimsMappingPolicyArrayOutput {
	return o
}

func (o ClaimsMappingPolicyArrayOutput) ToClaimsMappingPolicyArrayOutputWithContext(ctx context.Context) ClaimsMappingPolicyArrayOutput {
	return o
}

func (o ClaimsMappingPolicyArrayOutput) Index(i pulumi.IntInput) ClaimsMappingPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClaimsMappingPolicy {
		return vs[0].([]*ClaimsMappingPolicy)[vs[1].(int)]
	}).(ClaimsMappingPolicyOutput)
}

type ClaimsMappingPolicyMapOutput struct{ *pulumi.OutputState }

func (ClaimsMappingPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClaimsMappingPolicy)(nil)).Elem()
}

func (o ClaimsMappingPolicyMapOutput) ToClaimsMappingPolicyMapOutput() ClaimsMappingPolicyMapOutput {
	return o
}

func (o ClaimsMappingPolicyMapOutput) ToClaimsMappingPolicyMapOutputWithContext(ctx context.Context) ClaimsMappingPolicyMapOutput {
	return o
}

func (o ClaimsMappingPolicyMapOutput) MapIndex(k pulumi.StringInput) ClaimsMappingPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClaimsMappingPolicy {
		return vs[0].(map[string]*ClaimsMappingPolicy)[vs[1].(string)]
	}).(ClaimsMappingPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClaimsMappingPolicyInput)(nil)).Elem(), &ClaimsMappingPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClaimsMappingPolicyArrayInput)(nil)).Elem(), ClaimsMappingPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClaimsMappingPolicyMapInput)(nil)).Elem(), ClaimsMappingPolicyMap{})
	pulumi.RegisterOutputType(ClaimsMappingPolicyOutput{})
	pulumi.RegisterOutputType(ClaimsMappingPolicyArrayOutput{})
	pulumi.RegisterOutputType(ClaimsMappingPolicyMapOutput{})
}

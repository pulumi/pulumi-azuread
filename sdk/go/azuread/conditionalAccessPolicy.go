// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azuread/sdk/v5/go/azuread/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ### All users except guests or external users
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
//			_, err := azuread.NewConditionalAccessPolicy(ctx, "example", &azuread.ConditionalAccessPolicyArgs{
//				DisplayName: pulumi.String("example policy"),
//				State:       pulumi.String("disabled"),
//				Conditions: &azuread.ConditionalAccessPolicyConditionsArgs{
//					ClientAppTypes: pulumi.StringArray{
//						pulumi.String("all"),
//					},
//					SignInRiskLevels: pulumi.StringArray{
//						pulumi.String("medium"),
//					},
//					UserRiskLevels: pulumi.StringArray{
//						pulumi.String("medium"),
//					},
//					Applications: &azuread.ConditionalAccessPolicyConditionsApplicationsArgs{
//						IncludedApplications: pulumi.StringArray{
//							pulumi.String("All"),
//						},
//						ExcludedApplications: pulumi.StringArray{},
//					},
//					Devices: &azuread.ConditionalAccessPolicyConditionsDevicesArgs{
//						Filter: &azuread.ConditionalAccessPolicyConditionsDevicesFilterArgs{
//							Mode: pulumi.String("exclude"),
//							Rule: pulumi.String("device.operatingSystem eq \"Doors\""),
//						},
//					},
//					Locations: &azuread.ConditionalAccessPolicyConditionsLocationsArgs{
//						IncludedLocations: pulumi.StringArray{
//							pulumi.String("All"),
//						},
//						ExcludedLocations: pulumi.StringArray{
//							pulumi.String("AllTrusted"),
//						},
//					},
//					Platforms: &azuread.ConditionalAccessPolicyConditionsPlatformsArgs{
//						IncludedPlatforms: pulumi.StringArray{
//							pulumi.String("android"),
//						},
//						ExcludedPlatforms: pulumi.StringArray{
//							pulumi.String("iOS"),
//						},
//					},
//					Users: &azuread.ConditionalAccessPolicyConditionsUsersArgs{
//						IncludedUsers: pulumi.StringArray{
//							pulumi.String("All"),
//						},
//						ExcludedUsers: pulumi.StringArray{
//							pulumi.String("GuestsOrExternalUsers"),
//						},
//					},
//				},
//				GrantControls: &azuread.ConditionalAccessPolicyGrantControlsArgs{
//					Operator: pulumi.String("OR"),
//					BuiltInControls: pulumi.StringArray{
//						pulumi.String("mfa"),
//					},
//				},
//				SessionControls: &azuread.ConditionalAccessPolicySessionControlsArgs{
//					ApplicationEnforcedRestrictionsEnabled: pulumi.Bool(true),
//					DisableResilienceDefaults:              pulumi.Bool(false),
//					SignInFrequency:                        pulumi.Int(10),
//					SignInFrequencyPeriod:                  pulumi.String("hours"),
//					CloudAppSecurityPolicy:                 pulumi.String("monitorOnly"),
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
//
// ### Included client applications / service principals
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
//			current, err := azuread.GetClientConfig(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = azuread.NewConditionalAccessPolicy(ctx, "example", &azuread.ConditionalAccessPolicyArgs{
//				DisplayName: pulumi.String("example policy"),
//				State:       pulumi.String("disabled"),
//				Conditions: &azuread.ConditionalAccessPolicyConditionsArgs{
//					ClientAppTypes: pulumi.StringArray{
//						pulumi.String("all"),
//					},
//					Applications: &azuread.ConditionalAccessPolicyConditionsApplicationsArgs{
//						IncludedApplications: pulumi.StringArray{
//							pulumi.String("All"),
//						},
//					},
//					ClientApplications: &azuread.ConditionalAccessPolicyConditionsClientApplicationsArgs{
//						IncludedServicePrincipals: pulumi.StringArray{
//							pulumi.String(current.ObjectId),
//						},
//						ExcludedServicePrincipals: pulumi.StringArray{},
//					},
//					Users: &azuread.ConditionalAccessPolicyConditionsUsersArgs{
//						IncludedUsers: pulumi.StringArray{
//							pulumi.String("None"),
//						},
//					},
//				},
//				GrantControls: &azuread.ConditionalAccessPolicyGrantControlsArgs{
//					Operator: pulumi.String("OR"),
//					BuiltInControls: pulumi.StringArray{
//						pulumi.String("block"),
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
//
// ### Excluded client applications / service principals
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
//			current, err := azuread.GetClientConfig(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = azuread.NewConditionalAccessPolicy(ctx, "example", &azuread.ConditionalAccessPolicyArgs{
//				DisplayName: pulumi.String("example policy"),
//				State:       pulumi.String("disabled"),
//				Conditions: &azuread.ConditionalAccessPolicyConditionsArgs{
//					ClientAppTypes: pulumi.StringArray{
//						pulumi.String("all"),
//					},
//					Applications: &azuread.ConditionalAccessPolicyConditionsApplicationsArgs{
//						IncludedApplications: pulumi.StringArray{
//							pulumi.String("All"),
//						},
//					},
//					ClientApplications: &azuread.ConditionalAccessPolicyConditionsClientApplicationsArgs{
//						IncludedServicePrincipals: pulumi.StringArray{
//							pulumi.String("ServicePrincipalsInMyTenant"),
//						},
//						ExcludedServicePrincipals: pulumi.StringArray{
//							pulumi.String(current.ObjectId),
//						},
//					},
//					Users: &azuread.ConditionalAccessPolicyConditionsUsersArgs{
//						IncludedUsers: pulumi.StringArray{
//							pulumi.String("None"),
//						},
//					},
//				},
//				GrantControls: &azuread.ConditionalAccessPolicyGrantControlsArgs{
//					Operator: pulumi.String("OR"),
//					BuiltInControls: pulumi.StringArray{
//						pulumi.String("block"),
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
//
// ## Import
//
// Conditional Access Policies can be imported using the `id`, e.g.
//
// ```sh
// $ pulumi import azuread:index/conditionalAccessPolicy:ConditionalAccessPolicy my_location 00000000-0000-0000-0000-000000000000
// ```
type ConditionalAccessPolicy struct {
	pulumi.CustomResourceState

	// A `conditions` block as documented below, which specifies the rules that must be met for the policy to apply.
	Conditions ConditionalAccessPolicyConditionsOutput `pulumi:"conditions"`
	// The friendly name for this Conditional Access Policy.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// A `grantControls` block as documented below, which specifies the grant controls that must be fulfilled to pass the policy.
	GrantControls ConditionalAccessPolicyGrantControlsPtrOutput `pulumi:"grantControls"`
	// A `sessionControls` block as documented below, which specifies the session controls that are enforced after sign-in.
	//
	// > Note: At least one of `grantControls` and/or `sessionControls` blocks must be specified.
	SessionControls ConditionalAccessPolicySessionControlsPtrOutput `pulumi:"sessionControls"`
	// Specifies the state of the policy object. Possible values are: `enabled`, `disabled` and `enabledForReportingButNotEnforced`
	State pulumi.StringOutput `pulumi:"state"`
}

// NewConditionalAccessPolicy registers a new resource with the given unique name, arguments, and options.
func NewConditionalAccessPolicy(ctx *pulumi.Context,
	name string, args *ConditionalAccessPolicyArgs, opts ...pulumi.ResourceOption) (*ConditionalAccessPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Conditions == nil {
		return nil, errors.New("invalid value for required argument 'Conditions'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.State == nil {
		return nil, errors.New("invalid value for required argument 'State'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ConditionalAccessPolicy
	err := ctx.RegisterResource("azuread:index/conditionalAccessPolicy:ConditionalAccessPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConditionalAccessPolicy gets an existing ConditionalAccessPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConditionalAccessPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConditionalAccessPolicyState, opts ...pulumi.ResourceOption) (*ConditionalAccessPolicy, error) {
	var resource ConditionalAccessPolicy
	err := ctx.ReadResource("azuread:index/conditionalAccessPolicy:ConditionalAccessPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConditionalAccessPolicy resources.
type conditionalAccessPolicyState struct {
	// A `conditions` block as documented below, which specifies the rules that must be met for the policy to apply.
	Conditions *ConditionalAccessPolicyConditions `pulumi:"conditions"`
	// The friendly name for this Conditional Access Policy.
	DisplayName *string `pulumi:"displayName"`
	// A `grantControls` block as documented below, which specifies the grant controls that must be fulfilled to pass the policy.
	GrantControls *ConditionalAccessPolicyGrantControls `pulumi:"grantControls"`
	// A `sessionControls` block as documented below, which specifies the session controls that are enforced after sign-in.
	//
	// > Note: At least one of `grantControls` and/or `sessionControls` blocks must be specified.
	SessionControls *ConditionalAccessPolicySessionControls `pulumi:"sessionControls"`
	// Specifies the state of the policy object. Possible values are: `enabled`, `disabled` and `enabledForReportingButNotEnforced`
	State *string `pulumi:"state"`
}

type ConditionalAccessPolicyState struct {
	// A `conditions` block as documented below, which specifies the rules that must be met for the policy to apply.
	Conditions ConditionalAccessPolicyConditionsPtrInput
	// The friendly name for this Conditional Access Policy.
	DisplayName pulumi.StringPtrInput
	// A `grantControls` block as documented below, which specifies the grant controls that must be fulfilled to pass the policy.
	GrantControls ConditionalAccessPolicyGrantControlsPtrInput
	// A `sessionControls` block as documented below, which specifies the session controls that are enforced after sign-in.
	//
	// > Note: At least one of `grantControls` and/or `sessionControls` blocks must be specified.
	SessionControls ConditionalAccessPolicySessionControlsPtrInput
	// Specifies the state of the policy object. Possible values are: `enabled`, `disabled` and `enabledForReportingButNotEnforced`
	State pulumi.StringPtrInput
}

func (ConditionalAccessPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*conditionalAccessPolicyState)(nil)).Elem()
}

type conditionalAccessPolicyArgs struct {
	// A `conditions` block as documented below, which specifies the rules that must be met for the policy to apply.
	Conditions ConditionalAccessPolicyConditions `pulumi:"conditions"`
	// The friendly name for this Conditional Access Policy.
	DisplayName string `pulumi:"displayName"`
	// A `grantControls` block as documented below, which specifies the grant controls that must be fulfilled to pass the policy.
	GrantControls *ConditionalAccessPolicyGrantControls `pulumi:"grantControls"`
	// A `sessionControls` block as documented below, which specifies the session controls that are enforced after sign-in.
	//
	// > Note: At least one of `grantControls` and/or `sessionControls` blocks must be specified.
	SessionControls *ConditionalAccessPolicySessionControls `pulumi:"sessionControls"`
	// Specifies the state of the policy object. Possible values are: `enabled`, `disabled` and `enabledForReportingButNotEnforced`
	State string `pulumi:"state"`
}

// The set of arguments for constructing a ConditionalAccessPolicy resource.
type ConditionalAccessPolicyArgs struct {
	// A `conditions` block as documented below, which specifies the rules that must be met for the policy to apply.
	Conditions ConditionalAccessPolicyConditionsInput
	// The friendly name for this Conditional Access Policy.
	DisplayName pulumi.StringInput
	// A `grantControls` block as documented below, which specifies the grant controls that must be fulfilled to pass the policy.
	GrantControls ConditionalAccessPolicyGrantControlsPtrInput
	// A `sessionControls` block as documented below, which specifies the session controls that are enforced after sign-in.
	//
	// > Note: At least one of `grantControls` and/or `sessionControls` blocks must be specified.
	SessionControls ConditionalAccessPolicySessionControlsPtrInput
	// Specifies the state of the policy object. Possible values are: `enabled`, `disabled` and `enabledForReportingButNotEnforced`
	State pulumi.StringInput
}

func (ConditionalAccessPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*conditionalAccessPolicyArgs)(nil)).Elem()
}

type ConditionalAccessPolicyInput interface {
	pulumi.Input

	ToConditionalAccessPolicyOutput() ConditionalAccessPolicyOutput
	ToConditionalAccessPolicyOutputWithContext(ctx context.Context) ConditionalAccessPolicyOutput
}

func (*ConditionalAccessPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ConditionalAccessPolicy)(nil)).Elem()
}

func (i *ConditionalAccessPolicy) ToConditionalAccessPolicyOutput() ConditionalAccessPolicyOutput {
	return i.ToConditionalAccessPolicyOutputWithContext(context.Background())
}

func (i *ConditionalAccessPolicy) ToConditionalAccessPolicyOutputWithContext(ctx context.Context) ConditionalAccessPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionalAccessPolicyOutput)
}

// ConditionalAccessPolicyArrayInput is an input type that accepts ConditionalAccessPolicyArray and ConditionalAccessPolicyArrayOutput values.
// You can construct a concrete instance of `ConditionalAccessPolicyArrayInput` via:
//
//	ConditionalAccessPolicyArray{ ConditionalAccessPolicyArgs{...} }
type ConditionalAccessPolicyArrayInput interface {
	pulumi.Input

	ToConditionalAccessPolicyArrayOutput() ConditionalAccessPolicyArrayOutput
	ToConditionalAccessPolicyArrayOutputWithContext(context.Context) ConditionalAccessPolicyArrayOutput
}

type ConditionalAccessPolicyArray []ConditionalAccessPolicyInput

func (ConditionalAccessPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConditionalAccessPolicy)(nil)).Elem()
}

func (i ConditionalAccessPolicyArray) ToConditionalAccessPolicyArrayOutput() ConditionalAccessPolicyArrayOutput {
	return i.ToConditionalAccessPolicyArrayOutputWithContext(context.Background())
}

func (i ConditionalAccessPolicyArray) ToConditionalAccessPolicyArrayOutputWithContext(ctx context.Context) ConditionalAccessPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionalAccessPolicyArrayOutput)
}

// ConditionalAccessPolicyMapInput is an input type that accepts ConditionalAccessPolicyMap and ConditionalAccessPolicyMapOutput values.
// You can construct a concrete instance of `ConditionalAccessPolicyMapInput` via:
//
//	ConditionalAccessPolicyMap{ "key": ConditionalAccessPolicyArgs{...} }
type ConditionalAccessPolicyMapInput interface {
	pulumi.Input

	ToConditionalAccessPolicyMapOutput() ConditionalAccessPolicyMapOutput
	ToConditionalAccessPolicyMapOutputWithContext(context.Context) ConditionalAccessPolicyMapOutput
}

type ConditionalAccessPolicyMap map[string]ConditionalAccessPolicyInput

func (ConditionalAccessPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConditionalAccessPolicy)(nil)).Elem()
}

func (i ConditionalAccessPolicyMap) ToConditionalAccessPolicyMapOutput() ConditionalAccessPolicyMapOutput {
	return i.ToConditionalAccessPolicyMapOutputWithContext(context.Background())
}

func (i ConditionalAccessPolicyMap) ToConditionalAccessPolicyMapOutputWithContext(ctx context.Context) ConditionalAccessPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionalAccessPolicyMapOutput)
}

type ConditionalAccessPolicyOutput struct{ *pulumi.OutputState }

func (ConditionalAccessPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConditionalAccessPolicy)(nil)).Elem()
}

func (o ConditionalAccessPolicyOutput) ToConditionalAccessPolicyOutput() ConditionalAccessPolicyOutput {
	return o
}

func (o ConditionalAccessPolicyOutput) ToConditionalAccessPolicyOutputWithContext(ctx context.Context) ConditionalAccessPolicyOutput {
	return o
}

// A `conditions` block as documented below, which specifies the rules that must be met for the policy to apply.
func (o ConditionalAccessPolicyOutput) Conditions() ConditionalAccessPolicyConditionsOutput {
	return o.ApplyT(func(v *ConditionalAccessPolicy) ConditionalAccessPolicyConditionsOutput { return v.Conditions }).(ConditionalAccessPolicyConditionsOutput)
}

// The friendly name for this Conditional Access Policy.
func (o ConditionalAccessPolicyOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *ConditionalAccessPolicy) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// A `grantControls` block as documented below, which specifies the grant controls that must be fulfilled to pass the policy.
func (o ConditionalAccessPolicyOutput) GrantControls() ConditionalAccessPolicyGrantControlsPtrOutput {
	return o.ApplyT(func(v *ConditionalAccessPolicy) ConditionalAccessPolicyGrantControlsPtrOutput { return v.GrantControls }).(ConditionalAccessPolicyGrantControlsPtrOutput)
}

// A `sessionControls` block as documented below, which specifies the session controls that are enforced after sign-in.
//
// > Note: At least one of `grantControls` and/or `sessionControls` blocks must be specified.
func (o ConditionalAccessPolicyOutput) SessionControls() ConditionalAccessPolicySessionControlsPtrOutput {
	return o.ApplyT(func(v *ConditionalAccessPolicy) ConditionalAccessPolicySessionControlsPtrOutput {
		return v.SessionControls
	}).(ConditionalAccessPolicySessionControlsPtrOutput)
}

// Specifies the state of the policy object. Possible values are: `enabled`, `disabled` and `enabledForReportingButNotEnforced`
func (o ConditionalAccessPolicyOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *ConditionalAccessPolicy) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

type ConditionalAccessPolicyArrayOutput struct{ *pulumi.OutputState }

func (ConditionalAccessPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConditionalAccessPolicy)(nil)).Elem()
}

func (o ConditionalAccessPolicyArrayOutput) ToConditionalAccessPolicyArrayOutput() ConditionalAccessPolicyArrayOutput {
	return o
}

func (o ConditionalAccessPolicyArrayOutput) ToConditionalAccessPolicyArrayOutputWithContext(ctx context.Context) ConditionalAccessPolicyArrayOutput {
	return o
}

func (o ConditionalAccessPolicyArrayOutput) Index(i pulumi.IntInput) ConditionalAccessPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ConditionalAccessPolicy {
		return vs[0].([]*ConditionalAccessPolicy)[vs[1].(int)]
	}).(ConditionalAccessPolicyOutput)
}

type ConditionalAccessPolicyMapOutput struct{ *pulumi.OutputState }

func (ConditionalAccessPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConditionalAccessPolicy)(nil)).Elem()
}

func (o ConditionalAccessPolicyMapOutput) ToConditionalAccessPolicyMapOutput() ConditionalAccessPolicyMapOutput {
	return o
}

func (o ConditionalAccessPolicyMapOutput) ToConditionalAccessPolicyMapOutputWithContext(ctx context.Context) ConditionalAccessPolicyMapOutput {
	return o
}

func (o ConditionalAccessPolicyMapOutput) MapIndex(k pulumi.StringInput) ConditionalAccessPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ConditionalAccessPolicy {
		return vs[0].(map[string]*ConditionalAccessPolicy)[vs[1].(string)]
	}).(ConditionalAccessPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConditionalAccessPolicyInput)(nil)).Elem(), &ConditionalAccessPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConditionalAccessPolicyArrayInput)(nil)).Elem(), ConditionalAccessPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConditionalAccessPolicyMapInput)(nil)).Elem(), ConditionalAccessPolicyMap{})
	pulumi.RegisterOutputType(ConditionalAccessPolicyOutput{})
	pulumi.RegisterOutputType(ConditionalAccessPolicyArrayOutput{})
	pulumi.RegisterOutputType(ConditionalAccessPolicyMapOutput{})
}

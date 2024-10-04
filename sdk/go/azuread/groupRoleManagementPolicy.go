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

// Manage a role policy for an Azure AD group.
//
// ## API Permissions
//
// The following API permissions are required in order to use this resource.
//
// When authenticated with a service principal, this resource requires the `RoleManagementPolicy.ReadWrite.AzureADGroup` Microsoft Graph API permissions.
//
// When authenticated with a user principal, this resource requires `Global Administrator` directory role, or the `Privileged Role Administrator` role in Identity Governance.
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
//			example, err := azuread.NewGroup(ctx, "example", &azuread.GroupArgs{
//				DisplayName:     pulumi.String("group-name"),
//				SecurityEnabled: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuread.NewUser(ctx, "member", &azuread.UserArgs{
//				UserPrincipalName: pulumi.String("jdoe@example.com"),
//				DisplayName:       pulumi.String("J. Doe"),
//				MailNickname:      pulumi.String("jdoe"),
//				Password:          pulumi.String("SecretP@sswd99!"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuread.NewGroupRoleManagementPolicy(ctx, "example", &azuread.GroupRoleManagementPolicyArgs{
//				GroupId: example.ID(),
//				RoleId:  pulumi.String("member"),
//				ActiveAssignmentRules: &azuread.GroupRoleManagementPolicyActiveAssignmentRulesArgs{
//					ExpireAfter: pulumi.String("P365D"),
//				},
//				EligibleAssignmentRules: &azuread.GroupRoleManagementPolicyEligibleAssignmentRulesArgs{
//					ExpirationRequired: pulumi.Bool(false),
//				},
//				NotificationRules: &azuread.GroupRoleManagementPolicyNotificationRulesArgs{
//					EligibleAssignments: &azuread.GroupRoleManagementPolicyNotificationRulesEligibleAssignmentsArgs{
//						ApproverNotifications: &azuread.GroupRoleManagementPolicyNotificationRulesEligibleAssignmentsApproverNotificationsArgs{
//							NotificationLevel: pulumi.String("Critical"),
//							DefaultRecipients: pulumi.Bool(false),
//							AdditionalRecipients: pulumi.StringArray{
//								pulumi.String("someone@example.com"),
//								pulumi.String("someone.else@example.com"),
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
//
// ## Import
//
// Because these policies are created automatically by Entra ID, they will auto-import on first use.
type GroupRoleManagementPolicy struct {
	pulumi.CustomResourceState

	// An `activationRules` block as defined below.
	ActivationRules GroupRoleManagementPolicyActivationRulesOutput `pulumi:"activationRules"`
	// An `activeAssignmentRules` block as defined below.
	ActiveAssignmentRules GroupRoleManagementPolicyActiveAssignmentRulesOutput `pulumi:"activeAssignmentRules"`
	// (String) The description of this policy.
	Description pulumi.StringOutput `pulumi:"description"`
	// (String) The display name of this policy.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// An `eligibleAssignmentRules` block as defined below.
	EligibleAssignmentRules GroupRoleManagementPolicyEligibleAssignmentRulesOutput `pulumi:"eligibleAssignmentRules"`
	// The ID of the Azure AD group for which the policy applies.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// A `notificationRules` block as defined below.
	NotificationRules GroupRoleManagementPolicyNotificationRulesOutput `pulumi:"notificationRules"`
	// The type of assignment this policy coveres. Can be either `member` or `owner`.
	RoleId pulumi.StringOutput `pulumi:"roleId"`
}

// NewGroupRoleManagementPolicy registers a new resource with the given unique name, arguments, and options.
func NewGroupRoleManagementPolicy(ctx *pulumi.Context,
	name string, args *GroupRoleManagementPolicyArgs, opts ...pulumi.ResourceOption) (*GroupRoleManagementPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	if args.RoleId == nil {
		return nil, errors.New("invalid value for required argument 'RoleId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GroupRoleManagementPolicy
	err := ctx.RegisterResource("azuread:index/groupRoleManagementPolicy:GroupRoleManagementPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupRoleManagementPolicy gets an existing GroupRoleManagementPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupRoleManagementPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupRoleManagementPolicyState, opts ...pulumi.ResourceOption) (*GroupRoleManagementPolicy, error) {
	var resource GroupRoleManagementPolicy
	err := ctx.ReadResource("azuread:index/groupRoleManagementPolicy:GroupRoleManagementPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupRoleManagementPolicy resources.
type groupRoleManagementPolicyState struct {
	// An `activationRules` block as defined below.
	ActivationRules *GroupRoleManagementPolicyActivationRules `pulumi:"activationRules"`
	// An `activeAssignmentRules` block as defined below.
	ActiveAssignmentRules *GroupRoleManagementPolicyActiveAssignmentRules `pulumi:"activeAssignmentRules"`
	// (String) The description of this policy.
	Description *string `pulumi:"description"`
	// (String) The display name of this policy.
	DisplayName *string `pulumi:"displayName"`
	// An `eligibleAssignmentRules` block as defined below.
	EligibleAssignmentRules *GroupRoleManagementPolicyEligibleAssignmentRules `pulumi:"eligibleAssignmentRules"`
	// The ID of the Azure AD group for which the policy applies.
	GroupId *string `pulumi:"groupId"`
	// A `notificationRules` block as defined below.
	NotificationRules *GroupRoleManagementPolicyNotificationRules `pulumi:"notificationRules"`
	// The type of assignment this policy coveres. Can be either `member` or `owner`.
	RoleId *string `pulumi:"roleId"`
}

type GroupRoleManagementPolicyState struct {
	// An `activationRules` block as defined below.
	ActivationRules GroupRoleManagementPolicyActivationRulesPtrInput
	// An `activeAssignmentRules` block as defined below.
	ActiveAssignmentRules GroupRoleManagementPolicyActiveAssignmentRulesPtrInput
	// (String) The description of this policy.
	Description pulumi.StringPtrInput
	// (String) The display name of this policy.
	DisplayName pulumi.StringPtrInput
	// An `eligibleAssignmentRules` block as defined below.
	EligibleAssignmentRules GroupRoleManagementPolicyEligibleAssignmentRulesPtrInput
	// The ID of the Azure AD group for which the policy applies.
	GroupId pulumi.StringPtrInput
	// A `notificationRules` block as defined below.
	NotificationRules GroupRoleManagementPolicyNotificationRulesPtrInput
	// The type of assignment this policy coveres. Can be either `member` or `owner`.
	RoleId pulumi.StringPtrInput
}

func (GroupRoleManagementPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupRoleManagementPolicyState)(nil)).Elem()
}

type groupRoleManagementPolicyArgs struct {
	// An `activationRules` block as defined below.
	ActivationRules *GroupRoleManagementPolicyActivationRules `pulumi:"activationRules"`
	// An `activeAssignmentRules` block as defined below.
	ActiveAssignmentRules *GroupRoleManagementPolicyActiveAssignmentRules `pulumi:"activeAssignmentRules"`
	// An `eligibleAssignmentRules` block as defined below.
	EligibleAssignmentRules *GroupRoleManagementPolicyEligibleAssignmentRules `pulumi:"eligibleAssignmentRules"`
	// The ID of the Azure AD group for which the policy applies.
	GroupId string `pulumi:"groupId"`
	// A `notificationRules` block as defined below.
	NotificationRules *GroupRoleManagementPolicyNotificationRules `pulumi:"notificationRules"`
	// The type of assignment this policy coveres. Can be either `member` or `owner`.
	RoleId string `pulumi:"roleId"`
}

// The set of arguments for constructing a GroupRoleManagementPolicy resource.
type GroupRoleManagementPolicyArgs struct {
	// An `activationRules` block as defined below.
	ActivationRules GroupRoleManagementPolicyActivationRulesPtrInput
	// An `activeAssignmentRules` block as defined below.
	ActiveAssignmentRules GroupRoleManagementPolicyActiveAssignmentRulesPtrInput
	// An `eligibleAssignmentRules` block as defined below.
	EligibleAssignmentRules GroupRoleManagementPolicyEligibleAssignmentRulesPtrInput
	// The ID of the Azure AD group for which the policy applies.
	GroupId pulumi.StringInput
	// A `notificationRules` block as defined below.
	NotificationRules GroupRoleManagementPolicyNotificationRulesPtrInput
	// The type of assignment this policy coveres. Can be either `member` or `owner`.
	RoleId pulumi.StringInput
}

func (GroupRoleManagementPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupRoleManagementPolicyArgs)(nil)).Elem()
}

type GroupRoleManagementPolicyInput interface {
	pulumi.Input

	ToGroupRoleManagementPolicyOutput() GroupRoleManagementPolicyOutput
	ToGroupRoleManagementPolicyOutputWithContext(ctx context.Context) GroupRoleManagementPolicyOutput
}

func (*GroupRoleManagementPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupRoleManagementPolicy)(nil)).Elem()
}

func (i *GroupRoleManagementPolicy) ToGroupRoleManagementPolicyOutput() GroupRoleManagementPolicyOutput {
	return i.ToGroupRoleManagementPolicyOutputWithContext(context.Background())
}

func (i *GroupRoleManagementPolicy) ToGroupRoleManagementPolicyOutputWithContext(ctx context.Context) GroupRoleManagementPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupRoleManagementPolicyOutput)
}

// GroupRoleManagementPolicyArrayInput is an input type that accepts GroupRoleManagementPolicyArray and GroupRoleManagementPolicyArrayOutput values.
// You can construct a concrete instance of `GroupRoleManagementPolicyArrayInput` via:
//
//	GroupRoleManagementPolicyArray{ GroupRoleManagementPolicyArgs{...} }
type GroupRoleManagementPolicyArrayInput interface {
	pulumi.Input

	ToGroupRoleManagementPolicyArrayOutput() GroupRoleManagementPolicyArrayOutput
	ToGroupRoleManagementPolicyArrayOutputWithContext(context.Context) GroupRoleManagementPolicyArrayOutput
}

type GroupRoleManagementPolicyArray []GroupRoleManagementPolicyInput

func (GroupRoleManagementPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupRoleManagementPolicy)(nil)).Elem()
}

func (i GroupRoleManagementPolicyArray) ToGroupRoleManagementPolicyArrayOutput() GroupRoleManagementPolicyArrayOutput {
	return i.ToGroupRoleManagementPolicyArrayOutputWithContext(context.Background())
}

func (i GroupRoleManagementPolicyArray) ToGroupRoleManagementPolicyArrayOutputWithContext(ctx context.Context) GroupRoleManagementPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupRoleManagementPolicyArrayOutput)
}

// GroupRoleManagementPolicyMapInput is an input type that accepts GroupRoleManagementPolicyMap and GroupRoleManagementPolicyMapOutput values.
// You can construct a concrete instance of `GroupRoleManagementPolicyMapInput` via:
//
//	GroupRoleManagementPolicyMap{ "key": GroupRoleManagementPolicyArgs{...} }
type GroupRoleManagementPolicyMapInput interface {
	pulumi.Input

	ToGroupRoleManagementPolicyMapOutput() GroupRoleManagementPolicyMapOutput
	ToGroupRoleManagementPolicyMapOutputWithContext(context.Context) GroupRoleManagementPolicyMapOutput
}

type GroupRoleManagementPolicyMap map[string]GroupRoleManagementPolicyInput

func (GroupRoleManagementPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupRoleManagementPolicy)(nil)).Elem()
}

func (i GroupRoleManagementPolicyMap) ToGroupRoleManagementPolicyMapOutput() GroupRoleManagementPolicyMapOutput {
	return i.ToGroupRoleManagementPolicyMapOutputWithContext(context.Background())
}

func (i GroupRoleManagementPolicyMap) ToGroupRoleManagementPolicyMapOutputWithContext(ctx context.Context) GroupRoleManagementPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupRoleManagementPolicyMapOutput)
}

type GroupRoleManagementPolicyOutput struct{ *pulumi.OutputState }

func (GroupRoleManagementPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupRoleManagementPolicy)(nil)).Elem()
}

func (o GroupRoleManagementPolicyOutput) ToGroupRoleManagementPolicyOutput() GroupRoleManagementPolicyOutput {
	return o
}

func (o GroupRoleManagementPolicyOutput) ToGroupRoleManagementPolicyOutputWithContext(ctx context.Context) GroupRoleManagementPolicyOutput {
	return o
}

// An `activationRules` block as defined below.
func (o GroupRoleManagementPolicyOutput) ActivationRules() GroupRoleManagementPolicyActivationRulesOutput {
	return o.ApplyT(func(v *GroupRoleManagementPolicy) GroupRoleManagementPolicyActivationRulesOutput {
		return v.ActivationRules
	}).(GroupRoleManagementPolicyActivationRulesOutput)
}

// An `activeAssignmentRules` block as defined below.
func (o GroupRoleManagementPolicyOutput) ActiveAssignmentRules() GroupRoleManagementPolicyActiveAssignmentRulesOutput {
	return o.ApplyT(func(v *GroupRoleManagementPolicy) GroupRoleManagementPolicyActiveAssignmentRulesOutput {
		return v.ActiveAssignmentRules
	}).(GroupRoleManagementPolicyActiveAssignmentRulesOutput)
}

// (String) The description of this policy.
func (o GroupRoleManagementPolicyOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupRoleManagementPolicy) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// (String) The display name of this policy.
func (o GroupRoleManagementPolicyOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupRoleManagementPolicy) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// An `eligibleAssignmentRules` block as defined below.
func (o GroupRoleManagementPolicyOutput) EligibleAssignmentRules() GroupRoleManagementPolicyEligibleAssignmentRulesOutput {
	return o.ApplyT(func(v *GroupRoleManagementPolicy) GroupRoleManagementPolicyEligibleAssignmentRulesOutput {
		return v.EligibleAssignmentRules
	}).(GroupRoleManagementPolicyEligibleAssignmentRulesOutput)
}

// The ID of the Azure AD group for which the policy applies.
func (o GroupRoleManagementPolicyOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupRoleManagementPolicy) pulumi.StringOutput { return v.GroupId }).(pulumi.StringOutput)
}

// A `notificationRules` block as defined below.
func (o GroupRoleManagementPolicyOutput) NotificationRules() GroupRoleManagementPolicyNotificationRulesOutput {
	return o.ApplyT(func(v *GroupRoleManagementPolicy) GroupRoleManagementPolicyNotificationRulesOutput {
		return v.NotificationRules
	}).(GroupRoleManagementPolicyNotificationRulesOutput)
}

// The type of assignment this policy coveres. Can be either `member` or `owner`.
func (o GroupRoleManagementPolicyOutput) RoleId() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupRoleManagementPolicy) pulumi.StringOutput { return v.RoleId }).(pulumi.StringOutput)
}

type GroupRoleManagementPolicyArrayOutput struct{ *pulumi.OutputState }

func (GroupRoleManagementPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupRoleManagementPolicy)(nil)).Elem()
}

func (o GroupRoleManagementPolicyArrayOutput) ToGroupRoleManagementPolicyArrayOutput() GroupRoleManagementPolicyArrayOutput {
	return o
}

func (o GroupRoleManagementPolicyArrayOutput) ToGroupRoleManagementPolicyArrayOutputWithContext(ctx context.Context) GroupRoleManagementPolicyArrayOutput {
	return o
}

func (o GroupRoleManagementPolicyArrayOutput) Index(i pulumi.IntInput) GroupRoleManagementPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GroupRoleManagementPolicy {
		return vs[0].([]*GroupRoleManagementPolicy)[vs[1].(int)]
	}).(GroupRoleManagementPolicyOutput)
}

type GroupRoleManagementPolicyMapOutput struct{ *pulumi.OutputState }

func (GroupRoleManagementPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupRoleManagementPolicy)(nil)).Elem()
}

func (o GroupRoleManagementPolicyMapOutput) ToGroupRoleManagementPolicyMapOutput() GroupRoleManagementPolicyMapOutput {
	return o
}

func (o GroupRoleManagementPolicyMapOutput) ToGroupRoleManagementPolicyMapOutputWithContext(ctx context.Context) GroupRoleManagementPolicyMapOutput {
	return o
}

func (o GroupRoleManagementPolicyMapOutput) MapIndex(k pulumi.StringInput) GroupRoleManagementPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GroupRoleManagementPolicy {
		return vs[0].(map[string]*GroupRoleManagementPolicy)[vs[1].(string)]
	}).(GroupRoleManagementPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupRoleManagementPolicyInput)(nil)).Elem(), &GroupRoleManagementPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupRoleManagementPolicyArrayInput)(nil)).Elem(), GroupRoleManagementPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupRoleManagementPolicyMapInput)(nil)).Elem(), GroupRoleManagementPolicyMap{})
	pulumi.RegisterOutputType(GroupRoleManagementPolicyOutput{})
	pulumi.RegisterOutputType(GroupRoleManagementPolicyArrayOutput{})
	pulumi.RegisterOutputType(GroupRoleManagementPolicyMapOutput{})
}

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

// Manages an assignment policy for an access package within Identity Governance in Azure Active Directory.
//
// ## API Permissions
//
// The following API permissions are required in order to use this resource.
//
// When authenticated with a service principal, this resource requires the following application role: `EntitlementManagement.ReadWrite.All`.
//
// When authenticated with a user principal, this resource requires `Global Administrator` directory role, or one of the `Catalog Owner` and `Access Package Manager` role in Identity Governance.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			example, err := azuread.NewGroup(ctx, "example", &azuread.GroupArgs{
//				DisplayName:     pulumi.String("group-name"),
//				SecurityEnabled: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			exampleAccessPackageCatalog, err := azuread.NewAccessPackageCatalog(ctx, "example", &azuread.AccessPackageCatalogArgs{
//				DisplayName: pulumi.String("example-catalog"),
//				Description: pulumi.String("Example catalog"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleAccessPackage, err := azuread.NewAccessPackage(ctx, "example", &azuread.AccessPackageArgs{
//				CatalogId:   exampleAccessPackageCatalog.ID(),
//				DisplayName: pulumi.String("access-package"),
//				Description: pulumi.String("Access Package"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuread.NewAccessPackageAssignmentPolicy(ctx, "example", &azuread.AccessPackageAssignmentPolicyArgs{
//				AccessPackageId: exampleAccessPackage.ID(),
//				DisplayName:     pulumi.String("assignment-policy"),
//				Description:     pulumi.String("My assignment policy"),
//				DurationInDays:  pulumi.Int(90),
//				RequestorSettings: &azuread.AccessPackageAssignmentPolicyRequestorSettingsArgs{
//					ScopeType: pulumi.String("AllExistingDirectoryMemberUsers"),
//				},
//				ApprovalSettings: &azuread.AccessPackageAssignmentPolicyApprovalSettingsArgs{
//					ApprovalRequired: pulumi.Bool(true),
//					ApprovalStages: azuread.AccessPackageAssignmentPolicyApprovalSettingsApprovalStageArray{
//						&azuread.AccessPackageAssignmentPolicyApprovalSettingsApprovalStageArgs{
//							ApprovalTimeoutInDays: pulumi.Int(14),
//							PrimaryApprovers: azuread.AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverArray{
//								&azuread.AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverArgs{
//									ObjectId:    example.ObjectId,
//									SubjectType: pulumi.String("groupMembers"),
//								},
//							},
//						},
//					},
//				},
//				AssignmentReviewSettings: &azuread.AccessPackageAssignmentPolicyAssignmentReviewSettingsArgs{
//					Enabled:                     pulumi.Bool(true),
//					ReviewFrequency:             pulumi.String("weekly"),
//					DurationInDays:              pulumi.Int(3),
//					ReviewType:                  pulumi.String("Self"),
//					AccessReviewTimeoutBehavior: pulumi.String("keepAccess"),
//				},
//				Questions: azuread.AccessPackageAssignmentPolicyQuestionArray{
//					&azuread.AccessPackageAssignmentPolicyQuestionArgs{
//						Text: &azuread.AccessPackageAssignmentPolicyQuestionTextArgs{
//							DefaultText: pulumi.String("hello, how are you?"),
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
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// An access package assignment policy can be imported using the ID, e.g.
//
// ```sh
// $ pulumi import azuread:index/accessPackageAssignmentPolicy:AccessPackageAssignmentPolicy example 00000000-0000-0000-0000-000000000000
// ```
type AccessPackageAssignmentPolicy struct {
	pulumi.CustomResourceState

	// The ID of the access package that will contain the policy.
	AccessPackageId pulumi.StringOutput `pulumi:"accessPackageId"`
	// An `approvalSettings` block to specify whether approvals are required and how they are obtained, as documented below.
	ApprovalSettings AccessPackageAssignmentPolicyApprovalSettingsPtrOutput `pulumi:"approvalSettings"`
	// An `assignmentReviewSettings` block, to specify whether assignment review is needed and how it is conducted, as documented below.
	AssignmentReviewSettings AccessPackageAssignmentPolicyAssignmentReviewSettingsPtrOutput `pulumi:"assignmentReviewSettings"`
	// The description of the policy.
	Description pulumi.StringOutput `pulumi:"description"`
	// The display name of the policy.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// How many days this assignment is valid for.
	DurationInDays pulumi.IntPtrOutput `pulumi:"durationInDays"`
	// The date that this assignment expires, formatted as an RFC3339 date string in UTC(e.g. 2018-01-01T01:02:03Z).
	ExpirationDate pulumi.StringPtrOutput `pulumi:"expirationDate"`
	// Whether users will be able to request extension of their access to this package before their access expires.
	ExtensionEnabled pulumi.BoolPtrOutput `pulumi:"extensionEnabled"`
	// One or more `question` blocks for the requestor, as documented below.
	Questions AccessPackageAssignmentPolicyQuestionArrayOutput `pulumi:"questions"`
	// A `requestorSettings` block to configure the users who can request access, as documented below.
	RequestorSettings AccessPackageAssignmentPolicyRequestorSettingsPtrOutput `pulumi:"requestorSettings"`
}

// NewAccessPackageAssignmentPolicy registers a new resource with the given unique name, arguments, and options.
func NewAccessPackageAssignmentPolicy(ctx *pulumi.Context,
	name string, args *AccessPackageAssignmentPolicyArgs, opts ...pulumi.ResourceOption) (*AccessPackageAssignmentPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessPackageId == nil {
		return nil, errors.New("invalid value for required argument 'AccessPackageId'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AccessPackageAssignmentPolicy
	err := ctx.RegisterResource("azuread:index/accessPackageAssignmentPolicy:AccessPackageAssignmentPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessPackageAssignmentPolicy gets an existing AccessPackageAssignmentPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessPackageAssignmentPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessPackageAssignmentPolicyState, opts ...pulumi.ResourceOption) (*AccessPackageAssignmentPolicy, error) {
	var resource AccessPackageAssignmentPolicy
	err := ctx.ReadResource("azuread:index/accessPackageAssignmentPolicy:AccessPackageAssignmentPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessPackageAssignmentPolicy resources.
type accessPackageAssignmentPolicyState struct {
	// The ID of the access package that will contain the policy.
	AccessPackageId *string `pulumi:"accessPackageId"`
	// An `approvalSettings` block to specify whether approvals are required and how they are obtained, as documented below.
	ApprovalSettings *AccessPackageAssignmentPolicyApprovalSettings `pulumi:"approvalSettings"`
	// An `assignmentReviewSettings` block, to specify whether assignment review is needed and how it is conducted, as documented below.
	AssignmentReviewSettings *AccessPackageAssignmentPolicyAssignmentReviewSettings `pulumi:"assignmentReviewSettings"`
	// The description of the policy.
	Description *string `pulumi:"description"`
	// The display name of the policy.
	DisplayName *string `pulumi:"displayName"`
	// How many days this assignment is valid for.
	DurationInDays *int `pulumi:"durationInDays"`
	// The date that this assignment expires, formatted as an RFC3339 date string in UTC(e.g. 2018-01-01T01:02:03Z).
	ExpirationDate *string `pulumi:"expirationDate"`
	// Whether users will be able to request extension of their access to this package before their access expires.
	ExtensionEnabled *bool `pulumi:"extensionEnabled"`
	// One or more `question` blocks for the requestor, as documented below.
	Questions []AccessPackageAssignmentPolicyQuestion `pulumi:"questions"`
	// A `requestorSettings` block to configure the users who can request access, as documented below.
	RequestorSettings *AccessPackageAssignmentPolicyRequestorSettings `pulumi:"requestorSettings"`
}

type AccessPackageAssignmentPolicyState struct {
	// The ID of the access package that will contain the policy.
	AccessPackageId pulumi.StringPtrInput
	// An `approvalSettings` block to specify whether approvals are required and how they are obtained, as documented below.
	ApprovalSettings AccessPackageAssignmentPolicyApprovalSettingsPtrInput
	// An `assignmentReviewSettings` block, to specify whether assignment review is needed and how it is conducted, as documented below.
	AssignmentReviewSettings AccessPackageAssignmentPolicyAssignmentReviewSettingsPtrInput
	// The description of the policy.
	Description pulumi.StringPtrInput
	// The display name of the policy.
	DisplayName pulumi.StringPtrInput
	// How many days this assignment is valid for.
	DurationInDays pulumi.IntPtrInput
	// The date that this assignment expires, formatted as an RFC3339 date string in UTC(e.g. 2018-01-01T01:02:03Z).
	ExpirationDate pulumi.StringPtrInput
	// Whether users will be able to request extension of their access to this package before their access expires.
	ExtensionEnabled pulumi.BoolPtrInput
	// One or more `question` blocks for the requestor, as documented below.
	Questions AccessPackageAssignmentPolicyQuestionArrayInput
	// A `requestorSettings` block to configure the users who can request access, as documented below.
	RequestorSettings AccessPackageAssignmentPolicyRequestorSettingsPtrInput
}

func (AccessPackageAssignmentPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPackageAssignmentPolicyState)(nil)).Elem()
}

type accessPackageAssignmentPolicyArgs struct {
	// The ID of the access package that will contain the policy.
	AccessPackageId string `pulumi:"accessPackageId"`
	// An `approvalSettings` block to specify whether approvals are required and how they are obtained, as documented below.
	ApprovalSettings *AccessPackageAssignmentPolicyApprovalSettings `pulumi:"approvalSettings"`
	// An `assignmentReviewSettings` block, to specify whether assignment review is needed and how it is conducted, as documented below.
	AssignmentReviewSettings *AccessPackageAssignmentPolicyAssignmentReviewSettings `pulumi:"assignmentReviewSettings"`
	// The description of the policy.
	Description string `pulumi:"description"`
	// The display name of the policy.
	DisplayName string `pulumi:"displayName"`
	// How many days this assignment is valid for.
	DurationInDays *int `pulumi:"durationInDays"`
	// The date that this assignment expires, formatted as an RFC3339 date string in UTC(e.g. 2018-01-01T01:02:03Z).
	ExpirationDate *string `pulumi:"expirationDate"`
	// Whether users will be able to request extension of their access to this package before their access expires.
	ExtensionEnabled *bool `pulumi:"extensionEnabled"`
	// One or more `question` blocks for the requestor, as documented below.
	Questions []AccessPackageAssignmentPolicyQuestion `pulumi:"questions"`
	// A `requestorSettings` block to configure the users who can request access, as documented below.
	RequestorSettings *AccessPackageAssignmentPolicyRequestorSettings `pulumi:"requestorSettings"`
}

// The set of arguments for constructing a AccessPackageAssignmentPolicy resource.
type AccessPackageAssignmentPolicyArgs struct {
	// The ID of the access package that will contain the policy.
	AccessPackageId pulumi.StringInput
	// An `approvalSettings` block to specify whether approvals are required and how they are obtained, as documented below.
	ApprovalSettings AccessPackageAssignmentPolicyApprovalSettingsPtrInput
	// An `assignmentReviewSettings` block, to specify whether assignment review is needed and how it is conducted, as documented below.
	AssignmentReviewSettings AccessPackageAssignmentPolicyAssignmentReviewSettingsPtrInput
	// The description of the policy.
	Description pulumi.StringInput
	// The display name of the policy.
	DisplayName pulumi.StringInput
	// How many days this assignment is valid for.
	DurationInDays pulumi.IntPtrInput
	// The date that this assignment expires, formatted as an RFC3339 date string in UTC(e.g. 2018-01-01T01:02:03Z).
	ExpirationDate pulumi.StringPtrInput
	// Whether users will be able to request extension of their access to this package before their access expires.
	ExtensionEnabled pulumi.BoolPtrInput
	// One or more `question` blocks for the requestor, as documented below.
	Questions AccessPackageAssignmentPolicyQuestionArrayInput
	// A `requestorSettings` block to configure the users who can request access, as documented below.
	RequestorSettings AccessPackageAssignmentPolicyRequestorSettingsPtrInput
}

func (AccessPackageAssignmentPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPackageAssignmentPolicyArgs)(nil)).Elem()
}

type AccessPackageAssignmentPolicyInput interface {
	pulumi.Input

	ToAccessPackageAssignmentPolicyOutput() AccessPackageAssignmentPolicyOutput
	ToAccessPackageAssignmentPolicyOutputWithContext(ctx context.Context) AccessPackageAssignmentPolicyOutput
}

func (*AccessPackageAssignmentPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessPackageAssignmentPolicy)(nil)).Elem()
}

func (i *AccessPackageAssignmentPolicy) ToAccessPackageAssignmentPolicyOutput() AccessPackageAssignmentPolicyOutput {
	return i.ToAccessPackageAssignmentPolicyOutputWithContext(context.Background())
}

func (i *AccessPackageAssignmentPolicy) ToAccessPackageAssignmentPolicyOutputWithContext(ctx context.Context) AccessPackageAssignmentPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPackageAssignmentPolicyOutput)
}

// AccessPackageAssignmentPolicyArrayInput is an input type that accepts AccessPackageAssignmentPolicyArray and AccessPackageAssignmentPolicyArrayOutput values.
// You can construct a concrete instance of `AccessPackageAssignmentPolicyArrayInput` via:
//
//	AccessPackageAssignmentPolicyArray{ AccessPackageAssignmentPolicyArgs{...} }
type AccessPackageAssignmentPolicyArrayInput interface {
	pulumi.Input

	ToAccessPackageAssignmentPolicyArrayOutput() AccessPackageAssignmentPolicyArrayOutput
	ToAccessPackageAssignmentPolicyArrayOutputWithContext(context.Context) AccessPackageAssignmentPolicyArrayOutput
}

type AccessPackageAssignmentPolicyArray []AccessPackageAssignmentPolicyInput

func (AccessPackageAssignmentPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessPackageAssignmentPolicy)(nil)).Elem()
}

func (i AccessPackageAssignmentPolicyArray) ToAccessPackageAssignmentPolicyArrayOutput() AccessPackageAssignmentPolicyArrayOutput {
	return i.ToAccessPackageAssignmentPolicyArrayOutputWithContext(context.Background())
}

func (i AccessPackageAssignmentPolicyArray) ToAccessPackageAssignmentPolicyArrayOutputWithContext(ctx context.Context) AccessPackageAssignmentPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPackageAssignmentPolicyArrayOutput)
}

// AccessPackageAssignmentPolicyMapInput is an input type that accepts AccessPackageAssignmentPolicyMap and AccessPackageAssignmentPolicyMapOutput values.
// You can construct a concrete instance of `AccessPackageAssignmentPolicyMapInput` via:
//
//	AccessPackageAssignmentPolicyMap{ "key": AccessPackageAssignmentPolicyArgs{...} }
type AccessPackageAssignmentPolicyMapInput interface {
	pulumi.Input

	ToAccessPackageAssignmentPolicyMapOutput() AccessPackageAssignmentPolicyMapOutput
	ToAccessPackageAssignmentPolicyMapOutputWithContext(context.Context) AccessPackageAssignmentPolicyMapOutput
}

type AccessPackageAssignmentPolicyMap map[string]AccessPackageAssignmentPolicyInput

func (AccessPackageAssignmentPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessPackageAssignmentPolicy)(nil)).Elem()
}

func (i AccessPackageAssignmentPolicyMap) ToAccessPackageAssignmentPolicyMapOutput() AccessPackageAssignmentPolicyMapOutput {
	return i.ToAccessPackageAssignmentPolicyMapOutputWithContext(context.Background())
}

func (i AccessPackageAssignmentPolicyMap) ToAccessPackageAssignmentPolicyMapOutputWithContext(ctx context.Context) AccessPackageAssignmentPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPackageAssignmentPolicyMapOutput)
}

type AccessPackageAssignmentPolicyOutput struct{ *pulumi.OutputState }

func (AccessPackageAssignmentPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessPackageAssignmentPolicy)(nil)).Elem()
}

func (o AccessPackageAssignmentPolicyOutput) ToAccessPackageAssignmentPolicyOutput() AccessPackageAssignmentPolicyOutput {
	return o
}

func (o AccessPackageAssignmentPolicyOutput) ToAccessPackageAssignmentPolicyOutputWithContext(ctx context.Context) AccessPackageAssignmentPolicyOutput {
	return o
}

// The ID of the access package that will contain the policy.
func (o AccessPackageAssignmentPolicyOutput) AccessPackageId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPackageAssignmentPolicy) pulumi.StringOutput { return v.AccessPackageId }).(pulumi.StringOutput)
}

// An `approvalSettings` block to specify whether approvals are required and how they are obtained, as documented below.
func (o AccessPackageAssignmentPolicyOutput) ApprovalSettings() AccessPackageAssignmentPolicyApprovalSettingsPtrOutput {
	return o.ApplyT(func(v *AccessPackageAssignmentPolicy) AccessPackageAssignmentPolicyApprovalSettingsPtrOutput {
		return v.ApprovalSettings
	}).(AccessPackageAssignmentPolicyApprovalSettingsPtrOutput)
}

// An `assignmentReviewSettings` block, to specify whether assignment review is needed and how it is conducted, as documented below.
func (o AccessPackageAssignmentPolicyOutput) AssignmentReviewSettings() AccessPackageAssignmentPolicyAssignmentReviewSettingsPtrOutput {
	return o.ApplyT(func(v *AccessPackageAssignmentPolicy) AccessPackageAssignmentPolicyAssignmentReviewSettingsPtrOutput {
		return v.AssignmentReviewSettings
	}).(AccessPackageAssignmentPolicyAssignmentReviewSettingsPtrOutput)
}

// The description of the policy.
func (o AccessPackageAssignmentPolicyOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPackageAssignmentPolicy) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The display name of the policy.
func (o AccessPackageAssignmentPolicyOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPackageAssignmentPolicy) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// How many days this assignment is valid for.
func (o AccessPackageAssignmentPolicyOutput) DurationInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AccessPackageAssignmentPolicy) pulumi.IntPtrOutput { return v.DurationInDays }).(pulumi.IntPtrOutput)
}

// The date that this assignment expires, formatted as an RFC3339 date string in UTC(e.g. 2018-01-01T01:02:03Z).
func (o AccessPackageAssignmentPolicyOutput) ExpirationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessPackageAssignmentPolicy) pulumi.StringPtrOutput { return v.ExpirationDate }).(pulumi.StringPtrOutput)
}

// Whether users will be able to request extension of their access to this package before their access expires.
func (o AccessPackageAssignmentPolicyOutput) ExtensionEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AccessPackageAssignmentPolicy) pulumi.BoolPtrOutput { return v.ExtensionEnabled }).(pulumi.BoolPtrOutput)
}

// One or more `question` blocks for the requestor, as documented below.
func (o AccessPackageAssignmentPolicyOutput) Questions() AccessPackageAssignmentPolicyQuestionArrayOutput {
	return o.ApplyT(func(v *AccessPackageAssignmentPolicy) AccessPackageAssignmentPolicyQuestionArrayOutput {
		return v.Questions
	}).(AccessPackageAssignmentPolicyQuestionArrayOutput)
}

// A `requestorSettings` block to configure the users who can request access, as documented below.
func (o AccessPackageAssignmentPolicyOutput) RequestorSettings() AccessPackageAssignmentPolicyRequestorSettingsPtrOutput {
	return o.ApplyT(func(v *AccessPackageAssignmentPolicy) AccessPackageAssignmentPolicyRequestorSettingsPtrOutput {
		return v.RequestorSettings
	}).(AccessPackageAssignmentPolicyRequestorSettingsPtrOutput)
}

type AccessPackageAssignmentPolicyArrayOutput struct{ *pulumi.OutputState }

func (AccessPackageAssignmentPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessPackageAssignmentPolicy)(nil)).Elem()
}

func (o AccessPackageAssignmentPolicyArrayOutput) ToAccessPackageAssignmentPolicyArrayOutput() AccessPackageAssignmentPolicyArrayOutput {
	return o
}

func (o AccessPackageAssignmentPolicyArrayOutput) ToAccessPackageAssignmentPolicyArrayOutputWithContext(ctx context.Context) AccessPackageAssignmentPolicyArrayOutput {
	return o
}

func (o AccessPackageAssignmentPolicyArrayOutput) Index(i pulumi.IntInput) AccessPackageAssignmentPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AccessPackageAssignmentPolicy {
		return vs[0].([]*AccessPackageAssignmentPolicy)[vs[1].(int)]
	}).(AccessPackageAssignmentPolicyOutput)
}

type AccessPackageAssignmentPolicyMapOutput struct{ *pulumi.OutputState }

func (AccessPackageAssignmentPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessPackageAssignmentPolicy)(nil)).Elem()
}

func (o AccessPackageAssignmentPolicyMapOutput) ToAccessPackageAssignmentPolicyMapOutput() AccessPackageAssignmentPolicyMapOutput {
	return o
}

func (o AccessPackageAssignmentPolicyMapOutput) ToAccessPackageAssignmentPolicyMapOutputWithContext(ctx context.Context) AccessPackageAssignmentPolicyMapOutput {
	return o
}

func (o AccessPackageAssignmentPolicyMapOutput) MapIndex(k pulumi.StringInput) AccessPackageAssignmentPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AccessPackageAssignmentPolicy {
		return vs[0].(map[string]*AccessPackageAssignmentPolicy)[vs[1].(string)]
	}).(AccessPackageAssignmentPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessPackageAssignmentPolicyInput)(nil)).Elem(), &AccessPackageAssignmentPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessPackageAssignmentPolicyArrayInput)(nil)).Elem(), AccessPackageAssignmentPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessPackageAssignmentPolicyMapInput)(nil)).Elem(), AccessPackageAssignmentPolicyMap{})
	pulumi.RegisterOutputType(AccessPackageAssignmentPolicyOutput{})
	pulumi.RegisterOutputType(AccessPackageAssignmentPolicyArrayOutput{})
	pulumi.RegisterOutputType(AccessPackageAssignmentPolicyMapOutput{})
}

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

// Manages a single group membership within Azure Active Directory.
//
// > **Warning** Do not use this resource at the same time as the `members` property of the `Group` resource for the same group. Doing so will cause a conflict and group members will be removed.
//
// ## API Permissions
//
// The following API permissions are required in order to use this resource.
//
// When authenticated with a service principal, this resource requires one of the following application roles: `Group.ReadWrite.All` or `Directory.ReadWrite.All`.
//
// However, if the authenticated service principal is an owner of the group being managed, an application role is not required.
//
// When authenticated with a user principal, this resource requires one of the following directory roles: `Groups Administrator`, `User Administrator` or `Global Administrator`
//
// ## Example Usage
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
//			example, err := azuread.LookupUser(ctx, &azuread.LookupUserArgs{
//				UserPrincipalName: pulumi.StringRef("jdoe@example.com"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleGroup, err := azuread.NewGroup(ctx, "example", &azuread.GroupArgs{
//				DisplayName:     pulumi.String("my_group"),
//				SecurityEnabled: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuread.NewGroupMember(ctx, "example", &azuread.GroupMemberArgs{
//				GroupObjectId:  exampleGroup.ID(),
//				MemberObjectId: *pulumi.String(example.Id),
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
// Group members can be imported using the object ID of the group and the object ID of the member, e.g.
//
// ```sh
// $ pulumi import azuread:index/groupMember:GroupMember example 00000000-0000-0000-0000-000000000000/member/11111111-1111-1111-1111-111111111111
// ```
//
//	-> This ID format is unique to Terraform and is composed of the Azure AD Group Object ID and the target Member Object ID in the format `{GroupObjectID}/member/{MemberObjectID}`.
type GroupMember struct {
	pulumi.CustomResourceState

	// The object ID of the group you want to add the member to. Changing this forces a new resource to be created.
	GroupObjectId pulumi.StringOutput `pulumi:"groupObjectId"`
	// The object ID of the principal you want to add as a member to the group. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
	MemberObjectId pulumi.StringOutput `pulumi:"memberObjectId"`
}

// NewGroupMember registers a new resource with the given unique name, arguments, and options.
func NewGroupMember(ctx *pulumi.Context,
	name string, args *GroupMemberArgs, opts ...pulumi.ResourceOption) (*GroupMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupObjectId == nil {
		return nil, errors.New("invalid value for required argument 'GroupObjectId'")
	}
	if args.MemberObjectId == nil {
		return nil, errors.New("invalid value for required argument 'MemberObjectId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GroupMember
	err := ctx.RegisterResource("azuread:index/groupMember:GroupMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupMember gets an existing GroupMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupMemberState, opts ...pulumi.ResourceOption) (*GroupMember, error) {
	var resource GroupMember
	err := ctx.ReadResource("azuread:index/groupMember:GroupMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupMember resources.
type groupMemberState struct {
	// The object ID of the group you want to add the member to. Changing this forces a new resource to be created.
	GroupObjectId *string `pulumi:"groupObjectId"`
	// The object ID of the principal you want to add as a member to the group. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
	MemberObjectId *string `pulumi:"memberObjectId"`
}

type GroupMemberState struct {
	// The object ID of the group you want to add the member to. Changing this forces a new resource to be created.
	GroupObjectId pulumi.StringPtrInput
	// The object ID of the principal you want to add as a member to the group. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
	MemberObjectId pulumi.StringPtrInput
}

func (GroupMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMemberState)(nil)).Elem()
}

type groupMemberArgs struct {
	// The object ID of the group you want to add the member to. Changing this forces a new resource to be created.
	GroupObjectId string `pulumi:"groupObjectId"`
	// The object ID of the principal you want to add as a member to the group. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
	MemberObjectId string `pulumi:"memberObjectId"`
}

// The set of arguments for constructing a GroupMember resource.
type GroupMemberArgs struct {
	// The object ID of the group you want to add the member to. Changing this forces a new resource to be created.
	GroupObjectId pulumi.StringInput
	// The object ID of the principal you want to add as a member to the group. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
	MemberObjectId pulumi.StringInput
}

func (GroupMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMemberArgs)(nil)).Elem()
}

type GroupMemberInput interface {
	pulumi.Input

	ToGroupMemberOutput() GroupMemberOutput
	ToGroupMemberOutputWithContext(ctx context.Context) GroupMemberOutput
}

func (*GroupMember) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupMember)(nil)).Elem()
}

func (i *GroupMember) ToGroupMemberOutput() GroupMemberOutput {
	return i.ToGroupMemberOutputWithContext(context.Background())
}

func (i *GroupMember) ToGroupMemberOutputWithContext(ctx context.Context) GroupMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMemberOutput)
}

// GroupMemberArrayInput is an input type that accepts GroupMemberArray and GroupMemberArrayOutput values.
// You can construct a concrete instance of `GroupMemberArrayInput` via:
//
//	GroupMemberArray{ GroupMemberArgs{...} }
type GroupMemberArrayInput interface {
	pulumi.Input

	ToGroupMemberArrayOutput() GroupMemberArrayOutput
	ToGroupMemberArrayOutputWithContext(context.Context) GroupMemberArrayOutput
}

type GroupMemberArray []GroupMemberInput

func (GroupMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupMember)(nil)).Elem()
}

func (i GroupMemberArray) ToGroupMemberArrayOutput() GroupMemberArrayOutput {
	return i.ToGroupMemberArrayOutputWithContext(context.Background())
}

func (i GroupMemberArray) ToGroupMemberArrayOutputWithContext(ctx context.Context) GroupMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMemberArrayOutput)
}

// GroupMemberMapInput is an input type that accepts GroupMemberMap and GroupMemberMapOutput values.
// You can construct a concrete instance of `GroupMemberMapInput` via:
//
//	GroupMemberMap{ "key": GroupMemberArgs{...} }
type GroupMemberMapInput interface {
	pulumi.Input

	ToGroupMemberMapOutput() GroupMemberMapOutput
	ToGroupMemberMapOutputWithContext(context.Context) GroupMemberMapOutput
}

type GroupMemberMap map[string]GroupMemberInput

func (GroupMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupMember)(nil)).Elem()
}

func (i GroupMemberMap) ToGroupMemberMapOutput() GroupMemberMapOutput {
	return i.ToGroupMemberMapOutputWithContext(context.Background())
}

func (i GroupMemberMap) ToGroupMemberMapOutputWithContext(ctx context.Context) GroupMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMemberMapOutput)
}

type GroupMemberOutput struct{ *pulumi.OutputState }

func (GroupMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupMember)(nil)).Elem()
}

func (o GroupMemberOutput) ToGroupMemberOutput() GroupMemberOutput {
	return o
}

func (o GroupMemberOutput) ToGroupMemberOutputWithContext(ctx context.Context) GroupMemberOutput {
	return o
}

// The object ID of the group you want to add the member to. Changing this forces a new resource to be created.
func (o GroupMemberOutput) GroupObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupMember) pulumi.StringOutput { return v.GroupObjectId }).(pulumi.StringOutput)
}

// The object ID of the principal you want to add as a member to the group. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
func (o GroupMemberOutput) MemberObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupMember) pulumi.StringOutput { return v.MemberObjectId }).(pulumi.StringOutput)
}

type GroupMemberArrayOutput struct{ *pulumi.OutputState }

func (GroupMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupMember)(nil)).Elem()
}

func (o GroupMemberArrayOutput) ToGroupMemberArrayOutput() GroupMemberArrayOutput {
	return o
}

func (o GroupMemberArrayOutput) ToGroupMemberArrayOutputWithContext(ctx context.Context) GroupMemberArrayOutput {
	return o
}

func (o GroupMemberArrayOutput) Index(i pulumi.IntInput) GroupMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GroupMember {
		return vs[0].([]*GroupMember)[vs[1].(int)]
	}).(GroupMemberOutput)
}

type GroupMemberMapOutput struct{ *pulumi.OutputState }

func (GroupMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupMember)(nil)).Elem()
}

func (o GroupMemberMapOutput) ToGroupMemberMapOutput() GroupMemberMapOutput {
	return o
}

func (o GroupMemberMapOutput) ToGroupMemberMapOutputWithContext(ctx context.Context) GroupMemberMapOutput {
	return o
}

func (o GroupMemberMapOutput) MapIndex(k pulumi.StringInput) GroupMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GroupMember {
		return vs[0].(map[string]*GroupMember)[vs[1].(string)]
	}).(GroupMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMemberInput)(nil)).Elem(), &GroupMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMemberArrayInput)(nil)).Elem(), GroupMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMemberMapInput)(nil)).Elem(), GroupMemberMap{})
	pulumi.RegisterOutputType(GroupMemberOutput{})
	pulumi.RegisterOutputType(GroupMemberArrayOutput{})
	pulumi.RegisterOutputType(GroupMemberMapOutput{})
}

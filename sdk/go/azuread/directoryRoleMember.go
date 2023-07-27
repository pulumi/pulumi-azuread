// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuread/sdk/v5/go/azuread/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a single directory role membership (assignment) within Azure Active Directory.
//
// > **Deprecation Warning:** This resource has been superseded by the DirectoryRoleAssignment resource and will be removed in version 3.0 of the AzureAD provider
//
// ## API Permissions
//
// The following API permissions are required in order to use this resource.
//
// When authenticated with a service principal, this resource requires one of the following application roles: `RoleManagement.ReadWrite.Directory` or `Directory.ReadWrite.All`
//
// When authenticated with a user principal, this resource requires one of the following directory roles: `Privileged Role Administrator` or `Global Administrator`
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
//			exampleUser, err := azuread.LookupUser(ctx, &azuread.LookupUserArgs{
//				UserPrincipalName: pulumi.StringRef("jdoe@hashicorp.com"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleDirectoryRole, err := azuread.NewDirectoryRole(ctx, "exampleDirectoryRole", &azuread.DirectoryRoleArgs{
//				DisplayName: pulumi.String("Security administrator"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuread.NewDirectoryRoleMember(ctx, "exampleDirectoryRoleMember", &azuread.DirectoryRoleMemberArgs{
//				RoleObjectId:   exampleDirectoryRole.ObjectId,
//				MemberObjectId: *pulumi.String(exampleUser.ObjectId),
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
// Directory role members can be imported using the object ID of the role and the object ID of the member, e.g.
//
// ```sh
//
//	$ pulumi import azuread:index/directoryRoleMember:DirectoryRoleMember test 00000000-0000-0000-0000-000000000000/member/11111111-1111-1111-1111-111111111111
//
// ```
//
//	-> This ID format is unique to Terraform and is composed of the Directory Role Object ID and the target Member Object ID in the format `{RoleObjectID}/member/{MemberObjectID}`.
type DirectoryRoleMember struct {
	pulumi.CustomResourceState

	// The object ID of the principal you want to add as a member to the directory role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
	MemberObjectId pulumi.StringPtrOutput `pulumi:"memberObjectId"`
	// The object ID of the directory role you want to add the member to. Changing this forces a new resource to be created.
	RoleObjectId pulumi.StringPtrOutput `pulumi:"roleObjectId"`
}

// NewDirectoryRoleMember registers a new resource with the given unique name, arguments, and options.
func NewDirectoryRoleMember(ctx *pulumi.Context,
	name string, args *DirectoryRoleMemberArgs, opts ...pulumi.ResourceOption) (*DirectoryRoleMember, error) {
	if args == nil {
		args = &DirectoryRoleMemberArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DirectoryRoleMember
	err := ctx.RegisterResource("azuread:index/directoryRoleMember:DirectoryRoleMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDirectoryRoleMember gets an existing DirectoryRoleMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDirectoryRoleMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DirectoryRoleMemberState, opts ...pulumi.ResourceOption) (*DirectoryRoleMember, error) {
	var resource DirectoryRoleMember
	err := ctx.ReadResource("azuread:index/directoryRoleMember:DirectoryRoleMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DirectoryRoleMember resources.
type directoryRoleMemberState struct {
	// The object ID of the principal you want to add as a member to the directory role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
	MemberObjectId *string `pulumi:"memberObjectId"`
	// The object ID of the directory role you want to add the member to. Changing this forces a new resource to be created.
	RoleObjectId *string `pulumi:"roleObjectId"`
}

type DirectoryRoleMemberState struct {
	// The object ID of the principal you want to add as a member to the directory role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
	MemberObjectId pulumi.StringPtrInput
	// The object ID of the directory role you want to add the member to. Changing this forces a new resource to be created.
	RoleObjectId pulumi.StringPtrInput
}

func (DirectoryRoleMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*directoryRoleMemberState)(nil)).Elem()
}

type directoryRoleMemberArgs struct {
	// The object ID of the principal you want to add as a member to the directory role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
	MemberObjectId *string `pulumi:"memberObjectId"`
	// The object ID of the directory role you want to add the member to. Changing this forces a new resource to be created.
	RoleObjectId *string `pulumi:"roleObjectId"`
}

// The set of arguments for constructing a DirectoryRoleMember resource.
type DirectoryRoleMemberArgs struct {
	// The object ID of the principal you want to add as a member to the directory role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
	MemberObjectId pulumi.StringPtrInput
	// The object ID of the directory role you want to add the member to. Changing this forces a new resource to be created.
	RoleObjectId pulumi.StringPtrInput
}

func (DirectoryRoleMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*directoryRoleMemberArgs)(nil)).Elem()
}

type DirectoryRoleMemberInput interface {
	pulumi.Input

	ToDirectoryRoleMemberOutput() DirectoryRoleMemberOutput
	ToDirectoryRoleMemberOutputWithContext(ctx context.Context) DirectoryRoleMemberOutput
}

func (*DirectoryRoleMember) ElementType() reflect.Type {
	return reflect.TypeOf((**DirectoryRoleMember)(nil)).Elem()
}

func (i *DirectoryRoleMember) ToDirectoryRoleMemberOutput() DirectoryRoleMemberOutput {
	return i.ToDirectoryRoleMemberOutputWithContext(context.Background())
}

func (i *DirectoryRoleMember) ToDirectoryRoleMemberOutputWithContext(ctx context.Context) DirectoryRoleMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectoryRoleMemberOutput)
}

// DirectoryRoleMemberArrayInput is an input type that accepts DirectoryRoleMemberArray and DirectoryRoleMemberArrayOutput values.
// You can construct a concrete instance of `DirectoryRoleMemberArrayInput` via:
//
//	DirectoryRoleMemberArray{ DirectoryRoleMemberArgs{...} }
type DirectoryRoleMemberArrayInput interface {
	pulumi.Input

	ToDirectoryRoleMemberArrayOutput() DirectoryRoleMemberArrayOutput
	ToDirectoryRoleMemberArrayOutputWithContext(context.Context) DirectoryRoleMemberArrayOutput
}

type DirectoryRoleMemberArray []DirectoryRoleMemberInput

func (DirectoryRoleMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DirectoryRoleMember)(nil)).Elem()
}

func (i DirectoryRoleMemberArray) ToDirectoryRoleMemberArrayOutput() DirectoryRoleMemberArrayOutput {
	return i.ToDirectoryRoleMemberArrayOutputWithContext(context.Background())
}

func (i DirectoryRoleMemberArray) ToDirectoryRoleMemberArrayOutputWithContext(ctx context.Context) DirectoryRoleMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectoryRoleMemberArrayOutput)
}

// DirectoryRoleMemberMapInput is an input type that accepts DirectoryRoleMemberMap and DirectoryRoleMemberMapOutput values.
// You can construct a concrete instance of `DirectoryRoleMemberMapInput` via:
//
//	DirectoryRoleMemberMap{ "key": DirectoryRoleMemberArgs{...} }
type DirectoryRoleMemberMapInput interface {
	pulumi.Input

	ToDirectoryRoleMemberMapOutput() DirectoryRoleMemberMapOutput
	ToDirectoryRoleMemberMapOutputWithContext(context.Context) DirectoryRoleMemberMapOutput
}

type DirectoryRoleMemberMap map[string]DirectoryRoleMemberInput

func (DirectoryRoleMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DirectoryRoleMember)(nil)).Elem()
}

func (i DirectoryRoleMemberMap) ToDirectoryRoleMemberMapOutput() DirectoryRoleMemberMapOutput {
	return i.ToDirectoryRoleMemberMapOutputWithContext(context.Background())
}

func (i DirectoryRoleMemberMap) ToDirectoryRoleMemberMapOutputWithContext(ctx context.Context) DirectoryRoleMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectoryRoleMemberMapOutput)
}

type DirectoryRoleMemberOutput struct{ *pulumi.OutputState }

func (DirectoryRoleMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DirectoryRoleMember)(nil)).Elem()
}

func (o DirectoryRoleMemberOutput) ToDirectoryRoleMemberOutput() DirectoryRoleMemberOutput {
	return o
}

func (o DirectoryRoleMemberOutput) ToDirectoryRoleMemberOutputWithContext(ctx context.Context) DirectoryRoleMemberOutput {
	return o
}

// The object ID of the principal you want to add as a member to the directory role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
func (o DirectoryRoleMemberOutput) MemberObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DirectoryRoleMember) pulumi.StringPtrOutput { return v.MemberObjectId }).(pulumi.StringPtrOutput)
}

// The object ID of the directory role you want to add the member to. Changing this forces a new resource to be created.
func (o DirectoryRoleMemberOutput) RoleObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DirectoryRoleMember) pulumi.StringPtrOutput { return v.RoleObjectId }).(pulumi.StringPtrOutput)
}

type DirectoryRoleMemberArrayOutput struct{ *pulumi.OutputState }

func (DirectoryRoleMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DirectoryRoleMember)(nil)).Elem()
}

func (o DirectoryRoleMemberArrayOutput) ToDirectoryRoleMemberArrayOutput() DirectoryRoleMemberArrayOutput {
	return o
}

func (o DirectoryRoleMemberArrayOutput) ToDirectoryRoleMemberArrayOutputWithContext(ctx context.Context) DirectoryRoleMemberArrayOutput {
	return o
}

func (o DirectoryRoleMemberArrayOutput) Index(i pulumi.IntInput) DirectoryRoleMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DirectoryRoleMember {
		return vs[0].([]*DirectoryRoleMember)[vs[1].(int)]
	}).(DirectoryRoleMemberOutput)
}

type DirectoryRoleMemberMapOutput struct{ *pulumi.OutputState }

func (DirectoryRoleMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DirectoryRoleMember)(nil)).Elem()
}

func (o DirectoryRoleMemberMapOutput) ToDirectoryRoleMemberMapOutput() DirectoryRoleMemberMapOutput {
	return o
}

func (o DirectoryRoleMemberMapOutput) ToDirectoryRoleMemberMapOutputWithContext(ctx context.Context) DirectoryRoleMemberMapOutput {
	return o
}

func (o DirectoryRoleMemberMapOutput) MapIndex(k pulumi.StringInput) DirectoryRoleMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DirectoryRoleMember {
		return vs[0].(map[string]*DirectoryRoleMember)[vs[1].(string)]
	}).(DirectoryRoleMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DirectoryRoleMemberInput)(nil)).Elem(), &DirectoryRoleMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*DirectoryRoleMemberArrayInput)(nil)).Elem(), DirectoryRoleMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DirectoryRoleMemberMapInput)(nil)).Elem(), DirectoryRoleMemberMap{})
	pulumi.RegisterOutputType(DirectoryRoleMemberOutput{})
	pulumi.RegisterOutputType(DirectoryRoleMemberArrayOutput{})
	pulumi.RegisterOutputType(DirectoryRoleMemberMapOutput{})
}

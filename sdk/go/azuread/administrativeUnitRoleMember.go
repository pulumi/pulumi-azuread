// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azuread/sdk/v5/go/azuread/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Manages a single directory role assignment scoped to an administrative unit within Azure Active Directory.
//
// ## API Permissions
//
// The following API permissions are required in order to use this resource.
//
// When authenticated with a service principal, this resource requires one of the following application roles: `AdministrativeUnit.ReadWrite.All` and `RoleManagement.ReadWrite.Directory`, or `Directory.ReadWrite.All`
//
// When authenticated with a user principal, this resource requires one of the following directory roles: `Privileged Role Administrator` or `Global Administrator`
//
// ## Import
//
// Administrative unit role members can be imported using the object ID of the administrative unit and the unique ID of the role assignment, e.g.
//
// ```sh
//
//	$ pulumi import azuread:index/administrativeUnitRoleMember:AdministrativeUnitRoleMember example 00000000-0000-0000-0000-000000000000/roleMember/zX37MRLyF0uvE-xf2WH4B7x-6CPLfudNnxFGj800htpBXqkxW7bITqGb6Rj4kuTuS
//
// ```
//
//	-> This ID format is unique to Terraform and is composed of the Administrative Unit Object ID and the role assignment ID in the format `{AdministrativeUnitObjectID}/roleMember/{RoleAssignmentID}`.
type AdministrativeUnitRoleMember struct {
	pulumi.CustomResourceState

	// The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
	AdministrativeUnitObjectId pulumi.StringOutput `pulumi:"administrativeUnitObjectId"`
	// The object ID of the user, group or service principal you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
	MemberObjectId pulumi.StringOutput `pulumi:"memberObjectId"`
	// The object ID of the directory role you want to assign. Changing this forces a new resource to be created.
	RoleObjectId pulumi.StringOutput `pulumi:"roleObjectId"`
}

// NewAdministrativeUnitRoleMember registers a new resource with the given unique name, arguments, and options.
func NewAdministrativeUnitRoleMember(ctx *pulumi.Context,
	name string, args *AdministrativeUnitRoleMemberArgs, opts ...pulumi.ResourceOption) (*AdministrativeUnitRoleMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AdministrativeUnitObjectId == nil {
		return nil, errors.New("invalid value for required argument 'AdministrativeUnitObjectId'")
	}
	if args.MemberObjectId == nil {
		return nil, errors.New("invalid value for required argument 'MemberObjectId'")
	}
	if args.RoleObjectId == nil {
		return nil, errors.New("invalid value for required argument 'RoleObjectId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AdministrativeUnitRoleMember
	err := ctx.RegisterResource("azuread:index/administrativeUnitRoleMember:AdministrativeUnitRoleMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAdministrativeUnitRoleMember gets an existing AdministrativeUnitRoleMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAdministrativeUnitRoleMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AdministrativeUnitRoleMemberState, opts ...pulumi.ResourceOption) (*AdministrativeUnitRoleMember, error) {
	var resource AdministrativeUnitRoleMember
	err := ctx.ReadResource("azuread:index/administrativeUnitRoleMember:AdministrativeUnitRoleMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AdministrativeUnitRoleMember resources.
type administrativeUnitRoleMemberState struct {
	// The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
	AdministrativeUnitObjectId *string `pulumi:"administrativeUnitObjectId"`
	// The object ID of the user, group or service principal you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
	MemberObjectId *string `pulumi:"memberObjectId"`
	// The object ID of the directory role you want to assign. Changing this forces a new resource to be created.
	RoleObjectId *string `pulumi:"roleObjectId"`
}

type AdministrativeUnitRoleMemberState struct {
	// The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
	AdministrativeUnitObjectId pulumi.StringPtrInput
	// The object ID of the user, group or service principal you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
	MemberObjectId pulumi.StringPtrInput
	// The object ID of the directory role you want to assign. Changing this forces a new resource to be created.
	RoleObjectId pulumi.StringPtrInput
}

func (AdministrativeUnitRoleMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*administrativeUnitRoleMemberState)(nil)).Elem()
}

type administrativeUnitRoleMemberArgs struct {
	// The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
	AdministrativeUnitObjectId string `pulumi:"administrativeUnitObjectId"`
	// The object ID of the user, group or service principal you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
	MemberObjectId string `pulumi:"memberObjectId"`
	// The object ID of the directory role you want to assign. Changing this forces a new resource to be created.
	RoleObjectId string `pulumi:"roleObjectId"`
}

// The set of arguments for constructing a AdministrativeUnitRoleMember resource.
type AdministrativeUnitRoleMemberArgs struct {
	// The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
	AdministrativeUnitObjectId pulumi.StringInput
	// The object ID of the user, group or service principal you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
	MemberObjectId pulumi.StringInput
	// The object ID of the directory role you want to assign. Changing this forces a new resource to be created.
	RoleObjectId pulumi.StringInput
}

func (AdministrativeUnitRoleMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*administrativeUnitRoleMemberArgs)(nil)).Elem()
}

type AdministrativeUnitRoleMemberInput interface {
	pulumi.Input

	ToAdministrativeUnitRoleMemberOutput() AdministrativeUnitRoleMemberOutput
	ToAdministrativeUnitRoleMemberOutputWithContext(ctx context.Context) AdministrativeUnitRoleMemberOutput
}

func (*AdministrativeUnitRoleMember) ElementType() reflect.Type {
	return reflect.TypeOf((**AdministrativeUnitRoleMember)(nil)).Elem()
}

func (i *AdministrativeUnitRoleMember) ToAdministrativeUnitRoleMemberOutput() AdministrativeUnitRoleMemberOutput {
	return i.ToAdministrativeUnitRoleMemberOutputWithContext(context.Background())
}

func (i *AdministrativeUnitRoleMember) ToAdministrativeUnitRoleMemberOutputWithContext(ctx context.Context) AdministrativeUnitRoleMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdministrativeUnitRoleMemberOutput)
}

func (i *AdministrativeUnitRoleMember) ToOutput(ctx context.Context) pulumix.Output[*AdministrativeUnitRoleMember] {
	return pulumix.Output[*AdministrativeUnitRoleMember]{
		OutputState: i.ToAdministrativeUnitRoleMemberOutputWithContext(ctx).OutputState,
	}
}

// AdministrativeUnitRoleMemberArrayInput is an input type that accepts AdministrativeUnitRoleMemberArray and AdministrativeUnitRoleMemberArrayOutput values.
// You can construct a concrete instance of `AdministrativeUnitRoleMemberArrayInput` via:
//
//	AdministrativeUnitRoleMemberArray{ AdministrativeUnitRoleMemberArgs{...} }
type AdministrativeUnitRoleMemberArrayInput interface {
	pulumi.Input

	ToAdministrativeUnitRoleMemberArrayOutput() AdministrativeUnitRoleMemberArrayOutput
	ToAdministrativeUnitRoleMemberArrayOutputWithContext(context.Context) AdministrativeUnitRoleMemberArrayOutput
}

type AdministrativeUnitRoleMemberArray []AdministrativeUnitRoleMemberInput

func (AdministrativeUnitRoleMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AdministrativeUnitRoleMember)(nil)).Elem()
}

func (i AdministrativeUnitRoleMemberArray) ToAdministrativeUnitRoleMemberArrayOutput() AdministrativeUnitRoleMemberArrayOutput {
	return i.ToAdministrativeUnitRoleMemberArrayOutputWithContext(context.Background())
}

func (i AdministrativeUnitRoleMemberArray) ToAdministrativeUnitRoleMemberArrayOutputWithContext(ctx context.Context) AdministrativeUnitRoleMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdministrativeUnitRoleMemberArrayOutput)
}

func (i AdministrativeUnitRoleMemberArray) ToOutput(ctx context.Context) pulumix.Output[[]*AdministrativeUnitRoleMember] {
	return pulumix.Output[[]*AdministrativeUnitRoleMember]{
		OutputState: i.ToAdministrativeUnitRoleMemberArrayOutputWithContext(ctx).OutputState,
	}
}

// AdministrativeUnitRoleMemberMapInput is an input type that accepts AdministrativeUnitRoleMemberMap and AdministrativeUnitRoleMemberMapOutput values.
// You can construct a concrete instance of `AdministrativeUnitRoleMemberMapInput` via:
//
//	AdministrativeUnitRoleMemberMap{ "key": AdministrativeUnitRoleMemberArgs{...} }
type AdministrativeUnitRoleMemberMapInput interface {
	pulumi.Input

	ToAdministrativeUnitRoleMemberMapOutput() AdministrativeUnitRoleMemberMapOutput
	ToAdministrativeUnitRoleMemberMapOutputWithContext(context.Context) AdministrativeUnitRoleMemberMapOutput
}

type AdministrativeUnitRoleMemberMap map[string]AdministrativeUnitRoleMemberInput

func (AdministrativeUnitRoleMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AdministrativeUnitRoleMember)(nil)).Elem()
}

func (i AdministrativeUnitRoleMemberMap) ToAdministrativeUnitRoleMemberMapOutput() AdministrativeUnitRoleMemberMapOutput {
	return i.ToAdministrativeUnitRoleMemberMapOutputWithContext(context.Background())
}

func (i AdministrativeUnitRoleMemberMap) ToAdministrativeUnitRoleMemberMapOutputWithContext(ctx context.Context) AdministrativeUnitRoleMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdministrativeUnitRoleMemberMapOutput)
}

func (i AdministrativeUnitRoleMemberMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*AdministrativeUnitRoleMember] {
	return pulumix.Output[map[string]*AdministrativeUnitRoleMember]{
		OutputState: i.ToAdministrativeUnitRoleMemberMapOutputWithContext(ctx).OutputState,
	}
}

type AdministrativeUnitRoleMemberOutput struct{ *pulumi.OutputState }

func (AdministrativeUnitRoleMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AdministrativeUnitRoleMember)(nil)).Elem()
}

func (o AdministrativeUnitRoleMemberOutput) ToAdministrativeUnitRoleMemberOutput() AdministrativeUnitRoleMemberOutput {
	return o
}

func (o AdministrativeUnitRoleMemberOutput) ToAdministrativeUnitRoleMemberOutputWithContext(ctx context.Context) AdministrativeUnitRoleMemberOutput {
	return o
}

func (o AdministrativeUnitRoleMemberOutput) ToOutput(ctx context.Context) pulumix.Output[*AdministrativeUnitRoleMember] {
	return pulumix.Output[*AdministrativeUnitRoleMember]{
		OutputState: o.OutputState,
	}
}

// The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
func (o AdministrativeUnitRoleMemberOutput) AdministrativeUnitObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *AdministrativeUnitRoleMember) pulumi.StringOutput { return v.AdministrativeUnitObjectId }).(pulumi.StringOutput)
}

// The object ID of the user, group or service principal you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
func (o AdministrativeUnitRoleMemberOutput) MemberObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *AdministrativeUnitRoleMember) pulumi.StringOutput { return v.MemberObjectId }).(pulumi.StringOutput)
}

// The object ID of the directory role you want to assign. Changing this forces a new resource to be created.
func (o AdministrativeUnitRoleMemberOutput) RoleObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *AdministrativeUnitRoleMember) pulumi.StringOutput { return v.RoleObjectId }).(pulumi.StringOutput)
}

type AdministrativeUnitRoleMemberArrayOutput struct{ *pulumi.OutputState }

func (AdministrativeUnitRoleMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AdministrativeUnitRoleMember)(nil)).Elem()
}

func (o AdministrativeUnitRoleMemberArrayOutput) ToAdministrativeUnitRoleMemberArrayOutput() AdministrativeUnitRoleMemberArrayOutput {
	return o
}

func (o AdministrativeUnitRoleMemberArrayOutput) ToAdministrativeUnitRoleMemberArrayOutputWithContext(ctx context.Context) AdministrativeUnitRoleMemberArrayOutput {
	return o
}

func (o AdministrativeUnitRoleMemberArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*AdministrativeUnitRoleMember] {
	return pulumix.Output[[]*AdministrativeUnitRoleMember]{
		OutputState: o.OutputState,
	}
}

func (o AdministrativeUnitRoleMemberArrayOutput) Index(i pulumi.IntInput) AdministrativeUnitRoleMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AdministrativeUnitRoleMember {
		return vs[0].([]*AdministrativeUnitRoleMember)[vs[1].(int)]
	}).(AdministrativeUnitRoleMemberOutput)
}

type AdministrativeUnitRoleMemberMapOutput struct{ *pulumi.OutputState }

func (AdministrativeUnitRoleMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AdministrativeUnitRoleMember)(nil)).Elem()
}

func (o AdministrativeUnitRoleMemberMapOutput) ToAdministrativeUnitRoleMemberMapOutput() AdministrativeUnitRoleMemberMapOutput {
	return o
}

func (o AdministrativeUnitRoleMemberMapOutput) ToAdministrativeUnitRoleMemberMapOutputWithContext(ctx context.Context) AdministrativeUnitRoleMemberMapOutput {
	return o
}

func (o AdministrativeUnitRoleMemberMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*AdministrativeUnitRoleMember] {
	return pulumix.Output[map[string]*AdministrativeUnitRoleMember]{
		OutputState: o.OutputState,
	}
}

func (o AdministrativeUnitRoleMemberMapOutput) MapIndex(k pulumi.StringInput) AdministrativeUnitRoleMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AdministrativeUnitRoleMember {
		return vs[0].(map[string]*AdministrativeUnitRoleMember)[vs[1].(string)]
	}).(AdministrativeUnitRoleMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AdministrativeUnitRoleMemberInput)(nil)).Elem(), &AdministrativeUnitRoleMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*AdministrativeUnitRoleMemberArrayInput)(nil)).Elem(), AdministrativeUnitRoleMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AdministrativeUnitRoleMemberMapInput)(nil)).Elem(), AdministrativeUnitRoleMemberMap{})
	pulumi.RegisterOutputType(AdministrativeUnitRoleMemberOutput{})
	pulumi.RegisterOutputType(AdministrativeUnitRoleMemberArrayOutput{})
	pulumi.RegisterOutputType(AdministrativeUnitRoleMemberMapOutput{})
}

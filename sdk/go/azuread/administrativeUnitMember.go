// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuread/sdk/v5/go/azuread/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a single administrative unit membership within Azure Active Directory.
//
// > **Warning** Do not use this resource at the same time as the `members` property of the `AdministrativeUnit` resource for the same administrative unit. Doing so will cause a conflict and administrative unit members will be removed.
//
// ## API Permissions
//
// The following API permissions are required in order to use this resource.
//
// When authenticated with a service principal, this resource requires one of the following application roles: `AdministrativeUnit.ReadWrite.All` or `Directory.ReadWrite.All`
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
//			example, err := azuread.LookupUser(ctx, &azuread.LookupUserArgs{
//				UserPrincipalName: pulumi.StringRef("jdoe@example.com"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleAdministrativeUnit, err := azuread.NewAdministrativeUnit(ctx, "example", &azuread.AdministrativeUnitArgs{
//				DisplayName: pulumi.String("Example-AU"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuread.NewAdministrativeUnitMember(ctx, "example", &azuread.AdministrativeUnitMemberArgs{
//				AdministrativeUnitObjectId: exampleAdministrativeUnit.ID(),
//				MemberObjectId:             pulumi.String(example.Id),
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
// Administrative unit members can be imported using the object ID of the administrative unit and the object ID of the member, e.g.
//
// ```sh
// $ pulumi import azuread:index/administrativeUnitMember:AdministrativeUnitMember example 00000000-0000-0000-0000-000000000000/member/11111111-1111-1111-1111-111111111111
// ```
//
// -> This ID format is unique to Terraform and is composed of the Administrative Unit Object ID and the target Member Object ID in the format `{AdministrativeUnitObjectID}/member/{MemberObjectID}`.
type AdministrativeUnitMember struct {
	pulumi.CustomResourceState

	// The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
	AdministrativeUnitObjectId pulumi.StringPtrOutput `pulumi:"administrativeUnitObjectId"`
	// The object ID of the user or group you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
	MemberObjectId pulumi.StringPtrOutput `pulumi:"memberObjectId"`
}

// NewAdministrativeUnitMember registers a new resource with the given unique name, arguments, and options.
func NewAdministrativeUnitMember(ctx *pulumi.Context,
	name string, args *AdministrativeUnitMemberArgs, opts ...pulumi.ResourceOption) (*AdministrativeUnitMember, error) {
	if args == nil {
		args = &AdministrativeUnitMemberArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AdministrativeUnitMember
	err := ctx.RegisterResource("azuread:index/administrativeUnitMember:AdministrativeUnitMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAdministrativeUnitMember gets an existing AdministrativeUnitMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAdministrativeUnitMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AdministrativeUnitMemberState, opts ...pulumi.ResourceOption) (*AdministrativeUnitMember, error) {
	var resource AdministrativeUnitMember
	err := ctx.ReadResource("azuread:index/administrativeUnitMember:AdministrativeUnitMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AdministrativeUnitMember resources.
type administrativeUnitMemberState struct {
	// The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
	AdministrativeUnitObjectId *string `pulumi:"administrativeUnitObjectId"`
	// The object ID of the user or group you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
	MemberObjectId *string `pulumi:"memberObjectId"`
}

type AdministrativeUnitMemberState struct {
	// The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
	AdministrativeUnitObjectId pulumi.StringPtrInput
	// The object ID of the user or group you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
	MemberObjectId pulumi.StringPtrInput
}

func (AdministrativeUnitMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*administrativeUnitMemberState)(nil)).Elem()
}

type administrativeUnitMemberArgs struct {
	// The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
	AdministrativeUnitObjectId *string `pulumi:"administrativeUnitObjectId"`
	// The object ID of the user or group you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
	MemberObjectId *string `pulumi:"memberObjectId"`
}

// The set of arguments for constructing a AdministrativeUnitMember resource.
type AdministrativeUnitMemberArgs struct {
	// The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
	AdministrativeUnitObjectId pulumi.StringPtrInput
	// The object ID of the user or group you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
	MemberObjectId pulumi.StringPtrInput
}

func (AdministrativeUnitMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*administrativeUnitMemberArgs)(nil)).Elem()
}

type AdministrativeUnitMemberInput interface {
	pulumi.Input

	ToAdministrativeUnitMemberOutput() AdministrativeUnitMemberOutput
	ToAdministrativeUnitMemberOutputWithContext(ctx context.Context) AdministrativeUnitMemberOutput
}

func (*AdministrativeUnitMember) ElementType() reflect.Type {
	return reflect.TypeOf((**AdministrativeUnitMember)(nil)).Elem()
}

func (i *AdministrativeUnitMember) ToAdministrativeUnitMemberOutput() AdministrativeUnitMemberOutput {
	return i.ToAdministrativeUnitMemberOutputWithContext(context.Background())
}

func (i *AdministrativeUnitMember) ToAdministrativeUnitMemberOutputWithContext(ctx context.Context) AdministrativeUnitMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdministrativeUnitMemberOutput)
}

// AdministrativeUnitMemberArrayInput is an input type that accepts AdministrativeUnitMemberArray and AdministrativeUnitMemberArrayOutput values.
// You can construct a concrete instance of `AdministrativeUnitMemberArrayInput` via:
//
//	AdministrativeUnitMemberArray{ AdministrativeUnitMemberArgs{...} }
type AdministrativeUnitMemberArrayInput interface {
	pulumi.Input

	ToAdministrativeUnitMemberArrayOutput() AdministrativeUnitMemberArrayOutput
	ToAdministrativeUnitMemberArrayOutputWithContext(context.Context) AdministrativeUnitMemberArrayOutput
}

type AdministrativeUnitMemberArray []AdministrativeUnitMemberInput

func (AdministrativeUnitMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AdministrativeUnitMember)(nil)).Elem()
}

func (i AdministrativeUnitMemberArray) ToAdministrativeUnitMemberArrayOutput() AdministrativeUnitMemberArrayOutput {
	return i.ToAdministrativeUnitMemberArrayOutputWithContext(context.Background())
}

func (i AdministrativeUnitMemberArray) ToAdministrativeUnitMemberArrayOutputWithContext(ctx context.Context) AdministrativeUnitMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdministrativeUnitMemberArrayOutput)
}

// AdministrativeUnitMemberMapInput is an input type that accepts AdministrativeUnitMemberMap and AdministrativeUnitMemberMapOutput values.
// You can construct a concrete instance of `AdministrativeUnitMemberMapInput` via:
//
//	AdministrativeUnitMemberMap{ "key": AdministrativeUnitMemberArgs{...} }
type AdministrativeUnitMemberMapInput interface {
	pulumi.Input

	ToAdministrativeUnitMemberMapOutput() AdministrativeUnitMemberMapOutput
	ToAdministrativeUnitMemberMapOutputWithContext(context.Context) AdministrativeUnitMemberMapOutput
}

type AdministrativeUnitMemberMap map[string]AdministrativeUnitMemberInput

func (AdministrativeUnitMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AdministrativeUnitMember)(nil)).Elem()
}

func (i AdministrativeUnitMemberMap) ToAdministrativeUnitMemberMapOutput() AdministrativeUnitMemberMapOutput {
	return i.ToAdministrativeUnitMemberMapOutputWithContext(context.Background())
}

func (i AdministrativeUnitMemberMap) ToAdministrativeUnitMemberMapOutputWithContext(ctx context.Context) AdministrativeUnitMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdministrativeUnitMemberMapOutput)
}

type AdministrativeUnitMemberOutput struct{ *pulumi.OutputState }

func (AdministrativeUnitMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AdministrativeUnitMember)(nil)).Elem()
}

func (o AdministrativeUnitMemberOutput) ToAdministrativeUnitMemberOutput() AdministrativeUnitMemberOutput {
	return o
}

func (o AdministrativeUnitMemberOutput) ToAdministrativeUnitMemberOutputWithContext(ctx context.Context) AdministrativeUnitMemberOutput {
	return o
}

// The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
func (o AdministrativeUnitMemberOutput) AdministrativeUnitObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AdministrativeUnitMember) pulumi.StringPtrOutput { return v.AdministrativeUnitObjectId }).(pulumi.StringPtrOutput)
}

// The object ID of the user or group you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
func (o AdministrativeUnitMemberOutput) MemberObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AdministrativeUnitMember) pulumi.StringPtrOutput { return v.MemberObjectId }).(pulumi.StringPtrOutput)
}

type AdministrativeUnitMemberArrayOutput struct{ *pulumi.OutputState }

func (AdministrativeUnitMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AdministrativeUnitMember)(nil)).Elem()
}

func (o AdministrativeUnitMemberArrayOutput) ToAdministrativeUnitMemberArrayOutput() AdministrativeUnitMemberArrayOutput {
	return o
}

func (o AdministrativeUnitMemberArrayOutput) ToAdministrativeUnitMemberArrayOutputWithContext(ctx context.Context) AdministrativeUnitMemberArrayOutput {
	return o
}

func (o AdministrativeUnitMemberArrayOutput) Index(i pulumi.IntInput) AdministrativeUnitMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AdministrativeUnitMember {
		return vs[0].([]*AdministrativeUnitMember)[vs[1].(int)]
	}).(AdministrativeUnitMemberOutput)
}

type AdministrativeUnitMemberMapOutput struct{ *pulumi.OutputState }

func (AdministrativeUnitMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AdministrativeUnitMember)(nil)).Elem()
}

func (o AdministrativeUnitMemberMapOutput) ToAdministrativeUnitMemberMapOutput() AdministrativeUnitMemberMapOutput {
	return o
}

func (o AdministrativeUnitMemberMapOutput) ToAdministrativeUnitMemberMapOutputWithContext(ctx context.Context) AdministrativeUnitMemberMapOutput {
	return o
}

func (o AdministrativeUnitMemberMapOutput) MapIndex(k pulumi.StringInput) AdministrativeUnitMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AdministrativeUnitMember {
		return vs[0].(map[string]*AdministrativeUnitMember)[vs[1].(string)]
	}).(AdministrativeUnitMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AdministrativeUnitMemberInput)(nil)).Elem(), &AdministrativeUnitMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*AdministrativeUnitMemberArrayInput)(nil)).Elem(), AdministrativeUnitMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AdministrativeUnitMemberMapInput)(nil)).Elem(), AdministrativeUnitMemberMap{})
	pulumi.RegisterOutputType(AdministrativeUnitMemberOutput{})
	pulumi.RegisterOutputType(AdministrativeUnitMemberArrayOutput{})
	pulumi.RegisterOutputType(AdministrativeUnitMemberMapOutput{})
}

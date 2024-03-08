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

// Manages an Administrative Unit within Azure Active Directory.
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
//			_, err := azuread.NewAdministrativeUnit(ctx, "example", &azuread.AdministrativeUnitArgs{
//				DisplayName:             pulumi.String("Example-AU"),
//				Description:             pulumi.String("Just an example"),
//				HiddenMembershipEnabled: pulumi.Bool(false),
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
// Administrative units can be imported using their object ID, e.g.
//
// ```sh
// $ pulumi import azuread:index/administrativeUnit:AdministrativeUnit example 00000000-0000-0000-0000-000000000000
// ```
type AdministrativeUnit struct {
	pulumi.CustomResourceState

	// The description of the administrative unit.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name of the administrative unit.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Whether the administrative unit and its members are hidden or publicly viewable in the directory.
	HiddenMembershipEnabled pulumi.BoolPtrOutput `pulumi:"hiddenMembershipEnabled"`
	// A set of object IDs of members who should be present in this administrative unit. Supported object types are Users or Groups.
	//
	// !> **Warning** Do not use the `members` property at the same time as the AdministrativeUnitMember resource for the same administrative unit. Doing so will cause a conflict and administrative unit members will be removed.
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The object ID of the administrative unit.
	ObjectId pulumi.StringOutput `pulumi:"objectId"`
	// If `true`, will return an error if an existing administrative unit is found with the same name
	PreventDuplicateNames pulumi.BoolPtrOutput `pulumi:"preventDuplicateNames"`
}

// NewAdministrativeUnit registers a new resource with the given unique name, arguments, and options.
func NewAdministrativeUnit(ctx *pulumi.Context,
	name string, args *AdministrativeUnitArgs, opts ...pulumi.ResourceOption) (*AdministrativeUnit, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AdministrativeUnit
	err := ctx.RegisterResource("azuread:index/administrativeUnit:AdministrativeUnit", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAdministrativeUnit gets an existing AdministrativeUnit resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAdministrativeUnit(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AdministrativeUnitState, opts ...pulumi.ResourceOption) (*AdministrativeUnit, error) {
	var resource AdministrativeUnit
	err := ctx.ReadResource("azuread:index/administrativeUnit:AdministrativeUnit", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AdministrativeUnit resources.
type administrativeUnitState struct {
	// The description of the administrative unit.
	Description *string `pulumi:"description"`
	// The display name of the administrative unit.
	DisplayName *string `pulumi:"displayName"`
	// Whether the administrative unit and its members are hidden or publicly viewable in the directory.
	HiddenMembershipEnabled *bool `pulumi:"hiddenMembershipEnabled"`
	// A set of object IDs of members who should be present in this administrative unit. Supported object types are Users or Groups.
	//
	// !> **Warning** Do not use the `members` property at the same time as the AdministrativeUnitMember resource for the same administrative unit. Doing so will cause a conflict and administrative unit members will be removed.
	Members []string `pulumi:"members"`
	// The object ID of the administrative unit.
	ObjectId *string `pulumi:"objectId"`
	// If `true`, will return an error if an existing administrative unit is found with the same name
	PreventDuplicateNames *bool `pulumi:"preventDuplicateNames"`
}

type AdministrativeUnitState struct {
	// The description of the administrative unit.
	Description pulumi.StringPtrInput
	// The display name of the administrative unit.
	DisplayName pulumi.StringPtrInput
	// Whether the administrative unit and its members are hidden or publicly viewable in the directory.
	HiddenMembershipEnabled pulumi.BoolPtrInput
	// A set of object IDs of members who should be present in this administrative unit. Supported object types are Users or Groups.
	//
	// !> **Warning** Do not use the `members` property at the same time as the AdministrativeUnitMember resource for the same administrative unit. Doing so will cause a conflict and administrative unit members will be removed.
	Members pulumi.StringArrayInput
	// The object ID of the administrative unit.
	ObjectId pulumi.StringPtrInput
	// If `true`, will return an error if an existing administrative unit is found with the same name
	PreventDuplicateNames pulumi.BoolPtrInput
}

func (AdministrativeUnitState) ElementType() reflect.Type {
	return reflect.TypeOf((*administrativeUnitState)(nil)).Elem()
}

type administrativeUnitArgs struct {
	// The description of the administrative unit.
	Description *string `pulumi:"description"`
	// The display name of the administrative unit.
	DisplayName string `pulumi:"displayName"`
	// Whether the administrative unit and its members are hidden or publicly viewable in the directory.
	HiddenMembershipEnabled *bool `pulumi:"hiddenMembershipEnabled"`
	// A set of object IDs of members who should be present in this administrative unit. Supported object types are Users or Groups.
	//
	// !> **Warning** Do not use the `members` property at the same time as the AdministrativeUnitMember resource for the same administrative unit. Doing so will cause a conflict and administrative unit members will be removed.
	Members []string `pulumi:"members"`
	// If `true`, will return an error if an existing administrative unit is found with the same name
	PreventDuplicateNames *bool `pulumi:"preventDuplicateNames"`
}

// The set of arguments for constructing a AdministrativeUnit resource.
type AdministrativeUnitArgs struct {
	// The description of the administrative unit.
	Description pulumi.StringPtrInput
	// The display name of the administrative unit.
	DisplayName pulumi.StringInput
	// Whether the administrative unit and its members are hidden or publicly viewable in the directory.
	HiddenMembershipEnabled pulumi.BoolPtrInput
	// A set of object IDs of members who should be present in this administrative unit. Supported object types are Users or Groups.
	//
	// !> **Warning** Do not use the `members` property at the same time as the AdministrativeUnitMember resource for the same administrative unit. Doing so will cause a conflict and administrative unit members will be removed.
	Members pulumi.StringArrayInput
	// If `true`, will return an error if an existing administrative unit is found with the same name
	PreventDuplicateNames pulumi.BoolPtrInput
}

func (AdministrativeUnitArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*administrativeUnitArgs)(nil)).Elem()
}

type AdministrativeUnitInput interface {
	pulumi.Input

	ToAdministrativeUnitOutput() AdministrativeUnitOutput
	ToAdministrativeUnitOutputWithContext(ctx context.Context) AdministrativeUnitOutput
}

func (*AdministrativeUnit) ElementType() reflect.Type {
	return reflect.TypeOf((**AdministrativeUnit)(nil)).Elem()
}

func (i *AdministrativeUnit) ToAdministrativeUnitOutput() AdministrativeUnitOutput {
	return i.ToAdministrativeUnitOutputWithContext(context.Background())
}

func (i *AdministrativeUnit) ToAdministrativeUnitOutputWithContext(ctx context.Context) AdministrativeUnitOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdministrativeUnitOutput)
}

// AdministrativeUnitArrayInput is an input type that accepts AdministrativeUnitArray and AdministrativeUnitArrayOutput values.
// You can construct a concrete instance of `AdministrativeUnitArrayInput` via:
//
//	AdministrativeUnitArray{ AdministrativeUnitArgs{...} }
type AdministrativeUnitArrayInput interface {
	pulumi.Input

	ToAdministrativeUnitArrayOutput() AdministrativeUnitArrayOutput
	ToAdministrativeUnitArrayOutputWithContext(context.Context) AdministrativeUnitArrayOutput
}

type AdministrativeUnitArray []AdministrativeUnitInput

func (AdministrativeUnitArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AdministrativeUnit)(nil)).Elem()
}

func (i AdministrativeUnitArray) ToAdministrativeUnitArrayOutput() AdministrativeUnitArrayOutput {
	return i.ToAdministrativeUnitArrayOutputWithContext(context.Background())
}

func (i AdministrativeUnitArray) ToAdministrativeUnitArrayOutputWithContext(ctx context.Context) AdministrativeUnitArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdministrativeUnitArrayOutput)
}

// AdministrativeUnitMapInput is an input type that accepts AdministrativeUnitMap and AdministrativeUnitMapOutput values.
// You can construct a concrete instance of `AdministrativeUnitMapInput` via:
//
//	AdministrativeUnitMap{ "key": AdministrativeUnitArgs{...} }
type AdministrativeUnitMapInput interface {
	pulumi.Input

	ToAdministrativeUnitMapOutput() AdministrativeUnitMapOutput
	ToAdministrativeUnitMapOutputWithContext(context.Context) AdministrativeUnitMapOutput
}

type AdministrativeUnitMap map[string]AdministrativeUnitInput

func (AdministrativeUnitMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AdministrativeUnit)(nil)).Elem()
}

func (i AdministrativeUnitMap) ToAdministrativeUnitMapOutput() AdministrativeUnitMapOutput {
	return i.ToAdministrativeUnitMapOutputWithContext(context.Background())
}

func (i AdministrativeUnitMap) ToAdministrativeUnitMapOutputWithContext(ctx context.Context) AdministrativeUnitMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdministrativeUnitMapOutput)
}

type AdministrativeUnitOutput struct{ *pulumi.OutputState }

func (AdministrativeUnitOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AdministrativeUnit)(nil)).Elem()
}

func (o AdministrativeUnitOutput) ToAdministrativeUnitOutput() AdministrativeUnitOutput {
	return o
}

func (o AdministrativeUnitOutput) ToAdministrativeUnitOutputWithContext(ctx context.Context) AdministrativeUnitOutput {
	return o
}

// The description of the administrative unit.
func (o AdministrativeUnitOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AdministrativeUnit) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name of the administrative unit.
func (o AdministrativeUnitOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *AdministrativeUnit) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Whether the administrative unit and its members are hidden or publicly viewable in the directory.
func (o AdministrativeUnitOutput) HiddenMembershipEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AdministrativeUnit) pulumi.BoolPtrOutput { return v.HiddenMembershipEnabled }).(pulumi.BoolPtrOutput)
}

// A set of object IDs of members who should be present in this administrative unit. Supported object types are Users or Groups.
//
// !> **Warning** Do not use the `members` property at the same time as the AdministrativeUnitMember resource for the same administrative unit. Doing so will cause a conflict and administrative unit members will be removed.
func (o AdministrativeUnitOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AdministrativeUnit) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The object ID of the administrative unit.
func (o AdministrativeUnitOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *AdministrativeUnit) pulumi.StringOutput { return v.ObjectId }).(pulumi.StringOutput)
}

// If `true`, will return an error if an existing administrative unit is found with the same name
func (o AdministrativeUnitOutput) PreventDuplicateNames() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AdministrativeUnit) pulumi.BoolPtrOutput { return v.PreventDuplicateNames }).(pulumi.BoolPtrOutput)
}

type AdministrativeUnitArrayOutput struct{ *pulumi.OutputState }

func (AdministrativeUnitArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AdministrativeUnit)(nil)).Elem()
}

func (o AdministrativeUnitArrayOutput) ToAdministrativeUnitArrayOutput() AdministrativeUnitArrayOutput {
	return o
}

func (o AdministrativeUnitArrayOutput) ToAdministrativeUnitArrayOutputWithContext(ctx context.Context) AdministrativeUnitArrayOutput {
	return o
}

func (o AdministrativeUnitArrayOutput) Index(i pulumi.IntInput) AdministrativeUnitOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AdministrativeUnit {
		return vs[0].([]*AdministrativeUnit)[vs[1].(int)]
	}).(AdministrativeUnitOutput)
}

type AdministrativeUnitMapOutput struct{ *pulumi.OutputState }

func (AdministrativeUnitMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AdministrativeUnit)(nil)).Elem()
}

func (o AdministrativeUnitMapOutput) ToAdministrativeUnitMapOutput() AdministrativeUnitMapOutput {
	return o
}

func (o AdministrativeUnitMapOutput) ToAdministrativeUnitMapOutputWithContext(ctx context.Context) AdministrativeUnitMapOutput {
	return o
}

func (o AdministrativeUnitMapOutput) MapIndex(k pulumi.StringInput) AdministrativeUnitOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AdministrativeUnit {
		return vs[0].(map[string]*AdministrativeUnit)[vs[1].(string)]
	}).(AdministrativeUnitOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AdministrativeUnitInput)(nil)).Elem(), &AdministrativeUnit{})
	pulumi.RegisterInputType(reflect.TypeOf((*AdministrativeUnitArrayInput)(nil)).Elem(), AdministrativeUnitArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AdministrativeUnitMapInput)(nil)).Elem(), AdministrativeUnitMap{})
	pulumi.RegisterOutputType(AdministrativeUnitOutput{})
	pulumi.RegisterOutputType(AdministrativeUnitArrayOutput{})
	pulumi.RegisterOutputType(AdministrativeUnitMapOutput{})
}

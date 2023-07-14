// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccessPackageCatalogRoleAssignment struct {
	pulumi.CustomResourceState

	// The unique ID of the access package catalog.
	CatalogId pulumi.StringOutput `pulumi:"catalogId"`
	// The object ID of the member principal
	PrincipalObjectId pulumi.StringOutput `pulumi:"principalObjectId"`
	// The object ID of the catalog role for this assignment
	RoleId pulumi.StringOutput `pulumi:"roleId"`
}

// NewAccessPackageCatalogRoleAssignment registers a new resource with the given unique name, arguments, and options.
func NewAccessPackageCatalogRoleAssignment(ctx *pulumi.Context,
	name string, args *AccessPackageCatalogRoleAssignmentArgs, opts ...pulumi.ResourceOption) (*AccessPackageCatalogRoleAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CatalogId == nil {
		return nil, errors.New("invalid value for required argument 'CatalogId'")
	}
	if args.PrincipalObjectId == nil {
		return nil, errors.New("invalid value for required argument 'PrincipalObjectId'")
	}
	if args.RoleId == nil {
		return nil, errors.New("invalid value for required argument 'RoleId'")
	}
	var resource AccessPackageCatalogRoleAssignment
	err := ctx.RegisterResource("azuread:index/accessPackageCatalogRoleAssignment:AccessPackageCatalogRoleAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessPackageCatalogRoleAssignment gets an existing AccessPackageCatalogRoleAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessPackageCatalogRoleAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessPackageCatalogRoleAssignmentState, opts ...pulumi.ResourceOption) (*AccessPackageCatalogRoleAssignment, error) {
	var resource AccessPackageCatalogRoleAssignment
	err := ctx.ReadResource("azuread:index/accessPackageCatalogRoleAssignment:AccessPackageCatalogRoleAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessPackageCatalogRoleAssignment resources.
type accessPackageCatalogRoleAssignmentState struct {
	// The unique ID of the access package catalog.
	CatalogId *string `pulumi:"catalogId"`
	// The object ID of the member principal
	PrincipalObjectId *string `pulumi:"principalObjectId"`
	// The object ID of the catalog role for this assignment
	RoleId *string `pulumi:"roleId"`
}

type AccessPackageCatalogRoleAssignmentState struct {
	// The unique ID of the access package catalog.
	CatalogId pulumi.StringPtrInput
	// The object ID of the member principal
	PrincipalObjectId pulumi.StringPtrInput
	// The object ID of the catalog role for this assignment
	RoleId pulumi.StringPtrInput
}

func (AccessPackageCatalogRoleAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPackageCatalogRoleAssignmentState)(nil)).Elem()
}

type accessPackageCatalogRoleAssignmentArgs struct {
	// The unique ID of the access package catalog.
	CatalogId string `pulumi:"catalogId"`
	// The object ID of the member principal
	PrincipalObjectId string `pulumi:"principalObjectId"`
	// The object ID of the catalog role for this assignment
	RoleId string `pulumi:"roleId"`
}

// The set of arguments for constructing a AccessPackageCatalogRoleAssignment resource.
type AccessPackageCatalogRoleAssignmentArgs struct {
	// The unique ID of the access package catalog.
	CatalogId pulumi.StringInput
	// The object ID of the member principal
	PrincipalObjectId pulumi.StringInput
	// The object ID of the catalog role for this assignment
	RoleId pulumi.StringInput
}

func (AccessPackageCatalogRoleAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPackageCatalogRoleAssignmentArgs)(nil)).Elem()
}

type AccessPackageCatalogRoleAssignmentInput interface {
	pulumi.Input

	ToAccessPackageCatalogRoleAssignmentOutput() AccessPackageCatalogRoleAssignmentOutput
	ToAccessPackageCatalogRoleAssignmentOutputWithContext(ctx context.Context) AccessPackageCatalogRoleAssignmentOutput
}

func (*AccessPackageCatalogRoleAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessPackageCatalogRoleAssignment)(nil)).Elem()
}

func (i *AccessPackageCatalogRoleAssignment) ToAccessPackageCatalogRoleAssignmentOutput() AccessPackageCatalogRoleAssignmentOutput {
	return i.ToAccessPackageCatalogRoleAssignmentOutputWithContext(context.Background())
}

func (i *AccessPackageCatalogRoleAssignment) ToAccessPackageCatalogRoleAssignmentOutputWithContext(ctx context.Context) AccessPackageCatalogRoleAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPackageCatalogRoleAssignmentOutput)
}

// AccessPackageCatalogRoleAssignmentArrayInput is an input type that accepts AccessPackageCatalogRoleAssignmentArray and AccessPackageCatalogRoleAssignmentArrayOutput values.
// You can construct a concrete instance of `AccessPackageCatalogRoleAssignmentArrayInput` via:
//
//	AccessPackageCatalogRoleAssignmentArray{ AccessPackageCatalogRoleAssignmentArgs{...} }
type AccessPackageCatalogRoleAssignmentArrayInput interface {
	pulumi.Input

	ToAccessPackageCatalogRoleAssignmentArrayOutput() AccessPackageCatalogRoleAssignmentArrayOutput
	ToAccessPackageCatalogRoleAssignmentArrayOutputWithContext(context.Context) AccessPackageCatalogRoleAssignmentArrayOutput
}

type AccessPackageCatalogRoleAssignmentArray []AccessPackageCatalogRoleAssignmentInput

func (AccessPackageCatalogRoleAssignmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessPackageCatalogRoleAssignment)(nil)).Elem()
}

func (i AccessPackageCatalogRoleAssignmentArray) ToAccessPackageCatalogRoleAssignmentArrayOutput() AccessPackageCatalogRoleAssignmentArrayOutput {
	return i.ToAccessPackageCatalogRoleAssignmentArrayOutputWithContext(context.Background())
}

func (i AccessPackageCatalogRoleAssignmentArray) ToAccessPackageCatalogRoleAssignmentArrayOutputWithContext(ctx context.Context) AccessPackageCatalogRoleAssignmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPackageCatalogRoleAssignmentArrayOutput)
}

// AccessPackageCatalogRoleAssignmentMapInput is an input type that accepts AccessPackageCatalogRoleAssignmentMap and AccessPackageCatalogRoleAssignmentMapOutput values.
// You can construct a concrete instance of `AccessPackageCatalogRoleAssignmentMapInput` via:
//
//	AccessPackageCatalogRoleAssignmentMap{ "key": AccessPackageCatalogRoleAssignmentArgs{...} }
type AccessPackageCatalogRoleAssignmentMapInput interface {
	pulumi.Input

	ToAccessPackageCatalogRoleAssignmentMapOutput() AccessPackageCatalogRoleAssignmentMapOutput
	ToAccessPackageCatalogRoleAssignmentMapOutputWithContext(context.Context) AccessPackageCatalogRoleAssignmentMapOutput
}

type AccessPackageCatalogRoleAssignmentMap map[string]AccessPackageCatalogRoleAssignmentInput

func (AccessPackageCatalogRoleAssignmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessPackageCatalogRoleAssignment)(nil)).Elem()
}

func (i AccessPackageCatalogRoleAssignmentMap) ToAccessPackageCatalogRoleAssignmentMapOutput() AccessPackageCatalogRoleAssignmentMapOutput {
	return i.ToAccessPackageCatalogRoleAssignmentMapOutputWithContext(context.Background())
}

func (i AccessPackageCatalogRoleAssignmentMap) ToAccessPackageCatalogRoleAssignmentMapOutputWithContext(ctx context.Context) AccessPackageCatalogRoleAssignmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPackageCatalogRoleAssignmentMapOutput)
}

type AccessPackageCatalogRoleAssignmentOutput struct{ *pulumi.OutputState }

func (AccessPackageCatalogRoleAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessPackageCatalogRoleAssignment)(nil)).Elem()
}

func (o AccessPackageCatalogRoleAssignmentOutput) ToAccessPackageCatalogRoleAssignmentOutput() AccessPackageCatalogRoleAssignmentOutput {
	return o
}

func (o AccessPackageCatalogRoleAssignmentOutput) ToAccessPackageCatalogRoleAssignmentOutputWithContext(ctx context.Context) AccessPackageCatalogRoleAssignmentOutput {
	return o
}

// The unique ID of the access package catalog.
func (o AccessPackageCatalogRoleAssignmentOutput) CatalogId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPackageCatalogRoleAssignment) pulumi.StringOutput { return v.CatalogId }).(pulumi.StringOutput)
}

// The object ID of the member principal
func (o AccessPackageCatalogRoleAssignmentOutput) PrincipalObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPackageCatalogRoleAssignment) pulumi.StringOutput { return v.PrincipalObjectId }).(pulumi.StringOutput)
}

// The object ID of the catalog role for this assignment
func (o AccessPackageCatalogRoleAssignmentOutput) RoleId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPackageCatalogRoleAssignment) pulumi.StringOutput { return v.RoleId }).(pulumi.StringOutput)
}

type AccessPackageCatalogRoleAssignmentArrayOutput struct{ *pulumi.OutputState }

func (AccessPackageCatalogRoleAssignmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessPackageCatalogRoleAssignment)(nil)).Elem()
}

func (o AccessPackageCatalogRoleAssignmentArrayOutput) ToAccessPackageCatalogRoleAssignmentArrayOutput() AccessPackageCatalogRoleAssignmentArrayOutput {
	return o
}

func (o AccessPackageCatalogRoleAssignmentArrayOutput) ToAccessPackageCatalogRoleAssignmentArrayOutputWithContext(ctx context.Context) AccessPackageCatalogRoleAssignmentArrayOutput {
	return o
}

func (o AccessPackageCatalogRoleAssignmentArrayOutput) Index(i pulumi.IntInput) AccessPackageCatalogRoleAssignmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AccessPackageCatalogRoleAssignment {
		return vs[0].([]*AccessPackageCatalogRoleAssignment)[vs[1].(int)]
	}).(AccessPackageCatalogRoleAssignmentOutput)
}

type AccessPackageCatalogRoleAssignmentMapOutput struct{ *pulumi.OutputState }

func (AccessPackageCatalogRoleAssignmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessPackageCatalogRoleAssignment)(nil)).Elem()
}

func (o AccessPackageCatalogRoleAssignmentMapOutput) ToAccessPackageCatalogRoleAssignmentMapOutput() AccessPackageCatalogRoleAssignmentMapOutput {
	return o
}

func (o AccessPackageCatalogRoleAssignmentMapOutput) ToAccessPackageCatalogRoleAssignmentMapOutputWithContext(ctx context.Context) AccessPackageCatalogRoleAssignmentMapOutput {
	return o
}

func (o AccessPackageCatalogRoleAssignmentMapOutput) MapIndex(k pulumi.StringInput) AccessPackageCatalogRoleAssignmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AccessPackageCatalogRoleAssignment {
		return vs[0].(map[string]*AccessPackageCatalogRoleAssignment)[vs[1].(string)]
	}).(AccessPackageCatalogRoleAssignmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessPackageCatalogRoleAssignmentInput)(nil)).Elem(), &AccessPackageCatalogRoleAssignment{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessPackageCatalogRoleAssignmentArrayInput)(nil)).Elem(), AccessPackageCatalogRoleAssignmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessPackageCatalogRoleAssignmentMapInput)(nil)).Elem(), AccessPackageCatalogRoleAssignmentMap{})
	pulumi.RegisterOutputType(AccessPackageCatalogRoleAssignmentOutput{})
	pulumi.RegisterOutputType(AccessPackageCatalogRoleAssignmentArrayOutput{})
	pulumi.RegisterOutputType(AccessPackageCatalogRoleAssignmentMapOutput{})
}

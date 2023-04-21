// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages user flow attributes in an Azure Active Directory (Azure AD) tenant.
//
// ## API Permissions
//
// The following API permissions are required in order to use this resource.
//
// When authenticated with a service principal, this resource requires the following application role: `IdentityUserFlow.ReadWrite.All`
//
// ## Example Usage
//
// *Basic example*
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
//			_, err := azuread.NewUserFlowAttribute(ctx, "example", &azuread.UserFlowAttributeArgs{
//				DataType:    pulumi.String("string"),
//				Description: pulumi.String("Your hobby"),
//				DisplayName: pulumi.String("Hobby"),
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
// User flow attributes can be imported using the `id`, e.g.
//
// ```sh
//
//	$ pulumi import azuread:index/userFlowAttribute:UserFlowAttribute example extension_ecc9f88db2924942b8a96f44873616fe_Hobbyjkorv
//
// ```
//
//	-> This ID can be queried using the [User Flow Attributes API](https://learn.microsoft.com/en-us/graph/api/identityuserflowattribute-list?view=graph-rest-1.0&tabs=http).
type UserFlowAttribute struct {
	pulumi.CustomResourceState

	// The type of the user flow attribute. Values include `builtIn`, `custom` or `required`.
	AttributeType pulumi.StringOutput `pulumi:"attributeType"`
	// The data type of the user flow attribute. Possible values are `boolean`, `dateTime`, `int64`, `string` or `stringCollection`. Changing this forces a new resource to be created.
	DataType pulumi.StringOutput `pulumi:"dataType"`
	// The description of the user flow attribute that is shown to the user at the time of sign-up.
	Description pulumi.StringOutput `pulumi:"description"`
	// The display name of the user flow attribute. Changing this forces a new resource to be created.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
}

// NewUserFlowAttribute registers a new resource with the given unique name, arguments, and options.
func NewUserFlowAttribute(ctx *pulumi.Context,
	name string, args *UserFlowAttributeArgs, opts ...pulumi.ResourceOption) (*UserFlowAttribute, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataType == nil {
		return nil, errors.New("invalid value for required argument 'DataType'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	var resource UserFlowAttribute
	err := ctx.RegisterResource("azuread:index/userFlowAttribute:UserFlowAttribute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserFlowAttribute gets an existing UserFlowAttribute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserFlowAttribute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserFlowAttributeState, opts ...pulumi.ResourceOption) (*UserFlowAttribute, error) {
	var resource UserFlowAttribute
	err := ctx.ReadResource("azuread:index/userFlowAttribute:UserFlowAttribute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserFlowAttribute resources.
type userFlowAttributeState struct {
	// The type of the user flow attribute. Values include `builtIn`, `custom` or `required`.
	AttributeType *string `pulumi:"attributeType"`
	// The data type of the user flow attribute. Possible values are `boolean`, `dateTime`, `int64`, `string` or `stringCollection`. Changing this forces a new resource to be created.
	DataType *string `pulumi:"dataType"`
	// The description of the user flow attribute that is shown to the user at the time of sign-up.
	Description *string `pulumi:"description"`
	// The display name of the user flow attribute. Changing this forces a new resource to be created.
	DisplayName *string `pulumi:"displayName"`
}

type UserFlowAttributeState struct {
	// The type of the user flow attribute. Values include `builtIn`, `custom` or `required`.
	AttributeType pulumi.StringPtrInput
	// The data type of the user flow attribute. Possible values are `boolean`, `dateTime`, `int64`, `string` or `stringCollection`. Changing this forces a new resource to be created.
	DataType pulumi.StringPtrInput
	// The description of the user flow attribute that is shown to the user at the time of sign-up.
	Description pulumi.StringPtrInput
	// The display name of the user flow attribute. Changing this forces a new resource to be created.
	DisplayName pulumi.StringPtrInput
}

func (UserFlowAttributeState) ElementType() reflect.Type {
	return reflect.TypeOf((*userFlowAttributeState)(nil)).Elem()
}

type userFlowAttributeArgs struct {
	// The data type of the user flow attribute. Possible values are `boolean`, `dateTime`, `int64`, `string` or `stringCollection`. Changing this forces a new resource to be created.
	DataType string `pulumi:"dataType"`
	// The description of the user flow attribute that is shown to the user at the time of sign-up.
	Description string `pulumi:"description"`
	// The display name of the user flow attribute. Changing this forces a new resource to be created.
	DisplayName string `pulumi:"displayName"`
}

// The set of arguments for constructing a UserFlowAttribute resource.
type UserFlowAttributeArgs struct {
	// The data type of the user flow attribute. Possible values are `boolean`, `dateTime`, `int64`, `string` or `stringCollection`. Changing this forces a new resource to be created.
	DataType pulumi.StringInput
	// The description of the user flow attribute that is shown to the user at the time of sign-up.
	Description pulumi.StringInput
	// The display name of the user flow attribute. Changing this forces a new resource to be created.
	DisplayName pulumi.StringInput
}

func (UserFlowAttributeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userFlowAttributeArgs)(nil)).Elem()
}

type UserFlowAttributeInput interface {
	pulumi.Input

	ToUserFlowAttributeOutput() UserFlowAttributeOutput
	ToUserFlowAttributeOutputWithContext(ctx context.Context) UserFlowAttributeOutput
}

func (*UserFlowAttribute) ElementType() reflect.Type {
	return reflect.TypeOf((**UserFlowAttribute)(nil)).Elem()
}

func (i *UserFlowAttribute) ToUserFlowAttributeOutput() UserFlowAttributeOutput {
	return i.ToUserFlowAttributeOutputWithContext(context.Background())
}

func (i *UserFlowAttribute) ToUserFlowAttributeOutputWithContext(ctx context.Context) UserFlowAttributeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserFlowAttributeOutput)
}

// UserFlowAttributeArrayInput is an input type that accepts UserFlowAttributeArray and UserFlowAttributeArrayOutput values.
// You can construct a concrete instance of `UserFlowAttributeArrayInput` via:
//
//	UserFlowAttributeArray{ UserFlowAttributeArgs{...} }
type UserFlowAttributeArrayInput interface {
	pulumi.Input

	ToUserFlowAttributeArrayOutput() UserFlowAttributeArrayOutput
	ToUserFlowAttributeArrayOutputWithContext(context.Context) UserFlowAttributeArrayOutput
}

type UserFlowAttributeArray []UserFlowAttributeInput

func (UserFlowAttributeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserFlowAttribute)(nil)).Elem()
}

func (i UserFlowAttributeArray) ToUserFlowAttributeArrayOutput() UserFlowAttributeArrayOutput {
	return i.ToUserFlowAttributeArrayOutputWithContext(context.Background())
}

func (i UserFlowAttributeArray) ToUserFlowAttributeArrayOutputWithContext(ctx context.Context) UserFlowAttributeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserFlowAttributeArrayOutput)
}

// UserFlowAttributeMapInput is an input type that accepts UserFlowAttributeMap and UserFlowAttributeMapOutput values.
// You can construct a concrete instance of `UserFlowAttributeMapInput` via:
//
//	UserFlowAttributeMap{ "key": UserFlowAttributeArgs{...} }
type UserFlowAttributeMapInput interface {
	pulumi.Input

	ToUserFlowAttributeMapOutput() UserFlowAttributeMapOutput
	ToUserFlowAttributeMapOutputWithContext(context.Context) UserFlowAttributeMapOutput
}

type UserFlowAttributeMap map[string]UserFlowAttributeInput

func (UserFlowAttributeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserFlowAttribute)(nil)).Elem()
}

func (i UserFlowAttributeMap) ToUserFlowAttributeMapOutput() UserFlowAttributeMapOutput {
	return i.ToUserFlowAttributeMapOutputWithContext(context.Background())
}

func (i UserFlowAttributeMap) ToUserFlowAttributeMapOutputWithContext(ctx context.Context) UserFlowAttributeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserFlowAttributeMapOutput)
}

type UserFlowAttributeOutput struct{ *pulumi.OutputState }

func (UserFlowAttributeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserFlowAttribute)(nil)).Elem()
}

func (o UserFlowAttributeOutput) ToUserFlowAttributeOutput() UserFlowAttributeOutput {
	return o
}

func (o UserFlowAttributeOutput) ToUserFlowAttributeOutputWithContext(ctx context.Context) UserFlowAttributeOutput {
	return o
}

// The type of the user flow attribute. Values include `builtIn`, `custom` or `required`.
func (o UserFlowAttributeOutput) AttributeType() pulumi.StringOutput {
	return o.ApplyT(func(v *UserFlowAttribute) pulumi.StringOutput { return v.AttributeType }).(pulumi.StringOutput)
}

// The data type of the user flow attribute. Possible values are `boolean`, `dateTime`, `int64`, `string` or `stringCollection`. Changing this forces a new resource to be created.
func (o UserFlowAttributeOutput) DataType() pulumi.StringOutput {
	return o.ApplyT(func(v *UserFlowAttribute) pulumi.StringOutput { return v.DataType }).(pulumi.StringOutput)
}

// The description of the user flow attribute that is shown to the user at the time of sign-up.
func (o UserFlowAttributeOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *UserFlowAttribute) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The display name of the user flow attribute. Changing this forces a new resource to be created.
func (o UserFlowAttributeOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *UserFlowAttribute) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

type UserFlowAttributeArrayOutput struct{ *pulumi.OutputState }

func (UserFlowAttributeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserFlowAttribute)(nil)).Elem()
}

func (o UserFlowAttributeArrayOutput) ToUserFlowAttributeArrayOutput() UserFlowAttributeArrayOutput {
	return o
}

func (o UserFlowAttributeArrayOutput) ToUserFlowAttributeArrayOutputWithContext(ctx context.Context) UserFlowAttributeArrayOutput {
	return o
}

func (o UserFlowAttributeArrayOutput) Index(i pulumi.IntInput) UserFlowAttributeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserFlowAttribute {
		return vs[0].([]*UserFlowAttribute)[vs[1].(int)]
	}).(UserFlowAttributeOutput)
}

type UserFlowAttributeMapOutput struct{ *pulumi.OutputState }

func (UserFlowAttributeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserFlowAttribute)(nil)).Elem()
}

func (o UserFlowAttributeMapOutput) ToUserFlowAttributeMapOutput() UserFlowAttributeMapOutput {
	return o
}

func (o UserFlowAttributeMapOutput) ToUserFlowAttributeMapOutputWithContext(ctx context.Context) UserFlowAttributeMapOutput {
	return o
}

func (o UserFlowAttributeMapOutput) MapIndex(k pulumi.StringInput) UserFlowAttributeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserFlowAttribute {
		return vs[0].(map[string]*UserFlowAttribute)[vs[1].(string)]
	}).(UserFlowAttributeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserFlowAttributeInput)(nil)).Elem(), &UserFlowAttribute{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserFlowAttributeArrayInput)(nil)).Elem(), UserFlowAttributeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserFlowAttributeMapInput)(nil)).Elem(), UserFlowAttributeMap{})
	pulumi.RegisterOutputType(UserFlowAttributeOutput{})
	pulumi.RegisterOutputType(UserFlowAttributeArrayOutput{})
	pulumi.RegisterOutputType(UserFlowAttributeMapOutput{})
}
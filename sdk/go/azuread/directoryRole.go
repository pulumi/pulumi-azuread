// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Directory Role within Azure Active Directory. Directory Roles are also known as Administrator Roles.
//
// Directory Roles are built-in to Azure Active Directory and are immutable. However, by default they are not activated in a tenant (except for the Global Administrator role). This resource ensures a directory role is activated from its associated role template, and exports the object ID of the role, so that role assignments can be made for it.
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
// *Activate a directory role by its template ID*
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azuread/sdk/v5/go/azuread"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := azuread.NewDirectoryRole(ctx, "example", &azuread.DirectoryRoleArgs{
// 			TemplateId: pulumi.String("00000000-0000-0000-0000-000000000000"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// *Activate a directory role by display name*
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azuread/sdk/v5/go/azuread"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := azuread.NewDirectoryRole(ctx, "example", &azuread.DirectoryRoleArgs{
// 			DisplayName: pulumi.String("Printer administrator"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// This resource does not support importing.
type DirectoryRole struct {
	pulumi.CustomResourceState

	// The description of the directory role.
	Description pulumi.StringOutput `pulumi:"description"`
	// The display name of the directory role to activate. Changing this forces a new resource to be created.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The object ID of the directory role.
	ObjectId pulumi.StringOutput `pulumi:"objectId"`
	// The object ID of the role template from which to activate the directory role. Changing this forces a new resource to be created.
	TemplateId pulumi.StringOutput `pulumi:"templateId"`
}

// NewDirectoryRole registers a new resource with the given unique name, arguments, and options.
func NewDirectoryRole(ctx *pulumi.Context,
	name string, args *DirectoryRoleArgs, opts ...pulumi.ResourceOption) (*DirectoryRole, error) {
	if args == nil {
		args = &DirectoryRoleArgs{}
	}

	var resource DirectoryRole
	err := ctx.RegisterResource("azuread:index/directoryRole:DirectoryRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDirectoryRole gets an existing DirectoryRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDirectoryRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DirectoryRoleState, opts ...pulumi.ResourceOption) (*DirectoryRole, error) {
	var resource DirectoryRole
	err := ctx.ReadResource("azuread:index/directoryRole:DirectoryRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DirectoryRole resources.
type directoryRoleState struct {
	// The description of the directory role.
	Description *string `pulumi:"description"`
	// The display name of the directory role to activate. Changing this forces a new resource to be created.
	DisplayName *string `pulumi:"displayName"`
	// The object ID of the directory role.
	ObjectId *string `pulumi:"objectId"`
	// The object ID of the role template from which to activate the directory role. Changing this forces a new resource to be created.
	TemplateId *string `pulumi:"templateId"`
}

type DirectoryRoleState struct {
	// The description of the directory role.
	Description pulumi.StringPtrInput
	// The display name of the directory role to activate. Changing this forces a new resource to be created.
	DisplayName pulumi.StringPtrInput
	// The object ID of the directory role.
	ObjectId pulumi.StringPtrInput
	// The object ID of the role template from which to activate the directory role. Changing this forces a new resource to be created.
	TemplateId pulumi.StringPtrInput
}

func (DirectoryRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*directoryRoleState)(nil)).Elem()
}

type directoryRoleArgs struct {
	// The display name of the directory role to activate. Changing this forces a new resource to be created.
	DisplayName *string `pulumi:"displayName"`
	// The object ID of the role template from which to activate the directory role. Changing this forces a new resource to be created.
	TemplateId *string `pulumi:"templateId"`
}

// The set of arguments for constructing a DirectoryRole resource.
type DirectoryRoleArgs struct {
	// The display name of the directory role to activate. Changing this forces a new resource to be created.
	DisplayName pulumi.StringPtrInput
	// The object ID of the role template from which to activate the directory role. Changing this forces a new resource to be created.
	TemplateId pulumi.StringPtrInput
}

func (DirectoryRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*directoryRoleArgs)(nil)).Elem()
}

type DirectoryRoleInput interface {
	pulumi.Input

	ToDirectoryRoleOutput() DirectoryRoleOutput
	ToDirectoryRoleOutputWithContext(ctx context.Context) DirectoryRoleOutput
}

func (*DirectoryRole) ElementType() reflect.Type {
	return reflect.TypeOf((*DirectoryRole)(nil))
}

func (i *DirectoryRole) ToDirectoryRoleOutput() DirectoryRoleOutput {
	return i.ToDirectoryRoleOutputWithContext(context.Background())
}

func (i *DirectoryRole) ToDirectoryRoleOutputWithContext(ctx context.Context) DirectoryRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectoryRoleOutput)
}

func (i *DirectoryRole) ToDirectoryRolePtrOutput() DirectoryRolePtrOutput {
	return i.ToDirectoryRolePtrOutputWithContext(context.Background())
}

func (i *DirectoryRole) ToDirectoryRolePtrOutputWithContext(ctx context.Context) DirectoryRolePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectoryRolePtrOutput)
}

type DirectoryRolePtrInput interface {
	pulumi.Input

	ToDirectoryRolePtrOutput() DirectoryRolePtrOutput
	ToDirectoryRolePtrOutputWithContext(ctx context.Context) DirectoryRolePtrOutput
}

type directoryRolePtrType DirectoryRoleArgs

func (*directoryRolePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DirectoryRole)(nil))
}

func (i *directoryRolePtrType) ToDirectoryRolePtrOutput() DirectoryRolePtrOutput {
	return i.ToDirectoryRolePtrOutputWithContext(context.Background())
}

func (i *directoryRolePtrType) ToDirectoryRolePtrOutputWithContext(ctx context.Context) DirectoryRolePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectoryRolePtrOutput)
}

// DirectoryRoleArrayInput is an input type that accepts DirectoryRoleArray and DirectoryRoleArrayOutput values.
// You can construct a concrete instance of `DirectoryRoleArrayInput` via:
//
//          DirectoryRoleArray{ DirectoryRoleArgs{...} }
type DirectoryRoleArrayInput interface {
	pulumi.Input

	ToDirectoryRoleArrayOutput() DirectoryRoleArrayOutput
	ToDirectoryRoleArrayOutputWithContext(context.Context) DirectoryRoleArrayOutput
}

type DirectoryRoleArray []DirectoryRoleInput

func (DirectoryRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DirectoryRole)(nil)).Elem()
}

func (i DirectoryRoleArray) ToDirectoryRoleArrayOutput() DirectoryRoleArrayOutput {
	return i.ToDirectoryRoleArrayOutputWithContext(context.Background())
}

func (i DirectoryRoleArray) ToDirectoryRoleArrayOutputWithContext(ctx context.Context) DirectoryRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectoryRoleArrayOutput)
}

// DirectoryRoleMapInput is an input type that accepts DirectoryRoleMap and DirectoryRoleMapOutput values.
// You can construct a concrete instance of `DirectoryRoleMapInput` via:
//
//          DirectoryRoleMap{ "key": DirectoryRoleArgs{...} }
type DirectoryRoleMapInput interface {
	pulumi.Input

	ToDirectoryRoleMapOutput() DirectoryRoleMapOutput
	ToDirectoryRoleMapOutputWithContext(context.Context) DirectoryRoleMapOutput
}

type DirectoryRoleMap map[string]DirectoryRoleInput

func (DirectoryRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DirectoryRole)(nil)).Elem()
}

func (i DirectoryRoleMap) ToDirectoryRoleMapOutput() DirectoryRoleMapOutput {
	return i.ToDirectoryRoleMapOutputWithContext(context.Background())
}

func (i DirectoryRoleMap) ToDirectoryRoleMapOutputWithContext(ctx context.Context) DirectoryRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectoryRoleMapOutput)
}

type DirectoryRoleOutput struct{ *pulumi.OutputState }

func (DirectoryRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DirectoryRole)(nil))
}

func (o DirectoryRoleOutput) ToDirectoryRoleOutput() DirectoryRoleOutput {
	return o
}

func (o DirectoryRoleOutput) ToDirectoryRoleOutputWithContext(ctx context.Context) DirectoryRoleOutput {
	return o
}

func (o DirectoryRoleOutput) ToDirectoryRolePtrOutput() DirectoryRolePtrOutput {
	return o.ToDirectoryRolePtrOutputWithContext(context.Background())
}

func (o DirectoryRoleOutput) ToDirectoryRolePtrOutputWithContext(ctx context.Context) DirectoryRolePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DirectoryRole) *DirectoryRole {
		return &v
	}).(DirectoryRolePtrOutput)
}

type DirectoryRolePtrOutput struct{ *pulumi.OutputState }

func (DirectoryRolePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DirectoryRole)(nil))
}

func (o DirectoryRolePtrOutput) ToDirectoryRolePtrOutput() DirectoryRolePtrOutput {
	return o
}

func (o DirectoryRolePtrOutput) ToDirectoryRolePtrOutputWithContext(ctx context.Context) DirectoryRolePtrOutput {
	return o
}

func (o DirectoryRolePtrOutput) Elem() DirectoryRoleOutput {
	return o.ApplyT(func(v *DirectoryRole) DirectoryRole {
		if v != nil {
			return *v
		}
		var ret DirectoryRole
		return ret
	}).(DirectoryRoleOutput)
}

type DirectoryRoleArrayOutput struct{ *pulumi.OutputState }

func (DirectoryRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DirectoryRole)(nil))
}

func (o DirectoryRoleArrayOutput) ToDirectoryRoleArrayOutput() DirectoryRoleArrayOutput {
	return o
}

func (o DirectoryRoleArrayOutput) ToDirectoryRoleArrayOutputWithContext(ctx context.Context) DirectoryRoleArrayOutput {
	return o
}

func (o DirectoryRoleArrayOutput) Index(i pulumi.IntInput) DirectoryRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DirectoryRole {
		return vs[0].([]DirectoryRole)[vs[1].(int)]
	}).(DirectoryRoleOutput)
}

type DirectoryRoleMapOutput struct{ *pulumi.OutputState }

func (DirectoryRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DirectoryRole)(nil))
}

func (o DirectoryRoleMapOutput) ToDirectoryRoleMapOutput() DirectoryRoleMapOutput {
	return o
}

func (o DirectoryRoleMapOutput) ToDirectoryRoleMapOutputWithContext(ctx context.Context) DirectoryRoleMapOutput {
	return o
}

func (o DirectoryRoleMapOutput) MapIndex(k pulumi.StringInput) DirectoryRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DirectoryRole {
		return vs[0].(map[string]DirectoryRole)[vs[1].(string)]
	}).(DirectoryRoleOutput)
}

func init() {
	pulumi.RegisterOutputType(DirectoryRoleOutput{})
	pulumi.RegisterOutputType(DirectoryRolePtrOutput{})
	pulumi.RegisterOutputType(DirectoryRoleArrayOutput{})
	pulumi.RegisterOutputType(DirectoryRoleMapOutput{})
}

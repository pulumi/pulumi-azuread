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

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-azuread/sdk/v5/go/azuread"
//	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := azuread.NewApplicationRegistration(ctx, "example", &azuread.ApplicationRegistrationArgs{
//				DisplayName: pulumi.String("example"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleAdministrator, err := random.NewRandomUuid(ctx, "exampleAdministrator", nil)
//			if err != nil {
//				return err
//			}
//			_, err = azuread.NewApplicationAppRole(ctx, "exampleAdminister", &azuread.ApplicationAppRoleArgs{
//				ApplicationId: example.ID(),
//				RoleId:        exampleAdministrator.ID(),
//				AllowedMemberTypes: pulumi.StringArray{
//					pulumi.String("User"),
//				},
//				Description: pulumi.String("My role description"),
//				DisplayName: pulumi.String("Administer"),
//				Value:       pulumi.String("admin"),
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
// > **Tip** For managing more app roles, create additional instances of this resource
//
// *Usage with Application resource*
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
//			example, err := azuread.NewApplication(ctx, "example", &azuread.ApplicationArgs{
//				DisplayName: pulumi.String("example"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuread.NewApplicationAppRole(ctx, "exampleAdminister", &azuread.ApplicationAppRoleArgs{
//				ApplicationId: example.ID(),
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
// Application App Roles can be imported using the object ID of the application and the ID of the app role, in the following format.
//
// ```sh
// $ pulumi import azuread:index/applicationAppRole:ApplicationAppRole example /applications/00000000-0000-0000-0000-000000000000/appRoles/11111111-1111-1111-1111-111111111111
// ```
type ApplicationAppRole struct {
	pulumi.CustomResourceState

	// A set of values to specify whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications by setting to `Application`, or to both.
	AllowedMemberTypes pulumi.StringArrayOutput `pulumi:"allowedMemberTypes"`
	// The resource ID of the application registration. Changing this forces a new resource to be created.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// Description of the app role that appears when the role is being assigned, and if the role functions as an application permissions, during the consent experiences.
	Description pulumi.StringOutput `pulumi:"description"`
	// Display name for the app role that appears during app role assignment and in consent experiences.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The unique identifier of the app role
	RoleId pulumi.StringOutput `pulumi:"roleId"`
	// The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal.
	//
	// > **Roles and Permission Scopes** In Azure Active Directory, application roles and permission scopes exported by an application share the same namespace and cannot contain duplicate values.
	Value pulumi.StringPtrOutput `pulumi:"value"`
}

// NewApplicationAppRole registers a new resource with the given unique name, arguments, and options.
func NewApplicationAppRole(ctx *pulumi.Context,
	name string, args *ApplicationAppRoleArgs, opts ...pulumi.ResourceOption) (*ApplicationAppRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AllowedMemberTypes == nil {
		return nil, errors.New("invalid value for required argument 'AllowedMemberTypes'")
	}
	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.RoleId == nil {
		return nil, errors.New("invalid value for required argument 'RoleId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApplicationAppRole
	err := ctx.RegisterResource("azuread:index/applicationAppRole:ApplicationAppRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationAppRole gets an existing ApplicationAppRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationAppRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationAppRoleState, opts ...pulumi.ResourceOption) (*ApplicationAppRole, error) {
	var resource ApplicationAppRole
	err := ctx.ReadResource("azuread:index/applicationAppRole:ApplicationAppRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationAppRole resources.
type applicationAppRoleState struct {
	// A set of values to specify whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications by setting to `Application`, or to both.
	AllowedMemberTypes []string `pulumi:"allowedMemberTypes"`
	// The resource ID of the application registration. Changing this forces a new resource to be created.
	ApplicationId *string `pulumi:"applicationId"`
	// Description of the app role that appears when the role is being assigned, and if the role functions as an application permissions, during the consent experiences.
	Description *string `pulumi:"description"`
	// Display name for the app role that appears during app role assignment and in consent experiences.
	DisplayName *string `pulumi:"displayName"`
	// The unique identifier of the app role
	RoleId *string `pulumi:"roleId"`
	// The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal.
	//
	// > **Roles and Permission Scopes** In Azure Active Directory, application roles and permission scopes exported by an application share the same namespace and cannot contain duplicate values.
	Value *string `pulumi:"value"`
}

type ApplicationAppRoleState struct {
	// A set of values to specify whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications by setting to `Application`, or to both.
	AllowedMemberTypes pulumi.StringArrayInput
	// The resource ID of the application registration. Changing this forces a new resource to be created.
	ApplicationId pulumi.StringPtrInput
	// Description of the app role that appears when the role is being assigned, and if the role functions as an application permissions, during the consent experiences.
	Description pulumi.StringPtrInput
	// Display name for the app role that appears during app role assignment and in consent experiences.
	DisplayName pulumi.StringPtrInput
	// The unique identifier of the app role
	RoleId pulumi.StringPtrInput
	// The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal.
	//
	// > **Roles and Permission Scopes** In Azure Active Directory, application roles and permission scopes exported by an application share the same namespace and cannot contain duplicate values.
	Value pulumi.StringPtrInput
}

func (ApplicationAppRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationAppRoleState)(nil)).Elem()
}

type applicationAppRoleArgs struct {
	// A set of values to specify whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications by setting to `Application`, or to both.
	AllowedMemberTypes []string `pulumi:"allowedMemberTypes"`
	// The resource ID of the application registration. Changing this forces a new resource to be created.
	ApplicationId string `pulumi:"applicationId"`
	// Description of the app role that appears when the role is being assigned, and if the role functions as an application permissions, during the consent experiences.
	Description string `pulumi:"description"`
	// Display name for the app role that appears during app role assignment and in consent experiences.
	DisplayName string `pulumi:"displayName"`
	// The unique identifier of the app role
	RoleId string `pulumi:"roleId"`
	// The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal.
	//
	// > **Roles and Permission Scopes** In Azure Active Directory, application roles and permission scopes exported by an application share the same namespace and cannot contain duplicate values.
	Value *string `pulumi:"value"`
}

// The set of arguments for constructing a ApplicationAppRole resource.
type ApplicationAppRoleArgs struct {
	// A set of values to specify whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications by setting to `Application`, or to both.
	AllowedMemberTypes pulumi.StringArrayInput
	// The resource ID of the application registration. Changing this forces a new resource to be created.
	ApplicationId pulumi.StringInput
	// Description of the app role that appears when the role is being assigned, and if the role functions as an application permissions, during the consent experiences.
	Description pulumi.StringInput
	// Display name for the app role that appears during app role assignment and in consent experiences.
	DisplayName pulumi.StringInput
	// The unique identifier of the app role
	RoleId pulumi.StringInput
	// The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal.
	//
	// > **Roles and Permission Scopes** In Azure Active Directory, application roles and permission scopes exported by an application share the same namespace and cannot contain duplicate values.
	Value pulumi.StringPtrInput
}

func (ApplicationAppRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationAppRoleArgs)(nil)).Elem()
}

type ApplicationAppRoleInput interface {
	pulumi.Input

	ToApplicationAppRoleOutput() ApplicationAppRoleOutput
	ToApplicationAppRoleOutputWithContext(ctx context.Context) ApplicationAppRoleOutput
}

func (*ApplicationAppRole) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationAppRole)(nil)).Elem()
}

func (i *ApplicationAppRole) ToApplicationAppRoleOutput() ApplicationAppRoleOutput {
	return i.ToApplicationAppRoleOutputWithContext(context.Background())
}

func (i *ApplicationAppRole) ToApplicationAppRoleOutputWithContext(ctx context.Context) ApplicationAppRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationAppRoleOutput)
}

// ApplicationAppRoleArrayInput is an input type that accepts ApplicationAppRoleArray and ApplicationAppRoleArrayOutput values.
// You can construct a concrete instance of `ApplicationAppRoleArrayInput` via:
//
//	ApplicationAppRoleArray{ ApplicationAppRoleArgs{...} }
type ApplicationAppRoleArrayInput interface {
	pulumi.Input

	ToApplicationAppRoleArrayOutput() ApplicationAppRoleArrayOutput
	ToApplicationAppRoleArrayOutputWithContext(context.Context) ApplicationAppRoleArrayOutput
}

type ApplicationAppRoleArray []ApplicationAppRoleInput

func (ApplicationAppRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationAppRole)(nil)).Elem()
}

func (i ApplicationAppRoleArray) ToApplicationAppRoleArrayOutput() ApplicationAppRoleArrayOutput {
	return i.ToApplicationAppRoleArrayOutputWithContext(context.Background())
}

func (i ApplicationAppRoleArray) ToApplicationAppRoleArrayOutputWithContext(ctx context.Context) ApplicationAppRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationAppRoleArrayOutput)
}

// ApplicationAppRoleMapInput is an input type that accepts ApplicationAppRoleMap and ApplicationAppRoleMapOutput values.
// You can construct a concrete instance of `ApplicationAppRoleMapInput` via:
//
//	ApplicationAppRoleMap{ "key": ApplicationAppRoleArgs{...} }
type ApplicationAppRoleMapInput interface {
	pulumi.Input

	ToApplicationAppRoleMapOutput() ApplicationAppRoleMapOutput
	ToApplicationAppRoleMapOutputWithContext(context.Context) ApplicationAppRoleMapOutput
}

type ApplicationAppRoleMap map[string]ApplicationAppRoleInput

func (ApplicationAppRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationAppRole)(nil)).Elem()
}

func (i ApplicationAppRoleMap) ToApplicationAppRoleMapOutput() ApplicationAppRoleMapOutput {
	return i.ToApplicationAppRoleMapOutputWithContext(context.Background())
}

func (i ApplicationAppRoleMap) ToApplicationAppRoleMapOutputWithContext(ctx context.Context) ApplicationAppRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationAppRoleMapOutput)
}

type ApplicationAppRoleOutput struct{ *pulumi.OutputState }

func (ApplicationAppRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationAppRole)(nil)).Elem()
}

func (o ApplicationAppRoleOutput) ToApplicationAppRoleOutput() ApplicationAppRoleOutput {
	return o
}

func (o ApplicationAppRoleOutput) ToApplicationAppRoleOutputWithContext(ctx context.Context) ApplicationAppRoleOutput {
	return o
}

// A set of values to specify whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications by setting to `Application`, or to both.
func (o ApplicationAppRoleOutput) AllowedMemberTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationAppRole) pulumi.StringArrayOutput { return v.AllowedMemberTypes }).(pulumi.StringArrayOutput)
}

// The resource ID of the application registration. Changing this forces a new resource to be created.
func (o ApplicationAppRoleOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationAppRole) pulumi.StringOutput { return v.ApplicationId }).(pulumi.StringOutput)
}

// Description of the app role that appears when the role is being assigned, and if the role functions as an application permissions, during the consent experiences.
func (o ApplicationAppRoleOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationAppRole) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Display name for the app role that appears during app role assignment and in consent experiences.
func (o ApplicationAppRoleOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationAppRole) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The unique identifier of the app role
func (o ApplicationAppRoleOutput) RoleId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationAppRole) pulumi.StringOutput { return v.RoleId }).(pulumi.StringOutput)
}

// The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal.
//
// > **Roles and Permission Scopes** In Azure Active Directory, application roles and permission scopes exported by an application share the same namespace and cannot contain duplicate values.
func (o ApplicationAppRoleOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationAppRole) pulumi.StringPtrOutput { return v.Value }).(pulumi.StringPtrOutput)
}

type ApplicationAppRoleArrayOutput struct{ *pulumi.OutputState }

func (ApplicationAppRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationAppRole)(nil)).Elem()
}

func (o ApplicationAppRoleArrayOutput) ToApplicationAppRoleArrayOutput() ApplicationAppRoleArrayOutput {
	return o
}

func (o ApplicationAppRoleArrayOutput) ToApplicationAppRoleArrayOutputWithContext(ctx context.Context) ApplicationAppRoleArrayOutput {
	return o
}

func (o ApplicationAppRoleArrayOutput) Index(i pulumi.IntInput) ApplicationAppRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplicationAppRole {
		return vs[0].([]*ApplicationAppRole)[vs[1].(int)]
	}).(ApplicationAppRoleOutput)
}

type ApplicationAppRoleMapOutput struct{ *pulumi.OutputState }

func (ApplicationAppRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationAppRole)(nil)).Elem()
}

func (o ApplicationAppRoleMapOutput) ToApplicationAppRoleMapOutput() ApplicationAppRoleMapOutput {
	return o
}

func (o ApplicationAppRoleMapOutput) ToApplicationAppRoleMapOutputWithContext(ctx context.Context) ApplicationAppRoleMapOutput {
	return o
}

func (o ApplicationAppRoleMapOutput) MapIndex(k pulumi.StringInput) ApplicationAppRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplicationAppRole {
		return vs[0].(map[string]*ApplicationAppRole)[vs[1].(string)]
	}).(ApplicationAppRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationAppRoleInput)(nil)).Elem(), &ApplicationAppRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationAppRoleArrayInput)(nil)).Elem(), ApplicationAppRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationAppRoleMapInput)(nil)).Elem(), ApplicationAppRoleMap{})
	pulumi.RegisterOutputType(ApplicationAppRoleOutput{})
	pulumi.RegisterOutputType(ApplicationAppRoleArrayOutput{})
	pulumi.RegisterOutputType(ApplicationAppRoleMapOutput{})
}

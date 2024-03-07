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
// <!--Start PulumiCodeChooser -->
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
//			_, err := azuread.NewApplicationRegistration(ctx, "example", &azuread.ApplicationRegistrationArgs{
//				DisplayName: pulumi.String("example"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleAdminister, err := random.NewRandomUuid(ctx, "example_administer", nil)
//			if err != nil {
//				return err
//			}
//			_, err = azuread.NewApplicationPermissionScope(ctx, "example", &azuread.ApplicationPermissionScopeArgs{
//				ApplicationId:           pulumi.Any(test.Id),
//				ScopeId:                 exampleAdminister.ID(),
//				Value:                   pulumi.String("administer"),
//				AdminConsentDescription: pulumi.String("Administer the application"),
//				AdminConsentDisplayName: pulumi.String("Administer"),
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
// > **Tip** For managing more permissions scopes, create additional instances of this resource
//
// *Usage with Application resource*
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
//			example, err := azuread.NewApplication(ctx, "example", &azuread.ApplicationArgs{
//				DisplayName: pulumi.String("example"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuread.NewApplicationPermissionScope(ctx, "example", &azuread.ApplicationPermissionScopeArgs{
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
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// Application App Roles can be imported using the object ID of the application and the ID of the permission scope, in the following format.
//
// ```sh
// $ pulumi import azuread:index/applicationPermissionScope:ApplicationPermissionScope example /applications/00000000-0000-0000-0000-000000000000/permissionScopes/11111111-1111-1111-1111-111111111111
// ```
type ApplicationPermissionScope struct {
	pulumi.CustomResourceState

	// Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
	AdminConsentDescription pulumi.StringOutput `pulumi:"adminConsentDescription"`
	// Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
	AdminConsentDisplayName pulumi.StringOutput `pulumi:"adminConsentDisplayName"`
	// The resource ID of the application registration. Changing this forces a new resource to be created.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// The unique identifier of the permission scope. Must be a valid UUID. Changing this forces a new resource to be created.
	ScopeId pulumi.StringOutput `pulumi:"scopeId"`
	// Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions.
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
	UserConsentDescription pulumi.StringPtrOutput `pulumi:"userConsentDescription"`
	// Display name for the delegated permission that appears in the end user consent experience
	UserConsentDisplayName pulumi.StringPtrOutput `pulumi:"userConsentDisplayName"`
	// The value that is used for the `scp` claim in OAuth access tokens.
	//
	// > **Roles and Permission Scopes** In Azure Active Directory, application roles and permission scopes exported by an application share the same namespace and cannot contain duplicate values.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewApplicationPermissionScope registers a new resource with the given unique name, arguments, and options.
func NewApplicationPermissionScope(ctx *pulumi.Context,
	name string, args *ApplicationPermissionScopeArgs, opts ...pulumi.ResourceOption) (*ApplicationPermissionScope, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AdminConsentDescription == nil {
		return nil, errors.New("invalid value for required argument 'AdminConsentDescription'")
	}
	if args.AdminConsentDisplayName == nil {
		return nil, errors.New("invalid value for required argument 'AdminConsentDisplayName'")
	}
	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	if args.ScopeId == nil {
		return nil, errors.New("invalid value for required argument 'ScopeId'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApplicationPermissionScope
	err := ctx.RegisterResource("azuread:index/applicationPermissionScope:ApplicationPermissionScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationPermissionScope gets an existing ApplicationPermissionScope resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationPermissionScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationPermissionScopeState, opts ...pulumi.ResourceOption) (*ApplicationPermissionScope, error) {
	var resource ApplicationPermissionScope
	err := ctx.ReadResource("azuread:index/applicationPermissionScope:ApplicationPermissionScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationPermissionScope resources.
type applicationPermissionScopeState struct {
	// Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
	AdminConsentDescription *string `pulumi:"adminConsentDescription"`
	// Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
	AdminConsentDisplayName *string `pulumi:"adminConsentDisplayName"`
	// The resource ID of the application registration. Changing this forces a new resource to be created.
	ApplicationId *string `pulumi:"applicationId"`
	// The unique identifier of the permission scope. Must be a valid UUID. Changing this forces a new resource to be created.
	ScopeId *string `pulumi:"scopeId"`
	// Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions.
	Type *string `pulumi:"type"`
	// Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
	UserConsentDescription *string `pulumi:"userConsentDescription"`
	// Display name for the delegated permission that appears in the end user consent experience
	UserConsentDisplayName *string `pulumi:"userConsentDisplayName"`
	// The value that is used for the `scp` claim in OAuth access tokens.
	//
	// > **Roles and Permission Scopes** In Azure Active Directory, application roles and permission scopes exported by an application share the same namespace and cannot contain duplicate values.
	Value *string `pulumi:"value"`
}

type ApplicationPermissionScopeState struct {
	// Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
	AdminConsentDescription pulumi.StringPtrInput
	// Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
	AdminConsentDisplayName pulumi.StringPtrInput
	// The resource ID of the application registration. Changing this forces a new resource to be created.
	ApplicationId pulumi.StringPtrInput
	// The unique identifier of the permission scope. Must be a valid UUID. Changing this forces a new resource to be created.
	ScopeId pulumi.StringPtrInput
	// Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions.
	Type pulumi.StringPtrInput
	// Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
	UserConsentDescription pulumi.StringPtrInput
	// Display name for the delegated permission that appears in the end user consent experience
	UserConsentDisplayName pulumi.StringPtrInput
	// The value that is used for the `scp` claim in OAuth access tokens.
	//
	// > **Roles and Permission Scopes** In Azure Active Directory, application roles and permission scopes exported by an application share the same namespace and cannot contain duplicate values.
	Value pulumi.StringPtrInput
}

func (ApplicationPermissionScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationPermissionScopeState)(nil)).Elem()
}

type applicationPermissionScopeArgs struct {
	// Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
	AdminConsentDescription string `pulumi:"adminConsentDescription"`
	// Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
	AdminConsentDisplayName string `pulumi:"adminConsentDisplayName"`
	// The resource ID of the application registration. Changing this forces a new resource to be created.
	ApplicationId string `pulumi:"applicationId"`
	// The unique identifier of the permission scope. Must be a valid UUID. Changing this forces a new resource to be created.
	ScopeId string `pulumi:"scopeId"`
	// Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions.
	Type *string `pulumi:"type"`
	// Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
	UserConsentDescription *string `pulumi:"userConsentDescription"`
	// Display name for the delegated permission that appears in the end user consent experience
	UserConsentDisplayName *string `pulumi:"userConsentDisplayName"`
	// The value that is used for the `scp` claim in OAuth access tokens.
	//
	// > **Roles and Permission Scopes** In Azure Active Directory, application roles and permission scopes exported by an application share the same namespace and cannot contain duplicate values.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a ApplicationPermissionScope resource.
type ApplicationPermissionScopeArgs struct {
	// Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
	AdminConsentDescription pulumi.StringInput
	// Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
	AdminConsentDisplayName pulumi.StringInput
	// The resource ID of the application registration. Changing this forces a new resource to be created.
	ApplicationId pulumi.StringInput
	// The unique identifier of the permission scope. Must be a valid UUID. Changing this forces a new resource to be created.
	ScopeId pulumi.StringInput
	// Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions.
	Type pulumi.StringPtrInput
	// Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
	UserConsentDescription pulumi.StringPtrInput
	// Display name for the delegated permission that appears in the end user consent experience
	UserConsentDisplayName pulumi.StringPtrInput
	// The value that is used for the `scp` claim in OAuth access tokens.
	//
	// > **Roles and Permission Scopes** In Azure Active Directory, application roles and permission scopes exported by an application share the same namespace and cannot contain duplicate values.
	Value pulumi.StringInput
}

func (ApplicationPermissionScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationPermissionScopeArgs)(nil)).Elem()
}

type ApplicationPermissionScopeInput interface {
	pulumi.Input

	ToApplicationPermissionScopeOutput() ApplicationPermissionScopeOutput
	ToApplicationPermissionScopeOutputWithContext(ctx context.Context) ApplicationPermissionScopeOutput
}

func (*ApplicationPermissionScope) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationPermissionScope)(nil)).Elem()
}

func (i *ApplicationPermissionScope) ToApplicationPermissionScopeOutput() ApplicationPermissionScopeOutput {
	return i.ToApplicationPermissionScopeOutputWithContext(context.Background())
}

func (i *ApplicationPermissionScope) ToApplicationPermissionScopeOutputWithContext(ctx context.Context) ApplicationPermissionScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPermissionScopeOutput)
}

// ApplicationPermissionScopeArrayInput is an input type that accepts ApplicationPermissionScopeArray and ApplicationPermissionScopeArrayOutput values.
// You can construct a concrete instance of `ApplicationPermissionScopeArrayInput` via:
//
//	ApplicationPermissionScopeArray{ ApplicationPermissionScopeArgs{...} }
type ApplicationPermissionScopeArrayInput interface {
	pulumi.Input

	ToApplicationPermissionScopeArrayOutput() ApplicationPermissionScopeArrayOutput
	ToApplicationPermissionScopeArrayOutputWithContext(context.Context) ApplicationPermissionScopeArrayOutput
}

type ApplicationPermissionScopeArray []ApplicationPermissionScopeInput

func (ApplicationPermissionScopeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationPermissionScope)(nil)).Elem()
}

func (i ApplicationPermissionScopeArray) ToApplicationPermissionScopeArrayOutput() ApplicationPermissionScopeArrayOutput {
	return i.ToApplicationPermissionScopeArrayOutputWithContext(context.Background())
}

func (i ApplicationPermissionScopeArray) ToApplicationPermissionScopeArrayOutputWithContext(ctx context.Context) ApplicationPermissionScopeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPermissionScopeArrayOutput)
}

// ApplicationPermissionScopeMapInput is an input type that accepts ApplicationPermissionScopeMap and ApplicationPermissionScopeMapOutput values.
// You can construct a concrete instance of `ApplicationPermissionScopeMapInput` via:
//
//	ApplicationPermissionScopeMap{ "key": ApplicationPermissionScopeArgs{...} }
type ApplicationPermissionScopeMapInput interface {
	pulumi.Input

	ToApplicationPermissionScopeMapOutput() ApplicationPermissionScopeMapOutput
	ToApplicationPermissionScopeMapOutputWithContext(context.Context) ApplicationPermissionScopeMapOutput
}

type ApplicationPermissionScopeMap map[string]ApplicationPermissionScopeInput

func (ApplicationPermissionScopeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationPermissionScope)(nil)).Elem()
}

func (i ApplicationPermissionScopeMap) ToApplicationPermissionScopeMapOutput() ApplicationPermissionScopeMapOutput {
	return i.ToApplicationPermissionScopeMapOutputWithContext(context.Background())
}

func (i ApplicationPermissionScopeMap) ToApplicationPermissionScopeMapOutputWithContext(ctx context.Context) ApplicationPermissionScopeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPermissionScopeMapOutput)
}

type ApplicationPermissionScopeOutput struct{ *pulumi.OutputState }

func (ApplicationPermissionScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationPermissionScope)(nil)).Elem()
}

func (o ApplicationPermissionScopeOutput) ToApplicationPermissionScopeOutput() ApplicationPermissionScopeOutput {
	return o
}

func (o ApplicationPermissionScopeOutput) ToApplicationPermissionScopeOutputWithContext(ctx context.Context) ApplicationPermissionScopeOutput {
	return o
}

// Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
func (o ApplicationPermissionScopeOutput) AdminConsentDescription() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationPermissionScope) pulumi.StringOutput { return v.AdminConsentDescription }).(pulumi.StringOutput)
}

// Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
func (o ApplicationPermissionScopeOutput) AdminConsentDisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationPermissionScope) pulumi.StringOutput { return v.AdminConsentDisplayName }).(pulumi.StringOutput)
}

// The resource ID of the application registration. Changing this forces a new resource to be created.
func (o ApplicationPermissionScopeOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationPermissionScope) pulumi.StringOutput { return v.ApplicationId }).(pulumi.StringOutput)
}

// The unique identifier of the permission scope. Must be a valid UUID. Changing this forces a new resource to be created.
func (o ApplicationPermissionScopeOutput) ScopeId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationPermissionScope) pulumi.StringOutput { return v.ScopeId }).(pulumi.StringOutput)
}

// Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions.
func (o ApplicationPermissionScopeOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationPermissionScope) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

// Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
func (o ApplicationPermissionScopeOutput) UserConsentDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationPermissionScope) pulumi.StringPtrOutput { return v.UserConsentDescription }).(pulumi.StringPtrOutput)
}

// Display name for the delegated permission that appears in the end user consent experience
func (o ApplicationPermissionScopeOutput) UserConsentDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationPermissionScope) pulumi.StringPtrOutput { return v.UserConsentDisplayName }).(pulumi.StringPtrOutput)
}

// The value that is used for the `scp` claim in OAuth access tokens.
//
// > **Roles and Permission Scopes** In Azure Active Directory, application roles and permission scopes exported by an application share the same namespace and cannot contain duplicate values.
func (o ApplicationPermissionScopeOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationPermissionScope) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

type ApplicationPermissionScopeArrayOutput struct{ *pulumi.OutputState }

func (ApplicationPermissionScopeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationPermissionScope)(nil)).Elem()
}

func (o ApplicationPermissionScopeArrayOutput) ToApplicationPermissionScopeArrayOutput() ApplicationPermissionScopeArrayOutput {
	return o
}

func (o ApplicationPermissionScopeArrayOutput) ToApplicationPermissionScopeArrayOutputWithContext(ctx context.Context) ApplicationPermissionScopeArrayOutput {
	return o
}

func (o ApplicationPermissionScopeArrayOutput) Index(i pulumi.IntInput) ApplicationPermissionScopeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplicationPermissionScope {
		return vs[0].([]*ApplicationPermissionScope)[vs[1].(int)]
	}).(ApplicationPermissionScopeOutput)
}

type ApplicationPermissionScopeMapOutput struct{ *pulumi.OutputState }

func (ApplicationPermissionScopeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationPermissionScope)(nil)).Elem()
}

func (o ApplicationPermissionScopeMapOutput) ToApplicationPermissionScopeMapOutput() ApplicationPermissionScopeMapOutput {
	return o
}

func (o ApplicationPermissionScopeMapOutput) ToApplicationPermissionScopeMapOutputWithContext(ctx context.Context) ApplicationPermissionScopeMapOutput {
	return o
}

func (o ApplicationPermissionScopeMapOutput) MapIndex(k pulumi.StringInput) ApplicationPermissionScopeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplicationPermissionScope {
		return vs[0].(map[string]*ApplicationPermissionScope)[vs[1].(string)]
	}).(ApplicationPermissionScopeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationPermissionScopeInput)(nil)).Elem(), &ApplicationPermissionScope{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationPermissionScopeArrayInput)(nil)).Elem(), ApplicationPermissionScopeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationPermissionScopeMapInput)(nil)).Elem(), ApplicationPermissionScopeMap{})
	pulumi.RegisterOutputType(ApplicationPermissionScopeOutput{})
	pulumi.RegisterOutputType(ApplicationPermissionScopeArrayOutput{})
	pulumi.RegisterOutputType(ApplicationPermissionScopeMapOutput{})
}

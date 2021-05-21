// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an OAuth 2.0 Permission Scope associated with an application.
//
// > **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write all applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azuread/sdk/v4/go/azuread"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleApplication, err := azuread.NewApplication(ctx, "exampleApplication", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = azuread.NewApplicationOauth2PermissionScope(ctx, "exampleApplicationOauth2PermissionScope", &azuread.ApplicationOauth2PermissionScopeArgs{
// 			ApplicationObjectId:     exampleApplication.ID(),
// 			AdminConsentDescription: pulumi.String("Administer the application"),
// 			AdminConsentDisplayName: pulumi.String("Administer"),
// 			Enabled:                 pulumi.Bool(true),
// 			Type:                    pulumi.String("User"),
// 			UserConsentDescription:  pulumi.String("Administer the application"),
// 			UserConsentDisplayName:  pulumi.String("Administer"),
// 			Value:                   pulumi.String("administer"),
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
// OAuth2 Permission Scopes can be imported using the `object_id` of an Application and the `id` of the Permission Scope, e.g.
//
// ```sh
//  $ pulumi import azuread:index/applicationOauth2PermissionScope:ApplicationOauth2PermissionScope test 00000000-0000-0000-0000-000000000000/scope/11111111-1111-1111-1111-111111111111
// ```
type ApplicationOauth2PermissionScope struct {
	pulumi.CustomResourceState

	// Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
	AdminConsentDescription pulumi.StringOutput `pulumi:"adminConsentDescription"`
	// Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
	AdminConsentDisplayName pulumi.StringOutput `pulumi:"adminConsentDisplayName"`
	// The Object ID of the Application for which this Permission should be created. Changing this field forces a new resource to be created.
	ApplicationObjectId pulumi.StringOutput `pulumi:"applicationObjectId"`
	// Determines if the permission scope is enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Deprecated: [NOTE] This attribute has been renamed to `enabled` and will be removed in version 2.0 of the AzureAD provider
	IsEnabled pulumi.BoolPtrOutput `pulumi:"isEnabled"`
	// Deprecated: [NOTE] This attribute has been renamed to `scope_id` and will be removed in version 2.0 of the AzureAD provider
	PermissionId pulumi.StringOutput `pulumi:"permissionId"`
	// Specifies a custom UUID for the permission scope. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.
	ScopeId pulumi.StringOutput `pulumi:"scopeId"`
	// Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Defaults to `User`. Possible values are `User` or `Admin`.
	Type pulumi.StringOutput `pulumi:"type"`
	// Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
	UserConsentDescription pulumi.StringOutput `pulumi:"userConsentDescription"`
	// Display name for the delegated permission that appears in the end user consent experience.
	UserConsentDisplayName pulumi.StringOutput `pulumi:"userConsentDisplayName"`
	// The value that is used for the `scp` claim in OAuth 2.0 access tokens.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewApplicationOauth2PermissionScope registers a new resource with the given unique name, arguments, and options.
func NewApplicationOauth2PermissionScope(ctx *pulumi.Context,
	name string, args *ApplicationOauth2PermissionScopeArgs, opts ...pulumi.ResourceOption) (*ApplicationOauth2PermissionScope, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AdminConsentDescription == nil {
		return nil, errors.New("invalid value for required argument 'AdminConsentDescription'")
	}
	if args.AdminConsentDisplayName == nil {
		return nil, errors.New("invalid value for required argument 'AdminConsentDisplayName'")
	}
	if args.ApplicationObjectId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationObjectId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.UserConsentDescription == nil {
		return nil, errors.New("invalid value for required argument 'UserConsentDescription'")
	}
	if args.UserConsentDisplayName == nil {
		return nil, errors.New("invalid value for required argument 'UserConsentDisplayName'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	var resource ApplicationOauth2PermissionScope
	err := ctx.RegisterResource("azuread:index/applicationOauth2PermissionScope:ApplicationOauth2PermissionScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationOauth2PermissionScope gets an existing ApplicationOauth2PermissionScope resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationOauth2PermissionScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationOauth2PermissionScopeState, opts ...pulumi.ResourceOption) (*ApplicationOauth2PermissionScope, error) {
	var resource ApplicationOauth2PermissionScope
	err := ctx.ReadResource("azuread:index/applicationOauth2PermissionScope:ApplicationOauth2PermissionScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationOauth2PermissionScope resources.
type applicationOauth2PermissionScopeState struct {
	// Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
	AdminConsentDescription *string `pulumi:"adminConsentDescription"`
	// Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
	AdminConsentDisplayName *string `pulumi:"adminConsentDisplayName"`
	// The Object ID of the Application for which this Permission should be created. Changing this field forces a new resource to be created.
	ApplicationObjectId *string `pulumi:"applicationObjectId"`
	// Determines if the permission scope is enabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Deprecated: [NOTE] This attribute has been renamed to `enabled` and will be removed in version 2.0 of the AzureAD provider
	IsEnabled *bool `pulumi:"isEnabled"`
	// Deprecated: [NOTE] This attribute has been renamed to `scope_id` and will be removed in version 2.0 of the AzureAD provider
	PermissionId *string `pulumi:"permissionId"`
	// Specifies a custom UUID for the permission scope. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.
	ScopeId *string `pulumi:"scopeId"`
	// Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Defaults to `User`. Possible values are `User` or `Admin`.
	Type *string `pulumi:"type"`
	// Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
	UserConsentDescription *string `pulumi:"userConsentDescription"`
	// Display name for the delegated permission that appears in the end user consent experience.
	UserConsentDisplayName *string `pulumi:"userConsentDisplayName"`
	// The value that is used for the `scp` claim in OAuth 2.0 access tokens.
	Value *string `pulumi:"value"`
}

type ApplicationOauth2PermissionScopeState struct {
	// Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
	AdminConsentDescription pulumi.StringPtrInput
	// Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
	AdminConsentDisplayName pulumi.StringPtrInput
	// The Object ID of the Application for which this Permission should be created. Changing this field forces a new resource to be created.
	ApplicationObjectId pulumi.StringPtrInput
	// Determines if the permission scope is enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Deprecated: [NOTE] This attribute has been renamed to `enabled` and will be removed in version 2.0 of the AzureAD provider
	IsEnabled pulumi.BoolPtrInput
	// Deprecated: [NOTE] This attribute has been renamed to `scope_id` and will be removed in version 2.0 of the AzureAD provider
	PermissionId pulumi.StringPtrInput
	// Specifies a custom UUID for the permission scope. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.
	ScopeId pulumi.StringPtrInput
	// Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Defaults to `User`. Possible values are `User` or `Admin`.
	Type pulumi.StringPtrInput
	// Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
	UserConsentDescription pulumi.StringPtrInput
	// Display name for the delegated permission that appears in the end user consent experience.
	UserConsentDisplayName pulumi.StringPtrInput
	// The value that is used for the `scp` claim in OAuth 2.0 access tokens.
	Value pulumi.StringPtrInput
}

func (ApplicationOauth2PermissionScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationOauth2PermissionScopeState)(nil)).Elem()
}

type applicationOauth2PermissionScopeArgs struct {
	// Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
	AdminConsentDescription string `pulumi:"adminConsentDescription"`
	// Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
	AdminConsentDisplayName string `pulumi:"adminConsentDisplayName"`
	// The Object ID of the Application for which this Permission should be created. Changing this field forces a new resource to be created.
	ApplicationObjectId string `pulumi:"applicationObjectId"`
	// Determines if the permission scope is enabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Deprecated: [NOTE] This attribute has been renamed to `enabled` and will be removed in version 2.0 of the AzureAD provider
	IsEnabled *bool `pulumi:"isEnabled"`
	// Deprecated: [NOTE] This attribute has been renamed to `scope_id` and will be removed in version 2.0 of the AzureAD provider
	PermissionId *string `pulumi:"permissionId"`
	// Specifies a custom UUID for the permission scope. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.
	ScopeId *string `pulumi:"scopeId"`
	// Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Defaults to `User`. Possible values are `User` or `Admin`.
	Type string `pulumi:"type"`
	// Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
	UserConsentDescription string `pulumi:"userConsentDescription"`
	// Display name for the delegated permission that appears in the end user consent experience.
	UserConsentDisplayName string `pulumi:"userConsentDisplayName"`
	// The value that is used for the `scp` claim in OAuth 2.0 access tokens.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a ApplicationOauth2PermissionScope resource.
type ApplicationOauth2PermissionScopeArgs struct {
	// Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
	AdminConsentDescription pulumi.StringInput
	// Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
	AdminConsentDisplayName pulumi.StringInput
	// The Object ID of the Application for which this Permission should be created. Changing this field forces a new resource to be created.
	ApplicationObjectId pulumi.StringInput
	// Determines if the permission scope is enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Deprecated: [NOTE] This attribute has been renamed to `enabled` and will be removed in version 2.0 of the AzureAD provider
	IsEnabled pulumi.BoolPtrInput
	// Deprecated: [NOTE] This attribute has been renamed to `scope_id` and will be removed in version 2.0 of the AzureAD provider
	PermissionId pulumi.StringPtrInput
	// Specifies a custom UUID for the permission scope. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.
	ScopeId pulumi.StringPtrInput
	// Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Defaults to `User`. Possible values are `User` or `Admin`.
	Type pulumi.StringInput
	// Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
	UserConsentDescription pulumi.StringInput
	// Display name for the delegated permission that appears in the end user consent experience.
	UserConsentDisplayName pulumi.StringInput
	// The value that is used for the `scp` claim in OAuth 2.0 access tokens.
	Value pulumi.StringInput
}

func (ApplicationOauth2PermissionScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationOauth2PermissionScopeArgs)(nil)).Elem()
}

type ApplicationOauth2PermissionScopeInput interface {
	pulumi.Input

	ToApplicationOauth2PermissionScopeOutput() ApplicationOauth2PermissionScopeOutput
	ToApplicationOauth2PermissionScopeOutputWithContext(ctx context.Context) ApplicationOauth2PermissionScopeOutput
}

func (*ApplicationOauth2PermissionScope) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationOauth2PermissionScope)(nil))
}

func (i *ApplicationOauth2PermissionScope) ToApplicationOauth2PermissionScopeOutput() ApplicationOauth2PermissionScopeOutput {
	return i.ToApplicationOauth2PermissionScopeOutputWithContext(context.Background())
}

func (i *ApplicationOauth2PermissionScope) ToApplicationOauth2PermissionScopeOutputWithContext(ctx context.Context) ApplicationOauth2PermissionScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOauth2PermissionScopeOutput)
}

func (i *ApplicationOauth2PermissionScope) ToApplicationOauth2PermissionScopePtrOutput() ApplicationOauth2PermissionScopePtrOutput {
	return i.ToApplicationOauth2PermissionScopePtrOutputWithContext(context.Background())
}

func (i *ApplicationOauth2PermissionScope) ToApplicationOauth2PermissionScopePtrOutputWithContext(ctx context.Context) ApplicationOauth2PermissionScopePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOauth2PermissionScopePtrOutput)
}

type ApplicationOauth2PermissionScopePtrInput interface {
	pulumi.Input

	ToApplicationOauth2PermissionScopePtrOutput() ApplicationOauth2PermissionScopePtrOutput
	ToApplicationOauth2PermissionScopePtrOutputWithContext(ctx context.Context) ApplicationOauth2PermissionScopePtrOutput
}

type applicationOauth2PermissionScopePtrType ApplicationOauth2PermissionScopeArgs

func (*applicationOauth2PermissionScopePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationOauth2PermissionScope)(nil))
}

func (i *applicationOauth2PermissionScopePtrType) ToApplicationOauth2PermissionScopePtrOutput() ApplicationOauth2PermissionScopePtrOutput {
	return i.ToApplicationOauth2PermissionScopePtrOutputWithContext(context.Background())
}

func (i *applicationOauth2PermissionScopePtrType) ToApplicationOauth2PermissionScopePtrOutputWithContext(ctx context.Context) ApplicationOauth2PermissionScopePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOauth2PermissionScopePtrOutput)
}

// ApplicationOauth2PermissionScopeArrayInput is an input type that accepts ApplicationOauth2PermissionScopeArray and ApplicationOauth2PermissionScopeArrayOutput values.
// You can construct a concrete instance of `ApplicationOauth2PermissionScopeArrayInput` via:
//
//          ApplicationOauth2PermissionScopeArray{ ApplicationOauth2PermissionScopeArgs{...} }
type ApplicationOauth2PermissionScopeArrayInput interface {
	pulumi.Input

	ToApplicationOauth2PermissionScopeArrayOutput() ApplicationOauth2PermissionScopeArrayOutput
	ToApplicationOauth2PermissionScopeArrayOutputWithContext(context.Context) ApplicationOauth2PermissionScopeArrayOutput
}

type ApplicationOauth2PermissionScopeArray []ApplicationOauth2PermissionScopeInput

func (ApplicationOauth2PermissionScopeArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ApplicationOauth2PermissionScope)(nil))
}

func (i ApplicationOauth2PermissionScopeArray) ToApplicationOauth2PermissionScopeArrayOutput() ApplicationOauth2PermissionScopeArrayOutput {
	return i.ToApplicationOauth2PermissionScopeArrayOutputWithContext(context.Background())
}

func (i ApplicationOauth2PermissionScopeArray) ToApplicationOauth2PermissionScopeArrayOutputWithContext(ctx context.Context) ApplicationOauth2PermissionScopeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOauth2PermissionScopeArrayOutput)
}

// ApplicationOauth2PermissionScopeMapInput is an input type that accepts ApplicationOauth2PermissionScopeMap and ApplicationOauth2PermissionScopeMapOutput values.
// You can construct a concrete instance of `ApplicationOauth2PermissionScopeMapInput` via:
//
//          ApplicationOauth2PermissionScopeMap{ "key": ApplicationOauth2PermissionScopeArgs{...} }
type ApplicationOauth2PermissionScopeMapInput interface {
	pulumi.Input

	ToApplicationOauth2PermissionScopeMapOutput() ApplicationOauth2PermissionScopeMapOutput
	ToApplicationOauth2PermissionScopeMapOutputWithContext(context.Context) ApplicationOauth2PermissionScopeMapOutput
}

type ApplicationOauth2PermissionScopeMap map[string]ApplicationOauth2PermissionScopeInput

func (ApplicationOauth2PermissionScopeMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ApplicationOauth2PermissionScope)(nil))
}

func (i ApplicationOauth2PermissionScopeMap) ToApplicationOauth2PermissionScopeMapOutput() ApplicationOauth2PermissionScopeMapOutput {
	return i.ToApplicationOauth2PermissionScopeMapOutputWithContext(context.Background())
}

func (i ApplicationOauth2PermissionScopeMap) ToApplicationOauth2PermissionScopeMapOutputWithContext(ctx context.Context) ApplicationOauth2PermissionScopeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOauth2PermissionScopeMapOutput)
}

type ApplicationOauth2PermissionScopeOutput struct {
	*pulumi.OutputState
}

func (ApplicationOauth2PermissionScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationOauth2PermissionScope)(nil))
}

func (o ApplicationOauth2PermissionScopeOutput) ToApplicationOauth2PermissionScopeOutput() ApplicationOauth2PermissionScopeOutput {
	return o
}

func (o ApplicationOauth2PermissionScopeOutput) ToApplicationOauth2PermissionScopeOutputWithContext(ctx context.Context) ApplicationOauth2PermissionScopeOutput {
	return o
}

func (o ApplicationOauth2PermissionScopeOutput) ToApplicationOauth2PermissionScopePtrOutput() ApplicationOauth2PermissionScopePtrOutput {
	return o.ToApplicationOauth2PermissionScopePtrOutputWithContext(context.Background())
}

func (o ApplicationOauth2PermissionScopeOutput) ToApplicationOauth2PermissionScopePtrOutputWithContext(ctx context.Context) ApplicationOauth2PermissionScopePtrOutput {
	return o.ApplyT(func(v ApplicationOauth2PermissionScope) *ApplicationOauth2PermissionScope {
		return &v
	}).(ApplicationOauth2PermissionScopePtrOutput)
}

type ApplicationOauth2PermissionScopePtrOutput struct {
	*pulumi.OutputState
}

func (ApplicationOauth2PermissionScopePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationOauth2PermissionScope)(nil))
}

func (o ApplicationOauth2PermissionScopePtrOutput) ToApplicationOauth2PermissionScopePtrOutput() ApplicationOauth2PermissionScopePtrOutput {
	return o
}

func (o ApplicationOauth2PermissionScopePtrOutput) ToApplicationOauth2PermissionScopePtrOutputWithContext(ctx context.Context) ApplicationOauth2PermissionScopePtrOutput {
	return o
}

type ApplicationOauth2PermissionScopeArrayOutput struct{ *pulumi.OutputState }

func (ApplicationOauth2PermissionScopeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationOauth2PermissionScope)(nil))
}

func (o ApplicationOauth2PermissionScopeArrayOutput) ToApplicationOauth2PermissionScopeArrayOutput() ApplicationOauth2PermissionScopeArrayOutput {
	return o
}

func (o ApplicationOauth2PermissionScopeArrayOutput) ToApplicationOauth2PermissionScopeArrayOutputWithContext(ctx context.Context) ApplicationOauth2PermissionScopeArrayOutput {
	return o
}

func (o ApplicationOauth2PermissionScopeArrayOutput) Index(i pulumi.IntInput) ApplicationOauth2PermissionScopeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationOauth2PermissionScope {
		return vs[0].([]ApplicationOauth2PermissionScope)[vs[1].(int)]
	}).(ApplicationOauth2PermissionScopeOutput)
}

type ApplicationOauth2PermissionScopeMapOutput struct{ *pulumi.OutputState }

func (ApplicationOauth2PermissionScopeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApplicationOauth2PermissionScope)(nil))
}

func (o ApplicationOauth2PermissionScopeMapOutput) ToApplicationOauth2PermissionScopeMapOutput() ApplicationOauth2PermissionScopeMapOutput {
	return o
}

func (o ApplicationOauth2PermissionScopeMapOutput) ToApplicationOauth2PermissionScopeMapOutputWithContext(ctx context.Context) ApplicationOauth2PermissionScopeMapOutput {
	return o
}

func (o ApplicationOauth2PermissionScopeMapOutput) MapIndex(k pulumi.StringInput) ApplicationOauth2PermissionScopeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ApplicationOauth2PermissionScope {
		return vs[0].(map[string]ApplicationOauth2PermissionScope)[vs[1].(string)]
	}).(ApplicationOauth2PermissionScopeOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationOauth2PermissionScopeOutput{})
	pulumi.RegisterOutputType(ApplicationOauth2PermissionScopePtrOutput{})
	pulumi.RegisterOutputType(ApplicationOauth2PermissionScopeArrayOutput{})
	pulumi.RegisterOutputType(ApplicationOauth2PermissionScopeMapOutput{})
}

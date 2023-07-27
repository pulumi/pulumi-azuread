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

// Manages a single directory role assignment within Azure Active Directory.
//
// ## API Permissions
//
// The following API permissions are required in order to use this resource.
//
// When authenticated with a service principal, this resource requires one of the following application roles: `RoleManagement.ReadWrite.Directory` or `Directory.ReadWrite.All`
//
// When authenticated with a user principal, this resource requires one of the following directory roles: `Privileged Role Administrator` or `Global Administrator`
//
// ## Import
//
// Directory role assignments can be imported using the ID of the assignment, e.g.
//
// ```sh
//
//	$ pulumi import azuread:index/directoryRoleAssignment:DirectoryRoleAssignment test ePROZI_iKE653D_d6aoLHyr-lKgHI8ZGiIdz8CLVcng-1
//
// ```
type DirectoryRoleAssignment struct {
	pulumi.CustomResourceState

	// Identifier of the app-specific scope when the assignment scope is app-specific. Cannot be used with `directoryScopeId`. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&tabs=http) for example usage. Changing this forces a new resource to be created.
	AppScopeId pulumi.StringOutput `pulumi:"appScopeId"`
	// Identifier of the app-specific scope when the assignment scope is app-specific
	//
	// Deprecated: `app_scope_object_id` has been renamed to `app_scope_id` and will be removed in version 3.0 or the AzureAD Provider
	AppScopeObjectId pulumi.StringOutput `pulumi:"appScopeObjectId"`
	// Identifier of the directory object representing the scope of the assignment. Cannot be used with `appScopeId`. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&tabs=http) for example usage. Changing this forces a new resource to be created.
	DirectoryScopeId pulumi.StringOutput `pulumi:"directoryScopeId"`
	// Identifier of the directory object representing the scope of the assignment
	DirectoryScopeObjectId pulumi.StringOutput `pulumi:"directoryScopeObjectId"`
	// The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
	PrincipalObjectId pulumi.StringOutput `pulumi:"principalObjectId"`
	// The template ID (in the case of built-in roles) or object ID (in the case of custom roles) of the directory role you want to assign. Changing this forces a new resource to be created.
	RoleId pulumi.StringOutput `pulumi:"roleId"`
}

// NewDirectoryRoleAssignment registers a new resource with the given unique name, arguments, and options.
func NewDirectoryRoleAssignment(ctx *pulumi.Context,
	name string, args *DirectoryRoleAssignmentArgs, opts ...pulumi.ResourceOption) (*DirectoryRoleAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrincipalObjectId == nil {
		return nil, errors.New("invalid value for required argument 'PrincipalObjectId'")
	}
	if args.RoleId == nil {
		return nil, errors.New("invalid value for required argument 'RoleId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DirectoryRoleAssignment
	err := ctx.RegisterResource("azuread:index/directoryRoleAssignment:DirectoryRoleAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDirectoryRoleAssignment gets an existing DirectoryRoleAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDirectoryRoleAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DirectoryRoleAssignmentState, opts ...pulumi.ResourceOption) (*DirectoryRoleAssignment, error) {
	var resource DirectoryRoleAssignment
	err := ctx.ReadResource("azuread:index/directoryRoleAssignment:DirectoryRoleAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DirectoryRoleAssignment resources.
type directoryRoleAssignmentState struct {
	// Identifier of the app-specific scope when the assignment scope is app-specific. Cannot be used with `directoryScopeId`. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&tabs=http) for example usage. Changing this forces a new resource to be created.
	AppScopeId *string `pulumi:"appScopeId"`
	// Identifier of the app-specific scope when the assignment scope is app-specific
	//
	// Deprecated: `app_scope_object_id` has been renamed to `app_scope_id` and will be removed in version 3.0 or the AzureAD Provider
	AppScopeObjectId *string `pulumi:"appScopeObjectId"`
	// Identifier of the directory object representing the scope of the assignment. Cannot be used with `appScopeId`. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&tabs=http) for example usage. Changing this forces a new resource to be created.
	DirectoryScopeId *string `pulumi:"directoryScopeId"`
	// Identifier of the directory object representing the scope of the assignment
	DirectoryScopeObjectId *string `pulumi:"directoryScopeObjectId"`
	// The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
	PrincipalObjectId *string `pulumi:"principalObjectId"`
	// The template ID (in the case of built-in roles) or object ID (in the case of custom roles) of the directory role you want to assign. Changing this forces a new resource to be created.
	RoleId *string `pulumi:"roleId"`
}

type DirectoryRoleAssignmentState struct {
	// Identifier of the app-specific scope when the assignment scope is app-specific. Cannot be used with `directoryScopeId`. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&tabs=http) for example usage. Changing this forces a new resource to be created.
	AppScopeId pulumi.StringPtrInput
	// Identifier of the app-specific scope when the assignment scope is app-specific
	//
	// Deprecated: `app_scope_object_id` has been renamed to `app_scope_id` and will be removed in version 3.0 or the AzureAD Provider
	AppScopeObjectId pulumi.StringPtrInput
	// Identifier of the directory object representing the scope of the assignment. Cannot be used with `appScopeId`. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&tabs=http) for example usage. Changing this forces a new resource to be created.
	DirectoryScopeId pulumi.StringPtrInput
	// Identifier of the directory object representing the scope of the assignment
	DirectoryScopeObjectId pulumi.StringPtrInput
	// The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
	PrincipalObjectId pulumi.StringPtrInput
	// The template ID (in the case of built-in roles) or object ID (in the case of custom roles) of the directory role you want to assign. Changing this forces a new resource to be created.
	RoleId pulumi.StringPtrInput
}

func (DirectoryRoleAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*directoryRoleAssignmentState)(nil)).Elem()
}

type directoryRoleAssignmentArgs struct {
	// Identifier of the app-specific scope when the assignment scope is app-specific. Cannot be used with `directoryScopeId`. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&tabs=http) for example usage. Changing this forces a new resource to be created.
	AppScopeId *string `pulumi:"appScopeId"`
	// Identifier of the app-specific scope when the assignment scope is app-specific
	//
	// Deprecated: `app_scope_object_id` has been renamed to `app_scope_id` and will be removed in version 3.0 or the AzureAD Provider
	AppScopeObjectId *string `pulumi:"appScopeObjectId"`
	// Identifier of the directory object representing the scope of the assignment. Cannot be used with `appScopeId`. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&tabs=http) for example usage. Changing this forces a new resource to be created.
	DirectoryScopeId *string `pulumi:"directoryScopeId"`
	// Identifier of the directory object representing the scope of the assignment
	DirectoryScopeObjectId *string `pulumi:"directoryScopeObjectId"`
	// The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
	PrincipalObjectId string `pulumi:"principalObjectId"`
	// The template ID (in the case of built-in roles) or object ID (in the case of custom roles) of the directory role you want to assign. Changing this forces a new resource to be created.
	RoleId string `pulumi:"roleId"`
}

// The set of arguments for constructing a DirectoryRoleAssignment resource.
type DirectoryRoleAssignmentArgs struct {
	// Identifier of the app-specific scope when the assignment scope is app-specific. Cannot be used with `directoryScopeId`. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&tabs=http) for example usage. Changing this forces a new resource to be created.
	AppScopeId pulumi.StringPtrInput
	// Identifier of the app-specific scope when the assignment scope is app-specific
	//
	// Deprecated: `app_scope_object_id` has been renamed to `app_scope_id` and will be removed in version 3.0 or the AzureAD Provider
	AppScopeObjectId pulumi.StringPtrInput
	// Identifier of the directory object representing the scope of the assignment. Cannot be used with `appScopeId`. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&tabs=http) for example usage. Changing this forces a new resource to be created.
	DirectoryScopeId pulumi.StringPtrInput
	// Identifier of the directory object representing the scope of the assignment
	DirectoryScopeObjectId pulumi.StringPtrInput
	// The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
	PrincipalObjectId pulumi.StringInput
	// The template ID (in the case of built-in roles) or object ID (in the case of custom roles) of the directory role you want to assign. Changing this forces a new resource to be created.
	RoleId pulumi.StringInput
}

func (DirectoryRoleAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*directoryRoleAssignmentArgs)(nil)).Elem()
}

type DirectoryRoleAssignmentInput interface {
	pulumi.Input

	ToDirectoryRoleAssignmentOutput() DirectoryRoleAssignmentOutput
	ToDirectoryRoleAssignmentOutputWithContext(ctx context.Context) DirectoryRoleAssignmentOutput
}

func (*DirectoryRoleAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**DirectoryRoleAssignment)(nil)).Elem()
}

func (i *DirectoryRoleAssignment) ToDirectoryRoleAssignmentOutput() DirectoryRoleAssignmentOutput {
	return i.ToDirectoryRoleAssignmentOutputWithContext(context.Background())
}

func (i *DirectoryRoleAssignment) ToDirectoryRoleAssignmentOutputWithContext(ctx context.Context) DirectoryRoleAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectoryRoleAssignmentOutput)
}

// DirectoryRoleAssignmentArrayInput is an input type that accepts DirectoryRoleAssignmentArray and DirectoryRoleAssignmentArrayOutput values.
// You can construct a concrete instance of `DirectoryRoleAssignmentArrayInput` via:
//
//	DirectoryRoleAssignmentArray{ DirectoryRoleAssignmentArgs{...} }
type DirectoryRoleAssignmentArrayInput interface {
	pulumi.Input

	ToDirectoryRoleAssignmentArrayOutput() DirectoryRoleAssignmentArrayOutput
	ToDirectoryRoleAssignmentArrayOutputWithContext(context.Context) DirectoryRoleAssignmentArrayOutput
}

type DirectoryRoleAssignmentArray []DirectoryRoleAssignmentInput

func (DirectoryRoleAssignmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DirectoryRoleAssignment)(nil)).Elem()
}

func (i DirectoryRoleAssignmentArray) ToDirectoryRoleAssignmentArrayOutput() DirectoryRoleAssignmentArrayOutput {
	return i.ToDirectoryRoleAssignmentArrayOutputWithContext(context.Background())
}

func (i DirectoryRoleAssignmentArray) ToDirectoryRoleAssignmentArrayOutputWithContext(ctx context.Context) DirectoryRoleAssignmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectoryRoleAssignmentArrayOutput)
}

// DirectoryRoleAssignmentMapInput is an input type that accepts DirectoryRoleAssignmentMap and DirectoryRoleAssignmentMapOutput values.
// You can construct a concrete instance of `DirectoryRoleAssignmentMapInput` via:
//
//	DirectoryRoleAssignmentMap{ "key": DirectoryRoleAssignmentArgs{...} }
type DirectoryRoleAssignmentMapInput interface {
	pulumi.Input

	ToDirectoryRoleAssignmentMapOutput() DirectoryRoleAssignmentMapOutput
	ToDirectoryRoleAssignmentMapOutputWithContext(context.Context) DirectoryRoleAssignmentMapOutput
}

type DirectoryRoleAssignmentMap map[string]DirectoryRoleAssignmentInput

func (DirectoryRoleAssignmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DirectoryRoleAssignment)(nil)).Elem()
}

func (i DirectoryRoleAssignmentMap) ToDirectoryRoleAssignmentMapOutput() DirectoryRoleAssignmentMapOutput {
	return i.ToDirectoryRoleAssignmentMapOutputWithContext(context.Background())
}

func (i DirectoryRoleAssignmentMap) ToDirectoryRoleAssignmentMapOutputWithContext(ctx context.Context) DirectoryRoleAssignmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectoryRoleAssignmentMapOutput)
}

type DirectoryRoleAssignmentOutput struct{ *pulumi.OutputState }

func (DirectoryRoleAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DirectoryRoleAssignment)(nil)).Elem()
}

func (o DirectoryRoleAssignmentOutput) ToDirectoryRoleAssignmentOutput() DirectoryRoleAssignmentOutput {
	return o
}

func (o DirectoryRoleAssignmentOutput) ToDirectoryRoleAssignmentOutputWithContext(ctx context.Context) DirectoryRoleAssignmentOutput {
	return o
}

// Identifier of the app-specific scope when the assignment scope is app-specific. Cannot be used with `directoryScopeId`. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&tabs=http) for example usage. Changing this forces a new resource to be created.
func (o DirectoryRoleAssignmentOutput) AppScopeId() pulumi.StringOutput {
	return o.ApplyT(func(v *DirectoryRoleAssignment) pulumi.StringOutput { return v.AppScopeId }).(pulumi.StringOutput)
}

// Identifier of the app-specific scope when the assignment scope is app-specific
//
// Deprecated: `app_scope_object_id` has been renamed to `app_scope_id` and will be removed in version 3.0 or the AzureAD Provider
func (o DirectoryRoleAssignmentOutput) AppScopeObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *DirectoryRoleAssignment) pulumi.StringOutput { return v.AppScopeObjectId }).(pulumi.StringOutput)
}

// Identifier of the directory object representing the scope of the assignment. Cannot be used with `appScopeId`. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&tabs=http) for example usage. Changing this forces a new resource to be created.
func (o DirectoryRoleAssignmentOutput) DirectoryScopeId() pulumi.StringOutput {
	return o.ApplyT(func(v *DirectoryRoleAssignment) pulumi.StringOutput { return v.DirectoryScopeId }).(pulumi.StringOutput)
}

// Identifier of the directory object representing the scope of the assignment
func (o DirectoryRoleAssignmentOutput) DirectoryScopeObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *DirectoryRoleAssignment) pulumi.StringOutput { return v.DirectoryScopeObjectId }).(pulumi.StringOutput)
}

// The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
func (o DirectoryRoleAssignmentOutput) PrincipalObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *DirectoryRoleAssignment) pulumi.StringOutput { return v.PrincipalObjectId }).(pulumi.StringOutput)
}

// The template ID (in the case of built-in roles) or object ID (in the case of custom roles) of the directory role you want to assign. Changing this forces a new resource to be created.
func (o DirectoryRoleAssignmentOutput) RoleId() pulumi.StringOutput {
	return o.ApplyT(func(v *DirectoryRoleAssignment) pulumi.StringOutput { return v.RoleId }).(pulumi.StringOutput)
}

type DirectoryRoleAssignmentArrayOutput struct{ *pulumi.OutputState }

func (DirectoryRoleAssignmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DirectoryRoleAssignment)(nil)).Elem()
}

func (o DirectoryRoleAssignmentArrayOutput) ToDirectoryRoleAssignmentArrayOutput() DirectoryRoleAssignmentArrayOutput {
	return o
}

func (o DirectoryRoleAssignmentArrayOutput) ToDirectoryRoleAssignmentArrayOutputWithContext(ctx context.Context) DirectoryRoleAssignmentArrayOutput {
	return o
}

func (o DirectoryRoleAssignmentArrayOutput) Index(i pulumi.IntInput) DirectoryRoleAssignmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DirectoryRoleAssignment {
		return vs[0].([]*DirectoryRoleAssignment)[vs[1].(int)]
	}).(DirectoryRoleAssignmentOutput)
}

type DirectoryRoleAssignmentMapOutput struct{ *pulumi.OutputState }

func (DirectoryRoleAssignmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DirectoryRoleAssignment)(nil)).Elem()
}

func (o DirectoryRoleAssignmentMapOutput) ToDirectoryRoleAssignmentMapOutput() DirectoryRoleAssignmentMapOutput {
	return o
}

func (o DirectoryRoleAssignmentMapOutput) ToDirectoryRoleAssignmentMapOutputWithContext(ctx context.Context) DirectoryRoleAssignmentMapOutput {
	return o
}

func (o DirectoryRoleAssignmentMapOutput) MapIndex(k pulumi.StringInput) DirectoryRoleAssignmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DirectoryRoleAssignment {
		return vs[0].(map[string]*DirectoryRoleAssignment)[vs[1].(string)]
	}).(DirectoryRoleAssignmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DirectoryRoleAssignmentInput)(nil)).Elem(), &DirectoryRoleAssignment{})
	pulumi.RegisterInputType(reflect.TypeOf((*DirectoryRoleAssignmentArrayInput)(nil)).Elem(), DirectoryRoleAssignmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DirectoryRoleAssignmentMapInput)(nil)).Elem(), DirectoryRoleAssignmentMap{})
	pulumi.RegisterOutputType(DirectoryRoleAssignmentOutput{})
	pulumi.RegisterOutputType(DirectoryRoleAssignmentArrayOutput{})
	pulumi.RegisterOutputType(DirectoryRoleAssignmentMapOutput{})
}

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

// ## Import
//
// Application API Access can be imported using the object ID of the application and the client ID of the API, in the following format.
//
// ```sh
// $ pulumi import azuread:index/applicationApiAccess:ApplicationApiAccess example /applications/00000000-0000-0000-0000-000000000000/apiAccess/11111111-1111-1111-1111-111111111111
// ```
type ApplicationApiAccess struct {
	pulumi.CustomResourceState

	// The client ID of the API to which access is being granted. Changing this forces a new resource to be created.
	ApiClientId pulumi.StringOutput `pulumi:"apiClientId"`
	// The resource ID of the application registration. Changing this forces a new resource to be created.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// A set of role IDs to be granted to the application, as published by the API.
	RoleIds pulumi.StringArrayOutput `pulumi:"roleIds"`
	// A set of scope IDs to be granted to the application, as published by the API.
	//
	// > At least one of `roleIds` or `scopeIds` must be specified.
	ScopeIds pulumi.StringArrayOutput `pulumi:"scopeIds"`
}

// NewApplicationApiAccess registers a new resource with the given unique name, arguments, and options.
func NewApplicationApiAccess(ctx *pulumi.Context,
	name string, args *ApplicationApiAccessArgs, opts ...pulumi.ResourceOption) (*ApplicationApiAccess, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiClientId == nil {
		return nil, errors.New("invalid value for required argument 'ApiClientId'")
	}
	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApplicationApiAccess
	err := ctx.RegisterResource("azuread:index/applicationApiAccess:ApplicationApiAccess", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationApiAccess gets an existing ApplicationApiAccess resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationApiAccess(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationApiAccessState, opts ...pulumi.ResourceOption) (*ApplicationApiAccess, error) {
	var resource ApplicationApiAccess
	err := ctx.ReadResource("azuread:index/applicationApiAccess:ApplicationApiAccess", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationApiAccess resources.
type applicationApiAccessState struct {
	// The client ID of the API to which access is being granted. Changing this forces a new resource to be created.
	ApiClientId *string `pulumi:"apiClientId"`
	// The resource ID of the application registration. Changing this forces a new resource to be created.
	ApplicationId *string `pulumi:"applicationId"`
	// A set of role IDs to be granted to the application, as published by the API.
	RoleIds []string `pulumi:"roleIds"`
	// A set of scope IDs to be granted to the application, as published by the API.
	//
	// > At least one of `roleIds` or `scopeIds` must be specified.
	ScopeIds []string `pulumi:"scopeIds"`
}

type ApplicationApiAccessState struct {
	// The client ID of the API to which access is being granted. Changing this forces a new resource to be created.
	ApiClientId pulumi.StringPtrInput
	// The resource ID of the application registration. Changing this forces a new resource to be created.
	ApplicationId pulumi.StringPtrInput
	// A set of role IDs to be granted to the application, as published by the API.
	RoleIds pulumi.StringArrayInput
	// A set of scope IDs to be granted to the application, as published by the API.
	//
	// > At least one of `roleIds` or `scopeIds` must be specified.
	ScopeIds pulumi.StringArrayInput
}

func (ApplicationApiAccessState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationApiAccessState)(nil)).Elem()
}

type applicationApiAccessArgs struct {
	// The client ID of the API to which access is being granted. Changing this forces a new resource to be created.
	ApiClientId string `pulumi:"apiClientId"`
	// The resource ID of the application registration. Changing this forces a new resource to be created.
	ApplicationId string `pulumi:"applicationId"`
	// A set of role IDs to be granted to the application, as published by the API.
	RoleIds []string `pulumi:"roleIds"`
	// A set of scope IDs to be granted to the application, as published by the API.
	//
	// > At least one of `roleIds` or `scopeIds` must be specified.
	ScopeIds []string `pulumi:"scopeIds"`
}

// The set of arguments for constructing a ApplicationApiAccess resource.
type ApplicationApiAccessArgs struct {
	// The client ID of the API to which access is being granted. Changing this forces a new resource to be created.
	ApiClientId pulumi.StringInput
	// The resource ID of the application registration. Changing this forces a new resource to be created.
	ApplicationId pulumi.StringInput
	// A set of role IDs to be granted to the application, as published by the API.
	RoleIds pulumi.StringArrayInput
	// A set of scope IDs to be granted to the application, as published by the API.
	//
	// > At least one of `roleIds` or `scopeIds` must be specified.
	ScopeIds pulumi.StringArrayInput
}

func (ApplicationApiAccessArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationApiAccessArgs)(nil)).Elem()
}

type ApplicationApiAccessInput interface {
	pulumi.Input

	ToApplicationApiAccessOutput() ApplicationApiAccessOutput
	ToApplicationApiAccessOutputWithContext(ctx context.Context) ApplicationApiAccessOutput
}

func (*ApplicationApiAccess) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationApiAccess)(nil)).Elem()
}

func (i *ApplicationApiAccess) ToApplicationApiAccessOutput() ApplicationApiAccessOutput {
	return i.ToApplicationApiAccessOutputWithContext(context.Background())
}

func (i *ApplicationApiAccess) ToApplicationApiAccessOutputWithContext(ctx context.Context) ApplicationApiAccessOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationApiAccessOutput)
}

// ApplicationApiAccessArrayInput is an input type that accepts ApplicationApiAccessArray and ApplicationApiAccessArrayOutput values.
// You can construct a concrete instance of `ApplicationApiAccessArrayInput` via:
//
//	ApplicationApiAccessArray{ ApplicationApiAccessArgs{...} }
type ApplicationApiAccessArrayInput interface {
	pulumi.Input

	ToApplicationApiAccessArrayOutput() ApplicationApiAccessArrayOutput
	ToApplicationApiAccessArrayOutputWithContext(context.Context) ApplicationApiAccessArrayOutput
}

type ApplicationApiAccessArray []ApplicationApiAccessInput

func (ApplicationApiAccessArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationApiAccess)(nil)).Elem()
}

func (i ApplicationApiAccessArray) ToApplicationApiAccessArrayOutput() ApplicationApiAccessArrayOutput {
	return i.ToApplicationApiAccessArrayOutputWithContext(context.Background())
}

func (i ApplicationApiAccessArray) ToApplicationApiAccessArrayOutputWithContext(ctx context.Context) ApplicationApiAccessArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationApiAccessArrayOutput)
}

// ApplicationApiAccessMapInput is an input type that accepts ApplicationApiAccessMap and ApplicationApiAccessMapOutput values.
// You can construct a concrete instance of `ApplicationApiAccessMapInput` via:
//
//	ApplicationApiAccessMap{ "key": ApplicationApiAccessArgs{...} }
type ApplicationApiAccessMapInput interface {
	pulumi.Input

	ToApplicationApiAccessMapOutput() ApplicationApiAccessMapOutput
	ToApplicationApiAccessMapOutputWithContext(context.Context) ApplicationApiAccessMapOutput
}

type ApplicationApiAccessMap map[string]ApplicationApiAccessInput

func (ApplicationApiAccessMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationApiAccess)(nil)).Elem()
}

func (i ApplicationApiAccessMap) ToApplicationApiAccessMapOutput() ApplicationApiAccessMapOutput {
	return i.ToApplicationApiAccessMapOutputWithContext(context.Background())
}

func (i ApplicationApiAccessMap) ToApplicationApiAccessMapOutputWithContext(ctx context.Context) ApplicationApiAccessMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationApiAccessMapOutput)
}

type ApplicationApiAccessOutput struct{ *pulumi.OutputState }

func (ApplicationApiAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationApiAccess)(nil)).Elem()
}

func (o ApplicationApiAccessOutput) ToApplicationApiAccessOutput() ApplicationApiAccessOutput {
	return o
}

func (o ApplicationApiAccessOutput) ToApplicationApiAccessOutputWithContext(ctx context.Context) ApplicationApiAccessOutput {
	return o
}

// The client ID of the API to which access is being granted. Changing this forces a new resource to be created.
func (o ApplicationApiAccessOutput) ApiClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationApiAccess) pulumi.StringOutput { return v.ApiClientId }).(pulumi.StringOutput)
}

// The resource ID of the application registration. Changing this forces a new resource to be created.
func (o ApplicationApiAccessOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationApiAccess) pulumi.StringOutput { return v.ApplicationId }).(pulumi.StringOutput)
}

// A set of role IDs to be granted to the application, as published by the API.
func (o ApplicationApiAccessOutput) RoleIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationApiAccess) pulumi.StringArrayOutput { return v.RoleIds }).(pulumi.StringArrayOutput)
}

// A set of scope IDs to be granted to the application, as published by the API.
//
// > At least one of `roleIds` or `scopeIds` must be specified.
func (o ApplicationApiAccessOutput) ScopeIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationApiAccess) pulumi.StringArrayOutput { return v.ScopeIds }).(pulumi.StringArrayOutput)
}

type ApplicationApiAccessArrayOutput struct{ *pulumi.OutputState }

func (ApplicationApiAccessArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationApiAccess)(nil)).Elem()
}

func (o ApplicationApiAccessArrayOutput) ToApplicationApiAccessArrayOutput() ApplicationApiAccessArrayOutput {
	return o
}

func (o ApplicationApiAccessArrayOutput) ToApplicationApiAccessArrayOutputWithContext(ctx context.Context) ApplicationApiAccessArrayOutput {
	return o
}

func (o ApplicationApiAccessArrayOutput) Index(i pulumi.IntInput) ApplicationApiAccessOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplicationApiAccess {
		return vs[0].([]*ApplicationApiAccess)[vs[1].(int)]
	}).(ApplicationApiAccessOutput)
}

type ApplicationApiAccessMapOutput struct{ *pulumi.OutputState }

func (ApplicationApiAccessMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationApiAccess)(nil)).Elem()
}

func (o ApplicationApiAccessMapOutput) ToApplicationApiAccessMapOutput() ApplicationApiAccessMapOutput {
	return o
}

func (o ApplicationApiAccessMapOutput) ToApplicationApiAccessMapOutputWithContext(ctx context.Context) ApplicationApiAccessMapOutput {
	return o
}

func (o ApplicationApiAccessMapOutput) MapIndex(k pulumi.StringInput) ApplicationApiAccessOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplicationApiAccess {
		return vs[0].(map[string]*ApplicationApiAccess)[vs[1].(string)]
	}).(ApplicationApiAccessOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationApiAccessInput)(nil)).Elem(), &ApplicationApiAccess{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationApiAccessArrayInput)(nil)).Elem(), ApplicationApiAccessArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationApiAccessMapInput)(nil)).Elem(), ApplicationApiAccessMap{})
	pulumi.RegisterOutputType(ApplicationApiAccessOutput{})
	pulumi.RegisterOutputType(ApplicationApiAccessArrayOutput{})
	pulumi.RegisterOutputType(ApplicationApiAccessMapOutput{})
}

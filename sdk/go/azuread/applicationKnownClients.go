// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azuread/sdk/v6/go/azuread/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-azuread/sdk/v6/go/azuread"
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
//			client, err := azuread.NewApplicationRegistration(ctx, "client", &azuread.ApplicationRegistrationArgs{
//				DisplayName: pulumi.String("example client"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuread.NewApplicationKnownClients(ctx, "example", &azuread.ApplicationKnownClientsArgs{
//				ApplicationId: example.ID(),
//				KnownClientIds: pulumi.StringArray{
//					client.ClientId,
//				},
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
// Application Known Clients can be imported using the object ID of the application in the following format.
//
// ```sh
// $ pulumi import azuread:index/applicationKnownClients:ApplicationKnownClients example /applications/00000000-0000-0000-0000-000000000000/knownClients
// ```
type ApplicationKnownClients struct {
	pulumi.CustomResourceState

	// The resource ID of the application registration. Changing this forces a new resource to be created.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// A set of client IDs for the known applications.
	KnownClientIds pulumi.StringArrayOutput `pulumi:"knownClientIds"`
}

// NewApplicationKnownClients registers a new resource with the given unique name, arguments, and options.
func NewApplicationKnownClients(ctx *pulumi.Context,
	name string, args *ApplicationKnownClientsArgs, opts ...pulumi.ResourceOption) (*ApplicationKnownClients, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	if args.KnownClientIds == nil {
		return nil, errors.New("invalid value for required argument 'KnownClientIds'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApplicationKnownClients
	err := ctx.RegisterResource("azuread:index/applicationKnownClients:ApplicationKnownClients", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationKnownClients gets an existing ApplicationKnownClients resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationKnownClients(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationKnownClientsState, opts ...pulumi.ResourceOption) (*ApplicationKnownClients, error) {
	var resource ApplicationKnownClients
	err := ctx.ReadResource("azuread:index/applicationKnownClients:ApplicationKnownClients", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationKnownClients resources.
type applicationKnownClientsState struct {
	// The resource ID of the application registration. Changing this forces a new resource to be created.
	ApplicationId *string `pulumi:"applicationId"`
	// A set of client IDs for the known applications.
	KnownClientIds []string `pulumi:"knownClientIds"`
}

type ApplicationKnownClientsState struct {
	// The resource ID of the application registration. Changing this forces a new resource to be created.
	ApplicationId pulumi.StringPtrInput
	// A set of client IDs for the known applications.
	KnownClientIds pulumi.StringArrayInput
}

func (ApplicationKnownClientsState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationKnownClientsState)(nil)).Elem()
}

type applicationKnownClientsArgs struct {
	// The resource ID of the application registration. Changing this forces a new resource to be created.
	ApplicationId string `pulumi:"applicationId"`
	// A set of client IDs for the known applications.
	KnownClientIds []string `pulumi:"knownClientIds"`
}

// The set of arguments for constructing a ApplicationKnownClients resource.
type ApplicationKnownClientsArgs struct {
	// The resource ID of the application registration. Changing this forces a new resource to be created.
	ApplicationId pulumi.StringInput
	// A set of client IDs for the known applications.
	KnownClientIds pulumi.StringArrayInput
}

func (ApplicationKnownClientsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationKnownClientsArgs)(nil)).Elem()
}

type ApplicationKnownClientsInput interface {
	pulumi.Input

	ToApplicationKnownClientsOutput() ApplicationKnownClientsOutput
	ToApplicationKnownClientsOutputWithContext(ctx context.Context) ApplicationKnownClientsOutput
}

func (*ApplicationKnownClients) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationKnownClients)(nil)).Elem()
}

func (i *ApplicationKnownClients) ToApplicationKnownClientsOutput() ApplicationKnownClientsOutput {
	return i.ToApplicationKnownClientsOutputWithContext(context.Background())
}

func (i *ApplicationKnownClients) ToApplicationKnownClientsOutputWithContext(ctx context.Context) ApplicationKnownClientsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationKnownClientsOutput)
}

// ApplicationKnownClientsArrayInput is an input type that accepts ApplicationKnownClientsArray and ApplicationKnownClientsArrayOutput values.
// You can construct a concrete instance of `ApplicationKnownClientsArrayInput` via:
//
//	ApplicationKnownClientsArray{ ApplicationKnownClientsArgs{...} }
type ApplicationKnownClientsArrayInput interface {
	pulumi.Input

	ToApplicationKnownClientsArrayOutput() ApplicationKnownClientsArrayOutput
	ToApplicationKnownClientsArrayOutputWithContext(context.Context) ApplicationKnownClientsArrayOutput
}

type ApplicationKnownClientsArray []ApplicationKnownClientsInput

func (ApplicationKnownClientsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationKnownClients)(nil)).Elem()
}

func (i ApplicationKnownClientsArray) ToApplicationKnownClientsArrayOutput() ApplicationKnownClientsArrayOutput {
	return i.ToApplicationKnownClientsArrayOutputWithContext(context.Background())
}

func (i ApplicationKnownClientsArray) ToApplicationKnownClientsArrayOutputWithContext(ctx context.Context) ApplicationKnownClientsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationKnownClientsArrayOutput)
}

// ApplicationKnownClientsMapInput is an input type that accepts ApplicationKnownClientsMap and ApplicationKnownClientsMapOutput values.
// You can construct a concrete instance of `ApplicationKnownClientsMapInput` via:
//
//	ApplicationKnownClientsMap{ "key": ApplicationKnownClientsArgs{...} }
type ApplicationKnownClientsMapInput interface {
	pulumi.Input

	ToApplicationKnownClientsMapOutput() ApplicationKnownClientsMapOutput
	ToApplicationKnownClientsMapOutputWithContext(context.Context) ApplicationKnownClientsMapOutput
}

type ApplicationKnownClientsMap map[string]ApplicationKnownClientsInput

func (ApplicationKnownClientsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationKnownClients)(nil)).Elem()
}

func (i ApplicationKnownClientsMap) ToApplicationKnownClientsMapOutput() ApplicationKnownClientsMapOutput {
	return i.ToApplicationKnownClientsMapOutputWithContext(context.Background())
}

func (i ApplicationKnownClientsMap) ToApplicationKnownClientsMapOutputWithContext(ctx context.Context) ApplicationKnownClientsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationKnownClientsMapOutput)
}

type ApplicationKnownClientsOutput struct{ *pulumi.OutputState }

func (ApplicationKnownClientsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationKnownClients)(nil)).Elem()
}

func (o ApplicationKnownClientsOutput) ToApplicationKnownClientsOutput() ApplicationKnownClientsOutput {
	return o
}

func (o ApplicationKnownClientsOutput) ToApplicationKnownClientsOutputWithContext(ctx context.Context) ApplicationKnownClientsOutput {
	return o
}

// The resource ID of the application registration. Changing this forces a new resource to be created.
func (o ApplicationKnownClientsOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationKnownClients) pulumi.StringOutput { return v.ApplicationId }).(pulumi.StringOutput)
}

// A set of client IDs for the known applications.
func (o ApplicationKnownClientsOutput) KnownClientIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationKnownClients) pulumi.StringArrayOutput { return v.KnownClientIds }).(pulumi.StringArrayOutput)
}

type ApplicationKnownClientsArrayOutput struct{ *pulumi.OutputState }

func (ApplicationKnownClientsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationKnownClients)(nil)).Elem()
}

func (o ApplicationKnownClientsArrayOutput) ToApplicationKnownClientsArrayOutput() ApplicationKnownClientsArrayOutput {
	return o
}

func (o ApplicationKnownClientsArrayOutput) ToApplicationKnownClientsArrayOutputWithContext(ctx context.Context) ApplicationKnownClientsArrayOutput {
	return o
}

func (o ApplicationKnownClientsArrayOutput) Index(i pulumi.IntInput) ApplicationKnownClientsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplicationKnownClients {
		return vs[0].([]*ApplicationKnownClients)[vs[1].(int)]
	}).(ApplicationKnownClientsOutput)
}

type ApplicationKnownClientsMapOutput struct{ *pulumi.OutputState }

func (ApplicationKnownClientsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationKnownClients)(nil)).Elem()
}

func (o ApplicationKnownClientsMapOutput) ToApplicationKnownClientsMapOutput() ApplicationKnownClientsMapOutput {
	return o
}

func (o ApplicationKnownClientsMapOutput) ToApplicationKnownClientsMapOutputWithContext(ctx context.Context) ApplicationKnownClientsMapOutput {
	return o
}

func (o ApplicationKnownClientsMapOutput) MapIndex(k pulumi.StringInput) ApplicationKnownClientsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplicationKnownClients {
		return vs[0].(map[string]*ApplicationKnownClients)[vs[1].(string)]
	}).(ApplicationKnownClientsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationKnownClientsInput)(nil)).Elem(), &ApplicationKnownClients{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationKnownClientsArrayInput)(nil)).Elem(), ApplicationKnownClientsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationKnownClientsMapInput)(nil)).Elem(), ApplicationKnownClientsMap{})
	pulumi.RegisterOutputType(ApplicationKnownClientsOutput{})
	pulumi.RegisterOutputType(ApplicationKnownClientsArrayOutput{})
	pulumi.RegisterOutputType(ApplicationKnownClientsMapOutput{})
}

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
//			jane, err := azuread.NewUser(ctx, "jane", &azuread.UserArgs{
//				UserPrincipalName: pulumi.String("jane.fischer@hashitown.com"),
//				DisplayName:       pulumi.String("Jane Fischer"),
//				Password:          pulumi.String("Ch@ngeMe"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuread.NewApplicationOwner(ctx, "exampleJane", &azuread.ApplicationOwnerArgs{
//				ApplicationId: example.ID(),
//				OwnerObjectId: jane.ObjectId,
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
// > **Tip** For managing more application owners, create additional instances of this resource
//
// ## Import
//
// Application Owners can be imported using the object ID of the application and the object ID of the owner, in the following format.
//
// ```sh
// $ pulumi import azuread:index/applicationOwner:ApplicationOwner example /applications/00000000-0000-0000-0000-000000000000/owners/11111111-1111-1111-1111-111111111111
// ```
type ApplicationOwner struct {
	pulumi.CustomResourceState

	// The resource ID of the application registration. Changing this forces a new resource to be created.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// The object ID of the owner to assign to the application, typically a user or service principal. Changing this forces a new resource to be created.
	OwnerObjectId pulumi.StringOutput `pulumi:"ownerObjectId"`
}

// NewApplicationOwner registers a new resource with the given unique name, arguments, and options.
func NewApplicationOwner(ctx *pulumi.Context,
	name string, args *ApplicationOwnerArgs, opts ...pulumi.ResourceOption) (*ApplicationOwner, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	if args.OwnerObjectId == nil {
		return nil, errors.New("invalid value for required argument 'OwnerObjectId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApplicationOwner
	err := ctx.RegisterResource("azuread:index/applicationOwner:ApplicationOwner", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationOwner gets an existing ApplicationOwner resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationOwner(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationOwnerState, opts ...pulumi.ResourceOption) (*ApplicationOwner, error) {
	var resource ApplicationOwner
	err := ctx.ReadResource("azuread:index/applicationOwner:ApplicationOwner", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationOwner resources.
type applicationOwnerState struct {
	// The resource ID of the application registration. Changing this forces a new resource to be created.
	ApplicationId *string `pulumi:"applicationId"`
	// The object ID of the owner to assign to the application, typically a user or service principal. Changing this forces a new resource to be created.
	OwnerObjectId *string `pulumi:"ownerObjectId"`
}

type ApplicationOwnerState struct {
	// The resource ID of the application registration. Changing this forces a new resource to be created.
	ApplicationId pulumi.StringPtrInput
	// The object ID of the owner to assign to the application, typically a user or service principal. Changing this forces a new resource to be created.
	OwnerObjectId pulumi.StringPtrInput
}

func (ApplicationOwnerState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationOwnerState)(nil)).Elem()
}

type applicationOwnerArgs struct {
	// The resource ID of the application registration. Changing this forces a new resource to be created.
	ApplicationId string `pulumi:"applicationId"`
	// The object ID of the owner to assign to the application, typically a user or service principal. Changing this forces a new resource to be created.
	OwnerObjectId string `pulumi:"ownerObjectId"`
}

// The set of arguments for constructing a ApplicationOwner resource.
type ApplicationOwnerArgs struct {
	// The resource ID of the application registration. Changing this forces a new resource to be created.
	ApplicationId pulumi.StringInput
	// The object ID of the owner to assign to the application, typically a user or service principal. Changing this forces a new resource to be created.
	OwnerObjectId pulumi.StringInput
}

func (ApplicationOwnerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationOwnerArgs)(nil)).Elem()
}

type ApplicationOwnerInput interface {
	pulumi.Input

	ToApplicationOwnerOutput() ApplicationOwnerOutput
	ToApplicationOwnerOutputWithContext(ctx context.Context) ApplicationOwnerOutput
}

func (*ApplicationOwner) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationOwner)(nil)).Elem()
}

func (i *ApplicationOwner) ToApplicationOwnerOutput() ApplicationOwnerOutput {
	return i.ToApplicationOwnerOutputWithContext(context.Background())
}

func (i *ApplicationOwner) ToApplicationOwnerOutputWithContext(ctx context.Context) ApplicationOwnerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOwnerOutput)
}

// ApplicationOwnerArrayInput is an input type that accepts ApplicationOwnerArray and ApplicationOwnerArrayOutput values.
// You can construct a concrete instance of `ApplicationOwnerArrayInput` via:
//
//	ApplicationOwnerArray{ ApplicationOwnerArgs{...} }
type ApplicationOwnerArrayInput interface {
	pulumi.Input

	ToApplicationOwnerArrayOutput() ApplicationOwnerArrayOutput
	ToApplicationOwnerArrayOutputWithContext(context.Context) ApplicationOwnerArrayOutput
}

type ApplicationOwnerArray []ApplicationOwnerInput

func (ApplicationOwnerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationOwner)(nil)).Elem()
}

func (i ApplicationOwnerArray) ToApplicationOwnerArrayOutput() ApplicationOwnerArrayOutput {
	return i.ToApplicationOwnerArrayOutputWithContext(context.Background())
}

func (i ApplicationOwnerArray) ToApplicationOwnerArrayOutputWithContext(ctx context.Context) ApplicationOwnerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOwnerArrayOutput)
}

// ApplicationOwnerMapInput is an input type that accepts ApplicationOwnerMap and ApplicationOwnerMapOutput values.
// You can construct a concrete instance of `ApplicationOwnerMapInput` via:
//
//	ApplicationOwnerMap{ "key": ApplicationOwnerArgs{...} }
type ApplicationOwnerMapInput interface {
	pulumi.Input

	ToApplicationOwnerMapOutput() ApplicationOwnerMapOutput
	ToApplicationOwnerMapOutputWithContext(context.Context) ApplicationOwnerMapOutput
}

type ApplicationOwnerMap map[string]ApplicationOwnerInput

func (ApplicationOwnerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationOwner)(nil)).Elem()
}

func (i ApplicationOwnerMap) ToApplicationOwnerMapOutput() ApplicationOwnerMapOutput {
	return i.ToApplicationOwnerMapOutputWithContext(context.Background())
}

func (i ApplicationOwnerMap) ToApplicationOwnerMapOutputWithContext(ctx context.Context) ApplicationOwnerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOwnerMapOutput)
}

type ApplicationOwnerOutput struct{ *pulumi.OutputState }

func (ApplicationOwnerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationOwner)(nil)).Elem()
}

func (o ApplicationOwnerOutput) ToApplicationOwnerOutput() ApplicationOwnerOutput {
	return o
}

func (o ApplicationOwnerOutput) ToApplicationOwnerOutputWithContext(ctx context.Context) ApplicationOwnerOutput {
	return o
}

// The resource ID of the application registration. Changing this forces a new resource to be created.
func (o ApplicationOwnerOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationOwner) pulumi.StringOutput { return v.ApplicationId }).(pulumi.StringOutput)
}

// The object ID of the owner to assign to the application, typically a user or service principal. Changing this forces a new resource to be created.
func (o ApplicationOwnerOutput) OwnerObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationOwner) pulumi.StringOutput { return v.OwnerObjectId }).(pulumi.StringOutput)
}

type ApplicationOwnerArrayOutput struct{ *pulumi.OutputState }

func (ApplicationOwnerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationOwner)(nil)).Elem()
}

func (o ApplicationOwnerArrayOutput) ToApplicationOwnerArrayOutput() ApplicationOwnerArrayOutput {
	return o
}

func (o ApplicationOwnerArrayOutput) ToApplicationOwnerArrayOutputWithContext(ctx context.Context) ApplicationOwnerArrayOutput {
	return o
}

func (o ApplicationOwnerArrayOutput) Index(i pulumi.IntInput) ApplicationOwnerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplicationOwner {
		return vs[0].([]*ApplicationOwner)[vs[1].(int)]
	}).(ApplicationOwnerOutput)
}

type ApplicationOwnerMapOutput struct{ *pulumi.OutputState }

func (ApplicationOwnerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationOwner)(nil)).Elem()
}

func (o ApplicationOwnerMapOutput) ToApplicationOwnerMapOutput() ApplicationOwnerMapOutput {
	return o
}

func (o ApplicationOwnerMapOutput) ToApplicationOwnerMapOutputWithContext(ctx context.Context) ApplicationOwnerMapOutput {
	return o
}

func (o ApplicationOwnerMapOutput) MapIndex(k pulumi.StringInput) ApplicationOwnerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplicationOwner {
		return vs[0].(map[string]*ApplicationOwner)[vs[1].(string)]
	}).(ApplicationOwnerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationOwnerInput)(nil)).Elem(), &ApplicationOwner{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationOwnerArrayInput)(nil)).Elem(), ApplicationOwnerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationOwnerMapInput)(nil)).Elem(), ApplicationOwnerMap{})
	pulumi.RegisterOutputType(ApplicationOwnerOutput{})
	pulumi.RegisterOutputType(ApplicationOwnerArrayOutput{})
	pulumi.RegisterOutputType(ApplicationOwnerMapOutput{})
}

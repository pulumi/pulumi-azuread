// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azuread/sdk/v5/go/azuread/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
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
//			_, err = azuread.NewApplicationRedirectUris(ctx, "examplePublic", &azuread.ApplicationRedirectUrisArgs{
//				ApplicationId: example.ID(),
//				Type:          pulumi.String("PublicClient"),
//				RedirectUris: pulumi.StringArray{
//					pulumi.String("myapp://auth"),
//					pulumi.String("sample.mobile.app.bundie.id://auth"),
//					pulumi.String("https://login.microsoftonline.com/common/oauth2/nativeclient"),
//					pulumi.String("https://login.live.com/oauth20_desktop.srf"),
//					pulumi.String("ms-appx-web://Microsoft.AAD.BrokerPlugin/00000000-1111-1111-1111-222222222222"),
//					pulumi.String("urn:ietf:wg:oauth:2.0:foo"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuread.NewApplicationRedirectUris(ctx, "exampleSpa", &azuread.ApplicationRedirectUrisArgs{
//				ApplicationId: example.ID(),
//				Type:          pulumi.String("SPA"),
//				RedirectUris: pulumi.StringArray{
//					pulumi.String("https://mobile.hashitown.com/"),
//					pulumi.String("https://beta.hashitown.com/"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuread.NewApplicationRedirectUris(ctx, "exampleWeb", &azuread.ApplicationRedirectUrisArgs{
//				ApplicationId: example.ID(),
//				Type:          pulumi.String("Web"),
//				RedirectUris: pulumi.StringArray{
//					pulumi.String("https://app.hashitown.com/"),
//					pulumi.String("https://classic.hashitown.com/"),
//					pulumi.String("urn:ietf:wg:oauth:2.0:oob"),
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
// Application API Access can be imported using the object ID of the application and the URI type, in the following format.
//
// ```sh
//
//	$ pulumi import azuread:index/applicationRedirectUris:ApplicationRedirectUris example /applications/00000000-0000-0000-0000-000000000000/uriType/Web
//
// ```
type ApplicationRedirectUris struct {
	pulumi.CustomResourceState

	// The resource ID of the application registration. Changing this forces a new resource to be created.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// A set of redirect URIs to assign to the application.
	RedirectUris pulumi.StringArrayOutput `pulumi:"redirectUris"`
	// The type of redirect URIs to manage. Must be one of: `PublicClient`, `SPA`, or `Web`. Changing this forces a new resource to be created.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApplicationRedirectUris registers a new resource with the given unique name, arguments, and options.
func NewApplicationRedirectUris(ctx *pulumi.Context,
	name string, args *ApplicationRedirectUrisArgs, opts ...pulumi.ResourceOption) (*ApplicationRedirectUris, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	if args.RedirectUris == nil {
		return nil, errors.New("invalid value for required argument 'RedirectUris'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApplicationRedirectUris
	err := ctx.RegisterResource("azuread:index/applicationRedirectUris:ApplicationRedirectUris", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationRedirectUris gets an existing ApplicationRedirectUris resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationRedirectUris(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationRedirectUrisState, opts ...pulumi.ResourceOption) (*ApplicationRedirectUris, error) {
	var resource ApplicationRedirectUris
	err := ctx.ReadResource("azuread:index/applicationRedirectUris:ApplicationRedirectUris", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationRedirectUris resources.
type applicationRedirectUrisState struct {
	// The resource ID of the application registration. Changing this forces a new resource to be created.
	ApplicationId *string `pulumi:"applicationId"`
	// A set of redirect URIs to assign to the application.
	RedirectUris []string `pulumi:"redirectUris"`
	// The type of redirect URIs to manage. Must be one of: `PublicClient`, `SPA`, or `Web`. Changing this forces a new resource to be created.
	Type *string `pulumi:"type"`
}

type ApplicationRedirectUrisState struct {
	// The resource ID of the application registration. Changing this forces a new resource to be created.
	ApplicationId pulumi.StringPtrInput
	// A set of redirect URIs to assign to the application.
	RedirectUris pulumi.StringArrayInput
	// The type of redirect URIs to manage. Must be one of: `PublicClient`, `SPA`, or `Web`. Changing this forces a new resource to be created.
	Type pulumi.StringPtrInput
}

func (ApplicationRedirectUrisState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationRedirectUrisState)(nil)).Elem()
}

type applicationRedirectUrisArgs struct {
	// The resource ID of the application registration. Changing this forces a new resource to be created.
	ApplicationId string `pulumi:"applicationId"`
	// A set of redirect URIs to assign to the application.
	RedirectUris []string `pulumi:"redirectUris"`
	// The type of redirect URIs to manage. Must be one of: `PublicClient`, `SPA`, or `Web`. Changing this forces a new resource to be created.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a ApplicationRedirectUris resource.
type ApplicationRedirectUrisArgs struct {
	// The resource ID of the application registration. Changing this forces a new resource to be created.
	ApplicationId pulumi.StringInput
	// A set of redirect URIs to assign to the application.
	RedirectUris pulumi.StringArrayInput
	// The type of redirect URIs to manage. Must be one of: `PublicClient`, `SPA`, or `Web`. Changing this forces a new resource to be created.
	Type pulumi.StringInput
}

func (ApplicationRedirectUrisArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationRedirectUrisArgs)(nil)).Elem()
}

type ApplicationRedirectUrisInput interface {
	pulumi.Input

	ToApplicationRedirectUrisOutput() ApplicationRedirectUrisOutput
	ToApplicationRedirectUrisOutputWithContext(ctx context.Context) ApplicationRedirectUrisOutput
}

func (*ApplicationRedirectUris) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationRedirectUris)(nil)).Elem()
}

func (i *ApplicationRedirectUris) ToApplicationRedirectUrisOutput() ApplicationRedirectUrisOutput {
	return i.ToApplicationRedirectUrisOutputWithContext(context.Background())
}

func (i *ApplicationRedirectUris) ToApplicationRedirectUrisOutputWithContext(ctx context.Context) ApplicationRedirectUrisOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationRedirectUrisOutput)
}

func (i *ApplicationRedirectUris) ToOutput(ctx context.Context) pulumix.Output[*ApplicationRedirectUris] {
	return pulumix.Output[*ApplicationRedirectUris]{
		OutputState: i.ToApplicationRedirectUrisOutputWithContext(ctx).OutputState,
	}
}

// ApplicationRedirectUrisArrayInput is an input type that accepts ApplicationRedirectUrisArray and ApplicationRedirectUrisArrayOutput values.
// You can construct a concrete instance of `ApplicationRedirectUrisArrayInput` via:
//
//	ApplicationRedirectUrisArray{ ApplicationRedirectUrisArgs{...} }
type ApplicationRedirectUrisArrayInput interface {
	pulumi.Input

	ToApplicationRedirectUrisArrayOutput() ApplicationRedirectUrisArrayOutput
	ToApplicationRedirectUrisArrayOutputWithContext(context.Context) ApplicationRedirectUrisArrayOutput
}

type ApplicationRedirectUrisArray []ApplicationRedirectUrisInput

func (ApplicationRedirectUrisArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationRedirectUris)(nil)).Elem()
}

func (i ApplicationRedirectUrisArray) ToApplicationRedirectUrisArrayOutput() ApplicationRedirectUrisArrayOutput {
	return i.ToApplicationRedirectUrisArrayOutputWithContext(context.Background())
}

func (i ApplicationRedirectUrisArray) ToApplicationRedirectUrisArrayOutputWithContext(ctx context.Context) ApplicationRedirectUrisArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationRedirectUrisArrayOutput)
}

func (i ApplicationRedirectUrisArray) ToOutput(ctx context.Context) pulumix.Output[[]*ApplicationRedirectUris] {
	return pulumix.Output[[]*ApplicationRedirectUris]{
		OutputState: i.ToApplicationRedirectUrisArrayOutputWithContext(ctx).OutputState,
	}
}

// ApplicationRedirectUrisMapInput is an input type that accepts ApplicationRedirectUrisMap and ApplicationRedirectUrisMapOutput values.
// You can construct a concrete instance of `ApplicationRedirectUrisMapInput` via:
//
//	ApplicationRedirectUrisMap{ "key": ApplicationRedirectUrisArgs{...} }
type ApplicationRedirectUrisMapInput interface {
	pulumi.Input

	ToApplicationRedirectUrisMapOutput() ApplicationRedirectUrisMapOutput
	ToApplicationRedirectUrisMapOutputWithContext(context.Context) ApplicationRedirectUrisMapOutput
}

type ApplicationRedirectUrisMap map[string]ApplicationRedirectUrisInput

func (ApplicationRedirectUrisMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationRedirectUris)(nil)).Elem()
}

func (i ApplicationRedirectUrisMap) ToApplicationRedirectUrisMapOutput() ApplicationRedirectUrisMapOutput {
	return i.ToApplicationRedirectUrisMapOutputWithContext(context.Background())
}

func (i ApplicationRedirectUrisMap) ToApplicationRedirectUrisMapOutputWithContext(ctx context.Context) ApplicationRedirectUrisMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationRedirectUrisMapOutput)
}

func (i ApplicationRedirectUrisMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ApplicationRedirectUris] {
	return pulumix.Output[map[string]*ApplicationRedirectUris]{
		OutputState: i.ToApplicationRedirectUrisMapOutputWithContext(ctx).OutputState,
	}
}

type ApplicationRedirectUrisOutput struct{ *pulumi.OutputState }

func (ApplicationRedirectUrisOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationRedirectUris)(nil)).Elem()
}

func (o ApplicationRedirectUrisOutput) ToApplicationRedirectUrisOutput() ApplicationRedirectUrisOutput {
	return o
}

func (o ApplicationRedirectUrisOutput) ToApplicationRedirectUrisOutputWithContext(ctx context.Context) ApplicationRedirectUrisOutput {
	return o
}

func (o ApplicationRedirectUrisOutput) ToOutput(ctx context.Context) pulumix.Output[*ApplicationRedirectUris] {
	return pulumix.Output[*ApplicationRedirectUris]{
		OutputState: o.OutputState,
	}
}

// The resource ID of the application registration. Changing this forces a new resource to be created.
func (o ApplicationRedirectUrisOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationRedirectUris) pulumi.StringOutput { return v.ApplicationId }).(pulumi.StringOutput)
}

// A set of redirect URIs to assign to the application.
func (o ApplicationRedirectUrisOutput) RedirectUris() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationRedirectUris) pulumi.StringArrayOutput { return v.RedirectUris }).(pulumi.StringArrayOutput)
}

// The type of redirect URIs to manage. Must be one of: `PublicClient`, `SPA`, or `Web`. Changing this forces a new resource to be created.
func (o ApplicationRedirectUrisOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationRedirectUris) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type ApplicationRedirectUrisArrayOutput struct{ *pulumi.OutputState }

func (ApplicationRedirectUrisArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationRedirectUris)(nil)).Elem()
}

func (o ApplicationRedirectUrisArrayOutput) ToApplicationRedirectUrisArrayOutput() ApplicationRedirectUrisArrayOutput {
	return o
}

func (o ApplicationRedirectUrisArrayOutput) ToApplicationRedirectUrisArrayOutputWithContext(ctx context.Context) ApplicationRedirectUrisArrayOutput {
	return o
}

func (o ApplicationRedirectUrisArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ApplicationRedirectUris] {
	return pulumix.Output[[]*ApplicationRedirectUris]{
		OutputState: o.OutputState,
	}
}

func (o ApplicationRedirectUrisArrayOutput) Index(i pulumi.IntInput) ApplicationRedirectUrisOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplicationRedirectUris {
		return vs[0].([]*ApplicationRedirectUris)[vs[1].(int)]
	}).(ApplicationRedirectUrisOutput)
}

type ApplicationRedirectUrisMapOutput struct{ *pulumi.OutputState }

func (ApplicationRedirectUrisMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationRedirectUris)(nil)).Elem()
}

func (o ApplicationRedirectUrisMapOutput) ToApplicationRedirectUrisMapOutput() ApplicationRedirectUrisMapOutput {
	return o
}

func (o ApplicationRedirectUrisMapOutput) ToApplicationRedirectUrisMapOutputWithContext(ctx context.Context) ApplicationRedirectUrisMapOutput {
	return o
}

func (o ApplicationRedirectUrisMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ApplicationRedirectUris] {
	return pulumix.Output[map[string]*ApplicationRedirectUris]{
		OutputState: o.OutputState,
	}
}

func (o ApplicationRedirectUrisMapOutput) MapIndex(k pulumi.StringInput) ApplicationRedirectUrisOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplicationRedirectUris {
		return vs[0].(map[string]*ApplicationRedirectUris)[vs[1].(string)]
	}).(ApplicationRedirectUrisOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationRedirectUrisInput)(nil)).Elem(), &ApplicationRedirectUris{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationRedirectUrisArrayInput)(nil)).Elem(), ApplicationRedirectUrisArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationRedirectUrisMapInput)(nil)).Elem(), ApplicationRedirectUrisMap{})
	pulumi.RegisterOutputType(ApplicationRedirectUrisOutput{})
	pulumi.RegisterOutputType(ApplicationRedirectUrisArrayOutput{})
	pulumi.RegisterOutputType(ApplicationRedirectUrisMapOutput{})
}
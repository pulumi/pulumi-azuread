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
// *Basic example*
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
//			_, err = azuread.NewApplicationPassword(ctx, "example", &azuread.ApplicationPasswordArgs{
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
// *Time-based rotation*
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-azuread/sdk/v6/go/azuread"
//	"github.com/pulumi/pulumi-time/sdk/go/time"
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
//			exampleRotating, err := time.NewRotating(ctx, "example", &time.RotatingArgs{
//				RotationDays: pulumi.Int(7),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuread.NewApplicationPassword(ctx, "example", &azuread.ApplicationPasswordArgs{
//				ApplicationId: example.ID(),
//				RotateWhenChanged: pulumi.StringMap{
//					"rotation": exampleRotating.ID(),
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
// This resource does not support importing.
type ApplicationPassword struct {
	pulumi.CustomResourceState

	// The resource ID of the application for which this password should be created. Changing this field forces a new resource to be created.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// A display name for the password. Changing this field forces a new resource to be created.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
	EndDate pulumi.StringOutput `pulumi:"endDate"`
	// A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
	//
	// Deprecated: The `endDateRelative` property is deprecated and will be removed in a future version of the AzureAD provider. Please instead use the Terraform `timeadd()` function to calculate a value for the `endDate` property.
	EndDateRelative pulumi.StringPtrOutput `pulumi:"endDateRelative"`
	// A UUID used to uniquely identify this password credential.
	KeyId pulumi.StringOutput `pulumi:"keyId"`
	// A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
	RotateWhenChanged pulumi.StringMapOutput `pulumi:"rotateWhenChanged"`
	// The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
	StartDate pulumi.StringOutput `pulumi:"startDate"`
	// The password for this application, which is generated by Azure Active Directory.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewApplicationPassword registers a new resource with the given unique name, arguments, and options.
func NewApplicationPassword(ctx *pulumi.Context,
	name string, args *ApplicationPasswordArgs, opts ...pulumi.ResourceOption) (*ApplicationPassword, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"value",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApplicationPassword
	err := ctx.RegisterResource("azuread:index/applicationPassword:ApplicationPassword", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationPassword gets an existing ApplicationPassword resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationPassword(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationPasswordState, opts ...pulumi.ResourceOption) (*ApplicationPassword, error) {
	var resource ApplicationPassword
	err := ctx.ReadResource("azuread:index/applicationPassword:ApplicationPassword", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationPassword resources.
type applicationPasswordState struct {
	// The resource ID of the application for which this password should be created. Changing this field forces a new resource to be created.
	ApplicationId *string `pulumi:"applicationId"`
	// A display name for the password. Changing this field forces a new resource to be created.
	DisplayName *string `pulumi:"displayName"`
	// The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
	EndDate *string `pulumi:"endDate"`
	// A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
	//
	// Deprecated: The `endDateRelative` property is deprecated and will be removed in a future version of the AzureAD provider. Please instead use the Terraform `timeadd()` function to calculate a value for the `endDate` property.
	EndDateRelative *string `pulumi:"endDateRelative"`
	// A UUID used to uniquely identify this password credential.
	KeyId *string `pulumi:"keyId"`
	// A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
	RotateWhenChanged map[string]string `pulumi:"rotateWhenChanged"`
	// The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
	StartDate *string `pulumi:"startDate"`
	// The password for this application, which is generated by Azure Active Directory.
	Value *string `pulumi:"value"`
}

type ApplicationPasswordState struct {
	// The resource ID of the application for which this password should be created. Changing this field forces a new resource to be created.
	ApplicationId pulumi.StringPtrInput
	// A display name for the password. Changing this field forces a new resource to be created.
	DisplayName pulumi.StringPtrInput
	// The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
	EndDate pulumi.StringPtrInput
	// A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
	//
	// Deprecated: The `endDateRelative` property is deprecated and will be removed in a future version of the AzureAD provider. Please instead use the Terraform `timeadd()` function to calculate a value for the `endDate` property.
	EndDateRelative pulumi.StringPtrInput
	// A UUID used to uniquely identify this password credential.
	KeyId pulumi.StringPtrInput
	// A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
	RotateWhenChanged pulumi.StringMapInput
	// The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
	StartDate pulumi.StringPtrInput
	// The password for this application, which is generated by Azure Active Directory.
	Value pulumi.StringPtrInput
}

func (ApplicationPasswordState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationPasswordState)(nil)).Elem()
}

type applicationPasswordArgs struct {
	// The resource ID of the application for which this password should be created. Changing this field forces a new resource to be created.
	ApplicationId string `pulumi:"applicationId"`
	// A display name for the password. Changing this field forces a new resource to be created.
	DisplayName *string `pulumi:"displayName"`
	// The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
	EndDate *string `pulumi:"endDate"`
	// A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
	//
	// Deprecated: The `endDateRelative` property is deprecated and will be removed in a future version of the AzureAD provider. Please instead use the Terraform `timeadd()` function to calculate a value for the `endDate` property.
	EndDateRelative *string `pulumi:"endDateRelative"`
	// A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
	RotateWhenChanged map[string]string `pulumi:"rotateWhenChanged"`
	// The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
	StartDate *string `pulumi:"startDate"`
}

// The set of arguments for constructing a ApplicationPassword resource.
type ApplicationPasswordArgs struct {
	// The resource ID of the application for which this password should be created. Changing this field forces a new resource to be created.
	ApplicationId pulumi.StringInput
	// A display name for the password. Changing this field forces a new resource to be created.
	DisplayName pulumi.StringPtrInput
	// The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
	EndDate pulumi.StringPtrInput
	// A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
	//
	// Deprecated: The `endDateRelative` property is deprecated and will be removed in a future version of the AzureAD provider. Please instead use the Terraform `timeadd()` function to calculate a value for the `endDate` property.
	EndDateRelative pulumi.StringPtrInput
	// A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
	RotateWhenChanged pulumi.StringMapInput
	// The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
	StartDate pulumi.StringPtrInput
}

func (ApplicationPasswordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationPasswordArgs)(nil)).Elem()
}

type ApplicationPasswordInput interface {
	pulumi.Input

	ToApplicationPasswordOutput() ApplicationPasswordOutput
	ToApplicationPasswordOutputWithContext(ctx context.Context) ApplicationPasswordOutput
}

func (*ApplicationPassword) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationPassword)(nil)).Elem()
}

func (i *ApplicationPassword) ToApplicationPasswordOutput() ApplicationPasswordOutput {
	return i.ToApplicationPasswordOutputWithContext(context.Background())
}

func (i *ApplicationPassword) ToApplicationPasswordOutputWithContext(ctx context.Context) ApplicationPasswordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPasswordOutput)
}

// ApplicationPasswordArrayInput is an input type that accepts ApplicationPasswordArray and ApplicationPasswordArrayOutput values.
// You can construct a concrete instance of `ApplicationPasswordArrayInput` via:
//
//	ApplicationPasswordArray{ ApplicationPasswordArgs{...} }
type ApplicationPasswordArrayInput interface {
	pulumi.Input

	ToApplicationPasswordArrayOutput() ApplicationPasswordArrayOutput
	ToApplicationPasswordArrayOutputWithContext(context.Context) ApplicationPasswordArrayOutput
}

type ApplicationPasswordArray []ApplicationPasswordInput

func (ApplicationPasswordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationPassword)(nil)).Elem()
}

func (i ApplicationPasswordArray) ToApplicationPasswordArrayOutput() ApplicationPasswordArrayOutput {
	return i.ToApplicationPasswordArrayOutputWithContext(context.Background())
}

func (i ApplicationPasswordArray) ToApplicationPasswordArrayOutputWithContext(ctx context.Context) ApplicationPasswordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPasswordArrayOutput)
}

// ApplicationPasswordMapInput is an input type that accepts ApplicationPasswordMap and ApplicationPasswordMapOutput values.
// You can construct a concrete instance of `ApplicationPasswordMapInput` via:
//
//	ApplicationPasswordMap{ "key": ApplicationPasswordArgs{...} }
type ApplicationPasswordMapInput interface {
	pulumi.Input

	ToApplicationPasswordMapOutput() ApplicationPasswordMapOutput
	ToApplicationPasswordMapOutputWithContext(context.Context) ApplicationPasswordMapOutput
}

type ApplicationPasswordMap map[string]ApplicationPasswordInput

func (ApplicationPasswordMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationPassword)(nil)).Elem()
}

func (i ApplicationPasswordMap) ToApplicationPasswordMapOutput() ApplicationPasswordMapOutput {
	return i.ToApplicationPasswordMapOutputWithContext(context.Background())
}

func (i ApplicationPasswordMap) ToApplicationPasswordMapOutputWithContext(ctx context.Context) ApplicationPasswordMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPasswordMapOutput)
}

type ApplicationPasswordOutput struct{ *pulumi.OutputState }

func (ApplicationPasswordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationPassword)(nil)).Elem()
}

func (o ApplicationPasswordOutput) ToApplicationPasswordOutput() ApplicationPasswordOutput {
	return o
}

func (o ApplicationPasswordOutput) ToApplicationPasswordOutputWithContext(ctx context.Context) ApplicationPasswordOutput {
	return o
}

// The resource ID of the application for which this password should be created. Changing this field forces a new resource to be created.
func (o ApplicationPasswordOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationPassword) pulumi.StringOutput { return v.ApplicationId }).(pulumi.StringOutput)
}

// A display name for the password. Changing this field forces a new resource to be created.
func (o ApplicationPasswordOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationPassword) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
func (o ApplicationPasswordOutput) EndDate() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationPassword) pulumi.StringOutput { return v.EndDate }).(pulumi.StringOutput)
}

// A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
//
// Deprecated: The `endDateRelative` property is deprecated and will be removed in a future version of the AzureAD provider. Please instead use the Terraform `timeadd()` function to calculate a value for the `endDate` property.
func (o ApplicationPasswordOutput) EndDateRelative() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationPassword) pulumi.StringPtrOutput { return v.EndDateRelative }).(pulumi.StringPtrOutput)
}

// A UUID used to uniquely identify this password credential.
func (o ApplicationPasswordOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationPassword) pulumi.StringOutput { return v.KeyId }).(pulumi.StringOutput)
}

// A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
func (o ApplicationPasswordOutput) RotateWhenChanged() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ApplicationPassword) pulumi.StringMapOutput { return v.RotateWhenChanged }).(pulumi.StringMapOutput)
}

// The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
func (o ApplicationPasswordOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationPassword) pulumi.StringOutput { return v.StartDate }).(pulumi.StringOutput)
}

// The password for this application, which is generated by Azure Active Directory.
func (o ApplicationPasswordOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationPassword) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

type ApplicationPasswordArrayOutput struct{ *pulumi.OutputState }

func (ApplicationPasswordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationPassword)(nil)).Elem()
}

func (o ApplicationPasswordArrayOutput) ToApplicationPasswordArrayOutput() ApplicationPasswordArrayOutput {
	return o
}

func (o ApplicationPasswordArrayOutput) ToApplicationPasswordArrayOutputWithContext(ctx context.Context) ApplicationPasswordArrayOutput {
	return o
}

func (o ApplicationPasswordArrayOutput) Index(i pulumi.IntInput) ApplicationPasswordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplicationPassword {
		return vs[0].([]*ApplicationPassword)[vs[1].(int)]
	}).(ApplicationPasswordOutput)
}

type ApplicationPasswordMapOutput struct{ *pulumi.OutputState }

func (ApplicationPasswordMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationPassword)(nil)).Elem()
}

func (o ApplicationPasswordMapOutput) ToApplicationPasswordMapOutput() ApplicationPasswordMapOutput {
	return o
}

func (o ApplicationPasswordMapOutput) ToApplicationPasswordMapOutputWithContext(ctx context.Context) ApplicationPasswordMapOutput {
	return o
}

func (o ApplicationPasswordMapOutput) MapIndex(k pulumi.StringInput) ApplicationPasswordOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplicationPassword {
		return vs[0].(map[string]*ApplicationPassword)[vs[1].(string)]
	}).(ApplicationPasswordOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationPasswordInput)(nil)).Elem(), &ApplicationPassword{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationPasswordArrayInput)(nil)).Elem(), ApplicationPasswordArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationPasswordMapInput)(nil)).Elem(), ApplicationPasswordMap{})
	pulumi.RegisterOutputType(ApplicationPasswordOutput{})
	pulumi.RegisterOutputType(ApplicationPasswordArrayOutput{})
	pulumi.RegisterOutputType(ApplicationPasswordMapOutput{})
}

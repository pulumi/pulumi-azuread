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
// *Basic example*
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
//			exampleApplication, err := azuread.NewApplication(ctx, "exampleApplication", &azuread.ApplicationArgs{
//				DisplayName: pulumi.String("example"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleServicePrincipal, err := azuread.NewServicePrincipal(ctx, "exampleServicePrincipal", &azuread.ServicePrincipalArgs{
//				ApplicationId: exampleApplication.ApplicationId,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuread.NewServicePrincipalPassword(ctx, "exampleServicePrincipalPassword", &azuread.ServicePrincipalPasswordArgs{
//				ServicePrincipalId: exampleServicePrincipal.ObjectId,
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
//	"github.com/pulumi/pulumi-azuread/sdk/v5/go/azuread"
//	"github.com/pulumi/pulumi-time/sdk/go/time"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleApplication, err := azuread.NewApplication(ctx, "exampleApplication", &azuread.ApplicationArgs{
//				DisplayName: pulumi.String("example"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleServicePrincipal, err := azuread.NewServicePrincipal(ctx, "exampleServicePrincipal", &azuread.ServicePrincipalArgs{
//				ApplicationId: exampleApplication.ApplicationId,
//			})
//			if err != nil {
//				return err
//			}
//			exampleRotating, err := time.NewRotating(ctx, "exampleRotating", &time.RotatingArgs{
//				RotationDays: pulumi.Int(7),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuread.NewServicePrincipalPassword(ctx, "exampleServicePrincipalPassword", &azuread.ServicePrincipalPasswordArgs{
//				ServicePrincipalId: exampleServicePrincipal.ObjectId,
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
type ServicePrincipalPassword struct {
	pulumi.CustomResourceState

	// A display name for the password.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
	EndDate pulumi.StringOutput `pulumi:"endDate"`
	// A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
	EndDateRelative pulumi.StringPtrOutput `pulumi:"endDateRelative"`
	// A UUID used to uniquely identify this password credential.
	KeyId pulumi.StringOutput `pulumi:"keyId"`
	// A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
	RotateWhenChanged pulumi.StringMapOutput `pulumi:"rotateWhenChanged"`
	// The object ID of the service principal for which this password should be created. Changing this field forces a new resource to be created.
	ServicePrincipalId pulumi.StringOutput `pulumi:"servicePrincipalId"`
	// The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
	StartDate pulumi.StringOutput `pulumi:"startDate"`
	// The password for this service principal, which is generated by Azure Active Directory.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewServicePrincipalPassword registers a new resource with the given unique name, arguments, and options.
func NewServicePrincipalPassword(ctx *pulumi.Context,
	name string, args *ServicePrincipalPasswordArgs, opts ...pulumi.ResourceOption) (*ServicePrincipalPassword, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServicePrincipalId == nil {
		return nil, errors.New("invalid value for required argument 'ServicePrincipalId'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"value",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServicePrincipalPassword
	err := ctx.RegisterResource("azuread:index/servicePrincipalPassword:ServicePrincipalPassword", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServicePrincipalPassword gets an existing ServicePrincipalPassword resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServicePrincipalPassword(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServicePrincipalPasswordState, opts ...pulumi.ResourceOption) (*ServicePrincipalPassword, error) {
	var resource ServicePrincipalPassword
	err := ctx.ReadResource("azuread:index/servicePrincipalPassword:ServicePrincipalPassword", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServicePrincipalPassword resources.
type servicePrincipalPasswordState struct {
	// A display name for the password.
	DisplayName *string `pulumi:"displayName"`
	// The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
	EndDate *string `pulumi:"endDate"`
	// A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
	EndDateRelative *string `pulumi:"endDateRelative"`
	// A UUID used to uniquely identify this password credential.
	KeyId *string `pulumi:"keyId"`
	// A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
	RotateWhenChanged map[string]string `pulumi:"rotateWhenChanged"`
	// The object ID of the service principal for which this password should be created. Changing this field forces a new resource to be created.
	ServicePrincipalId *string `pulumi:"servicePrincipalId"`
	// The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
	StartDate *string `pulumi:"startDate"`
	// The password for this service principal, which is generated by Azure Active Directory.
	Value *string `pulumi:"value"`
}

type ServicePrincipalPasswordState struct {
	// A display name for the password.
	DisplayName pulumi.StringPtrInput
	// The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
	EndDate pulumi.StringPtrInput
	// A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
	EndDateRelative pulumi.StringPtrInput
	// A UUID used to uniquely identify this password credential.
	KeyId pulumi.StringPtrInput
	// A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
	RotateWhenChanged pulumi.StringMapInput
	// The object ID of the service principal for which this password should be created. Changing this field forces a new resource to be created.
	ServicePrincipalId pulumi.StringPtrInput
	// The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
	StartDate pulumi.StringPtrInput
	// The password for this service principal, which is generated by Azure Active Directory.
	Value pulumi.StringPtrInput
}

func (ServicePrincipalPasswordState) ElementType() reflect.Type {
	return reflect.TypeOf((*servicePrincipalPasswordState)(nil)).Elem()
}

type servicePrincipalPasswordArgs struct {
	// A display name for the password.
	DisplayName *string `pulumi:"displayName"`
	// The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
	EndDate *string `pulumi:"endDate"`
	// A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
	EndDateRelative *string `pulumi:"endDateRelative"`
	// A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
	RotateWhenChanged map[string]string `pulumi:"rotateWhenChanged"`
	// The object ID of the service principal for which this password should be created. Changing this field forces a new resource to be created.
	ServicePrincipalId string `pulumi:"servicePrincipalId"`
	// The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
	StartDate *string `pulumi:"startDate"`
}

// The set of arguments for constructing a ServicePrincipalPassword resource.
type ServicePrincipalPasswordArgs struct {
	// A display name for the password.
	DisplayName pulumi.StringPtrInput
	// The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
	EndDate pulumi.StringPtrInput
	// A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
	EndDateRelative pulumi.StringPtrInput
	// A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
	RotateWhenChanged pulumi.StringMapInput
	// The object ID of the service principal for which this password should be created. Changing this field forces a new resource to be created.
	ServicePrincipalId pulumi.StringInput
	// The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
	StartDate pulumi.StringPtrInput
}

func (ServicePrincipalPasswordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*servicePrincipalPasswordArgs)(nil)).Elem()
}

type ServicePrincipalPasswordInput interface {
	pulumi.Input

	ToServicePrincipalPasswordOutput() ServicePrincipalPasswordOutput
	ToServicePrincipalPasswordOutputWithContext(ctx context.Context) ServicePrincipalPasswordOutput
}

func (*ServicePrincipalPassword) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePrincipalPassword)(nil)).Elem()
}

func (i *ServicePrincipalPassword) ToServicePrincipalPasswordOutput() ServicePrincipalPasswordOutput {
	return i.ToServicePrincipalPasswordOutputWithContext(context.Background())
}

func (i *ServicePrincipalPassword) ToServicePrincipalPasswordOutputWithContext(ctx context.Context) ServicePrincipalPasswordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalPasswordOutput)
}

// ServicePrincipalPasswordArrayInput is an input type that accepts ServicePrincipalPasswordArray and ServicePrincipalPasswordArrayOutput values.
// You can construct a concrete instance of `ServicePrincipalPasswordArrayInput` via:
//
//	ServicePrincipalPasswordArray{ ServicePrincipalPasswordArgs{...} }
type ServicePrincipalPasswordArrayInput interface {
	pulumi.Input

	ToServicePrincipalPasswordArrayOutput() ServicePrincipalPasswordArrayOutput
	ToServicePrincipalPasswordArrayOutputWithContext(context.Context) ServicePrincipalPasswordArrayOutput
}

type ServicePrincipalPasswordArray []ServicePrincipalPasswordInput

func (ServicePrincipalPasswordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServicePrincipalPassword)(nil)).Elem()
}

func (i ServicePrincipalPasswordArray) ToServicePrincipalPasswordArrayOutput() ServicePrincipalPasswordArrayOutput {
	return i.ToServicePrincipalPasswordArrayOutputWithContext(context.Background())
}

func (i ServicePrincipalPasswordArray) ToServicePrincipalPasswordArrayOutputWithContext(ctx context.Context) ServicePrincipalPasswordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalPasswordArrayOutput)
}

// ServicePrincipalPasswordMapInput is an input type that accepts ServicePrincipalPasswordMap and ServicePrincipalPasswordMapOutput values.
// You can construct a concrete instance of `ServicePrincipalPasswordMapInput` via:
//
//	ServicePrincipalPasswordMap{ "key": ServicePrincipalPasswordArgs{...} }
type ServicePrincipalPasswordMapInput interface {
	pulumi.Input

	ToServicePrincipalPasswordMapOutput() ServicePrincipalPasswordMapOutput
	ToServicePrincipalPasswordMapOutputWithContext(context.Context) ServicePrincipalPasswordMapOutput
}

type ServicePrincipalPasswordMap map[string]ServicePrincipalPasswordInput

func (ServicePrincipalPasswordMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServicePrincipalPassword)(nil)).Elem()
}

func (i ServicePrincipalPasswordMap) ToServicePrincipalPasswordMapOutput() ServicePrincipalPasswordMapOutput {
	return i.ToServicePrincipalPasswordMapOutputWithContext(context.Background())
}

func (i ServicePrincipalPasswordMap) ToServicePrincipalPasswordMapOutputWithContext(ctx context.Context) ServicePrincipalPasswordMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalPasswordMapOutput)
}

type ServicePrincipalPasswordOutput struct{ *pulumi.OutputState }

func (ServicePrincipalPasswordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePrincipalPassword)(nil)).Elem()
}

func (o ServicePrincipalPasswordOutput) ToServicePrincipalPasswordOutput() ServicePrincipalPasswordOutput {
	return o
}

func (o ServicePrincipalPasswordOutput) ToServicePrincipalPasswordOutputWithContext(ctx context.Context) ServicePrincipalPasswordOutput {
	return o
}

// A display name for the password.
func (o ServicePrincipalPasswordOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServicePrincipalPassword) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
func (o ServicePrincipalPasswordOutput) EndDate() pulumi.StringOutput {
	return o.ApplyT(func(v *ServicePrincipalPassword) pulumi.StringOutput { return v.EndDate }).(pulumi.StringOutput)
}

// A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
func (o ServicePrincipalPasswordOutput) EndDateRelative() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalPassword) pulumi.StringPtrOutput { return v.EndDateRelative }).(pulumi.StringPtrOutput)
}

// A UUID used to uniquely identify this password credential.
func (o ServicePrincipalPasswordOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServicePrincipalPassword) pulumi.StringOutput { return v.KeyId }).(pulumi.StringOutput)
}

// A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
func (o ServicePrincipalPasswordOutput) RotateWhenChanged() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServicePrincipalPassword) pulumi.StringMapOutput { return v.RotateWhenChanged }).(pulumi.StringMapOutput)
}

// The object ID of the service principal for which this password should be created. Changing this field forces a new resource to be created.
func (o ServicePrincipalPasswordOutput) ServicePrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServicePrincipalPassword) pulumi.StringOutput { return v.ServicePrincipalId }).(pulumi.StringOutput)
}

// The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
func (o ServicePrincipalPasswordOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v *ServicePrincipalPassword) pulumi.StringOutput { return v.StartDate }).(pulumi.StringOutput)
}

// The password for this service principal, which is generated by Azure Active Directory.
func (o ServicePrincipalPasswordOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *ServicePrincipalPassword) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

type ServicePrincipalPasswordArrayOutput struct{ *pulumi.OutputState }

func (ServicePrincipalPasswordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServicePrincipalPassword)(nil)).Elem()
}

func (o ServicePrincipalPasswordArrayOutput) ToServicePrincipalPasswordArrayOutput() ServicePrincipalPasswordArrayOutput {
	return o
}

func (o ServicePrincipalPasswordArrayOutput) ToServicePrincipalPasswordArrayOutputWithContext(ctx context.Context) ServicePrincipalPasswordArrayOutput {
	return o
}

func (o ServicePrincipalPasswordArrayOutput) Index(i pulumi.IntInput) ServicePrincipalPasswordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServicePrincipalPassword {
		return vs[0].([]*ServicePrincipalPassword)[vs[1].(int)]
	}).(ServicePrincipalPasswordOutput)
}

type ServicePrincipalPasswordMapOutput struct{ *pulumi.OutputState }

func (ServicePrincipalPasswordMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServicePrincipalPassword)(nil)).Elem()
}

func (o ServicePrincipalPasswordMapOutput) ToServicePrincipalPasswordMapOutput() ServicePrincipalPasswordMapOutput {
	return o
}

func (o ServicePrincipalPasswordMapOutput) ToServicePrincipalPasswordMapOutputWithContext(ctx context.Context) ServicePrincipalPasswordMapOutput {
	return o
}

func (o ServicePrincipalPasswordMapOutput) MapIndex(k pulumi.StringInput) ServicePrincipalPasswordOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServicePrincipalPassword {
		return vs[0].(map[string]*ServicePrincipalPassword)[vs[1].(string)]
	}).(ServicePrincipalPasswordOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePrincipalPasswordInput)(nil)).Elem(), &ServicePrincipalPassword{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePrincipalPasswordArrayInput)(nil)).Elem(), ServicePrincipalPasswordArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePrincipalPasswordMapInput)(nil)).Elem(), ServicePrincipalPasswordMap{})
	pulumi.RegisterOutputType(ServicePrincipalPasswordOutput{})
	pulumi.RegisterOutputType(ServicePrincipalPasswordArrayOutput{})
	pulumi.RegisterOutputType(ServicePrincipalPasswordMapOutput{})
}

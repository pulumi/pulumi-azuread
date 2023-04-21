// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a token signing certificate associated with a service principal within Azure Active Directory.
//
// ## API Permissions
//
// The following API permissions are required in order to use this resource.
//
// When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.All` or `Directory.ReadWrite.All`
//
// When authenticated with a user principal, this resource requires one of the following directory roles: `Application Administrator` or `Global Administrator`
//
// ## Example Usage
//
// *Using default settings*
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
//			_, err = azuread.NewServicePrincipalTokenSigningCertificate(ctx, "exampleServicePrincipalTokenSigningCertificate", &azuread.ServicePrincipalTokenSigningCertificateArgs{
//				ServicePrincipalId: exampleServicePrincipal.ID(),
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
// *Using custom settings*
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
//			_, err = azuread.NewServicePrincipalTokenSigningCertificate(ctx, "exampleServicePrincipalTokenSigningCertificate", &azuread.ServicePrincipalTokenSigningCertificateArgs{
//				ServicePrincipalId: exampleServicePrincipal.ID(),
//				DisplayName:        pulumi.String("CN=example.com"),
//				EndDate:            pulumi.String("2023-05-01T01:02:03Z"),
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
// Token signing certificates can be imported using the object ID of the associated service principal and the key ID of the verify certificate credential, e.g.
//
// ```sh
//
//	$ pulumi import azuread:index/servicePrincipalTokenSigningCertificate:ServicePrincipalTokenSigningCertificate example 00000000-0000-0000-0000-000000000000/tokenSigningCertificate/11111111-1111-1111-1111-111111111111
//
// ```
//
//	-> This ID format is unique to Terraform and is composed of the service principal's object ID, the string "tokenSigningCertificate" and the verify certificate's key ID in the format `{ServicePrincipalObjectId}/tokenSigningCertificate/{CertificateKeyId}`.
type ServicePrincipalTokenSigningCertificate struct {
	pulumi.CustomResourceState

	// Specifies a friendly name for the certificate.
	// Must start with `CN=`. Changing this field forces a new resource to be created.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The end date until which the token signing certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
	EndDate pulumi.StringOutput `pulumi:"endDate"`
	// A UUID used to uniquely identify the verify certificate.
	KeyId pulumi.StringOutput `pulumi:"keyId"`
	// The object ID of the service principal for which this certificate should be created. Changing this field forces a new resource to be created.
	ServicePrincipalId pulumi.StringOutput `pulumi:"servicePrincipalId"`
	// The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
	StartDate pulumi.StringOutput `pulumi:"startDate"`
	// A SHA-1 generated thumbprint of the token signing certificate, which can be used to set the preferred signing certificate for a service principal.
	Thumbprint pulumi.StringOutput `pulumi:"thumbprint"`
	// The certificate data, which is PEM encoded but does not include the
	// header `-----BEGIN CERTIFICATE-----\n` or the footer `\n-----END CERTIFICATE-----`.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewServicePrincipalTokenSigningCertificate registers a new resource with the given unique name, arguments, and options.
func NewServicePrincipalTokenSigningCertificate(ctx *pulumi.Context,
	name string, args *ServicePrincipalTokenSigningCertificateArgs, opts ...pulumi.ResourceOption) (*ServicePrincipalTokenSigningCertificate, error) {
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
	var resource ServicePrincipalTokenSigningCertificate
	err := ctx.RegisterResource("azuread:index/servicePrincipalTokenSigningCertificate:ServicePrincipalTokenSigningCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServicePrincipalTokenSigningCertificate gets an existing ServicePrincipalTokenSigningCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServicePrincipalTokenSigningCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServicePrincipalTokenSigningCertificateState, opts ...pulumi.ResourceOption) (*ServicePrincipalTokenSigningCertificate, error) {
	var resource ServicePrincipalTokenSigningCertificate
	err := ctx.ReadResource("azuread:index/servicePrincipalTokenSigningCertificate:ServicePrincipalTokenSigningCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServicePrincipalTokenSigningCertificate resources.
type servicePrincipalTokenSigningCertificateState struct {
	// Specifies a friendly name for the certificate.
	// Must start with `CN=`. Changing this field forces a new resource to be created.
	DisplayName *string `pulumi:"displayName"`
	// The end date until which the token signing certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
	EndDate *string `pulumi:"endDate"`
	// A UUID used to uniquely identify the verify certificate.
	KeyId *string `pulumi:"keyId"`
	// The object ID of the service principal for which this certificate should be created. Changing this field forces a new resource to be created.
	ServicePrincipalId *string `pulumi:"servicePrincipalId"`
	// The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
	StartDate *string `pulumi:"startDate"`
	// A SHA-1 generated thumbprint of the token signing certificate, which can be used to set the preferred signing certificate for a service principal.
	Thumbprint *string `pulumi:"thumbprint"`
	// The certificate data, which is PEM encoded but does not include the
	// header `-----BEGIN CERTIFICATE-----\n` or the footer `\n-----END CERTIFICATE-----`.
	Value *string `pulumi:"value"`
}

type ServicePrincipalTokenSigningCertificateState struct {
	// Specifies a friendly name for the certificate.
	// Must start with `CN=`. Changing this field forces a new resource to be created.
	DisplayName pulumi.StringPtrInput
	// The end date until which the token signing certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
	EndDate pulumi.StringPtrInput
	// A UUID used to uniquely identify the verify certificate.
	KeyId pulumi.StringPtrInput
	// The object ID of the service principal for which this certificate should be created. Changing this field forces a new resource to be created.
	ServicePrincipalId pulumi.StringPtrInput
	// The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
	StartDate pulumi.StringPtrInput
	// A SHA-1 generated thumbprint of the token signing certificate, which can be used to set the preferred signing certificate for a service principal.
	Thumbprint pulumi.StringPtrInput
	// The certificate data, which is PEM encoded but does not include the
	// header `-----BEGIN CERTIFICATE-----\n` or the footer `\n-----END CERTIFICATE-----`.
	Value pulumi.StringPtrInput
}

func (ServicePrincipalTokenSigningCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*servicePrincipalTokenSigningCertificateState)(nil)).Elem()
}

type servicePrincipalTokenSigningCertificateArgs struct {
	// Specifies a friendly name for the certificate.
	// Must start with `CN=`. Changing this field forces a new resource to be created.
	DisplayName *string `pulumi:"displayName"`
	// The end date until which the token signing certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
	EndDate *string `pulumi:"endDate"`
	// The object ID of the service principal for which this certificate should be created. Changing this field forces a new resource to be created.
	ServicePrincipalId string `pulumi:"servicePrincipalId"`
}

// The set of arguments for constructing a ServicePrincipalTokenSigningCertificate resource.
type ServicePrincipalTokenSigningCertificateArgs struct {
	// Specifies a friendly name for the certificate.
	// Must start with `CN=`. Changing this field forces a new resource to be created.
	DisplayName pulumi.StringPtrInput
	// The end date until which the token signing certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
	EndDate pulumi.StringPtrInput
	// The object ID of the service principal for which this certificate should be created. Changing this field forces a new resource to be created.
	ServicePrincipalId pulumi.StringInput
}

func (ServicePrincipalTokenSigningCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*servicePrincipalTokenSigningCertificateArgs)(nil)).Elem()
}

type ServicePrincipalTokenSigningCertificateInput interface {
	pulumi.Input

	ToServicePrincipalTokenSigningCertificateOutput() ServicePrincipalTokenSigningCertificateOutput
	ToServicePrincipalTokenSigningCertificateOutputWithContext(ctx context.Context) ServicePrincipalTokenSigningCertificateOutput
}

func (*ServicePrincipalTokenSigningCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePrincipalTokenSigningCertificate)(nil)).Elem()
}

func (i *ServicePrincipalTokenSigningCertificate) ToServicePrincipalTokenSigningCertificateOutput() ServicePrincipalTokenSigningCertificateOutput {
	return i.ToServicePrincipalTokenSigningCertificateOutputWithContext(context.Background())
}

func (i *ServicePrincipalTokenSigningCertificate) ToServicePrincipalTokenSigningCertificateOutputWithContext(ctx context.Context) ServicePrincipalTokenSigningCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalTokenSigningCertificateOutput)
}

// ServicePrincipalTokenSigningCertificateArrayInput is an input type that accepts ServicePrincipalTokenSigningCertificateArray and ServicePrincipalTokenSigningCertificateArrayOutput values.
// You can construct a concrete instance of `ServicePrincipalTokenSigningCertificateArrayInput` via:
//
//	ServicePrincipalTokenSigningCertificateArray{ ServicePrincipalTokenSigningCertificateArgs{...} }
type ServicePrincipalTokenSigningCertificateArrayInput interface {
	pulumi.Input

	ToServicePrincipalTokenSigningCertificateArrayOutput() ServicePrincipalTokenSigningCertificateArrayOutput
	ToServicePrincipalTokenSigningCertificateArrayOutputWithContext(context.Context) ServicePrincipalTokenSigningCertificateArrayOutput
}

type ServicePrincipalTokenSigningCertificateArray []ServicePrincipalTokenSigningCertificateInput

func (ServicePrincipalTokenSigningCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServicePrincipalTokenSigningCertificate)(nil)).Elem()
}

func (i ServicePrincipalTokenSigningCertificateArray) ToServicePrincipalTokenSigningCertificateArrayOutput() ServicePrincipalTokenSigningCertificateArrayOutput {
	return i.ToServicePrincipalTokenSigningCertificateArrayOutputWithContext(context.Background())
}

func (i ServicePrincipalTokenSigningCertificateArray) ToServicePrincipalTokenSigningCertificateArrayOutputWithContext(ctx context.Context) ServicePrincipalTokenSigningCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalTokenSigningCertificateArrayOutput)
}

// ServicePrincipalTokenSigningCertificateMapInput is an input type that accepts ServicePrincipalTokenSigningCertificateMap and ServicePrincipalTokenSigningCertificateMapOutput values.
// You can construct a concrete instance of `ServicePrincipalTokenSigningCertificateMapInput` via:
//
//	ServicePrincipalTokenSigningCertificateMap{ "key": ServicePrincipalTokenSigningCertificateArgs{...} }
type ServicePrincipalTokenSigningCertificateMapInput interface {
	pulumi.Input

	ToServicePrincipalTokenSigningCertificateMapOutput() ServicePrincipalTokenSigningCertificateMapOutput
	ToServicePrincipalTokenSigningCertificateMapOutputWithContext(context.Context) ServicePrincipalTokenSigningCertificateMapOutput
}

type ServicePrincipalTokenSigningCertificateMap map[string]ServicePrincipalTokenSigningCertificateInput

func (ServicePrincipalTokenSigningCertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServicePrincipalTokenSigningCertificate)(nil)).Elem()
}

func (i ServicePrincipalTokenSigningCertificateMap) ToServicePrincipalTokenSigningCertificateMapOutput() ServicePrincipalTokenSigningCertificateMapOutput {
	return i.ToServicePrincipalTokenSigningCertificateMapOutputWithContext(context.Background())
}

func (i ServicePrincipalTokenSigningCertificateMap) ToServicePrincipalTokenSigningCertificateMapOutputWithContext(ctx context.Context) ServicePrincipalTokenSigningCertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalTokenSigningCertificateMapOutput)
}

type ServicePrincipalTokenSigningCertificateOutput struct{ *pulumi.OutputState }

func (ServicePrincipalTokenSigningCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePrincipalTokenSigningCertificate)(nil)).Elem()
}

func (o ServicePrincipalTokenSigningCertificateOutput) ToServicePrincipalTokenSigningCertificateOutput() ServicePrincipalTokenSigningCertificateOutput {
	return o
}

func (o ServicePrincipalTokenSigningCertificateOutput) ToServicePrincipalTokenSigningCertificateOutputWithContext(ctx context.Context) ServicePrincipalTokenSigningCertificateOutput {
	return o
}

// Specifies a friendly name for the certificate.
// Must start with `CN=`. Changing this field forces a new resource to be created.
func (o ServicePrincipalTokenSigningCertificateOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServicePrincipalTokenSigningCertificate) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The end date until which the token signing certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
func (o ServicePrincipalTokenSigningCertificateOutput) EndDate() pulumi.StringOutput {
	return o.ApplyT(func(v *ServicePrincipalTokenSigningCertificate) pulumi.StringOutput { return v.EndDate }).(pulumi.StringOutput)
}

// A UUID used to uniquely identify the verify certificate.
func (o ServicePrincipalTokenSigningCertificateOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServicePrincipalTokenSigningCertificate) pulumi.StringOutput { return v.KeyId }).(pulumi.StringOutput)
}

// The object ID of the service principal for which this certificate should be created. Changing this field forces a new resource to be created.
func (o ServicePrincipalTokenSigningCertificateOutput) ServicePrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServicePrincipalTokenSigningCertificate) pulumi.StringOutput { return v.ServicePrincipalId }).(pulumi.StringOutput)
}

// The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
func (o ServicePrincipalTokenSigningCertificateOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v *ServicePrincipalTokenSigningCertificate) pulumi.StringOutput { return v.StartDate }).(pulumi.StringOutput)
}

// A SHA-1 generated thumbprint of the token signing certificate, which can be used to set the preferred signing certificate for a service principal.
func (o ServicePrincipalTokenSigningCertificateOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v *ServicePrincipalTokenSigningCertificate) pulumi.StringOutput { return v.Thumbprint }).(pulumi.StringOutput)
}

// The certificate data, which is PEM encoded but does not include the
// header `-----BEGIN CERTIFICATE-----\n` or the footer `\n-----END CERTIFICATE-----`.
func (o ServicePrincipalTokenSigningCertificateOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *ServicePrincipalTokenSigningCertificate) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

type ServicePrincipalTokenSigningCertificateArrayOutput struct{ *pulumi.OutputState }

func (ServicePrincipalTokenSigningCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServicePrincipalTokenSigningCertificate)(nil)).Elem()
}

func (o ServicePrincipalTokenSigningCertificateArrayOutput) ToServicePrincipalTokenSigningCertificateArrayOutput() ServicePrincipalTokenSigningCertificateArrayOutput {
	return o
}

func (o ServicePrincipalTokenSigningCertificateArrayOutput) ToServicePrincipalTokenSigningCertificateArrayOutputWithContext(ctx context.Context) ServicePrincipalTokenSigningCertificateArrayOutput {
	return o
}

func (o ServicePrincipalTokenSigningCertificateArrayOutput) Index(i pulumi.IntInput) ServicePrincipalTokenSigningCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServicePrincipalTokenSigningCertificate {
		return vs[0].([]*ServicePrincipalTokenSigningCertificate)[vs[1].(int)]
	}).(ServicePrincipalTokenSigningCertificateOutput)
}

type ServicePrincipalTokenSigningCertificateMapOutput struct{ *pulumi.OutputState }

func (ServicePrincipalTokenSigningCertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServicePrincipalTokenSigningCertificate)(nil)).Elem()
}

func (o ServicePrincipalTokenSigningCertificateMapOutput) ToServicePrincipalTokenSigningCertificateMapOutput() ServicePrincipalTokenSigningCertificateMapOutput {
	return o
}

func (o ServicePrincipalTokenSigningCertificateMapOutput) ToServicePrincipalTokenSigningCertificateMapOutputWithContext(ctx context.Context) ServicePrincipalTokenSigningCertificateMapOutput {
	return o
}

func (o ServicePrincipalTokenSigningCertificateMapOutput) MapIndex(k pulumi.StringInput) ServicePrincipalTokenSigningCertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServicePrincipalTokenSigningCertificate {
		return vs[0].(map[string]*ServicePrincipalTokenSigningCertificate)[vs[1].(string)]
	}).(ServicePrincipalTokenSigningCertificateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePrincipalTokenSigningCertificateInput)(nil)).Elem(), &ServicePrincipalTokenSigningCertificate{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePrincipalTokenSigningCertificateArrayInput)(nil)).Elem(), ServicePrincipalTokenSigningCertificateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePrincipalTokenSigningCertificateMapInput)(nil)).Elem(), ServicePrincipalTokenSigningCertificateMap{})
	pulumi.RegisterOutputType(ServicePrincipalTokenSigningCertificateOutput{})
	pulumi.RegisterOutputType(ServicePrincipalTokenSigningCertificateArrayOutput{})
	pulumi.RegisterOutputType(ServicePrincipalTokenSigningCertificateMapOutput{})
}
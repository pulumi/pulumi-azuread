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
// *Using a PEM certificate*
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-azuread/sdk/v6/go/azuread"
//	"github.com/pulumi/pulumi-std/sdk/go/std"
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
//			exampleServicePrincipal, err := azuread.NewServicePrincipal(ctx, "example", &azuread.ServicePrincipalArgs{
//				ClientId: example.ClientId,
//			})
//			if err != nil {
//				return err
//			}
//			invokeFile, err := std.File(ctx, &std.FileArgs{
//				Input: "cert.pem",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = azuread.NewServicePrincipalCertificate(ctx, "example", &azuread.ServicePrincipalCertificateArgs{
//				ServicePrincipalId: exampleServicePrincipal.ID(),
//				Type:               pulumi.String("AsymmetricX509Cert"),
//				Value:              pulumi.String(invokeFile.Result),
//				EndDate:            pulumi.String("2021-05-01T01:02:03Z"),
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
// *Using a DER certificate*
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-azuread/sdk/v6/go/azuread"
//	"github.com/pulumi/pulumi-std/sdk/go/std"
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
//			exampleServicePrincipal, err := azuread.NewServicePrincipal(ctx, "example", &azuread.ServicePrincipalArgs{
//				ClientId: example.ClientId,
//			})
//			if err != nil {
//				return err
//			}
//			invokeBase64encode, err := std.Base64encode(ctx, &std.Base64encodeArgs{
//				Input: std.File(ctx, &std.FileArgs{
//					Input: "cert.der",
//				}, nil).Result,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = azuread.NewServicePrincipalCertificate(ctx, "example", &azuread.ServicePrincipalCertificateArgs{
//				ServicePrincipalId: exampleServicePrincipal.ID(),
//				Type:               pulumi.String("AsymmetricX509Cert"),
//				Encoding:           pulumi.String("base64"),
//				Value:              pulumi.String(invokeBase64encode.Result),
//				EndDate:            pulumi.String("2021-05-01T01:02:03Z"),
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
// Certificates can be imported using the object ID of the associated service principal and the key ID of the certificate credential, e.g.
//
// ```sh
// $ pulumi import azuread:index/servicePrincipalCertificate:ServicePrincipalCertificate example 00000000-0000-0000-0000-000000000000/certificate/11111111-1111-1111-1111-111111111111
// ```
//
// -> This ID format is unique to Terraform and is composed of the service principal's object ID, the string "certificate" and the certificate's key ID in the format `{ServicePrincipalObjectId}/certificate/{CertificateKeyId}`.
type ServicePrincipalCertificate struct {
	pulumi.CustomResourceState

	// Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.
	//
	// > **Tip for Azure Key Vault** The `hex` encoding option is useful for consuming certificate data from the azurermKeyVaultCertificate resource.
	Encoding pulumi.StringPtrOutput `pulumi:"encoding"`
	// The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
	EndDate pulumi.StringOutput `pulumi:"endDate"`
	// A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". Changing this field forces a new resource to be created.
	//
	// > One of `endDate` or `endDateRelative` must be set. The maximum duration is determined by Azure AD.
	//
	// Deprecated: The `endDateRelative` property is deprecated and will be removed in a future version of the AzureAD provider. Please instead use the Terraform `timeadd()` function to calculate a value for the `endDate` property.
	EndDateRelative pulumi.StringPtrOutput `pulumi:"endDateRelative"`
	// A UUID used to uniquely identify this certificate. If not specified a UUID will be automatically generated. Changing this field forces a new resource to be created.
	KeyId pulumi.StringOutput `pulumi:"keyId"`
	// The ID of the service principal for which this certificate should be created. Changing this field forces a new resource to be created.
	ServicePrincipalId pulumi.StringOutput `pulumi:"servicePrincipalId"`
	// The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the value is determined by Azure Active Directory and is usually the start date of the certificate for asymmetric keys, or the current timestamp for symmetric keys. Changing this field forces a new resource to be created.
	StartDate pulumi.StringOutput `pulumi:"startDate"`
	// The type of key/certificate. Must be one of `AsymmetricX509Cert` or `Symmetric`. Changing this fields forces a new resource to be created.
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding` argument.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewServicePrincipalCertificate registers a new resource with the given unique name, arguments, and options.
func NewServicePrincipalCertificate(ctx *pulumi.Context,
	name string, args *ServicePrincipalCertificateArgs, opts ...pulumi.ResourceOption) (*ServicePrincipalCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServicePrincipalId == nil {
		return nil, errors.New("invalid value for required argument 'ServicePrincipalId'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	if args.Value != nil {
		args.Value = pulumi.ToSecret(args.Value).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"value",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServicePrincipalCertificate
	err := ctx.RegisterResource("azuread:index/servicePrincipalCertificate:ServicePrincipalCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServicePrincipalCertificate gets an existing ServicePrincipalCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServicePrincipalCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServicePrincipalCertificateState, opts ...pulumi.ResourceOption) (*ServicePrincipalCertificate, error) {
	var resource ServicePrincipalCertificate
	err := ctx.ReadResource("azuread:index/servicePrincipalCertificate:ServicePrincipalCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServicePrincipalCertificate resources.
type servicePrincipalCertificateState struct {
	// Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.
	//
	// > **Tip for Azure Key Vault** The `hex` encoding option is useful for consuming certificate data from the azurermKeyVaultCertificate resource.
	Encoding *string `pulumi:"encoding"`
	// The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
	EndDate *string `pulumi:"endDate"`
	// A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". Changing this field forces a new resource to be created.
	//
	// > One of `endDate` or `endDateRelative` must be set. The maximum duration is determined by Azure AD.
	//
	// Deprecated: The `endDateRelative` property is deprecated and will be removed in a future version of the AzureAD provider. Please instead use the Terraform `timeadd()` function to calculate a value for the `endDate` property.
	EndDateRelative *string `pulumi:"endDateRelative"`
	// A UUID used to uniquely identify this certificate. If not specified a UUID will be automatically generated. Changing this field forces a new resource to be created.
	KeyId *string `pulumi:"keyId"`
	// The ID of the service principal for which this certificate should be created. Changing this field forces a new resource to be created.
	ServicePrincipalId *string `pulumi:"servicePrincipalId"`
	// The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the value is determined by Azure Active Directory and is usually the start date of the certificate for asymmetric keys, or the current timestamp for symmetric keys. Changing this field forces a new resource to be created.
	StartDate *string `pulumi:"startDate"`
	// The type of key/certificate. Must be one of `AsymmetricX509Cert` or `Symmetric`. Changing this fields forces a new resource to be created.
	Type *string `pulumi:"type"`
	// The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding` argument.
	Value *string `pulumi:"value"`
}

type ServicePrincipalCertificateState struct {
	// Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.
	//
	// > **Tip for Azure Key Vault** The `hex` encoding option is useful for consuming certificate data from the azurermKeyVaultCertificate resource.
	Encoding pulumi.StringPtrInput
	// The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
	EndDate pulumi.StringPtrInput
	// A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". Changing this field forces a new resource to be created.
	//
	// > One of `endDate` or `endDateRelative` must be set. The maximum duration is determined by Azure AD.
	//
	// Deprecated: The `endDateRelative` property is deprecated and will be removed in a future version of the AzureAD provider. Please instead use the Terraform `timeadd()` function to calculate a value for the `endDate` property.
	EndDateRelative pulumi.StringPtrInput
	// A UUID used to uniquely identify this certificate. If not specified a UUID will be automatically generated. Changing this field forces a new resource to be created.
	KeyId pulumi.StringPtrInput
	// The ID of the service principal for which this certificate should be created. Changing this field forces a new resource to be created.
	ServicePrincipalId pulumi.StringPtrInput
	// The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the value is determined by Azure Active Directory and is usually the start date of the certificate for asymmetric keys, or the current timestamp for symmetric keys. Changing this field forces a new resource to be created.
	StartDate pulumi.StringPtrInput
	// The type of key/certificate. Must be one of `AsymmetricX509Cert` or `Symmetric`. Changing this fields forces a new resource to be created.
	Type pulumi.StringPtrInput
	// The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding` argument.
	Value pulumi.StringPtrInput
}

func (ServicePrincipalCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*servicePrincipalCertificateState)(nil)).Elem()
}

type servicePrincipalCertificateArgs struct {
	// Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.
	//
	// > **Tip for Azure Key Vault** The `hex` encoding option is useful for consuming certificate data from the azurermKeyVaultCertificate resource.
	Encoding *string `pulumi:"encoding"`
	// The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
	EndDate *string `pulumi:"endDate"`
	// A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". Changing this field forces a new resource to be created.
	//
	// > One of `endDate` or `endDateRelative` must be set. The maximum duration is determined by Azure AD.
	//
	// Deprecated: The `endDateRelative` property is deprecated and will be removed in a future version of the AzureAD provider. Please instead use the Terraform `timeadd()` function to calculate a value for the `endDate` property.
	EndDateRelative *string `pulumi:"endDateRelative"`
	// A UUID used to uniquely identify this certificate. If not specified a UUID will be automatically generated. Changing this field forces a new resource to be created.
	KeyId *string `pulumi:"keyId"`
	// The ID of the service principal for which this certificate should be created. Changing this field forces a new resource to be created.
	ServicePrincipalId string `pulumi:"servicePrincipalId"`
	// The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the value is determined by Azure Active Directory and is usually the start date of the certificate for asymmetric keys, or the current timestamp for symmetric keys. Changing this field forces a new resource to be created.
	StartDate *string `pulumi:"startDate"`
	// The type of key/certificate. Must be one of `AsymmetricX509Cert` or `Symmetric`. Changing this fields forces a new resource to be created.
	Type *string `pulumi:"type"`
	// The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding` argument.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a ServicePrincipalCertificate resource.
type ServicePrincipalCertificateArgs struct {
	// Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.
	//
	// > **Tip for Azure Key Vault** The `hex` encoding option is useful for consuming certificate data from the azurermKeyVaultCertificate resource.
	Encoding pulumi.StringPtrInput
	// The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
	EndDate pulumi.StringPtrInput
	// A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". Changing this field forces a new resource to be created.
	//
	// > One of `endDate` or `endDateRelative` must be set. The maximum duration is determined by Azure AD.
	//
	// Deprecated: The `endDateRelative` property is deprecated and will be removed in a future version of the AzureAD provider. Please instead use the Terraform `timeadd()` function to calculate a value for the `endDate` property.
	EndDateRelative pulumi.StringPtrInput
	// A UUID used to uniquely identify this certificate. If not specified a UUID will be automatically generated. Changing this field forces a new resource to be created.
	KeyId pulumi.StringPtrInput
	// The ID of the service principal for which this certificate should be created. Changing this field forces a new resource to be created.
	ServicePrincipalId pulumi.StringInput
	// The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the value is determined by Azure Active Directory and is usually the start date of the certificate for asymmetric keys, or the current timestamp for symmetric keys. Changing this field forces a new resource to be created.
	StartDate pulumi.StringPtrInput
	// The type of key/certificate. Must be one of `AsymmetricX509Cert` or `Symmetric`. Changing this fields forces a new resource to be created.
	Type pulumi.StringPtrInput
	// The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding` argument.
	Value pulumi.StringInput
}

func (ServicePrincipalCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*servicePrincipalCertificateArgs)(nil)).Elem()
}

type ServicePrincipalCertificateInput interface {
	pulumi.Input

	ToServicePrincipalCertificateOutput() ServicePrincipalCertificateOutput
	ToServicePrincipalCertificateOutputWithContext(ctx context.Context) ServicePrincipalCertificateOutput
}

func (*ServicePrincipalCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePrincipalCertificate)(nil)).Elem()
}

func (i *ServicePrincipalCertificate) ToServicePrincipalCertificateOutput() ServicePrincipalCertificateOutput {
	return i.ToServicePrincipalCertificateOutputWithContext(context.Background())
}

func (i *ServicePrincipalCertificate) ToServicePrincipalCertificateOutputWithContext(ctx context.Context) ServicePrincipalCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalCertificateOutput)
}

// ServicePrincipalCertificateArrayInput is an input type that accepts ServicePrincipalCertificateArray and ServicePrincipalCertificateArrayOutput values.
// You can construct a concrete instance of `ServicePrincipalCertificateArrayInput` via:
//
//	ServicePrincipalCertificateArray{ ServicePrincipalCertificateArgs{...} }
type ServicePrincipalCertificateArrayInput interface {
	pulumi.Input

	ToServicePrincipalCertificateArrayOutput() ServicePrincipalCertificateArrayOutput
	ToServicePrincipalCertificateArrayOutputWithContext(context.Context) ServicePrincipalCertificateArrayOutput
}

type ServicePrincipalCertificateArray []ServicePrincipalCertificateInput

func (ServicePrincipalCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServicePrincipalCertificate)(nil)).Elem()
}

func (i ServicePrincipalCertificateArray) ToServicePrincipalCertificateArrayOutput() ServicePrincipalCertificateArrayOutput {
	return i.ToServicePrincipalCertificateArrayOutputWithContext(context.Background())
}

func (i ServicePrincipalCertificateArray) ToServicePrincipalCertificateArrayOutputWithContext(ctx context.Context) ServicePrincipalCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalCertificateArrayOutput)
}

// ServicePrincipalCertificateMapInput is an input type that accepts ServicePrincipalCertificateMap and ServicePrincipalCertificateMapOutput values.
// You can construct a concrete instance of `ServicePrincipalCertificateMapInput` via:
//
//	ServicePrincipalCertificateMap{ "key": ServicePrincipalCertificateArgs{...} }
type ServicePrincipalCertificateMapInput interface {
	pulumi.Input

	ToServicePrincipalCertificateMapOutput() ServicePrincipalCertificateMapOutput
	ToServicePrincipalCertificateMapOutputWithContext(context.Context) ServicePrincipalCertificateMapOutput
}

type ServicePrincipalCertificateMap map[string]ServicePrincipalCertificateInput

func (ServicePrincipalCertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServicePrincipalCertificate)(nil)).Elem()
}

func (i ServicePrincipalCertificateMap) ToServicePrincipalCertificateMapOutput() ServicePrincipalCertificateMapOutput {
	return i.ToServicePrincipalCertificateMapOutputWithContext(context.Background())
}

func (i ServicePrincipalCertificateMap) ToServicePrincipalCertificateMapOutputWithContext(ctx context.Context) ServicePrincipalCertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalCertificateMapOutput)
}

type ServicePrincipalCertificateOutput struct{ *pulumi.OutputState }

func (ServicePrincipalCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePrincipalCertificate)(nil)).Elem()
}

func (o ServicePrincipalCertificateOutput) ToServicePrincipalCertificateOutput() ServicePrincipalCertificateOutput {
	return o
}

func (o ServicePrincipalCertificateOutput) ToServicePrincipalCertificateOutputWithContext(ctx context.Context) ServicePrincipalCertificateOutput {
	return o
}

// Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.
//
// > **Tip for Azure Key Vault** The `hex` encoding option is useful for consuming certificate data from the azurermKeyVaultCertificate resource.
func (o ServicePrincipalCertificateOutput) Encoding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalCertificate) pulumi.StringPtrOutput { return v.Encoding }).(pulumi.StringPtrOutput)
}

// The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
func (o ServicePrincipalCertificateOutput) EndDate() pulumi.StringOutput {
	return o.ApplyT(func(v *ServicePrincipalCertificate) pulumi.StringOutput { return v.EndDate }).(pulumi.StringOutput)
}

// A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". Changing this field forces a new resource to be created.
//
// > One of `endDate` or `endDateRelative` must be set. The maximum duration is determined by Azure AD.
//
// Deprecated: The `endDateRelative` property is deprecated and will be removed in a future version of the AzureAD provider. Please instead use the Terraform `timeadd()` function to calculate a value for the `endDate` property.
func (o ServicePrincipalCertificateOutput) EndDateRelative() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalCertificate) pulumi.StringPtrOutput { return v.EndDateRelative }).(pulumi.StringPtrOutput)
}

// A UUID used to uniquely identify this certificate. If not specified a UUID will be automatically generated. Changing this field forces a new resource to be created.
func (o ServicePrincipalCertificateOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServicePrincipalCertificate) pulumi.StringOutput { return v.KeyId }).(pulumi.StringOutput)
}

// The ID of the service principal for which this certificate should be created. Changing this field forces a new resource to be created.
func (o ServicePrincipalCertificateOutput) ServicePrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServicePrincipalCertificate) pulumi.StringOutput { return v.ServicePrincipalId }).(pulumi.StringOutput)
}

// The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the value is determined by Azure Active Directory and is usually the start date of the certificate for asymmetric keys, or the current timestamp for symmetric keys. Changing this field forces a new resource to be created.
func (o ServicePrincipalCertificateOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v *ServicePrincipalCertificate) pulumi.StringOutput { return v.StartDate }).(pulumi.StringOutput)
}

// The type of key/certificate. Must be one of `AsymmetricX509Cert` or `Symmetric`. Changing this fields forces a new resource to be created.
func (o ServicePrincipalCertificateOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalCertificate) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

// The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding` argument.
func (o ServicePrincipalCertificateOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *ServicePrincipalCertificate) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

type ServicePrincipalCertificateArrayOutput struct{ *pulumi.OutputState }

func (ServicePrincipalCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServicePrincipalCertificate)(nil)).Elem()
}

func (o ServicePrincipalCertificateArrayOutput) ToServicePrincipalCertificateArrayOutput() ServicePrincipalCertificateArrayOutput {
	return o
}

func (o ServicePrincipalCertificateArrayOutput) ToServicePrincipalCertificateArrayOutputWithContext(ctx context.Context) ServicePrincipalCertificateArrayOutput {
	return o
}

func (o ServicePrincipalCertificateArrayOutput) Index(i pulumi.IntInput) ServicePrincipalCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServicePrincipalCertificate {
		return vs[0].([]*ServicePrincipalCertificate)[vs[1].(int)]
	}).(ServicePrincipalCertificateOutput)
}

type ServicePrincipalCertificateMapOutput struct{ *pulumi.OutputState }

func (ServicePrincipalCertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServicePrincipalCertificate)(nil)).Elem()
}

func (o ServicePrincipalCertificateMapOutput) ToServicePrincipalCertificateMapOutput() ServicePrincipalCertificateMapOutput {
	return o
}

func (o ServicePrincipalCertificateMapOutput) ToServicePrincipalCertificateMapOutputWithContext(ctx context.Context) ServicePrincipalCertificateMapOutput {
	return o
}

func (o ServicePrincipalCertificateMapOutput) MapIndex(k pulumi.StringInput) ServicePrincipalCertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServicePrincipalCertificate {
		return vs[0].(map[string]*ServicePrincipalCertificate)[vs[1].(string)]
	}).(ServicePrincipalCertificateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePrincipalCertificateInput)(nil)).Elem(), &ServicePrincipalCertificate{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePrincipalCertificateArrayInput)(nil)).Elem(), ServicePrincipalCertificateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePrincipalCertificateMapInput)(nil)).Elem(), ServicePrincipalCertificateMap{})
	pulumi.RegisterOutputType(ServicePrincipalCertificateOutput{})
	pulumi.RegisterOutputType(ServicePrincipalCertificateArrayOutput{})
	pulumi.RegisterOutputType(ServicePrincipalCertificateMapOutput{})
}

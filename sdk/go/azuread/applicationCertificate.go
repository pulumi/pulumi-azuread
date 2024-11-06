// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
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
// ### Using a certificate from Azure Key Vault
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-azure/sdk/go/azure/keyvault"
//	"github.com/pulumi/pulumi-azuread/sdk/v6/go/azuread"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleApplication, err := azuread.NewApplication(ctx, "example", &azuread.ApplicationArgs{
//				DisplayName: pulumi.String("example"),
//			})
//			if err != nil {
//				return err
//			}
//			example, err := keyvault / certificate.NewCertificate(ctx, "example", &keyvault/certificate.CertificateArgs{
//				Name:       "generated-cert",
//				KeyVaultId: exampleAzurermKeyVault.Id,
//				CertificatePolicy: map[string]interface{}{
//					"issuerParameters": map[string]interface{}{
//						"name": "Self",
//					},
//					"keyProperties": map[string]interface{}{
//						"exportable": true,
//						"keySize":    2048,
//						"keyType":    "RSA",
//						"reuseKey":   true,
//					},
//					"lifetimeActions": []map[string]interface{}{
//						map[string]interface{}{
//							"action": map[string]interface{}{
//								"actionType": "AutoRenew",
//							},
//							"trigger": map[string]interface{}{
//								"daysBeforeExpiry": 30,
//							},
//						},
//					},
//					"secretProperties": map[string]interface{}{
//						"contentType": "application/x-pkcs12",
//					},
//					"x509CertificateProperties": map[string]interface{}{
//						"extendedKeyUsages": []string{
//							"1.3.6.1.5.5.7.3.2",
//						},
//						"keyUsages": []string{
//							"dataEncipherment",
//							"digitalSignature",
//							"keyCertSign",
//							"keyEncipherment",
//						},
//						"subjectAlternativeNames": map[string]interface{}{
//							"dnsNames": []string{
//								"internal.contoso.com",
//								"domain.hello.world",
//							},
//						},
//						"subject":          fmt.Sprintf("CN=%v", exampleApplication.Name),
//						"validityInMonths": 12,
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuread.NewApplicationCertificate(ctx, "example", &azuread.ApplicationCertificateArgs{
//				ApplicationId: exampleApplication.ID(),
//				Type:          pulumi.String("AsymmetricX509Cert"),
//				Encoding:      pulumi.String("hex"),
//				Value:         example.CertificateData,
//				EndDate:       example.CertificateAttributes[0].Expires,
//				StartDate:     example.CertificateAttributes[0].NotBefore,
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
// Certificates can be imported using the object ID of the associated application and the key ID of the certificate credential, e.g.
//
// ```sh
// $ pulumi import azuread:index/applicationCertificate:ApplicationCertificate example 00000000-0000-0000-0000-000000000000/certificate/11111111-1111-1111-1111-111111111111
// ```
//
// -> This ID format is unique to Terraform and is composed of the application's object ID, the string "certificate" and the certificate's key ID in the format `{ObjectId}/certificate/{CertificateKeyId}`.
type ApplicationCertificate struct {
	pulumi.CustomResourceState

	// The resource ID of the application for which this certificate should be created. Changing this field forces a new resource to be created.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.
	//
	// > **Tip for Azure Key Vault** The `hex` encoding option is useful for consuming certificate data from the azurermKeyVaultCertificate resource.
	Encoding pulumi.StringPtrOutput `pulumi:"encoding"`
	// The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If omitted, the API will decide a suitable expiry date, which is typically around 2 years from the start date. Changing this field forces a new resource to be created.
	EndDate pulumi.StringOutput `pulumi:"endDate"`
	// A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
	//
	// > One of `endDate` or `endDateRelative` must be specified. The maximum allowed duration is determined by Azure AD and is typically around 2 years from the creation date.
	//
	// Deprecated: The `endDateRelative` property is deprecated and will be removed in a future version of the AzureAD provider. Please instead use the Terraform `timeadd()` function to calculate a value for the `endDate` property.
	EndDateRelative pulumi.StringPtrOutput `pulumi:"endDateRelative"`
	// A UUID used to uniquely identify this certificate. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.
	KeyId pulumi.StringOutput `pulumi:"keyId"`
	// The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the value is determined by Azure Active Directory and is usually the start date of the certificate for asymmetric keys, or the current timestamp for symmetric keys. Changing this field forces a new resource to be created.
	StartDate pulumi.StringOutput `pulumi:"startDate"`
	// The type of key/certificate. Must be one of `AsymmetricX509Cert` or `Symmetric`. Changing this fields forces a new resource to be created.
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding` argument.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewApplicationCertificate registers a new resource with the given unique name, arguments, and options.
func NewApplicationCertificate(ctx *pulumi.Context,
	name string, args *ApplicationCertificateArgs, opts ...pulumi.ResourceOption) (*ApplicationCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
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
	var resource ApplicationCertificate
	err := ctx.RegisterResource("azuread:index/applicationCertificate:ApplicationCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationCertificate gets an existing ApplicationCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationCertificateState, opts ...pulumi.ResourceOption) (*ApplicationCertificate, error) {
	var resource ApplicationCertificate
	err := ctx.ReadResource("azuread:index/applicationCertificate:ApplicationCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationCertificate resources.
type applicationCertificateState struct {
	// The resource ID of the application for which this certificate should be created. Changing this field forces a new resource to be created.
	ApplicationId *string `pulumi:"applicationId"`
	// Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.
	//
	// > **Tip for Azure Key Vault** The `hex` encoding option is useful for consuming certificate data from the azurermKeyVaultCertificate resource.
	Encoding *string `pulumi:"encoding"`
	// The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If omitted, the API will decide a suitable expiry date, which is typically around 2 years from the start date. Changing this field forces a new resource to be created.
	EndDate *string `pulumi:"endDate"`
	// A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
	//
	// > One of `endDate` or `endDateRelative` must be specified. The maximum allowed duration is determined by Azure AD and is typically around 2 years from the creation date.
	//
	// Deprecated: The `endDateRelative` property is deprecated and will be removed in a future version of the AzureAD provider. Please instead use the Terraform `timeadd()` function to calculate a value for the `endDate` property.
	EndDateRelative *string `pulumi:"endDateRelative"`
	// A UUID used to uniquely identify this certificate. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.
	KeyId *string `pulumi:"keyId"`
	// The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the value is determined by Azure Active Directory and is usually the start date of the certificate for asymmetric keys, or the current timestamp for symmetric keys. Changing this field forces a new resource to be created.
	StartDate *string `pulumi:"startDate"`
	// The type of key/certificate. Must be one of `AsymmetricX509Cert` or `Symmetric`. Changing this fields forces a new resource to be created.
	Type *string `pulumi:"type"`
	// The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding` argument.
	Value *string `pulumi:"value"`
}

type ApplicationCertificateState struct {
	// The resource ID of the application for which this certificate should be created. Changing this field forces a new resource to be created.
	ApplicationId pulumi.StringPtrInput
	// Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.
	//
	// > **Tip for Azure Key Vault** The `hex` encoding option is useful for consuming certificate data from the azurermKeyVaultCertificate resource.
	Encoding pulumi.StringPtrInput
	// The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If omitted, the API will decide a suitable expiry date, which is typically around 2 years from the start date. Changing this field forces a new resource to be created.
	EndDate pulumi.StringPtrInput
	// A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
	//
	// > One of `endDate` or `endDateRelative` must be specified. The maximum allowed duration is determined by Azure AD and is typically around 2 years from the creation date.
	//
	// Deprecated: The `endDateRelative` property is deprecated and will be removed in a future version of the AzureAD provider. Please instead use the Terraform `timeadd()` function to calculate a value for the `endDate` property.
	EndDateRelative pulumi.StringPtrInput
	// A UUID used to uniquely identify this certificate. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.
	KeyId pulumi.StringPtrInput
	// The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the value is determined by Azure Active Directory and is usually the start date of the certificate for asymmetric keys, or the current timestamp for symmetric keys. Changing this field forces a new resource to be created.
	StartDate pulumi.StringPtrInput
	// The type of key/certificate. Must be one of `AsymmetricX509Cert` or `Symmetric`. Changing this fields forces a new resource to be created.
	Type pulumi.StringPtrInput
	// The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding` argument.
	Value pulumi.StringPtrInput
}

func (ApplicationCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationCertificateState)(nil)).Elem()
}

type applicationCertificateArgs struct {
	// The resource ID of the application for which this certificate should be created. Changing this field forces a new resource to be created.
	ApplicationId string `pulumi:"applicationId"`
	// Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.
	//
	// > **Tip for Azure Key Vault** The `hex` encoding option is useful for consuming certificate data from the azurermKeyVaultCertificate resource.
	Encoding *string `pulumi:"encoding"`
	// The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If omitted, the API will decide a suitable expiry date, which is typically around 2 years from the start date. Changing this field forces a new resource to be created.
	EndDate *string `pulumi:"endDate"`
	// A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
	//
	// > One of `endDate` or `endDateRelative` must be specified. The maximum allowed duration is determined by Azure AD and is typically around 2 years from the creation date.
	//
	// Deprecated: The `endDateRelative` property is deprecated and will be removed in a future version of the AzureAD provider. Please instead use the Terraform `timeadd()` function to calculate a value for the `endDate` property.
	EndDateRelative *string `pulumi:"endDateRelative"`
	// A UUID used to uniquely identify this certificate. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.
	KeyId *string `pulumi:"keyId"`
	// The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the value is determined by Azure Active Directory and is usually the start date of the certificate for asymmetric keys, or the current timestamp for symmetric keys. Changing this field forces a new resource to be created.
	StartDate *string `pulumi:"startDate"`
	// The type of key/certificate. Must be one of `AsymmetricX509Cert` or `Symmetric`. Changing this fields forces a new resource to be created.
	Type *string `pulumi:"type"`
	// The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding` argument.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a ApplicationCertificate resource.
type ApplicationCertificateArgs struct {
	// The resource ID of the application for which this certificate should be created. Changing this field forces a new resource to be created.
	ApplicationId pulumi.StringInput
	// Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.
	//
	// > **Tip for Azure Key Vault** The `hex` encoding option is useful for consuming certificate data from the azurermKeyVaultCertificate resource.
	Encoding pulumi.StringPtrInput
	// The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If omitted, the API will decide a suitable expiry date, which is typically around 2 years from the start date. Changing this field forces a new resource to be created.
	EndDate pulumi.StringPtrInput
	// A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
	//
	// > One of `endDate` or `endDateRelative` must be specified. The maximum allowed duration is determined by Azure AD and is typically around 2 years from the creation date.
	//
	// Deprecated: The `endDateRelative` property is deprecated and will be removed in a future version of the AzureAD provider. Please instead use the Terraform `timeadd()` function to calculate a value for the `endDate` property.
	EndDateRelative pulumi.StringPtrInput
	// A UUID used to uniquely identify this certificate. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.
	KeyId pulumi.StringPtrInput
	// The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the value is determined by Azure Active Directory and is usually the start date of the certificate for asymmetric keys, or the current timestamp for symmetric keys. Changing this field forces a new resource to be created.
	StartDate pulumi.StringPtrInput
	// The type of key/certificate. Must be one of `AsymmetricX509Cert` or `Symmetric`. Changing this fields forces a new resource to be created.
	Type pulumi.StringPtrInput
	// The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding` argument.
	Value pulumi.StringInput
}

func (ApplicationCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationCertificateArgs)(nil)).Elem()
}

type ApplicationCertificateInput interface {
	pulumi.Input

	ToApplicationCertificateOutput() ApplicationCertificateOutput
	ToApplicationCertificateOutputWithContext(ctx context.Context) ApplicationCertificateOutput
}

func (*ApplicationCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationCertificate)(nil)).Elem()
}

func (i *ApplicationCertificate) ToApplicationCertificateOutput() ApplicationCertificateOutput {
	return i.ToApplicationCertificateOutputWithContext(context.Background())
}

func (i *ApplicationCertificate) ToApplicationCertificateOutputWithContext(ctx context.Context) ApplicationCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationCertificateOutput)
}

// ApplicationCertificateArrayInput is an input type that accepts ApplicationCertificateArray and ApplicationCertificateArrayOutput values.
// You can construct a concrete instance of `ApplicationCertificateArrayInput` via:
//
//	ApplicationCertificateArray{ ApplicationCertificateArgs{...} }
type ApplicationCertificateArrayInput interface {
	pulumi.Input

	ToApplicationCertificateArrayOutput() ApplicationCertificateArrayOutput
	ToApplicationCertificateArrayOutputWithContext(context.Context) ApplicationCertificateArrayOutput
}

type ApplicationCertificateArray []ApplicationCertificateInput

func (ApplicationCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationCertificate)(nil)).Elem()
}

func (i ApplicationCertificateArray) ToApplicationCertificateArrayOutput() ApplicationCertificateArrayOutput {
	return i.ToApplicationCertificateArrayOutputWithContext(context.Background())
}

func (i ApplicationCertificateArray) ToApplicationCertificateArrayOutputWithContext(ctx context.Context) ApplicationCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationCertificateArrayOutput)
}

// ApplicationCertificateMapInput is an input type that accepts ApplicationCertificateMap and ApplicationCertificateMapOutput values.
// You can construct a concrete instance of `ApplicationCertificateMapInput` via:
//
//	ApplicationCertificateMap{ "key": ApplicationCertificateArgs{...} }
type ApplicationCertificateMapInput interface {
	pulumi.Input

	ToApplicationCertificateMapOutput() ApplicationCertificateMapOutput
	ToApplicationCertificateMapOutputWithContext(context.Context) ApplicationCertificateMapOutput
}

type ApplicationCertificateMap map[string]ApplicationCertificateInput

func (ApplicationCertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationCertificate)(nil)).Elem()
}

func (i ApplicationCertificateMap) ToApplicationCertificateMapOutput() ApplicationCertificateMapOutput {
	return i.ToApplicationCertificateMapOutputWithContext(context.Background())
}

func (i ApplicationCertificateMap) ToApplicationCertificateMapOutputWithContext(ctx context.Context) ApplicationCertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationCertificateMapOutput)
}

type ApplicationCertificateOutput struct{ *pulumi.OutputState }

func (ApplicationCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationCertificate)(nil)).Elem()
}

func (o ApplicationCertificateOutput) ToApplicationCertificateOutput() ApplicationCertificateOutput {
	return o
}

func (o ApplicationCertificateOutput) ToApplicationCertificateOutputWithContext(ctx context.Context) ApplicationCertificateOutput {
	return o
}

// The resource ID of the application for which this certificate should be created. Changing this field forces a new resource to be created.
func (o ApplicationCertificateOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationCertificate) pulumi.StringOutput { return v.ApplicationId }).(pulumi.StringOutput)
}

// Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.
//
// > **Tip for Azure Key Vault** The `hex` encoding option is useful for consuming certificate data from the azurermKeyVaultCertificate resource.
func (o ApplicationCertificateOutput) Encoding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationCertificate) pulumi.StringPtrOutput { return v.Encoding }).(pulumi.StringPtrOutput)
}

// The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If omitted, the API will decide a suitable expiry date, which is typically around 2 years from the start date. Changing this field forces a new resource to be created.
func (o ApplicationCertificateOutput) EndDate() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationCertificate) pulumi.StringOutput { return v.EndDate }).(pulumi.StringOutput)
}

// A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
//
// > One of `endDate` or `endDateRelative` must be specified. The maximum allowed duration is determined by Azure AD and is typically around 2 years from the creation date.
//
// Deprecated: The `endDateRelative` property is deprecated and will be removed in a future version of the AzureAD provider. Please instead use the Terraform `timeadd()` function to calculate a value for the `endDate` property.
func (o ApplicationCertificateOutput) EndDateRelative() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationCertificate) pulumi.StringPtrOutput { return v.EndDateRelative }).(pulumi.StringPtrOutput)
}

// A UUID used to uniquely identify this certificate. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.
func (o ApplicationCertificateOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationCertificate) pulumi.StringOutput { return v.KeyId }).(pulumi.StringOutput)
}

// The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the value is determined by Azure Active Directory and is usually the start date of the certificate for asymmetric keys, or the current timestamp for symmetric keys. Changing this field forces a new resource to be created.
func (o ApplicationCertificateOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationCertificate) pulumi.StringOutput { return v.StartDate }).(pulumi.StringOutput)
}

// The type of key/certificate. Must be one of `AsymmetricX509Cert` or `Symmetric`. Changing this fields forces a new resource to be created.
func (o ApplicationCertificateOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationCertificate) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

// The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding` argument.
func (o ApplicationCertificateOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationCertificate) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

type ApplicationCertificateArrayOutput struct{ *pulumi.OutputState }

func (ApplicationCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationCertificate)(nil)).Elem()
}

func (o ApplicationCertificateArrayOutput) ToApplicationCertificateArrayOutput() ApplicationCertificateArrayOutput {
	return o
}

func (o ApplicationCertificateArrayOutput) ToApplicationCertificateArrayOutputWithContext(ctx context.Context) ApplicationCertificateArrayOutput {
	return o
}

func (o ApplicationCertificateArrayOutput) Index(i pulumi.IntInput) ApplicationCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplicationCertificate {
		return vs[0].([]*ApplicationCertificate)[vs[1].(int)]
	}).(ApplicationCertificateOutput)
}

type ApplicationCertificateMapOutput struct{ *pulumi.OutputState }

func (ApplicationCertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationCertificate)(nil)).Elem()
}

func (o ApplicationCertificateMapOutput) ToApplicationCertificateMapOutput() ApplicationCertificateMapOutput {
	return o
}

func (o ApplicationCertificateMapOutput) ToApplicationCertificateMapOutputWithContext(ctx context.Context) ApplicationCertificateMapOutput {
	return o
}

func (o ApplicationCertificateMapOutput) MapIndex(k pulumi.StringInput) ApplicationCertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplicationCertificate {
		return vs[0].(map[string]*ApplicationCertificate)[vs[1].(string)]
	}).(ApplicationCertificateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationCertificateInput)(nil)).Elem(), &ApplicationCertificate{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationCertificateArrayInput)(nil)).Elem(), ApplicationCertificateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationCertificateMapInput)(nil)).Elem(), ApplicationCertificateMap{})
	pulumi.RegisterOutputType(ApplicationCertificateOutput{})
	pulumi.RegisterOutputType(ApplicationCertificateArrayOutput{})
	pulumi.RegisterOutputType(ApplicationCertificateMapOutput{})
}

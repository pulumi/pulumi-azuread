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
// *Using a PEM certificate*
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-azuread/sdk/v5/go/azuread"
//	"github.com/pulumi/pulumi-std/sdk/go/std"
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
//			invokeFile, err := std.File(ctx, &std.FileArgs{
//				Input: "cert.pem",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = azuread.NewApplicationCertificate(ctx, "example", &azuread.ApplicationCertificateArgs{
//				ApplicationId: example.ID(),
//				Type:          pulumi.String("AsymmetricX509Cert"),
//				Value:         invokeFile.Result,
//				EndDate:       pulumi.String("2021-05-01T01:02:03Z"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// *Using a DER certificate*
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-azuread/sdk/v5/go/azuread"
//	"github.com/pulumi/pulumi-std/sdk/go/std"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
// func main() {
// pulumi.Run(func(ctx *pulumi.Context) error {
// example, err := azuread.NewApplicationRegistration(ctx, "example", &azuread.ApplicationRegistrationArgs{
// DisplayName: pulumi.String("example"),
// })
// if err != nil {
// return err
// }
// invokeBase64encode, err := std.Base64encode(ctx, invokeFile1, err := std.File(ctx, &std.FileArgs{
// Input: "cert.der",
// }, nil)
// if err != nil {
// return err
// }
// &std.Base64encodeArgs{
// Input: invokeFile1.Result,
// }, nil)
// if err != nil {
// return err
// }
// _, err = azuread.NewApplicationCertificate(ctx, "example", &azuread.ApplicationCertificateArgs{
// ApplicationId: example.ID(),
// Type: pulumi.String("AsymmetricX509Cert"),
// Encoding: pulumi.String("base64"),
// Value: invokeBase64encode.Result,
// EndDate: pulumi.String("2021-05-01T01:02:03Z"),
// })
// if err != nil {
// return err
// }
// return nil
// })
// }
// ```
// <!--End PulumiCodeChooser -->
//
// ### Using a certificate from Azure Key Vault
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-azure/sdk/v5/go/azure/keyvault"
//	"github.com/pulumi/pulumi-azuread/sdk/v5/go/azuread"
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
//			example, err := keyvault.NewCertificate(ctx, "example", &keyvault.CertificateArgs{
//				Name:       pulumi.String("generated-cert"),
//				KeyVaultId: pulumi.Any(exampleAzurermKeyVault.Id),
//				CertificatePolicy: &keyvault.CertificateCertificatePolicyArgs{
//					IssuerParameters: &keyvault.CertificateCertificatePolicyIssuerParametersArgs{
//						Name: pulumi.String("Self"),
//					},
//					KeyProperties: &keyvault.CertificateCertificatePolicyKeyPropertiesArgs{
//						Exportable: pulumi.Bool(true),
//						KeySize:    pulumi.Int(2048),
//						KeyType:    pulumi.String("RSA"),
//						ReuseKey:   pulumi.Bool(true),
//					},
//					LifetimeActions: keyvault.CertificateCertificatePolicyLifetimeActionArray{
//						&keyvault.CertificateCertificatePolicyLifetimeActionArgs{
//							Action: &keyvault.CertificateCertificatePolicyLifetimeActionActionArgs{
//								ActionType: pulumi.String("AutoRenew"),
//							},
//							Trigger: &keyvault.CertificateCertificatePolicyLifetimeActionTriggerArgs{
//								DaysBeforeExpiry: pulumi.Int(30),
//							},
//						},
//					},
//					SecretProperties: &keyvault.CertificateCertificatePolicySecretPropertiesArgs{
//						ContentType: pulumi.String("application/x-pkcs12"),
//					},
//					X509CertificateProperties: &keyvault.CertificateCertificatePolicyX509CertificatePropertiesArgs{
//						ExtendedKeyUsages: pulumi.StringArray{
//							pulumi.String("1.3.6.1.5.5.7.3.2"),
//						},
//						KeyUsages: pulumi.StringArray{
//							pulumi.String("dataEncipherment"),
//							pulumi.String("digitalSignature"),
//							pulumi.String("keyCertSign"),
//							pulumi.String("keyEncipherment"),
//						},
//						SubjectAlternativeNames: &keyvault.CertificateCertificatePolicyX509CertificatePropertiesSubjectAlternativeNamesArgs{
//							DnsNames: pulumi.StringArray{
//								pulumi.String("internal.contoso.com"),
//								pulumi.String("domain.hello.world"),
//							},
//						},
//						Subject:          pulumi.String(fmt.Sprintf("CN=%v", exampleApplication.Name)),
//						ValidityInMonths: pulumi.Int(12),
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
//				EndDate: example.CertificateAttributes.ApplyT(func(certificateAttributes []keyvault.CertificateCertificateAttribute) (*string, error) {
//					return &certificateAttributes[0].Expires, nil
//				}).(pulumi.StringPtrOutput),
//				StartDate: example.CertificateAttributes.ApplyT(func(certificateAttributes []keyvault.CertificateCertificateAttribute) (*string, error) {
//					return &certificateAttributes[0].NotBefore, nil
//				}).(pulumi.StringPtrOutput),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
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
	// The object ID of the application for which this certificate should be created
	//
	// Deprecated: The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider
	ApplicationObjectId pulumi.StringOutput `pulumi:"applicationObjectId"`
	// Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.
	//
	// > **Tip for Azure Key Vault** The `hex` encoding option is useful for consuming certificate data from the azurermKeyVaultCertificate resource.
	Encoding pulumi.StringPtrOutput `pulumi:"encoding"`
	// The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If omitted, the API will decide a suitable expiry date, which is typically around 2 years from the start date. Changing this field forces a new resource to be created.
	EndDate pulumi.StringOutput `pulumi:"endDate"`
	// A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
	//
	// > One of `endDate` or `endDateRelative` must be specified. The maximum allowed duration is determined by Azure AD and is typically around 2 years from the creation date.
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
	// The object ID of the application for which this certificate should be created
	//
	// Deprecated: The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider
	ApplicationObjectId *string `pulumi:"applicationObjectId"`
	// Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.
	//
	// > **Tip for Azure Key Vault** The `hex` encoding option is useful for consuming certificate data from the azurermKeyVaultCertificate resource.
	Encoding *string `pulumi:"encoding"`
	// The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If omitted, the API will decide a suitable expiry date, which is typically around 2 years from the start date. Changing this field forces a new resource to be created.
	EndDate *string `pulumi:"endDate"`
	// A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
	//
	// > One of `endDate` or `endDateRelative` must be specified. The maximum allowed duration is determined by Azure AD and is typically around 2 years from the creation date.
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
	// The object ID of the application for which this certificate should be created
	//
	// Deprecated: The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider
	ApplicationObjectId pulumi.StringPtrInput
	// Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.
	//
	// > **Tip for Azure Key Vault** The `hex` encoding option is useful for consuming certificate data from the azurermKeyVaultCertificate resource.
	Encoding pulumi.StringPtrInput
	// The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If omitted, the API will decide a suitable expiry date, which is typically around 2 years from the start date. Changing this field forces a new resource to be created.
	EndDate pulumi.StringPtrInput
	// A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
	//
	// > One of `endDate` or `endDateRelative` must be specified. The maximum allowed duration is determined by Azure AD and is typically around 2 years from the creation date.
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
	ApplicationId *string `pulumi:"applicationId"`
	// The object ID of the application for which this certificate should be created
	//
	// Deprecated: The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider
	ApplicationObjectId *string `pulumi:"applicationObjectId"`
	// Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.
	//
	// > **Tip for Azure Key Vault** The `hex` encoding option is useful for consuming certificate data from the azurermKeyVaultCertificate resource.
	Encoding *string `pulumi:"encoding"`
	// The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If omitted, the API will decide a suitable expiry date, which is typically around 2 years from the start date. Changing this field forces a new resource to be created.
	EndDate *string `pulumi:"endDate"`
	// A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
	//
	// > One of `endDate` or `endDateRelative` must be specified. The maximum allowed duration is determined by Azure AD and is typically around 2 years from the creation date.
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
	ApplicationId pulumi.StringPtrInput
	// The object ID of the application for which this certificate should be created
	//
	// Deprecated: The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider
	ApplicationObjectId pulumi.StringPtrInput
	// Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.
	//
	// > **Tip for Azure Key Vault** The `hex` encoding option is useful for consuming certificate data from the azurermKeyVaultCertificate resource.
	Encoding pulumi.StringPtrInput
	// The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If omitted, the API will decide a suitable expiry date, which is typically around 2 years from the start date. Changing this field forces a new resource to be created.
	EndDate pulumi.StringPtrInput
	// A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
	//
	// > One of `endDate` or `endDateRelative` must be specified. The maximum allowed duration is determined by Azure AD and is typically around 2 years from the creation date.
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

// The object ID of the application for which this certificate should be created
//
// Deprecated: The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider
func (o ApplicationCertificateOutput) ApplicationObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationCertificate) pulumi.StringOutput { return v.ApplicationObjectId }).(pulumi.StringOutput)
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

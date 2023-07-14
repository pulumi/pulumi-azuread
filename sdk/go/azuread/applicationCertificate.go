// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplicationCertificate struct {
	pulumi.CustomResourceState

	// The object ID of the application for which this certificate should be created
	ApplicationObjectId pulumi.StringOutput `pulumi:"applicationObjectId"`
	// Specifies the encoding used for the supplied certificate data
	Encoding pulumi.StringPtrOutput `pulumi:"encoding"`
	// The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If
	// omitted, the API will decide a suitable expiry date, which is typically around 2 years from the start date
	EndDate pulumi.StringOutput `pulumi:"endDate"`
	// A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`
	EndDateRelative pulumi.StringPtrOutput `pulumi:"endDateRelative"`
	// A UUID used to uniquely identify this certificate. If omitted, a random UUID will be automatically generated
	KeyId pulumi.StringOutput `pulumi:"keyId"`
	// The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
	// If this isn't specified, the current date and time are use
	StartDate pulumi.StringOutput `pulumi:"startDate"`
	// The type of key/certificate
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding`
	// argumen
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewApplicationCertificate registers a new resource with the given unique name, arguments, and options.
func NewApplicationCertificate(ctx *pulumi.Context,
	name string, args *ApplicationCertificateArgs, opts ...pulumi.ResourceOption) (*ApplicationCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationObjectId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationObjectId'")
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
	// The object ID of the application for which this certificate should be created
	ApplicationObjectId *string `pulumi:"applicationObjectId"`
	// Specifies the encoding used for the supplied certificate data
	Encoding *string `pulumi:"encoding"`
	// The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If
	// omitted, the API will decide a suitable expiry date, which is typically around 2 years from the start date
	EndDate *string `pulumi:"endDate"`
	// A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`
	EndDateRelative *string `pulumi:"endDateRelative"`
	// A UUID used to uniquely identify this certificate. If omitted, a random UUID will be automatically generated
	KeyId *string `pulumi:"keyId"`
	// The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
	// If this isn't specified, the current date and time are use
	StartDate *string `pulumi:"startDate"`
	// The type of key/certificate
	Type *string `pulumi:"type"`
	// The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding`
	// argumen
	Value *string `pulumi:"value"`
}

type ApplicationCertificateState struct {
	// The object ID of the application for which this certificate should be created
	ApplicationObjectId pulumi.StringPtrInput
	// Specifies the encoding used for the supplied certificate data
	Encoding pulumi.StringPtrInput
	// The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If
	// omitted, the API will decide a suitable expiry date, which is typically around 2 years from the start date
	EndDate pulumi.StringPtrInput
	// A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`
	EndDateRelative pulumi.StringPtrInput
	// A UUID used to uniquely identify this certificate. If omitted, a random UUID will be automatically generated
	KeyId pulumi.StringPtrInput
	// The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
	// If this isn't specified, the current date and time are use
	StartDate pulumi.StringPtrInput
	// The type of key/certificate
	Type pulumi.StringPtrInput
	// The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding`
	// argumen
	Value pulumi.StringPtrInput
}

func (ApplicationCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationCertificateState)(nil)).Elem()
}

type applicationCertificateArgs struct {
	// The object ID of the application for which this certificate should be created
	ApplicationObjectId string `pulumi:"applicationObjectId"`
	// Specifies the encoding used for the supplied certificate data
	Encoding *string `pulumi:"encoding"`
	// The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If
	// omitted, the API will decide a suitable expiry date, which is typically around 2 years from the start date
	EndDate *string `pulumi:"endDate"`
	// A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`
	EndDateRelative *string `pulumi:"endDateRelative"`
	// A UUID used to uniquely identify this certificate. If omitted, a random UUID will be automatically generated
	KeyId *string `pulumi:"keyId"`
	// The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
	// If this isn't specified, the current date and time are use
	StartDate *string `pulumi:"startDate"`
	// The type of key/certificate
	Type *string `pulumi:"type"`
	// The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding`
	// argumen
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a ApplicationCertificate resource.
type ApplicationCertificateArgs struct {
	// The object ID of the application for which this certificate should be created
	ApplicationObjectId pulumi.StringInput
	// Specifies the encoding used for the supplied certificate data
	Encoding pulumi.StringPtrInput
	// The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If
	// omitted, the API will decide a suitable expiry date, which is typically around 2 years from the start date
	EndDate pulumi.StringPtrInput
	// A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`
	EndDateRelative pulumi.StringPtrInput
	// A UUID used to uniquely identify this certificate. If omitted, a random UUID will be automatically generated
	KeyId pulumi.StringPtrInput
	// The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
	// If this isn't specified, the current date and time are use
	StartDate pulumi.StringPtrInput
	// The type of key/certificate
	Type pulumi.StringPtrInput
	// The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding`
	// argumen
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

// The object ID of the application for which this certificate should be created
func (o ApplicationCertificateOutput) ApplicationObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationCertificate) pulumi.StringOutput { return v.ApplicationObjectId }).(pulumi.StringOutput)
}

// Specifies the encoding used for the supplied certificate data
func (o ApplicationCertificateOutput) Encoding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationCertificate) pulumi.StringPtrOutput { return v.Encoding }).(pulumi.StringPtrOutput)
}

// The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If
// omitted, the API will decide a suitable expiry date, which is typically around 2 years from the start date
func (o ApplicationCertificateOutput) EndDate() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationCertificate) pulumi.StringOutput { return v.EndDate }).(pulumi.StringOutput)
}

// A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`
func (o ApplicationCertificateOutput) EndDateRelative() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationCertificate) pulumi.StringPtrOutput { return v.EndDateRelative }).(pulumi.StringPtrOutput)
}

// A UUID used to uniquely identify this certificate. If omitted, a random UUID will be automatically generated
func (o ApplicationCertificateOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationCertificate) pulumi.StringOutput { return v.KeyId }).(pulumi.StringOutput)
}

// The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
// If this isn't specified, the current date and time are use
func (o ApplicationCertificateOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationCertificate) pulumi.StringOutput { return v.StartDate }).(pulumi.StringOutput)
}

// The type of key/certificate
func (o ApplicationCertificateOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationCertificate) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

// The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding`
// argumen
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

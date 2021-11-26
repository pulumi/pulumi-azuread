// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The provider type for the azuread package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState

	// Base64 encoded PKCS#12 certificate bundle to use when authenticating as a Service Principal using a Client Certificate
	ClientCertificate pulumi.StringPtrOutput `pulumi:"clientCertificate"`
	// The password to decrypt the Client Certificate. For use when authenticating as a Service Principal using a Client
	// Certificate
	ClientCertificatePassword pulumi.StringPtrOutput `pulumi:"clientCertificatePassword"`
	// The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
	// Principal using a Client Certificate
	ClientCertificatePath pulumi.StringPtrOutput `pulumi:"clientCertificatePath"`
	// The Client ID which should be used for service principal authentication
	ClientId pulumi.StringPtrOutput `pulumi:"clientId"`
	// The application password to use when authenticating as a Service Principal using a Client Secret
	ClientSecret pulumi.StringPtrOutput `pulumi:"clientSecret"`
	// The cloud environment which should be used. Possible values are: `global` (also `public`), `usgovernmentl4` (also
	// `usgovernment`), `usgovernmentl5` (also `dod`), and `china`. Defaults to `global`
	Environment pulumi.StringPtrOutput `pulumi:"environment"`
	// The path to a custom endpoint for Managed Identity - in most circumstances this should be detected automatically
	MsiEndpoint pulumi.StringPtrOutput `pulumi:"msiEndpoint"`
	// A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution
	PartnerId pulumi.StringPtrOutput `pulumi:"partnerId"`
	// The Tenant ID which should be used. Works with all authentication methods except Managed Identity
	TenantId pulumi.StringPtrOutput `pulumi:"tenantId"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}

	if args.Environment == nil {
		args.Environment = pulumi.StringPtr(getEnvOrDefault("public", nil, "ARM_ENVIRONMENT").(string))
	}
	if args.MsiEndpoint == nil {
		args.MsiEndpoint = pulumi.StringPtr(getEnvOrDefault("", nil, "ARM_MSI_ENDPOINT").(string))
	}
	if args.UseMsi == nil {
		args.UseMsi = pulumi.BoolPtr(getEnvOrDefault(false, parseEnvBool, "ARM_USE_MSI").(bool))
	}
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:azuread", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// Base64 encoded PKCS#12 certificate bundle to use when authenticating as a Service Principal using a Client Certificate
	ClientCertificate *string `pulumi:"clientCertificate"`
	// The password to decrypt the Client Certificate. For use when authenticating as a Service Principal using a Client
	// Certificate
	ClientCertificatePassword *string `pulumi:"clientCertificatePassword"`
	// The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
	// Principal using a Client Certificate
	ClientCertificatePath *string `pulumi:"clientCertificatePath"`
	// The Client ID which should be used for service principal authentication
	ClientId *string `pulumi:"clientId"`
	// The application password to use when authenticating as a Service Principal using a Client Secret
	ClientSecret *string `pulumi:"clientSecret"`
	// Disable the Terraform Partner ID, which is used if a custom `partner_id` isn't specified
	DisableTerraformPartnerId *bool `pulumi:"disableTerraformPartnerId"`
	// The cloud environment which should be used. Possible values are: `global` (also `public`), `usgovernmentl4` (also
	// `usgovernment`), `usgovernmentl5` (also `dod`), and `china`. Defaults to `global`
	Environment *string `pulumi:"environment"`
	// The path to a custom endpoint for Managed Identity - in most circumstances this should be detected automatically
	MsiEndpoint *string `pulumi:"msiEndpoint"`
	// A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution
	PartnerId *string `pulumi:"partnerId"`
	// The Tenant ID which should be used. Works with all authentication methods except Managed Identity
	TenantId *string `pulumi:"tenantId"`
	// Allow Azure CLI to be used for Authentication
	UseCli *bool `pulumi:"useCli"`
	// Allow Managed Identity to be used for Authentication
	UseMsi *bool `pulumi:"useMsi"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// Base64 encoded PKCS#12 certificate bundle to use when authenticating as a Service Principal using a Client Certificate
	ClientCertificate pulumi.StringPtrInput
	// The password to decrypt the Client Certificate. For use when authenticating as a Service Principal using a Client
	// Certificate
	ClientCertificatePassword pulumi.StringPtrInput
	// The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
	// Principal using a Client Certificate
	ClientCertificatePath pulumi.StringPtrInput
	// The Client ID which should be used for service principal authentication
	ClientId pulumi.StringPtrInput
	// The application password to use when authenticating as a Service Principal using a Client Secret
	ClientSecret pulumi.StringPtrInput
	// Disable the Terraform Partner ID, which is used if a custom `partner_id` isn't specified
	DisableTerraformPartnerId pulumi.BoolPtrInput
	// The cloud environment which should be used. Possible values are: `global` (also `public`), `usgovernmentl4` (also
	// `usgovernment`), `usgovernmentl5` (also `dod`), and `china`. Defaults to `global`
	Environment pulumi.StringPtrInput
	// The path to a custom endpoint for Managed Identity - in most circumstances this should be detected automatically
	MsiEndpoint pulumi.StringPtrInput
	// A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution
	PartnerId pulumi.StringPtrInput
	// The Tenant ID which should be used. Works with all authentication methods except Managed Identity
	TenantId pulumi.StringPtrInput
	// Allow Azure CLI to be used for Authentication
	UseCli pulumi.BoolPtrInput
	// Allow Managed Identity to be used for Authentication
	UseMsi pulumi.BoolPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((*Provider)(nil))
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

func (i *Provider) ToProviderPtrOutput() ProviderPtrOutput {
	return i.ToProviderPtrOutputWithContext(context.Background())
}

func (i *Provider) ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderPtrOutput)
}

type ProviderPtrInput interface {
	pulumi.Input

	ToProviderPtrOutput() ProviderPtrOutput
	ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput
}

type providerPtrType ProviderArgs

func (*providerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil))
}

func (i *providerPtrType) ToProviderPtrOutput() ProviderPtrOutput {
	return i.ToProviderPtrOutputWithContext(context.Background())
}

func (i *providerPtrType) ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderPtrOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Provider)(nil))
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderPtrOutput() ProviderPtrOutput {
	return o.ToProviderPtrOutputWithContext(context.Background())
}

func (o ProviderOutput) ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Provider) *Provider {
		return &v
	}).(ProviderPtrOutput)
}

type ProviderPtrOutput struct{ *pulumi.OutputState }

func (ProviderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil))
}

func (o ProviderPtrOutput) ToProviderPtrOutput() ProviderPtrOutput {
	return o
}

func (o ProviderPtrOutput) ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput {
	return o
}

func (o ProviderPtrOutput) Elem() ProviderOutput {
	return o.ApplyT(func(v *Provider) Provider {
		if v != nil {
			return *v
		}
		var ret Provider
		return ret
	}).(ProviderOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderPtrInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
	pulumi.RegisterOutputType(ProviderPtrOutput{})
}

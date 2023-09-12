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
	// The Hostname which should be used for the Azure Metadata Service.
	MetadataHost pulumi.StringOutput `pulumi:"metadataHost"`
	// The path to a custom endpoint for Managed Identity - in most circumstances this should be detected automatically
	MsiEndpoint pulumi.StringPtrOutput `pulumi:"msiEndpoint"`
	// The bearer token for the request to the OIDC provider. For use when authenticating as a Service Principal using OpenID
	// Connect.
	OidcRequestToken pulumi.StringPtrOutput `pulumi:"oidcRequestToken"`
	// The URL for the OIDC provider from which to request an ID token. For use when authenticating as a Service Principal
	// using OpenID Connect.
	OidcRequestUrl pulumi.StringPtrOutput `pulumi:"oidcRequestUrl"`
	// The ID token for use when authenticating as a Service Principal using OpenID Connect.
	OidcToken pulumi.StringPtrOutput `pulumi:"oidcToken"`
	// The path to a file containing an ID token for use when authenticating as a Service Principal using OpenID Connect.
	OidcTokenFilePath pulumi.StringPtrOutput `pulumi:"oidcTokenFilePath"`
	// A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution
	PartnerId pulumi.StringPtrOutput `pulumi:"partnerId"`
	// The Tenant ID which should be used. Works with all authentication methods except Managed Identity
	TenantId pulumi.StringPtrOutput `pulumi:"tenantId"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MetadataHost == nil {
		return nil, errors.New("invalid value for required argument 'MetadataHost'")
	}
	if args.Environment == nil {
		if d := internal.GetEnvOrDefault("public", nil, "ARM_ENVIRONMENT"); d != nil {
			args.Environment = pulumi.StringPtr(d.(string))
		}
	}
	if args.MsiEndpoint == nil {
		if d := internal.GetEnvOrDefault(nil, nil, "ARM_MSI_ENDPOINT"); d != nil {
			args.MsiEndpoint = pulumi.StringPtr(d.(string))
		}
	}
	if args.UseMsi == nil {
		if d := internal.GetEnvOrDefault(false, internal.ParseEnvBool, "ARM_USE_MSI"); d != nil {
			args.UseMsi = pulumi.BoolPtr(d.(bool))
		}
	}
	opts = internal.PkgResourceDefaultOpts(opts)
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
	// The Hostname which should be used for the Azure Metadata Service.
	MetadataHost string `pulumi:"metadataHost"`
	// The path to a custom endpoint for Managed Identity - in most circumstances this should be detected automatically
	MsiEndpoint *string `pulumi:"msiEndpoint"`
	// The bearer token for the request to the OIDC provider. For use when authenticating as a Service Principal using OpenID
	// Connect.
	OidcRequestToken *string `pulumi:"oidcRequestToken"`
	// The URL for the OIDC provider from which to request an ID token. For use when authenticating as a Service Principal
	// using OpenID Connect.
	OidcRequestUrl *string `pulumi:"oidcRequestUrl"`
	// The ID token for use when authenticating as a Service Principal using OpenID Connect.
	OidcToken *string `pulumi:"oidcToken"`
	// The path to a file containing an ID token for use when authenticating as a Service Principal using OpenID Connect.
	OidcTokenFilePath *string `pulumi:"oidcTokenFilePath"`
	// A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution
	PartnerId *string `pulumi:"partnerId"`
	// The Tenant ID which should be used. Works with all authentication methods except Managed Identity
	TenantId *string `pulumi:"tenantId"`
	// Allow Azure CLI to be used for Authentication
	UseCli *bool `pulumi:"useCli"`
	// Allow Managed Identity to be used for Authentication
	UseMsi *bool `pulumi:"useMsi"`
	// Allow OpenID Connect to be used for authentication
	UseOidc *bool `pulumi:"useOidc"`
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
	// The Hostname which should be used for the Azure Metadata Service.
	MetadataHost pulumi.StringInput
	// The path to a custom endpoint for Managed Identity - in most circumstances this should be detected automatically
	MsiEndpoint pulumi.StringPtrInput
	// The bearer token for the request to the OIDC provider. For use when authenticating as a Service Principal using OpenID
	// Connect.
	OidcRequestToken pulumi.StringPtrInput
	// The URL for the OIDC provider from which to request an ID token. For use when authenticating as a Service Principal
	// using OpenID Connect.
	OidcRequestUrl pulumi.StringPtrInput
	// The ID token for use when authenticating as a Service Principal using OpenID Connect.
	OidcToken pulumi.StringPtrInput
	// The path to a file containing an ID token for use when authenticating as a Service Principal using OpenID Connect.
	OidcTokenFilePath pulumi.StringPtrInput
	// A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution
	PartnerId pulumi.StringPtrInput
	// The Tenant ID which should be used. Works with all authentication methods except Managed Identity
	TenantId pulumi.StringPtrInput
	// Allow Azure CLI to be used for Authentication
	UseCli pulumi.BoolPtrInput
	// Allow Managed Identity to be used for Authentication
	UseMsi pulumi.BoolPtrInput
	// Allow OpenID Connect to be used for authentication
	UseOidc pulumi.BoolPtrInput
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
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

func (i *Provider) ToOutput(ctx context.Context) pulumix.Output[*Provider] {
	return pulumix.Output[*Provider]{
		OutputState: i.ToProviderOutputWithContext(ctx).OutputState,
	}
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

func (o ProviderOutput) ToOutput(ctx context.Context) pulumix.Output[*Provider] {
	return pulumix.Output[*Provider]{
		OutputState: o.OutputState,
	}
}

// Base64 encoded PKCS#12 certificate bundle to use when authenticating as a Service Principal using a Client Certificate
func (o ProviderOutput) ClientCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.ClientCertificate }).(pulumi.StringPtrOutput)
}

// The password to decrypt the Client Certificate. For use when authenticating as a Service Principal using a Client
// Certificate
func (o ProviderOutput) ClientCertificatePassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.ClientCertificatePassword }).(pulumi.StringPtrOutput)
}

// The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
// Principal using a Client Certificate
func (o ProviderOutput) ClientCertificatePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.ClientCertificatePath }).(pulumi.StringPtrOutput)
}

// The Client ID which should be used for service principal authentication
func (o ProviderOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.ClientId }).(pulumi.StringPtrOutput)
}

// The application password to use when authenticating as a Service Principal using a Client Secret
func (o ProviderOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

// The cloud environment which should be used. Possible values are: `global` (also `public`), `usgovernmentl4` (also
// `usgovernment`), `usgovernmentl5` (also `dod`), and `china`. Defaults to `global`
func (o ProviderOutput) Environment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Environment }).(pulumi.StringPtrOutput)
}

// The Hostname which should be used for the Azure Metadata Service.
func (o ProviderOutput) MetadataHost() pulumi.StringOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringOutput { return v.MetadataHost }).(pulumi.StringOutput)
}

// The path to a custom endpoint for Managed Identity - in most circumstances this should be detected automatically
func (o ProviderOutput) MsiEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.MsiEndpoint }).(pulumi.StringPtrOutput)
}

// The bearer token for the request to the OIDC provider. For use when authenticating as a Service Principal using OpenID
// Connect.
func (o ProviderOutput) OidcRequestToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.OidcRequestToken }).(pulumi.StringPtrOutput)
}

// The URL for the OIDC provider from which to request an ID token. For use when authenticating as a Service Principal
// using OpenID Connect.
func (o ProviderOutput) OidcRequestUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.OidcRequestUrl }).(pulumi.StringPtrOutput)
}

// The ID token for use when authenticating as a Service Principal using OpenID Connect.
func (o ProviderOutput) OidcToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.OidcToken }).(pulumi.StringPtrOutput)
}

// The path to a file containing an ID token for use when authenticating as a Service Principal using OpenID Connect.
func (o ProviderOutput) OidcTokenFilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.OidcTokenFilePath }).(pulumi.StringPtrOutput)
}

// A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution
func (o ProviderOutput) PartnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.PartnerId }).(pulumi.StringPtrOutput)
}

// The Tenant ID which should be used. Works with all authentication methods except Managed Identity
func (o ProviderOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}

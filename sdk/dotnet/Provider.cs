// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD
{
    /// <summary>
    /// The provider type for the azuread package. By default, resources use package-wide configuration
    /// settings, however an explicit `Provider` instance may be created and passed during resource
    /// construction to achieve fine-grained programmatic control over provider settings. See the
    /// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
    /// </summary>
    [AzureADResourceType("pulumi:providers:azuread")]
    public partial class Provider : global::Pulumi.ProviderResource
    {
        /// <summary>
        /// Base64 encoded PKCS#12 certificate bundle to use when authenticating as a Service Principal using a Client Certificate
        /// </summary>
        [Output("clientCertificate")]
        public Output<string?> ClientCertificate { get; private set; } = null!;

        /// <summary>
        /// The password to decrypt the Client Certificate. For use when authenticating as a Service Principal using a Client
        /// Certificate
        /// </summary>
        [Output("clientCertificatePassword")]
        public Output<string?> ClientCertificatePassword { get; private set; } = null!;

        /// <summary>
        /// The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
        /// Principal using a Client Certificate
        /// </summary>
        [Output("clientCertificatePath")]
        public Output<string?> ClientCertificatePath { get; private set; } = null!;

        /// <summary>
        /// The Client ID which should be used for service principal authentication
        /// </summary>
        [Output("clientId")]
        public Output<string?> ClientId { get; private set; } = null!;

        /// <summary>
        /// The application password to use when authenticating as a Service Principal using a Client Secret
        /// </summary>
        [Output("clientSecret")]
        public Output<string?> ClientSecret { get; private set; } = null!;

        /// <summary>
        /// The cloud environment which should be used. Possible values are: `global` (also `public`), `usgovernmentl4` (also
        /// `usgovernment`), `usgovernmentl5` (also `dod`), and `china`. Defaults to `global`
        /// </summary>
        [Output("environment")]
        public Output<string?> Environment { get; private set; } = null!;

        /// <summary>
        /// The path to a custom endpoint for Managed Identity - in most circumstances this should be detected automatically
        /// </summary>
        [Output("msiEndpoint")]
        public Output<string?> MsiEndpoint { get; private set; } = null!;

        /// <summary>
        /// The bearer token for the request to the OIDC provider. For use when authenticating as a Service Principal using OpenID
        /// Connect.
        /// </summary>
        [Output("oidcRequestToken")]
        public Output<string?> OidcRequestToken { get; private set; } = null!;

        /// <summary>
        /// The URL for the OIDC provider from which to request an ID token. For use when authenticating as a Service Principal
        /// using OpenID Connect.
        /// </summary>
        [Output("oidcRequestUrl")]
        public Output<string?> OidcRequestUrl { get; private set; } = null!;

        /// <summary>
        /// The ID token for use when authenticating as a Service Principal using OpenID Connect.
        /// </summary>
        [Output("oidcToken")]
        public Output<string?> OidcToken { get; private set; } = null!;

        /// <summary>
        /// A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution
        /// </summary>
        [Output("partnerId")]
        public Output<string?> PartnerId { get; private set; } = null!;

        /// <summary>
        /// The Tenant ID which should be used. Works with all authentication methods except Managed Identity
        /// </summary>
        [Output("tenantId")]
        public Output<string?> TenantId { get; private set; } = null!;


        /// <summary>
        /// Create a Provider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Provider(string name, ProviderArgs? args = null, CustomResourceOptions? options = null)
            : base("azuread", name, args ?? new ProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class ProviderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Base64 encoded PKCS#12 certificate bundle to use when authenticating as a Service Principal using a Client Certificate
        /// </summary>
        [Input("clientCertificate")]
        public Input<string>? ClientCertificate { get; set; }

        /// <summary>
        /// The password to decrypt the Client Certificate. For use when authenticating as a Service Principal using a Client
        /// Certificate
        /// </summary>
        [Input("clientCertificatePassword")]
        public Input<string>? ClientCertificatePassword { get; set; }

        /// <summary>
        /// The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
        /// Principal using a Client Certificate
        /// </summary>
        [Input("clientCertificatePath")]
        public Input<string>? ClientCertificatePath { get; set; }

        /// <summary>
        /// The Client ID which should be used for service principal authentication
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// The application password to use when authenticating as a Service Principal using a Client Secret
        /// </summary>
        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

        /// <summary>
        /// Disable the Terraform Partner ID, which is used if a custom `partner_id` isn't specified
        /// </summary>
        [Input("disableTerraformPartnerId", json: true)]
        public Input<bool>? DisableTerraformPartnerId { get; set; }

        /// <summary>
        /// The cloud environment which should be used. Possible values are: `global` (also `public`), `usgovernmentl4` (also
        /// `usgovernment`), `usgovernmentl5` (also `dod`), and `china`. Defaults to `global`
        /// </summary>
        [Input("environment")]
        public Input<string>? Environment { get; set; }

        /// <summary>
        /// The path to a custom endpoint for Managed Identity - in most circumstances this should be detected automatically
        /// </summary>
        [Input("msiEndpoint")]
        public Input<string>? MsiEndpoint { get; set; }

        /// <summary>
        /// The bearer token for the request to the OIDC provider. For use when authenticating as a Service Principal using OpenID
        /// Connect.
        /// </summary>
        [Input("oidcRequestToken")]
        public Input<string>? OidcRequestToken { get; set; }

        /// <summary>
        /// The URL for the OIDC provider from which to request an ID token. For use when authenticating as a Service Principal
        /// using OpenID Connect.
        /// </summary>
        [Input("oidcRequestUrl")]
        public Input<string>? OidcRequestUrl { get; set; }

        /// <summary>
        /// The ID token for use when authenticating as a Service Principal using OpenID Connect.
        /// </summary>
        [Input("oidcToken")]
        public Input<string>? OidcToken { get; set; }

        /// <summary>
        /// A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution
        /// </summary>
        [Input("partnerId")]
        public Input<string>? PartnerId { get; set; }

        /// <summary>
        /// The Tenant ID which should be used. Works with all authentication methods except Managed Identity
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// Allow Azure CLI to be used for Authentication
        /// </summary>
        [Input("useCli", json: true)]
        public Input<bool>? UseCli { get; set; }

        /// <summary>
        /// Allow Managed Identity to be used for Authentication
        /// </summary>
        [Input("useMsi", json: true)]
        public Input<bool>? UseMsi { get; set; }

        /// <summary>
        /// Allow OpenID Connect to be used for authentication
        /// </summary>
        [Input("useOidc", json: true)]
        public Input<bool>? UseOidc { get; set; }

        public ProviderArgs()
        {
            Environment = Utilities.GetEnv("ARM_ENVIRONMENT") ?? "public";
            MsiEndpoint = Utilities.GetEnv("ARM_MSI_ENDPOINT");
            UseMsi = Utilities.GetEnvBoolean("ARM_USE_MSI") ?? false;
        }
        public static new ProviderArgs Empty => new ProviderArgs();
    }
}

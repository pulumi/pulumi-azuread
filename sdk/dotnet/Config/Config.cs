// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumi.AzureAD
{
    public static class Config
    {
        [System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly Pulumi.Config __config = new Pulumi.Config("azuread");

        private static readonly __Value<string?> _clientCertificate = new __Value<string?>(() => __config.Get("clientCertificate"));
        /// <summary>
        /// Base64 encoded PKCS#12 certificate bundle to use when authenticating as a Service Principal using a Client Certificate
        /// </summary>
        public static string? ClientCertificate
        {
            get => _clientCertificate.Get();
            set => _clientCertificate.Set(value);
        }

        private static readonly __Value<string?> _clientCertificatePassword = new __Value<string?>(() => __config.Get("clientCertificatePassword"));
        /// <summary>
        /// The password to decrypt the Client Certificate. For use when authenticating as a Service Principal using a Client
        /// Certificate
        /// </summary>
        public static string? ClientCertificatePassword
        {
            get => _clientCertificatePassword.Get();
            set => _clientCertificatePassword.Set(value);
        }

        private static readonly __Value<string?> _clientCertificatePath = new __Value<string?>(() => __config.Get("clientCertificatePath"));
        /// <summary>
        /// The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
        /// Principal using a Client Certificate
        /// </summary>
        public static string? ClientCertificatePath
        {
            get => _clientCertificatePath.Get();
            set => _clientCertificatePath.Set(value);
        }

        private static readonly __Value<string?> _clientId = new __Value<string?>(() => __config.Get("clientId"));
        /// <summary>
        /// The Client ID which should be used for service principal authentication
        /// </summary>
        public static string? ClientId
        {
            get => _clientId.Get();
            set => _clientId.Set(value);
        }

        private static readonly __Value<string?> _clientSecret = new __Value<string?>(() => __config.Get("clientSecret"));
        /// <summary>
        /// The application password to use when authenticating as a Service Principal using a Client Secret
        /// </summary>
        public static string? ClientSecret
        {
            get => _clientSecret.Get();
            set => _clientSecret.Set(value);
        }

        private static readonly __Value<bool?> _disableTerraformPartnerId = new __Value<bool?>(() => __config.GetBoolean("disableTerraformPartnerId"));
        /// <summary>
        /// Disable the Terraform Partner ID, which is used if a custom `partner_id` isn't specified
        /// </summary>
        public static bool? DisableTerraformPartnerId
        {
            get => _disableTerraformPartnerId.Get();
            set => _disableTerraformPartnerId.Set(value);
        }

        private static readonly __Value<string?> _environment = new __Value<string?>(() => __config.Get("environment") ?? Utilities.GetEnv("ARM_ENVIRONMENT") ?? "public");
        /// <summary>
        /// The cloud environment which should be used. Possible values are: `global` (also `public`), `usgovernmentl4` (also
        /// `usgovernment`), `usgovernmentl5` (also `dod`), and `china`. Defaults to `global`
        /// </summary>
        public static string? Environment
        {
            get => _environment.Get();
            set => _environment.Set(value);
        }

        private static readonly __Value<string?> _msiEndpoint = new __Value<string?>(() => __config.Get("msiEndpoint") ?? Utilities.GetEnv("ARM_MSI_ENDPOINT"));
        /// <summary>
        /// The path to a custom endpoint for Managed Identity - in most circumstances this should be detected automatically
        /// </summary>
        public static string? MsiEndpoint
        {
            get => _msiEndpoint.Get();
            set => _msiEndpoint.Set(value);
        }

        private static readonly __Value<string?> _oidcRequestToken = new __Value<string?>(() => __config.Get("oidcRequestToken"));
        /// <summary>
        /// The bearer token for the request to the OIDC provider. For use when authenticating as a Service Principal using OpenID
        /// Connect.
        /// </summary>
        public static string? OidcRequestToken
        {
            get => _oidcRequestToken.Get();
            set => _oidcRequestToken.Set(value);
        }

        private static readonly __Value<string?> _oidcRequestUrl = new __Value<string?>(() => __config.Get("oidcRequestUrl"));
        /// <summary>
        /// The URL for the OIDC provider from which to request an ID token. For use when authenticating as a Service Principal
        /// using OpenID Connect.
        /// </summary>
        public static string? OidcRequestUrl
        {
            get => _oidcRequestUrl.Get();
            set => _oidcRequestUrl.Set(value);
        }

        private static readonly __Value<string?> _partnerId = new __Value<string?>(() => __config.Get("partnerId"));
        /// <summary>
        /// A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution
        /// </summary>
        public static string? PartnerId
        {
            get => _partnerId.Get();
            set => _partnerId.Set(value);
        }

        private static readonly __Value<string?> _tenantId = new __Value<string?>(() => __config.Get("tenantId"));
        /// <summary>
        /// The Tenant ID which should be used. Works with all authentication methods except Managed Identity
        /// </summary>
        public static string? TenantId
        {
            get => _tenantId.Get();
            set => _tenantId.Set(value);
        }

        private static readonly __Value<bool?> _useCli = new __Value<bool?>(() => __config.GetBoolean("useCli"));
        /// <summary>
        /// Allow Azure CLI to be used for Authentication
        /// </summary>
        public static bool? UseCli
        {
            get => _useCli.Get();
            set => _useCli.Set(value);
        }

        private static readonly __Value<bool?> _useMsi = new __Value<bool?>(() => __config.GetBoolean("useMsi") ?? Utilities.GetEnvBoolean("ARM_USE_MSI") ?? false);
        /// <summary>
        /// Allow Managed Identity to be used for Authentication
        /// </summary>
        public static bool? UseMsi
        {
            get => _useMsi.Get();
            set => _useMsi.Set(value);
        }

        private static readonly __Value<bool?> _useOidc = new __Value<bool?>(() => __config.GetBoolean("useOidc"));
        /// <summary>
        /// Allow OpenID Connect to be used for authentication
        /// </summary>
        public static bool? UseOidc
        {
            get => _useOidc.Get();
            set => _useOidc.Set(value);
        }

    }
}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProviderArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProviderArgs Empty = new ProviderArgs();

    /**
     * Base64 encoded PKCS#12 certificate bundle to use when authenticating as a Service Principal using a Client Certificate
     * 
     */
    @Import(name="clientCertificate")
    private @Nullable Output<String> clientCertificate;

    /**
     * @return Base64 encoded PKCS#12 certificate bundle to use when authenticating as a Service Principal using a Client Certificate
     * 
     */
    public Optional<Output<String>> clientCertificate() {
        return Optional.ofNullable(this.clientCertificate);
    }

    /**
     * The password to decrypt the Client Certificate. For use when authenticating as a Service Principal using a Client
     * Certificate
     * 
     */
    @Import(name="clientCertificatePassword")
    private @Nullable Output<String> clientCertificatePassword;

    /**
     * @return The password to decrypt the Client Certificate. For use when authenticating as a Service Principal using a Client
     * Certificate
     * 
     */
    public Optional<Output<String>> clientCertificatePassword() {
        return Optional.ofNullable(this.clientCertificatePassword);
    }

    /**
     * The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
     * Principal using a Client Certificate
     * 
     */
    @Import(name="clientCertificatePath")
    private @Nullable Output<String> clientCertificatePath;

    /**
     * @return The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
     * Principal using a Client Certificate
     * 
     */
    public Optional<Output<String>> clientCertificatePath() {
        return Optional.ofNullable(this.clientCertificatePath);
    }

    /**
     * The Client ID which should be used for service principal authentication
     * 
     */
    @Import(name="clientId")
    private @Nullable Output<String> clientId;

    /**
     * @return The Client ID which should be used for service principal authentication
     * 
     */
    public Optional<Output<String>> clientId() {
        return Optional.ofNullable(this.clientId);
    }

    /**
     * The application password to use when authenticating as a Service Principal using a Client Secret
     * 
     */
    @Import(name="clientSecret")
    private @Nullable Output<String> clientSecret;

    /**
     * @return The application password to use when authenticating as a Service Principal using a Client Secret
     * 
     */
    public Optional<Output<String>> clientSecret() {
        return Optional.ofNullable(this.clientSecret);
    }

    /**
     * Disable the Terraform Partner ID, which is used if a custom `partner_id` isn&#39;t specified
     * 
     */
    @Import(name="disableTerraformPartnerId", json=true)
    private @Nullable Output<Boolean> disableTerraformPartnerId;

    /**
     * @return Disable the Terraform Partner ID, which is used if a custom `partner_id` isn&#39;t specified
     * 
     */
    public Optional<Output<Boolean>> disableTerraformPartnerId() {
        return Optional.ofNullable(this.disableTerraformPartnerId);
    }

    /**
     * The cloud environment which should be used. Possible values are: `global` (also `public`), `usgovernmentl4` (also
     * `usgovernment`), `usgovernmentl5` (also `dod`), and `china`. Defaults to `global`
     * 
     */
    @Import(name="environment")
    private @Nullable Output<String> environment;

    /**
     * @return The cloud environment which should be used. Possible values are: `global` (also `public`), `usgovernmentl4` (also
     * `usgovernment`), `usgovernmentl5` (also `dod`), and `china`. Defaults to `global`
     * 
     */
    public Optional<Output<String>> environment() {
        return Optional.ofNullable(this.environment);
    }

    /**
     * The path to a custom endpoint for Managed Identity - in most circumstances this should be detected automatically
     * 
     */
    @Import(name="msiEndpoint")
    private @Nullable Output<String> msiEndpoint;

    /**
     * @return The path to a custom endpoint for Managed Identity - in most circumstances this should be detected automatically
     * 
     */
    public Optional<Output<String>> msiEndpoint() {
        return Optional.ofNullable(this.msiEndpoint);
    }

    /**
     * The bearer token for the request to the OIDC provider. For use when authenticating as a Service Principal using OpenID
     * Connect.
     * 
     */
    @Import(name="oidcRequestToken")
    private @Nullable Output<String> oidcRequestToken;

    /**
     * @return The bearer token for the request to the OIDC provider. For use when authenticating as a Service Principal using OpenID
     * Connect.
     * 
     */
    public Optional<Output<String>> oidcRequestToken() {
        return Optional.ofNullable(this.oidcRequestToken);
    }

    /**
     * The URL for the OIDC provider from which to request an ID token. For use when authenticating as a Service Principal
     * using OpenID Connect.
     * 
     */
    @Import(name="oidcRequestUrl")
    private @Nullable Output<String> oidcRequestUrl;

    /**
     * @return The URL for the OIDC provider from which to request an ID token. For use when authenticating as a Service Principal
     * using OpenID Connect.
     * 
     */
    public Optional<Output<String>> oidcRequestUrl() {
        return Optional.ofNullable(this.oidcRequestUrl);
    }

    /**
     * A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution
     * 
     */
    @Import(name="partnerId")
    private @Nullable Output<String> partnerId;

    /**
     * @return A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution
     * 
     */
    public Optional<Output<String>> partnerId() {
        return Optional.ofNullable(this.partnerId);
    }

    /**
     * The Tenant ID which should be used. Works with all authentication methods except Managed Identity
     * 
     */
    @Import(name="tenantId")
    private @Nullable Output<String> tenantId;

    /**
     * @return The Tenant ID which should be used. Works with all authentication methods except Managed Identity
     * 
     */
    public Optional<Output<String>> tenantId() {
        return Optional.ofNullable(this.tenantId);
    }

    /**
     * Allow Azure CLI to be used for Authentication
     * 
     */
    @Import(name="useCli", json=true)
    private @Nullable Output<Boolean> useCli;

    /**
     * @return Allow Azure CLI to be used for Authentication
     * 
     */
    public Optional<Output<Boolean>> useCli() {
        return Optional.ofNullable(this.useCli);
    }

    /**
     * Allow Managed Identity to be used for Authentication
     * 
     */
    @Import(name="useMsi", json=true)
    private @Nullable Output<Boolean> useMsi;

    /**
     * @return Allow Managed Identity to be used for Authentication
     * 
     */
    public Optional<Output<Boolean>> useMsi() {
        return Optional.ofNullable(this.useMsi);
    }

    /**
     * Allow OpenID Connect to be used for authentication
     * 
     */
    @Import(name="useOidc", json=true)
    private @Nullable Output<Boolean> useOidc;

    /**
     * @return Allow OpenID Connect to be used for authentication
     * 
     */
    public Optional<Output<Boolean>> useOidc() {
        return Optional.ofNullable(this.useOidc);
    }

    private ProviderArgs() {}

    private ProviderArgs(ProviderArgs $) {
        this.clientCertificate = $.clientCertificate;
        this.clientCertificatePassword = $.clientCertificatePassword;
        this.clientCertificatePath = $.clientCertificatePath;
        this.clientId = $.clientId;
        this.clientSecret = $.clientSecret;
        this.disableTerraformPartnerId = $.disableTerraformPartnerId;
        this.environment = $.environment;
        this.msiEndpoint = $.msiEndpoint;
        this.oidcRequestToken = $.oidcRequestToken;
        this.oidcRequestUrl = $.oidcRequestUrl;
        this.partnerId = $.partnerId;
        this.tenantId = $.tenantId;
        this.useCli = $.useCli;
        this.useMsi = $.useMsi;
        this.useOidc = $.useOidc;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProviderArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProviderArgs $;

        public Builder() {
            $ = new ProviderArgs();
        }

        public Builder(ProviderArgs defaults) {
            $ = new ProviderArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clientCertificate Base64 encoded PKCS#12 certificate bundle to use when authenticating as a Service Principal using a Client Certificate
         * 
         * @return builder
         * 
         */
        public Builder clientCertificate(@Nullable Output<String> clientCertificate) {
            $.clientCertificate = clientCertificate;
            return this;
        }

        /**
         * @param clientCertificate Base64 encoded PKCS#12 certificate bundle to use when authenticating as a Service Principal using a Client Certificate
         * 
         * @return builder
         * 
         */
        public Builder clientCertificate(String clientCertificate) {
            return clientCertificate(Output.of(clientCertificate));
        }

        /**
         * @param clientCertificatePassword The password to decrypt the Client Certificate. For use when authenticating as a Service Principal using a Client
         * Certificate
         * 
         * @return builder
         * 
         */
        public Builder clientCertificatePassword(@Nullable Output<String> clientCertificatePassword) {
            $.clientCertificatePassword = clientCertificatePassword;
            return this;
        }

        /**
         * @param clientCertificatePassword The password to decrypt the Client Certificate. For use when authenticating as a Service Principal using a Client
         * Certificate
         * 
         * @return builder
         * 
         */
        public Builder clientCertificatePassword(String clientCertificatePassword) {
            return clientCertificatePassword(Output.of(clientCertificatePassword));
        }

        /**
         * @param clientCertificatePath The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
         * Principal using a Client Certificate
         * 
         * @return builder
         * 
         */
        public Builder clientCertificatePath(@Nullable Output<String> clientCertificatePath) {
            $.clientCertificatePath = clientCertificatePath;
            return this;
        }

        /**
         * @param clientCertificatePath The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
         * Principal using a Client Certificate
         * 
         * @return builder
         * 
         */
        public Builder clientCertificatePath(String clientCertificatePath) {
            return clientCertificatePath(Output.of(clientCertificatePath));
        }

        /**
         * @param clientId The Client ID which should be used for service principal authentication
         * 
         * @return builder
         * 
         */
        public Builder clientId(@Nullable Output<String> clientId) {
            $.clientId = clientId;
            return this;
        }

        /**
         * @param clientId The Client ID which should be used for service principal authentication
         * 
         * @return builder
         * 
         */
        public Builder clientId(String clientId) {
            return clientId(Output.of(clientId));
        }

        /**
         * @param clientSecret The application password to use when authenticating as a Service Principal using a Client Secret
         * 
         * @return builder
         * 
         */
        public Builder clientSecret(@Nullable Output<String> clientSecret) {
            $.clientSecret = clientSecret;
            return this;
        }

        /**
         * @param clientSecret The application password to use when authenticating as a Service Principal using a Client Secret
         * 
         * @return builder
         * 
         */
        public Builder clientSecret(String clientSecret) {
            return clientSecret(Output.of(clientSecret));
        }

        /**
         * @param disableTerraformPartnerId Disable the Terraform Partner ID, which is used if a custom `partner_id` isn&#39;t specified
         * 
         * @return builder
         * 
         */
        public Builder disableTerraformPartnerId(@Nullable Output<Boolean> disableTerraformPartnerId) {
            $.disableTerraformPartnerId = disableTerraformPartnerId;
            return this;
        }

        /**
         * @param disableTerraformPartnerId Disable the Terraform Partner ID, which is used if a custom `partner_id` isn&#39;t specified
         * 
         * @return builder
         * 
         */
        public Builder disableTerraformPartnerId(Boolean disableTerraformPartnerId) {
            return disableTerraformPartnerId(Output.of(disableTerraformPartnerId));
        }

        /**
         * @param environment The cloud environment which should be used. Possible values are: `global` (also `public`), `usgovernmentl4` (also
         * `usgovernment`), `usgovernmentl5` (also `dod`), and `china`. Defaults to `global`
         * 
         * @return builder
         * 
         */
        public Builder environment(@Nullable Output<String> environment) {
            $.environment = environment;
            return this;
        }

        /**
         * @param environment The cloud environment which should be used. Possible values are: `global` (also `public`), `usgovernmentl4` (also
         * `usgovernment`), `usgovernmentl5` (also `dod`), and `china`. Defaults to `global`
         * 
         * @return builder
         * 
         */
        public Builder environment(String environment) {
            return environment(Output.of(environment));
        }

        /**
         * @param msiEndpoint The path to a custom endpoint for Managed Identity - in most circumstances this should be detected automatically
         * 
         * @return builder
         * 
         */
        public Builder msiEndpoint(@Nullable Output<String> msiEndpoint) {
            $.msiEndpoint = msiEndpoint;
            return this;
        }

        /**
         * @param msiEndpoint The path to a custom endpoint for Managed Identity - in most circumstances this should be detected automatically
         * 
         * @return builder
         * 
         */
        public Builder msiEndpoint(String msiEndpoint) {
            return msiEndpoint(Output.of(msiEndpoint));
        }

        /**
         * @param oidcRequestToken The bearer token for the request to the OIDC provider. For use when authenticating as a Service Principal using OpenID
         * Connect.
         * 
         * @return builder
         * 
         */
        public Builder oidcRequestToken(@Nullable Output<String> oidcRequestToken) {
            $.oidcRequestToken = oidcRequestToken;
            return this;
        }

        /**
         * @param oidcRequestToken The bearer token for the request to the OIDC provider. For use when authenticating as a Service Principal using OpenID
         * Connect.
         * 
         * @return builder
         * 
         */
        public Builder oidcRequestToken(String oidcRequestToken) {
            return oidcRequestToken(Output.of(oidcRequestToken));
        }

        /**
         * @param oidcRequestUrl The URL for the OIDC provider from which to request an ID token. For use when authenticating as a Service Principal
         * using OpenID Connect.
         * 
         * @return builder
         * 
         */
        public Builder oidcRequestUrl(@Nullable Output<String> oidcRequestUrl) {
            $.oidcRequestUrl = oidcRequestUrl;
            return this;
        }

        /**
         * @param oidcRequestUrl The URL for the OIDC provider from which to request an ID token. For use when authenticating as a Service Principal
         * using OpenID Connect.
         * 
         * @return builder
         * 
         */
        public Builder oidcRequestUrl(String oidcRequestUrl) {
            return oidcRequestUrl(Output.of(oidcRequestUrl));
        }

        /**
         * @param partnerId A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution
         * 
         * @return builder
         * 
         */
        public Builder partnerId(@Nullable Output<String> partnerId) {
            $.partnerId = partnerId;
            return this;
        }

        /**
         * @param partnerId A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution
         * 
         * @return builder
         * 
         */
        public Builder partnerId(String partnerId) {
            return partnerId(Output.of(partnerId));
        }

        /**
         * @param tenantId The Tenant ID which should be used. Works with all authentication methods except Managed Identity
         * 
         * @return builder
         * 
         */
        public Builder tenantId(@Nullable Output<String> tenantId) {
            $.tenantId = tenantId;
            return this;
        }

        /**
         * @param tenantId The Tenant ID which should be used. Works with all authentication methods except Managed Identity
         * 
         * @return builder
         * 
         */
        public Builder tenantId(String tenantId) {
            return tenantId(Output.of(tenantId));
        }

        /**
         * @param useCli Allow Azure CLI to be used for Authentication
         * 
         * @return builder
         * 
         */
        public Builder useCli(@Nullable Output<Boolean> useCli) {
            $.useCli = useCli;
            return this;
        }

        /**
         * @param useCli Allow Azure CLI to be used for Authentication
         * 
         * @return builder
         * 
         */
        public Builder useCli(Boolean useCli) {
            return useCli(Output.of(useCli));
        }

        /**
         * @param useMsi Allow Managed Identity to be used for Authentication
         * 
         * @return builder
         * 
         */
        public Builder useMsi(@Nullable Output<Boolean> useMsi) {
            $.useMsi = useMsi;
            return this;
        }

        /**
         * @param useMsi Allow Managed Identity to be used for Authentication
         * 
         * @return builder
         * 
         */
        public Builder useMsi(Boolean useMsi) {
            return useMsi(Output.of(useMsi));
        }

        /**
         * @param useOidc Allow OpenID Connect to be used for authentication
         * 
         * @return builder
         * 
         */
        public Builder useOidc(@Nullable Output<Boolean> useOidc) {
            $.useOidc = useOidc;
            return this;
        }

        /**
         * @param useOidc Allow OpenID Connect to be used for authentication
         * 
         * @return builder
         * 
         */
        public Builder useOidc(Boolean useOidc) {
            return useOidc(Output.of(useOidc));
        }

        public ProviderArgs build() {
            $.environment = Codegen.stringProp("environment").output().arg($.environment).env("ARM_ENVIRONMENT").def("public").getNullable();
            $.msiEndpoint = Codegen.stringProp("msiEndpoint").output().arg($.msiEndpoint).env("ARM_MSI_ENDPOINT").getNullable();
            $.useMsi = Codegen.booleanProp("useMsi").output().arg($.useMsi).env("ARM_USE_MSI").def(false).getNullable();
            return $;
        }
    }

}

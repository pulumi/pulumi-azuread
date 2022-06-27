// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.ProviderArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The provider type for the azuread package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 * 
 */
@ResourceType(type="pulumi:providers:azuread")
public class Provider extends com.pulumi.resources.ProviderResource {
    /**
     * Base64 encoded PKCS#12 certificate bundle to use when authenticating as a Service Principal using a Client Certificate
     * 
     */
    @Export(name="clientCertificate", type=String.class, parameters={})
    private Output</* @Nullable */ String> clientCertificate;

    /**
     * @return Base64 encoded PKCS#12 certificate bundle to use when authenticating as a Service Principal using a Client Certificate
     * 
     */
    public Output<Optional<String>> clientCertificate() {
        return Codegen.optional(this.clientCertificate);
    }
    /**
     * The password to decrypt the Client Certificate. For use when authenticating as a Service Principal using a Client
     * Certificate
     * 
     */
    @Export(name="clientCertificatePassword", type=String.class, parameters={})
    private Output</* @Nullable */ String> clientCertificatePassword;

    /**
     * @return The password to decrypt the Client Certificate. For use when authenticating as a Service Principal using a Client
     * Certificate
     * 
     */
    public Output<Optional<String>> clientCertificatePassword() {
        return Codegen.optional(this.clientCertificatePassword);
    }
    /**
     * The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
     * Principal using a Client Certificate
     * 
     */
    @Export(name="clientCertificatePath", type=String.class, parameters={})
    private Output</* @Nullable */ String> clientCertificatePath;

    /**
     * @return The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
     * Principal using a Client Certificate
     * 
     */
    public Output<Optional<String>> clientCertificatePath() {
        return Codegen.optional(this.clientCertificatePath);
    }
    /**
     * The Client ID which should be used for service principal authentication
     * 
     */
    @Export(name="clientId", type=String.class, parameters={})
    private Output</* @Nullable */ String> clientId;

    /**
     * @return The Client ID which should be used for service principal authentication
     * 
     */
    public Output<Optional<String>> clientId() {
        return Codegen.optional(this.clientId);
    }
    /**
     * The application password to use when authenticating as a Service Principal using a Client Secret
     * 
     */
    @Export(name="clientSecret", type=String.class, parameters={})
    private Output</* @Nullable */ String> clientSecret;

    /**
     * @return The application password to use when authenticating as a Service Principal using a Client Secret
     * 
     */
    public Output<Optional<String>> clientSecret() {
        return Codegen.optional(this.clientSecret);
    }
    /**
     * The cloud environment which should be used. Possible values are: `global` (also `public`), `usgovernmentl4` (also
     * `usgovernment`), `usgovernmentl5` (also `dod`), and `china`. Defaults to `global`
     * 
     */
    @Export(name="environment", type=String.class, parameters={})
    private Output</* @Nullable */ String> environment;

    /**
     * @return The cloud environment which should be used. Possible values are: `global` (also `public`), `usgovernmentl4` (also
     * `usgovernment`), `usgovernmentl5` (also `dod`), and `china`. Defaults to `global`
     * 
     */
    public Output<Optional<String>> environment() {
        return Codegen.optional(this.environment);
    }
    /**
     * The path to a custom endpoint for Managed Identity - in most circumstances this should be detected automatically
     * 
     */
    @Export(name="msiEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> msiEndpoint;

    /**
     * @return The path to a custom endpoint for Managed Identity - in most circumstances this should be detected automatically
     * 
     */
    public Output<Optional<String>> msiEndpoint() {
        return Codegen.optional(this.msiEndpoint);
    }
    /**
     * The bearer token for the request to the OIDC provider. For use when authenticating as a Service Principal using OpenID
     * Connect.
     * 
     */
    @Export(name="oidcRequestToken", type=String.class, parameters={})
    private Output</* @Nullable */ String> oidcRequestToken;

    /**
     * @return The bearer token for the request to the OIDC provider. For use when authenticating as a Service Principal using OpenID
     * Connect.
     * 
     */
    public Output<Optional<String>> oidcRequestToken() {
        return Codegen.optional(this.oidcRequestToken);
    }
    /**
     * The URL for the OIDC provider from which to request an ID token. For use when authenticating as a Service Principal
     * using OpenID Connect.
     * 
     */
    @Export(name="oidcRequestUrl", type=String.class, parameters={})
    private Output</* @Nullable */ String> oidcRequestUrl;

    /**
     * @return The URL for the OIDC provider from which to request an ID token. For use when authenticating as a Service Principal
     * using OpenID Connect.
     * 
     */
    public Output<Optional<String>> oidcRequestUrl() {
        return Codegen.optional(this.oidcRequestUrl);
    }
    /**
     * A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution
     * 
     */
    @Export(name="partnerId", type=String.class, parameters={})
    private Output</* @Nullable */ String> partnerId;

    /**
     * @return A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution
     * 
     */
    public Output<Optional<String>> partnerId() {
        return Codegen.optional(this.partnerId);
    }
    /**
     * The Tenant ID which should be used. Works with all authentication methods except Managed Identity
     * 
     */
    @Export(name="tenantId", type=String.class, parameters={})
    private Output</* @Nullable */ String> tenantId;

    /**
     * @return The Tenant ID which should be used. Works with all authentication methods except Managed Identity
     * 
     */
    public Output<Optional<String>> tenantId() {
        return Codegen.optional(this.tenantId);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Provider(String name) {
        this(name, ProviderArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Provider(String name, @Nullable ProviderArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Provider(String name, @Nullable ProviderArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread", name, args == null ? ProviderArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

}

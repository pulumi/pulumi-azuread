// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.ServicePrincipalTokenSigningCertificateArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.ServicePrincipalTokenSigningCertificateState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * *Using default settings*
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuread.Application;
 * import com.pulumi.azuread.ApplicationArgs;
 * import com.pulumi.azuread.ServicePrincipal;
 * import com.pulumi.azuread.ServicePrincipalArgs;
 * import com.pulumi.azuread.ServicePrincipalTokenSigningCertificate;
 * import com.pulumi.azuread.ServicePrincipalTokenSigningCertificateArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var example = new Application(&#34;example&#34;, ApplicationArgs.builder()        
 *             .displayName(&#34;example&#34;)
 *             .build());
 * 
 *         var exampleServicePrincipal = new ServicePrincipal(&#34;exampleServicePrincipal&#34;, ServicePrincipalArgs.builder()        
 *             .applicationId(example.applicationId())
 *             .build());
 * 
 *         var exampleServicePrincipalTokenSigningCertificate = new ServicePrincipalTokenSigningCertificate(&#34;exampleServicePrincipalTokenSigningCertificate&#34;, ServicePrincipalTokenSigningCertificateArgs.builder()        
 *             .servicePrincipalId(exampleServicePrincipal.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * *Using custom settings*
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuread.Application;
 * import com.pulumi.azuread.ApplicationArgs;
 * import com.pulumi.azuread.ServicePrincipal;
 * import com.pulumi.azuread.ServicePrincipalArgs;
 * import com.pulumi.azuread.ServicePrincipalTokenSigningCertificate;
 * import com.pulumi.azuread.ServicePrincipalTokenSigningCertificateArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var example = new Application(&#34;example&#34;, ApplicationArgs.builder()        
 *             .displayName(&#34;example&#34;)
 *             .build());
 * 
 *         var exampleServicePrincipal = new ServicePrincipal(&#34;exampleServicePrincipal&#34;, ServicePrincipalArgs.builder()        
 *             .applicationId(example.applicationId())
 *             .build());
 * 
 *         var exampleServicePrincipalTokenSigningCertificate = new ServicePrincipalTokenSigningCertificate(&#34;exampleServicePrincipalTokenSigningCertificate&#34;, ServicePrincipalTokenSigningCertificateArgs.builder()        
 *             .servicePrincipalId(exampleServicePrincipal.id())
 *             .displayName(&#34;CN=example.com&#34;)
 *             .endDate(&#34;2023-05-01T01:02:03Z&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Token signing certificates can be imported using the object ID of the associated service principal and the key ID of the verify certificate credential, e.g.
 * 
 * ```sh
 * $ pulumi import azuread:index/servicePrincipalTokenSigningCertificate:ServicePrincipalTokenSigningCertificate example 00000000-0000-0000-0000-000000000000/tokenSigningCertificate/11111111-1111-1111-1111-111111111111
 * ```
 * 
 * -&gt; This ID format is unique to Terraform and is composed of the service principal&#39;s object ID, the string &#34;tokenSigningCertificate&#34; and the verify certificate&#39;s key ID in the format `{ServicePrincipalObjectId}/tokenSigningCertificate/{CertificateKeyId}`.
 * 
 */
@ResourceType(type="azuread:index/servicePrincipalTokenSigningCertificate:ServicePrincipalTokenSigningCertificate")
public class ServicePrincipalTokenSigningCertificate extends com.pulumi.resources.CustomResource {
    /**
     * Specifies a friendly name for the certificate.
     * Must start with `CN=`. Changing this field forces a new resource to be created.
     * 
     * &gt; If not specified, it will default to `CN=Microsoft Azure Federated SSO Certificate`.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    /**
     * @return Specifies a friendly name for the certificate.
     * Must start with `CN=`. Changing this field forces a new resource to be created.
     * 
     * &gt; If not specified, it will default to `CN=Microsoft Azure Federated SSO Certificate`.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * The end date until which the token signing certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
     * 
     */
    @Export(name="endDate", refs={String.class}, tree="[0]")
    private Output<String> endDate;

    /**
     * @return The end date until which the token signing certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
     * 
     */
    public Output<String> endDate() {
        return this.endDate;
    }
    /**
     * A UUID used to uniquely identify the verify certificate.
     * 
     */
    @Export(name="keyId", refs={String.class}, tree="[0]")
    private Output<String> keyId;

    /**
     * @return A UUID used to uniquely identify the verify certificate.
     * 
     */
    public Output<String> keyId() {
        return this.keyId;
    }
    /**
     * The object ID of the service principal for which this certificate should be created. Changing this field forces a new resource to be created.
     * 
     */
    @Export(name="servicePrincipalId", refs={String.class}, tree="[0]")
    private Output<String> servicePrincipalId;

    /**
     * @return The object ID of the service principal for which this certificate should be created. Changing this field forces a new resource to be created.
     * 
     */
    public Output<String> servicePrincipalId() {
        return this.servicePrincipalId;
    }
    /**
     * The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
     * 
     */
    @Export(name="startDate", refs={String.class}, tree="[0]")
    private Output<String> startDate;

    /**
     * @return The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
     * 
     */
    public Output<String> startDate() {
        return this.startDate;
    }
    /**
     * A SHA-1 generated thumbprint of the token signing certificate, which can be used to set the preferred signing certificate for a service principal.
     * 
     */
    @Export(name="thumbprint", refs={String.class}, tree="[0]")
    private Output<String> thumbprint;

    /**
     * @return A SHA-1 generated thumbprint of the token signing certificate, which can be used to set the preferred signing certificate for a service principal.
     * 
     */
    public Output<String> thumbprint() {
        return this.thumbprint;
    }
    /**
     * The certificate data, which is PEM encoded but does not include the
     * header `-----BEGIN CERTIFICATE-----\n` or the footer `\n-----END CERTIFICATE-----`.
     * 
     */
    @Export(name="value", refs={String.class}, tree="[0]")
    private Output<String> value;

    /**
     * @return The certificate data, which is PEM encoded but does not include the
     * header `-----BEGIN CERTIFICATE-----\n` or the footer `\n-----END CERTIFICATE-----`.
     * 
     */
    public Output<String> value() {
        return this.value;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServicePrincipalTokenSigningCertificate(String name) {
        this(name, ServicePrincipalTokenSigningCertificateArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServicePrincipalTokenSigningCertificate(String name, ServicePrincipalTokenSigningCertificateArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServicePrincipalTokenSigningCertificate(String name, ServicePrincipalTokenSigningCertificateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/servicePrincipalTokenSigningCertificate:ServicePrincipalTokenSigningCertificate", name, args == null ? ServicePrincipalTokenSigningCertificateArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ServicePrincipalTokenSigningCertificate(String name, Output<String> id, @Nullable ServicePrincipalTokenSigningCertificateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/servicePrincipalTokenSigningCertificate:ServicePrincipalTokenSigningCertificate", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "value"
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static ServicePrincipalTokenSigningCertificate get(String name, Output<String> id, @Nullable ServicePrincipalTokenSigningCertificateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServicePrincipalTokenSigningCertificate(name, id, state, options);
    }
}

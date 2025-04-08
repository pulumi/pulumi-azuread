// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.ApplicationCertificateArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.ApplicationCertificateState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * *Using a PEM certificate*
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuread.ApplicationRegistration;
 * import com.pulumi.azuread.ApplicationRegistrationArgs;
 * import com.pulumi.azuread.ApplicationCertificate;
 * import com.pulumi.azuread.ApplicationCertificateArgs;
 * import com.pulumi.std.StdFunctions;
 * import com.pulumi.std.inputs.FileArgs;
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
 *         var example = new ApplicationRegistration("example", ApplicationRegistrationArgs.builder()
 *             .displayName("example")
 *             .build());
 * 
 *         var exampleApplicationCertificate = new ApplicationCertificate("exampleApplicationCertificate", ApplicationCertificateArgs.builder()
 *             .applicationId(example.id())
 *             .type("AsymmetricX509Cert")
 *             .value(StdFunctions.file(FileArgs.builder()
 *                 .input("cert.pem")
 *                 .build()).result())
 *             .endDate("2021-05-01T01:02:03Z")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * *Using a DER certificate*
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuread.ApplicationRegistration;
 * import com.pulumi.azuread.ApplicationRegistrationArgs;
 * import com.pulumi.azuread.ApplicationCertificate;
 * import com.pulumi.azuread.ApplicationCertificateArgs;
 * import com.pulumi.std.StdFunctions;
 * import com.pulumi.std.inputs.FileArgs;
 * import com.pulumi.std.inputs.Base64encodeArgs;
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
 *         var example = new ApplicationRegistration("example", ApplicationRegistrationArgs.builder()
 *             .displayName("example")
 *             .build());
 * 
 *         var exampleApplicationCertificate = new ApplicationCertificate("exampleApplicationCertificate", ApplicationCertificateArgs.builder()
 *             .applicationId(example.id())
 *             .type("AsymmetricX509Cert")
 *             .encoding("base64")
 *             .value(StdFunctions.base64encode(Base64encodeArgs.builder()
 *                 .input(StdFunctions.file(FileArgs.builder()
 *                     .input("cert.der")
 *                     .build()).result())
 *                 .build()).result())
 *             .endDate("2021-05-01T01:02:03Z")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Using a certificate from Azure Key Vault
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuread.Application;
 * import com.pulumi.azuread.ApplicationArgs;
 * import com.pulumi.azure.keyvault.Certificate;
 * import com.pulumi.azure.keyvault.CertificateArgs;
 * import com.pulumi.azure.keyvault.inputs.CertificateCertificatePolicyArgs;
 * import com.pulumi.azure.keyvault.inputs.CertificateCertificatePolicyIssuerParametersArgs;
 * import com.pulumi.azure.keyvault.inputs.CertificateCertificatePolicyKeyPropertiesArgs;
 * import com.pulumi.azure.keyvault.inputs.CertificateCertificatePolicySecretPropertiesArgs;
 * import com.pulumi.azure.keyvault.inputs.CertificateCertificatePolicyX509CertificatePropertiesArgs;
 * import com.pulumi.azure.keyvault.inputs.CertificateCertificatePolicyX509CertificatePropertiesSubjectAlternativeNamesArgs;
 * import com.pulumi.azuread.ApplicationCertificate;
 * import com.pulumi.azuread.ApplicationCertificateArgs;
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
 *         var exampleApplication = new Application("exampleApplication", ApplicationArgs.builder()
 *             .displayName("example")
 *             .build());
 * 
 *         var example = new Certificate("example", CertificateArgs.builder()
 *             .name("generated-cert")
 *             .keyVaultId(exampleAzurermKeyVault.id())
 *             .certificatePolicy(CertificateCertificatePolicyArgs.builder()
 *                 .issuerParameters(CertificateCertificatePolicyIssuerParametersArgs.builder()
 *                     .name("Self")
 *                     .build())
 *                 .keyProperties(CertificateCertificatePolicyKeyPropertiesArgs.builder()
 *                     .exportable(true)
 *                     .keySize(2048)
 *                     .keyType("RSA")
 *                     .reuseKey(true)
 *                     .build())
 *                 .lifetimeActions(CertificateCertificatePolicyLifetimeActionArgs.builder()
 *                     .action(CertificateCertificatePolicyLifetimeActionActionArgs.builder()
 *                         .actionType("AutoRenew")
 *                         .build())
 *                     .trigger(CertificateCertificatePolicyLifetimeActionTriggerArgs.builder()
 *                         .daysBeforeExpiry(30)
 *                         .build())
 *                     .build())
 *                 .secretProperties(CertificateCertificatePolicySecretPropertiesArgs.builder()
 *                     .contentType("application/x-pkcs12")
 *                     .build())
 *                 .x509CertificateProperties(CertificateCertificatePolicyX509CertificatePropertiesArgs.builder()
 *                     .extendedKeyUsages("1.3.6.1.5.5.7.3.2")
 *                     .keyUsages(                    
 *                         "dataEncipherment",
 *                         "digitalSignature",
 *                         "keyCertSign",
 *                         "keyEncipherment")
 *                     .subjectAlternativeNames(CertificateCertificatePolicyX509CertificatePropertiesSubjectAlternativeNamesArgs.builder()
 *                         .dnsNames(                        
 *                             "internal.contoso.com",
 *                             "domain.hello.world")
 *                         .build())
 *                     .subject(String.format("CN=%s", exampleApplication.name()))
 *                     .validityInMonths(12)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var exampleApplicationCertificate = new ApplicationCertificate("exampleApplicationCertificate", ApplicationCertificateArgs.builder()
 *             .applicationId(exampleApplication.id())
 *             .type("AsymmetricX509Cert")
 *             .encoding("hex")
 *             .value(example.certificateData())
 *             .endDate(example.certificateAttributes().applyValue(_certificateAttributes -> _certificateAttributes[0].expires()))
 *             .startDate(example.certificateAttributes().applyValue(_certificateAttributes -> _certificateAttributes[0].notBefore()))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Certificates can be imported using the object ID of the associated application and the key ID of the certificate credential, e.g.
 * 
 * ```sh
 * $ pulumi import azuread:index/applicationCertificate:ApplicationCertificate example 00000000-0000-0000-0000-000000000000/certificate/11111111-1111-1111-1111-111111111111
 * ```
 * 
 * -&gt; This ID format is unique to Terraform and is composed of the application&#39;s object ID, the string &#34;certificate&#34; and the certificate&#39;s key ID in the format `{ObjectId}/certificate/{CertificateKeyId}`.
 * 
 */
@ResourceType(type="azuread:index/applicationCertificate:ApplicationCertificate")
public class ApplicationCertificate extends com.pulumi.resources.CustomResource {
    /**
     * The resource ID of the application for which this certificate should be created. Changing this field forces a new resource to be created.
     * 
     */
    @Export(name="applicationId", refs={String.class}, tree="[0]")
    private Output<String> applicationId;

    /**
     * @return The resource ID of the application for which this certificate should be created. Changing this field forces a new resource to be created.
     * 
     */
    public Output<String> applicationId() {
        return this.applicationId;
    }
    /**
     * Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.
     * 
     * &gt; **Tip for Azure Key Vault** The `hex` encoding option is useful for consuming certificate data from the azurerm_key_vault_certificate resource.
     * 
     */
    @Export(name="encoding", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> encoding;

    /**
     * @return Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.
     * 
     * &gt; **Tip for Azure Key Vault** The `hex` encoding option is useful for consuming certificate data from the azurerm_key_vault_certificate resource.
     * 
     */
    public Output<Optional<String>> encoding() {
        return Codegen.optional(this.encoding);
    }
    /**
     * The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If omitted, the API will decide a suitable expiry date, which is typically around 2 years from the start date. Changing this field forces a new resource to be created.
     * 
     */
    @Export(name="endDate", refs={String.class}, tree="[0]")
    private Output<String> endDate;

    /**
     * @return The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If omitted, the API will decide a suitable expiry date, which is typically around 2 years from the start date. Changing this field forces a new resource to be created.
     * 
     */
    public Output<String> endDate() {
        return this.endDate;
    }
    /**
     * A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
     * 
     * &gt; One of `end_date` or `end_date_relative` must be specified. The maximum allowed duration is determined by Azure AD and is typically around 2 years from the creation date.
     * 
     * @deprecated
     * The `end_date_relative` property is deprecated and will be removed in a future version of the AzureAD provider. Please instead use the Terraform `timeadd()` function to calculate a value for the `end_date` property.
     * 
     */
    @Deprecated /* The `end_date_relative` property is deprecated and will be removed in a future version of the AzureAD provider. Please instead use the Terraform `timeadd()` function to calculate a value for the `end_date` property. */
    @Export(name="endDateRelative", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> endDateRelative;

    /**
     * @return A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
     * 
     * &gt; One of `end_date` or `end_date_relative` must be specified. The maximum allowed duration is determined by Azure AD and is typically around 2 years from the creation date.
     * 
     */
    public Output<Optional<String>> endDateRelative() {
        return Codegen.optional(this.endDateRelative);
    }
    /**
     * A UUID used to uniquely identify this certificate. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.
     * 
     */
    @Export(name="keyId", refs={String.class}, tree="[0]")
    private Output<String> keyId;

    /**
     * @return A UUID used to uniquely identify this certificate. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.
     * 
     */
    public Output<String> keyId() {
        return this.keyId;
    }
    /**
     * The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn&#39;t specified, the value is determined by Azure Active Directory and is usually the start date of the certificate for asymmetric keys, or the current timestamp for symmetric keys. Changing this field forces a new resource to be created.
     * 
     */
    @Export(name="startDate", refs={String.class}, tree="[0]")
    private Output<String> startDate;

    /**
     * @return The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn&#39;t specified, the value is determined by Azure Active Directory and is usually the start date of the certificate for asymmetric keys, or the current timestamp for symmetric keys. Changing this field forces a new resource to be created.
     * 
     */
    public Output<String> startDate() {
        return this.startDate;
    }
    /**
     * The type of key/certificate. Must be one of `AsymmetricX509Cert` or `Symmetric`. Changing this fields forces a new resource to be created.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> type;

    /**
     * @return The type of key/certificate. Must be one of `AsymmetricX509Cert` or `Symmetric`. Changing this fields forces a new resource to be created.
     * 
     */
    public Output<Optional<String>> type() {
        return Codegen.optional(this.type);
    }
    /**
     * The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding` argument.
     * 
     */
    @Export(name="value", refs={String.class}, tree="[0]")
    private Output<String> value;

    /**
     * @return The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding` argument.
     * 
     */
    public Output<String> value() {
        return this.value;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ApplicationCertificate(java.lang.String name) {
        this(name, ApplicationCertificateArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ApplicationCertificate(java.lang.String name, ApplicationCertificateArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ApplicationCertificate(java.lang.String name, ApplicationCertificateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/applicationCertificate:ApplicationCertificate", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ApplicationCertificate(java.lang.String name, Output<java.lang.String> id, @Nullable ApplicationCertificateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/applicationCertificate:ApplicationCertificate", name, state, makeResourceOptions(options, id), false);
    }

    private static ApplicationCertificateArgs makeArgs(ApplicationCertificateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ApplicationCertificateArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static ApplicationCertificate get(java.lang.String name, Output<java.lang.String> id, @Nullable ApplicationCertificateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ApplicationCertificate(name, id, state, options);
    }
}

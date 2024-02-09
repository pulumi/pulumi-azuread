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
 * ## Import
 * 
 * Certificates can be imported using the object ID of the associated application and the key ID of the certificate credential, e.g.
 * 
 * ```sh
 * $ pulumi import azuread:index/applicationCertificate:ApplicationCertificate example 00000000-0000-0000-0000-000000000000/certificate/11111111-1111-1111-1111-111111111111
 * ```
 * 
 *  -&gt; This ID format is unique to Terraform and is composed of the application&#39;s object ID, the string &#34;certificate&#34; and the certificate&#39;s key ID in the format `{ObjectId}/certificate/{CertificateKeyId}`.
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
     * The object ID of the application for which this certificate should be created
     * 
     * @deprecated
     * The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider
     * 
     */
    @Deprecated /* The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider */
    @Export(name="applicationObjectId", refs={String.class}, tree="[0]")
    private Output<String> applicationObjectId;

    /**
     * @return The object ID of the application for which this certificate should be created
     * 
     */
    public Output<String> applicationObjectId() {
        return this.applicationObjectId;
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
     */
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
    public ApplicationCertificate(String name) {
        this(name, ApplicationCertificateArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ApplicationCertificate(String name, ApplicationCertificateArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ApplicationCertificate(String name, ApplicationCertificateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/applicationCertificate:ApplicationCertificate", name, args == null ? ApplicationCertificateArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ApplicationCertificate(String name, Output<String> id, @Nullable ApplicationCertificateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/applicationCertificate:ApplicationCertificate", name, state, makeResourceOptions(options, id));
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
    public static ApplicationCertificate get(String name, Output<String> id, @Nullable ApplicationCertificateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ApplicationCertificate(name, id, state, options);
    }
}

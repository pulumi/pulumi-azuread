// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * *Using a PEM certificate*
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 * import * as std from "@pulumi/std";
 *
 * const example = new azuread.ApplicationRegistration("example", {displayName: "example"});
 * const exampleApplicationCertificate = new azuread.ApplicationCertificate("example", {
 *     applicationId: example.id,
 *     type: "AsymmetricX509Cert",
 *     value: std.file({
 *         input: "cert.pem",
 *     }).then(invoke => invoke.result),
 *     endDate: "2021-05-01T01:02:03Z",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * *Using a DER certificate*
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 * import * as std from "@pulumi/std";
 *
 * const example = new azuread.ApplicationRegistration("example", {displayName: "example"});
 * const exampleApplicationCertificate = new azuread.ApplicationCertificate("example", {
 *     applicationId: example.id,
 *     type: "AsymmetricX509Cert",
 *     encoding: "base64",
 *     value: std.file({
 *         input: "cert.der",
 *     }).then(invoke => std.base64encode({
 *         input: invoke.result,
 *     })).then(invoke => invoke.result),
 *     endDate: "2021-05-01T01:02:03Z",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### Using a certificate from Azure Key Vault
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * import * as azuread from "@pulumi/azuread";
 *
 * const exampleApplication = new azuread.Application("example", {displayName: "example"});
 * const example = new azure.keyvault.Certificate("example", {
 *     name: "generated-cert",
 *     keyVaultId: exampleAzurermKeyVault.id,
 *     certificatePolicy: {
 *         issuerParameters: {
 *             name: "Self",
 *         },
 *         keyProperties: {
 *             exportable: true,
 *             keySize: 2048,
 *             keyType: "RSA",
 *             reuseKey: true,
 *         },
 *         lifetimeActions: [{
 *             action: {
 *                 actionType: "AutoRenew",
 *             },
 *             trigger: {
 *                 daysBeforeExpiry: 30,
 *             },
 *         }],
 *         secretProperties: {
 *             contentType: "application/x-pkcs12",
 *         },
 *         x509CertificateProperties: {
 *             extendedKeyUsages: ["1.3.6.1.5.5.7.3.2"],
 *             keyUsages: [
 *                 "dataEncipherment",
 *                 "digitalSignature",
 *                 "keyCertSign",
 *                 "keyEncipherment",
 *             ],
 *             subjectAlternativeNames: {
 *                 dnsNames: [
 *                     "internal.contoso.com",
 *                     "domain.hello.world",
 *                 ],
 *             },
 *             subject: `CN=${exampleApplication.name}`,
 *             validityInMonths: 12,
 *         },
 *     },
 * });
 * const exampleApplicationCertificate = new azuread.ApplicationCertificate("example", {
 *     applicationId: exampleApplication.id,
 *     type: "AsymmetricX509Cert",
 *     encoding: "hex",
 *     value: example.certificateData,
 *     endDate: example.certificateAttributes.apply(certificateAttributes => certificateAttributes[0].expires),
 *     startDate: example.certificateAttributes.apply(certificateAttributes => certificateAttributes[0].notBefore),
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Certificates can be imported using the object ID of the associated application and the key ID of the certificate credential, e.g.
 *
 * ```sh
 * $ pulumi import azuread:index/applicationCertificate:ApplicationCertificate example 00000000-0000-0000-0000-000000000000/certificate/11111111-1111-1111-1111-111111111111
 * ```
 *
 * -> This ID format is unique to Terraform and is composed of the application's object ID, the string "certificate" and the certificate's key ID in the format `{ObjectId}/certificate/{CertificateKeyId}`.
 */
export class ApplicationCertificate extends pulumi.CustomResource {
    /**
     * Get an existing ApplicationCertificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApplicationCertificateState, opts?: pulumi.CustomResourceOptions): ApplicationCertificate {
        return new ApplicationCertificate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/applicationCertificate:ApplicationCertificate';

    /**
     * Returns true if the given object is an instance of ApplicationCertificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApplicationCertificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApplicationCertificate.__pulumiType;
    }

    /**
     * The resource ID of the application for which this certificate should be created. Changing this field forces a new resource to be created.
     */
    public readonly applicationId!: pulumi.Output<string>;
    /**
     * The object ID of the application for which this certificate should be created
     *
     * @deprecated The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider
     */
    public readonly applicationObjectId!: pulumi.Output<string>;
    /**
     * Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.
     *
     * > **Tip for Azure Key Vault** The `hex` encoding option is useful for consuming certificate data from the azurermKeyVaultCertificate resource.
     */
    public readonly encoding!: pulumi.Output<string | undefined>;
    /**
     * The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If omitted, the API will decide a suitable expiry date, which is typically around 2 years from the start date. Changing this field forces a new resource to be created.
     */
    public readonly endDate!: pulumi.Output<string>;
    /**
     * A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
     *
     * > One of `endDate` or `endDateRelative` must be specified. The maximum allowed duration is determined by Azure AD and is typically around 2 years from the creation date.
     */
    public readonly endDateRelative!: pulumi.Output<string | undefined>;
    /**
     * A UUID used to uniquely identify this certificate. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.
     */
    public readonly keyId!: pulumi.Output<string>;
    /**
     * The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the value is determined by Azure Active Directory and is usually the start date of the certificate for asymmetric keys, or the current timestamp for symmetric keys. Changing this field forces a new resource to be created.
     */
    public readonly startDate!: pulumi.Output<string>;
    /**
     * The type of key/certificate. Must be one of `AsymmetricX509Cert` or `Symmetric`. Changing this fields forces a new resource to be created.
     */
    public readonly type!: pulumi.Output<string | undefined>;
    /**
     * The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding` argument.
     */
    public readonly value!: pulumi.Output<string>;

    /**
     * Create a ApplicationCertificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationCertificateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationCertificateArgs | ApplicationCertificateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApplicationCertificateState | undefined;
            resourceInputs["applicationId"] = state ? state.applicationId : undefined;
            resourceInputs["applicationObjectId"] = state ? state.applicationObjectId : undefined;
            resourceInputs["encoding"] = state ? state.encoding : undefined;
            resourceInputs["endDate"] = state ? state.endDate : undefined;
            resourceInputs["endDateRelative"] = state ? state.endDateRelative : undefined;
            resourceInputs["keyId"] = state ? state.keyId : undefined;
            resourceInputs["startDate"] = state ? state.startDate : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as ApplicationCertificateArgs | undefined;
            if ((!args || args.value === undefined) && !opts.urn) {
                throw new Error("Missing required property 'value'");
            }
            resourceInputs["applicationId"] = args ? args.applicationId : undefined;
            resourceInputs["applicationObjectId"] = args ? args.applicationObjectId : undefined;
            resourceInputs["encoding"] = args ? args.encoding : undefined;
            resourceInputs["endDate"] = args ? args.endDate : undefined;
            resourceInputs["endDateRelative"] = args ? args.endDateRelative : undefined;
            resourceInputs["keyId"] = args ? args.keyId : undefined;
            resourceInputs["startDate"] = args ? args.startDate : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["value"] = args?.value ? pulumi.secret(args.value) : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["value"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ApplicationCertificate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApplicationCertificate resources.
 */
export interface ApplicationCertificateState {
    /**
     * The resource ID of the application for which this certificate should be created. Changing this field forces a new resource to be created.
     */
    applicationId?: pulumi.Input<string>;
    /**
     * The object ID of the application for which this certificate should be created
     *
     * @deprecated The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider
     */
    applicationObjectId?: pulumi.Input<string>;
    /**
     * Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.
     *
     * > **Tip for Azure Key Vault** The `hex` encoding option is useful for consuming certificate data from the azurermKeyVaultCertificate resource.
     */
    encoding?: pulumi.Input<string>;
    /**
     * The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If omitted, the API will decide a suitable expiry date, which is typically around 2 years from the start date. Changing this field forces a new resource to be created.
     */
    endDate?: pulumi.Input<string>;
    /**
     * A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
     *
     * > One of `endDate` or `endDateRelative` must be specified. The maximum allowed duration is determined by Azure AD and is typically around 2 years from the creation date.
     */
    endDateRelative?: pulumi.Input<string>;
    /**
     * A UUID used to uniquely identify this certificate. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.
     */
    keyId?: pulumi.Input<string>;
    /**
     * The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the value is determined by Azure Active Directory and is usually the start date of the certificate for asymmetric keys, or the current timestamp for symmetric keys. Changing this field forces a new resource to be created.
     */
    startDate?: pulumi.Input<string>;
    /**
     * The type of key/certificate. Must be one of `AsymmetricX509Cert` or `Symmetric`. Changing this fields forces a new resource to be created.
     */
    type?: pulumi.Input<string>;
    /**
     * The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding` argument.
     */
    value?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ApplicationCertificate resource.
 */
export interface ApplicationCertificateArgs {
    /**
     * The resource ID of the application for which this certificate should be created. Changing this field forces a new resource to be created.
     */
    applicationId?: pulumi.Input<string>;
    /**
     * The object ID of the application for which this certificate should be created
     *
     * @deprecated The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider
     */
    applicationObjectId?: pulumi.Input<string>;
    /**
     * Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.
     *
     * > **Tip for Azure Key Vault** The `hex` encoding option is useful for consuming certificate data from the azurermKeyVaultCertificate resource.
     */
    encoding?: pulumi.Input<string>;
    /**
     * The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If omitted, the API will decide a suitable expiry date, which is typically around 2 years from the start date. Changing this field forces a new resource to be created.
     */
    endDate?: pulumi.Input<string>;
    /**
     * A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
     *
     * > One of `endDate` or `endDateRelative` must be specified. The maximum allowed duration is determined by Azure AD and is typically around 2 years from the creation date.
     */
    endDateRelative?: pulumi.Input<string>;
    /**
     * A UUID used to uniquely identify this certificate. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.
     */
    keyId?: pulumi.Input<string>;
    /**
     * The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the value is determined by Azure Active Directory and is usually the start date of the certificate for asymmetric keys, or the current timestamp for symmetric keys. Changing this field forces a new resource to be created.
     */
    startDate?: pulumi.Input<string>;
    /**
     * The type of key/certificate. Must be one of `AsymmetricX509Cert` or `Symmetric`. Changing this fields forces a new resource to be created.
     */
    type?: pulumi.Input<string>;
    /**
     * The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding` argument.
     */
    value: pulumi.Input<string>;
}

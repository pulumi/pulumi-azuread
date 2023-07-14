// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class ServicePrincipalTokenSigningCertificate extends pulumi.CustomResource {
    /**
     * Get an existing ServicePrincipalTokenSigningCertificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServicePrincipalTokenSigningCertificateState, opts?: pulumi.CustomResourceOptions): ServicePrincipalTokenSigningCertificate {
        return new ServicePrincipalTokenSigningCertificate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/servicePrincipalTokenSigningCertificate:ServicePrincipalTokenSigningCertificate';

    /**
     * Returns true if the given object is an instance of ServicePrincipalTokenSigningCertificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServicePrincipalTokenSigningCertificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServicePrincipalTokenSigningCertificate.__pulumiType;
    }

    /**
     * A friendly name for the certificate
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
     * Default is 3 years from current date.
     */
    public readonly endDate!: pulumi.Output<string>;
    /**
     * A UUID used to uniquely identify the verify certificate.
     */
    public /*out*/ readonly keyId!: pulumi.Output<string>;
    /**
     * The object ID of the service principal for which this certificate should be created
     */
    public readonly servicePrincipalId!: pulumi.Output<string>;
    /**
     * The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
     */
    public /*out*/ readonly startDate!: pulumi.Output<string>;
    /**
     * The thumbprint of the certificate.
     */
    public /*out*/ readonly thumbprint!: pulumi.Output<string>;
    /**
     * The certificate data, which is PEM encoded but does not include the header/footer
     */
    public /*out*/ readonly value!: pulumi.Output<string>;

    /**
     * Create a ServicePrincipalTokenSigningCertificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServicePrincipalTokenSigningCertificateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServicePrincipalTokenSigningCertificateArgs | ServicePrincipalTokenSigningCertificateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServicePrincipalTokenSigningCertificateState | undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["endDate"] = state ? state.endDate : undefined;
            resourceInputs["keyId"] = state ? state.keyId : undefined;
            resourceInputs["servicePrincipalId"] = state ? state.servicePrincipalId : undefined;
            resourceInputs["startDate"] = state ? state.startDate : undefined;
            resourceInputs["thumbprint"] = state ? state.thumbprint : undefined;
            resourceInputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as ServicePrincipalTokenSigningCertificateArgs | undefined;
            if ((!args || args.servicePrincipalId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'servicePrincipalId'");
            }
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["endDate"] = args ? args.endDate : undefined;
            resourceInputs["servicePrincipalId"] = args ? args.servicePrincipalId : undefined;
            resourceInputs["keyId"] = undefined /*out*/;
            resourceInputs["startDate"] = undefined /*out*/;
            resourceInputs["thumbprint"] = undefined /*out*/;
            resourceInputs["value"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["value"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ServicePrincipalTokenSigningCertificate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServicePrincipalTokenSigningCertificate resources.
 */
export interface ServicePrincipalTokenSigningCertificateState {
    /**
     * A friendly name for the certificate
     */
    displayName?: pulumi.Input<string>;
    /**
     * The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
     * Default is 3 years from current date.
     */
    endDate?: pulumi.Input<string>;
    /**
     * A UUID used to uniquely identify the verify certificate.
     */
    keyId?: pulumi.Input<string>;
    /**
     * The object ID of the service principal for which this certificate should be created
     */
    servicePrincipalId?: pulumi.Input<string>;
    /**
     * The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
     */
    startDate?: pulumi.Input<string>;
    /**
     * The thumbprint of the certificate.
     */
    thumbprint?: pulumi.Input<string>;
    /**
     * The certificate data, which is PEM encoded but does not include the header/footer
     */
    value?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServicePrincipalTokenSigningCertificate resource.
 */
export interface ServicePrincipalTokenSigningCertificateArgs {
    /**
     * A friendly name for the certificate
     */
    displayName?: pulumi.Input<string>;
    /**
     * The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
     * Default is 3 years from current date.
     */
    endDate?: pulumi.Input<string>;
    /**
     * The object ID of the service principal for which this certificate should be created
     */
    servicePrincipalId: pulumi.Input<string>;
}

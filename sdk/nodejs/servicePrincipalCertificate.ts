// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a certificate associated with a Service Principal within Azure Active Directory.
 *
 * > **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write all applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.
 *
 * ## Import
 *
 * Certificates can be imported using the `object id` of the Service Principal and the `key id` of the certificate, e.g.
 *
 * ```sh
 *  $ pulumi import azuread:index/servicePrincipalCertificate:ServicePrincipalCertificate test 00000000-0000-0000-0000-000000000000/certificate/11111111-1111-1111-1111-111111111111
 * ```
 */
export class ServicePrincipalCertificate extends pulumi.CustomResource {
    /**
     * Get an existing ServicePrincipalCertificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServicePrincipalCertificateState, opts?: pulumi.CustomResourceOptions): ServicePrincipalCertificate {
        return new ServicePrincipalCertificate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/servicePrincipalCertificate:ServicePrincipalCertificate';

    /**
     * Returns true if the given object is an instance of ServicePrincipalCertificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServicePrincipalCertificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServicePrincipalCertificate.__pulumiType;
    }

    /**
     * Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.
     */
    public readonly encoding!: pulumi.Output<string | undefined>;
    /**
     * The End Date which the Certificate is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
     */
    public readonly endDate!: pulumi.Output<string>;
    /**
     * A relative duration for which the Certificate is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". Changing this field forces a new resource to be created.
     */
    public readonly endDateRelative!: pulumi.Output<string | undefined>;
    /**
     * A GUID used to uniquely identify this Certificate. If not specified a GUID will be created. Changing this field forces a new resource to be created.
     */
    public readonly keyId!: pulumi.Output<string>;
    /**
     * The ID of the Service Principal for which this certificate should be created. Changing this field forces a new resource to be created.
     */
    public readonly servicePrincipalId!: pulumi.Output<string>;
    /**
     * The Start Date which the Certificate is valid from, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
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
     * Create a ServicePrincipalCertificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServicePrincipalCertificateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServicePrincipalCertificateArgs | ServicePrincipalCertificateState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServicePrincipalCertificateState | undefined;
            inputs["encoding"] = state ? state.encoding : undefined;
            inputs["endDate"] = state ? state.endDate : undefined;
            inputs["endDateRelative"] = state ? state.endDateRelative : undefined;
            inputs["keyId"] = state ? state.keyId : undefined;
            inputs["servicePrincipalId"] = state ? state.servicePrincipalId : undefined;
            inputs["startDate"] = state ? state.startDate : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as ServicePrincipalCertificateArgs | undefined;
            if ((!args || args.servicePrincipalId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'servicePrincipalId'");
            }
            if ((!args || args.value === undefined) && !opts.urn) {
                throw new Error("Missing required property 'value'");
            }
            inputs["encoding"] = args ? args.encoding : undefined;
            inputs["endDate"] = args ? args.endDate : undefined;
            inputs["endDateRelative"] = args ? args.endDateRelative : undefined;
            inputs["keyId"] = args ? args.keyId : undefined;
            inputs["servicePrincipalId"] = args ? args.servicePrincipalId : undefined;
            inputs["startDate"] = args ? args.startDate : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["value"] = args ? args.value : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ServicePrincipalCertificate.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServicePrincipalCertificate resources.
 */
export interface ServicePrincipalCertificateState {
    /**
     * Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.
     */
    encoding?: pulumi.Input<string>;
    /**
     * The End Date which the Certificate is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
     */
    endDate?: pulumi.Input<string>;
    /**
     * A relative duration for which the Certificate is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". Changing this field forces a new resource to be created.
     */
    endDateRelative?: pulumi.Input<string>;
    /**
     * A GUID used to uniquely identify this Certificate. If not specified a GUID will be created. Changing this field forces a new resource to be created.
     */
    keyId?: pulumi.Input<string>;
    /**
     * The ID of the Service Principal for which this certificate should be created. Changing this field forces a new resource to be created.
     */
    servicePrincipalId?: pulumi.Input<string>;
    /**
     * The Start Date which the Certificate is valid from, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
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
 * The set of arguments for constructing a ServicePrincipalCertificate resource.
 */
export interface ServicePrincipalCertificateArgs {
    /**
     * Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.
     */
    encoding?: pulumi.Input<string>;
    /**
     * The End Date which the Certificate is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
     */
    endDate?: pulumi.Input<string>;
    /**
     * A relative duration for which the Certificate is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". Changing this field forces a new resource to be created.
     */
    endDateRelative?: pulumi.Input<string>;
    /**
     * A GUID used to uniquely identify this Certificate. If not specified a GUID will be created. Changing this field forces a new resource to be created.
     */
    keyId?: pulumi.Input<string>;
    /**
     * The ID of the Service Principal for which this certificate should be created. Changing this field forces a new resource to be created.
     */
    servicePrincipalId: pulumi.Input<string>;
    /**
     * The Start Date which the Certificate is valid from, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
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

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class SynchronizationSecret extends pulumi.CustomResource {
    /**
     * Get an existing SynchronizationSecret resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SynchronizationSecretState, opts?: pulumi.CustomResourceOptions): SynchronizationSecret {
        return new SynchronizationSecret(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/synchronizationSecret:SynchronizationSecret';

    /**
     * Returns true if the given object is an instance of SynchronizationSecret.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SynchronizationSecret {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SynchronizationSecret.__pulumiType;
    }

    public readonly credentials!: pulumi.Output<outputs.SynchronizationSecretCredential[] | undefined>;
    /**
     * The object ID of the service principal for which this synchronization secret should be created
     */
    public readonly servicePrincipalId!: pulumi.Output<string>;

    /**
     * Create a SynchronizationSecret resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SynchronizationSecretArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SynchronizationSecretArgs | SynchronizationSecretState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SynchronizationSecretState | undefined;
            resourceInputs["credentials"] = state ? state.credentials : undefined;
            resourceInputs["servicePrincipalId"] = state ? state.servicePrincipalId : undefined;
        } else {
            const args = argsOrState as SynchronizationSecretArgs | undefined;
            if ((!args || args.servicePrincipalId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'servicePrincipalId'");
            }
            resourceInputs["credentials"] = args ? args.credentials : undefined;
            resourceInputs["servicePrincipalId"] = args ? args.servicePrincipalId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SynchronizationSecret.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SynchronizationSecret resources.
 */
export interface SynchronizationSecretState {
    credentials?: pulumi.Input<pulumi.Input<inputs.SynchronizationSecretCredential>[]>;
    /**
     * The object ID of the service principal for which this synchronization secret should be created
     */
    servicePrincipalId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SynchronizationSecret resource.
 */
export interface SynchronizationSecretArgs {
    credentials?: pulumi.Input<pulumi.Input<inputs.SynchronizationSecretCredential>[]>;
    /**
     * The object ID of the service principal for which this synchronization secret should be created
     */
    servicePrincipalId: pulumi.Input<string>;
}

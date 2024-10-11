// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Manages synchronization secrets associated with a service principal (enterprise application) within Azure Active Directory.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this resource.
 *
 * When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.All` or `Directory.ReadWrite.All`
 *
 * ## Example Usage
 *
 * *Basic example*
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = azuread.getApplicationTemplate({
 *     displayName: "Azure Databricks SCIM Provisioning Connector",
 * });
 * const exampleApplicationFromTemplate = new azuread.ApplicationFromTemplate("example", {
 *     displayName: "example",
 *     templateId: example.then(example => example.templateId),
 * });
 * const exampleGetServicePrincipal = azuread.getServicePrincipalOutput({
 *     objectId: exampleApplicationFromTemplate.servicePrincipalObjectId,
 * });
 * const exampleSynchronizationSecret = new azuread.SynchronizationSecret("example", {
 *     servicePrincipalId: exampleGetServicePrincipal.apply(exampleGetServicePrincipal => exampleGetServicePrincipal.id),
 *     credentials: [
 *         {
 *             key: "BaseAddress",
 *             value: "abc",
 *         },
 *         {
 *             key: "SecretToken",
 *             value: "some-token",
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * This resource does not support importing.
 */
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

    /**
     * One or more `credential` blocks as documented below.
     */
    public readonly credentials!: pulumi.Output<outputs.SynchronizationSecretCredential[] | undefined>;
    /**
     * The ID of the service principal for which this synchronization secrets should be stored. Changing this field forces a new resource to be created.
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
    /**
     * One or more `credential` blocks as documented below.
     */
    credentials?: pulumi.Input<pulumi.Input<inputs.SynchronizationSecretCredential>[]>;
    /**
     * The ID of the service principal for which this synchronization secrets should be stored. Changing this field forces a new resource to be created.
     */
    servicePrincipalId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SynchronizationSecret resource.
 */
export interface SynchronizationSecretArgs {
    /**
     * One or more `credential` blocks as documented below.
     */
    credentials?: pulumi.Input<pulumi.Input<inputs.SynchronizationSecretCredential>[]>;
    /**
     * The ID of the service principal for which this synchronization secrets should be stored. Changing this field forces a new resource to be created.
     */
    servicePrincipalId: pulumi.Input<string>;
}

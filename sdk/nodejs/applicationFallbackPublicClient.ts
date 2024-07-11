// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = new azuread.ApplicationRegistration("example", {displayName: "example"});
 * const exampleApplicationFallbackPublicClient = new azuread.ApplicationFallbackPublicClient("example", {
 *     applicationId: example.id,
 *     enabled: true,
 * });
 * ```
 *
 * ## Import
 *
 * The Application Fallback Public Client setting can be imported using the object ID of the application, in the following format.
 *
 * ```sh
 * $ pulumi import azuread:index/applicationFallbackPublicClient:ApplicationFallbackPublicClient example /applications/00000000-0000-0000-0000-000000000000/fallbackPublicClient
 * ```
 */
export class ApplicationFallbackPublicClient extends pulumi.CustomResource {
    /**
     * Get an existing ApplicationFallbackPublicClient resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApplicationFallbackPublicClientState, opts?: pulumi.CustomResourceOptions): ApplicationFallbackPublicClient {
        return new ApplicationFallbackPublicClient(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/applicationFallbackPublicClient:ApplicationFallbackPublicClient';

    /**
     * Returns true if the given object is an instance of ApplicationFallbackPublicClient.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApplicationFallbackPublicClient {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApplicationFallbackPublicClient.__pulumiType;
    }

    /**
     * The resource ID of the application registration. Changing this forces a new resource to be created.
     */
    public readonly applicationId!: pulumi.Output<string>;
    /**
     * Whether to enable the application as a fallback public client.
     *
     * > Some configurations may require the Fallback Public Client setting to be `null`, for this case simply destroy this resource (or don't use it)
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;

    /**
     * Create a ApplicationFallbackPublicClient resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationFallbackPublicClientArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationFallbackPublicClientArgs | ApplicationFallbackPublicClientState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApplicationFallbackPublicClientState | undefined;
            resourceInputs["applicationId"] = state ? state.applicationId : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
        } else {
            const args = argsOrState as ApplicationFallbackPublicClientArgs | undefined;
            if ((!args || args.applicationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationId'");
            }
            resourceInputs["applicationId"] = args ? args.applicationId : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ApplicationFallbackPublicClient.__pulumiType, name, resourceInputs, opts, false /*remote*/);
    }
}

/**
 * Input properties used for looking up and filtering ApplicationFallbackPublicClient resources.
 */
export interface ApplicationFallbackPublicClientState {
    /**
     * The resource ID of the application registration. Changing this forces a new resource to be created.
     */
    applicationId?: pulumi.Input<string>;
    /**
     * Whether to enable the application as a fallback public client.
     *
     * > Some configurations may require the Fallback Public Client setting to be `null`, for this case simply destroy this resource (or don't use it)
     */
    enabled?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a ApplicationFallbackPublicClient resource.
 */
export interface ApplicationFallbackPublicClientArgs {
    /**
     * The resource ID of the application registration. Changing this forces a new resource to be created.
     */
    applicationId: pulumi.Input<string>;
    /**
     * Whether to enable the application as a fallback public client.
     *
     * > Some configurations may require the Fallback Public Client setting to be `null`, for this case simply destroy this resource (or don't use it)
     */
    enabled?: pulumi.Input<boolean>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Import
 *
 * Application API Access can be imported using the object ID of the application and the client ID of the API, in the following format.
 *
 * ```sh
 *  $ pulumi import azuread:index/applicationApiAccess:ApplicationApiAccess example /applications/00000000-0000-0000-0000-000000000000/apiAccess/11111111-1111-1111-1111-111111111111
 * ```
 */
export class ApplicationApiAccess extends pulumi.CustomResource {
    /**
     * Get an existing ApplicationApiAccess resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApplicationApiAccessState, opts?: pulumi.CustomResourceOptions): ApplicationApiAccess {
        return new ApplicationApiAccess(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/applicationApiAccess:ApplicationApiAccess';

    /**
     * Returns true if the given object is an instance of ApplicationApiAccess.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApplicationApiAccess {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApplicationApiAccess.__pulumiType;
    }

    /**
     * The client ID of the API to which access is being granted. Changing this forces a new resource to be created.
     */
    public readonly apiClientId!: pulumi.Output<string>;
    /**
     * The resource ID of the application registration. Changing this forces a new resource to be created.
     */
    public readonly applicationId!: pulumi.Output<string>;
    /**
     * A set of role IDs to be granted to the application, as published by the API.
     */
    public readonly roleIds!: pulumi.Output<string[] | undefined>;
    /**
     * A set of scope IDs to be granted to the application, as published by the API.
     *
     * > At least one of `roleIds` or `scopeIds` must be specified.
     */
    public readonly scopeIds!: pulumi.Output<string[] | undefined>;

    /**
     * Create a ApplicationApiAccess resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationApiAccessArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationApiAccessArgs | ApplicationApiAccessState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApplicationApiAccessState | undefined;
            resourceInputs["apiClientId"] = state ? state.apiClientId : undefined;
            resourceInputs["applicationId"] = state ? state.applicationId : undefined;
            resourceInputs["roleIds"] = state ? state.roleIds : undefined;
            resourceInputs["scopeIds"] = state ? state.scopeIds : undefined;
        } else {
            const args = argsOrState as ApplicationApiAccessArgs | undefined;
            if ((!args || args.apiClientId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiClientId'");
            }
            if ((!args || args.applicationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationId'");
            }
            resourceInputs["apiClientId"] = args ? args.apiClientId : undefined;
            resourceInputs["applicationId"] = args ? args.applicationId : undefined;
            resourceInputs["roleIds"] = args ? args.roleIds : undefined;
            resourceInputs["scopeIds"] = args ? args.scopeIds : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ApplicationApiAccess.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApplicationApiAccess resources.
 */
export interface ApplicationApiAccessState {
    /**
     * The client ID of the API to which access is being granted. Changing this forces a new resource to be created.
     */
    apiClientId?: pulumi.Input<string>;
    /**
     * The resource ID of the application registration. Changing this forces a new resource to be created.
     */
    applicationId?: pulumi.Input<string>;
    /**
     * A set of role IDs to be granted to the application, as published by the API.
     */
    roleIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A set of scope IDs to be granted to the application, as published by the API.
     *
     * > At least one of `roleIds` or `scopeIds` must be specified.
     */
    scopeIds?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a ApplicationApiAccess resource.
 */
export interface ApplicationApiAccessArgs {
    /**
     * The client ID of the API to which access is being granted. Changing this forces a new resource to be created.
     */
    apiClientId: pulumi.Input<string>;
    /**
     * The resource ID of the application registration. Changing this forces a new resource to be created.
     */
    applicationId: pulumi.Input<string>;
    /**
     * A set of role IDs to be granted to the application, as published by the API.
     */
    roleIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A set of scope IDs to be granted to the application, as published by the API.
     *
     * > At least one of `roleIds` or `scopeIds` must be specified.
     */
    scopeIds?: pulumi.Input<pulumi.Input<string>[]>;
}

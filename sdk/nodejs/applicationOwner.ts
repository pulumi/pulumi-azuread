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
 * const jane = new azuread.User("jane", {
 *     userPrincipalName: "jane.fischer@hashitown.com",
 *     displayName: "Jane Fischer",
 *     password: "Ch@ngeMe",
 * });
 * const exampleJane = new azuread.ApplicationOwner("exampleJane", {
 *     applicationId: example.id,
 *     ownerObjectId: jane.objectId,
 * });
 * ```
 *
 * > **Tip** For managing more application owners, create additional instances of this resource
 *
 * ## Import
 *
 * Application Owners can be imported using the object ID of the application and the object ID of the owner, in the following format.
 *
 * ```sh
 *  $ pulumi import azuread:index/applicationOwner:ApplicationOwner example /applications/00000000-0000-0000-0000-000000000000/owners/11111111-1111-1111-1111-111111111111
 * ```
 */
export class ApplicationOwner extends pulumi.CustomResource {
    /**
     * Get an existing ApplicationOwner resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApplicationOwnerState, opts?: pulumi.CustomResourceOptions): ApplicationOwner {
        return new ApplicationOwner(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/applicationOwner:ApplicationOwner';

    /**
     * Returns true if the given object is an instance of ApplicationOwner.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApplicationOwner {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApplicationOwner.__pulumiType;
    }

    /**
     * The resource ID of the application registration. Changing this forces a new resource to be created.
     */
    public readonly applicationId!: pulumi.Output<string>;
    /**
     * The object ID of the owner to assign to the application, typically a user or service principal. Changing this forces a new resource to be created.
     */
    public readonly ownerObjectId!: pulumi.Output<string>;

    /**
     * Create a ApplicationOwner resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationOwnerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationOwnerArgs | ApplicationOwnerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApplicationOwnerState | undefined;
            resourceInputs["applicationId"] = state ? state.applicationId : undefined;
            resourceInputs["ownerObjectId"] = state ? state.ownerObjectId : undefined;
        } else {
            const args = argsOrState as ApplicationOwnerArgs | undefined;
            if ((!args || args.applicationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationId'");
            }
            if ((!args || args.ownerObjectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ownerObjectId'");
            }
            resourceInputs["applicationId"] = args ? args.applicationId : undefined;
            resourceInputs["ownerObjectId"] = args ? args.ownerObjectId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ApplicationOwner.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApplicationOwner resources.
 */
export interface ApplicationOwnerState {
    /**
     * The resource ID of the application registration. Changing this forces a new resource to be created.
     */
    applicationId?: pulumi.Input<string>;
    /**
     * The object ID of the owner to assign to the application, typically a user or service principal. Changing this forces a new resource to be created.
     */
    ownerObjectId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ApplicationOwner resource.
 */
export interface ApplicationOwnerArgs {
    /**
     * The resource ID of the application registration. Changing this forces a new resource to be created.
     */
    applicationId: pulumi.Input<string>;
    /**
     * The object ID of the owner to assign to the application, typically a user or service principal. Changing this forces a new resource to be created.
     */
    ownerObjectId: pulumi.Input<string>;
}
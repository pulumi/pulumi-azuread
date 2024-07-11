// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages an Administrative Unit within Azure Active Directory.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this resource.
 *
 * When authenticated with a service principal, this resource requires one of the following application roles: `AdministrativeUnit.ReadWrite.All` or `Directory.ReadWrite.All`
 *
 * When authenticated with a user principal, this resource requires one of the following directory roles: `Privileged Role Administrator` or `Global Administrator`
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = new azuread.AdministrativeUnit("example", {
 *     displayName: "Example-AU",
 *     description: "Just an example",
 *     hiddenMembershipEnabled: false,
 * });
 * ```
 *
 * ## Import
 *
 * Administrative units can be imported using their object ID, e.g.
 *
 * ```sh
 * $ pulumi import azuread:index/administrativeUnit:AdministrativeUnit example 00000000-0000-0000-0000-000000000000
 * ```
 */
export class AdministrativeUnit extends pulumi.CustomResource {
    /**
     * Get an existing AdministrativeUnit resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AdministrativeUnitState, opts?: pulumi.CustomResourceOptions): AdministrativeUnit {
        return new AdministrativeUnit(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/administrativeUnit:AdministrativeUnit';

    /**
     * Returns true if the given object is an instance of AdministrativeUnit.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AdministrativeUnit {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AdministrativeUnit.__pulumiType;
    }

    /**
     * The description of the administrative unit.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The display name of the administrative unit.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Whether the administrative unit and its members are hidden or publicly viewable in the directory.
     */
    public readonly hiddenMembershipEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * A set of object IDs of members who should be present in this administrative unit. Supported object types are Users or Groups.
     *
     * > **Caution** When using the `members` property of the azuread.AdministrativeUnit resource, to manage Administrative Unit membership for a group, you will need to use an `ignoreChanges = [administrativeUnitIds]` lifecycle meta argument for the `azuread.Group` resource, in order to avoid a persistent diff.
     *
     * !> **Warning** Do not use the `members` property at the same time as the azuread.AdministrativeUnitMember resource for the same administrative unit. Doing so will cause a conflict and administrative unit members will be removed.
     */
    public readonly members!: pulumi.Output<string[]>;
    /**
     * The object ID of the administrative unit.
     */
    public /*out*/ readonly objectId!: pulumi.Output<string>;
    /**
     * If `true`, will return an error if an existing administrative unit is found with the same name
     */
    public readonly preventDuplicateNames!: pulumi.Output<boolean | undefined>;

    /**
     * Create a AdministrativeUnit resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AdministrativeUnitArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AdministrativeUnitArgs | AdministrativeUnitState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AdministrativeUnitState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["hiddenMembershipEnabled"] = state ? state.hiddenMembershipEnabled : undefined;
            resourceInputs["members"] = state ? state.members : undefined;
            resourceInputs["objectId"] = state ? state.objectId : undefined;
            resourceInputs["preventDuplicateNames"] = state ? state.preventDuplicateNames : undefined;
        } else {
            const args = argsOrState as AdministrativeUnitArgs | undefined;
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["hiddenMembershipEnabled"] = args ? args.hiddenMembershipEnabled : undefined;
            resourceInputs["members"] = args ? args.members : undefined;
            resourceInputs["preventDuplicateNames"] = args ? args.preventDuplicateNames : undefined;
            resourceInputs["objectId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AdministrativeUnit.__pulumiType, name, resourceInputs, opts, false /*remote*/);
    }
}

/**
 * Input properties used for looking up and filtering AdministrativeUnit resources.
 */
export interface AdministrativeUnitState {
    /**
     * The description of the administrative unit.
     */
    description?: pulumi.Input<string>;
    /**
     * The display name of the administrative unit.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Whether the administrative unit and its members are hidden or publicly viewable in the directory.
     */
    hiddenMembershipEnabled?: pulumi.Input<boolean>;
    /**
     * A set of object IDs of members who should be present in this administrative unit. Supported object types are Users or Groups.
     *
     * > **Caution** When using the `members` property of the azuread.AdministrativeUnit resource, to manage Administrative Unit membership for a group, you will need to use an `ignoreChanges = [administrativeUnitIds]` lifecycle meta argument for the `azuread.Group` resource, in order to avoid a persistent diff.
     *
     * !> **Warning** Do not use the `members` property at the same time as the azuread.AdministrativeUnitMember resource for the same administrative unit. Doing so will cause a conflict and administrative unit members will be removed.
     */
    members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The object ID of the administrative unit.
     */
    objectId?: pulumi.Input<string>;
    /**
     * If `true`, will return an error if an existing administrative unit is found with the same name
     */
    preventDuplicateNames?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a AdministrativeUnit resource.
 */
export interface AdministrativeUnitArgs {
    /**
     * The description of the administrative unit.
     */
    description?: pulumi.Input<string>;
    /**
     * The display name of the administrative unit.
     */
    displayName: pulumi.Input<string>;
    /**
     * Whether the administrative unit and its members are hidden or publicly viewable in the directory.
     */
    hiddenMembershipEnabled?: pulumi.Input<boolean>;
    /**
     * A set of object IDs of members who should be present in this administrative unit. Supported object types are Users or Groups.
     *
     * > **Caution** When using the `members` property of the azuread.AdministrativeUnit resource, to manage Administrative Unit membership for a group, you will need to use an `ignoreChanges = [administrativeUnitIds]` lifecycle meta argument for the `azuread.Group` resource, in order to avoid a persistent diff.
     *
     * !> **Warning** Do not use the `members` property at the same time as the azuread.AdministrativeUnitMember resource for the same administrative unit. Doing so will cause a conflict and administrative unit members will be removed.
     */
    members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If `true`, will return an error if an existing administrative unit is found with the same name
     */
    preventDuplicateNames?: pulumi.Input<boolean>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a single directory role assignment scoped to an administrative unit within Azure Active Directory.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this resource.
 *
 * When authenticated with a service principal, this resource requires one of the following application roles: `AdministrativeUnit.ReadWrite.All` and `RoleManagement.ReadWrite.Directory`, or `Directory.ReadWrite.All`
 *
 * When authenticated with a user principal, this resource requires one of the following directory roles: `Privileged Role Administrator` or `Global Administrator`
 *
 * ## Import
 *
 * Administrative unit role members can be imported using the object ID of the administrative unit and the unique ID of the role assignment, e.g.
 *
 * ```sh
 * $ pulumi import azuread:index/administrativeUnitRoleMember:AdministrativeUnitRoleMember example 00000000-0000-0000-0000-000000000000/roleMember/zX37MRLyF0uvE-xf2WH4B7x-6CPLfudNnxFGj800htpBXqkxW7bITqGb6Rj4kuTuS
 * ```
 *
 *  -> This ID format is unique to Terraform and is composed of the Administrative Unit Object ID and the role assignment ID in the format `{AdministrativeUnitObjectID}/roleMember/{RoleAssignmentID}`.
 */
export class AdministrativeUnitRoleMember extends pulumi.CustomResource {
    /**
     * Get an existing AdministrativeUnitRoleMember resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AdministrativeUnitRoleMemberState, opts?: pulumi.CustomResourceOptions): AdministrativeUnitRoleMember {
        return new AdministrativeUnitRoleMember(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/administrativeUnitRoleMember:AdministrativeUnitRoleMember';

    /**
     * Returns true if the given object is an instance of AdministrativeUnitRoleMember.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AdministrativeUnitRoleMember {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AdministrativeUnitRoleMember.__pulumiType;
    }

    /**
     * The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
     */
    public readonly administrativeUnitObjectId!: pulumi.Output<string>;
    /**
     * The object ID of the user, group or service principal you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
     */
    public readonly memberObjectId!: pulumi.Output<string>;
    /**
     * The object ID of the directory role you want to assign. Changing this forces a new resource to be created.
     */
    public readonly roleObjectId!: pulumi.Output<string>;

    /**
     * Create a AdministrativeUnitRoleMember resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AdministrativeUnitRoleMemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AdministrativeUnitRoleMemberArgs | AdministrativeUnitRoleMemberState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AdministrativeUnitRoleMemberState | undefined;
            resourceInputs["administrativeUnitObjectId"] = state ? state.administrativeUnitObjectId : undefined;
            resourceInputs["memberObjectId"] = state ? state.memberObjectId : undefined;
            resourceInputs["roleObjectId"] = state ? state.roleObjectId : undefined;
        } else {
            const args = argsOrState as AdministrativeUnitRoleMemberArgs | undefined;
            if ((!args || args.administrativeUnitObjectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'administrativeUnitObjectId'");
            }
            if ((!args || args.memberObjectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'memberObjectId'");
            }
            if ((!args || args.roleObjectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleObjectId'");
            }
            resourceInputs["administrativeUnitObjectId"] = args ? args.administrativeUnitObjectId : undefined;
            resourceInputs["memberObjectId"] = args ? args.memberObjectId : undefined;
            resourceInputs["roleObjectId"] = args ? args.roleObjectId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AdministrativeUnitRoleMember.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AdministrativeUnitRoleMember resources.
 */
export interface AdministrativeUnitRoleMemberState {
    /**
     * The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
     */
    administrativeUnitObjectId?: pulumi.Input<string>;
    /**
     * The object ID of the user, group or service principal you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
     */
    memberObjectId?: pulumi.Input<string>;
    /**
     * The object ID of the directory role you want to assign. Changing this forces a new resource to be created.
     */
    roleObjectId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AdministrativeUnitRoleMember resource.
 */
export interface AdministrativeUnitRoleMemberArgs {
    /**
     * The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
     */
    administrativeUnitObjectId: pulumi.Input<string>;
    /**
     * The object ID of the user, group or service principal you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
     */
    memberObjectId: pulumi.Input<string>;
    /**
     * The object ID of the directory role you want to assign. Changing this forces a new resource to be created.
     */
    roleObjectId: pulumi.Input<string>;
}

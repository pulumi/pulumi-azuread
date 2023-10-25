// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a single administrative unit membership within Azure Active Directory.
 *
 * > **Warning** Do not use this resource at the same time as the `members` property of the `azuread.AdministrativeUnit` resource for the same administrative unit. Doing so will cause a conflict and administrative unit members will be removed.
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
 * const exampleUser = azuread.getUser({
 *     userPrincipalName: "jdoe@hashicorp.com",
 * });
 * const exampleAdministrativeUnit = new azuread.AdministrativeUnit("exampleAdministrativeUnit", {displayName: "Example-AU"});
 * const exampleAdministrativeUnitMember = new azuread.AdministrativeUnitMember("exampleAdministrativeUnitMember", {
 *     administrativeUnitObjectId: exampleAdministrativeUnit.id,
 *     memberObjectId: exampleUser.then(exampleUser => exampleUser.id),
 * });
 * ```
 *
 * ## Import
 *
 * Administrative unit members can be imported using the object ID of the administrative unit and the object ID of the member, e.g.
 *
 * ```sh
 *  $ pulumi import azuread:index/administrativeUnitMember:AdministrativeUnitMember example 00000000-0000-0000-0000-000000000000/member/11111111-1111-1111-1111-111111111111
 * ```
 *
 *  -> This ID format is unique to Terraform and is composed of the Administrative Unit Object ID and the target Member Object ID in the format `{AdministrativeUnitObjectID}/member/{MemberObjectID}`.
 */
export class AdministrativeUnitMember extends pulumi.CustomResource {
    /**
     * Get an existing AdministrativeUnitMember resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AdministrativeUnitMemberState, opts?: pulumi.CustomResourceOptions): AdministrativeUnitMember {
        return new AdministrativeUnitMember(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/administrativeUnitMember:AdministrativeUnitMember';

    /**
     * Returns true if the given object is an instance of AdministrativeUnitMember.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AdministrativeUnitMember {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AdministrativeUnitMember.__pulumiType;
    }

    /**
     * The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
     */
    public readonly administrativeUnitObjectId!: pulumi.Output<string | undefined>;
    /**
     * The object ID of the user or group you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
     */
    public readonly memberObjectId!: pulumi.Output<string | undefined>;

    /**
     * Create a AdministrativeUnitMember resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AdministrativeUnitMemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AdministrativeUnitMemberArgs | AdministrativeUnitMemberState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AdministrativeUnitMemberState | undefined;
            resourceInputs["administrativeUnitObjectId"] = state ? state.administrativeUnitObjectId : undefined;
            resourceInputs["memberObjectId"] = state ? state.memberObjectId : undefined;
        } else {
            const args = argsOrState as AdministrativeUnitMemberArgs | undefined;
            resourceInputs["administrativeUnitObjectId"] = args ? args.administrativeUnitObjectId : undefined;
            resourceInputs["memberObjectId"] = args ? args.memberObjectId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AdministrativeUnitMember.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AdministrativeUnitMember resources.
 */
export interface AdministrativeUnitMemberState {
    /**
     * The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
     */
    administrativeUnitObjectId?: pulumi.Input<string>;
    /**
     * The object ID of the user or group you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
     */
    memberObjectId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AdministrativeUnitMember resource.
 */
export interface AdministrativeUnitMemberArgs {
    /**
     * The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
     */
    administrativeUnitObjectId?: pulumi.Input<string>;
    /**
     * The object ID of the user or group you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
     */
    memberObjectId?: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a single group membership within Azure Active Directory.
 *
 * > **Warning** Do not use this resource at the same time as the `members` property of the `azuread.Group` resource for the same group. Doing so will cause a conflict and group members will be removed.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this resource.
 *
 * When authenticated with a service principal, this resource requires one of the following application roles: `Group.ReadWrite.All` or `Directory.ReadWrite.All`
 *
 * When authenticated with a user principal, this resource requires one of the following directory roles: `Groups Administrator`, `User Administrator` or `Global Administrator`
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
 * const exampleGroup = new azuread.Group("exampleGroup", {
 *     displayName: "my_group",
 *     securityEnabled: true,
 * });
 * const exampleGroupMember = new azuread.GroupMember("exampleGroupMember", {
 *     groupObjectId: exampleGroup.id,
 *     memberObjectId: exampleUser.then(exampleUser => exampleUser.id),
 * });
 * ```
 *
 * ## Import
 *
 * Group members can be imported using the object ID of the group and the object ID of the member, e.g.
 *
 * ```sh
 *  $ pulumi import azuread:index/groupMember:GroupMember test 00000000-0000-0000-0000-000000000000/member/11111111-1111-1111-1111-111111111111
 * ```
 *
 *  -> This ID format is unique to Terraform and is composed of the Azure AD Group Object ID and the target Member Object ID in the format `{GroupObjectID}/member/{MemberObjectID}`.
 */
export class GroupMember extends pulumi.CustomResource {
    /**
     * Get an existing GroupMember resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupMemberState, opts?: pulumi.CustomResourceOptions): GroupMember {
        return new GroupMember(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/groupMember:GroupMember';

    /**
     * Returns true if the given object is an instance of GroupMember.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupMember {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupMember.__pulumiType;
    }

    /**
     * The object ID of the group you want to add the member to. Changing this forces a new resource to be created.
     */
    public readonly groupObjectId!: pulumi.Output<string>;
    /**
     * The object ID of the principal you want to add as a member to the group. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
     */
    public readonly memberObjectId!: pulumi.Output<string>;

    /**
     * Create a GroupMember resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupMemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupMemberArgs | GroupMemberState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupMemberState | undefined;
            resourceInputs["groupObjectId"] = state ? state.groupObjectId : undefined;
            resourceInputs["memberObjectId"] = state ? state.memberObjectId : undefined;
        } else {
            const args = argsOrState as GroupMemberArgs | undefined;
            if ((!args || args.groupObjectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupObjectId'");
            }
            if ((!args || args.memberObjectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'memberObjectId'");
            }
            resourceInputs["groupObjectId"] = args ? args.groupObjectId : undefined;
            resourceInputs["memberObjectId"] = args ? args.memberObjectId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(GroupMember.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupMember resources.
 */
export interface GroupMemberState {
    /**
     * The object ID of the group you want to add the member to. Changing this forces a new resource to be created.
     */
    groupObjectId?: pulumi.Input<string>;
    /**
     * The object ID of the principal you want to add as a member to the group. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
     */
    memberObjectId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GroupMember resource.
 */
export interface GroupMemberArgs {
    /**
     * The object ID of the group you want to add the member to. Changing this forces a new resource to be created.
     */
    groupObjectId: pulumi.Input<string>;
    /**
     * The object ID of the principal you want to add as a member to the group. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
     */
    memberObjectId: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a single directory role membership (assignment) within Azure Active Directory.
 *
 * > **Deprecation Warning:** This resource has been superseded by the azuread.DirectoryRoleAssignment resource and will be removed in version 3.0 of the AzureAD provider
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this resource.
 *
 * When authenticated with a service principal, this resource requires one of the following application roles: `RoleManagement.ReadWrite.Directory` or `Directory.ReadWrite.All`
 *
 * When authenticated with a user principal, this resource requires one of the following directory roles: `Privileged Role Administrator` or `Global Administrator`
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = azuread.getUser({
 *     userPrincipalName: "jdoe@example.com",
 * });
 * const exampleDirectoryRole = new azuread.DirectoryRole("example", {displayName: "Security administrator"});
 * const exampleDirectoryRoleMember = new azuread.DirectoryRoleMember("example", {
 *     roleObjectId: exampleDirectoryRole.objectId,
 *     memberObjectId: example.then(example => example.objectId),
 * });
 * ```
 *
 * ## Import
 *
 * Directory role members can be imported using the object ID of the role and the object ID of the member, e.g.
 *
 * ```sh
 * $ pulumi import azuread:index/directoryRoleMember:DirectoryRoleMember example 00000000-0000-0000-0000-000000000000/member/11111111-1111-1111-1111-111111111111
 * ```
 *
 * -> This ID format is unique to Terraform and is composed of the Directory Role Object ID and the target Member Object ID in the format `{RoleObjectID}/member/{MemberObjectID}`.
 */
export class DirectoryRoleMember extends pulumi.CustomResource {
    /**
     * Get an existing DirectoryRoleMember resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DirectoryRoleMemberState, opts?: pulumi.CustomResourceOptions): DirectoryRoleMember {
        return new DirectoryRoleMember(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/directoryRoleMember:DirectoryRoleMember';

    /**
     * Returns true if the given object is an instance of DirectoryRoleMember.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DirectoryRoleMember {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DirectoryRoleMember.__pulumiType;
    }

    /**
     * The object ID of the principal you want to add as a member to the directory role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
     */
    public readonly memberObjectId!: pulumi.Output<string | undefined>;
    /**
     * The object ID of the directory role you want to add the member to. Changing this forces a new resource to be created.
     */
    public readonly roleObjectId!: pulumi.Output<string | undefined>;

    /**
     * Create a DirectoryRoleMember resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DirectoryRoleMemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DirectoryRoleMemberArgs | DirectoryRoleMemberState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DirectoryRoleMemberState | undefined;
            resourceInputs["memberObjectId"] = state ? state.memberObjectId : undefined;
            resourceInputs["roleObjectId"] = state ? state.roleObjectId : undefined;
        } else {
            const args = argsOrState as DirectoryRoleMemberArgs | undefined;
            resourceInputs["memberObjectId"] = args ? args.memberObjectId : undefined;
            resourceInputs["roleObjectId"] = args ? args.roleObjectId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DirectoryRoleMember.__pulumiType, name, resourceInputs, opts, false /*remote*/);
    }
}

/**
 * Input properties used for looking up and filtering DirectoryRoleMember resources.
 */
export interface DirectoryRoleMemberState {
    /**
     * The object ID of the principal you want to add as a member to the directory role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
     */
    memberObjectId?: pulumi.Input<string>;
    /**
     * The object ID of the directory role you want to add the member to. Changing this forces a new resource to be created.
     */
    roleObjectId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DirectoryRoleMember resource.
 */
export interface DirectoryRoleMemberArgs {
    /**
     * The object ID of the principal you want to add as a member to the directory role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
     */
    memberObjectId?: pulumi.Input<string>;
    /**
     * The object ID of the directory role you want to add the member to. Changing this forces a new resource to be created.
     */
    roleObjectId?: pulumi.Input<string>;
}

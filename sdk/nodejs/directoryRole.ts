// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a Directory Role within Azure Active Directory. Directory Roles are also known as Administrator Roles.
 *
 * Directory Roles are built-in to Azure Active Directory and are immutable. However, by default they are not activated in a tenant (except for the Global Administrator role). This resource ensures a directory role is activated from its associated role template, and exports the object ID of the role, so that role assignments can be made for it.
 *
 * Once activated, directory roles cannot be deactivated and so this resource does not perform any actions on destroy.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this resource.
 *
 * When authenticated with a service principal, this resource requires one of the following application roles: `RoleManagement.ReadWrite.Directory` or `Directory.ReadWrite.All`
 *
 * When authenticated with a user principal, this resource requires one of the following directory roles: `Privileged Role Administrator` or `Global Administrator`
 *
 * ## Import
 *
 * This resource does not support importing.
 */
export class DirectoryRole extends pulumi.CustomResource {
    /**
     * Get an existing DirectoryRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DirectoryRoleState, opts?: pulumi.CustomResourceOptions): DirectoryRole {
        return new DirectoryRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/directoryRole:DirectoryRole';

    /**
     * Returns true if the given object is an instance of DirectoryRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DirectoryRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DirectoryRole.__pulumiType;
    }

    /**
     * The description of the directory role.
     */
    public /*out*/ readonly description!: pulumi.Output<string>;
    /**
     * The display name of the directory role to activate. Changing this forces a new resource to be created.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The object ID of the directory role.
     */
    public /*out*/ readonly objectId!: pulumi.Output<string>;
    /**
     * The object ID of the role template from which to activate the directory role. Changing this forces a new resource to be created.
     *
     * > Either `displayName` or `templateId` must be specified.
     */
    public readonly templateId!: pulumi.Output<string>;

    /**
     * Create a DirectoryRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DirectoryRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DirectoryRoleArgs | DirectoryRoleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DirectoryRoleState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["objectId"] = state ? state.objectId : undefined;
            resourceInputs["templateId"] = state ? state.templateId : undefined;
        } else {
            const args = argsOrState as DirectoryRoleArgs | undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["templateId"] = args ? args.templateId : undefined;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["objectId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DirectoryRole.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DirectoryRole resources.
 */
export interface DirectoryRoleState {
    /**
     * The description of the directory role.
     */
    description?: pulumi.Input<string>;
    /**
     * The display name of the directory role to activate. Changing this forces a new resource to be created.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The object ID of the directory role.
     */
    objectId?: pulumi.Input<string>;
    /**
     * The object ID of the role template from which to activate the directory role. Changing this forces a new resource to be created.
     *
     * > Either `displayName` or `templateId` must be specified.
     */
    templateId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DirectoryRole resource.
 */
export interface DirectoryRoleArgs {
    /**
     * The display name of the directory role to activate. Changing this forces a new resource to be created.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The object ID of the role template from which to activate the directory role. Changing this forces a new resource to be created.
     *
     * > Either `displayName` or `templateId` must be specified.
     */
    templateId?: pulumi.Input<string>;
}

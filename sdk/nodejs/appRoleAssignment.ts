// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages an app role assignment for a group, user or service principal. Can be used to grant admin consent for application permissions.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this resource.
 *
 * When authenticated with a service principal, this resource requires one of the following application roles: `AppRoleAssignment.ReadWrite.All` and `Application.Read.All`, or `AppRoleAssignment.ReadWrite.All` and `Directory.Read.All`, or `Application.ReadWrite.All`, or `Directory.ReadWrite.All`
 *
 * When authenticated with a user principal, this resource requires one of the following directory roles: `Application Administrator` or `Global Administrator`
 *
 * ## Import
 *
 * App role assignments can be imported using the object ID of the service principal representing the resource and the ID of the app role assignment (note: _not_ the ID of the app role), e.g.
 *
 * ```sh
 * $ pulumi import azuread:index/appRoleAssignment:AppRoleAssignment example 00000000-0000-0000-0000-000000000000/appRoleAssignment/aaBBcDDeFG6h5JKLMN2PQrrssTTUUvWWxxxxxyyyzzz
 * ```
 *
 *  -> This ID format is unique to Terraform and is composed of the Resource Service Principal Object ID and the ID of the App Role Assignment in the format `{ResourcePrincipalID}/appRoleAssignment/{AppRoleAssignmentID}`.
 */
export class AppRoleAssignment extends pulumi.CustomResource {
    /**
     * Get an existing AppRoleAssignment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AppRoleAssignmentState, opts?: pulumi.CustomResourceOptions): AppRoleAssignment {
        return new AppRoleAssignment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/appRoleAssignment:AppRoleAssignment';

    /**
     * Returns true if the given object is an instance of AppRoleAssignment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppRoleAssignment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppRoleAssignment.__pulumiType;
    }

    /**
     * The ID of the app role to be assigned, or the default role ID `00000000-0000-0000-0000-000000000000`. Changing this forces a new resource to be created.
     */
    public readonly appRoleId!: pulumi.Output<string>;
    /**
     * The display name of the principal to which the app role is assigned.
     */
    public /*out*/ readonly principalDisplayName!: pulumi.Output<string>;
    /**
     * The object ID of the user, group or service principal to be assigned this app role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
     */
    public readonly principalObjectId!: pulumi.Output<string>;
    /**
     * The object type of the principal to which the app role is assigned.
     */
    public /*out*/ readonly principalType!: pulumi.Output<string>;
    /**
     * The display name of the application representing the resource.
     */
    public /*out*/ readonly resourceDisplayName!: pulumi.Output<string>;
    /**
     * The object ID of the service principal representing the resource. Changing this forces a new resource to be created.
     */
    public readonly resourceObjectId!: pulumi.Output<string>;

    /**
     * Create a AppRoleAssignment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppRoleAssignmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AppRoleAssignmentArgs | AppRoleAssignmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AppRoleAssignmentState | undefined;
            resourceInputs["appRoleId"] = state ? state.appRoleId : undefined;
            resourceInputs["principalDisplayName"] = state ? state.principalDisplayName : undefined;
            resourceInputs["principalObjectId"] = state ? state.principalObjectId : undefined;
            resourceInputs["principalType"] = state ? state.principalType : undefined;
            resourceInputs["resourceDisplayName"] = state ? state.resourceDisplayName : undefined;
            resourceInputs["resourceObjectId"] = state ? state.resourceObjectId : undefined;
        } else {
            const args = argsOrState as AppRoleAssignmentArgs | undefined;
            if ((!args || args.appRoleId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'appRoleId'");
            }
            if ((!args || args.principalObjectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'principalObjectId'");
            }
            if ((!args || args.resourceObjectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceObjectId'");
            }
            resourceInputs["appRoleId"] = args ? args.appRoleId : undefined;
            resourceInputs["principalObjectId"] = args ? args.principalObjectId : undefined;
            resourceInputs["resourceObjectId"] = args ? args.resourceObjectId : undefined;
            resourceInputs["principalDisplayName"] = undefined /*out*/;
            resourceInputs["principalType"] = undefined /*out*/;
            resourceInputs["resourceDisplayName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AppRoleAssignment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AppRoleAssignment resources.
 */
export interface AppRoleAssignmentState {
    /**
     * The ID of the app role to be assigned, or the default role ID `00000000-0000-0000-0000-000000000000`. Changing this forces a new resource to be created.
     */
    appRoleId?: pulumi.Input<string>;
    /**
     * The display name of the principal to which the app role is assigned.
     */
    principalDisplayName?: pulumi.Input<string>;
    /**
     * The object ID of the user, group or service principal to be assigned this app role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
     */
    principalObjectId?: pulumi.Input<string>;
    /**
     * The object type of the principal to which the app role is assigned.
     */
    principalType?: pulumi.Input<string>;
    /**
     * The display name of the application representing the resource.
     */
    resourceDisplayName?: pulumi.Input<string>;
    /**
     * The object ID of the service principal representing the resource. Changing this forces a new resource to be created.
     */
    resourceObjectId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AppRoleAssignment resource.
 */
export interface AppRoleAssignmentArgs {
    /**
     * The ID of the app role to be assigned, or the default role ID `00000000-0000-0000-0000-000000000000`. Changing this forces a new resource to be created.
     */
    appRoleId: pulumi.Input<string>;
    /**
     * The object ID of the user, group or service principal to be assigned this app role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
     */
    principalObjectId: pulumi.Input<string>;
    /**
     * The object ID of the service principal representing the resource. Changing this forces a new resource to be created.
     */
    resourceObjectId: pulumi.Input<string>;
}

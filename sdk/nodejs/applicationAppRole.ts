// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages an App Role associated with an Application within Azure Active Directory.
 *
 * > **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write all applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const exampleApplication = new azuread.Application("exampleApplication", {});
 * const exampleApplicationAppRole = new azuread.ApplicationAppRole("exampleApplicationAppRole", {
 *     applicationObjectId: exampleApplication.id,
 *     allowedMemberTypes: ["User"],
 *     description: "Admins can manage roles and perform all task actions",
 *     displayName: "Admin",
 *     enabled: true,
 *     value: "administer",
 * });
 * ```
 *
 * ## Import
 *
 * App Roles can be imported using the `object_id` of an Application and the `id` of the App Role, e.g.
 *
 * ```sh
 *  $ pulumi import azuread:index/applicationAppRole:ApplicationAppRole test 00000000-0000-0000-0000-000000000000/role/11111111-1111-1111-1111-111111111111
 * ```
 */
export class ApplicationAppRole extends pulumi.CustomResource {
    /**
     * Get an existing ApplicationAppRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApplicationAppRoleState, opts?: pulumi.CustomResourceOptions): ApplicationAppRole {
        return new ApplicationAppRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/applicationAppRole:ApplicationAppRole';

    /**
     * Returns true if the given object is an instance of ApplicationAppRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApplicationAppRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApplicationAppRole.__pulumiType;
    }

    /**
     * Specifies whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications (that are accessing this application in a standalone scenario) by setting to `Application`, or to both.
     */
    public readonly allowedMemberTypes!: pulumi.Output<string[]>;
    /**
     * The Object ID of the Application for which this App Role should be created. Changing this field forces a new resource to be created.
     */
    public readonly applicationObjectId!: pulumi.Output<string>;
    /**
     * Description of the app role that appears when the role is being assigned and, if the role functions as an application permissions, during the consent experiences.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Display name for the app role that appears during app role assignment and in consent experiences.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Determines if the app role is enabled: Defaults to `true`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * @deprecated [NOTE] This attribute has been renamed to `enabled` and will be removed in version 2.0 of the AzureAD provider
     */
    public readonly isEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The unique identifier for the app role. If omitted, a random UUID will be automatically generated. Must be a valid UUID. Changing this field forces a new resource to be created.
     */
    public readonly roleId!: pulumi.Output<string>;
    /**
     * The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal.
     */
    public readonly value!: pulumi.Output<string | undefined>;

    /**
     * Create a ApplicationAppRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationAppRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationAppRoleArgs | ApplicationAppRoleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApplicationAppRoleState | undefined;
            inputs["allowedMemberTypes"] = state ? state.allowedMemberTypes : undefined;
            inputs["applicationObjectId"] = state ? state.applicationObjectId : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["isEnabled"] = state ? state.isEnabled : undefined;
            inputs["roleId"] = state ? state.roleId : undefined;
            inputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as ApplicationAppRoleArgs | undefined;
            if ((!args || args.allowedMemberTypes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'allowedMemberTypes'");
            }
            if ((!args || args.applicationObjectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationObjectId'");
            }
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            inputs["allowedMemberTypes"] = args ? args.allowedMemberTypes : undefined;
            inputs["applicationObjectId"] = args ? args.applicationObjectId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["isEnabled"] = args ? args.isEnabled : undefined;
            inputs["roleId"] = args ? args.roleId : undefined;
            inputs["value"] = args ? args.value : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ApplicationAppRole.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApplicationAppRole resources.
 */
export interface ApplicationAppRoleState {
    /**
     * Specifies whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications (that are accessing this application in a standalone scenario) by setting to `Application`, or to both.
     */
    readonly allowedMemberTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Object ID of the Application for which this App Role should be created. Changing this field forces a new resource to be created.
     */
    readonly applicationObjectId?: pulumi.Input<string>;
    /**
     * Description of the app role that appears when the role is being assigned and, if the role functions as an application permissions, during the consent experiences.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Display name for the app role that appears during app role assignment and in consent experiences.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * Determines if the app role is enabled: Defaults to `true`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * @deprecated [NOTE] This attribute has been renamed to `enabled` and will be removed in version 2.0 of the AzureAD provider
     */
    readonly isEnabled?: pulumi.Input<boolean>;
    /**
     * The unique identifier for the app role. If omitted, a random UUID will be automatically generated. Must be a valid UUID. Changing this field forces a new resource to be created.
     */
    readonly roleId?: pulumi.Input<string>;
    /**
     * The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal.
     */
    readonly value?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ApplicationAppRole resource.
 */
export interface ApplicationAppRoleArgs {
    /**
     * Specifies whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications (that are accessing this application in a standalone scenario) by setting to `Application`, or to both.
     */
    readonly allowedMemberTypes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Object ID of the Application for which this App Role should be created. Changing this field forces a new resource to be created.
     */
    readonly applicationObjectId: pulumi.Input<string>;
    /**
     * Description of the app role that appears when the role is being assigned and, if the role functions as an application permissions, during the consent experiences.
     */
    readonly description: pulumi.Input<string>;
    /**
     * Display name for the app role that appears during app role assignment and in consent experiences.
     */
    readonly displayName: pulumi.Input<string>;
    /**
     * Determines if the app role is enabled: Defaults to `true`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * @deprecated [NOTE] This attribute has been renamed to `enabled` and will be removed in version 2.0 of the AzureAD provider
     */
    readonly isEnabled?: pulumi.Input<boolean>;
    /**
     * The unique identifier for the app role. If omitted, a random UUID will be automatically generated. Must be a valid UUID. Changing this field forces a new resource to be created.
     */
    readonly roleId?: pulumi.Input<string>;
    /**
     * The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal.
     */
    readonly value?: pulumi.Input<string>;
}

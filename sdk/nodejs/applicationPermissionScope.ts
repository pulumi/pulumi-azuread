// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 * import * as random from "@pulumi/random";
 *
 * const example = new azuread.ApplicationRegistration("example", {displayName: "example"});
 * const exampleAdminister = new random.RandomUuid("example_administer", {});
 * const exampleApplicationPermissionScope = new azuread.ApplicationPermissionScope("example", {
 *     applicationId: test.id,
 *     scopeId: exampleAdminister.id,
 *     value: "administer",
 *     adminConsentDescription: "Administer the application",
 *     adminConsentDisplayName: "Administer",
 * });
 * ```
 *
 * > **Tip** For managing more permissions scopes, create additional instances of this resource
 *
 * *Usage with azuread.Application resource*
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = new azuread.Application("example", {displayName: "example"});
 * const exampleApplicationPermissionScope = new azuread.ApplicationPermissionScope("example", {applicationId: example.id});
 * ```
 *
 * ## Import
 *
 * Application App Roles can be imported using the object ID of the application and the ID of the permission scope, in the following format.
 *
 * ```sh
 * $ pulumi import azuread:index/applicationPermissionScope:ApplicationPermissionScope example /applications/00000000-0000-0000-0000-000000000000/permissionScopes/11111111-1111-1111-1111-111111111111
 * ```
 */
export class ApplicationPermissionScope extends pulumi.CustomResource {
    /**
     * Get an existing ApplicationPermissionScope resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApplicationPermissionScopeState, opts?: pulumi.CustomResourceOptions): ApplicationPermissionScope {
        return new ApplicationPermissionScope(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/applicationPermissionScope:ApplicationPermissionScope';

    /**
     * Returns true if the given object is an instance of ApplicationPermissionScope.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApplicationPermissionScope {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApplicationPermissionScope.__pulumiType;
    }

    /**
     * Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
     */
    public readonly adminConsentDescription!: pulumi.Output<string>;
    /**
     * Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
     */
    public readonly adminConsentDisplayName!: pulumi.Output<string>;
    /**
     * The resource ID of the application registration. Changing this forces a new resource to be created.
     */
    public readonly applicationId!: pulumi.Output<string>;
    /**
     * The unique identifier of the permission scope. Must be a valid UUID. Changing this forces a new resource to be created.
     */
    public readonly scopeId!: pulumi.Output<string>;
    /**
     * Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions.
     */
    public readonly type!: pulumi.Output<string | undefined>;
    /**
     * Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
     */
    public readonly userConsentDescription!: pulumi.Output<string | undefined>;
    /**
     * Display name for the delegated permission that appears in the end user consent experience
     */
    public readonly userConsentDisplayName!: pulumi.Output<string | undefined>;
    /**
     * The value that is used for the `scp` claim in OAuth access tokens.
     *
     * > **Roles and Permission Scopes** In Azure Active Directory, application roles and permission scopes exported by an application share the same namespace and cannot contain duplicate values.
     */
    public readonly value!: pulumi.Output<string>;

    /**
     * Create a ApplicationPermissionScope resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationPermissionScopeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationPermissionScopeArgs | ApplicationPermissionScopeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApplicationPermissionScopeState | undefined;
            resourceInputs["adminConsentDescription"] = state ? state.adminConsentDescription : undefined;
            resourceInputs["adminConsentDisplayName"] = state ? state.adminConsentDisplayName : undefined;
            resourceInputs["applicationId"] = state ? state.applicationId : undefined;
            resourceInputs["scopeId"] = state ? state.scopeId : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["userConsentDescription"] = state ? state.userConsentDescription : undefined;
            resourceInputs["userConsentDisplayName"] = state ? state.userConsentDisplayName : undefined;
            resourceInputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as ApplicationPermissionScopeArgs | undefined;
            if ((!args || args.adminConsentDescription === undefined) && !opts.urn) {
                throw new Error("Missing required property 'adminConsentDescription'");
            }
            if ((!args || args.adminConsentDisplayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'adminConsentDisplayName'");
            }
            if ((!args || args.applicationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationId'");
            }
            if ((!args || args.scopeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scopeId'");
            }
            if ((!args || args.value === undefined) && !opts.urn) {
                throw new Error("Missing required property 'value'");
            }
            resourceInputs["adminConsentDescription"] = args ? args.adminConsentDescription : undefined;
            resourceInputs["adminConsentDisplayName"] = args ? args.adminConsentDisplayName : undefined;
            resourceInputs["applicationId"] = args ? args.applicationId : undefined;
            resourceInputs["scopeId"] = args ? args.scopeId : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["userConsentDescription"] = args ? args.userConsentDescription : undefined;
            resourceInputs["userConsentDisplayName"] = args ? args.userConsentDisplayName : undefined;
            resourceInputs["value"] = args ? args.value : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ApplicationPermissionScope.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApplicationPermissionScope resources.
 */
export interface ApplicationPermissionScopeState {
    /**
     * Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
     */
    adminConsentDescription?: pulumi.Input<string>;
    /**
     * Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
     */
    adminConsentDisplayName?: pulumi.Input<string>;
    /**
     * The resource ID of the application registration. Changing this forces a new resource to be created.
     */
    applicationId?: pulumi.Input<string>;
    /**
     * The unique identifier of the permission scope. Must be a valid UUID. Changing this forces a new resource to be created.
     */
    scopeId?: pulumi.Input<string>;
    /**
     * Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions.
     */
    type?: pulumi.Input<string>;
    /**
     * Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
     */
    userConsentDescription?: pulumi.Input<string>;
    /**
     * Display name for the delegated permission that appears in the end user consent experience
     */
    userConsentDisplayName?: pulumi.Input<string>;
    /**
     * The value that is used for the `scp` claim in OAuth access tokens.
     *
     * > **Roles and Permission Scopes** In Azure Active Directory, application roles and permission scopes exported by an application share the same namespace and cannot contain duplicate values.
     */
    value?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ApplicationPermissionScope resource.
 */
export interface ApplicationPermissionScopeArgs {
    /**
     * Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
     */
    adminConsentDescription: pulumi.Input<string>;
    /**
     * Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
     */
    adminConsentDisplayName: pulumi.Input<string>;
    /**
     * The resource ID of the application registration. Changing this forces a new resource to be created.
     */
    applicationId: pulumi.Input<string>;
    /**
     * The unique identifier of the permission scope. Must be a valid UUID. Changing this forces a new resource to be created.
     */
    scopeId: pulumi.Input<string>;
    /**
     * Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions.
     */
    type?: pulumi.Input<string>;
    /**
     * Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
     */
    userConsentDescription?: pulumi.Input<string>;
    /**
     * Display name for the delegated permission that appears in the end user consent experience
     */
    userConsentDisplayName?: pulumi.Input<string>;
    /**
     * The value that is used for the `scp` claim in OAuth access tokens.
     *
     * > **Roles and Permission Scopes** In Azure Active Directory, application roles and permission scopes exported by an application share the same namespace and cannot contain duplicate values.
     */
    value: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a single catalog role assignment within Azure Active Directory.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this resource.
 *
 * When authenticated with a service principal, this resource requires one of the following application roles: `EntitlementManagement.ReadWrite.All` or `Directory.ReadWrite.All`
 *
 * When authenticated with a user principal, this resource requires one of the following directory roles: `Identity Governance administrator` or `Global Administrator`
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
 * const exampleAccessPackageCatalogRole = azuread.getAccessPackageCatalogRole({
 *     displayName: "Catalog owner",
 * });
 * const exampleAccessPackageCatalog = new azuread.AccessPackageCatalog("exampleAccessPackageCatalog", {
 *     displayName: "example-access-package-catalog",
 *     description: "Example access package catalog",
 * });
 * const exampleAccessPackageCatalogRoleAssignment = new azuread.AccessPackageCatalogRoleAssignment("exampleAccessPackageCatalogRoleAssignment", {
 *     roleId: exampleAccessPackageCatalogRole.then(exampleAccessPackageCatalogRole => exampleAccessPackageCatalogRole.objectId),
 *     principalObjectId: exampleUser.then(exampleUser => exampleUser.objectId),
 *     catalogId: exampleAccessPackageCatalog.id,
 * });
 * ```
 *
 * ## Import
 *
 * Catalog role assignments can be imported using the ID of the assignment, e.g.
 *
 * ```sh
 *  $ pulumi import azuread:index/accessPackageCatalogRoleAssignment:AccessPackageCatalogRoleAssignment test 00000000-0000-0000-0000-000000000000
 * ```
 */
export class AccessPackageCatalogRoleAssignment extends pulumi.CustomResource {
    /**
     * Get an existing AccessPackageCatalogRoleAssignment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccessPackageCatalogRoleAssignmentState, opts?: pulumi.CustomResourceOptions): AccessPackageCatalogRoleAssignment {
        return new AccessPackageCatalogRoleAssignment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/accessPackageCatalogRoleAssignment:AccessPackageCatalogRoleAssignment';

    /**
     * Returns true if the given object is an instance of AccessPackageCatalogRoleAssignment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccessPackageCatalogRoleAssignment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccessPackageCatalogRoleAssignment.__pulumiType;
    }

    /**
     * The ID of the Catalog this role assignment will be scoped to. Changing this forces a new resource to be created.
     */
    public readonly catalogId!: pulumi.Output<string>;
    /**
     * The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
     */
    public readonly principalObjectId!: pulumi.Output<string>;
    /**
     * The object ID of the catalog role you want to assign. Changing this forces a new resource to be created.
     */
    public readonly roleId!: pulumi.Output<string>;

    /**
     * Create a AccessPackageCatalogRoleAssignment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccessPackageCatalogRoleAssignmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccessPackageCatalogRoleAssignmentArgs | AccessPackageCatalogRoleAssignmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AccessPackageCatalogRoleAssignmentState | undefined;
            resourceInputs["catalogId"] = state ? state.catalogId : undefined;
            resourceInputs["principalObjectId"] = state ? state.principalObjectId : undefined;
            resourceInputs["roleId"] = state ? state.roleId : undefined;
        } else {
            const args = argsOrState as AccessPackageCatalogRoleAssignmentArgs | undefined;
            if ((!args || args.catalogId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'catalogId'");
            }
            if ((!args || args.principalObjectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'principalObjectId'");
            }
            if ((!args || args.roleId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleId'");
            }
            resourceInputs["catalogId"] = args ? args.catalogId : undefined;
            resourceInputs["principalObjectId"] = args ? args.principalObjectId : undefined;
            resourceInputs["roleId"] = args ? args.roleId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AccessPackageCatalogRoleAssignment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AccessPackageCatalogRoleAssignment resources.
 */
export interface AccessPackageCatalogRoleAssignmentState {
    /**
     * The ID of the Catalog this role assignment will be scoped to. Changing this forces a new resource to be created.
     */
    catalogId?: pulumi.Input<string>;
    /**
     * The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
     */
    principalObjectId?: pulumi.Input<string>;
    /**
     * The object ID of the catalog role you want to assign. Changing this forces a new resource to be created.
     */
    roleId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AccessPackageCatalogRoleAssignment resource.
 */
export interface AccessPackageCatalogRoleAssignmentArgs {
    /**
     * The ID of the Catalog this role assignment will be scoped to. Changing this forces a new resource to be created.
     */
    catalogId: pulumi.Input<string>;
    /**
     * The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
     */
    principalObjectId: pulumi.Input<string>;
    /**
     * The object ID of the catalog role you want to assign. Changing this forces a new resource to be created.
     */
    roleId: pulumi.Input<string>;
}

// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages the resources added to access packages within Identity Governance in Azure Active Directory.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this resource.
 *
 * When authenticated with a service principal, this resource requires the following application role: `EntitlementManagement.ReadWrite.All`.
 *
 * When authenticated with a user principal, this resource requires one of the following directory roles: `Catalog owner`, `Access package manager` or `Global Administrator`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = new azuread.Group("example", {
 *     displayName: "example-group",
 *     securityEnabled: true,
 * });
 * const exampleAccessPackageCatalog = new azuread.AccessPackageCatalog("example", {
 *     displayName: "example-catalog",
 *     description: "Example catalog",
 * });
 * const exampleAccessPackageResourceCatalogAssociation = new azuread.AccessPackageResourceCatalogAssociation("example", {
 *     catalogId: exampleCatalog.id,
 *     resourceOriginId: exampleGroup.objectId,
 *     resourceOriginSystem: "AadGroup",
 * });
 * const exampleAccessPackage = new azuread.AccessPackage("example", {
 *     displayName: "example-package",
 *     description: "Example Package",
 *     catalogId: exampleCatalog.id,
 * });
 * const exampleAccessPackageResourcePackageAssociation = new azuread.AccessPackageResourcePackageAssociation("example", {
 *     accessPackageId: exampleAccessPackage.id,
 *     catalogResourceAssociationId: exampleAccessPackageResourceCatalogAssociation.id,
 * });
 * ```
 *
 * ## Import
 *
 * The resource and catalog association can be imported using the access package ID, the access package ResourceRoleScope, the resource origin ID, and the access type, e.g.
 *
 * ```sh
 * $ pulumi import azuread:index/accessPackageResourcePackageAssociation:AccessPackageResourcePackageAssociation example 00000000-0000-0000-0000-000000000000/11111111-1111-1111-1111-111111111111_22222222-2222-2222-2222-22222222/33333333-3333-3333-3333-33333333/Member
 * ```
 *
 * -> This ID format is unique to Terraform and is composed of the Access Package ID, the access package ResourceRoleScope (in the format Role_Scope), the Resource Origin ID, and the Access Type, in the format `{AccessPackageID}/{ResourceRoleScope}/{ResourceOriginID}/{AccessType}`.
 */
export class AccessPackageResourcePackageAssociation extends pulumi.CustomResource {
    /**
     * Get an existing AccessPackageResourcePackageAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccessPackageResourcePackageAssociationState, opts?: pulumi.CustomResourceOptions): AccessPackageResourcePackageAssociation {
        return new AccessPackageResourcePackageAssociation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/accessPackageResourcePackageAssociation:AccessPackageResourcePackageAssociation';

    /**
     * Returns true if the given object is an instance of AccessPackageResourcePackageAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccessPackageResourcePackageAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccessPackageResourcePackageAssociation.__pulumiType;
    }

    /**
     * The ID of access package this resource association is configured to. Changing this forces a new resource to be created.
     */
    public readonly accessPackageId!: pulumi.Output<string>;
    /**
     * The role of access type to the specified resource. Valid values are `Member`, or `Owner` The default is `Member`. Changing this forces a new resource to be created.
     */
    public readonly accessType!: pulumi.Output<string | undefined>;
    /**
     * The ID of the catalog association from the `azuread.AccessPackageResourceCatalogAssociation` resource. Changing this forces a new resource to be created.
     */
    public readonly catalogResourceAssociationId!: pulumi.Output<string>;

    /**
     * Create a AccessPackageResourcePackageAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccessPackageResourcePackageAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccessPackageResourcePackageAssociationArgs | AccessPackageResourcePackageAssociationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AccessPackageResourcePackageAssociationState | undefined;
            resourceInputs["accessPackageId"] = state ? state.accessPackageId : undefined;
            resourceInputs["accessType"] = state ? state.accessType : undefined;
            resourceInputs["catalogResourceAssociationId"] = state ? state.catalogResourceAssociationId : undefined;
        } else {
            const args = argsOrState as AccessPackageResourcePackageAssociationArgs | undefined;
            if ((!args || args.accessPackageId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accessPackageId'");
            }
            if ((!args || args.catalogResourceAssociationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'catalogResourceAssociationId'");
            }
            resourceInputs["accessPackageId"] = args ? args.accessPackageId : undefined;
            resourceInputs["accessType"] = args ? args.accessType : undefined;
            resourceInputs["catalogResourceAssociationId"] = args ? args.catalogResourceAssociationId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AccessPackageResourcePackageAssociation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AccessPackageResourcePackageAssociation resources.
 */
export interface AccessPackageResourcePackageAssociationState {
    /**
     * The ID of access package this resource association is configured to. Changing this forces a new resource to be created.
     */
    accessPackageId?: pulumi.Input<string>;
    /**
     * The role of access type to the specified resource. Valid values are `Member`, or `Owner` The default is `Member`. Changing this forces a new resource to be created.
     */
    accessType?: pulumi.Input<string>;
    /**
     * The ID of the catalog association from the `azuread.AccessPackageResourceCatalogAssociation` resource. Changing this forces a new resource to be created.
     */
    catalogResourceAssociationId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AccessPackageResourcePackageAssociation resource.
 */
export interface AccessPackageResourcePackageAssociationArgs {
    /**
     * The ID of access package this resource association is configured to. Changing this forces a new resource to be created.
     */
    accessPackageId: pulumi.Input<string>;
    /**
     * The role of access type to the specified resource. Valid values are `Member`, or `Owner` The default is `Member`. Changing this forces a new resource to be created.
     */
    accessType?: pulumi.Input<string>;
    /**
     * The ID of the catalog association from the `azuread.AccessPackageResourceCatalogAssociation` resource. Changing this forces a new resource to be created.
     */
    catalogResourceAssociationId: pulumi.Input<string>;
}

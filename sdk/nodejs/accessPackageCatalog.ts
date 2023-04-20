// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages an access package catalog within Identity Governance in Azure Active Directory.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this resource.
 *
 * When authenticated with a service principal, this resource requires the following application role: `EntitlementManagement.ReadWrite.All`.
 *
 * When authenticated with a user principal, this resource requires one of the following directory roles: `Catalog owner`, `Catalog creator` or `Global Administrator`
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = new azuread.AccessPackageCatalog("example", {
 *     description: "Example access package catalog",
 *     displayName: "example-access-package-catalog",
 * });
 * ```
 *
 * ## Import
 *
 * An Access Package Catalog can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import azuread:index/accessPackageCatalog:AccessPackageCatalog example 00000000-0000-0000-0000-000000000000
 * ```
 */
export class AccessPackageCatalog extends pulumi.CustomResource {
    /**
     * Get an existing AccessPackageCatalog resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccessPackageCatalogState, opts?: pulumi.CustomResourceOptions): AccessPackageCatalog {
        return new AccessPackageCatalog(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/accessPackageCatalog:AccessPackageCatalog';

    /**
     * Returns true if the given object is an instance of AccessPackageCatalog.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccessPackageCatalog {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccessPackageCatalog.__pulumiType;
    }

    /**
     * The description of the access package catalog.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The display name of the access package catalog.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Whether the access packages in this catalog can be requested by users outside the tenant.
     */
    public readonly externallyVisible!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the access packages in this catalog are available for management.
     */
    public readonly published!: pulumi.Output<boolean | undefined>;

    /**
     * Create a AccessPackageCatalog resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccessPackageCatalogArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccessPackageCatalogArgs | AccessPackageCatalogState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AccessPackageCatalogState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["externallyVisible"] = state ? state.externallyVisible : undefined;
            resourceInputs["published"] = state ? state.published : undefined;
        } else {
            const args = argsOrState as AccessPackageCatalogArgs | undefined;
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["externallyVisible"] = args ? args.externallyVisible : undefined;
            resourceInputs["published"] = args ? args.published : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AccessPackageCatalog.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AccessPackageCatalog resources.
 */
export interface AccessPackageCatalogState {
    /**
     * The description of the access package catalog.
     */
    description?: pulumi.Input<string>;
    /**
     * The display name of the access package catalog.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Whether the access packages in this catalog can be requested by users outside the tenant.
     */
    externallyVisible?: pulumi.Input<boolean>;
    /**
     * Whether the access packages in this catalog are available for management.
     */
    published?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a AccessPackageCatalog resource.
 */
export interface AccessPackageCatalogArgs {
    /**
     * The description of the access package catalog.
     */
    description: pulumi.Input<string>;
    /**
     * The display name of the access package catalog.
     */
    displayName: pulumi.Input<string>;
    /**
     * Whether the access packages in this catalog can be requested by users outside the tenant.
     */
    externallyVisible?: pulumi.Input<boolean>;
    /**
     * Whether the access packages in this catalog are available for management.
     */
    published?: pulumi.Input<boolean>;
}

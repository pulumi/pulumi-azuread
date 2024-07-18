// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages an Access Package within Identity Governance in Azure Active Directory.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this resource.
 *
 * When authenticated with a service principal, this resource requires the following application role: `EntitlementManagement.ReadWrite.All`.
 *
 * When authenticated with a user principal, this resource requires one of the following directory roles: `Catalog owner`, `Access package manager` or `Global Administrator`
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = new azuread.AccessPackageCatalog("example", {
 *     displayName: "example-catalog",
 *     description: "Example catalog",
 * });
 * const exampleAccessPackage = new azuread.AccessPackage("example", {
 *     catalogId: example.id,
 *     displayName: "access-package",
 *     description: "Access Package",
 * });
 * ```
 *
 * ## Import
 *
 * Access Packages can be imported using the `id`, e.g.
 *
 * ```sh
 * $ pulumi import azuread:index/accessPackage:AccessPackage example_package 00000000-0000-0000-0000-000000000000
 * ```
 */
export class AccessPackage extends pulumi.CustomResource {
    /**
     * Get an existing AccessPackage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccessPackageState, opts?: pulumi.CustomResourceOptions): AccessPackage {
        return new AccessPackage(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/accessPackage:AccessPackage';

    /**
     * Returns true if the given object is an instance of AccessPackage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccessPackage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccessPackage.__pulumiType;
    }

    /**
     * The ID of the Catalog this access package will be created in.
     */
    public readonly catalogId!: pulumi.Output<string>;
    /**
     * The description of the access package.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The display name of the access package.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Whether the access package is hidden from the requestor.
     */
    public readonly hidden!: pulumi.Output<boolean | undefined>;

    /**
     * Create a AccessPackage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccessPackageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccessPackageArgs | AccessPackageState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AccessPackageState | undefined;
            resourceInputs["catalogId"] = state ? state.catalogId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["hidden"] = state ? state.hidden : undefined;
        } else {
            const args = argsOrState as AccessPackageArgs | undefined;
            if ((!args || args.catalogId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'catalogId'");
            }
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            resourceInputs["catalogId"] = args ? args.catalogId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["hidden"] = args ? args.hidden : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AccessPackage.__pulumiType, name, resourceInputs, opts, false /*remote*/);
    }
}

/**
 * Input properties used for looking up and filtering AccessPackage resources.
 */
export interface AccessPackageState {
    /**
     * The ID of the Catalog this access package will be created in.
     */
    catalogId?: pulumi.Input<string>;
    /**
     * The description of the access package.
     */
    description?: pulumi.Input<string>;
    /**
     * The display name of the access package.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Whether the access package is hidden from the requestor.
     */
    hidden?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a AccessPackage resource.
 */
export interface AccessPackageArgs {
    /**
     * The ID of the Catalog this access package will be created in.
     */
    catalogId: pulumi.Input<string>;
    /**
     * The description of the access package.
     */
    description: pulumi.Input<string>;
    /**
     * The display name of the access package.
     */
    displayName: pulumi.Input<string>;
    /**
     * Whether the access package is hidden from the requestor.
     */
    hidden?: pulumi.Input<boolean>;
}

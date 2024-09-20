// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Gets information about an access package catalog role.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this data source.
 *
 * When authenticated with a service principal, this data source requires one of the following application roles: `EntitlementManagement.Read.All` or `Directory.Read.All`
 *
 * When authenticated with a user principal, this data source does not require any additional roles.
 *
 * ## Example Usage
 *
 * ### By Group Display Name)
 *
 * *Look up by display name*
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = azuread.getAccessPackageCatalogRole({
 *     displayName: "Catalog owner",
 * });
 * ```
 *
 * *Look up by object ID*
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = azuread.getAccessPackageCatalogRole({
 *     objectId: "00000000-0000-0000-0000-000000000000",
 * });
 * ```
 */
export function getAccessPackageCatalogRole(args?: GetAccessPackageCatalogRoleArgs, opts?: pulumi.InvokeOptions): Promise<GetAccessPackageCatalogRoleResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azuread:index/getAccessPackageCatalogRole:getAccessPackageCatalogRole", {
        "displayName": args.displayName,
        "objectId": args.objectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAccessPackageCatalogRole.
 */
export interface GetAccessPackageCatalogRoleArgs {
    /**
     * Specifies the display name of the role.
     */
    displayName?: string;
    /**
     * Specifies the object ID of the role.
     *
     * > One of `displayName` or `objectId` must be specified.
     */
    objectId?: string;
}

/**
 * A collection of values returned by getAccessPackageCatalogRole.
 */
export interface GetAccessPackageCatalogRoleResult {
    /**
     * The description of the role.
     */
    readonly description: string;
    /**
     * The display name of the role.
     */
    readonly displayName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The object ID of the role.
     */
    readonly objectId: string;
    /**
     * The object ID of the role.
     */
    readonly templateId: string;
}
/**
 * Gets information about an access package catalog role.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this data source.
 *
 * When authenticated with a service principal, this data source requires one of the following application roles: `EntitlementManagement.Read.All` or `Directory.Read.All`
 *
 * When authenticated with a user principal, this data source does not require any additional roles.
 *
 * ## Example Usage
 *
 * ### By Group Display Name)
 *
 * *Look up by display name*
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = azuread.getAccessPackageCatalogRole({
 *     displayName: "Catalog owner",
 * });
 * ```
 *
 * *Look up by object ID*
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = azuread.getAccessPackageCatalogRole({
 *     objectId: "00000000-0000-0000-0000-000000000000",
 * });
 * ```
 */
export function getAccessPackageCatalogRoleOutput(args?: GetAccessPackageCatalogRoleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAccessPackageCatalogRoleResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azuread:index/getAccessPackageCatalogRole:getAccessPackageCatalogRole", {
        "displayName": args.displayName,
        "objectId": args.objectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAccessPackageCatalogRole.
 */
export interface GetAccessPackageCatalogRoleOutputArgs {
    /**
     * Specifies the display name of the role.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Specifies the object ID of the role.
     *
     * > One of `displayName` or `objectId` must be specified.
     */
    objectId?: pulumi.Input<string>;
}

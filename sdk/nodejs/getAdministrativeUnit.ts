// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Gets information about an adminisrative unit in Azure Active Directory.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this data source.
 *
 * When authenticated with a service principal, this data source requires one of the following application roles: `AdministrativeUnit.Read.All` or `Directory.Read.All`
 *
 * When authenticated with a user principal, this data source does not require any additional roles.
 *
 * ## Example Usage
 *
 * ### by Group Display Name)
 *
 * *Look up by display name*
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = azuread.getAdministrativeUnit({
 *     displayName: "Example-AU",
 * });
 * ```
 *
 * *Look up by object ID*
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = azuread.getAdministrativeUnit({
 *     objectId: "00000000-0000-0000-0000-000000000000",
 * });
 * ```
 */
export function getAdministrativeUnit(args?: GetAdministrativeUnitArgs, opts?: pulumi.InvokeOptions): Promise<GetAdministrativeUnitResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azuread:index/getAdministrativeUnit:getAdministrativeUnit", {
        "displayName": args.displayName,
        "objectId": args.objectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAdministrativeUnit.
 */
export interface GetAdministrativeUnitArgs {
    /**
     * Specifies the display name of the administrative unit.
     */
    displayName?: string;
    /**
     * Specifies the object ID of the administrative unit.
     *
     * > One of `displayName` or `objectId` must be specified.
     */
    objectId?: string;
}

/**
 * A collection of values returned by getAdministrativeUnit.
 */
export interface GetAdministrativeUnitResult {
    /**
     * The description of the administrative unit.
     */
    readonly description: string;
    /**
     * The display name of the administrative unit.
     */
    readonly displayName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of object IDs of members who are present in this administrative unit.
     */
    readonly members: string[];
    /**
     * The object ID of the administrative unit.
     */
    readonly objectId: string;
    /**
     * Whether the administrative unit *and* its members are hidden or publicly viewable in the directory. One of: `Hiddenmembership` or `Public`.
     */
    readonly visibility: string;
}
/**
 * Gets information about an adminisrative unit in Azure Active Directory.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this data source.
 *
 * When authenticated with a service principal, this data source requires one of the following application roles: `AdministrativeUnit.Read.All` or `Directory.Read.All`
 *
 * When authenticated with a user principal, this data source does not require any additional roles.
 *
 * ## Example Usage
 *
 * ### by Group Display Name)
 *
 * *Look up by display name*
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = azuread.getAdministrativeUnit({
 *     displayName: "Example-AU",
 * });
 * ```
 *
 * *Look up by object ID*
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = azuread.getAdministrativeUnit({
 *     objectId: "00000000-0000-0000-0000-000000000000",
 * });
 * ```
 */
export function getAdministrativeUnitOutput(args?: GetAdministrativeUnitOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAdministrativeUnitResult> {
    return pulumi.output(args).apply((a: any) => getAdministrativeUnit(a, opts))
}

/**
 * A collection of arguments for invoking getAdministrativeUnit.
 */
export interface GetAdministrativeUnitOutputArgs {
    /**
     * Specifies the display name of the administrative unit.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Specifies the object ID of the administrative unit.
     *
     * > One of `displayName` or `objectId` must be specified.
     */
    objectId?: pulumi.Input<string>;
}

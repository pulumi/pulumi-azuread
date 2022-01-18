// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Gets Object IDs or Display Names for multiple Azure Active Directory groups.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this data source.
 *
 * When authenticated with a service principal, this data source requires one of the following application roles: `Group.Read.All` or `Directory.Read.All`
 *
 * When authenticated with a user principal, this data source does not require any additional roles.
 *
 * ## Example Usage
 *
 * *Look up by group name*
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = pulumi.output(azuread.getGroups({
 *     displayNames: [
 *         "group-a",
 *         "group-b",
 *     ],
 * }));
 * ```
 *
 * *Look up by display name prefix*
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const sales = pulumi.output(azuread.getGroups({
 *     displayNamePrefix: "sales-",
 * }));
 * ```
 *
 * *Look up all groups*
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const all = pulumi.output(azuread.getGroups({
 *     returnAll: true,
 * }));
 * ```
 *
 * *Look up all mail-enabled groups*
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const mailEnabled = pulumi.output(azuread.getGroups({
 *     mailEnabled: true,
 *     returnAll: true,
 * }));
 * ```
 *
 * *Look up all security-enabled groups that are not mail-enabled*
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const securityOnly = pulumi.output(azuread.getGroups({
 *     mailEnabled: false,
 *     returnAll: true,
 *     securityEnabled: true,
 * }));
 * ```
 */
export function getGroups(args?: GetGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azuread:index/getGroups:getGroups", {
        "displayNamePrefix": args.displayNamePrefix,
        "displayNames": args.displayNames,
        "mailEnabled": args.mailEnabled,
        "objectIds": args.objectIds,
        "returnAll": args.returnAll,
        "securityEnabled": args.securityEnabled,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroups.
 */
export interface GetGroupsArgs {
    /**
     * A common display name prefix to match when returning groups.
     */
    displayNamePrefix?: string;
    /**
     * The display names of the groups.
     */
    displayNames?: string[];
    /**
     * Whether the returned groups should be mail-enabled. By itself this does not exclude security-enabled groups. Setting this to `true` ensures all groups are mail-enabled, and setting to `false` ensures that all groups are _not_ mail-enabled. To ignore this filter, omit the property or set it to null. Cannot be specified together with `objectIds`.
     */
    mailEnabled?: boolean;
    /**
     * The object IDs of the groups.
     */
    objectIds?: string[];
    /**
     * A flag to denote if all groups should be fetched and returned.
     */
    returnAll?: boolean;
    /**
     * Whether the returned groups should be security-enabled. By itself this does not exclude mail-enabled groups. Setting this to `true` ensures all groups are security-enabled, and setting to `false` ensures that all groups are _not_ security-enabled. To ignore this filter, omit the property or set it to null. Cannot be specified together with `objectIds`.
     */
    securityEnabled?: boolean;
}

/**
 * A collection of values returned by getGroups.
 */
export interface GetGroupsResult {
    readonly displayNamePrefix: string;
    /**
     * The display names of the groups.
     */
    readonly displayNames: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly mailEnabled: boolean;
    /**
     * The object IDs of the groups.
     */
    readonly objectIds: string[];
    readonly returnAll?: boolean;
    readonly securityEnabled: boolean;
}

export function getGroupsOutput(args?: GetGroupsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGroupsResult> {
    return pulumi.output(args).apply(a => getGroups(a, opts))
}

/**
 * A collection of arguments for invoking getGroups.
 */
export interface GetGroupsOutputArgs {
    /**
     * A common display name prefix to match when returning groups.
     */
    displayNamePrefix?: pulumi.Input<string>;
    /**
     * The display names of the groups.
     */
    displayNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether the returned groups should be mail-enabled. By itself this does not exclude security-enabled groups. Setting this to `true` ensures all groups are mail-enabled, and setting to `false` ensures that all groups are _not_ mail-enabled. To ignore this filter, omit the property or set it to null. Cannot be specified together with `objectIds`.
     */
    mailEnabled?: pulumi.Input<boolean>;
    /**
     * The object IDs of the groups.
     */
    objectIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A flag to denote if all groups should be fetched and returned.
     */
    returnAll?: pulumi.Input<boolean>;
    /**
     * Whether the returned groups should be security-enabled. By itself this does not exclude mail-enabled groups. Setting this to `true` ensures all groups are security-enabled, and setting to `false` ensures that all groups are _not_ security-enabled. To ignore this filter, omit the property or set it to null. Cannot be specified together with `objectIds`.
     */
    securityEnabled?: pulumi.Input<boolean>;
}

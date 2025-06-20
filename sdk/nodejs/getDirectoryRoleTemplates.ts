// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to access information about directory role templates within Azure Active Directory.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this resource.
 *
 * When authenticated with a service principal, this resource requires one of the following application roles: `RoleManagement.Read.Directory` or `Directory.Read.All`
 *
 * When authenticated with a user principal, this data source does not require any additional roles.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const current = azuread.getDirectoryRoleTemplates({});
 * export const roles = current.then(current => current.objectIds);
 * ```
 */
export function getDirectoryRoleTemplates(opts?: pulumi.InvokeOptions): Promise<GetDirectoryRoleTemplatesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azuread:index/getDirectoryRoleTemplates:getDirectoryRoleTemplates", {
    }, opts);
}

/**
 * A collection of values returned by getDirectoryRoleTemplates.
 */
export interface GetDirectoryRoleTemplatesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The object IDs of the role templates.
     */
    readonly objectIds: string[];
    /**
     * A list of role templates. Each `roleTemplate` object provides the attributes documented below.
     */
    readonly roleTemplates: outputs.GetDirectoryRoleTemplatesRoleTemplate[];
}
/**
 * Use this data source to access information about directory role templates within Azure Active Directory.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this resource.
 *
 * When authenticated with a service principal, this resource requires one of the following application roles: `RoleManagement.Read.Directory` or `Directory.Read.All`
 *
 * When authenticated with a user principal, this data source does not require any additional roles.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const current = azuread.getDirectoryRoleTemplates({});
 * export const roles = current.then(current => current.objectIds);
 * ```
 */
export function getDirectoryRoleTemplatesOutput(opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetDirectoryRoleTemplatesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azuread:index/getDirectoryRoleTemplates:getDirectoryRoleTemplates", {
    }, opts);
}

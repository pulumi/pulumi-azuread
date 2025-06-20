// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * *Listing well-known application IDs*
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const wellKnown = azuread.getApplicationPublishedAppIds({});
 * export const publishedAppIds = wellKnown.then(wellKnown => wellKnown.result);
 * ```
 *
 * *Granting access to an application*
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const wellKnown = azuread.getApplicationPublishedAppIds({});
 * const msgraph = new azuread.ServicePrincipal("msgraph", {
 *     clientId: wellKnown.then(wellKnown => wellKnown.result?.MicrosoftGraph),
 *     useExisting: true,
 * });
 * const example = new azuread.Application("example", {
 *     displayName: "example",
 *     requiredResourceAccesses: [{
 *         resourceAppId: wellKnown.then(wellKnown => wellKnown.result?.MicrosoftGraph),
 *         resourceAccesses: [
 *             {
 *                 id: msgraph.appRoleIds["User.Read.All"],
 *                 type: "Role",
 *             },
 *             {
 *                 id: msgraph.oauth2PermissionScopeIds["User.ReadWrite"],
 *                 type: "Scope",
 *             },
 *         ],
 *     }],
 * });
 * ```
 */
export function getApplicationPublishedAppIds(opts?: pulumi.InvokeOptions): Promise<GetApplicationPublishedAppIdsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azuread:index/getApplicationPublishedAppIds:getApplicationPublishedAppIds", {
    }, opts);
}

/**
 * A collection of values returned by getApplicationPublishedAppIds.
 */
export interface GetApplicationPublishedAppIdsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A map of application names to application IDs.
     */
    readonly result: {[key: string]: string};
}
/**
 * ## Example Usage
 *
 * *Listing well-known application IDs*
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const wellKnown = azuread.getApplicationPublishedAppIds({});
 * export const publishedAppIds = wellKnown.then(wellKnown => wellKnown.result);
 * ```
 *
 * *Granting access to an application*
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const wellKnown = azuread.getApplicationPublishedAppIds({});
 * const msgraph = new azuread.ServicePrincipal("msgraph", {
 *     clientId: wellKnown.then(wellKnown => wellKnown.result?.MicrosoftGraph),
 *     useExisting: true,
 * });
 * const example = new azuread.Application("example", {
 *     displayName: "example",
 *     requiredResourceAccesses: [{
 *         resourceAppId: wellKnown.then(wellKnown => wellKnown.result?.MicrosoftGraph),
 *         resourceAccesses: [
 *             {
 *                 id: msgraph.appRoleIds["User.Read.All"],
 *                 type: "Role",
 *             },
 *             {
 *                 id: msgraph.oauth2PermissionScopeIds["User.ReadWrite"],
 *                 type: "Scope",
 *             },
 *         ],
 *     }],
 * });
 * ```
 */
export function getApplicationPublishedAppIdsOutput(opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetApplicationPublishedAppIdsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azuread:index/getApplicationPublishedAppIds:getApplicationPublishedAppIds", {
    }, opts);
}

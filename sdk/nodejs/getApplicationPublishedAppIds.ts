// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to discover application IDs for APIs published by Microsoft.
 *
 * This data source uses an [unofficial source of application IDs](https://github.com/hashicorp/go-azure-sdk/blob/main/sdk/environments/application_ids.go), as there is currently no available official indexed source for applications or APIs published by Microsoft.
 *
 * The app IDs returned by this data source are sourced from the Azure Global (Public) Cloud, however some of them are known to work in government and national clouds.
 *
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
 *     applicationId: wellKnown.then(wellKnown => wellKnown.result?.MicrosoftGraph),
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
 * Use this data source to discover application IDs for APIs published by Microsoft.
 *
 * This data source uses an [unofficial source of application IDs](https://github.com/hashicorp/go-azure-sdk/blob/main/sdk/environments/application_ids.go), as there is currently no available official indexed source for applications or APIs published by Microsoft.
 *
 * The app IDs returned by this data source are sourced from the Azure Global (Public) Cloud, however some of them are known to work in government and national clouds.
 *
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
 *     applicationId: wellKnown.then(wellKnown => wellKnown.result?.MicrosoftGraph),
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
export function getApplicationPublishedAppIdsOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetApplicationPublishedAppIdsResult> {
    return pulumi.output(getApplicationPublishedAppIds(opts))
}

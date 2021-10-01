// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to access the configuration of the AzureAD provider.
 *
 * ## API Permissions
 *
 * No additional roles are required to use this data source.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const current = azuread.getClientConfig({});
 * export const objectId = current.then(current => current.objectId);
 * ```
 */
export function getClientConfig(opts?: pulumi.InvokeOptions): Promise<GetClientConfigResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azuread:index/getClientConfig:getClientConfig", {
    }, opts);
}

/**
 * A collection of values returned by getClientConfig.
 */
export interface GetClientConfigResult {
    /**
     * The client ID (application ID) linked to the authenticated principal, or the application used for delegated authentication.
     */
    readonly clientId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The object ID of the authenticated principal.
     */
    readonly objectId: string;
    /**
     * The tenant ID of the authenticated principal.
     */
    readonly tenantId: string;
}

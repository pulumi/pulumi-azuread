// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

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
    displayName?: string;
    objectId?: string;
}

/**
 * A collection of values returned by getAccessPackageCatalogRole.
 */
export interface GetAccessPackageCatalogRoleResult {
    readonly description: string;
    readonly displayName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly objectId: string;
    readonly templateId: string;
}
export function getAccessPackageCatalogRoleOutput(args?: GetAccessPackageCatalogRoleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAccessPackageCatalogRoleResult> {
    return pulumi.output(args).apply((a: any) => getAccessPackageCatalogRole(a, opts))
}

/**
 * A collection of arguments for invoking getAccessPackageCatalogRole.
 */
export interface GetAccessPackageCatalogRoleOutputArgs {
    displayName?: pulumi.Input<string>;
    objectId?: pulumi.Input<string>;
}

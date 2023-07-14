// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getAccessPackage(args?: GetAccessPackageArgs, opts?: pulumi.InvokeOptions): Promise<GetAccessPackageResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azuread:index/getAccessPackage:getAccessPackage", {
        "catalogId": args.catalogId,
        "displayName": args.displayName,
        "objectId": args.objectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAccessPackage.
 */
export interface GetAccessPackageArgs {
    catalogId?: string;
    displayName?: string;
    objectId?: string;
}

/**
 * A collection of values returned by getAccessPackage.
 */
export interface GetAccessPackageResult {
    readonly catalogId?: string;
    readonly description: string;
    readonly displayName: string;
    readonly hidden: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly objectId: string;
}
export function getAccessPackageOutput(args?: GetAccessPackageOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAccessPackageResult> {
    return pulumi.output(args).apply((a: any) => getAccessPackage(a, opts))
}

/**
 * A collection of arguments for invoking getAccessPackage.
 */
export interface GetAccessPackageOutputArgs {
    catalogId?: pulumi.Input<string>;
    displayName?: pulumi.Input<string>;
    objectId?: pulumi.Input<string>;
}

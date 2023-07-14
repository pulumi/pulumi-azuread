// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getServicePrincipals(args?: GetServicePrincipalsArgs, opts?: pulumi.InvokeOptions): Promise<GetServicePrincipalsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azuread:index/getServicePrincipals:getServicePrincipals", {
        "applicationIds": args.applicationIds,
        "displayNames": args.displayNames,
        "ignoreMissing": args.ignoreMissing,
        "objectIds": args.objectIds,
        "returnAll": args.returnAll,
    }, opts);
}

/**
 * A collection of arguments for invoking getServicePrincipals.
 */
export interface GetServicePrincipalsArgs {
    applicationIds?: string[];
    displayNames?: string[];
    ignoreMissing?: boolean;
    objectIds?: string[];
    returnAll?: boolean;
}

/**
 * A collection of values returned by getServicePrincipals.
 */
export interface GetServicePrincipalsResult {
    readonly applicationIds: string[];
    readonly displayNames: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ignoreMissing?: boolean;
    readonly objectIds: string[];
    readonly returnAll?: boolean;
    readonly servicePrincipals: outputs.GetServicePrincipalsServicePrincipal[];
}
export function getServicePrincipalsOutput(args?: GetServicePrincipalsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServicePrincipalsResult> {
    return pulumi.output(args).apply((a: any) => getServicePrincipals(a, opts))
}

/**
 * A collection of arguments for invoking getServicePrincipals.
 */
export interface GetServicePrincipalsOutputArgs {
    applicationIds?: pulumi.Input<pulumi.Input<string>[]>;
    displayNames?: pulumi.Input<pulumi.Input<string>[]>;
    ignoreMissing?: pulumi.Input<boolean>;
    objectIds?: pulumi.Input<pulumi.Input<string>[]>;
    returnAll?: pulumi.Input<boolean>;
}

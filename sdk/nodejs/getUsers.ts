// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getUsers(args?: GetUsersArgs, opts?: pulumi.InvokeOptions): Promise<GetUsersResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azuread:index/getUsers:getUsers", {
        "employeeIds": args.employeeIds,
        "ignoreMissing": args.ignoreMissing,
        "mailNicknames": args.mailNicknames,
        "objectIds": args.objectIds,
        "returnAll": args.returnAll,
        "userPrincipalNames": args.userPrincipalNames,
    }, opts);
}

/**
 * A collection of arguments for invoking getUsers.
 */
export interface GetUsersArgs {
    employeeIds?: string[];
    ignoreMissing?: boolean;
    mailNicknames?: string[];
    objectIds?: string[];
    returnAll?: boolean;
    userPrincipalNames?: string[];
}

/**
 * A collection of values returned by getUsers.
 */
export interface GetUsersResult {
    readonly employeeIds: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ignoreMissing?: boolean;
    readonly mailNicknames: string[];
    readonly objectIds: string[];
    readonly returnAll?: boolean;
    readonly userPrincipalNames: string[];
    readonly users: outputs.GetUsersUser[];
}
export function getUsersOutput(args?: GetUsersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetUsersResult> {
    return pulumi.output(args).apply((a: any) => getUsers(a, opts))
}

/**
 * A collection of arguments for invoking getUsers.
 */
export interface GetUsersOutputArgs {
    employeeIds?: pulumi.Input<pulumi.Input<string>[]>;
    ignoreMissing?: pulumi.Input<boolean>;
    mailNicknames?: pulumi.Input<pulumi.Input<string>[]>;
    objectIds?: pulumi.Input<pulumi.Input<string>[]>;
    returnAll?: pulumi.Input<boolean>;
    userPrincipalNames?: pulumi.Input<pulumi.Input<string>[]>;
}

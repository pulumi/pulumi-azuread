// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Gets information about an Azure Active Directory group.
 *
 * > **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `Read directory data` within the `Windows Azure Active Directory` API.
 *
 * ## Example Usage
 * ### By Group Display Name)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = pulumi.output(azuread.getGroup({
 *     displayName: "MyGroupName",
 *     securityEnabled: true,
 * }));
 * ```
 */
export function getGroup(args?: GetGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azuread:index/getGroup:getGroup", {
        "displayName": args.displayName,
        "mailEnabled": args.mailEnabled,
        "name": args.name,
        "objectId": args.objectId,
        "securityEnabled": args.securityEnabled,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroup.
 */
export interface GetGroupArgs {
    /**
     * The display name for the Group.
     */
    displayName?: string;
    /**
     * Whether the group is mail-enabled.
     */
    mailEnabled?: boolean;
    /**
     * @deprecated This property has been renamed to `display_name` and will be removed in version 2.0 of the AzureAD provider.
     */
    name?: string;
    /**
     * Specifies the Object ID of the Group.
     */
    objectId?: string;
    /**
     * Whether the group is a security group.
     */
    securityEnabled?: boolean;
}

/**
 * A collection of values returned by getGroup.
 */
export interface GetGroupResult {
    /**
     * The optional description of the Group.
     */
    readonly description: string;
    /**
     * The display name for the Group.
     */
    readonly displayName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Whether the group is mail-enabled.
     */
    readonly mailEnabled: boolean;
    /**
     * The Object IDs of the Group members.
     */
    readonly members: string[];
    /**
     * @deprecated This property has been renamed to `display_name` and will be removed in version 2.0 of the AzureAD provider.
     */
    readonly name: string;
    readonly objectId: string;
    /**
     * The Object IDs of the Group owners.
     */
    readonly owners: string[];
    /**
     * Whether the group is a security group.
     */
    readonly securityEnabled: boolean;
}

// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve a role policy for an Azure AD group.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this resource.
 *
 * When authenticated with a service principal, this resource requires the `RoleManagementPolicy.Read.AzureADGroup` Microsoft Graph API permissions.
 *
 * When authenticated with a user principal, this resource requires `Global Administrator` directory role, or the `Privileged Role Administrator` role in Identity Governance.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = new azuread.Group("example", {
 *     displayName: "group-name",
 *     securityEnabled: true,
 * });
 * const ownersPolicy = azuread.getGroupRoleManagementPolicyOutput({
 *     groupId: example.id,
 *     roleId: "owner",
 * });
 * ```
 */
export function getGroupRoleManagementPolicy(args: GetGroupRoleManagementPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupRoleManagementPolicyResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azuread:index/getGroupRoleManagementPolicy:getGroupRoleManagementPolicy", {
        "groupId": args.groupId,
        "roleId": args.roleId,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroupRoleManagementPolicy.
 */
export interface GetGroupRoleManagementPolicyArgs {
    /**
     * The ID of the Azure AD group for which the policy applies.
     */
    groupId: string;
    /**
     * The type of assignment this policy coveres. Can be either `member` or `owner`.
     */
    roleId: string;
}

/**
 * A collection of values returned by getGroupRoleManagementPolicy.
 */
export interface GetGroupRoleManagementPolicyResult {
    /**
     * (String) The description of this policy.
     */
    readonly description: string;
    /**
     * (String) The display name of this policy.
     */
    readonly displayName: string;
    readonly groupId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly roleId: string;
}
/**
 * Use this data source to retrieve a role policy for an Azure AD group.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this resource.
 *
 * When authenticated with a service principal, this resource requires the `RoleManagementPolicy.Read.AzureADGroup` Microsoft Graph API permissions.
 *
 * When authenticated with a user principal, this resource requires `Global Administrator` directory role, or the `Privileged Role Administrator` role in Identity Governance.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = new azuread.Group("example", {
 *     displayName: "group-name",
 *     securityEnabled: true,
 * });
 * const ownersPolicy = azuread.getGroupRoleManagementPolicyOutput({
 *     groupId: example.id,
 *     roleId: "owner",
 * });
 * ```
 */
export function getGroupRoleManagementPolicyOutput(args: GetGroupRoleManagementPolicyOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetGroupRoleManagementPolicyResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azuread:index/getGroupRoleManagementPolicy:getGroupRoleManagementPolicy", {
        "groupId": args.groupId,
        "roleId": args.roleId,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroupRoleManagementPolicy.
 */
export interface GetGroupRoleManagementPolicyOutputArgs {
    /**
     * The ID of the Azure AD group for which the policy applies.
     */
    groupId: pulumi.Input<string>;
    /**
     * The type of assignment this policy coveres. Can be either `member` or `owner`.
     */
    roleId: pulumi.Input<string>;
}

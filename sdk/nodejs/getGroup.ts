// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Gets information about an Azure Active Directory group.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this data source.
 *
 * When authenticated with a service principal, this data source requires one of the following application roles: `Group.Read.All` or `Directory.Read.All`
 *
 * When authenticated with a user principal, this data source does not require any additional roles.
 *
 * ## Example Usage
 * ### By Group Display Name)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = azuread.getGroup({
 *     displayName: "MyGroupName",
 *     securityEnabled: true,
 * });
 * ```
 */
export function getGroup(args?: GetGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azuread:index/getGroup:getGroup", {
        "displayName": args.displayName,
        "mailEnabled": args.mailEnabled,
        "objectId": args.objectId,
        "securityEnabled": args.securityEnabled,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroup.
 */
export interface GetGroupArgs {
    /**
     * The display name for the group.
     */
    displayName?: string;
    /**
     * Whether the group is mail-enabled.
     */
    mailEnabled?: boolean;
    /**
     * Specifies the object ID of the group.
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
     * Indicates whether this group can be assigned to an Azure Active Directory role.
     */
    readonly assignableToRole: boolean;
    /**
     * Indicates whether new members added to the group will be auto-subscribed to receive email notifications. Only set for Unified groups.
     */
    readonly autoSubscribeNewMembers: boolean;
    /**
     * A list of behaviors for a Microsoft 365 group, such as `AllowOnlyMembersToPost`, `HideGroupInOutlook`, `SubscribeNewGroupMembers` and `WelcomeEmailDisabled`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for more details.
     */
    readonly behaviors: string[];
    /**
     * The optional description of the group.
     */
    readonly description: string;
    /**
     * The display name for the group.
     */
    readonly displayName: string;
    /**
     * A `dynamicMembership` block as documented below.
     */
    readonly dynamicMemberships: outputs.GetGroupDynamicMembership[];
    /**
     * Indicates whether people external to the organization can send messages to the group. Only set for Unified groups.
     */
    readonly externalSendersAllowed: boolean;
    /**
     * Indicates whether the group is displayed in certain parts of the Outlook user interface: in the Address Book, in address lists for selecting message recipients, and in the Browse Groups dialog for searching groups. Only set for Unified groups.
     */
    readonly hideFromAddressLists: boolean;
    /**
     * Indicates whether the group is displayed in Outlook clients, such as Outlook for Windows and Outlook on the web. Only set for Unified groups.
     */
    readonly hideFromOutlookClients: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The SMTP address for the group.
     */
    readonly mail: string;
    /**
     * Whether the group is mail-enabled.
     */
    readonly mailEnabled: boolean;
    /**
     * The mail alias for the group, unique in the organisation.
     */
    readonly mailNickname: string;
    /**
     * List of object IDs of the group members.
     */
    readonly members: string[];
    /**
     * The object ID of the group.
     */
    readonly objectId: string;
    /**
     * The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.
     */
    readonly onpremisesDomainName: string;
    /**
     * The on-premises NetBIOS name, synchronised from the on-premises directory when Azure AD Connect is used.
     */
    readonly onpremisesNetbiosName: string;
    /**
     * The on-premises SAM account name, synchronised from the on-premises directory when Azure AD Connect is used.
     */
    readonly onpremisesSamAccountName: string;
    /**
     * The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.
     */
    readonly onpremisesSecurityIdentifier: string;
    /**
     * Whether this group is synchronised from an on-premises directory (`true`), no longer synchronised (`false`), or has never been synchronised (`null`).
     */
    readonly onpremisesSyncEnabled: boolean;
    /**
     * List of object IDs of the group owners.
     */
    readonly owners: string[];
    /**
     * The preferred language for a Microsoft 365 group, in ISO 639-1 notation.
     */
    readonly preferredLanguage: string;
    /**
     * A list of provisioning options for a Microsoft 365 group, such as `Team`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for details.
     */
    readonly provisioningOptions: string[];
    /**
     * List of email addresses for the group that direct to the same group mailbox.
     */
    readonly proxyAddresses: string[];
    /**
     * Whether the group is a security group.
     */
    readonly securityEnabled: boolean;
    /**
     * The colour theme for a Microsoft 365 group. Possible values are `Blue`, `Green`, `Orange`, `Pink`, `Purple`, `Red` or `Teal`. When no theme is set, the value is `null`.
     */
    readonly theme: string;
    /**
     * A list of group types configured for the group. Supported values are `DynamicMembership`, which denotes a group with dynamic membership, and `Unified`, which specifies a Microsoft 365 group.
     */
    readonly types: string[];
    /**
     * The group join policy and group content visibility. Possible values are `Private`, `Public`, or `Hiddenmembership`. Only Microsoft 365 groups can have `Hiddenmembership` visibility.
     */
    readonly visibility: string;
}
/**
 * Gets information about an Azure Active Directory group.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this data source.
 *
 * When authenticated with a service principal, this data source requires one of the following application roles: `Group.Read.All` or `Directory.Read.All`
 *
 * When authenticated with a user principal, this data source does not require any additional roles.
 *
 * ## Example Usage
 * ### By Group Display Name)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = azuread.getGroup({
 *     displayName: "MyGroupName",
 *     securityEnabled: true,
 * });
 * ```
 */
export function getGroupOutput(args?: GetGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGroupResult> {
    return pulumi.output(args).apply((a: any) => getGroup(a, opts))
}

/**
 * A collection of arguments for invoking getGroup.
 */
export interface GetGroupOutputArgs {
    /**
     * The display name for the group.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Whether the group is mail-enabled.
     */
    mailEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies the object ID of the group.
     */
    objectId?: pulumi.Input<string>;
    /**
     * Whether the group is a security group.
     */
    securityEnabled?: pulumi.Input<boolean>;
}

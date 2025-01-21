// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Manages a group within Azure Active Directory.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this resource.
 *
 * When authenticated with a service principal, this resource requires one of the following application roles: `Group.ReadWrite.All` or `Directory.ReadWrite.All`.
 *
 * Alternatively, if the authenticated service principal is also an owner of the group being managed, this resource can use the application role: `Group.Create`.
 *
 * If using the `assignableToRole` property, this resource additionally requires the `RoleManagement.ReadWrite.Directory` application role.
 *
 * If specifying owners for a group, which are user principals, this resource additionally requires one of the following application roles: `User.Read.All`, `User.ReadWrite.All`, `Directory.Read.All` or `Directory.ReadWrite.All`
 *
 * When authenticated with a user principal, this resource requires one of the following directory roles: `Groups Administrator`, `User Administrator` or `Global Administrator`
 *
 * When creating this resource in administrative units exclusively, the directory role `Groups Administrator` is required to be scoped on any administrative unit used. Additionally, it must be possible to read the administrative units being used, which can be granted through the `AdministrativeUnit.Read.All` or `Directory.Read.All` application roles.
 *
 * The `externalSendersAllowed`, `autoSubscribeNewMembers`, `hideFromAddressLists` and `hideFromOutlookClients` properties can only be configured when authenticating as a user and cannot be configured when authenticating as a service principal. Additionally, the user being used for authentication must be a Member of the tenant where the group is being managed and _not_ a Guest. This is a known API issue; please see the [Microsoft Graph Known Issues](https://docs.microsoft.com/en-us/graph/known-issues#groups) official documentation.
 *
 * ## Example Usage
 *
 * *Basic example*
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const current = azuread.getClientConfig({});
 * const example = new azuread.Group("example", {
 *     displayName: "example",
 *     owners: [current.then(current => current.objectId)],
 *     securityEnabled: true,
 * });
 * ```
 *
 * *Microsoft 365 group*
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const current = azuread.getClientConfig({});
 * const groupOwner = new azuread.User("group_owner", {
 *     userPrincipalName: "example-group-owner@example.com",
 *     displayName: "Group Owner",
 *     mailNickname: "example-group-owner",
 *     password: "SecretP@sswd99!",
 * });
 * const example = new azuread.Group("example", {
 *     displayName: "example",
 *     mailEnabled: true,
 *     mailNickname: "ExampleGroup",
 *     securityEnabled: true,
 *     types: ["Unified"],
 *     owners: [
 *         current.then(current => current.objectId),
 *         groupOwner.objectId,
 *     ],
 * });
 * ```
 *
 * *Group with members*
 *
 * ## Import
 *
 * Groups can be imported using their object ID, e.g.
 *
 * ```sh
 * $ pulumi import azuread:index/group:Group my_group /groups/00000000-0000-0000-0000-000000000000
 * ```
 */
export class Group extends pulumi.CustomResource {
    /**
     * Get an existing Group resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupState, opts?: pulumi.CustomResourceOptions): Group {
        return new Group(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/group:Group';

    /**
     * Returns true if the given object is an instance of Group.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Group {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Group.__pulumiType;
    }

    /**
     * The object IDs of administrative units in which the group is a member. If specified, new groups will be created in the scope of the first administrative unit and added to the others. If empty, new groups will be created at the tenant level.
     *
     * > **Caution** When using the azuread.AdministrativeUnitMember resource, or the `members` property of the azuread.AdministrativeUnit resource, to manage Administrative Unit membership for a group, you will need to use an `ignoreChanges = [administrativeUnitIds]` lifecycle meta argument for the `azuread.Group` resource, in order to avoid a persistent diff.
     */
    public readonly administrativeUnitIds!: pulumi.Output<string[] | undefined>;
    /**
     * Indicates whether this group can be assigned to an Azure Active Directory role. Defaults to `false`. Can only be set to `true` for security-enabled groups. Changing this forces a new resource to be created.
     */
    public readonly assignableToRole!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates whether new members added to the group will be auto-subscribed to receive email notifications. Can only be set for Unified groups.
     *
     * > **Known Permissions Issue** The `autoSubscribeNewMembers` property can only be set when authenticating as a Member user of the tenant and _not_ when authenticating as a Guest user or as a service principal. Please see the [Microsoft Graph Known Issues](https://docs.microsoft.com/en-us/graph/known-issues#groups) documentation.
     */
    public readonly autoSubscribeNewMembers!: pulumi.Output<boolean>;
    /**
     * A set of behaviors for a Microsoft 365 group. Possible values are `AllowOnlyMembersToPost`, `HideGroupInOutlook`, `SkipExchangeInstantOn`, `SubscribeMembersToCalendarEventsDisabled`, `SubscribeNewGroupMembers` and `WelcomeEmailDisabled`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for more details. Changing this forces a new resource to be created.
     */
    public readonly behaviors!: pulumi.Output<string[] | undefined>;
    /**
     * The description for the group.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The display name for the group.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * A `dynamicMembership` block as documented below. Required when `types` contains `DynamicMembership`. Cannot be used with the `members` property.
     */
    public readonly dynamicMembership!: pulumi.Output<outputs.GroupDynamicMembership | undefined>;
    /**
     * Indicates whether people external to the organization can send messages to the group. Can only be set for Unified groups.
     *
     * > **Known Permissions Issue** The `externalSendersAllowed` property can only be set when authenticating as a Member user of the tenant and _not_ when authenticating as a Guest user or as a service principal. Please see the [Microsoft Graph Known Issues](https://docs.microsoft.com/en-us/graph/known-issues#groups) documentation.
     */
    public readonly externalSendersAllowed!: pulumi.Output<boolean>;
    /**
     * Indicates whether the group is displayed in certain parts of the Outlook user interface: in the Address Book, in address lists for selecting message recipients, and in the Browse Groups dialog for searching groups. Can only be set for Unified groups.
     *
     * > **Known Permissions Issue** The `hideFromAddressLists` property can only be set when authenticating as a Member user of the tenant and _not_ when authenticating as a Guest user or as a service principal. Please see the [Microsoft Graph Known Issues](https://docs.microsoft.com/en-us/graph/known-issues#groups) documentation.
     */
    public readonly hideFromAddressLists!: pulumi.Output<boolean>;
    /**
     * Indicates whether the group is displayed in Outlook clients, such as Outlook for Windows and Outlook on the web. Can only be set for Unified groups.
     *
     * > **Known Permissions Issue** The `hideFromOutlookClients` property can only be set when authenticating as a Member user of the tenant and _not_ when authenticating as a Guest user or as a service principal. Please see the [Microsoft Graph Known Issues](https://docs.microsoft.com/en-us/graph/known-issues#groups) documentation.
     */
    public readonly hideFromOutlookClients!: pulumi.Output<boolean>;
    /**
     * The SMTP address for the group.
     */
    public /*out*/ readonly mail!: pulumi.Output<string>;
    /**
     * Whether the group is a mail enabled, with a shared group mailbox. At least one of `mailEnabled` or `securityEnabled` must be specified. Only Microsoft 365 groups can be mail enabled (see the `types` property).
     */
    public readonly mailEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The mail alias for the group, unique in the organisation. Required for mail-enabled groups. Changing this forces a new resource to be created.
     */
    public readonly mailNickname!: pulumi.Output<string>;
    /**
     * A set of members who should be present in this group. Supported object types are Users, Groups or Service Principals. Cannot be used with the `dynamicMembership` block.
     *
     * !> **Warning** Do not use the `members` property at the same time as the azuread.GroupMember resource for the same group. Doing so will cause a conflict and group members will be removed.
     */
    public readonly members!: pulumi.Output<string[]>;
    /**
     * The object ID of the group.
     */
    public /*out*/ readonly objectId!: pulumi.Output<string>;
    /**
     * The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.
     */
    public /*out*/ readonly onpremisesDomainName!: pulumi.Output<string>;
    /**
     * The on-premises group type that the AAD group will be written as, when writeback is enabled. Possible values are `UniversalDistributionGroup`, `UniversalMailEnabledSecurityGroup`, or `UniversalSecurityGroup`.
     */
    public readonly onpremisesGroupType!: pulumi.Output<string>;
    /**
     * The on-premises NetBIOS name, synchronised from the on-premises directory when Azure AD Connect is used.
     */
    public /*out*/ readonly onpremisesNetbiosName!: pulumi.Output<string>;
    /**
     * The on-premises SAM account name, synchronised from the on-premises directory when Azure AD Connect is used.
     */
    public /*out*/ readonly onpremisesSamAccountName!: pulumi.Output<string>;
    /**
     * The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.
     */
    public /*out*/ readonly onpremisesSecurityIdentifier!: pulumi.Output<string>;
    /**
     * Whether this group is synchronised from an on-premises directory (`true`), no longer synchronised (`false`), or has never been synchronised (`null`).
     */
    public /*out*/ readonly onpremisesSyncEnabled!: pulumi.Output<boolean>;
    /**
     * A set of owners who own this group. Supported object types are Users or Service Principals
     */
    public readonly owners!: pulumi.Output<string[]>;
    /**
     * The preferred language for a Microsoft 365 group, in ISO 639-1 notation.
     */
    public /*out*/ readonly preferredLanguage!: pulumi.Output<string>;
    /**
     * If `true`, will return an error if an existing group is found with the same name. Defaults to `false`.
     */
    public readonly preventDuplicateNames!: pulumi.Output<boolean | undefined>;
    /**
     * A set of provisioning options for a Microsoft 365 group. The only supported value is `Team`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for details. Changing this forces a new resource to be created.
     */
    public readonly provisioningOptions!: pulumi.Output<string[] | undefined>;
    /**
     * List of email addresses for the group that direct to the same group mailbox.
     */
    public /*out*/ readonly proxyAddresses!: pulumi.Output<string[]>;
    /**
     * Whether the group is a security group for controlling access to in-app resources. At least one of `securityEnabled` or `mailEnabled` must be specified. A Microsoft 365 group can be security enabled _and_ mail enabled (see the `types` property).
     */
    public readonly securityEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The colour theme for a Microsoft 365 group. Possible values are `Blue`, `Green`, `Orange`, `Pink`, `Purple`, `Red` or `Teal`. By default, no theme is set.
     */
    public readonly theme!: pulumi.Output<string | undefined>;
    /**
     * A set of group types to configure for the group. Supported values are `DynamicMembership`, which denotes a group with dynamic membership, and `Unified`, which specifies a Microsoft 365 group. Required when `mailEnabled` is true. Changing this forces a new resource to be created.
     *
     * > **Supported Group Types** At present, only security groups and Microsoft 365 groups can be created or managed with this resource. Distribution groups and mail-enabled security groups are not supported. Microsoft 365 groups can be security-enabled.
     */
    public readonly types!: pulumi.Output<string[] | undefined>;
    /**
     * The group join policy and group content visibility. Possible values are `Private`, `Public`, or `Hiddenmembership`. Only Microsoft 365 groups can have `Hiddenmembership` visibility and this value must be set when the group is created. By default, security groups will receive `Private` visibility and Microsoft 365 groups will receive `Public` visibility.
     *
     * > **Group Name Uniqueness** Group names are not unique within Azure Active Directory. Use the `preventDuplicateNames` argument to check for existing groups if you want to avoid name collisions.
     */
    public readonly visibility!: pulumi.Output<string>;
    /**
     * Whether the group will be written back to the configured on-premises Active Directory when Azure AD Connect is used.
     */
    public readonly writebackEnabled!: pulumi.Output<boolean | undefined>;

    /**
     * Create a Group resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupArgs | GroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupState | undefined;
            resourceInputs["administrativeUnitIds"] = state ? state.administrativeUnitIds : undefined;
            resourceInputs["assignableToRole"] = state ? state.assignableToRole : undefined;
            resourceInputs["autoSubscribeNewMembers"] = state ? state.autoSubscribeNewMembers : undefined;
            resourceInputs["behaviors"] = state ? state.behaviors : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["dynamicMembership"] = state ? state.dynamicMembership : undefined;
            resourceInputs["externalSendersAllowed"] = state ? state.externalSendersAllowed : undefined;
            resourceInputs["hideFromAddressLists"] = state ? state.hideFromAddressLists : undefined;
            resourceInputs["hideFromOutlookClients"] = state ? state.hideFromOutlookClients : undefined;
            resourceInputs["mail"] = state ? state.mail : undefined;
            resourceInputs["mailEnabled"] = state ? state.mailEnabled : undefined;
            resourceInputs["mailNickname"] = state ? state.mailNickname : undefined;
            resourceInputs["members"] = state ? state.members : undefined;
            resourceInputs["objectId"] = state ? state.objectId : undefined;
            resourceInputs["onpremisesDomainName"] = state ? state.onpremisesDomainName : undefined;
            resourceInputs["onpremisesGroupType"] = state ? state.onpremisesGroupType : undefined;
            resourceInputs["onpremisesNetbiosName"] = state ? state.onpremisesNetbiosName : undefined;
            resourceInputs["onpremisesSamAccountName"] = state ? state.onpremisesSamAccountName : undefined;
            resourceInputs["onpremisesSecurityIdentifier"] = state ? state.onpremisesSecurityIdentifier : undefined;
            resourceInputs["onpremisesSyncEnabled"] = state ? state.onpremisesSyncEnabled : undefined;
            resourceInputs["owners"] = state ? state.owners : undefined;
            resourceInputs["preferredLanguage"] = state ? state.preferredLanguage : undefined;
            resourceInputs["preventDuplicateNames"] = state ? state.preventDuplicateNames : undefined;
            resourceInputs["provisioningOptions"] = state ? state.provisioningOptions : undefined;
            resourceInputs["proxyAddresses"] = state ? state.proxyAddresses : undefined;
            resourceInputs["securityEnabled"] = state ? state.securityEnabled : undefined;
            resourceInputs["theme"] = state ? state.theme : undefined;
            resourceInputs["types"] = state ? state.types : undefined;
            resourceInputs["visibility"] = state ? state.visibility : undefined;
            resourceInputs["writebackEnabled"] = state ? state.writebackEnabled : undefined;
        } else {
            const args = argsOrState as GroupArgs | undefined;
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            resourceInputs["administrativeUnitIds"] = args ? args.administrativeUnitIds : undefined;
            resourceInputs["assignableToRole"] = args ? args.assignableToRole : undefined;
            resourceInputs["autoSubscribeNewMembers"] = args ? args.autoSubscribeNewMembers : undefined;
            resourceInputs["behaviors"] = args ? args.behaviors : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["dynamicMembership"] = args ? args.dynamicMembership : undefined;
            resourceInputs["externalSendersAllowed"] = args ? args.externalSendersAllowed : undefined;
            resourceInputs["hideFromAddressLists"] = args ? args.hideFromAddressLists : undefined;
            resourceInputs["hideFromOutlookClients"] = args ? args.hideFromOutlookClients : undefined;
            resourceInputs["mailEnabled"] = args ? args.mailEnabled : undefined;
            resourceInputs["mailNickname"] = args ? args.mailNickname : undefined;
            resourceInputs["members"] = args ? args.members : undefined;
            resourceInputs["onpremisesGroupType"] = args ? args.onpremisesGroupType : undefined;
            resourceInputs["owners"] = args ? args.owners : undefined;
            resourceInputs["preventDuplicateNames"] = args ? args.preventDuplicateNames : undefined;
            resourceInputs["provisioningOptions"] = args ? args.provisioningOptions : undefined;
            resourceInputs["securityEnabled"] = args ? args.securityEnabled : undefined;
            resourceInputs["theme"] = args ? args.theme : undefined;
            resourceInputs["types"] = args ? args.types : undefined;
            resourceInputs["visibility"] = args ? args.visibility : undefined;
            resourceInputs["writebackEnabled"] = args ? args.writebackEnabled : undefined;
            resourceInputs["mail"] = undefined /*out*/;
            resourceInputs["objectId"] = undefined /*out*/;
            resourceInputs["onpremisesDomainName"] = undefined /*out*/;
            resourceInputs["onpremisesNetbiosName"] = undefined /*out*/;
            resourceInputs["onpremisesSamAccountName"] = undefined /*out*/;
            resourceInputs["onpremisesSecurityIdentifier"] = undefined /*out*/;
            resourceInputs["onpremisesSyncEnabled"] = undefined /*out*/;
            resourceInputs["preferredLanguage"] = undefined /*out*/;
            resourceInputs["proxyAddresses"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Group.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Group resources.
 */
export interface GroupState {
    /**
     * The object IDs of administrative units in which the group is a member. If specified, new groups will be created in the scope of the first administrative unit and added to the others. If empty, new groups will be created at the tenant level.
     *
     * > **Caution** When using the azuread.AdministrativeUnitMember resource, or the `members` property of the azuread.AdministrativeUnit resource, to manage Administrative Unit membership for a group, you will need to use an `ignoreChanges = [administrativeUnitIds]` lifecycle meta argument for the `azuread.Group` resource, in order to avoid a persistent diff.
     */
    administrativeUnitIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Indicates whether this group can be assigned to an Azure Active Directory role. Defaults to `false`. Can only be set to `true` for security-enabled groups. Changing this forces a new resource to be created.
     */
    assignableToRole?: pulumi.Input<boolean>;
    /**
     * Indicates whether new members added to the group will be auto-subscribed to receive email notifications. Can only be set for Unified groups.
     *
     * > **Known Permissions Issue** The `autoSubscribeNewMembers` property can only be set when authenticating as a Member user of the tenant and _not_ when authenticating as a Guest user or as a service principal. Please see the [Microsoft Graph Known Issues](https://docs.microsoft.com/en-us/graph/known-issues#groups) documentation.
     */
    autoSubscribeNewMembers?: pulumi.Input<boolean>;
    /**
     * A set of behaviors for a Microsoft 365 group. Possible values are `AllowOnlyMembersToPost`, `HideGroupInOutlook`, `SkipExchangeInstantOn`, `SubscribeMembersToCalendarEventsDisabled`, `SubscribeNewGroupMembers` and `WelcomeEmailDisabled`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for more details. Changing this forces a new resource to be created.
     */
    behaviors?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The description for the group.
     */
    description?: pulumi.Input<string>;
    /**
     * The display name for the group.
     */
    displayName?: pulumi.Input<string>;
    /**
     * A `dynamicMembership` block as documented below. Required when `types` contains `DynamicMembership`. Cannot be used with the `members` property.
     */
    dynamicMembership?: pulumi.Input<inputs.GroupDynamicMembership>;
    /**
     * Indicates whether people external to the organization can send messages to the group. Can only be set for Unified groups.
     *
     * > **Known Permissions Issue** The `externalSendersAllowed` property can only be set when authenticating as a Member user of the tenant and _not_ when authenticating as a Guest user or as a service principal. Please see the [Microsoft Graph Known Issues](https://docs.microsoft.com/en-us/graph/known-issues#groups) documentation.
     */
    externalSendersAllowed?: pulumi.Input<boolean>;
    /**
     * Indicates whether the group is displayed in certain parts of the Outlook user interface: in the Address Book, in address lists for selecting message recipients, and in the Browse Groups dialog for searching groups. Can only be set for Unified groups.
     *
     * > **Known Permissions Issue** The `hideFromAddressLists` property can only be set when authenticating as a Member user of the tenant and _not_ when authenticating as a Guest user or as a service principal. Please see the [Microsoft Graph Known Issues](https://docs.microsoft.com/en-us/graph/known-issues#groups) documentation.
     */
    hideFromAddressLists?: pulumi.Input<boolean>;
    /**
     * Indicates whether the group is displayed in Outlook clients, such as Outlook for Windows and Outlook on the web. Can only be set for Unified groups.
     *
     * > **Known Permissions Issue** The `hideFromOutlookClients` property can only be set when authenticating as a Member user of the tenant and _not_ when authenticating as a Guest user or as a service principal. Please see the [Microsoft Graph Known Issues](https://docs.microsoft.com/en-us/graph/known-issues#groups) documentation.
     */
    hideFromOutlookClients?: pulumi.Input<boolean>;
    /**
     * The SMTP address for the group.
     */
    mail?: pulumi.Input<string>;
    /**
     * Whether the group is a mail enabled, with a shared group mailbox. At least one of `mailEnabled` or `securityEnabled` must be specified. Only Microsoft 365 groups can be mail enabled (see the `types` property).
     */
    mailEnabled?: pulumi.Input<boolean>;
    /**
     * The mail alias for the group, unique in the organisation. Required for mail-enabled groups. Changing this forces a new resource to be created.
     */
    mailNickname?: pulumi.Input<string>;
    /**
     * A set of members who should be present in this group. Supported object types are Users, Groups or Service Principals. Cannot be used with the `dynamicMembership` block.
     *
     * !> **Warning** Do not use the `members` property at the same time as the azuread.GroupMember resource for the same group. Doing so will cause a conflict and group members will be removed.
     */
    members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The object ID of the group.
     */
    objectId?: pulumi.Input<string>;
    /**
     * The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.
     */
    onpremisesDomainName?: pulumi.Input<string>;
    /**
     * The on-premises group type that the AAD group will be written as, when writeback is enabled. Possible values are `UniversalDistributionGroup`, `UniversalMailEnabledSecurityGroup`, or `UniversalSecurityGroup`.
     */
    onpremisesGroupType?: pulumi.Input<string>;
    /**
     * The on-premises NetBIOS name, synchronised from the on-premises directory when Azure AD Connect is used.
     */
    onpremisesNetbiosName?: pulumi.Input<string>;
    /**
     * The on-premises SAM account name, synchronised from the on-premises directory when Azure AD Connect is used.
     */
    onpremisesSamAccountName?: pulumi.Input<string>;
    /**
     * The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.
     */
    onpremisesSecurityIdentifier?: pulumi.Input<string>;
    /**
     * Whether this group is synchronised from an on-premises directory (`true`), no longer synchronised (`false`), or has never been synchronised (`null`).
     */
    onpremisesSyncEnabled?: pulumi.Input<boolean>;
    /**
     * A set of owners who own this group. Supported object types are Users or Service Principals
     */
    owners?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The preferred language for a Microsoft 365 group, in ISO 639-1 notation.
     */
    preferredLanguage?: pulumi.Input<string>;
    /**
     * If `true`, will return an error if an existing group is found with the same name. Defaults to `false`.
     */
    preventDuplicateNames?: pulumi.Input<boolean>;
    /**
     * A set of provisioning options for a Microsoft 365 group. The only supported value is `Team`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for details. Changing this forces a new resource to be created.
     */
    provisioningOptions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of email addresses for the group that direct to the same group mailbox.
     */
    proxyAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether the group is a security group for controlling access to in-app resources. At least one of `securityEnabled` or `mailEnabled` must be specified. A Microsoft 365 group can be security enabled _and_ mail enabled (see the `types` property).
     */
    securityEnabled?: pulumi.Input<boolean>;
    /**
     * The colour theme for a Microsoft 365 group. Possible values are `Blue`, `Green`, `Orange`, `Pink`, `Purple`, `Red` or `Teal`. By default, no theme is set.
     */
    theme?: pulumi.Input<string>;
    /**
     * A set of group types to configure for the group. Supported values are `DynamicMembership`, which denotes a group with dynamic membership, and `Unified`, which specifies a Microsoft 365 group. Required when `mailEnabled` is true. Changing this forces a new resource to be created.
     *
     * > **Supported Group Types** At present, only security groups and Microsoft 365 groups can be created or managed with this resource. Distribution groups and mail-enabled security groups are not supported. Microsoft 365 groups can be security-enabled.
     */
    types?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The group join policy and group content visibility. Possible values are `Private`, `Public`, or `Hiddenmembership`. Only Microsoft 365 groups can have `Hiddenmembership` visibility and this value must be set when the group is created. By default, security groups will receive `Private` visibility and Microsoft 365 groups will receive `Public` visibility.
     *
     * > **Group Name Uniqueness** Group names are not unique within Azure Active Directory. Use the `preventDuplicateNames` argument to check for existing groups if you want to avoid name collisions.
     */
    visibility?: pulumi.Input<string>;
    /**
     * Whether the group will be written back to the configured on-premises Active Directory when Azure AD Connect is used.
     */
    writebackEnabled?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Group resource.
 */
export interface GroupArgs {
    /**
     * The object IDs of administrative units in which the group is a member. If specified, new groups will be created in the scope of the first administrative unit and added to the others. If empty, new groups will be created at the tenant level.
     *
     * > **Caution** When using the azuread.AdministrativeUnitMember resource, or the `members` property of the azuread.AdministrativeUnit resource, to manage Administrative Unit membership for a group, you will need to use an `ignoreChanges = [administrativeUnitIds]` lifecycle meta argument for the `azuread.Group` resource, in order to avoid a persistent diff.
     */
    administrativeUnitIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Indicates whether this group can be assigned to an Azure Active Directory role. Defaults to `false`. Can only be set to `true` for security-enabled groups. Changing this forces a new resource to be created.
     */
    assignableToRole?: pulumi.Input<boolean>;
    /**
     * Indicates whether new members added to the group will be auto-subscribed to receive email notifications. Can only be set for Unified groups.
     *
     * > **Known Permissions Issue** The `autoSubscribeNewMembers` property can only be set when authenticating as a Member user of the tenant and _not_ when authenticating as a Guest user or as a service principal. Please see the [Microsoft Graph Known Issues](https://docs.microsoft.com/en-us/graph/known-issues#groups) documentation.
     */
    autoSubscribeNewMembers?: pulumi.Input<boolean>;
    /**
     * A set of behaviors for a Microsoft 365 group. Possible values are `AllowOnlyMembersToPost`, `HideGroupInOutlook`, `SkipExchangeInstantOn`, `SubscribeMembersToCalendarEventsDisabled`, `SubscribeNewGroupMembers` and `WelcomeEmailDisabled`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for more details. Changing this forces a new resource to be created.
     */
    behaviors?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The description for the group.
     */
    description?: pulumi.Input<string>;
    /**
     * The display name for the group.
     */
    displayName: pulumi.Input<string>;
    /**
     * A `dynamicMembership` block as documented below. Required when `types` contains `DynamicMembership`. Cannot be used with the `members` property.
     */
    dynamicMembership?: pulumi.Input<inputs.GroupDynamicMembership>;
    /**
     * Indicates whether people external to the organization can send messages to the group. Can only be set for Unified groups.
     *
     * > **Known Permissions Issue** The `externalSendersAllowed` property can only be set when authenticating as a Member user of the tenant and _not_ when authenticating as a Guest user or as a service principal. Please see the [Microsoft Graph Known Issues](https://docs.microsoft.com/en-us/graph/known-issues#groups) documentation.
     */
    externalSendersAllowed?: pulumi.Input<boolean>;
    /**
     * Indicates whether the group is displayed in certain parts of the Outlook user interface: in the Address Book, in address lists for selecting message recipients, and in the Browse Groups dialog for searching groups. Can only be set for Unified groups.
     *
     * > **Known Permissions Issue** The `hideFromAddressLists` property can only be set when authenticating as a Member user of the tenant and _not_ when authenticating as a Guest user or as a service principal. Please see the [Microsoft Graph Known Issues](https://docs.microsoft.com/en-us/graph/known-issues#groups) documentation.
     */
    hideFromAddressLists?: pulumi.Input<boolean>;
    /**
     * Indicates whether the group is displayed in Outlook clients, such as Outlook for Windows and Outlook on the web. Can only be set for Unified groups.
     *
     * > **Known Permissions Issue** The `hideFromOutlookClients` property can only be set when authenticating as a Member user of the tenant and _not_ when authenticating as a Guest user or as a service principal. Please see the [Microsoft Graph Known Issues](https://docs.microsoft.com/en-us/graph/known-issues#groups) documentation.
     */
    hideFromOutlookClients?: pulumi.Input<boolean>;
    /**
     * Whether the group is a mail enabled, with a shared group mailbox. At least one of `mailEnabled` or `securityEnabled` must be specified. Only Microsoft 365 groups can be mail enabled (see the `types` property).
     */
    mailEnabled?: pulumi.Input<boolean>;
    /**
     * The mail alias for the group, unique in the organisation. Required for mail-enabled groups. Changing this forces a new resource to be created.
     */
    mailNickname?: pulumi.Input<string>;
    /**
     * A set of members who should be present in this group. Supported object types are Users, Groups or Service Principals. Cannot be used with the `dynamicMembership` block.
     *
     * !> **Warning** Do not use the `members` property at the same time as the azuread.GroupMember resource for the same group. Doing so will cause a conflict and group members will be removed.
     */
    members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The on-premises group type that the AAD group will be written as, when writeback is enabled. Possible values are `UniversalDistributionGroup`, `UniversalMailEnabledSecurityGroup`, or `UniversalSecurityGroup`.
     */
    onpremisesGroupType?: pulumi.Input<string>;
    /**
     * A set of owners who own this group. Supported object types are Users or Service Principals
     */
    owners?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If `true`, will return an error if an existing group is found with the same name. Defaults to `false`.
     */
    preventDuplicateNames?: pulumi.Input<boolean>;
    /**
     * A set of provisioning options for a Microsoft 365 group. The only supported value is `Team`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for details. Changing this forces a new resource to be created.
     */
    provisioningOptions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether the group is a security group for controlling access to in-app resources. At least one of `securityEnabled` or `mailEnabled` must be specified. A Microsoft 365 group can be security enabled _and_ mail enabled (see the `types` property).
     */
    securityEnabled?: pulumi.Input<boolean>;
    /**
     * The colour theme for a Microsoft 365 group. Possible values are `Blue`, `Green`, `Orange`, `Pink`, `Purple`, `Red` or `Teal`. By default, no theme is set.
     */
    theme?: pulumi.Input<string>;
    /**
     * A set of group types to configure for the group. Supported values are `DynamicMembership`, which denotes a group with dynamic membership, and `Unified`, which specifies a Microsoft 365 group. Required when `mailEnabled` is true. Changing this forces a new resource to be created.
     *
     * > **Supported Group Types** At present, only security groups and Microsoft 365 groups can be created or managed with this resource. Distribution groups and mail-enabled security groups are not supported. Microsoft 365 groups can be security-enabled.
     */
    types?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The group join policy and group content visibility. Possible values are `Private`, `Public`, or `Hiddenmembership`. Only Microsoft 365 groups can have `Hiddenmembership` visibility and this value must be set when the group is created. By default, security groups will receive `Private` visibility and Microsoft 365 groups will receive `Public` visibility.
     *
     * > **Group Name Uniqueness** Group names are not unique within Azure Active Directory. Use the `preventDuplicateNames` argument to check for existing groups if you want to avoid name collisions.
     */
    visibility?: pulumi.Input<string>;
    /**
     * Whether the group will be written back to the configured on-premises Active Directory when Azure AD Connect is used.
     */
    writebackEnabled?: pulumi.Input<boolean>;
}

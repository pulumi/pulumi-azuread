// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.GroupArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.GroupState;
import com.pulumi.azuread.outputs.GroupDynamicMembership;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

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
 * If using the `assignable_to_role` property, this resource additionally requires one of the following application roles: `RoleManagement.ReadWrite.Directory` or `Directory.ReadWrite.All`
 * 
 * If specifying owners for a group, which are user principals, this resource additionally requires one of the following application roles: `User.Read.All`, `User.ReadWrite.All`, `Directory.Read.All` or `Directory.ReadWrite.All`
 * 
 * When authenticated with a user principal, this resource requires one of the following directory roles: `Groups Administrator`, `User Administrator` or `Global Administrator`
 * 
 * When creating this resource in administrative units exclusively, the role `Groups Administrator` is required to be scoped on any administrative unit used.
 * 
 * The `external_senders_allowed`, `auto_subscribe_new_members`, `hide_from_address_lists` and `hide_from_outlook_clients` properties can only be configured when authenticating as a user and cannot be configured when authenticating as a service principal. Additionally, the user being used for authentication must be a Member of the tenant where the group is being managed and _not_ a Guest. This is a known API issue; please see the [Microsoft Graph Known Issues](https://docs.microsoft.com/en-us/graph/known-issues#groups) official documentation.
 * 
 * ## Example Usage
 * 
 * *Basic example*
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuread.AzureadFunctions;
 * import com.pulumi.azuread.Group;
 * import com.pulumi.azuread.GroupArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var current = AzureadFunctions.getClientConfig();
 * 
 *         var example = new Group(&#34;example&#34;, GroupArgs.builder()        
 *             .displayName(&#34;example&#34;)
 *             .owners(current.applyValue(getClientConfigResult -&gt; getClientConfigResult.objectId()))
 *             .securityEnabled(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * *Microsoft 365 group*
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuread.AzureadFunctions;
 * import com.pulumi.azuread.User;
 * import com.pulumi.azuread.UserArgs;
 * import com.pulumi.azuread.Group;
 * import com.pulumi.azuread.GroupArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var current = AzureadFunctions.getClientConfig();
 * 
 *         var groupOwner = new User(&#34;groupOwner&#34;, UserArgs.builder()        
 *             .userPrincipalName(&#34;example-group-owner@example.com&#34;)
 *             .displayName(&#34;Group Owner&#34;)
 *             .mailNickname(&#34;example-group-owner&#34;)
 *             .password(&#34;SecretP@sswd99!&#34;)
 *             .build());
 * 
 *         var example = new Group(&#34;example&#34;, GroupArgs.builder()        
 *             .displayName(&#34;example&#34;)
 *             .mailEnabled(true)
 *             .mailNickname(&#34;ExampleGroup&#34;)
 *             .securityEnabled(true)
 *             .types(&#34;Unified&#34;)
 *             .owners(            
 *                 current.applyValue(getClientConfigResult -&gt; getClientConfigResult.objectId()),
 *                 groupOwner.objectId())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * *Group with members*
 * 
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuread.AzureadFunctions;
 * import com.pulumi.azuread.User;
 * import com.pulumi.azuread.UserArgs;
 * import com.pulumi.azuread.Group;
 * import com.pulumi.azuread.GroupArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var current = AzureadFunctions.getClientConfig();
 * 
 *         var exampleUser = new User(&#34;exampleUser&#34;, UserArgs.builder()        
 *             .displayName(&#34;J Doe&#34;)
 *             .owners(current.applyValue(getClientConfigResult -&gt; getClientConfigResult.objectId()))
 *             .password(&#34;notSecure123&#34;)
 *             .userPrincipalName(&#34;jdoe@example.com&#34;)
 *             .build());
 * 
 *         var exampleGroup = new Group(&#34;exampleGroup&#34;, GroupArgs.builder()        
 *             .displayName(&#34;MyGroup&#34;)
 *             .owners(current.applyValue(getClientConfigResult -&gt; getClientConfigResult.objectId()))
 *             .securityEnabled(true)
 *             .members(exampleUser.objectId())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * *Group with dynamic membership*
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuread.AzureadFunctions;
 * import com.pulumi.azuread.Group;
 * import com.pulumi.azuread.GroupArgs;
 * import com.pulumi.azuread.inputs.GroupDynamicMembershipArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var current = AzureadFunctions.getClientConfig();
 * 
 *         var example = new Group(&#34;example&#34;, GroupArgs.builder()        
 *             .displayName(&#34;MyGroup&#34;)
 *             .owners(current.applyValue(getClientConfigResult -&gt; getClientConfigResult.objectId()))
 *             .securityEnabled(true)
 *             .types(&#34;DynamicMembership&#34;)
 *             .dynamicMembership(GroupDynamicMembershipArgs.builder()
 *                 .enabled(true)
 *                 .rule(&#34;user.department -eq \&#34;Sales\&#34;&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Groups can be imported using their object ID, e.g.
 * 
 * ```sh
 * $ pulumi import azuread:index/group:Group my_group 00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuread:index/group:Group")
public class Group extends com.pulumi.resources.CustomResource {
    /**
     * The object IDs of administrative units in which the group is a member. If specified, new groups will be created in the scope of the first administrative unit and added to the others. If empty, new groups will be created at the tenant level.
     * 
     * !&gt; **Warning** Do not use the `administrative_unit_ids` property at the same time as the azuread.AdministrativeUnitMember resource, _for the same group_. Doing so will cause a conflict and administrative unit members will be removed.
     * 
     */
    @Export(name="administrativeUnitIds", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> administrativeUnitIds;

    /**
     * @return The object IDs of administrative units in which the group is a member. If specified, new groups will be created in the scope of the first administrative unit and added to the others. If empty, new groups will be created at the tenant level.
     * 
     * !&gt; **Warning** Do not use the `administrative_unit_ids` property at the same time as the azuread.AdministrativeUnitMember resource, _for the same group_. Doing so will cause a conflict and administrative unit members will be removed.
     * 
     */
    public Output<Optional<List<String>>> administrativeUnitIds() {
        return Codegen.optional(this.administrativeUnitIds);
    }
    /**
     * Indicates whether this group can be assigned to an Azure Active Directory role. Defaults to `false`. Can only be set to `true` for security-enabled groups. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="assignableToRole", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> assignableToRole;

    /**
     * @return Indicates whether this group can be assigned to an Azure Active Directory role. Defaults to `false`. Can only be set to `true` for security-enabled groups. Changing this forces a new resource to be created.
     * 
     */
    public Output<Optional<Boolean>> assignableToRole() {
        return Codegen.optional(this.assignableToRole);
    }
    /**
     * Indicates whether new members added to the group will be auto-subscribed to receive email notifications. Can only be set for Unified groups.
     * 
     * &gt; **Known Permissions Issue** The `auto_subscribe_new_members` property can only be set when authenticating as a Member user of the tenant and _not_ when authenticating as a Guest user or as a service principal. Please see the [Microsoft Graph Known Issues](https://docs.microsoft.com/en-us/graph/known-issues#groups) documentation.
     * 
     */
    @Export(name="autoSubscribeNewMembers", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> autoSubscribeNewMembers;

    /**
     * @return Indicates whether new members added to the group will be auto-subscribed to receive email notifications. Can only be set for Unified groups.
     * 
     * &gt; **Known Permissions Issue** The `auto_subscribe_new_members` property can only be set when authenticating as a Member user of the tenant and _not_ when authenticating as a Guest user or as a service principal. Please see the [Microsoft Graph Known Issues](https://docs.microsoft.com/en-us/graph/known-issues#groups) documentation.
     * 
     */
    public Output<Boolean> autoSubscribeNewMembers() {
        return this.autoSubscribeNewMembers;
    }
    /**
     * A set of behaviors for a Microsoft 365 group. Possible values are `AllowOnlyMembersToPost`, `HideGroupInOutlook`, `SubscribeMembersToCalendarEventsDisabled`, `SubscribeNewGroupMembers` and `WelcomeEmailDisabled`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for more details. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="behaviors", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> behaviors;

    /**
     * @return A set of behaviors for a Microsoft 365 group. Possible values are `AllowOnlyMembersToPost`, `HideGroupInOutlook`, `SubscribeMembersToCalendarEventsDisabled`, `SubscribeNewGroupMembers` and `WelcomeEmailDisabled`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for more details. Changing this forces a new resource to be created.
     * 
     */
    public Output<Optional<List<String>>> behaviors() {
        return Codegen.optional(this.behaviors);
    }
    /**
     * The description for the group.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description for the group.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The display name for the group.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    /**
     * @return The display name for the group.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * A `dynamic_membership` block as documented below. Required when `types` contains `DynamicMembership`. Cannot be used with the `members` property.
     * 
     */
    @Export(name="dynamicMembership", refs={GroupDynamicMembership.class}, tree="[0]")
    private Output</* @Nullable */ GroupDynamicMembership> dynamicMembership;

    /**
     * @return A `dynamic_membership` block as documented below. Required when `types` contains `DynamicMembership`. Cannot be used with the `members` property.
     * 
     */
    public Output<Optional<GroupDynamicMembership>> dynamicMembership() {
        return Codegen.optional(this.dynamicMembership);
    }
    /**
     * Indicates whether people external to the organization can send messages to the group. Can only be set for Unified groups.
     * 
     * &gt; **Known Permissions Issue** The `external_senders_allowed` property can only be set when authenticating as a Member user of the tenant and _not_ when authenticating as a Guest user or as a service principal. Please see the [Microsoft Graph Known Issues](https://docs.microsoft.com/en-us/graph/known-issues#groups) documentation.
     * 
     */
    @Export(name="externalSendersAllowed", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> externalSendersAllowed;

    /**
     * @return Indicates whether people external to the organization can send messages to the group. Can only be set for Unified groups.
     * 
     * &gt; **Known Permissions Issue** The `external_senders_allowed` property can only be set when authenticating as a Member user of the tenant and _not_ when authenticating as a Guest user or as a service principal. Please see the [Microsoft Graph Known Issues](https://docs.microsoft.com/en-us/graph/known-issues#groups) documentation.
     * 
     */
    public Output<Boolean> externalSendersAllowed() {
        return this.externalSendersAllowed;
    }
    /**
     * Indicates whether the group is displayed in certain parts of the Outlook user interface: in the Address Book, in address lists for selecting message recipients, and in the Browse Groups dialog for searching groups. Can only be set for Unified groups.
     * 
     * &gt; **Known Permissions Issue** The `hide_from_address_lists` property can only be set when authenticating as a Member user of the tenant and _not_ when authenticating as a Guest user or as a service principal. Please see the [Microsoft Graph Known Issues](https://docs.microsoft.com/en-us/graph/known-issues#groups) documentation.
     * 
     */
    @Export(name="hideFromAddressLists", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> hideFromAddressLists;

    /**
     * @return Indicates whether the group is displayed in certain parts of the Outlook user interface: in the Address Book, in address lists for selecting message recipients, and in the Browse Groups dialog for searching groups. Can only be set for Unified groups.
     * 
     * &gt; **Known Permissions Issue** The `hide_from_address_lists` property can only be set when authenticating as a Member user of the tenant and _not_ when authenticating as a Guest user or as a service principal. Please see the [Microsoft Graph Known Issues](https://docs.microsoft.com/en-us/graph/known-issues#groups) documentation.
     * 
     */
    public Output<Boolean> hideFromAddressLists() {
        return this.hideFromAddressLists;
    }
    /**
     * Indicates whether the group is displayed in Outlook clients, such as Outlook for Windows and Outlook on the web. Can only be set for Unified groups.
     * 
     * &gt; **Known Permissions Issue** The `hide_from_outlook_clients` property can only be set when authenticating as a Member user of the tenant and _not_ when authenticating as a Guest user or as a service principal. Please see the [Microsoft Graph Known Issues](https://docs.microsoft.com/en-us/graph/known-issues#groups) documentation.
     * 
     */
    @Export(name="hideFromOutlookClients", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> hideFromOutlookClients;

    /**
     * @return Indicates whether the group is displayed in Outlook clients, such as Outlook for Windows and Outlook on the web. Can only be set for Unified groups.
     * 
     * &gt; **Known Permissions Issue** The `hide_from_outlook_clients` property can only be set when authenticating as a Member user of the tenant and _not_ when authenticating as a Guest user or as a service principal. Please see the [Microsoft Graph Known Issues](https://docs.microsoft.com/en-us/graph/known-issues#groups) documentation.
     * 
     */
    public Output<Boolean> hideFromOutlookClients() {
        return this.hideFromOutlookClients;
    }
    /**
     * The SMTP address for the group.
     * 
     */
    @Export(name="mail", refs={String.class}, tree="[0]")
    private Output<String> mail;

    /**
     * @return The SMTP address for the group.
     * 
     */
    public Output<String> mail() {
        return this.mail;
    }
    /**
     * Whether the group is a mail enabled, with a shared group mailbox. At least one of `mail_enabled` or `security_enabled` must be specified. Only Microsoft 365 groups can be mail enabled (see the `types` property).
     * 
     */
    @Export(name="mailEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> mailEnabled;

    /**
     * @return Whether the group is a mail enabled, with a shared group mailbox. At least one of `mail_enabled` or `security_enabled` must be specified. Only Microsoft 365 groups can be mail enabled (see the `types` property).
     * 
     */
    public Output<Optional<Boolean>> mailEnabled() {
        return Codegen.optional(this.mailEnabled);
    }
    /**
     * The mail alias for the group, unique in the organisation. Required for mail-enabled groups. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="mailNickname", refs={String.class}, tree="[0]")
    private Output<String> mailNickname;

    /**
     * @return The mail alias for the group, unique in the organisation. Required for mail-enabled groups. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> mailNickname() {
        return this.mailNickname;
    }
    /**
     * A set of members who should be present in this group. Supported object types are Users, Groups or Service Principals. Cannot be used with the `dynamic_membership` block.
     * 
     * !&gt; **Warning** Do not use the `members` property at the same time as the azuread.GroupMember resource for the same group. Doing so will cause a conflict and group members will be removed.
     * 
     */
    @Export(name="members", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> members;

    /**
     * @return A set of members who should be present in this group. Supported object types are Users, Groups or Service Principals. Cannot be used with the `dynamic_membership` block.
     * 
     * !&gt; **Warning** Do not use the `members` property at the same time as the azuread.GroupMember resource for the same group. Doing so will cause a conflict and group members will be removed.
     * 
     */
    public Output<List<String>> members() {
        return this.members;
    }
    /**
     * The object ID of the group.
     * 
     */
    @Export(name="objectId", refs={String.class}, tree="[0]")
    private Output<String> objectId;

    /**
     * @return The object ID of the group.
     * 
     */
    public Output<String> objectId() {
        return this.objectId;
    }
    /**
     * The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.
     * 
     */
    @Export(name="onpremisesDomainName", refs={String.class}, tree="[0]")
    private Output<String> onpremisesDomainName;

    /**
     * @return The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.
     * 
     */
    public Output<String> onpremisesDomainName() {
        return this.onpremisesDomainName;
    }
    /**
     * The on-premises group type that the AAD group will be written as, when writeback is enabled. Possible values are `UniversalDistributionGroup`, `UniversalMailEnabledSecurityGroup`, or `UniversalSecurityGroup`.
     * 
     */
    @Export(name="onpremisesGroupType", refs={String.class}, tree="[0]")
    private Output<String> onpremisesGroupType;

    /**
     * @return The on-premises group type that the AAD group will be written as, when writeback is enabled. Possible values are `UniversalDistributionGroup`, `UniversalMailEnabledSecurityGroup`, or `UniversalSecurityGroup`.
     * 
     */
    public Output<String> onpremisesGroupType() {
        return this.onpremisesGroupType;
    }
    /**
     * The on-premises NetBIOS name, synchronised from the on-premises directory when Azure AD Connect is used.
     * 
     */
    @Export(name="onpremisesNetbiosName", refs={String.class}, tree="[0]")
    private Output<String> onpremisesNetbiosName;

    /**
     * @return The on-premises NetBIOS name, synchronised from the on-premises directory when Azure AD Connect is used.
     * 
     */
    public Output<String> onpremisesNetbiosName() {
        return this.onpremisesNetbiosName;
    }
    /**
     * The on-premises SAM account name, synchronised from the on-premises directory when Azure AD Connect is used.
     * 
     */
    @Export(name="onpremisesSamAccountName", refs={String.class}, tree="[0]")
    private Output<String> onpremisesSamAccountName;

    /**
     * @return The on-premises SAM account name, synchronised from the on-premises directory when Azure AD Connect is used.
     * 
     */
    public Output<String> onpremisesSamAccountName() {
        return this.onpremisesSamAccountName;
    }
    /**
     * The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.
     * 
     */
    @Export(name="onpremisesSecurityIdentifier", refs={String.class}, tree="[0]")
    private Output<String> onpremisesSecurityIdentifier;

    /**
     * @return The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.
     * 
     */
    public Output<String> onpremisesSecurityIdentifier() {
        return this.onpremisesSecurityIdentifier;
    }
    /**
     * Whether this group is synchronised from an on-premises directory (`true`), no longer synchronised (`false`), or has never been synchronised (`null`).
     * 
     */
    @Export(name="onpremisesSyncEnabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> onpremisesSyncEnabled;

    /**
     * @return Whether this group is synchronised from an on-premises directory (`true`), no longer synchronised (`false`), or has never been synchronised (`null`).
     * 
     */
    public Output<Boolean> onpremisesSyncEnabled() {
        return this.onpremisesSyncEnabled;
    }
    /**
     * A set of owners who own this group. Supported object types are Users or Service Principals
     * 
     */
    @Export(name="owners", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> owners;

    /**
     * @return A set of owners who own this group. Supported object types are Users or Service Principals
     * 
     */
    public Output<List<String>> owners() {
        return this.owners;
    }
    /**
     * The preferred language for a Microsoft 365 group, in ISO 639-1 notation.
     * 
     */
    @Export(name="preferredLanguage", refs={String.class}, tree="[0]")
    private Output<String> preferredLanguage;

    /**
     * @return The preferred language for a Microsoft 365 group, in ISO 639-1 notation.
     * 
     */
    public Output<String> preferredLanguage() {
        return this.preferredLanguage;
    }
    /**
     * If `true`, will return an error if an existing group is found with the same name. Defaults to `false`.
     * 
     */
    @Export(name="preventDuplicateNames", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> preventDuplicateNames;

    /**
     * @return If `true`, will return an error if an existing group is found with the same name. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> preventDuplicateNames() {
        return Codegen.optional(this.preventDuplicateNames);
    }
    /**
     * A set of provisioning options for a Microsoft 365 group. The only supported value is `Team`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for details. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="provisioningOptions", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> provisioningOptions;

    /**
     * @return A set of provisioning options for a Microsoft 365 group. The only supported value is `Team`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for details. Changing this forces a new resource to be created.
     * 
     */
    public Output<Optional<List<String>>> provisioningOptions() {
        return Codegen.optional(this.provisioningOptions);
    }
    /**
     * List of email addresses for the group that direct to the same group mailbox.
     * 
     */
    @Export(name="proxyAddresses", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> proxyAddresses;

    /**
     * @return List of email addresses for the group that direct to the same group mailbox.
     * 
     */
    public Output<List<String>> proxyAddresses() {
        return this.proxyAddresses;
    }
    /**
     * Whether the group is a security group for controlling access to in-app resources. At least one of `security_enabled` or `mail_enabled` must be specified. A Microsoft 365 group can be security enabled _and_ mail enabled (see the `types` property).
     * 
     */
    @Export(name="securityEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> securityEnabled;

    /**
     * @return Whether the group is a security group for controlling access to in-app resources. At least one of `security_enabled` or `mail_enabled` must be specified. A Microsoft 365 group can be security enabled _and_ mail enabled (see the `types` property).
     * 
     */
    public Output<Optional<Boolean>> securityEnabled() {
        return Codegen.optional(this.securityEnabled);
    }
    /**
     * The colour theme for a Microsoft 365 group. Possible values are `Blue`, `Green`, `Orange`, `Pink`, `Purple`, `Red` or `Teal`. By default, no theme is set.
     * 
     */
    @Export(name="theme", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> theme;

    /**
     * @return The colour theme for a Microsoft 365 group. Possible values are `Blue`, `Green`, `Orange`, `Pink`, `Purple`, `Red` or `Teal`. By default, no theme is set.
     * 
     */
    public Output<Optional<String>> theme() {
        return Codegen.optional(this.theme);
    }
    /**
     * A set of group types to configure for the group. Supported values are `DynamicMembership`, which denotes a group with dynamic membership, and `Unified`, which specifies a Microsoft 365 group. Required when `mail_enabled` is true. Changing this forces a new resource to be created.
     * 
     * &gt; **Supported Group Types** At present, only security groups and Microsoft 365 groups can be created or managed with this resource. Distribution groups and mail-enabled security groups are not supported. Microsoft 365 groups can be security-enabled.
     * 
     */
    @Export(name="types", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> types;

    /**
     * @return A set of group types to configure for the group. Supported values are `DynamicMembership`, which denotes a group with dynamic membership, and `Unified`, which specifies a Microsoft 365 group. Required when `mail_enabled` is true. Changing this forces a new resource to be created.
     * 
     * &gt; **Supported Group Types** At present, only security groups and Microsoft 365 groups can be created or managed with this resource. Distribution groups and mail-enabled security groups are not supported. Microsoft 365 groups can be security-enabled.
     * 
     */
    public Output<Optional<List<String>>> types() {
        return Codegen.optional(this.types);
    }
    /**
     * The group join policy and group content visibility. Possible values are `Private`, `Public`, or `Hiddenmembership`. Only Microsoft 365 groups can have `Hiddenmembership` visibility and this value must be set when the group is created. By default, security groups will receive `Private` visibility and Microsoft 365 groups will receive `Public` visibility.
     * 
     * &gt; **Group Name Uniqueness** Group names are not unique within Azure Active Directory. Use the `prevent_duplicate_names` argument to check for existing groups if you want to avoid name collisions.
     * 
     */
    @Export(name="visibility", refs={String.class}, tree="[0]")
    private Output<String> visibility;

    /**
     * @return The group join policy and group content visibility. Possible values are `Private`, `Public`, or `Hiddenmembership`. Only Microsoft 365 groups can have `Hiddenmembership` visibility and this value must be set when the group is created. By default, security groups will receive `Private` visibility and Microsoft 365 groups will receive `Public` visibility.
     * 
     * &gt; **Group Name Uniqueness** Group names are not unique within Azure Active Directory. Use the `prevent_duplicate_names` argument to check for existing groups if you want to avoid name collisions.
     * 
     */
    public Output<String> visibility() {
        return this.visibility;
    }
    /**
     * Whether the group will be written back to the configured on-premises Active Directory when Azure AD Connect is used.
     * 
     */
    @Export(name="writebackEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> writebackEnabled;

    /**
     * @return Whether the group will be written back to the configured on-premises Active Directory when Azure AD Connect is used.
     * 
     */
    public Output<Optional<Boolean>> writebackEnabled() {
        return Codegen.optional(this.writebackEnabled);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Group(String name) {
        this(name, GroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Group(String name, GroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Group(String name, GroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/group:Group", name, args == null ? GroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Group(String name, Output<String> id, @Nullable GroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/group:Group", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Group get(String name, Output<String> id, @Nullable GroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Group(name, id, state, options);
    }
}

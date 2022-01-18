// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD
{
    /// <summary>
    /// Manages a group within Azure Active Directory.
    /// 
    /// ## API Permissions
    /// 
    /// The following API permissions are required in order to use this resource.
    /// 
    /// When authenticated with a service principal, this resource requires one of the following application roles: `Group.ReadWrite.All` or `Directory.ReadWrite.All`
    /// 
    /// If using the `assignable_to_role` property, this resource additionally requires one of the following application roles: `RoleManagement.ReadWrite.Directory` or `Directory.ReadWrite.All`
    /// 
    /// If specifying owners for a group, which are user principals, this resource additionally requires one of the following application roles: `User.Read.All`, `User.ReadWrite.All`, `Directory.Read.All` or `Directory.ReadWrite.All`
    /// 
    /// When authenticated with a user principal, this resource requires one of the following directory roles: `Groups Administrator`, `User Administrator` or `Global Administrator`
    /// 
    /// ## Import
    /// 
    /// Groups can be imported using their object ID, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azuread:index/group:Group my_group 00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureADResourceType("azuread:index/group:Group")]
    public partial class Group : Pulumi.CustomResource
    {
        /// <summary>
        /// Indicates whether this group can be assigned to an Azure Active Directory role. Can only be `true` for security-enabled groups. Changing this forces a new resource to be created.
        /// </summary>
        [Output("assignableToRole")]
        public Output<bool?> AssignableToRole { get; private set; } = null!;

        /// <summary>
        /// A set of behaviors for a Microsoft 365 group. Possible values are `AllowOnlyMembersToPost`, `HideGroupInOutlook`, `SubscribeNewGroupMembers` and `WelcomeEmailDisabled`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for more details. Changing this forces a new resource to be created.
        /// </summary>
        [Output("behaviors")]
        public Output<ImmutableArray<string>> Behaviors { get; private set; } = null!;

        /// <summary>
        /// The description for the group.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The display name for the group.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// A `dynamic_membership` block as documented below. Required when `types` contains `DynamicMembership`. Cannot be used with the `members` property.
        /// </summary>
        [Output("dynamicMembership")]
        public Output<Outputs.GroupDynamicMembership?> DynamicMembership { get; private set; } = null!;

        /// <summary>
        /// The SMTP address for the group.
        /// </summary>
        [Output("mail")]
        public Output<string> Mail { get; private set; } = null!;

        /// <summary>
        /// Whether the group is a mail enabled, with a shared group mailbox. At least one of `mail_enabled` or `security_enabled` must be specified. Only Microsoft 365 groups can be mail enabled (see the `types` property).
        /// </summary>
        [Output("mailEnabled")]
        public Output<bool?> MailEnabled { get; private set; } = null!;

        /// <summary>
        /// The mail alias for the group, unique in the organisation. Required for mail-enabled groups. Changing this forces a new resource to be created.
        /// </summary>
        [Output("mailNickname")]
        public Output<string> MailNickname { get; private set; } = null!;

        /// <summary>
        /// A set of members who should be present in this group. Supported object types are Users, Groups or Service Principals. Cannot be used with the `dynamic_membership` block.
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<string>> Members { get; private set; } = null!;

        /// <summary>
        /// The object ID of the group.
        /// </summary>
        [Output("objectId")]
        public Output<string> ObjectId { get; private set; } = null!;

        /// <summary>
        /// The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.
        /// </summary>
        [Output("onpremisesDomainName")]
        public Output<string> OnpremisesDomainName { get; private set; } = null!;

        /// <summary>
        /// The on-premises NetBIOS name, synchronised from the on-premises directory when Azure AD Connect is used.
        /// </summary>
        [Output("onpremisesNetbiosName")]
        public Output<string> OnpremisesNetbiosName { get; private set; } = null!;

        /// <summary>
        /// The on-premises SAM account name, synchronised from the on-premises directory when Azure AD Connect is used.
        /// </summary>
        [Output("onpremisesSamAccountName")]
        public Output<string> OnpremisesSamAccountName { get; private set; } = null!;

        /// <summary>
        /// The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.
        /// </summary>
        [Output("onpremisesSecurityIdentifier")]
        public Output<string> OnpremisesSecurityIdentifier { get; private set; } = null!;

        /// <summary>
        /// Whether this group is synchronised from an on-premises directory (`true`), no longer synchronised (`false`), or has never been synchronised (`null`).
        /// </summary>
        [Output("onpremisesSyncEnabled")]
        public Output<bool> OnpremisesSyncEnabled { get; private set; } = null!;

        /// <summary>
        /// A set of owners who own this group. Supported object types are Users or Service Principals
        /// </summary>
        [Output("owners")]
        public Output<ImmutableArray<string>> Owners { get; private set; } = null!;

        /// <summary>
        /// The preferred language for a Microsoft 365 group, in ISO 639-1 notation.
        /// </summary>
        [Output("preferredLanguage")]
        public Output<string> PreferredLanguage { get; private set; } = null!;

        /// <summary>
        /// If `true`, will return an error if an existing group is found with the same name. Defaults to `false`.
        /// </summary>
        [Output("preventDuplicateNames")]
        public Output<bool?> PreventDuplicateNames { get; private set; } = null!;

        /// <summary>
        /// A set of provisioning options for a Microsoft 365 group. The only supported value is `Team`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for details. Changing this forces a new resource to be created.
        /// </summary>
        [Output("provisioningOptions")]
        public Output<ImmutableArray<string>> ProvisioningOptions { get; private set; } = null!;

        /// <summary>
        /// List of email addresses for the group that direct to the same group mailbox.
        /// </summary>
        [Output("proxyAddresses")]
        public Output<ImmutableArray<string>> ProxyAddresses { get; private set; } = null!;

        /// <summary>
        /// Whether the group is a security group for controlling access to in-app resources. At least one of `security_enabled` or `mail_enabled` must be specified. A Microsoft 365 group can be security enabled _and_ mail enabled (see the `types` property).
        /// </summary>
        [Output("securityEnabled")]
        public Output<bool?> SecurityEnabled { get; private set; } = null!;

        /// <summary>
        /// The colour theme for a Microsoft 365 group. Possible values are `Blue`, `Green`, `Orange`, `Pink`, `Purple`, `Red` or `Teal`. By default, no theme is set.
        /// </summary>
        [Output("theme")]
        public Output<string?> Theme { get; private set; } = null!;

        /// <summary>
        /// A set of group types to configure for the group. Supported values are `DynamicMembership`, which denotes a group with dynamic membership, and `Unified`, which specifies a Microsoft 365 group. Required when `mail_enabled` is true. Changing this forces a new resource to be created.
        /// </summary>
        [Output("types")]
        public Output<ImmutableArray<string>> Types { get; private set; } = null!;

        /// <summary>
        /// The group join policy and group content visibility. Possible values are `Private`, `Public`, or `Hiddenmembership`. Only Microsoft 365 groups can have `Hiddenmembership` visibility and this value must be set when the group is created. By default, security groups will receive `Private` visibility and Microsoft 365 groups will receive `Public` visibility.
        /// </summary>
        [Output("visibility")]
        public Output<string> Visibility { get; private set; } = null!;


        /// <summary>
        /// Create a Group resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Group(string name, GroupArgs args, CustomResourceOptions? options = null)
            : base("azuread:index/group:Group", name, args ?? new GroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Group(string name, Input<string> id, GroupState? state = null, CustomResourceOptions? options = null)
            : base("azuread:index/group:Group", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Group resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Group Get(string name, Input<string> id, GroupState? state = null, CustomResourceOptions? options = null)
        {
            return new Group(name, id, state, options);
        }
    }

    public sealed class GroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether this group can be assigned to an Azure Active Directory role. Can only be `true` for security-enabled groups. Changing this forces a new resource to be created.
        /// </summary>
        [Input("assignableToRole")]
        public Input<bool>? AssignableToRole { get; set; }

        [Input("behaviors")]
        private InputList<string>? _behaviors;

        /// <summary>
        /// A set of behaviors for a Microsoft 365 group. Possible values are `AllowOnlyMembersToPost`, `HideGroupInOutlook`, `SubscribeNewGroupMembers` and `WelcomeEmailDisabled`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for more details. Changing this forces a new resource to be created.
        /// </summary>
        public InputList<string> Behaviors
        {
            get => _behaviors ?? (_behaviors = new InputList<string>());
            set => _behaviors = value;
        }

        /// <summary>
        /// The description for the group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name for the group.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// A `dynamic_membership` block as documented below. Required when `types` contains `DynamicMembership`. Cannot be used with the `members` property.
        /// </summary>
        [Input("dynamicMembership")]
        public Input<Inputs.GroupDynamicMembershipArgs>? DynamicMembership { get; set; }

        /// <summary>
        /// Whether the group is a mail enabled, with a shared group mailbox. At least one of `mail_enabled` or `security_enabled` must be specified. Only Microsoft 365 groups can be mail enabled (see the `types` property).
        /// </summary>
        [Input("mailEnabled")]
        public Input<bool>? MailEnabled { get; set; }

        /// <summary>
        /// The mail alias for the group, unique in the organisation. Required for mail-enabled groups. Changing this forces a new resource to be created.
        /// </summary>
        [Input("mailNickname")]
        public Input<string>? MailNickname { get; set; }

        [Input("members")]
        private InputList<string>? _members;

        /// <summary>
        /// A set of members who should be present in this group. Supported object types are Users, Groups or Service Principals. Cannot be used with the `dynamic_membership` block.
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        [Input("owners")]
        private InputList<string>? _owners;

        /// <summary>
        /// A set of owners who own this group. Supported object types are Users or Service Principals
        /// </summary>
        public InputList<string> Owners
        {
            get => _owners ?? (_owners = new InputList<string>());
            set => _owners = value;
        }

        /// <summary>
        /// If `true`, will return an error if an existing group is found with the same name. Defaults to `false`.
        /// </summary>
        [Input("preventDuplicateNames")]
        public Input<bool>? PreventDuplicateNames { get; set; }

        [Input("provisioningOptions")]
        private InputList<string>? _provisioningOptions;

        /// <summary>
        /// A set of provisioning options for a Microsoft 365 group. The only supported value is `Team`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for details. Changing this forces a new resource to be created.
        /// </summary>
        public InputList<string> ProvisioningOptions
        {
            get => _provisioningOptions ?? (_provisioningOptions = new InputList<string>());
            set => _provisioningOptions = value;
        }

        /// <summary>
        /// Whether the group is a security group for controlling access to in-app resources. At least one of `security_enabled` or `mail_enabled` must be specified. A Microsoft 365 group can be security enabled _and_ mail enabled (see the `types` property).
        /// </summary>
        [Input("securityEnabled")]
        public Input<bool>? SecurityEnabled { get; set; }

        /// <summary>
        /// The colour theme for a Microsoft 365 group. Possible values are `Blue`, `Green`, `Orange`, `Pink`, `Purple`, `Red` or `Teal`. By default, no theme is set.
        /// </summary>
        [Input("theme")]
        public Input<string>? Theme { get; set; }

        [Input("types")]
        private InputList<string>? _types;

        /// <summary>
        /// A set of group types to configure for the group. Supported values are `DynamicMembership`, which denotes a group with dynamic membership, and `Unified`, which specifies a Microsoft 365 group. Required when `mail_enabled` is true. Changing this forces a new resource to be created.
        /// </summary>
        public InputList<string> Types
        {
            get => _types ?? (_types = new InputList<string>());
            set => _types = value;
        }

        /// <summary>
        /// The group join policy and group content visibility. Possible values are `Private`, `Public`, or `Hiddenmembership`. Only Microsoft 365 groups can have `Hiddenmembership` visibility and this value must be set when the group is created. By default, security groups will receive `Private` visibility and Microsoft 365 groups will receive `Public` visibility.
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        public GroupArgs()
        {
        }
    }

    public sealed class GroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether this group can be assigned to an Azure Active Directory role. Can only be `true` for security-enabled groups. Changing this forces a new resource to be created.
        /// </summary>
        [Input("assignableToRole")]
        public Input<bool>? AssignableToRole { get; set; }

        [Input("behaviors")]
        private InputList<string>? _behaviors;

        /// <summary>
        /// A set of behaviors for a Microsoft 365 group. Possible values are `AllowOnlyMembersToPost`, `HideGroupInOutlook`, `SubscribeNewGroupMembers` and `WelcomeEmailDisabled`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for more details. Changing this forces a new resource to be created.
        /// </summary>
        public InputList<string> Behaviors
        {
            get => _behaviors ?? (_behaviors = new InputList<string>());
            set => _behaviors = value;
        }

        /// <summary>
        /// The description for the group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name for the group.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// A `dynamic_membership` block as documented below. Required when `types` contains `DynamicMembership`. Cannot be used with the `members` property.
        /// </summary>
        [Input("dynamicMembership")]
        public Input<Inputs.GroupDynamicMembershipGetArgs>? DynamicMembership { get; set; }

        /// <summary>
        /// The SMTP address for the group.
        /// </summary>
        [Input("mail")]
        public Input<string>? Mail { get; set; }

        /// <summary>
        /// Whether the group is a mail enabled, with a shared group mailbox. At least one of `mail_enabled` or `security_enabled` must be specified. Only Microsoft 365 groups can be mail enabled (see the `types` property).
        /// </summary>
        [Input("mailEnabled")]
        public Input<bool>? MailEnabled { get; set; }

        /// <summary>
        /// The mail alias for the group, unique in the organisation. Required for mail-enabled groups. Changing this forces a new resource to be created.
        /// </summary>
        [Input("mailNickname")]
        public Input<string>? MailNickname { get; set; }

        [Input("members")]
        private InputList<string>? _members;

        /// <summary>
        /// A set of members who should be present in this group. Supported object types are Users, Groups or Service Principals. Cannot be used with the `dynamic_membership` block.
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// The object ID of the group.
        /// </summary>
        [Input("objectId")]
        public Input<string>? ObjectId { get; set; }

        /// <summary>
        /// The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.
        /// </summary>
        [Input("onpremisesDomainName")]
        public Input<string>? OnpremisesDomainName { get; set; }

        /// <summary>
        /// The on-premises NetBIOS name, synchronised from the on-premises directory when Azure AD Connect is used.
        /// </summary>
        [Input("onpremisesNetbiosName")]
        public Input<string>? OnpremisesNetbiosName { get; set; }

        /// <summary>
        /// The on-premises SAM account name, synchronised from the on-premises directory when Azure AD Connect is used.
        /// </summary>
        [Input("onpremisesSamAccountName")]
        public Input<string>? OnpremisesSamAccountName { get; set; }

        /// <summary>
        /// The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.
        /// </summary>
        [Input("onpremisesSecurityIdentifier")]
        public Input<string>? OnpremisesSecurityIdentifier { get; set; }

        /// <summary>
        /// Whether this group is synchronised from an on-premises directory (`true`), no longer synchronised (`false`), or has never been synchronised (`null`).
        /// </summary>
        [Input("onpremisesSyncEnabled")]
        public Input<bool>? OnpremisesSyncEnabled { get; set; }

        [Input("owners")]
        private InputList<string>? _owners;

        /// <summary>
        /// A set of owners who own this group. Supported object types are Users or Service Principals
        /// </summary>
        public InputList<string> Owners
        {
            get => _owners ?? (_owners = new InputList<string>());
            set => _owners = value;
        }

        /// <summary>
        /// The preferred language for a Microsoft 365 group, in ISO 639-1 notation.
        /// </summary>
        [Input("preferredLanguage")]
        public Input<string>? PreferredLanguage { get; set; }

        /// <summary>
        /// If `true`, will return an error if an existing group is found with the same name. Defaults to `false`.
        /// </summary>
        [Input("preventDuplicateNames")]
        public Input<bool>? PreventDuplicateNames { get; set; }

        [Input("provisioningOptions")]
        private InputList<string>? _provisioningOptions;

        /// <summary>
        /// A set of provisioning options for a Microsoft 365 group. The only supported value is `Team`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for details. Changing this forces a new resource to be created.
        /// </summary>
        public InputList<string> ProvisioningOptions
        {
            get => _provisioningOptions ?? (_provisioningOptions = new InputList<string>());
            set => _provisioningOptions = value;
        }

        [Input("proxyAddresses")]
        private InputList<string>? _proxyAddresses;

        /// <summary>
        /// List of email addresses for the group that direct to the same group mailbox.
        /// </summary>
        public InputList<string> ProxyAddresses
        {
            get => _proxyAddresses ?? (_proxyAddresses = new InputList<string>());
            set => _proxyAddresses = value;
        }

        /// <summary>
        /// Whether the group is a security group for controlling access to in-app resources. At least one of `security_enabled` or `mail_enabled` must be specified. A Microsoft 365 group can be security enabled _and_ mail enabled (see the `types` property).
        /// </summary>
        [Input("securityEnabled")]
        public Input<bool>? SecurityEnabled { get; set; }

        /// <summary>
        /// The colour theme for a Microsoft 365 group. Possible values are `Blue`, `Green`, `Orange`, `Pink`, `Purple`, `Red` or `Teal`. By default, no theme is set.
        /// </summary>
        [Input("theme")]
        public Input<string>? Theme { get; set; }

        [Input("types")]
        private InputList<string>? _types;

        /// <summary>
        /// A set of group types to configure for the group. Supported values are `DynamicMembership`, which denotes a group with dynamic membership, and `Unified`, which specifies a Microsoft 365 group. Required when `mail_enabled` is true. Changing this forces a new resource to be created.
        /// </summary>
        public InputList<string> Types
        {
            get => _types ?? (_types = new InputList<string>());
            set => _types = value;
        }

        /// <summary>
        /// The group join policy and group content visibility. Possible values are `Private`, `Public`, or `Hiddenmembership`. Only Microsoft 365 groups can have `Hiddenmembership` visibility and this value must be set when the group is created. By default, security groups will receive `Private` visibility and Microsoft 365 groups will receive `Public` visibility.
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        public GroupState()
        {
        }
    }
}

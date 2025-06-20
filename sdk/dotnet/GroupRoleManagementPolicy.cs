// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD
{
    /// <summary>
    /// Manage a role policy for an Azure AD group.
    /// 
    /// ## API Permissions
    /// 
    /// The following API permissions are required in order to use this resource.
    /// 
    /// When authenticated with a service principal, this resource requires the `RoleManagementPolicy.ReadWrite.AzureADGroup` Microsoft Graph API permissions.
    /// 
    /// When authenticated with a user principal, this resource requires `Global Administrator` directory role, or the `Privileged Role Administrator` role in Identity Governance.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureAD = Pulumi.AzureAD;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new AzureAD.Group("example", new()
    ///     {
    ///         DisplayName = "group-name",
    ///         SecurityEnabled = true,
    ///     });
    /// 
    ///     var member = new AzureAD.User("member", new()
    ///     {
    ///         UserPrincipalName = "jdoe@example.com",
    ///         DisplayName = "J. Doe",
    ///         MailNickname = "jdoe",
    ///         Password = "SecretP@sswd99!",
    ///     });
    /// 
    ///     var exampleGroupRoleManagementPolicy = new AzureAD.GroupRoleManagementPolicy("example", new()
    ///     {
    ///         GroupId = example.Id,
    ///         RoleId = "member",
    ///         ActiveAssignmentRules = new AzureAD.Inputs.GroupRoleManagementPolicyActiveAssignmentRulesArgs
    ///         {
    ///             ExpireAfter = "P365D",
    ///         },
    ///         EligibleAssignmentRules = new AzureAD.Inputs.GroupRoleManagementPolicyEligibleAssignmentRulesArgs
    ///         {
    ///             ExpirationRequired = false,
    ///         },
    ///         NotificationRules = new AzureAD.Inputs.GroupRoleManagementPolicyNotificationRulesArgs
    ///         {
    ///             EligibleAssignments = new AzureAD.Inputs.GroupRoleManagementPolicyNotificationRulesEligibleAssignmentsArgs
    ///             {
    ///                 ApproverNotifications = new AzureAD.Inputs.GroupRoleManagementPolicyNotificationRulesEligibleAssignmentsApproverNotificationsArgs
    ///                 {
    ///                     NotificationLevel = "Critical",
    ///                     DefaultRecipients = false,
    ///                     AdditionalRecipients = new[]
    ///                     {
    ///                         "someone@example.com",
    ///                         "someone.else@example.com",
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Because these policies are created automatically by Entra ID, they will auto-import on first use.
    /// </summary>
    [AzureADResourceType("azuread:index/groupRoleManagementPolicy:GroupRoleManagementPolicy")]
    public partial class GroupRoleManagementPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// An `activation_rules` block as defined below.
        /// </summary>
        [Output("activationRules")]
        public Output<Outputs.GroupRoleManagementPolicyActivationRules> ActivationRules { get; private set; } = null!;

        /// <summary>
        /// An `active_assignment_rules` block as defined below.
        /// </summary>
        [Output("activeAssignmentRules")]
        public Output<Outputs.GroupRoleManagementPolicyActiveAssignmentRules> ActiveAssignmentRules { get; private set; } = null!;

        /// <summary>
        /// (String) The description of this policy.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// (String) The display name of this policy.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// An `eligible_assignment_rules` block as defined below.
        /// </summary>
        [Output("eligibleAssignmentRules")]
        public Output<Outputs.GroupRoleManagementPolicyEligibleAssignmentRules> EligibleAssignmentRules { get; private set; } = null!;

        /// <summary>
        /// The ID of the Azure AD group for which the policy applies.
        /// </summary>
        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        /// <summary>
        /// A `notification_rules` block as defined below.
        /// </summary>
        [Output("notificationRules")]
        public Output<Outputs.GroupRoleManagementPolicyNotificationRules> NotificationRules { get; private set; } = null!;

        /// <summary>
        /// The type of assignment this policy coveres. Can be either `member` or `owner`.
        /// </summary>
        [Output("roleId")]
        public Output<string> RoleId { get; private set; } = null!;


        /// <summary>
        /// Create a GroupRoleManagementPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupRoleManagementPolicy(string name, GroupRoleManagementPolicyArgs args, CustomResourceOptions? options = null)
            : base("azuread:index/groupRoleManagementPolicy:GroupRoleManagementPolicy", name, args ?? new GroupRoleManagementPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GroupRoleManagementPolicy(string name, Input<string> id, GroupRoleManagementPolicyState? state = null, CustomResourceOptions? options = null)
            : base("azuread:index/groupRoleManagementPolicy:GroupRoleManagementPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GroupRoleManagementPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GroupRoleManagementPolicy Get(string name, Input<string> id, GroupRoleManagementPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new GroupRoleManagementPolicy(name, id, state, options);
        }
    }

    public sealed class GroupRoleManagementPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An `activation_rules` block as defined below.
        /// </summary>
        [Input("activationRules")]
        public Input<Inputs.GroupRoleManagementPolicyActivationRulesArgs>? ActivationRules { get; set; }

        /// <summary>
        /// An `active_assignment_rules` block as defined below.
        /// </summary>
        [Input("activeAssignmentRules")]
        public Input<Inputs.GroupRoleManagementPolicyActiveAssignmentRulesArgs>? ActiveAssignmentRules { get; set; }

        /// <summary>
        /// An `eligible_assignment_rules` block as defined below.
        /// </summary>
        [Input("eligibleAssignmentRules")]
        public Input<Inputs.GroupRoleManagementPolicyEligibleAssignmentRulesArgs>? EligibleAssignmentRules { get; set; }

        /// <summary>
        /// The ID of the Azure AD group for which the policy applies.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        /// <summary>
        /// A `notification_rules` block as defined below.
        /// </summary>
        [Input("notificationRules")]
        public Input<Inputs.GroupRoleManagementPolicyNotificationRulesArgs>? NotificationRules { get; set; }

        /// <summary>
        /// The type of assignment this policy coveres. Can be either `member` or `owner`.
        /// </summary>
        [Input("roleId", required: true)]
        public Input<string> RoleId { get; set; } = null!;

        public GroupRoleManagementPolicyArgs()
        {
        }
        public static new GroupRoleManagementPolicyArgs Empty => new GroupRoleManagementPolicyArgs();
    }

    public sealed class GroupRoleManagementPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An `activation_rules` block as defined below.
        /// </summary>
        [Input("activationRules")]
        public Input<Inputs.GroupRoleManagementPolicyActivationRulesGetArgs>? ActivationRules { get; set; }

        /// <summary>
        /// An `active_assignment_rules` block as defined below.
        /// </summary>
        [Input("activeAssignmentRules")]
        public Input<Inputs.GroupRoleManagementPolicyActiveAssignmentRulesGetArgs>? ActiveAssignmentRules { get; set; }

        /// <summary>
        /// (String) The description of this policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// (String) The display name of this policy.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// An `eligible_assignment_rules` block as defined below.
        /// </summary>
        [Input("eligibleAssignmentRules")]
        public Input<Inputs.GroupRoleManagementPolicyEligibleAssignmentRulesGetArgs>? EligibleAssignmentRules { get; set; }

        /// <summary>
        /// The ID of the Azure AD group for which the policy applies.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// A `notification_rules` block as defined below.
        /// </summary>
        [Input("notificationRules")]
        public Input<Inputs.GroupRoleManagementPolicyNotificationRulesGetArgs>? NotificationRules { get; set; }

        /// <summary>
        /// The type of assignment this policy coveres. Can be either `member` or `owner`.
        /// </summary>
        [Input("roleId")]
        public Input<string>? RoleId { get; set; }

        public GroupRoleManagementPolicyState()
        {
        }
        public static new GroupRoleManagementPolicyState Empty => new GroupRoleManagementPolicyState();
    }
}

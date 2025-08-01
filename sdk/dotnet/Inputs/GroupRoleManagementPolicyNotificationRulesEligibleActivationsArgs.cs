// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class GroupRoleManagementPolicyNotificationRulesEligibleActivationsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Admin notification settings
        /// </summary>
        [Input("adminNotifications")]
        public Input<Inputs.GroupRoleManagementPolicyNotificationRulesEligibleActivationsAdminNotificationsArgs>? AdminNotifications { get; set; }

        /// <summary>
        /// Approver notification settings
        /// </summary>
        [Input("approverNotifications")]
        public Input<Inputs.GroupRoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsArgs>? ApproverNotifications { get; set; }

        /// <summary>
        /// Assignee notification settings
        /// </summary>
        [Input("assigneeNotifications")]
        public Input<Inputs.GroupRoleManagementPolicyNotificationRulesEligibleActivationsAssigneeNotificationsArgs>? AssigneeNotifications { get; set; }

        public GroupRoleManagementPolicyNotificationRulesEligibleActivationsArgs()
        {
        }
        public static new GroupRoleManagementPolicyNotificationRulesEligibleActivationsArgs Empty => new GroupRoleManagementPolicyNotificationRulesEligibleActivationsArgs();
    }
}

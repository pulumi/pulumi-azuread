// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class GroupRoleManagementPolicyNotificationRulesEligibleActivationsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Admin notification settings
        /// </summary>
        [Input("adminNotifications")]
        public Input<Inputs.GroupRoleManagementPolicyNotificationRulesEligibleActivationsAdminNotificationsGetArgs>? AdminNotifications { get; set; }

        /// <summary>
        /// Approver notification settings
        /// </summary>
        [Input("approverNotifications")]
        public Input<Inputs.GroupRoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsGetArgs>? ApproverNotifications { get; set; }

        /// <summary>
        /// Assignee notification settings
        /// </summary>
        [Input("assigneeNotifications")]
        public Input<Inputs.GroupRoleManagementPolicyNotificationRulesEligibleActivationsAssigneeNotificationsGetArgs>? AssigneeNotifications { get; set; }

        public GroupRoleManagementPolicyNotificationRulesEligibleActivationsGetArgs()
        {
        }
        public static new GroupRoleManagementPolicyNotificationRulesEligibleActivationsGetArgs Empty => new GroupRoleManagementPolicyNotificationRulesEligibleActivationsGetArgs();
    }
}

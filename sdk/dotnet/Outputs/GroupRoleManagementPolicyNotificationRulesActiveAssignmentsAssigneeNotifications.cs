// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Outputs
{

    [OutputType]
    public sealed class GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotifications
    {
        /// <summary>
        /// The additional recipients to notify
        /// </summary>
        public readonly ImmutableArray<string> AdditionalRecipients;
        /// <summary>
        /// Whether the default recipients are notified
        /// </summary>
        public readonly bool DefaultRecipients;
        /// <summary>
        /// What level of notifications are sent
        /// </summary>
        public readonly string NotificationLevel;

        [OutputConstructor]
        private GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotifications(
            ImmutableArray<string> additionalRecipients,

            bool defaultRecipients,

            string notificationLevel)
        {
            AdditionalRecipients = additionalRecipients;
            DefaultRecipients = defaultRecipients;
            NotificationLevel = notificationLevel;
        }
    }
}

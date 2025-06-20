// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Outputs
{

    [OutputType]
    public sealed class GroupRoleManagementPolicyActivationRulesApprovalStagePrimaryApprover
    {
        /// <summary>
        /// The ID of the object which will act as an approver.
        /// </summary>
        public readonly string ObjectId;
        /// <summary>
        /// The type of object acting as an approver. Possible options are `singleUser` and `groupMembers`.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private GroupRoleManagementPolicyActivationRulesApprovalStagePrimaryApprover(
            string objectId,

            string? type)
        {
            ObjectId = objectId;
            Type = type;
        }
    }
}

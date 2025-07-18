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
    public sealed class GroupRoleManagementPolicyActivationRulesApprovalStage
    {
        /// <summary>
        /// The IDs of the users or groups who can approve the activation
        /// </summary>
        public readonly ImmutableArray<Outputs.GroupRoleManagementPolicyActivationRulesApprovalStagePrimaryApprover> PrimaryApprovers;

        [OutputConstructor]
        private GroupRoleManagementPolicyActivationRulesApprovalStage(ImmutableArray<Outputs.GroupRoleManagementPolicyActivationRulesApprovalStagePrimaryApprover> primaryApprovers)
        {
            PrimaryApprovers = primaryApprovers;
        }
    }
}

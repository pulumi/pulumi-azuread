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
    public sealed class AccessPackageAssignmentPolicyApprovalSettingsApprovalStage
    {
        /// <summary>
        /// Whether alternative approvers are enabled.
        /// </summary>
        public readonly bool? AlternativeApprovalEnabled;
        /// <summary>
        /// A block specifying alternative approvers when escalation is enabled and the primary approvers do not respond before the escalation time, as documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.AccessPackageAssignmentPolicyApprovalSettingsApprovalStageAlternativeApprover> AlternativeApprovers;
        /// <summary>
        /// Maximum number of days within which a request must be approved. If a request is not approved within this time period after it is made, it will be automatically rejected.
        /// </summary>
        public readonly int ApprovalTimeoutInDays;
        /// <summary>
        /// Whether an approver must provide a justification for their decision. Justification is visible to other approvers and the requestor.
        /// </summary>
        public readonly bool? ApproverJustificationRequired;
        /// <summary>
        /// Number of days before the request is forwarded to alternative approvers.
        /// </summary>
        public readonly int? EnableAlternativeApprovalInDays;
        /// <summary>
        /// A block specifying the users who will be asked to approve requests, as documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApprover> PrimaryApprovers;

        [OutputConstructor]
        private AccessPackageAssignmentPolicyApprovalSettingsApprovalStage(
            bool? alternativeApprovalEnabled,

            ImmutableArray<Outputs.AccessPackageAssignmentPolicyApprovalSettingsApprovalStageAlternativeApprover> alternativeApprovers,

            int approvalTimeoutInDays,

            bool? approverJustificationRequired,

            int? enableAlternativeApprovalInDays,

            ImmutableArray<Outputs.AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApprover> primaryApprovers)
        {
            AlternativeApprovalEnabled = alternativeApprovalEnabled;
            AlternativeApprovers = alternativeApprovers;
            ApprovalTimeoutInDays = approvalTimeoutInDays;
            ApproverJustificationRequired = approverJustificationRequired;
            EnableAlternativeApprovalInDays = enableAlternativeApprovalInDays;
            PrimaryApprovers = primaryApprovers;
        }
    }
}

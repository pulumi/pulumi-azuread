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
    public sealed class AccessPackageAssignmentPolicyApprovalSettingsApprovalStageAlternativeApprover
    {
        /// <summary>
        /// For a user in an approval stage, this property indicates whether the user is a backup fallback approver
        /// </summary>
        public readonly bool? Backup;
        /// <summary>
        /// The object ID of the subject
        /// </summary>
        public readonly string? ObjectId;
        /// <summary>
        /// Type of users
        /// </summary>
        public readonly string SubjectType;

        [OutputConstructor]
        private AccessPackageAssignmentPolicyApprovalSettingsApprovalStageAlternativeApprover(
            bool? backup,

            string? objectId,

            string subjectType)
        {
            Backup = backup;
            ObjectId = objectId;
            SubjectType = subjectType;
        }
    }
}

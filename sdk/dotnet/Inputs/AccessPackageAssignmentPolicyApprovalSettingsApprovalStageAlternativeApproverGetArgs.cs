// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class AccessPackageAssignmentPolicyApprovalSettingsApprovalStageAlternativeApproverGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("backup")]
        public Input<bool>? Backup { get; set; }

        [Input("objectId")]
        public Input<string>? ObjectId { get; set; }

        [Input("subjectType", required: true)]
        public Input<string> SubjectType { get; set; } = null!;

        public AccessPackageAssignmentPolicyApprovalSettingsApprovalStageAlternativeApproverGetArgs()
        {
        }
        public static new AccessPackageAssignmentPolicyApprovalSettingsApprovalStageAlternativeApproverGetArgs Empty => new AccessPackageAssignmentPolicyApprovalSettingsApprovalStageAlternativeApproverGetArgs();
    }
}

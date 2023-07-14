// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class AccessPackageAssignmentPolicyAssignmentReviewSettingsGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessRecommendationEnabled")]
        public Input<bool>? AccessRecommendationEnabled { get; set; }

        [Input("accessReviewTimeoutBehavior")]
        public Input<string>? AccessReviewTimeoutBehavior { get; set; }

        [Input("approverJustificationRequired")]
        public Input<bool>? ApproverJustificationRequired { get; set; }

        [Input("durationInDays")]
        public Input<int>? DurationInDays { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("reviewFrequency")]
        public Input<string>? ReviewFrequency { get; set; }

        [Input("reviewType")]
        public Input<string>? ReviewType { get; set; }

        [Input("reviewers")]
        private InputList<Inputs.AccessPackageAssignmentPolicyAssignmentReviewSettingsReviewerGetArgs>? _reviewers;
        public InputList<Inputs.AccessPackageAssignmentPolicyAssignmentReviewSettingsReviewerGetArgs> Reviewers
        {
            get => _reviewers ?? (_reviewers = new InputList<Inputs.AccessPackageAssignmentPolicyAssignmentReviewSettingsReviewerGetArgs>());
            set => _reviewers = value;
        }

        [Input("startingOn")]
        public Input<string>? StartingOn { get; set; }

        public AccessPackageAssignmentPolicyAssignmentReviewSettingsGetArgs()
        {
        }
        public static new AccessPackageAssignmentPolicyAssignmentReviewSettingsGetArgs Empty => new AccessPackageAssignmentPolicyAssignmentReviewSettingsGetArgs();
    }
}

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
    public sealed class AccessPackageAssignmentPolicyAssignmentReviewSettings
    {
        /// <summary>
        /// Whether to show the reviewer decision helpers. If enabled, system recommendations based on users' access information will be shown to the reviewers. The reviewer will be recommended to approve the review if the user has signed-in at least once during the last 30 days. The reviewer will be recommended to deny the review if the user has not signed-in during the last 30 days.
        /// </summary>
        public readonly bool? AccessRecommendationEnabled;
        /// <summary>
        /// Specifies the actions the system takes if reviewers don't respond in time. Valid values are `keepAccess`, `removeAccess`, or `acceptAccessRecommendation`.
        /// </summary>
        public readonly string? AccessReviewTimeoutBehavior;
        /// <summary>
        /// Whether a reviewer needs to provide a justification for their decision. Justification is visible to other reviewers and the requestor.
        /// </summary>
        public readonly bool? ApproverJustificationRequired;
        /// <summary>
        /// How many days each occurrence of the access review series will run.
        /// </summary>
        public readonly int? DurationInDays;
        /// <summary>
        /// Whether to enable assignment review.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// This will determine how often the access review campaign runs, valid values are `weekly`, `monthly`, `quarterly`, `halfyearly`, or `annual`.
        /// </summary>
        public readonly string? ReviewFrequency;
        /// <summary>
        /// Self-review or specific reviewers. Valid values are `Manager`, `Reviewers`, or `Self`.
        /// </summary>
        public readonly string? ReviewType;
        /// <summary>
        /// One or more `reviewer` blocks to specify the users who will be reviewers (when `review_type` is `Reviewers`), as documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.AccessPackageAssignmentPolicyAssignmentReviewSettingsReviewer> Reviewers;
        /// <summary>
        /// This is the date the access review campaign will start on, formatted as an RFC3339 date string in UTC(e.g. 2018-01-01T01:02:03Z), default is now. Once an access review has been created, you cannot update its start date
        /// </summary>
        public readonly string? StartingOn;

        [OutputConstructor]
        private AccessPackageAssignmentPolicyAssignmentReviewSettings(
            bool? accessRecommendationEnabled,

            string? accessReviewTimeoutBehavior,

            bool? approverJustificationRequired,

            int? durationInDays,

            bool? enabled,

            string? reviewFrequency,

            string? reviewType,

            ImmutableArray<Outputs.AccessPackageAssignmentPolicyAssignmentReviewSettingsReviewer> reviewers,

            string? startingOn)
        {
            AccessRecommendationEnabled = accessRecommendationEnabled;
            AccessReviewTimeoutBehavior = accessReviewTimeoutBehavior;
            ApproverJustificationRequired = approverJustificationRequired;
            DurationInDays = durationInDays;
            Enabled = enabled;
            ReviewFrequency = reviewFrequency;
            ReviewType = reviewType;
            Reviewers = reviewers;
            StartingOn = startingOn;
        }
    }
}

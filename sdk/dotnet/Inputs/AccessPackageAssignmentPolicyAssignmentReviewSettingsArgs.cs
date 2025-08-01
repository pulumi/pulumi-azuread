// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class AccessPackageAssignmentPolicyAssignmentReviewSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to show the reviewer decision helpers. If enabled, system recommendations based on users' access information will be shown to the reviewers. The reviewer will be recommended to approve the review if the user has signed-in at least once during the last 30 days. The reviewer will be recommended to deny the review if the user has not signed-in during the last 30 days.
        /// </summary>
        [Input("accessRecommendationEnabled")]
        public Input<bool>? AccessRecommendationEnabled { get; set; }

        /// <summary>
        /// Specifies the actions the system takes if reviewers don't respond in time. Valid values are `keepAccess`, `removeAccess`, or `acceptAccessRecommendation`.
        /// </summary>
        [Input("accessReviewTimeoutBehavior")]
        public Input<string>? AccessReviewTimeoutBehavior { get; set; }

        /// <summary>
        /// Whether a reviewer needs to provide a justification for their decision. Justification is visible to other reviewers and the requestor.
        /// </summary>
        [Input("approverJustificationRequired")]
        public Input<bool>? ApproverJustificationRequired { get; set; }

        /// <summary>
        /// How many days each occurrence of the access review series will run.
        /// </summary>
        [Input("durationInDays")]
        public Input<int>? DurationInDays { get; set; }

        /// <summary>
        /// Whether to enable assignment review.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// This will determine how often the access review campaign runs, valid values are `weekly`, `monthly`, `quarterly`, `halfyearly`, or `annual`.
        /// </summary>
        [Input("reviewFrequency")]
        public Input<string>? ReviewFrequency { get; set; }

        /// <summary>
        /// Self-review or specific reviewers. Valid values are `Manager`, `Reviewers`, or `Self`.
        /// </summary>
        [Input("reviewType")]
        public Input<string>? ReviewType { get; set; }

        [Input("reviewers")]
        private InputList<Inputs.AccessPackageAssignmentPolicyAssignmentReviewSettingsReviewerArgs>? _reviewers;

        /// <summary>
        /// One or more `reviewer` blocks to specify the users who will be reviewers (when `review_type` is `Reviewers`), as documented below.
        /// </summary>
        public InputList<Inputs.AccessPackageAssignmentPolicyAssignmentReviewSettingsReviewerArgs> Reviewers
        {
            get => _reviewers ?? (_reviewers = new InputList<Inputs.AccessPackageAssignmentPolicyAssignmentReviewSettingsReviewerArgs>());
            set => _reviewers = value;
        }

        /// <summary>
        /// This is the date the access review campaign will start on, formatted as an RFC3339 date string in UTC(e.g. 2018-01-01T01:02:03Z), default is now. Once an access review has been created, you cannot update its start date
        /// </summary>
        [Input("startingOn")]
        public Input<string>? StartingOn { get; set; }

        public AccessPackageAssignmentPolicyAssignmentReviewSettingsArgs()
        {
        }
        public static new AccessPackageAssignmentPolicyAssignmentReviewSettingsArgs Empty => new AccessPackageAssignmentPolicyAssignmentReviewSettingsArgs();
    }
}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.azuread.outputs.AccessPackageAssignmentPolicyAssignmentReviewSettingsReviewer;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class AccessPackageAssignmentPolicyAssignmentReviewSettings {
    /**
     * @return Whether to show the reviewer decision helpers. If enabled, system recommendations based on users&#39; access information will be shown to the reviewers. The reviewer will be recommended to approve the review if the user has signed-in at least once during the last 30 days. The reviewer will be recommended to deny the review if the user has not signed-in during the last 30 days.
     * 
     */
    private @Nullable Boolean accessRecommendationEnabled;
    /**
     * @return Specifies the actions the system takes if reviewers don&#39;t respond in time. Valid values are `keepAccess`, `removeAccess`, or `acceptAccessRecommendation`.
     * 
     */
    private @Nullable String accessReviewTimeoutBehavior;
    /**
     * @return Whether a reviewer needs to provide a justification for their decision. Justification is visible to other reviewers and the requestor.
     * 
     */
    private @Nullable Boolean approverJustificationRequired;
    /**
     * @return How many days each occurrence of the access review series will run.
     * 
     */
    private @Nullable Integer durationInDays;
    /**
     * @return Whether to enable assignment review.
     * 
     */
    private @Nullable Boolean enabled;
    /**
     * @return This will determine how often the access review campaign runs, valid values are `weekly`, `monthly`, `quarterly`, `halfyearly`, or `annual`.
     * 
     */
    private @Nullable String reviewFrequency;
    /**
     * @return Self-review or specific reviewers. Valid values are `Manager`, `Reviewers`, or `Self`.
     * 
     */
    private @Nullable String reviewType;
    /**
     * @return One or more `reviewer` blocks to specify the users who will be reviewers (when `review_type` is `Reviewers`), as documented below.
     * 
     */
    private @Nullable List<AccessPackageAssignmentPolicyAssignmentReviewSettingsReviewer> reviewers;
    /**
     * @return This is the date the access review campaign will start on, formatted as an RFC3339 date string in UTC(e.g. 2018-01-01T01:02:03Z), default is now. Once an access review has been created, you cannot update its start date
     * 
     */
    private @Nullable String startingOn;

    private AccessPackageAssignmentPolicyAssignmentReviewSettings() {}
    /**
     * @return Whether to show the reviewer decision helpers. If enabled, system recommendations based on users&#39; access information will be shown to the reviewers. The reviewer will be recommended to approve the review if the user has signed-in at least once during the last 30 days. The reviewer will be recommended to deny the review if the user has not signed-in during the last 30 days.
     * 
     */
    public Optional<Boolean> accessRecommendationEnabled() {
        return Optional.ofNullable(this.accessRecommendationEnabled);
    }
    /**
     * @return Specifies the actions the system takes if reviewers don&#39;t respond in time. Valid values are `keepAccess`, `removeAccess`, or `acceptAccessRecommendation`.
     * 
     */
    public Optional<String> accessReviewTimeoutBehavior() {
        return Optional.ofNullable(this.accessReviewTimeoutBehavior);
    }
    /**
     * @return Whether a reviewer needs to provide a justification for their decision. Justification is visible to other reviewers and the requestor.
     * 
     */
    public Optional<Boolean> approverJustificationRequired() {
        return Optional.ofNullable(this.approverJustificationRequired);
    }
    /**
     * @return How many days each occurrence of the access review series will run.
     * 
     */
    public Optional<Integer> durationInDays() {
        return Optional.ofNullable(this.durationInDays);
    }
    /**
     * @return Whether to enable assignment review.
     * 
     */
    public Optional<Boolean> enabled() {
        return Optional.ofNullable(this.enabled);
    }
    /**
     * @return This will determine how often the access review campaign runs, valid values are `weekly`, `monthly`, `quarterly`, `halfyearly`, or `annual`.
     * 
     */
    public Optional<String> reviewFrequency() {
        return Optional.ofNullable(this.reviewFrequency);
    }
    /**
     * @return Self-review or specific reviewers. Valid values are `Manager`, `Reviewers`, or `Self`.
     * 
     */
    public Optional<String> reviewType() {
        return Optional.ofNullable(this.reviewType);
    }
    /**
     * @return One or more `reviewer` blocks to specify the users who will be reviewers (when `review_type` is `Reviewers`), as documented below.
     * 
     */
    public List<AccessPackageAssignmentPolicyAssignmentReviewSettingsReviewer> reviewers() {
        return this.reviewers == null ? List.of() : this.reviewers;
    }
    /**
     * @return This is the date the access review campaign will start on, formatted as an RFC3339 date string in UTC(e.g. 2018-01-01T01:02:03Z), default is now. Once an access review has been created, you cannot update its start date
     * 
     */
    public Optional<String> startingOn() {
        return Optional.ofNullable(this.startingOn);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AccessPackageAssignmentPolicyAssignmentReviewSettings defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean accessRecommendationEnabled;
        private @Nullable String accessReviewTimeoutBehavior;
        private @Nullable Boolean approverJustificationRequired;
        private @Nullable Integer durationInDays;
        private @Nullable Boolean enabled;
        private @Nullable String reviewFrequency;
        private @Nullable String reviewType;
        private @Nullable List<AccessPackageAssignmentPolicyAssignmentReviewSettingsReviewer> reviewers;
        private @Nullable String startingOn;
        public Builder() {}
        public Builder(AccessPackageAssignmentPolicyAssignmentReviewSettings defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accessRecommendationEnabled = defaults.accessRecommendationEnabled;
    	      this.accessReviewTimeoutBehavior = defaults.accessReviewTimeoutBehavior;
    	      this.approverJustificationRequired = defaults.approverJustificationRequired;
    	      this.durationInDays = defaults.durationInDays;
    	      this.enabled = defaults.enabled;
    	      this.reviewFrequency = defaults.reviewFrequency;
    	      this.reviewType = defaults.reviewType;
    	      this.reviewers = defaults.reviewers;
    	      this.startingOn = defaults.startingOn;
        }

        @CustomType.Setter
        public Builder accessRecommendationEnabled(@Nullable Boolean accessRecommendationEnabled) {

            this.accessRecommendationEnabled = accessRecommendationEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder accessReviewTimeoutBehavior(@Nullable String accessReviewTimeoutBehavior) {

            this.accessReviewTimeoutBehavior = accessReviewTimeoutBehavior;
            return this;
        }
        @CustomType.Setter
        public Builder approverJustificationRequired(@Nullable Boolean approverJustificationRequired) {

            this.approverJustificationRequired = approverJustificationRequired;
            return this;
        }
        @CustomType.Setter
        public Builder durationInDays(@Nullable Integer durationInDays) {

            this.durationInDays = durationInDays;
            return this;
        }
        @CustomType.Setter
        public Builder enabled(@Nullable Boolean enabled) {

            this.enabled = enabled;
            return this;
        }
        @CustomType.Setter
        public Builder reviewFrequency(@Nullable String reviewFrequency) {

            this.reviewFrequency = reviewFrequency;
            return this;
        }
        @CustomType.Setter
        public Builder reviewType(@Nullable String reviewType) {

            this.reviewType = reviewType;
            return this;
        }
        @CustomType.Setter
        public Builder reviewers(@Nullable List<AccessPackageAssignmentPolicyAssignmentReviewSettingsReviewer> reviewers) {

            this.reviewers = reviewers;
            return this;
        }
        public Builder reviewers(AccessPackageAssignmentPolicyAssignmentReviewSettingsReviewer... reviewers) {
            return reviewers(List.of(reviewers));
        }
        @CustomType.Setter
        public Builder startingOn(@Nullable String startingOn) {

            this.startingOn = startingOn;
            return this;
        }
        public AccessPackageAssignmentPolicyAssignmentReviewSettings build() {
            final var _resultValue = new AccessPackageAssignmentPolicyAssignmentReviewSettings();
            _resultValue.accessRecommendationEnabled = accessRecommendationEnabled;
            _resultValue.accessReviewTimeoutBehavior = accessReviewTimeoutBehavior;
            _resultValue.approverJustificationRequired = approverJustificationRequired;
            _resultValue.durationInDays = durationInDays;
            _resultValue.enabled = enabled;
            _resultValue.reviewFrequency = reviewFrequency;
            _resultValue.reviewType = reviewType;
            _resultValue.reviewers = reviewers;
            _resultValue.startingOn = startingOn;
            return _resultValue;
        }
    }
}

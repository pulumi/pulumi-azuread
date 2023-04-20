// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.inputs.AccessPackageAssignmentPolicyApprovalSettingsArgs;
import com.pulumi.azuread.inputs.AccessPackageAssignmentPolicyAssignmentReviewSettingsArgs;
import com.pulumi.azuread.inputs.AccessPackageAssignmentPolicyQuestionArgs;
import com.pulumi.azuread.inputs.AccessPackageAssignmentPolicyRequestorSettingsArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AccessPackageAssignmentPolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final AccessPackageAssignmentPolicyArgs Empty = new AccessPackageAssignmentPolicyArgs();

    /**
     * The ID of the access package that will contain the policy.
     * 
     */
    @Import(name="accessPackageId", required=true)
    private Output<String> accessPackageId;

    /**
     * @return The ID of the access package that will contain the policy.
     * 
     */
    public Output<String> accessPackageId() {
        return this.accessPackageId;
    }

    /**
     * An `approval_settings` block to specify whether approvals are required and how they are obtained, as documented below.
     * 
     */
    @Import(name="approvalSettings")
    private @Nullable Output<AccessPackageAssignmentPolicyApprovalSettingsArgs> approvalSettings;

    /**
     * @return An `approval_settings` block to specify whether approvals are required and how they are obtained, as documented below.
     * 
     */
    public Optional<Output<AccessPackageAssignmentPolicyApprovalSettingsArgs>> approvalSettings() {
        return Optional.ofNullable(this.approvalSettings);
    }

    /**
     * An `assignment_review_settings` block, to specify whether assignment review is needed and how it is conducted, as documented below.
     * 
     */
    @Import(name="assignmentReviewSettings")
    private @Nullable Output<AccessPackageAssignmentPolicyAssignmentReviewSettingsArgs> assignmentReviewSettings;

    /**
     * @return An `assignment_review_settings` block, to specify whether assignment review is needed and how it is conducted, as documented below.
     * 
     */
    public Optional<Output<AccessPackageAssignmentPolicyAssignmentReviewSettingsArgs>> assignmentReviewSettings() {
        return Optional.ofNullable(this.assignmentReviewSettings);
    }

    /**
     * The description of the policy.
     * 
     */
    @Import(name="description", required=true)
    private Output<String> description;

    /**
     * @return The description of the policy.
     * 
     */
    public Output<String> description() {
        return this.description;
    }

    /**
     * The display name of the policy.
     * 
     */
    @Import(name="displayName", required=true)
    private Output<String> displayName;

    /**
     * @return The display name of the policy.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }

    /**
     * How many days this assignment is valid for.
     * 
     */
    @Import(name="durationInDays")
    private @Nullable Output<Integer> durationInDays;

    /**
     * @return How many days this assignment is valid for.
     * 
     */
    public Optional<Output<Integer>> durationInDays() {
        return Optional.ofNullable(this.durationInDays);
    }

    /**
     * The date that this assignment expires, formatted as an RFC3339 date string in UTC(e.g. 2018-01-01T01:02:03Z).
     * 
     */
    @Import(name="expirationDate")
    private @Nullable Output<String> expirationDate;

    /**
     * @return The date that this assignment expires, formatted as an RFC3339 date string in UTC(e.g. 2018-01-01T01:02:03Z).
     * 
     */
    public Optional<Output<String>> expirationDate() {
        return Optional.ofNullable(this.expirationDate);
    }

    /**
     * Whether users will be able to request extension of their access to this package before their access expires.
     * 
     */
    @Import(name="extensionEnabled")
    private @Nullable Output<Boolean> extensionEnabled;

    /**
     * @return Whether users will be able to request extension of their access to this package before their access expires.
     * 
     */
    public Optional<Output<Boolean>> extensionEnabled() {
        return Optional.ofNullable(this.extensionEnabled);
    }

    /**
     * One or more `question` blocks for the requestor, as documented below.
     * 
     */
    @Import(name="questions")
    private @Nullable Output<List<AccessPackageAssignmentPolicyQuestionArgs>> questions;

    /**
     * @return One or more `question` blocks for the requestor, as documented below.
     * 
     */
    public Optional<Output<List<AccessPackageAssignmentPolicyQuestionArgs>>> questions() {
        return Optional.ofNullable(this.questions);
    }

    /**
     * A `requestor_settings` block to configure the users who can request access, as documented below.
     * 
     */
    @Import(name="requestorSettings")
    private @Nullable Output<AccessPackageAssignmentPolicyRequestorSettingsArgs> requestorSettings;

    /**
     * @return A `requestor_settings` block to configure the users who can request access, as documented below.
     * 
     */
    public Optional<Output<AccessPackageAssignmentPolicyRequestorSettingsArgs>> requestorSettings() {
        return Optional.ofNullable(this.requestorSettings);
    }

    private AccessPackageAssignmentPolicyArgs() {}

    private AccessPackageAssignmentPolicyArgs(AccessPackageAssignmentPolicyArgs $) {
        this.accessPackageId = $.accessPackageId;
        this.approvalSettings = $.approvalSettings;
        this.assignmentReviewSettings = $.assignmentReviewSettings;
        this.description = $.description;
        this.displayName = $.displayName;
        this.durationInDays = $.durationInDays;
        this.expirationDate = $.expirationDate;
        this.extensionEnabled = $.extensionEnabled;
        this.questions = $.questions;
        this.requestorSettings = $.requestorSettings;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AccessPackageAssignmentPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AccessPackageAssignmentPolicyArgs $;

        public Builder() {
            $ = new AccessPackageAssignmentPolicyArgs();
        }

        public Builder(AccessPackageAssignmentPolicyArgs defaults) {
            $ = new AccessPackageAssignmentPolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessPackageId The ID of the access package that will contain the policy.
         * 
         * @return builder
         * 
         */
        public Builder accessPackageId(Output<String> accessPackageId) {
            $.accessPackageId = accessPackageId;
            return this;
        }

        /**
         * @param accessPackageId The ID of the access package that will contain the policy.
         * 
         * @return builder
         * 
         */
        public Builder accessPackageId(String accessPackageId) {
            return accessPackageId(Output.of(accessPackageId));
        }

        /**
         * @param approvalSettings An `approval_settings` block to specify whether approvals are required and how they are obtained, as documented below.
         * 
         * @return builder
         * 
         */
        public Builder approvalSettings(@Nullable Output<AccessPackageAssignmentPolicyApprovalSettingsArgs> approvalSettings) {
            $.approvalSettings = approvalSettings;
            return this;
        }

        /**
         * @param approvalSettings An `approval_settings` block to specify whether approvals are required and how they are obtained, as documented below.
         * 
         * @return builder
         * 
         */
        public Builder approvalSettings(AccessPackageAssignmentPolicyApprovalSettingsArgs approvalSettings) {
            return approvalSettings(Output.of(approvalSettings));
        }

        /**
         * @param assignmentReviewSettings An `assignment_review_settings` block, to specify whether assignment review is needed and how it is conducted, as documented below.
         * 
         * @return builder
         * 
         */
        public Builder assignmentReviewSettings(@Nullable Output<AccessPackageAssignmentPolicyAssignmentReviewSettingsArgs> assignmentReviewSettings) {
            $.assignmentReviewSettings = assignmentReviewSettings;
            return this;
        }

        /**
         * @param assignmentReviewSettings An `assignment_review_settings` block, to specify whether assignment review is needed and how it is conducted, as documented below.
         * 
         * @return builder
         * 
         */
        public Builder assignmentReviewSettings(AccessPackageAssignmentPolicyAssignmentReviewSettingsArgs assignmentReviewSettings) {
            return assignmentReviewSettings(Output.of(assignmentReviewSettings));
        }

        /**
         * @param description The description of the policy.
         * 
         * @return builder
         * 
         */
        public Builder description(Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the policy.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param displayName The display name of the policy.
         * 
         * @return builder
         * 
         */
        public Builder displayName(Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName The display name of the policy.
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param durationInDays How many days this assignment is valid for.
         * 
         * @return builder
         * 
         */
        public Builder durationInDays(@Nullable Output<Integer> durationInDays) {
            $.durationInDays = durationInDays;
            return this;
        }

        /**
         * @param durationInDays How many days this assignment is valid for.
         * 
         * @return builder
         * 
         */
        public Builder durationInDays(Integer durationInDays) {
            return durationInDays(Output.of(durationInDays));
        }

        /**
         * @param expirationDate The date that this assignment expires, formatted as an RFC3339 date string in UTC(e.g. 2018-01-01T01:02:03Z).
         * 
         * @return builder
         * 
         */
        public Builder expirationDate(@Nullable Output<String> expirationDate) {
            $.expirationDate = expirationDate;
            return this;
        }

        /**
         * @param expirationDate The date that this assignment expires, formatted as an RFC3339 date string in UTC(e.g. 2018-01-01T01:02:03Z).
         * 
         * @return builder
         * 
         */
        public Builder expirationDate(String expirationDate) {
            return expirationDate(Output.of(expirationDate));
        }

        /**
         * @param extensionEnabled Whether users will be able to request extension of their access to this package before their access expires.
         * 
         * @return builder
         * 
         */
        public Builder extensionEnabled(@Nullable Output<Boolean> extensionEnabled) {
            $.extensionEnabled = extensionEnabled;
            return this;
        }

        /**
         * @param extensionEnabled Whether users will be able to request extension of their access to this package before their access expires.
         * 
         * @return builder
         * 
         */
        public Builder extensionEnabled(Boolean extensionEnabled) {
            return extensionEnabled(Output.of(extensionEnabled));
        }

        /**
         * @param questions One or more `question` blocks for the requestor, as documented below.
         * 
         * @return builder
         * 
         */
        public Builder questions(@Nullable Output<List<AccessPackageAssignmentPolicyQuestionArgs>> questions) {
            $.questions = questions;
            return this;
        }

        /**
         * @param questions One or more `question` blocks for the requestor, as documented below.
         * 
         * @return builder
         * 
         */
        public Builder questions(List<AccessPackageAssignmentPolicyQuestionArgs> questions) {
            return questions(Output.of(questions));
        }

        /**
         * @param questions One or more `question` blocks for the requestor, as documented below.
         * 
         * @return builder
         * 
         */
        public Builder questions(AccessPackageAssignmentPolicyQuestionArgs... questions) {
            return questions(List.of(questions));
        }

        /**
         * @param requestorSettings A `requestor_settings` block to configure the users who can request access, as documented below.
         * 
         * @return builder
         * 
         */
        public Builder requestorSettings(@Nullable Output<AccessPackageAssignmentPolicyRequestorSettingsArgs> requestorSettings) {
            $.requestorSettings = requestorSettings;
            return this;
        }

        /**
         * @param requestorSettings A `requestor_settings` block to configure the users who can request access, as documented below.
         * 
         * @return builder
         * 
         */
        public Builder requestorSettings(AccessPackageAssignmentPolicyRequestorSettingsArgs requestorSettings) {
            return requestorSettings(Output.of(requestorSettings));
        }

        public AccessPackageAssignmentPolicyArgs build() {
            $.accessPackageId = Objects.requireNonNull($.accessPackageId, "expected parameter 'accessPackageId' to be non-null");
            $.description = Objects.requireNonNull($.description, "expected parameter 'description' to be non-null");
            $.displayName = Objects.requireNonNull($.displayName, "expected parameter 'displayName' to be non-null");
            return $;
        }
    }

}

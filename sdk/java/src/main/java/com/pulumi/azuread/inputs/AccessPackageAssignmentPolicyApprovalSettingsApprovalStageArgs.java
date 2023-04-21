// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.inputs;

import com.pulumi.azuread.inputs.AccessPackageAssignmentPolicyApprovalSettingsApprovalStageAlternativeApproverArgs;
import com.pulumi.azuread.inputs.AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AccessPackageAssignmentPolicyApprovalSettingsApprovalStageArgs extends com.pulumi.resources.ResourceArgs {

    public static final AccessPackageAssignmentPolicyApprovalSettingsApprovalStageArgs Empty = new AccessPackageAssignmentPolicyApprovalSettingsApprovalStageArgs();

    /**
     * Whether alternative approvers are enabled.
     * 
     */
    @Import(name="alternativeApprovalEnabled")
    private @Nullable Output<Boolean> alternativeApprovalEnabled;

    /**
     * @return Whether alternative approvers are enabled.
     * 
     */
    public Optional<Output<Boolean>> alternativeApprovalEnabled() {
        return Optional.ofNullable(this.alternativeApprovalEnabled);
    }

    /**
     * A block specifying alternative approvers when escalation is enabled and the primary approvers do not respond before the escalation time, as documented below.
     * 
     */
    @Import(name="alternativeApprovers")
    private @Nullable Output<List<AccessPackageAssignmentPolicyApprovalSettingsApprovalStageAlternativeApproverArgs>> alternativeApprovers;

    /**
     * @return A block specifying alternative approvers when escalation is enabled and the primary approvers do not respond before the escalation time, as documented below.
     * 
     */
    public Optional<Output<List<AccessPackageAssignmentPolicyApprovalSettingsApprovalStageAlternativeApproverArgs>>> alternativeApprovers() {
        return Optional.ofNullable(this.alternativeApprovers);
    }

    /**
     * Maximum number of days within which a request must be approved. If a request is not approved within this time period after it is made, it will be automatically rejected.
     * 
     */
    @Import(name="approvalTimeoutInDays", required=true)
    private Output<Integer> approvalTimeoutInDays;

    /**
     * @return Maximum number of days within which a request must be approved. If a request is not approved within this time period after it is made, it will be automatically rejected.
     * 
     */
    public Output<Integer> approvalTimeoutInDays() {
        return this.approvalTimeoutInDays;
    }

    /**
     * Whether an approver must provide a justification for their decision. Justification is visible to other approvers and the requestor.
     * 
     */
    @Import(name="approverJustificationRequired")
    private @Nullable Output<Boolean> approverJustificationRequired;

    /**
     * @return Whether an approver must provide a justification for their decision. Justification is visible to other approvers and the requestor.
     * 
     */
    public Optional<Output<Boolean>> approverJustificationRequired() {
        return Optional.ofNullable(this.approverJustificationRequired);
    }

    /**
     * Number of days before the request is forwarded to alternative approvers.
     * 
     */
    @Import(name="enableAlternativeApprovalInDays")
    private @Nullable Output<Integer> enableAlternativeApprovalInDays;

    /**
     * @return Number of days before the request is forwarded to alternative approvers.
     * 
     */
    public Optional<Output<Integer>> enableAlternativeApprovalInDays() {
        return Optional.ofNullable(this.enableAlternativeApprovalInDays);
    }

    /**
     * A block specifying the users who will be asked to approve requests, as documented below.
     * 
     */
    @Import(name="primaryApprovers")
    private @Nullable Output<List<AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverArgs>> primaryApprovers;

    /**
     * @return A block specifying the users who will be asked to approve requests, as documented below.
     * 
     */
    public Optional<Output<List<AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverArgs>>> primaryApprovers() {
        return Optional.ofNullable(this.primaryApprovers);
    }

    private AccessPackageAssignmentPolicyApprovalSettingsApprovalStageArgs() {}

    private AccessPackageAssignmentPolicyApprovalSettingsApprovalStageArgs(AccessPackageAssignmentPolicyApprovalSettingsApprovalStageArgs $) {
        this.alternativeApprovalEnabled = $.alternativeApprovalEnabled;
        this.alternativeApprovers = $.alternativeApprovers;
        this.approvalTimeoutInDays = $.approvalTimeoutInDays;
        this.approverJustificationRequired = $.approverJustificationRequired;
        this.enableAlternativeApprovalInDays = $.enableAlternativeApprovalInDays;
        this.primaryApprovers = $.primaryApprovers;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AccessPackageAssignmentPolicyApprovalSettingsApprovalStageArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AccessPackageAssignmentPolicyApprovalSettingsApprovalStageArgs $;

        public Builder() {
            $ = new AccessPackageAssignmentPolicyApprovalSettingsApprovalStageArgs();
        }

        public Builder(AccessPackageAssignmentPolicyApprovalSettingsApprovalStageArgs defaults) {
            $ = new AccessPackageAssignmentPolicyApprovalSettingsApprovalStageArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param alternativeApprovalEnabled Whether alternative approvers are enabled.
         * 
         * @return builder
         * 
         */
        public Builder alternativeApprovalEnabled(@Nullable Output<Boolean> alternativeApprovalEnabled) {
            $.alternativeApprovalEnabled = alternativeApprovalEnabled;
            return this;
        }

        /**
         * @param alternativeApprovalEnabled Whether alternative approvers are enabled.
         * 
         * @return builder
         * 
         */
        public Builder alternativeApprovalEnabled(Boolean alternativeApprovalEnabled) {
            return alternativeApprovalEnabled(Output.of(alternativeApprovalEnabled));
        }

        /**
         * @param alternativeApprovers A block specifying alternative approvers when escalation is enabled and the primary approvers do not respond before the escalation time, as documented below.
         * 
         * @return builder
         * 
         */
        public Builder alternativeApprovers(@Nullable Output<List<AccessPackageAssignmentPolicyApprovalSettingsApprovalStageAlternativeApproverArgs>> alternativeApprovers) {
            $.alternativeApprovers = alternativeApprovers;
            return this;
        }

        /**
         * @param alternativeApprovers A block specifying alternative approvers when escalation is enabled and the primary approvers do not respond before the escalation time, as documented below.
         * 
         * @return builder
         * 
         */
        public Builder alternativeApprovers(List<AccessPackageAssignmentPolicyApprovalSettingsApprovalStageAlternativeApproverArgs> alternativeApprovers) {
            return alternativeApprovers(Output.of(alternativeApprovers));
        }

        /**
         * @param alternativeApprovers A block specifying alternative approvers when escalation is enabled and the primary approvers do not respond before the escalation time, as documented below.
         * 
         * @return builder
         * 
         */
        public Builder alternativeApprovers(AccessPackageAssignmentPolicyApprovalSettingsApprovalStageAlternativeApproverArgs... alternativeApprovers) {
            return alternativeApprovers(List.of(alternativeApprovers));
        }

        /**
         * @param approvalTimeoutInDays Maximum number of days within which a request must be approved. If a request is not approved within this time period after it is made, it will be automatically rejected.
         * 
         * @return builder
         * 
         */
        public Builder approvalTimeoutInDays(Output<Integer> approvalTimeoutInDays) {
            $.approvalTimeoutInDays = approvalTimeoutInDays;
            return this;
        }

        /**
         * @param approvalTimeoutInDays Maximum number of days within which a request must be approved. If a request is not approved within this time period after it is made, it will be automatically rejected.
         * 
         * @return builder
         * 
         */
        public Builder approvalTimeoutInDays(Integer approvalTimeoutInDays) {
            return approvalTimeoutInDays(Output.of(approvalTimeoutInDays));
        }

        /**
         * @param approverJustificationRequired Whether an approver must provide a justification for their decision. Justification is visible to other approvers and the requestor.
         * 
         * @return builder
         * 
         */
        public Builder approverJustificationRequired(@Nullable Output<Boolean> approverJustificationRequired) {
            $.approverJustificationRequired = approverJustificationRequired;
            return this;
        }

        /**
         * @param approverJustificationRequired Whether an approver must provide a justification for their decision. Justification is visible to other approvers and the requestor.
         * 
         * @return builder
         * 
         */
        public Builder approverJustificationRequired(Boolean approverJustificationRequired) {
            return approverJustificationRequired(Output.of(approverJustificationRequired));
        }

        /**
         * @param enableAlternativeApprovalInDays Number of days before the request is forwarded to alternative approvers.
         * 
         * @return builder
         * 
         */
        public Builder enableAlternativeApprovalInDays(@Nullable Output<Integer> enableAlternativeApprovalInDays) {
            $.enableAlternativeApprovalInDays = enableAlternativeApprovalInDays;
            return this;
        }

        /**
         * @param enableAlternativeApprovalInDays Number of days before the request is forwarded to alternative approvers.
         * 
         * @return builder
         * 
         */
        public Builder enableAlternativeApprovalInDays(Integer enableAlternativeApprovalInDays) {
            return enableAlternativeApprovalInDays(Output.of(enableAlternativeApprovalInDays));
        }

        /**
         * @param primaryApprovers A block specifying the users who will be asked to approve requests, as documented below.
         * 
         * @return builder
         * 
         */
        public Builder primaryApprovers(@Nullable Output<List<AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverArgs>> primaryApprovers) {
            $.primaryApprovers = primaryApprovers;
            return this;
        }

        /**
         * @param primaryApprovers A block specifying the users who will be asked to approve requests, as documented below.
         * 
         * @return builder
         * 
         */
        public Builder primaryApprovers(List<AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverArgs> primaryApprovers) {
            return primaryApprovers(Output.of(primaryApprovers));
        }

        /**
         * @param primaryApprovers A block specifying the users who will be asked to approve requests, as documented below.
         * 
         * @return builder
         * 
         */
        public Builder primaryApprovers(AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverArgs... primaryApprovers) {
            return primaryApprovers(List.of(primaryApprovers));
        }

        public AccessPackageAssignmentPolicyApprovalSettingsApprovalStageArgs build() {
            $.approvalTimeoutInDays = Objects.requireNonNull($.approvalTimeoutInDays, "expected parameter 'approvalTimeoutInDays' to be non-null");
            return $;
        }
    }

}

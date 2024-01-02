// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.azuread.outputs.AccessPackageAssignmentPolicyApprovalSettingsApprovalStageAlternativeApprover;
import com.pulumi.azuread.outputs.AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApprover;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class AccessPackageAssignmentPolicyApprovalSettingsApprovalStage {
    /**
     * @return Whether alternative approvers are enabled.
     * 
     */
    private @Nullable Boolean alternativeApprovalEnabled;
    /**
     * @return A block specifying alternative approvers when escalation is enabled and the primary approvers do not respond before the escalation time, as documented below.
     * 
     */
    private @Nullable List<AccessPackageAssignmentPolicyApprovalSettingsApprovalStageAlternativeApprover> alternativeApprovers;
    /**
     * @return Maximum number of days within which a request must be approved. If a request is not approved within this time period after it is made, it will be automatically rejected.
     * 
     */
    private Integer approvalTimeoutInDays;
    /**
     * @return Whether an approver must provide a justification for their decision. Justification is visible to other approvers and the requestor.
     * 
     */
    private @Nullable Boolean approverJustificationRequired;
    /**
     * @return Number of days before the request is forwarded to alternative approvers.
     * 
     */
    private @Nullable Integer enableAlternativeApprovalInDays;
    /**
     * @return A block specifying the users who will be asked to approve requests, as documented below.
     * 
     */
    private @Nullable List<AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApprover> primaryApprovers;

    private AccessPackageAssignmentPolicyApprovalSettingsApprovalStage() {}
    /**
     * @return Whether alternative approvers are enabled.
     * 
     */
    public Optional<Boolean> alternativeApprovalEnabled() {
        return Optional.ofNullable(this.alternativeApprovalEnabled);
    }
    /**
     * @return A block specifying alternative approvers when escalation is enabled and the primary approvers do not respond before the escalation time, as documented below.
     * 
     */
    public List<AccessPackageAssignmentPolicyApprovalSettingsApprovalStageAlternativeApprover> alternativeApprovers() {
        return this.alternativeApprovers == null ? List.of() : this.alternativeApprovers;
    }
    /**
     * @return Maximum number of days within which a request must be approved. If a request is not approved within this time period after it is made, it will be automatically rejected.
     * 
     */
    public Integer approvalTimeoutInDays() {
        return this.approvalTimeoutInDays;
    }
    /**
     * @return Whether an approver must provide a justification for their decision. Justification is visible to other approvers and the requestor.
     * 
     */
    public Optional<Boolean> approverJustificationRequired() {
        return Optional.ofNullable(this.approverJustificationRequired);
    }
    /**
     * @return Number of days before the request is forwarded to alternative approvers.
     * 
     */
    public Optional<Integer> enableAlternativeApprovalInDays() {
        return Optional.ofNullable(this.enableAlternativeApprovalInDays);
    }
    /**
     * @return A block specifying the users who will be asked to approve requests, as documented below.
     * 
     */
    public List<AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApprover> primaryApprovers() {
        return this.primaryApprovers == null ? List.of() : this.primaryApprovers;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AccessPackageAssignmentPolicyApprovalSettingsApprovalStage defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean alternativeApprovalEnabled;
        private @Nullable List<AccessPackageAssignmentPolicyApprovalSettingsApprovalStageAlternativeApprover> alternativeApprovers;
        private Integer approvalTimeoutInDays;
        private @Nullable Boolean approverJustificationRequired;
        private @Nullable Integer enableAlternativeApprovalInDays;
        private @Nullable List<AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApprover> primaryApprovers;
        public Builder() {}
        public Builder(AccessPackageAssignmentPolicyApprovalSettingsApprovalStage defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.alternativeApprovalEnabled = defaults.alternativeApprovalEnabled;
    	      this.alternativeApprovers = defaults.alternativeApprovers;
    	      this.approvalTimeoutInDays = defaults.approvalTimeoutInDays;
    	      this.approverJustificationRequired = defaults.approverJustificationRequired;
    	      this.enableAlternativeApprovalInDays = defaults.enableAlternativeApprovalInDays;
    	      this.primaryApprovers = defaults.primaryApprovers;
        }

        @CustomType.Setter
        public Builder alternativeApprovalEnabled(@Nullable Boolean alternativeApprovalEnabled) {

            this.alternativeApprovalEnabled = alternativeApprovalEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder alternativeApprovers(@Nullable List<AccessPackageAssignmentPolicyApprovalSettingsApprovalStageAlternativeApprover> alternativeApprovers) {

            this.alternativeApprovers = alternativeApprovers;
            return this;
        }
        public Builder alternativeApprovers(AccessPackageAssignmentPolicyApprovalSettingsApprovalStageAlternativeApprover... alternativeApprovers) {
            return alternativeApprovers(List.of(alternativeApprovers));
        }
        @CustomType.Setter
        public Builder approvalTimeoutInDays(Integer approvalTimeoutInDays) {
            if (approvalTimeoutInDays == null) {
              throw new MissingRequiredPropertyException("AccessPackageAssignmentPolicyApprovalSettingsApprovalStage", "approvalTimeoutInDays");
            }
            this.approvalTimeoutInDays = approvalTimeoutInDays;
            return this;
        }
        @CustomType.Setter
        public Builder approverJustificationRequired(@Nullable Boolean approverJustificationRequired) {

            this.approverJustificationRequired = approverJustificationRequired;
            return this;
        }
        @CustomType.Setter
        public Builder enableAlternativeApprovalInDays(@Nullable Integer enableAlternativeApprovalInDays) {

            this.enableAlternativeApprovalInDays = enableAlternativeApprovalInDays;
            return this;
        }
        @CustomType.Setter
        public Builder primaryApprovers(@Nullable List<AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApprover> primaryApprovers) {

            this.primaryApprovers = primaryApprovers;
            return this;
        }
        public Builder primaryApprovers(AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApprover... primaryApprovers) {
            return primaryApprovers(List.of(primaryApprovers));
        }
        public AccessPackageAssignmentPolicyApprovalSettingsApprovalStage build() {
            final var _resultValue = new AccessPackageAssignmentPolicyApprovalSettingsApprovalStage();
            _resultValue.alternativeApprovalEnabled = alternativeApprovalEnabled;
            _resultValue.alternativeApprovers = alternativeApprovers;
            _resultValue.approvalTimeoutInDays = approvalTimeoutInDays;
            _resultValue.approverJustificationRequired = approverJustificationRequired;
            _resultValue.enableAlternativeApprovalInDays = enableAlternativeApprovalInDays;
            _resultValue.primaryApprovers = primaryApprovers;
            return _resultValue;
        }
    }
}

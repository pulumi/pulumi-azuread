// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.azuread.outputs.GroupRoleManagementPolicyNotificationRulesActiveAssignments;
import com.pulumi.azuread.outputs.GroupRoleManagementPolicyNotificationRulesEligibleActivations;
import com.pulumi.azuread.outputs.GroupRoleManagementPolicyNotificationRulesEligibleAssignments;
import com.pulumi.core.annotations.CustomType;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GroupRoleManagementPolicyNotificationRules {
    /**
     * @return A `notification_target` block as defined below to configure notfications on active role assignments.
     * 
     */
    private @Nullable GroupRoleManagementPolicyNotificationRulesActiveAssignments activeAssignments;
    /**
     * @return A `notification_target` block as defined below for configuring notifications on activation of eligible role.
     * 
     */
    private @Nullable GroupRoleManagementPolicyNotificationRulesEligibleActivations eligibleActivations;
    /**
     * @return A `notification_target` block as defined below to configure notification on eligible role assignments.
     * 
     * At least one `notification_target` block must be provided.
     * 
     */
    private @Nullable GroupRoleManagementPolicyNotificationRulesEligibleAssignments eligibleAssignments;

    private GroupRoleManagementPolicyNotificationRules() {}
    /**
     * @return A `notification_target` block as defined below to configure notfications on active role assignments.
     * 
     */
    public Optional<GroupRoleManagementPolicyNotificationRulesActiveAssignments> activeAssignments() {
        return Optional.ofNullable(this.activeAssignments);
    }
    /**
     * @return A `notification_target` block as defined below for configuring notifications on activation of eligible role.
     * 
     */
    public Optional<GroupRoleManagementPolicyNotificationRulesEligibleActivations> eligibleActivations() {
        return Optional.ofNullable(this.eligibleActivations);
    }
    /**
     * @return A `notification_target` block as defined below to configure notification on eligible role assignments.
     * 
     * At least one `notification_target` block must be provided.
     * 
     */
    public Optional<GroupRoleManagementPolicyNotificationRulesEligibleAssignments> eligibleAssignments() {
        return Optional.ofNullable(this.eligibleAssignments);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GroupRoleManagementPolicyNotificationRules defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable GroupRoleManagementPolicyNotificationRulesActiveAssignments activeAssignments;
        private @Nullable GroupRoleManagementPolicyNotificationRulesEligibleActivations eligibleActivations;
        private @Nullable GroupRoleManagementPolicyNotificationRulesEligibleAssignments eligibleAssignments;
        public Builder() {}
        public Builder(GroupRoleManagementPolicyNotificationRules defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.activeAssignments = defaults.activeAssignments;
    	      this.eligibleActivations = defaults.eligibleActivations;
    	      this.eligibleAssignments = defaults.eligibleAssignments;
        }

        @CustomType.Setter
        public Builder activeAssignments(@Nullable GroupRoleManagementPolicyNotificationRulesActiveAssignments activeAssignments) {

            this.activeAssignments = activeAssignments;
            return this;
        }
        @CustomType.Setter
        public Builder eligibleActivations(@Nullable GroupRoleManagementPolicyNotificationRulesEligibleActivations eligibleActivations) {

            this.eligibleActivations = eligibleActivations;
            return this;
        }
        @CustomType.Setter
        public Builder eligibleAssignments(@Nullable GroupRoleManagementPolicyNotificationRulesEligibleAssignments eligibleAssignments) {

            this.eligibleAssignments = eligibleAssignments;
            return this;
        }
        public GroupRoleManagementPolicyNotificationRules build() {
            final var _resultValue = new GroupRoleManagementPolicyNotificationRules();
            _resultValue.activeAssignments = activeAssignments;
            _resultValue.eligibleActivations = eligibleActivations;
            _resultValue.eligibleAssignments = eligibleAssignments;
            return _resultValue;
        }
    }
}
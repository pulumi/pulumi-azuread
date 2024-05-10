// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.inputs;

import com.pulumi.azuread.inputs.GroupRoleManagementPolicyNotificationRulesEligibleActivationsAdminNotificationsArgs;
import com.pulumi.azuread.inputs.GroupRoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsArgs;
import com.pulumi.azuread.inputs.GroupRoleManagementPolicyNotificationRulesEligibleActivationsAssigneeNotificationsArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GroupRoleManagementPolicyNotificationRulesEligibleActivationsArgs extends com.pulumi.resources.ResourceArgs {

    public static final GroupRoleManagementPolicyNotificationRulesEligibleActivationsArgs Empty = new GroupRoleManagementPolicyNotificationRulesEligibleActivationsArgs();

    /**
     * Admin notification settings
     * 
     */
    @Import(name="adminNotifications")
    private @Nullable Output<GroupRoleManagementPolicyNotificationRulesEligibleActivationsAdminNotificationsArgs> adminNotifications;

    /**
     * @return Admin notification settings
     * 
     */
    public Optional<Output<GroupRoleManagementPolicyNotificationRulesEligibleActivationsAdminNotificationsArgs>> adminNotifications() {
        return Optional.ofNullable(this.adminNotifications);
    }

    /**
     * Approver notification settings
     * 
     */
    @Import(name="approverNotifications")
    private @Nullable Output<GroupRoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsArgs> approverNotifications;

    /**
     * @return Approver notification settings
     * 
     */
    public Optional<Output<GroupRoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsArgs>> approverNotifications() {
        return Optional.ofNullable(this.approverNotifications);
    }

    /**
     * Assignee notification settings
     * 
     */
    @Import(name="assigneeNotifications")
    private @Nullable Output<GroupRoleManagementPolicyNotificationRulesEligibleActivationsAssigneeNotificationsArgs> assigneeNotifications;

    /**
     * @return Assignee notification settings
     * 
     */
    public Optional<Output<GroupRoleManagementPolicyNotificationRulesEligibleActivationsAssigneeNotificationsArgs>> assigneeNotifications() {
        return Optional.ofNullable(this.assigneeNotifications);
    }

    private GroupRoleManagementPolicyNotificationRulesEligibleActivationsArgs() {}

    private GroupRoleManagementPolicyNotificationRulesEligibleActivationsArgs(GroupRoleManagementPolicyNotificationRulesEligibleActivationsArgs $) {
        this.adminNotifications = $.adminNotifications;
        this.approverNotifications = $.approverNotifications;
        this.assigneeNotifications = $.assigneeNotifications;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GroupRoleManagementPolicyNotificationRulesEligibleActivationsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GroupRoleManagementPolicyNotificationRulesEligibleActivationsArgs $;

        public Builder() {
            $ = new GroupRoleManagementPolicyNotificationRulesEligibleActivationsArgs();
        }

        public Builder(GroupRoleManagementPolicyNotificationRulesEligibleActivationsArgs defaults) {
            $ = new GroupRoleManagementPolicyNotificationRulesEligibleActivationsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param adminNotifications Admin notification settings
         * 
         * @return builder
         * 
         */
        public Builder adminNotifications(@Nullable Output<GroupRoleManagementPolicyNotificationRulesEligibleActivationsAdminNotificationsArgs> adminNotifications) {
            $.adminNotifications = adminNotifications;
            return this;
        }

        /**
         * @param adminNotifications Admin notification settings
         * 
         * @return builder
         * 
         */
        public Builder adminNotifications(GroupRoleManagementPolicyNotificationRulesEligibleActivationsAdminNotificationsArgs adminNotifications) {
            return adminNotifications(Output.of(adminNotifications));
        }

        /**
         * @param approverNotifications Approver notification settings
         * 
         * @return builder
         * 
         */
        public Builder approverNotifications(@Nullable Output<GroupRoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsArgs> approverNotifications) {
            $.approverNotifications = approverNotifications;
            return this;
        }

        /**
         * @param approverNotifications Approver notification settings
         * 
         * @return builder
         * 
         */
        public Builder approverNotifications(GroupRoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsArgs approverNotifications) {
            return approverNotifications(Output.of(approverNotifications));
        }

        /**
         * @param assigneeNotifications Assignee notification settings
         * 
         * @return builder
         * 
         */
        public Builder assigneeNotifications(@Nullable Output<GroupRoleManagementPolicyNotificationRulesEligibleActivationsAssigneeNotificationsArgs> assigneeNotifications) {
            $.assigneeNotifications = assigneeNotifications;
            return this;
        }

        /**
         * @param assigneeNotifications Assignee notification settings
         * 
         * @return builder
         * 
         */
        public Builder assigneeNotifications(GroupRoleManagementPolicyNotificationRulesEligibleActivationsAssigneeNotificationsArgs assigneeNotifications) {
            return assigneeNotifications(Output.of(assigneeNotifications));
        }

        public GroupRoleManagementPolicyNotificationRulesEligibleActivationsArgs build() {
            return $;
        }
    }

}

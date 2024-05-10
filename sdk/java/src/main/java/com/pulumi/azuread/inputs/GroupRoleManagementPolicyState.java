// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.inputs;

import com.pulumi.azuread.inputs.GroupRoleManagementPolicyActivationRulesArgs;
import com.pulumi.azuread.inputs.GroupRoleManagementPolicyActiveAssignmentRulesArgs;
import com.pulumi.azuread.inputs.GroupRoleManagementPolicyEligibleAssignmentRulesArgs;
import com.pulumi.azuread.inputs.GroupRoleManagementPolicyNotificationRulesArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GroupRoleManagementPolicyState extends com.pulumi.resources.ResourceArgs {

    public static final GroupRoleManagementPolicyState Empty = new GroupRoleManagementPolicyState();

    /**
     * An `activation_rules` block as defined below.
     * 
     */
    @Import(name="activationRules")
    private @Nullable Output<GroupRoleManagementPolicyActivationRulesArgs> activationRules;

    /**
     * @return An `activation_rules` block as defined below.
     * 
     */
    public Optional<Output<GroupRoleManagementPolicyActivationRulesArgs>> activationRules() {
        return Optional.ofNullable(this.activationRules);
    }

    /**
     * An `active_assignment_rules` block as defined below.
     * 
     */
    @Import(name="activeAssignmentRules")
    private @Nullable Output<GroupRoleManagementPolicyActiveAssignmentRulesArgs> activeAssignmentRules;

    /**
     * @return An `active_assignment_rules` block as defined below.
     * 
     */
    public Optional<Output<GroupRoleManagementPolicyActiveAssignmentRulesArgs>> activeAssignmentRules() {
        return Optional.ofNullable(this.activeAssignmentRules);
    }

    /**
     * (String) The description of this policy.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return (String) The description of this policy.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * (String) The display name of this policy.
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return (String) The display name of this policy.
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * An `eligible_assignment_rules` block as defined below.
     * 
     */
    @Import(name="eligibleAssignmentRules")
    private @Nullable Output<GroupRoleManagementPolicyEligibleAssignmentRulesArgs> eligibleAssignmentRules;

    /**
     * @return An `eligible_assignment_rules` block as defined below.
     * 
     */
    public Optional<Output<GroupRoleManagementPolicyEligibleAssignmentRulesArgs>> eligibleAssignmentRules() {
        return Optional.ofNullable(this.eligibleAssignmentRules);
    }

    /**
     * The ID of the Azure AD group for which the policy applies.
     * 
     */
    @Import(name="groupId")
    private @Nullable Output<String> groupId;

    /**
     * @return The ID of the Azure AD group for which the policy applies.
     * 
     */
    public Optional<Output<String>> groupId() {
        return Optional.ofNullable(this.groupId);
    }

    /**
     * A `notification_rules` block as defined below.
     * 
     */
    @Import(name="notificationRules")
    private @Nullable Output<GroupRoleManagementPolicyNotificationRulesArgs> notificationRules;

    /**
     * @return A `notification_rules` block as defined below.
     * 
     */
    public Optional<Output<GroupRoleManagementPolicyNotificationRulesArgs>> notificationRules() {
        return Optional.ofNullable(this.notificationRules);
    }

    /**
     * The type of assignment this policy coveres. Can be either `member` or `owner`.
     * 
     */
    @Import(name="roleId")
    private @Nullable Output<String> roleId;

    /**
     * @return The type of assignment this policy coveres. Can be either `member` or `owner`.
     * 
     */
    public Optional<Output<String>> roleId() {
        return Optional.ofNullable(this.roleId);
    }

    private GroupRoleManagementPolicyState() {}

    private GroupRoleManagementPolicyState(GroupRoleManagementPolicyState $) {
        this.activationRules = $.activationRules;
        this.activeAssignmentRules = $.activeAssignmentRules;
        this.description = $.description;
        this.displayName = $.displayName;
        this.eligibleAssignmentRules = $.eligibleAssignmentRules;
        this.groupId = $.groupId;
        this.notificationRules = $.notificationRules;
        this.roleId = $.roleId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GroupRoleManagementPolicyState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GroupRoleManagementPolicyState $;

        public Builder() {
            $ = new GroupRoleManagementPolicyState();
        }

        public Builder(GroupRoleManagementPolicyState defaults) {
            $ = new GroupRoleManagementPolicyState(Objects.requireNonNull(defaults));
        }

        /**
         * @param activationRules An `activation_rules` block as defined below.
         * 
         * @return builder
         * 
         */
        public Builder activationRules(@Nullable Output<GroupRoleManagementPolicyActivationRulesArgs> activationRules) {
            $.activationRules = activationRules;
            return this;
        }

        /**
         * @param activationRules An `activation_rules` block as defined below.
         * 
         * @return builder
         * 
         */
        public Builder activationRules(GroupRoleManagementPolicyActivationRulesArgs activationRules) {
            return activationRules(Output.of(activationRules));
        }

        /**
         * @param activeAssignmentRules An `active_assignment_rules` block as defined below.
         * 
         * @return builder
         * 
         */
        public Builder activeAssignmentRules(@Nullable Output<GroupRoleManagementPolicyActiveAssignmentRulesArgs> activeAssignmentRules) {
            $.activeAssignmentRules = activeAssignmentRules;
            return this;
        }

        /**
         * @param activeAssignmentRules An `active_assignment_rules` block as defined below.
         * 
         * @return builder
         * 
         */
        public Builder activeAssignmentRules(GroupRoleManagementPolicyActiveAssignmentRulesArgs activeAssignmentRules) {
            return activeAssignmentRules(Output.of(activeAssignmentRules));
        }

        /**
         * @param description (String) The description of this policy.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description (String) The description of this policy.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param displayName (String) The display name of this policy.
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName (String) The display name of this policy.
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param eligibleAssignmentRules An `eligible_assignment_rules` block as defined below.
         * 
         * @return builder
         * 
         */
        public Builder eligibleAssignmentRules(@Nullable Output<GroupRoleManagementPolicyEligibleAssignmentRulesArgs> eligibleAssignmentRules) {
            $.eligibleAssignmentRules = eligibleAssignmentRules;
            return this;
        }

        /**
         * @param eligibleAssignmentRules An `eligible_assignment_rules` block as defined below.
         * 
         * @return builder
         * 
         */
        public Builder eligibleAssignmentRules(GroupRoleManagementPolicyEligibleAssignmentRulesArgs eligibleAssignmentRules) {
            return eligibleAssignmentRules(Output.of(eligibleAssignmentRules));
        }

        /**
         * @param groupId The ID of the Azure AD group for which the policy applies.
         * 
         * @return builder
         * 
         */
        public Builder groupId(@Nullable Output<String> groupId) {
            $.groupId = groupId;
            return this;
        }

        /**
         * @param groupId The ID of the Azure AD group for which the policy applies.
         * 
         * @return builder
         * 
         */
        public Builder groupId(String groupId) {
            return groupId(Output.of(groupId));
        }

        /**
         * @param notificationRules A `notification_rules` block as defined below.
         * 
         * @return builder
         * 
         */
        public Builder notificationRules(@Nullable Output<GroupRoleManagementPolicyNotificationRulesArgs> notificationRules) {
            $.notificationRules = notificationRules;
            return this;
        }

        /**
         * @param notificationRules A `notification_rules` block as defined below.
         * 
         * @return builder
         * 
         */
        public Builder notificationRules(GroupRoleManagementPolicyNotificationRulesArgs notificationRules) {
            return notificationRules(Output.of(notificationRules));
        }

        /**
         * @param roleId The type of assignment this policy coveres. Can be either `member` or `owner`.
         * 
         * @return builder
         * 
         */
        public Builder roleId(@Nullable Output<String> roleId) {
            $.roleId = roleId;
            return this;
        }

        /**
         * @param roleId The type of assignment this policy coveres. Can be either `member` or `owner`.
         * 
         * @return builder
         * 
         */
        public Builder roleId(String roleId) {
            return roleId(Output.of(roleId));
        }

        public GroupRoleManagementPolicyState build() {
            return $;
        }
    }

}

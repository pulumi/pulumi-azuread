// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GroupRoleManagementPolicyNotificationRulesEligibleActivationsAdminNotificationsArgs extends com.pulumi.resources.ResourceArgs {

    public static final GroupRoleManagementPolicyNotificationRulesEligibleActivationsAdminNotificationsArgs Empty = new GroupRoleManagementPolicyNotificationRulesEligibleActivationsAdminNotificationsArgs();

    /**
     * The additional recipients to notify
     * 
     */
    @Import(name="additionalRecipients")
    private @Nullable Output<List<String>> additionalRecipients;

    /**
     * @return The additional recipients to notify
     * 
     */
    public Optional<Output<List<String>>> additionalRecipients() {
        return Optional.ofNullable(this.additionalRecipients);
    }

    /**
     * Whether the default recipients are notified
     * 
     */
    @Import(name="defaultRecipients", required=true)
    private Output<Boolean> defaultRecipients;

    /**
     * @return Whether the default recipients are notified
     * 
     */
    public Output<Boolean> defaultRecipients() {
        return this.defaultRecipients;
    }

    /**
     * What level of notifications are sent
     * 
     */
    @Import(name="notificationLevel", required=true)
    private Output<String> notificationLevel;

    /**
     * @return What level of notifications are sent
     * 
     */
    public Output<String> notificationLevel() {
        return this.notificationLevel;
    }

    private GroupRoleManagementPolicyNotificationRulesEligibleActivationsAdminNotificationsArgs() {}

    private GroupRoleManagementPolicyNotificationRulesEligibleActivationsAdminNotificationsArgs(GroupRoleManagementPolicyNotificationRulesEligibleActivationsAdminNotificationsArgs $) {
        this.additionalRecipients = $.additionalRecipients;
        this.defaultRecipients = $.defaultRecipients;
        this.notificationLevel = $.notificationLevel;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GroupRoleManagementPolicyNotificationRulesEligibleActivationsAdminNotificationsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GroupRoleManagementPolicyNotificationRulesEligibleActivationsAdminNotificationsArgs $;

        public Builder() {
            $ = new GroupRoleManagementPolicyNotificationRulesEligibleActivationsAdminNotificationsArgs();
        }

        public Builder(GroupRoleManagementPolicyNotificationRulesEligibleActivationsAdminNotificationsArgs defaults) {
            $ = new GroupRoleManagementPolicyNotificationRulesEligibleActivationsAdminNotificationsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param additionalRecipients The additional recipients to notify
         * 
         * @return builder
         * 
         */
        public Builder additionalRecipients(@Nullable Output<List<String>> additionalRecipients) {
            $.additionalRecipients = additionalRecipients;
            return this;
        }

        /**
         * @param additionalRecipients The additional recipients to notify
         * 
         * @return builder
         * 
         */
        public Builder additionalRecipients(List<String> additionalRecipients) {
            return additionalRecipients(Output.of(additionalRecipients));
        }

        /**
         * @param additionalRecipients The additional recipients to notify
         * 
         * @return builder
         * 
         */
        public Builder additionalRecipients(String... additionalRecipients) {
            return additionalRecipients(List.of(additionalRecipients));
        }

        /**
         * @param defaultRecipients Whether the default recipients are notified
         * 
         * @return builder
         * 
         */
        public Builder defaultRecipients(Output<Boolean> defaultRecipients) {
            $.defaultRecipients = defaultRecipients;
            return this;
        }

        /**
         * @param defaultRecipients Whether the default recipients are notified
         * 
         * @return builder
         * 
         */
        public Builder defaultRecipients(Boolean defaultRecipients) {
            return defaultRecipients(Output.of(defaultRecipients));
        }

        /**
         * @param notificationLevel What level of notifications are sent
         * 
         * @return builder
         * 
         */
        public Builder notificationLevel(Output<String> notificationLevel) {
            $.notificationLevel = notificationLevel;
            return this;
        }

        /**
         * @param notificationLevel What level of notifications are sent
         * 
         * @return builder
         * 
         */
        public Builder notificationLevel(String notificationLevel) {
            return notificationLevel(Output.of(notificationLevel));
        }

        public GroupRoleManagementPolicyNotificationRulesEligibleActivationsAdminNotificationsArgs build() {
            if ($.defaultRecipients == null) {
                throw new MissingRequiredPropertyException("GroupRoleManagementPolicyNotificationRulesEligibleActivationsAdminNotificationsArgs", "defaultRecipients");
            }
            if ($.notificationLevel == null) {
                throw new MissingRequiredPropertyException("GroupRoleManagementPolicyNotificationRulesEligibleActivationsAdminNotificationsArgs", "notificationLevel");
            }
            return $;
        }
    }

}

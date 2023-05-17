// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ConditionalAccessPolicyConditionsUsersArgs extends com.pulumi.resources.ResourceArgs {

    public static final ConditionalAccessPolicyConditionsUsersArgs Empty = new ConditionalAccessPolicyConditionsUsersArgs();

    /**
     * A list of group IDs excluded from scope of policy.
     * 
     */
    @Import(name="excludedGroups")
    private @Nullable Output<List<String>> excludedGroups;

    /**
     * @return A list of group IDs excluded from scope of policy.
     * 
     */
    public Optional<Output<List<String>>> excludedGroups() {
        return Optional.ofNullable(this.excludedGroups);
    }

    /**
     * A list of role IDs excluded from scope of policy.
     * 
     */
    @Import(name="excludedRoles")
    private @Nullable Output<List<String>> excludedRoles;

    /**
     * @return A list of role IDs excluded from scope of policy.
     * 
     */
    public Optional<Output<List<String>>> excludedRoles() {
        return Optional.ofNullable(this.excludedRoles);
    }

    /**
     * A list of user IDs excluded from scope of policy and/or `GuestsOrExternalUsers`.
     * 
     */
    @Import(name="excludedUsers")
    private @Nullable Output<List<String>> excludedUsers;

    /**
     * @return A list of user IDs excluded from scope of policy and/or `GuestsOrExternalUsers`.
     * 
     */
    public Optional<Output<List<String>>> excludedUsers() {
        return Optional.ofNullable(this.excludedUsers);
    }

    /**
     * A list of group IDs in scope of policy unless explicitly excluded.
     * 
     */
    @Import(name="includedGroups")
    private @Nullable Output<List<String>> includedGroups;

    /**
     * @return A list of group IDs in scope of policy unless explicitly excluded.
     * 
     */
    public Optional<Output<List<String>>> includedGroups() {
        return Optional.ofNullable(this.includedGroups);
    }

    /**
     * A list of role IDs in scope of policy unless explicitly excluded.
     * 
     */
    @Import(name="includedRoles")
    private @Nullable Output<List<String>> includedRoles;

    /**
     * @return A list of role IDs in scope of policy unless explicitly excluded.
     * 
     */
    public Optional<Output<List<String>>> includedRoles() {
        return Optional.ofNullable(this.includedRoles);
    }

    /**
     * A list of user IDs in scope of policy unless explicitly excluded, or `None` or `All` or `GuestsOrExternalUsers`.
     * 
     * &gt; At least one of `included_groups`, `included_roles` or `included_users` must be specified.
     * 
     */
    @Import(name="includedUsers")
    private @Nullable Output<List<String>> includedUsers;

    /**
     * @return A list of user IDs in scope of policy unless explicitly excluded, or `None` or `All` or `GuestsOrExternalUsers`.
     * 
     * &gt; At least one of `included_groups`, `included_roles` or `included_users` must be specified.
     * 
     */
    public Optional<Output<List<String>>> includedUsers() {
        return Optional.ofNullable(this.includedUsers);
    }

    private ConditionalAccessPolicyConditionsUsersArgs() {}

    private ConditionalAccessPolicyConditionsUsersArgs(ConditionalAccessPolicyConditionsUsersArgs $) {
        this.excludedGroups = $.excludedGroups;
        this.excludedRoles = $.excludedRoles;
        this.excludedUsers = $.excludedUsers;
        this.includedGroups = $.includedGroups;
        this.includedRoles = $.includedRoles;
        this.includedUsers = $.includedUsers;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ConditionalAccessPolicyConditionsUsersArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ConditionalAccessPolicyConditionsUsersArgs $;

        public Builder() {
            $ = new ConditionalAccessPolicyConditionsUsersArgs();
        }

        public Builder(ConditionalAccessPolicyConditionsUsersArgs defaults) {
            $ = new ConditionalAccessPolicyConditionsUsersArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param excludedGroups A list of group IDs excluded from scope of policy.
         * 
         * @return builder
         * 
         */
        public Builder excludedGroups(@Nullable Output<List<String>> excludedGroups) {
            $.excludedGroups = excludedGroups;
            return this;
        }

        /**
         * @param excludedGroups A list of group IDs excluded from scope of policy.
         * 
         * @return builder
         * 
         */
        public Builder excludedGroups(List<String> excludedGroups) {
            return excludedGroups(Output.of(excludedGroups));
        }

        /**
         * @param excludedGroups A list of group IDs excluded from scope of policy.
         * 
         * @return builder
         * 
         */
        public Builder excludedGroups(String... excludedGroups) {
            return excludedGroups(List.of(excludedGroups));
        }

        /**
         * @param excludedRoles A list of role IDs excluded from scope of policy.
         * 
         * @return builder
         * 
         */
        public Builder excludedRoles(@Nullable Output<List<String>> excludedRoles) {
            $.excludedRoles = excludedRoles;
            return this;
        }

        /**
         * @param excludedRoles A list of role IDs excluded from scope of policy.
         * 
         * @return builder
         * 
         */
        public Builder excludedRoles(List<String> excludedRoles) {
            return excludedRoles(Output.of(excludedRoles));
        }

        /**
         * @param excludedRoles A list of role IDs excluded from scope of policy.
         * 
         * @return builder
         * 
         */
        public Builder excludedRoles(String... excludedRoles) {
            return excludedRoles(List.of(excludedRoles));
        }

        /**
         * @param excludedUsers A list of user IDs excluded from scope of policy and/or `GuestsOrExternalUsers`.
         * 
         * @return builder
         * 
         */
        public Builder excludedUsers(@Nullable Output<List<String>> excludedUsers) {
            $.excludedUsers = excludedUsers;
            return this;
        }

        /**
         * @param excludedUsers A list of user IDs excluded from scope of policy and/or `GuestsOrExternalUsers`.
         * 
         * @return builder
         * 
         */
        public Builder excludedUsers(List<String> excludedUsers) {
            return excludedUsers(Output.of(excludedUsers));
        }

        /**
         * @param excludedUsers A list of user IDs excluded from scope of policy and/or `GuestsOrExternalUsers`.
         * 
         * @return builder
         * 
         */
        public Builder excludedUsers(String... excludedUsers) {
            return excludedUsers(List.of(excludedUsers));
        }

        /**
         * @param includedGroups A list of group IDs in scope of policy unless explicitly excluded.
         * 
         * @return builder
         * 
         */
        public Builder includedGroups(@Nullable Output<List<String>> includedGroups) {
            $.includedGroups = includedGroups;
            return this;
        }

        /**
         * @param includedGroups A list of group IDs in scope of policy unless explicitly excluded.
         * 
         * @return builder
         * 
         */
        public Builder includedGroups(List<String> includedGroups) {
            return includedGroups(Output.of(includedGroups));
        }

        /**
         * @param includedGroups A list of group IDs in scope of policy unless explicitly excluded.
         * 
         * @return builder
         * 
         */
        public Builder includedGroups(String... includedGroups) {
            return includedGroups(List.of(includedGroups));
        }

        /**
         * @param includedRoles A list of role IDs in scope of policy unless explicitly excluded.
         * 
         * @return builder
         * 
         */
        public Builder includedRoles(@Nullable Output<List<String>> includedRoles) {
            $.includedRoles = includedRoles;
            return this;
        }

        /**
         * @param includedRoles A list of role IDs in scope of policy unless explicitly excluded.
         * 
         * @return builder
         * 
         */
        public Builder includedRoles(List<String> includedRoles) {
            return includedRoles(Output.of(includedRoles));
        }

        /**
         * @param includedRoles A list of role IDs in scope of policy unless explicitly excluded.
         * 
         * @return builder
         * 
         */
        public Builder includedRoles(String... includedRoles) {
            return includedRoles(List.of(includedRoles));
        }

        /**
         * @param includedUsers A list of user IDs in scope of policy unless explicitly excluded, or `None` or `All` or `GuestsOrExternalUsers`.
         * 
         * &gt; At least one of `included_groups`, `included_roles` or `included_users` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder includedUsers(@Nullable Output<List<String>> includedUsers) {
            $.includedUsers = includedUsers;
            return this;
        }

        /**
         * @param includedUsers A list of user IDs in scope of policy unless explicitly excluded, or `None` or `All` or `GuestsOrExternalUsers`.
         * 
         * &gt; At least one of `included_groups`, `included_roles` or `included_users` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder includedUsers(List<String> includedUsers) {
            return includedUsers(Output.of(includedUsers));
        }

        /**
         * @param includedUsers A list of user IDs in scope of policy unless explicitly excluded, or `None` or `All` or `GuestsOrExternalUsers`.
         * 
         * &gt; At least one of `included_groups`, `included_roles` or `included_users` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder includedUsers(String... includedUsers) {
            return includedUsers(List.of(includedUsers));
        }

        public ConditionalAccessPolicyConditionsUsersArgs build() {
            return $;
        }
    }

}

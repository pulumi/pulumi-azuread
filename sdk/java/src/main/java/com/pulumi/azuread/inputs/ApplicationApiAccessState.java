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


public final class ApplicationApiAccessState extends com.pulumi.resources.ResourceArgs {

    public static final ApplicationApiAccessState Empty = new ApplicationApiAccessState();

    /**
     * The client ID of the API to which access is being granted. Changing this forces a new resource to be created.
     * 
     */
    @Import(name="apiClientId")
    private @Nullable Output<String> apiClientId;

    /**
     * @return The client ID of the API to which access is being granted. Changing this forces a new resource to be created.
     * 
     */
    public Optional<Output<String>> apiClientId() {
        return Optional.ofNullable(this.apiClientId);
    }

    /**
     * The resource ID of the application registration. Changing this forces a new resource to be created.
     * 
     */
    @Import(name="applicationId")
    private @Nullable Output<String> applicationId;

    /**
     * @return The resource ID of the application registration. Changing this forces a new resource to be created.
     * 
     */
    public Optional<Output<String>> applicationId() {
        return Optional.ofNullable(this.applicationId);
    }

    /**
     * A set of role IDs to be granted to the application, as published by the API.
     * 
     */
    @Import(name="roleIds")
    private @Nullable Output<List<String>> roleIds;

    /**
     * @return A set of role IDs to be granted to the application, as published by the API.
     * 
     */
    public Optional<Output<List<String>>> roleIds() {
        return Optional.ofNullable(this.roleIds);
    }

    /**
     * A set of scope IDs to be granted to the application, as published by the API.
     * 
     * &gt; At least one of `role_ids` or `scope_ids` must be specified.
     * 
     */
    @Import(name="scopeIds")
    private @Nullable Output<List<String>> scopeIds;

    /**
     * @return A set of scope IDs to be granted to the application, as published by the API.
     * 
     * &gt; At least one of `role_ids` or `scope_ids` must be specified.
     * 
     */
    public Optional<Output<List<String>>> scopeIds() {
        return Optional.ofNullable(this.scopeIds);
    }

    private ApplicationApiAccessState() {}

    private ApplicationApiAccessState(ApplicationApiAccessState $) {
        this.apiClientId = $.apiClientId;
        this.applicationId = $.applicationId;
        this.roleIds = $.roleIds;
        this.scopeIds = $.scopeIds;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ApplicationApiAccessState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ApplicationApiAccessState $;

        public Builder() {
            $ = new ApplicationApiAccessState();
        }

        public Builder(ApplicationApiAccessState defaults) {
            $ = new ApplicationApiAccessState(Objects.requireNonNull(defaults));
        }

        /**
         * @param apiClientId The client ID of the API to which access is being granted. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder apiClientId(@Nullable Output<String> apiClientId) {
            $.apiClientId = apiClientId;
            return this;
        }

        /**
         * @param apiClientId The client ID of the API to which access is being granted. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder apiClientId(String apiClientId) {
            return apiClientId(Output.of(apiClientId));
        }

        /**
         * @param applicationId The resource ID of the application registration. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder applicationId(@Nullable Output<String> applicationId) {
            $.applicationId = applicationId;
            return this;
        }

        /**
         * @param applicationId The resource ID of the application registration. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder applicationId(String applicationId) {
            return applicationId(Output.of(applicationId));
        }

        /**
         * @param roleIds A set of role IDs to be granted to the application, as published by the API.
         * 
         * @return builder
         * 
         */
        public Builder roleIds(@Nullable Output<List<String>> roleIds) {
            $.roleIds = roleIds;
            return this;
        }

        /**
         * @param roleIds A set of role IDs to be granted to the application, as published by the API.
         * 
         * @return builder
         * 
         */
        public Builder roleIds(List<String> roleIds) {
            return roleIds(Output.of(roleIds));
        }

        /**
         * @param roleIds A set of role IDs to be granted to the application, as published by the API.
         * 
         * @return builder
         * 
         */
        public Builder roleIds(String... roleIds) {
            return roleIds(List.of(roleIds));
        }

        /**
         * @param scopeIds A set of scope IDs to be granted to the application, as published by the API.
         * 
         * &gt; At least one of `role_ids` or `scope_ids` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder scopeIds(@Nullable Output<List<String>> scopeIds) {
            $.scopeIds = scopeIds;
            return this;
        }

        /**
         * @param scopeIds A set of scope IDs to be granted to the application, as published by the API.
         * 
         * &gt; At least one of `role_ids` or `scope_ids` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder scopeIds(List<String> scopeIds) {
            return scopeIds(Output.of(scopeIds));
        }

        /**
         * @param scopeIds A set of scope IDs to be granted to the application, as published by the API.
         * 
         * &gt; At least one of `role_ids` or `scope_ids` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder scopeIds(String... scopeIds) {
            return scopeIds(List.of(scopeIds));
        }

        public ApplicationApiAccessState build() {
            return $;
        }
    }

}

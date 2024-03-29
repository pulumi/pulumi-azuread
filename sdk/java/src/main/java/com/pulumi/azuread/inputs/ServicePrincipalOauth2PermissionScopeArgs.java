// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServicePrincipalOauth2PermissionScopeArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServicePrincipalOauth2PermissionScopeArgs Empty = new ServicePrincipalOauth2PermissionScopeArgs();

    /**
     * Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
     * 
     */
    @Import(name="adminConsentDescription")
    private @Nullable Output<String> adminConsentDescription;

    /**
     * @return Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
     * 
     */
    public Optional<Output<String>> adminConsentDescription() {
        return Optional.ofNullable(this.adminConsentDescription);
    }

    /**
     * Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
     * 
     */
    @Import(name="adminConsentDisplayName")
    private @Nullable Output<String> adminConsentDisplayName;

    /**
     * @return Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
     * 
     */
    public Optional<Output<String>> adminConsentDisplayName() {
        return Optional.ofNullable(this.adminConsentDisplayName);
    }

    /**
     * Specifies whether the permission scope is enabled.
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return Specifies whether the permission scope is enabled.
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * The unique identifier of the delegated permission.
     * 
     */
    @Import(name="id")
    private @Nullable Output<String> id;

    /**
     * @return The unique identifier of the delegated permission.
     * 
     */
    public Optional<Output<String>> id() {
        return Optional.ofNullable(this.id);
    }

    /**
     * Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Possible values are `User` or `Admin`.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Possible values are `User` or `Admin`.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    /**
     * Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
     * 
     */
    @Import(name="userConsentDescription")
    private @Nullable Output<String> userConsentDescription;

    /**
     * @return Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
     * 
     */
    public Optional<Output<String>> userConsentDescription() {
        return Optional.ofNullable(this.userConsentDescription);
    }

    /**
     * Display name for the delegated permission that appears in the end user consent experience.
     * 
     */
    @Import(name="userConsentDisplayName")
    private @Nullable Output<String> userConsentDisplayName;

    /**
     * @return Display name for the delegated permission that appears in the end user consent experience.
     * 
     */
    public Optional<Output<String>> userConsentDisplayName() {
        return Optional.ofNullable(this.userConsentDisplayName);
    }

    /**
     * The value that is used for the `scp` claim in OAuth 2.0 access tokens.
     * 
     */
    @Import(name="value")
    private @Nullable Output<String> value;

    /**
     * @return The value that is used for the `scp` claim in OAuth 2.0 access tokens.
     * 
     */
    public Optional<Output<String>> value() {
        return Optional.ofNullable(this.value);
    }

    private ServicePrincipalOauth2PermissionScopeArgs() {}

    private ServicePrincipalOauth2PermissionScopeArgs(ServicePrincipalOauth2PermissionScopeArgs $) {
        this.adminConsentDescription = $.adminConsentDescription;
        this.adminConsentDisplayName = $.adminConsentDisplayName;
        this.enabled = $.enabled;
        this.id = $.id;
        this.type = $.type;
        this.userConsentDescription = $.userConsentDescription;
        this.userConsentDisplayName = $.userConsentDisplayName;
        this.value = $.value;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServicePrincipalOauth2PermissionScopeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServicePrincipalOauth2PermissionScopeArgs $;

        public Builder() {
            $ = new ServicePrincipalOauth2PermissionScopeArgs();
        }

        public Builder(ServicePrincipalOauth2PermissionScopeArgs defaults) {
            $ = new ServicePrincipalOauth2PermissionScopeArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param adminConsentDescription Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
         * 
         * @return builder
         * 
         */
        public Builder adminConsentDescription(@Nullable Output<String> adminConsentDescription) {
            $.adminConsentDescription = adminConsentDescription;
            return this;
        }

        /**
         * @param adminConsentDescription Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
         * 
         * @return builder
         * 
         */
        public Builder adminConsentDescription(String adminConsentDescription) {
            return adminConsentDescription(Output.of(adminConsentDescription));
        }

        /**
         * @param adminConsentDisplayName Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
         * 
         * @return builder
         * 
         */
        public Builder adminConsentDisplayName(@Nullable Output<String> adminConsentDisplayName) {
            $.adminConsentDisplayName = adminConsentDisplayName;
            return this;
        }

        /**
         * @param adminConsentDisplayName Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
         * 
         * @return builder
         * 
         */
        public Builder adminConsentDisplayName(String adminConsentDisplayName) {
            return adminConsentDisplayName(Output.of(adminConsentDisplayName));
        }

        /**
         * @param enabled Specifies whether the permission scope is enabled.
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Specifies whether the permission scope is enabled.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param id The unique identifier of the delegated permission.
         * 
         * @return builder
         * 
         */
        public Builder id(@Nullable Output<String> id) {
            $.id = id;
            return this;
        }

        /**
         * @param id The unique identifier of the delegated permission.
         * 
         * @return builder
         * 
         */
        public Builder id(String id) {
            return id(Output.of(id));
        }

        /**
         * @param type Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Possible values are `User` or `Admin`.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Possible values are `User` or `Admin`.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        /**
         * @param userConsentDescription Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
         * 
         * @return builder
         * 
         */
        public Builder userConsentDescription(@Nullable Output<String> userConsentDescription) {
            $.userConsentDescription = userConsentDescription;
            return this;
        }

        /**
         * @param userConsentDescription Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
         * 
         * @return builder
         * 
         */
        public Builder userConsentDescription(String userConsentDescription) {
            return userConsentDescription(Output.of(userConsentDescription));
        }

        /**
         * @param userConsentDisplayName Display name for the delegated permission that appears in the end user consent experience.
         * 
         * @return builder
         * 
         */
        public Builder userConsentDisplayName(@Nullable Output<String> userConsentDisplayName) {
            $.userConsentDisplayName = userConsentDisplayName;
            return this;
        }

        /**
         * @param userConsentDisplayName Display name for the delegated permission that appears in the end user consent experience.
         * 
         * @return builder
         * 
         */
        public Builder userConsentDisplayName(String userConsentDisplayName) {
            return userConsentDisplayName(Output.of(userConsentDisplayName));
        }

        /**
         * @param value The value that is used for the `scp` claim in OAuth 2.0 access tokens.
         * 
         * @return builder
         * 
         */
        public Builder value(@Nullable Output<String> value) {
            $.value = value;
            return this;
        }

        /**
         * @param value The value that is used for the `scp` claim in OAuth 2.0 access tokens.
         * 
         * @return builder
         * 
         */
        public Builder value(String value) {
            return value(Output.of(value));
        }

        public ServicePrincipalOauth2PermissionScopeArgs build() {
            return $;
        }
    }

}

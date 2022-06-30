// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetApplicationApiOauth2PermissionScope {
    /**
     * @return Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
     * 
     */
    private final String adminConsentDescription;
    /**
     * @return Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
     * 
     */
    private final String adminConsentDisplayName;
    /**
     * @return Determines if the app role is enabled.
     * 
     */
    private final Boolean enabled;
    /**
     * @return The unique identifier for an app role or OAuth2 permission scope published by the resource application.
     * 
     */
    private final String id;
    /**
     * @return Specifies whether the `id` property references an app role or an OAuth2 permission scope. Possible values are `Role` or `Scope`.
     * 
     */
    private final String type;
    /**
     * @return Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
     * 
     */
    private final String userConsentDescription;
    /**
     * @return Display name for the delegated permission that appears in the end user consent experience.
     * 
     */
    private final String userConsentDisplayName;
    /**
     * @return The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal.
     * 
     */
    private final String value;

    @CustomType.Constructor
    private GetApplicationApiOauth2PermissionScope(
        @CustomType.Parameter("adminConsentDescription") String adminConsentDescription,
        @CustomType.Parameter("adminConsentDisplayName") String adminConsentDisplayName,
        @CustomType.Parameter("enabled") Boolean enabled,
        @CustomType.Parameter("id") String id,
        @CustomType.Parameter("type") String type,
        @CustomType.Parameter("userConsentDescription") String userConsentDescription,
        @CustomType.Parameter("userConsentDisplayName") String userConsentDisplayName,
        @CustomType.Parameter("value") String value) {
        this.adminConsentDescription = adminConsentDescription;
        this.adminConsentDisplayName = adminConsentDisplayName;
        this.enabled = enabled;
        this.id = id;
        this.type = type;
        this.userConsentDescription = userConsentDescription;
        this.userConsentDisplayName = userConsentDisplayName;
        this.value = value;
    }

    /**
     * @return Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
     * 
     */
    public String adminConsentDescription() {
        return this.adminConsentDescription;
    }
    /**
     * @return Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
     * 
     */
    public String adminConsentDisplayName() {
        return this.adminConsentDisplayName;
    }
    /**
     * @return Determines if the app role is enabled.
     * 
     */
    public Boolean enabled() {
        return this.enabled;
    }
    /**
     * @return The unique identifier for an app role or OAuth2 permission scope published by the resource application.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Specifies whether the `id` property references an app role or an OAuth2 permission scope. Possible values are `Role` or `Scope`.
     * 
     */
    public String type() {
        return this.type;
    }
    /**
     * @return Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
     * 
     */
    public String userConsentDescription() {
        return this.userConsentDescription;
    }
    /**
     * @return Display name for the delegated permission that appears in the end user consent experience.
     * 
     */
    public String userConsentDisplayName() {
        return this.userConsentDisplayName;
    }
    /**
     * @return The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal.
     * 
     */
    public String value() {
        return this.value;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetApplicationApiOauth2PermissionScope defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String adminConsentDescription;
        private String adminConsentDisplayName;
        private Boolean enabled;
        private String id;
        private String type;
        private String userConsentDescription;
        private String userConsentDisplayName;
        private String value;

        public Builder() {
    	      // Empty
        }

        public Builder(GetApplicationApiOauth2PermissionScope defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.adminConsentDescription = defaults.adminConsentDescription;
    	      this.adminConsentDisplayName = defaults.adminConsentDisplayName;
    	      this.enabled = defaults.enabled;
    	      this.id = defaults.id;
    	      this.type = defaults.type;
    	      this.userConsentDescription = defaults.userConsentDescription;
    	      this.userConsentDisplayName = defaults.userConsentDisplayName;
    	      this.value = defaults.value;
        }

        public Builder adminConsentDescription(String adminConsentDescription) {
            this.adminConsentDescription = Objects.requireNonNull(adminConsentDescription);
            return this;
        }
        public Builder adminConsentDisplayName(String adminConsentDisplayName) {
            this.adminConsentDisplayName = Objects.requireNonNull(adminConsentDisplayName);
            return this;
        }
        public Builder enabled(Boolean enabled) {
            this.enabled = Objects.requireNonNull(enabled);
            return this;
        }
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        public Builder userConsentDescription(String userConsentDescription) {
            this.userConsentDescription = Objects.requireNonNull(userConsentDescription);
            return this;
        }
        public Builder userConsentDisplayName(String userConsentDisplayName) {
            this.userConsentDisplayName = Objects.requireNonNull(userConsentDisplayName);
            return this;
        }
        public Builder value(String value) {
            this.value = Objects.requireNonNull(value);
            return this;
        }        public GetApplicationApiOauth2PermissionScope build() {
            return new GetApplicationApiOauth2PermissionScope(adminConsentDescription, adminConsentDisplayName, enabled, id, type, userConsentDescription, userConsentDisplayName, value);
        }
    }
}

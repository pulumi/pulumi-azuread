// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetApplicationApiOauth2PermissionScope {
    /**
     * @return Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
     * 
     */
    private String adminConsentDescription;
    /**
     * @return Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
     * 
     */
    private String adminConsentDisplayName;
    /**
     * @return Determines if the app role is enabled.
     * 
     */
    private Boolean enabled;
    /**
     * @return The unique identifier for an app role or OAuth2 permission scope published by the resource application.
     * 
     */
    private String id;
    /**
     * @return Specifies whether the `id` property references an app role or an OAuth2 permission scope. Possible values are `Role` or `Scope`.
     * 
     */
    private String type;
    /**
     * @return Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
     * 
     */
    private String userConsentDescription;
    /**
     * @return Display name for the delegated permission that appears in the end user consent experience.
     * 
     */
    private String userConsentDisplayName;
    /**
     * @return The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal.
     * 
     */
    private String value;

    private GetApplicationApiOauth2PermissionScope() {}
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
    @CustomType.Builder
    public static final class Builder {
        private String adminConsentDescription;
        private String adminConsentDisplayName;
        private Boolean enabled;
        private String id;
        private String type;
        private String userConsentDescription;
        private String userConsentDisplayName;
        private String value;
        public Builder() {}
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

        @CustomType.Setter
        public Builder adminConsentDescription(String adminConsentDescription) {
            if (adminConsentDescription == null) {
              throw new MissingRequiredPropertyException("GetApplicationApiOauth2PermissionScope", "adminConsentDescription");
            }
            this.adminConsentDescription = adminConsentDescription;
            return this;
        }
        @CustomType.Setter
        public Builder adminConsentDisplayName(String adminConsentDisplayName) {
            if (adminConsentDisplayName == null) {
              throw new MissingRequiredPropertyException("GetApplicationApiOauth2PermissionScope", "adminConsentDisplayName");
            }
            this.adminConsentDisplayName = adminConsentDisplayName;
            return this;
        }
        @CustomType.Setter
        public Builder enabled(Boolean enabled) {
            if (enabled == null) {
              throw new MissingRequiredPropertyException("GetApplicationApiOauth2PermissionScope", "enabled");
            }
            this.enabled = enabled;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetApplicationApiOauth2PermissionScope", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            if (type == null) {
              throw new MissingRequiredPropertyException("GetApplicationApiOauth2PermissionScope", "type");
            }
            this.type = type;
            return this;
        }
        @CustomType.Setter
        public Builder userConsentDescription(String userConsentDescription) {
            if (userConsentDescription == null) {
              throw new MissingRequiredPropertyException("GetApplicationApiOauth2PermissionScope", "userConsentDescription");
            }
            this.userConsentDescription = userConsentDescription;
            return this;
        }
        @CustomType.Setter
        public Builder userConsentDisplayName(String userConsentDisplayName) {
            if (userConsentDisplayName == null) {
              throw new MissingRequiredPropertyException("GetApplicationApiOauth2PermissionScope", "userConsentDisplayName");
            }
            this.userConsentDisplayName = userConsentDisplayName;
            return this;
        }
        @CustomType.Setter
        public Builder value(String value) {
            if (value == null) {
              throw new MissingRequiredPropertyException("GetApplicationApiOauth2PermissionScope", "value");
            }
            this.value = value;
            return this;
        }
        public GetApplicationApiOauth2PermissionScope build() {
            final var _resultValue = new GetApplicationApiOauth2PermissionScope();
            _resultValue.adminConsentDescription = adminConsentDescription;
            _resultValue.adminConsentDisplayName = adminConsentDisplayName;
            _resultValue.enabled = enabled;
            _resultValue.id = id;
            _resultValue.type = type;
            _resultValue.userConsentDescription = userConsentDescription;
            _resultValue.userConsentDisplayName = userConsentDisplayName;
            _resultValue.value = value;
            return _resultValue;
        }
    }
}

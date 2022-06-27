// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetApplicationAppRole {
    /**
     * @return Specifies whether this app role definition can be assigned to users and groups, or to other applications (that are accessing this application in a standalone scenario). Possible values are `User` or `Application`, or both.
     * 
     */
    private final List<String> allowedMemberTypes;
    /**
     * @return Description of the app role that appears when the role is being assigned and, if the role functions as an application permissions, during the consent experiences.
     * 
     */
    private final String description;
    /**
     * @return Specifies the display name of the application.
     * 
     */
    private final String displayName;
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
     * @return The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal.
     * 
     */
    private final String value;

    @CustomType.Constructor
    private GetApplicationAppRole(
        @CustomType.Parameter("allowedMemberTypes") List<String> allowedMemberTypes,
        @CustomType.Parameter("description") String description,
        @CustomType.Parameter("displayName") String displayName,
        @CustomType.Parameter("enabled") Boolean enabled,
        @CustomType.Parameter("id") String id,
        @CustomType.Parameter("value") String value) {
        this.allowedMemberTypes = allowedMemberTypes;
        this.description = description;
        this.displayName = displayName;
        this.enabled = enabled;
        this.id = id;
        this.value = value;
    }

    /**
     * @return Specifies whether this app role definition can be assigned to users and groups, or to other applications (that are accessing this application in a standalone scenario). Possible values are `User` or `Application`, or both.
     * 
     */
    public List<String> allowedMemberTypes() {
        return this.allowedMemberTypes;
    }
    /**
     * @return Description of the app role that appears when the role is being assigned and, if the role functions as an application permissions, during the consent experiences.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return Specifies the display name of the application.
     * 
     */
    public String displayName() {
        return this.displayName;
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
     * @return The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal.
     * 
     */
    public String value() {
        return this.value;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetApplicationAppRole defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private List<String> allowedMemberTypes;
        private String description;
        private String displayName;
        private Boolean enabled;
        private String id;
        private String value;

        public Builder() {
    	      // Empty
        }

        public Builder(GetApplicationAppRole defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.allowedMemberTypes = defaults.allowedMemberTypes;
    	      this.description = defaults.description;
    	      this.displayName = defaults.displayName;
    	      this.enabled = defaults.enabled;
    	      this.id = defaults.id;
    	      this.value = defaults.value;
        }

        public Builder allowedMemberTypes(List<String> allowedMemberTypes) {
            this.allowedMemberTypes = Objects.requireNonNull(allowedMemberTypes);
            return this;
        }
        public Builder allowedMemberTypes(String... allowedMemberTypes) {
            return allowedMemberTypes(List.of(allowedMemberTypes));
        }
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        public Builder displayName(String displayName) {
            this.displayName = Objects.requireNonNull(displayName);
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
        public Builder value(String value) {
            this.value = Objects.requireNonNull(value);
            return this;
        }        public GetApplicationAppRole build() {
            return new GetApplicationAppRole(allowedMemberTypes, description, displayName, enabled, id, value);
        }
    }
}

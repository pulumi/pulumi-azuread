// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.azuread.outputs.GetDirectoryRolesRole;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetDirectoryRolesResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The object IDs of the roles.
     * 
     */
    private List<String> objectIds;
    /**
     * @return A list of users. Each `role` object provides the attributes documented below.
     * 
     */
    private List<GetDirectoryRolesRole> roles;
    /**
     * @return The template IDs of the roles.
     * 
     */
    private List<String> templateIds;

    private GetDirectoryRolesResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The object IDs of the roles.
     * 
     */
    public List<String> objectIds() {
        return this.objectIds;
    }
    /**
     * @return A list of users. Each `role` object provides the attributes documented below.
     * 
     */
    public List<GetDirectoryRolesRole> roles() {
        return this.roles;
    }
    /**
     * @return The template IDs of the roles.
     * 
     */
    public List<String> templateIds() {
        return this.templateIds;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDirectoryRolesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private List<String> objectIds;
        private List<GetDirectoryRolesRole> roles;
        private List<String> templateIds;
        public Builder() {}
        public Builder(GetDirectoryRolesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.objectIds = defaults.objectIds;
    	      this.roles = defaults.roles;
    	      this.templateIds = defaults.templateIds;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder objectIds(List<String> objectIds) {
            this.objectIds = Objects.requireNonNull(objectIds);
            return this;
        }
        public Builder objectIds(String... objectIds) {
            return objectIds(List.of(objectIds));
        }
        @CustomType.Setter
        public Builder roles(List<GetDirectoryRolesRole> roles) {
            this.roles = Objects.requireNonNull(roles);
            return this;
        }
        public Builder roles(GetDirectoryRolesRole... roles) {
            return roles(List.of(roles));
        }
        @CustomType.Setter
        public Builder templateIds(List<String> templateIds) {
            this.templateIds = Objects.requireNonNull(templateIds);
            return this;
        }
        public Builder templateIds(String... templateIds) {
            return templateIds(List.of(templateIds));
        }
        public GetDirectoryRolesResult build() {
            final var o = new GetDirectoryRolesResult();
            o.id = id;
            o.objectIds = objectIds;
            o.roles = roles;
            o.templateIds = templateIds;
            return o;
        }
    }
}

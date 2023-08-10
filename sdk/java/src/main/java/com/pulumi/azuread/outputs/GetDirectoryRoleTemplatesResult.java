// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.azuread.outputs.GetDirectoryRoleTemplatesRoleTemplate;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetDirectoryRoleTemplatesResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The object IDs of the role templates.
     * 
     */
    private List<String> objectIds;
    /**
     * @return A list of role templates. Each `role_template` object provides the attributes documented below.
     * 
     */
    private List<GetDirectoryRoleTemplatesRoleTemplate> roleTemplates;

    private GetDirectoryRoleTemplatesResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The object IDs of the role templates.
     * 
     */
    public List<String> objectIds() {
        return this.objectIds;
    }
    /**
     * @return A list of role templates. Each `role_template` object provides the attributes documented below.
     * 
     */
    public List<GetDirectoryRoleTemplatesRoleTemplate> roleTemplates() {
        return this.roleTemplates;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDirectoryRoleTemplatesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private List<String> objectIds;
        private List<GetDirectoryRoleTemplatesRoleTemplate> roleTemplates;
        public Builder() {}
        public Builder(GetDirectoryRoleTemplatesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.objectIds = defaults.objectIds;
    	      this.roleTemplates = defaults.roleTemplates;
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
        public Builder roleTemplates(List<GetDirectoryRoleTemplatesRoleTemplate> roleTemplates) {
            this.roleTemplates = Objects.requireNonNull(roleTemplates);
            return this;
        }
        public Builder roleTemplates(GetDirectoryRoleTemplatesRoleTemplate... roleTemplates) {
            return roleTemplates(List.of(roleTemplates));
        }
        public GetDirectoryRoleTemplatesResult build() {
            final var o = new GetDirectoryRoleTemplatesResult();
            o.id = id;
            o.objectIds = objectIds;
            o.roleTemplates = roleTemplates;
            return o;
        }
    }
}
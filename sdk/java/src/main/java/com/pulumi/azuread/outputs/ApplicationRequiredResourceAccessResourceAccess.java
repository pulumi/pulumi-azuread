// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class ApplicationRequiredResourceAccessResourceAccess {
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

    @CustomType.Constructor
    private ApplicationRequiredResourceAccessResourceAccess(
        @CustomType.Parameter("id") String id,
        @CustomType.Parameter("type") String type) {
        this.id = id;
        this.type = type;
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

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ApplicationRequiredResourceAccessResourceAccess defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String id;
        private String type;

        public Builder() {
    	      // Empty
        }

        public Builder(ApplicationRequiredResourceAccessResourceAccess defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.type = defaults.type;
        }

        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }        public ApplicationRequiredResourceAccessResourceAccess build() {
            return new ApplicationRequiredResourceAccessResourceAccess(id, type);
        }
    }
}

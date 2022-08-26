// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetClientConfigResult {
    /**
     * @return The client ID (application ID) linked to the authenticated principal, or the application used for delegated authentication.
     * 
     */
    private String clientId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The object ID of the authenticated principal.
     * 
     */
    private String objectId;
    /**
     * @return The tenant ID of the authenticated principal.
     * 
     */
    private String tenantId;

    private GetClientConfigResult() {}
    /**
     * @return The client ID (application ID) linked to the authenticated principal, or the application used for delegated authentication.
     * 
     */
    public String clientId() {
        return this.clientId;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The object ID of the authenticated principal.
     * 
     */
    public String objectId() {
        return this.objectId;
    }
    /**
     * @return The tenant ID of the authenticated principal.
     * 
     */
    public String tenantId() {
        return this.tenantId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetClientConfigResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String clientId;
        private String id;
        private String objectId;
        private String tenantId;
        public Builder() {}
        public Builder(GetClientConfigResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clientId = defaults.clientId;
    	      this.id = defaults.id;
    	      this.objectId = defaults.objectId;
    	      this.tenantId = defaults.tenantId;
        }

        @CustomType.Setter
        public Builder clientId(String clientId) {
            this.clientId = Objects.requireNonNull(clientId);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder objectId(String objectId) {
            this.objectId = Objects.requireNonNull(objectId);
            return this;
        }
        @CustomType.Setter
        public Builder tenantId(String tenantId) {
            this.tenantId = Objects.requireNonNull(tenantId);
            return this;
        }
        public GetClientConfigResult build() {
            final var o = new GetClientConfigResult();
            o.clientId = clientId;
            o.id = id;
            o.objectId = objectId;
            o.tenantId = tenantId;
            return o;
        }
    }
}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.azuread.outputs.GetApplicationRequiredResourceAccessResourceAccess;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetApplicationRequiredResourceAccess {
    /**
     * @return A collection of `resource_access` blocks as documented below, describing OAuth2.0 permission scopes and app roles that the application requires from the specified resource.
     * 
     */
    private List<GetApplicationRequiredResourceAccessResourceAccess> resourceAccesses;
    /**
     * @return The unique identifier for the resource that the application requires access to. This is the Application ID of the target application.
     * 
     */
    private String resourceAppId;

    private GetApplicationRequiredResourceAccess() {}
    /**
     * @return A collection of `resource_access` blocks as documented below, describing OAuth2.0 permission scopes and app roles that the application requires from the specified resource.
     * 
     */
    public List<GetApplicationRequiredResourceAccessResourceAccess> resourceAccesses() {
        return this.resourceAccesses;
    }
    /**
     * @return The unique identifier for the resource that the application requires access to. This is the Application ID of the target application.
     * 
     */
    public String resourceAppId() {
        return this.resourceAppId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetApplicationRequiredResourceAccess defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetApplicationRequiredResourceAccessResourceAccess> resourceAccesses;
        private String resourceAppId;
        public Builder() {}
        public Builder(GetApplicationRequiredResourceAccess defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.resourceAccesses = defaults.resourceAccesses;
    	      this.resourceAppId = defaults.resourceAppId;
        }

        @CustomType.Setter
        public Builder resourceAccesses(List<GetApplicationRequiredResourceAccessResourceAccess> resourceAccesses) {
            if (resourceAccesses == null) {
              throw new MissingRequiredPropertyException("GetApplicationRequiredResourceAccess", "resourceAccesses");
            }
            this.resourceAccesses = resourceAccesses;
            return this;
        }
        public Builder resourceAccesses(GetApplicationRequiredResourceAccessResourceAccess... resourceAccesses) {
            return resourceAccesses(List.of(resourceAccesses));
        }
        @CustomType.Setter
        public Builder resourceAppId(String resourceAppId) {
            if (resourceAppId == null) {
              throw new MissingRequiredPropertyException("GetApplicationRequiredResourceAccess", "resourceAppId");
            }
            this.resourceAppId = resourceAppId;
            return this;
        }
        public GetApplicationRequiredResourceAccess build() {
            final var _resultValue = new GetApplicationRequiredResourceAccess();
            _resultValue.resourceAccesses = resourceAccesses;
            _resultValue.resourceAppId = resourceAppId;
            return _resultValue;
        }
    }
}

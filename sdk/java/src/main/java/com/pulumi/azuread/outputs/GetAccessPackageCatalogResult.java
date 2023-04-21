// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetAccessPackageCatalogResult {
    /**
     * @return The description of the access package catalog.
     * 
     */
    private String description;
    private String displayName;
    /**
     * @return Whether the access packages in this catalog can be requested by users outside the tenant.
     * 
     */
    private Boolean externallyVisible;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String objectId;
    /**
     * @return Whether the access packages in this catalog are available for management.
     * 
     */
    private Boolean published;

    private GetAccessPackageCatalogResult() {}
    /**
     * @return The description of the access package catalog.
     * 
     */
    public String description() {
        return this.description;
    }
    public String displayName() {
        return this.displayName;
    }
    /**
     * @return Whether the access packages in this catalog can be requested by users outside the tenant.
     * 
     */
    public Boolean externallyVisible() {
        return this.externallyVisible;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String objectId() {
        return this.objectId;
    }
    /**
     * @return Whether the access packages in this catalog are available for management.
     * 
     */
    public Boolean published() {
        return this.published;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAccessPackageCatalogResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String description;
        private String displayName;
        private Boolean externallyVisible;
        private String id;
        private String objectId;
        private Boolean published;
        public Builder() {}
        public Builder(GetAccessPackageCatalogResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.displayName = defaults.displayName;
    	      this.externallyVisible = defaults.externallyVisible;
    	      this.id = defaults.id;
    	      this.objectId = defaults.objectId;
    	      this.published = defaults.published;
        }

        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder displayName(String displayName) {
            this.displayName = Objects.requireNonNull(displayName);
            return this;
        }
        @CustomType.Setter
        public Builder externallyVisible(Boolean externallyVisible) {
            this.externallyVisible = Objects.requireNonNull(externallyVisible);
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
        public Builder published(Boolean published) {
            this.published = Objects.requireNonNull(published);
            return this;
        }
        public GetAccessPackageCatalogResult build() {
            final var o = new GetAccessPackageCatalogResult();
            o.description = description;
            o.displayName = displayName;
            o.externallyVisible = externallyVisible;
            o.id = id;
            o.objectId = objectId;
            o.published = published;
            return o;
        }
    }
}
// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetAccessPackageResult {
    private @Nullable String catalogId;
    /**
     * @return The description of the access package.
     * 
     */
    private String description;
    private String displayName;
    /**
     * @return Whether the access package is hidden from the requestor.
     * 
     */
    private Boolean hidden;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String objectId;

    private GetAccessPackageResult() {}
    public Optional<String> catalogId() {
        return Optional.ofNullable(this.catalogId);
    }
    /**
     * @return The description of the access package.
     * 
     */
    public String description() {
        return this.description;
    }
    public String displayName() {
        return this.displayName;
    }
    /**
     * @return Whether the access package is hidden from the requestor.
     * 
     */
    public Boolean hidden() {
        return this.hidden;
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

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAccessPackageResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String catalogId;
        private String description;
        private String displayName;
        private Boolean hidden;
        private String id;
        private String objectId;
        public Builder() {}
        public Builder(GetAccessPackageResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.catalogId = defaults.catalogId;
    	      this.description = defaults.description;
    	      this.displayName = defaults.displayName;
    	      this.hidden = defaults.hidden;
    	      this.id = defaults.id;
    	      this.objectId = defaults.objectId;
        }

        @CustomType.Setter
        public Builder catalogId(@Nullable String catalogId) {

            this.catalogId = catalogId;
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetAccessPackageResult", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder displayName(String displayName) {
            if (displayName == null) {
              throw new MissingRequiredPropertyException("GetAccessPackageResult", "displayName");
            }
            this.displayName = displayName;
            return this;
        }
        @CustomType.Setter
        public Builder hidden(Boolean hidden) {
            if (hidden == null) {
              throw new MissingRequiredPropertyException("GetAccessPackageResult", "hidden");
            }
            this.hidden = hidden;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetAccessPackageResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder objectId(String objectId) {
            if (objectId == null) {
              throw new MissingRequiredPropertyException("GetAccessPackageResult", "objectId");
            }
            this.objectId = objectId;
            return this;
        }
        public GetAccessPackageResult build() {
            final var _resultValue = new GetAccessPackageResult();
            _resultValue.catalogId = catalogId;
            _resultValue.description = description;
            _resultValue.displayName = displayName;
            _resultValue.hidden = hidden;
            _resultValue.id = id;
            _resultValue.objectId = objectId;
            return _resultValue;
        }
    }
}

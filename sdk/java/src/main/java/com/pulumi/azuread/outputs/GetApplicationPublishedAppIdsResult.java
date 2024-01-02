// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetApplicationPublishedAppIdsResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return A map of application names to application IDs.
     * 
     */
    private Map<String,String> result;

    private GetApplicationPublishedAppIdsResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return A map of application names to application IDs.
     * 
     */
    public Map<String,String> result() {
        return this.result;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetApplicationPublishedAppIdsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private Map<String,String> result;
        public Builder() {}
        public Builder(GetApplicationPublishedAppIdsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.result = defaults.result;
        }

        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetApplicationPublishedAppIdsResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder result(Map<String,String> result) {
            if (result == null) {
              throw new MissingRequiredPropertyException("GetApplicationPublishedAppIdsResult", "result");
            }
            this.result = result;
            return this;
        }
        public GetApplicationPublishedAppIdsResult build() {
            final var _resultValue = new GetApplicationPublishedAppIdsResult();
            _resultValue.id = id;
            _resultValue.result = result;
            return _resultValue;
        }
    }
}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetApplicationPublishedAppIdsResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private final String id;
    /**
     * @return A map of application names to application IDs.
     * 
     */
    private final Map<String,String> result;

    @CustomType.Constructor
    private GetApplicationPublishedAppIdsResult(
        @CustomType.Parameter("id") String id,
        @CustomType.Parameter("result") Map<String,String> result) {
        this.id = id;
        this.result = result;
    }

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

    public static final class Builder {
        private String id;
        private Map<String,String> result;

        public Builder() {
    	      // Empty
        }

        public Builder(GetApplicationPublishedAppIdsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.result = defaults.result;
        }

        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        public Builder result(Map<String,String> result) {
            this.result = Objects.requireNonNull(result);
            return this;
        }        public GetApplicationPublishedAppIdsResult build() {
            return new GetApplicationPublishedAppIdsResult(id, result);
        }
    }
}

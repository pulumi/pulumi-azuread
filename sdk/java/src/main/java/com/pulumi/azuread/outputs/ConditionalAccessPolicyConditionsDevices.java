// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.azuread.outputs.ConditionalAccessPolicyConditionsDevicesFilter;
import com.pulumi.core.annotations.CustomType;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ConditionalAccessPolicyConditionsDevices {
    /**
     * @return A `filter` block as described below.
     * 
     */
    private @Nullable ConditionalAccessPolicyConditionsDevicesFilter filter;

    private ConditionalAccessPolicyConditionsDevices() {}
    /**
     * @return A `filter` block as described below.
     * 
     */
    public Optional<ConditionalAccessPolicyConditionsDevicesFilter> filter() {
        return Optional.ofNullable(this.filter);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ConditionalAccessPolicyConditionsDevices defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable ConditionalAccessPolicyConditionsDevicesFilter filter;
        public Builder() {}
        public Builder(ConditionalAccessPolicyConditionsDevices defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.filter = defaults.filter;
        }

        @CustomType.Setter
        public Builder filter(@Nullable ConditionalAccessPolicyConditionsDevicesFilter filter) {
            this.filter = filter;
            return this;
        }
        public ConditionalAccessPolicyConditionsDevices build() {
            final var _resultValue = new ConditionalAccessPolicyConditionsDevices();
            _resultValue.filter = filter;
            return _resultValue;
        }
    }
}

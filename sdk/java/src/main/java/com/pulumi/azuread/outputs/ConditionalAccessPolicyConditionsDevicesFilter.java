// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class ConditionalAccessPolicyConditionsDevicesFilter {
    /**
     * @return Whether to include in, or exclude from, matching devices from the policy. Supported values are `include` or `exclude`.
     * 
     */
    private String mode;
    /**
     * @return Condition filter to match devices. For more information, see [official documentation](https://docs.microsoft.com/en-us/azure/active-directory/conditional-access/concept-condition-filters-for-devices#supported-operators-and-device-properties-for-filters).
     * 
     */
    private String rule;

    private ConditionalAccessPolicyConditionsDevicesFilter() {}
    /**
     * @return Whether to include in, or exclude from, matching devices from the policy. Supported values are `include` or `exclude`.
     * 
     */
    public String mode() {
        return this.mode;
    }
    /**
     * @return Condition filter to match devices. For more information, see [official documentation](https://docs.microsoft.com/en-us/azure/active-directory/conditional-access/concept-condition-filters-for-devices#supported-operators-and-device-properties-for-filters).
     * 
     */
    public String rule() {
        return this.rule;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ConditionalAccessPolicyConditionsDevicesFilter defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String mode;
        private String rule;
        public Builder() {}
        public Builder(ConditionalAccessPolicyConditionsDevicesFilter defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.mode = defaults.mode;
    	      this.rule = defaults.rule;
        }

        @CustomType.Setter
        public Builder mode(String mode) {
            if (mode == null) {
              throw new MissingRequiredPropertyException("ConditionalAccessPolicyConditionsDevicesFilter", "mode");
            }
            this.mode = mode;
            return this;
        }
        @CustomType.Setter
        public Builder rule(String rule) {
            if (rule == null) {
              throw new MissingRequiredPropertyException("ConditionalAccessPolicyConditionsDevicesFilter", "rule");
            }
            this.rule = rule;
            return this;
        }
        public ConditionalAccessPolicyConditionsDevicesFilter build() {
            final var _resultValue = new ConditionalAccessPolicyConditionsDevicesFilter();
            _resultValue.mode = mode;
            _resultValue.rule = rule;
            return _resultValue;
        }
    }
}

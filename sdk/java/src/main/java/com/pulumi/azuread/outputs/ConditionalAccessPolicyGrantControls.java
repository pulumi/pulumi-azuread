// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ConditionalAccessPolicyGrantControls {
    /**
     * @return ID of an Authentication Strength Policy to use in this policy. When using a hard-coded ID, the UUID value should be prefixed with: `/policies/authenticationStrengthPolicies/`.
     * 
     */
    private @Nullable String authenticationStrengthPolicyId;
    /**
     * @return List of built-in controls required by the policy. Possible values are: `block`, `mfa`, `approvedApplication`, `compliantApplication`, `compliantDevice`, `domainJoinedDevice`, `passwordChange` or `unknownFutureValue`.
     * 
     */
    private @Nullable List<String> builtInControls;
    /**
     * @return List of custom controls IDs required by the policy.
     * 
     */
    private @Nullable List<String> customAuthenticationFactors;
    /**
     * @return Defines the relationship of the grant controls. Possible values are: `AND`, `OR`.
     * 
     */
    private String operator;
    /**
     * @return List of terms of use IDs required by the policy.
     * 
     * &gt; At least one of `authentication_strength_policy_id`, `built_in_controls` or `terms_of_use` must be specified.
     * 
     */
    private @Nullable List<String> termsOfUses;

    private ConditionalAccessPolicyGrantControls() {}
    /**
     * @return ID of an Authentication Strength Policy to use in this policy. When using a hard-coded ID, the UUID value should be prefixed with: `/policies/authenticationStrengthPolicies/`.
     * 
     */
    public Optional<String> authenticationStrengthPolicyId() {
        return Optional.ofNullable(this.authenticationStrengthPolicyId);
    }
    /**
     * @return List of built-in controls required by the policy. Possible values are: `block`, `mfa`, `approvedApplication`, `compliantApplication`, `compliantDevice`, `domainJoinedDevice`, `passwordChange` or `unknownFutureValue`.
     * 
     */
    public List<String> builtInControls() {
        return this.builtInControls == null ? List.of() : this.builtInControls;
    }
    /**
     * @return List of custom controls IDs required by the policy.
     * 
     */
    public List<String> customAuthenticationFactors() {
        return this.customAuthenticationFactors == null ? List.of() : this.customAuthenticationFactors;
    }
    /**
     * @return Defines the relationship of the grant controls. Possible values are: `AND`, `OR`.
     * 
     */
    public String operator() {
        return this.operator;
    }
    /**
     * @return List of terms of use IDs required by the policy.
     * 
     * &gt; At least one of `authentication_strength_policy_id`, `built_in_controls` or `terms_of_use` must be specified.
     * 
     */
    public List<String> termsOfUses() {
        return this.termsOfUses == null ? List.of() : this.termsOfUses;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ConditionalAccessPolicyGrantControls defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String authenticationStrengthPolicyId;
        private @Nullable List<String> builtInControls;
        private @Nullable List<String> customAuthenticationFactors;
        private String operator;
        private @Nullable List<String> termsOfUses;
        public Builder() {}
        public Builder(ConditionalAccessPolicyGrantControls defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.authenticationStrengthPolicyId = defaults.authenticationStrengthPolicyId;
    	      this.builtInControls = defaults.builtInControls;
    	      this.customAuthenticationFactors = defaults.customAuthenticationFactors;
    	      this.operator = defaults.operator;
    	      this.termsOfUses = defaults.termsOfUses;
        }

        @CustomType.Setter
        public Builder authenticationStrengthPolicyId(@Nullable String authenticationStrengthPolicyId) {

            this.authenticationStrengthPolicyId = authenticationStrengthPolicyId;
            return this;
        }
        @CustomType.Setter
        public Builder builtInControls(@Nullable List<String> builtInControls) {

            this.builtInControls = builtInControls;
            return this;
        }
        public Builder builtInControls(String... builtInControls) {
            return builtInControls(List.of(builtInControls));
        }
        @CustomType.Setter
        public Builder customAuthenticationFactors(@Nullable List<String> customAuthenticationFactors) {

            this.customAuthenticationFactors = customAuthenticationFactors;
            return this;
        }
        public Builder customAuthenticationFactors(String... customAuthenticationFactors) {
            return customAuthenticationFactors(List.of(customAuthenticationFactors));
        }
        @CustomType.Setter
        public Builder operator(String operator) {
            if (operator == null) {
              throw new MissingRequiredPropertyException("ConditionalAccessPolicyGrantControls", "operator");
            }
            this.operator = operator;
            return this;
        }
        @CustomType.Setter
        public Builder termsOfUses(@Nullable List<String> termsOfUses) {

            this.termsOfUses = termsOfUses;
            return this;
        }
        public Builder termsOfUses(String... termsOfUses) {
            return termsOfUses(List.of(termsOfUses));
        }
        public ConditionalAccessPolicyGrantControls build() {
            final var _resultValue = new ConditionalAccessPolicyGrantControls();
            _resultValue.authenticationStrengthPolicyId = authenticationStrengthPolicyId;
            _resultValue.builtInControls = builtInControls;
            _resultValue.customAuthenticationFactors = customAuthenticationFactors;
            _resultValue.operator = operator;
            _resultValue.termsOfUses = termsOfUses;
            return _resultValue;
        }
    }
}

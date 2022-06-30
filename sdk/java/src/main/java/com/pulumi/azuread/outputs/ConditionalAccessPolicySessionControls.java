// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ConditionalAccessPolicySessionControls {
    /**
     * @return Whether or not application enforced restrictions are enabled. Defaults to `false`.
     * 
     */
    private final @Nullable Boolean applicationEnforcedRestrictionsEnabled;
    /**
     * @return Enables cloud app security and specifies the cloud app security policy to use. Possible values are: `blockDownloads`, `mcasConfigured`, `monitorOnly` or `unknownFutureValue`.
     * 
     */
    private final @Nullable String cloudAppSecurityPolicy;
    /**
     * @return Session control to define whether to persist cookies or not. Possible values are: `always` or `never`.
     * 
     */
    private final @Nullable String persistentBrowserMode;
    /**
     * @return Number of days or hours to enforce sign-in frequency. Required when `sign_in_frequency_period` is specified. Due to an API issue, removing this property forces a new resource to be created.
     * 
     */
    private final @Nullable Integer signInFrequency;
    /**
     * @return The time period to enforce sign-in frequency. Possible values are: `hours` or `days`. Required when `sign_in_frequency_period` is specified. Due to an API issue, removing this property forces a new resource to be created.
     * 
     */
    private final @Nullable String signInFrequencyPeriod;

    @CustomType.Constructor
    private ConditionalAccessPolicySessionControls(
        @CustomType.Parameter("applicationEnforcedRestrictionsEnabled") @Nullable Boolean applicationEnforcedRestrictionsEnabled,
        @CustomType.Parameter("cloudAppSecurityPolicy") @Nullable String cloudAppSecurityPolicy,
        @CustomType.Parameter("persistentBrowserMode") @Nullable String persistentBrowserMode,
        @CustomType.Parameter("signInFrequency") @Nullable Integer signInFrequency,
        @CustomType.Parameter("signInFrequencyPeriod") @Nullable String signInFrequencyPeriod) {
        this.applicationEnforcedRestrictionsEnabled = applicationEnforcedRestrictionsEnabled;
        this.cloudAppSecurityPolicy = cloudAppSecurityPolicy;
        this.persistentBrowserMode = persistentBrowserMode;
        this.signInFrequency = signInFrequency;
        this.signInFrequencyPeriod = signInFrequencyPeriod;
    }

    /**
     * @return Whether or not application enforced restrictions are enabled. Defaults to `false`.
     * 
     */
    public Optional<Boolean> applicationEnforcedRestrictionsEnabled() {
        return Optional.ofNullable(this.applicationEnforcedRestrictionsEnabled);
    }
    /**
     * @return Enables cloud app security and specifies the cloud app security policy to use. Possible values are: `blockDownloads`, `mcasConfigured`, `monitorOnly` or `unknownFutureValue`.
     * 
     */
    public Optional<String> cloudAppSecurityPolicy() {
        return Optional.ofNullable(this.cloudAppSecurityPolicy);
    }
    /**
     * @return Session control to define whether to persist cookies or not. Possible values are: `always` or `never`.
     * 
     */
    public Optional<String> persistentBrowserMode() {
        return Optional.ofNullable(this.persistentBrowserMode);
    }
    /**
     * @return Number of days or hours to enforce sign-in frequency. Required when `sign_in_frequency_period` is specified. Due to an API issue, removing this property forces a new resource to be created.
     * 
     */
    public Optional<Integer> signInFrequency() {
        return Optional.ofNullable(this.signInFrequency);
    }
    /**
     * @return The time period to enforce sign-in frequency. Possible values are: `hours` or `days`. Required when `sign_in_frequency_period` is specified. Due to an API issue, removing this property forces a new resource to be created.
     * 
     */
    public Optional<String> signInFrequencyPeriod() {
        return Optional.ofNullable(this.signInFrequencyPeriod);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ConditionalAccessPolicySessionControls defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable Boolean applicationEnforcedRestrictionsEnabled;
        private @Nullable String cloudAppSecurityPolicy;
        private @Nullable String persistentBrowserMode;
        private @Nullable Integer signInFrequency;
        private @Nullable String signInFrequencyPeriod;

        public Builder() {
    	      // Empty
        }

        public Builder(ConditionalAccessPolicySessionControls defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.applicationEnforcedRestrictionsEnabled = defaults.applicationEnforcedRestrictionsEnabled;
    	      this.cloudAppSecurityPolicy = defaults.cloudAppSecurityPolicy;
    	      this.persistentBrowserMode = defaults.persistentBrowserMode;
    	      this.signInFrequency = defaults.signInFrequency;
    	      this.signInFrequencyPeriod = defaults.signInFrequencyPeriod;
        }

        public Builder applicationEnforcedRestrictionsEnabled(@Nullable Boolean applicationEnforcedRestrictionsEnabled) {
            this.applicationEnforcedRestrictionsEnabled = applicationEnforcedRestrictionsEnabled;
            return this;
        }
        public Builder cloudAppSecurityPolicy(@Nullable String cloudAppSecurityPolicy) {
            this.cloudAppSecurityPolicy = cloudAppSecurityPolicy;
            return this;
        }
        public Builder persistentBrowserMode(@Nullable String persistentBrowserMode) {
            this.persistentBrowserMode = persistentBrowserMode;
            return this;
        }
        public Builder signInFrequency(@Nullable Integer signInFrequency) {
            this.signInFrequency = signInFrequency;
            return this;
        }
        public Builder signInFrequencyPeriod(@Nullable String signInFrequencyPeriod) {
            this.signInFrequencyPeriod = signInFrequencyPeriod;
            return this;
        }        public ConditionalAccessPolicySessionControls build() {
            return new ConditionalAccessPolicySessionControls(applicationEnforcedRestrictionsEnabled, cloudAppSecurityPolicy, persistentBrowserMode, signInFrequency, signInFrequencyPeriod);
        }
    }
}

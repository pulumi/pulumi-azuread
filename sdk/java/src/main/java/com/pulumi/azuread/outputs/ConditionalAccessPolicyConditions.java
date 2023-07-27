// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.azuread.outputs.ConditionalAccessPolicyConditionsApplications;
import com.pulumi.azuread.outputs.ConditionalAccessPolicyConditionsClientApplications;
import com.pulumi.azuread.outputs.ConditionalAccessPolicyConditionsDevices;
import com.pulumi.azuread.outputs.ConditionalAccessPolicyConditionsLocations;
import com.pulumi.azuread.outputs.ConditionalAccessPolicyConditionsPlatforms;
import com.pulumi.azuread.outputs.ConditionalAccessPolicyConditionsUsers;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ConditionalAccessPolicyConditions {
    /**
     * @return An `applications` block as documented below, which specifies applications and user actions included in and excluded from the policy.
     * 
     */
    private ConditionalAccessPolicyConditionsApplications applications;
    /**
     * @return A list of client application types included in the policy. Possible values are: `all`, `browser`, `mobileAppsAndDesktopClients`, `exchangeActiveSync`, `easSupported` and `other`.
     * 
     */
    private List<String> clientAppTypes;
    /**
     * @return An `client_applications` block as documented below, which specifies service principals included in and excluded from the policy.
     * 
     */
    private @Nullable ConditionalAccessPolicyConditionsClientApplications clientApplications;
    /**
     * @return A `devices` block as documented below, which describes devices to be included in and excluded from the policy. A `devices` block can be added to an existing policy, but removing the `devices` block forces a new resource to be created.
     * 
     */
    private @Nullable ConditionalAccessPolicyConditionsDevices devices;
    /**
     * @return A `locations` block as documented below, which specifies locations included in and excluded from the policy.
     * 
     */
    private @Nullable ConditionalAccessPolicyConditionsLocations locations;
    /**
     * @return A `platforms` block as documented below, which specifies platforms included in and excluded from the policy.
     * 
     */
    private @Nullable ConditionalAccessPolicyConditionsPlatforms platforms;
    /**
     * @return A list of service principal sign-in risk levels included in the policy. Possible values are: `low`, `medium`, `high`, `none`, `unknownFutureValue`.
     * 
     */
    private @Nullable List<String> servicePrincipalRiskLevels;
    /**
     * @return A list of user sign-in risk levels included in the policy. Possible values are: `low`, `medium`, `high`, `hidden`, `none`, `unknownFutureValue`.
     * 
     */
    private @Nullable List<String> signInRiskLevels;
    /**
     * @return A list of user risk levels included in the policy. Possible values are: `low`, `medium`, `high`, `hidden`, `none`, `unknownFutureValue`.
     * 
     */
    private @Nullable List<String> userRiskLevels;
    /**
     * @return A `users` block as documented below, which specifies users, groups, and roles included in and excluded from the policy.
     * 
     */
    private ConditionalAccessPolicyConditionsUsers users;

    private ConditionalAccessPolicyConditions() {}
    /**
     * @return An `applications` block as documented below, which specifies applications and user actions included in and excluded from the policy.
     * 
     */
    public ConditionalAccessPolicyConditionsApplications applications() {
        return this.applications;
    }
    /**
     * @return A list of client application types included in the policy. Possible values are: `all`, `browser`, `mobileAppsAndDesktopClients`, `exchangeActiveSync`, `easSupported` and `other`.
     * 
     */
    public List<String> clientAppTypes() {
        return this.clientAppTypes;
    }
    /**
     * @return An `client_applications` block as documented below, which specifies service principals included in and excluded from the policy.
     * 
     */
    public Optional<ConditionalAccessPolicyConditionsClientApplications> clientApplications() {
        return Optional.ofNullable(this.clientApplications);
    }
    /**
     * @return A `devices` block as documented below, which describes devices to be included in and excluded from the policy. A `devices` block can be added to an existing policy, but removing the `devices` block forces a new resource to be created.
     * 
     */
    public Optional<ConditionalAccessPolicyConditionsDevices> devices() {
        return Optional.ofNullable(this.devices);
    }
    /**
     * @return A `locations` block as documented below, which specifies locations included in and excluded from the policy.
     * 
     */
    public Optional<ConditionalAccessPolicyConditionsLocations> locations() {
        return Optional.ofNullable(this.locations);
    }
    /**
     * @return A `platforms` block as documented below, which specifies platforms included in and excluded from the policy.
     * 
     */
    public Optional<ConditionalAccessPolicyConditionsPlatforms> platforms() {
        return Optional.ofNullable(this.platforms);
    }
    /**
     * @return A list of service principal sign-in risk levels included in the policy. Possible values are: `low`, `medium`, `high`, `none`, `unknownFutureValue`.
     * 
     */
    public List<String> servicePrincipalRiskLevels() {
        return this.servicePrincipalRiskLevels == null ? List.of() : this.servicePrincipalRiskLevels;
    }
    /**
     * @return A list of user sign-in risk levels included in the policy. Possible values are: `low`, `medium`, `high`, `hidden`, `none`, `unknownFutureValue`.
     * 
     */
    public List<String> signInRiskLevels() {
        return this.signInRiskLevels == null ? List.of() : this.signInRiskLevels;
    }
    /**
     * @return A list of user risk levels included in the policy. Possible values are: `low`, `medium`, `high`, `hidden`, `none`, `unknownFutureValue`.
     * 
     */
    public List<String> userRiskLevels() {
        return this.userRiskLevels == null ? List.of() : this.userRiskLevels;
    }
    /**
     * @return A `users` block as documented below, which specifies users, groups, and roles included in and excluded from the policy.
     * 
     */
    public ConditionalAccessPolicyConditionsUsers users() {
        return this.users;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ConditionalAccessPolicyConditions defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private ConditionalAccessPolicyConditionsApplications applications;
        private List<String> clientAppTypes;
        private @Nullable ConditionalAccessPolicyConditionsClientApplications clientApplications;
        private @Nullable ConditionalAccessPolicyConditionsDevices devices;
        private @Nullable ConditionalAccessPolicyConditionsLocations locations;
        private @Nullable ConditionalAccessPolicyConditionsPlatforms platforms;
        private @Nullable List<String> servicePrincipalRiskLevels;
        private @Nullable List<String> signInRiskLevels;
        private @Nullable List<String> userRiskLevels;
        private ConditionalAccessPolicyConditionsUsers users;
        public Builder() {}
        public Builder(ConditionalAccessPolicyConditions defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.applications = defaults.applications;
    	      this.clientAppTypes = defaults.clientAppTypes;
    	      this.clientApplications = defaults.clientApplications;
    	      this.devices = defaults.devices;
    	      this.locations = defaults.locations;
    	      this.platforms = defaults.platforms;
    	      this.servicePrincipalRiskLevels = defaults.servicePrincipalRiskLevels;
    	      this.signInRiskLevels = defaults.signInRiskLevels;
    	      this.userRiskLevels = defaults.userRiskLevels;
    	      this.users = defaults.users;
        }

        @CustomType.Setter
        public Builder applications(ConditionalAccessPolicyConditionsApplications applications) {
            this.applications = Objects.requireNonNull(applications);
            return this;
        }
        @CustomType.Setter
        public Builder clientAppTypes(List<String> clientAppTypes) {
            this.clientAppTypes = Objects.requireNonNull(clientAppTypes);
            return this;
        }
        public Builder clientAppTypes(String... clientAppTypes) {
            return clientAppTypes(List.of(clientAppTypes));
        }
        @CustomType.Setter
        public Builder clientApplications(@Nullable ConditionalAccessPolicyConditionsClientApplications clientApplications) {
            this.clientApplications = clientApplications;
            return this;
        }
        @CustomType.Setter
        public Builder devices(@Nullable ConditionalAccessPolicyConditionsDevices devices) {
            this.devices = devices;
            return this;
        }
        @CustomType.Setter
        public Builder locations(@Nullable ConditionalAccessPolicyConditionsLocations locations) {
            this.locations = locations;
            return this;
        }
        @CustomType.Setter
        public Builder platforms(@Nullable ConditionalAccessPolicyConditionsPlatforms platforms) {
            this.platforms = platforms;
            return this;
        }
        @CustomType.Setter
        public Builder servicePrincipalRiskLevels(@Nullable List<String> servicePrincipalRiskLevels) {
            this.servicePrincipalRiskLevels = servicePrincipalRiskLevels;
            return this;
        }
        public Builder servicePrincipalRiskLevels(String... servicePrincipalRiskLevels) {
            return servicePrincipalRiskLevels(List.of(servicePrincipalRiskLevels));
        }
        @CustomType.Setter
        public Builder signInRiskLevels(@Nullable List<String> signInRiskLevels) {
            this.signInRiskLevels = signInRiskLevels;
            return this;
        }
        public Builder signInRiskLevels(String... signInRiskLevels) {
            return signInRiskLevels(List.of(signInRiskLevels));
        }
        @CustomType.Setter
        public Builder userRiskLevels(@Nullable List<String> userRiskLevels) {
            this.userRiskLevels = userRiskLevels;
            return this;
        }
        public Builder userRiskLevels(String... userRiskLevels) {
            return userRiskLevels(List.of(userRiskLevels));
        }
        @CustomType.Setter
        public Builder users(ConditionalAccessPolicyConditionsUsers users) {
            this.users = Objects.requireNonNull(users);
            return this;
        }
        public ConditionalAccessPolicyConditions build() {
            final var o = new ConditionalAccessPolicyConditions();
            o.applications = applications;
            o.clientAppTypes = clientAppTypes;
            o.clientApplications = clientApplications;
            o.devices = devices;
            o.locations = locations;
            o.platforms = platforms;
            o.servicePrincipalRiskLevels = servicePrincipalRiskLevels;
            o.signInRiskLevels = signInRiskLevels;
            o.userRiskLevels = userRiskLevels;
            o.users = users;
            return o;
        }
    }
}

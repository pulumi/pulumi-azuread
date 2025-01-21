// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class NamedLocationCountry {
    /**
     * @return List of countries and/or regions in two-letter format specified by ISO 3166-2.
     * 
     */
    private List<String> countriesAndRegions;
    /**
     * @return Method of detecting country the user is located in. Possible values are `clientIpAddress` for IP-based location and `authenticatorAppGps` for Authenticator app GPS-based location.  Defaults to `clientIpAddress`.
     * 
     */
    private @Nullable String countryLookupMethod;
    /**
     * @return Whether IP addresses that don&#39;t map to a country or region should be included in the named location. Defaults to `false`.
     * 
     */
    private @Nullable Boolean includeUnknownCountriesAndRegions;

    private NamedLocationCountry() {}
    /**
     * @return List of countries and/or regions in two-letter format specified by ISO 3166-2.
     * 
     */
    public List<String> countriesAndRegions() {
        return this.countriesAndRegions;
    }
    /**
     * @return Method of detecting country the user is located in. Possible values are `clientIpAddress` for IP-based location and `authenticatorAppGps` for Authenticator app GPS-based location.  Defaults to `clientIpAddress`.
     * 
     */
    public Optional<String> countryLookupMethod() {
        return Optional.ofNullable(this.countryLookupMethod);
    }
    /**
     * @return Whether IP addresses that don&#39;t map to a country or region should be included in the named location. Defaults to `false`.
     * 
     */
    public Optional<Boolean> includeUnknownCountriesAndRegions() {
        return Optional.ofNullable(this.includeUnknownCountriesAndRegions);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(NamedLocationCountry defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> countriesAndRegions;
        private @Nullable String countryLookupMethod;
        private @Nullable Boolean includeUnknownCountriesAndRegions;
        public Builder() {}
        public Builder(NamedLocationCountry defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.countriesAndRegions = defaults.countriesAndRegions;
    	      this.countryLookupMethod = defaults.countryLookupMethod;
    	      this.includeUnknownCountriesAndRegions = defaults.includeUnknownCountriesAndRegions;
        }

        @CustomType.Setter
        public Builder countriesAndRegions(List<String> countriesAndRegions) {
            if (countriesAndRegions == null) {
              throw new MissingRequiredPropertyException("NamedLocationCountry", "countriesAndRegions");
            }
            this.countriesAndRegions = countriesAndRegions;
            return this;
        }
        public Builder countriesAndRegions(String... countriesAndRegions) {
            return countriesAndRegions(List.of(countriesAndRegions));
        }
        @CustomType.Setter
        public Builder countryLookupMethod(@Nullable String countryLookupMethod) {

            this.countryLookupMethod = countryLookupMethod;
            return this;
        }
        @CustomType.Setter
        public Builder includeUnknownCountriesAndRegions(@Nullable Boolean includeUnknownCountriesAndRegions) {

            this.includeUnknownCountriesAndRegions = includeUnknownCountriesAndRegions;
            return this;
        }
        public NamedLocationCountry build() {
            final var _resultValue = new NamedLocationCountry();
            _resultValue.countriesAndRegions = countriesAndRegions;
            _resultValue.countryLookupMethod = countryLookupMethod;
            _resultValue.includeUnknownCountriesAndRegions = includeUnknownCountriesAndRegions;
            return _resultValue;
        }
    }
}

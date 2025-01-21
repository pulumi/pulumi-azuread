// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NamedLocationCountryArgs extends com.pulumi.resources.ResourceArgs {

    public static final NamedLocationCountryArgs Empty = new NamedLocationCountryArgs();

    /**
     * List of countries and/or regions in two-letter format specified by ISO 3166-2.
     * 
     */
    @Import(name="countriesAndRegions", required=true)
    private Output<List<String>> countriesAndRegions;

    /**
     * @return List of countries and/or regions in two-letter format specified by ISO 3166-2.
     * 
     */
    public Output<List<String>> countriesAndRegions() {
        return this.countriesAndRegions;
    }

    /**
     * Method of detecting country the user is located in. Possible values are `clientIpAddress` for IP-based location and `authenticatorAppGps` for Authenticator app GPS-based location.  Defaults to `clientIpAddress`.
     * 
     */
    @Import(name="countryLookupMethod")
    private @Nullable Output<String> countryLookupMethod;

    /**
     * @return Method of detecting country the user is located in. Possible values are `clientIpAddress` for IP-based location and `authenticatorAppGps` for Authenticator app GPS-based location.  Defaults to `clientIpAddress`.
     * 
     */
    public Optional<Output<String>> countryLookupMethod() {
        return Optional.ofNullable(this.countryLookupMethod);
    }

    /**
     * Whether IP addresses that don&#39;t map to a country or region should be included in the named location. Defaults to `false`.
     * 
     */
    @Import(name="includeUnknownCountriesAndRegions")
    private @Nullable Output<Boolean> includeUnknownCountriesAndRegions;

    /**
     * @return Whether IP addresses that don&#39;t map to a country or region should be included in the named location. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> includeUnknownCountriesAndRegions() {
        return Optional.ofNullable(this.includeUnknownCountriesAndRegions);
    }

    private NamedLocationCountryArgs() {}

    private NamedLocationCountryArgs(NamedLocationCountryArgs $) {
        this.countriesAndRegions = $.countriesAndRegions;
        this.countryLookupMethod = $.countryLookupMethod;
        this.includeUnknownCountriesAndRegions = $.includeUnknownCountriesAndRegions;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NamedLocationCountryArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NamedLocationCountryArgs $;

        public Builder() {
            $ = new NamedLocationCountryArgs();
        }

        public Builder(NamedLocationCountryArgs defaults) {
            $ = new NamedLocationCountryArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param countriesAndRegions List of countries and/or regions in two-letter format specified by ISO 3166-2.
         * 
         * @return builder
         * 
         */
        public Builder countriesAndRegions(Output<List<String>> countriesAndRegions) {
            $.countriesAndRegions = countriesAndRegions;
            return this;
        }

        /**
         * @param countriesAndRegions List of countries and/or regions in two-letter format specified by ISO 3166-2.
         * 
         * @return builder
         * 
         */
        public Builder countriesAndRegions(List<String> countriesAndRegions) {
            return countriesAndRegions(Output.of(countriesAndRegions));
        }

        /**
         * @param countriesAndRegions List of countries and/or regions in two-letter format specified by ISO 3166-2.
         * 
         * @return builder
         * 
         */
        public Builder countriesAndRegions(String... countriesAndRegions) {
            return countriesAndRegions(List.of(countriesAndRegions));
        }

        /**
         * @param countryLookupMethod Method of detecting country the user is located in. Possible values are `clientIpAddress` for IP-based location and `authenticatorAppGps` for Authenticator app GPS-based location.  Defaults to `clientIpAddress`.
         * 
         * @return builder
         * 
         */
        public Builder countryLookupMethod(@Nullable Output<String> countryLookupMethod) {
            $.countryLookupMethod = countryLookupMethod;
            return this;
        }

        /**
         * @param countryLookupMethod Method of detecting country the user is located in. Possible values are `clientIpAddress` for IP-based location and `authenticatorAppGps` for Authenticator app GPS-based location.  Defaults to `clientIpAddress`.
         * 
         * @return builder
         * 
         */
        public Builder countryLookupMethod(String countryLookupMethod) {
            return countryLookupMethod(Output.of(countryLookupMethod));
        }

        /**
         * @param includeUnknownCountriesAndRegions Whether IP addresses that don&#39;t map to a country or region should be included in the named location. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder includeUnknownCountriesAndRegions(@Nullable Output<Boolean> includeUnknownCountriesAndRegions) {
            $.includeUnknownCountriesAndRegions = includeUnknownCountriesAndRegions;
            return this;
        }

        /**
         * @param includeUnknownCountriesAndRegions Whether IP addresses that don&#39;t map to a country or region should be included in the named location. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder includeUnknownCountriesAndRegions(Boolean includeUnknownCountriesAndRegions) {
            return includeUnknownCountriesAndRegions(Output.of(includeUnknownCountriesAndRegions));
        }

        public NamedLocationCountryArgs build() {
            if ($.countriesAndRegions == null) {
                throw new MissingRequiredPropertyException("NamedLocationCountryArgs", "countriesAndRegions");
            }
            return $;
        }
    }

}

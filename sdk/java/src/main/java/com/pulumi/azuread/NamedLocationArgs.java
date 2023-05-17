// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.inputs.NamedLocationCountryArgs;
import com.pulumi.azuread.inputs.NamedLocationIpArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NamedLocationArgs extends com.pulumi.resources.ResourceArgs {

    public static final NamedLocationArgs Empty = new NamedLocationArgs();

    /**
     * A `country` block as documented below, which configures a country-based named location.
     * 
     */
    @Import(name="country")
    private @Nullable Output<NamedLocationCountryArgs> country;

    /**
     * @return A `country` block as documented below, which configures a country-based named location.
     * 
     */
    public Optional<Output<NamedLocationCountryArgs>> country() {
        return Optional.ofNullable(this.country);
    }

    /**
     * The friendly name for this named location.
     * 
     */
    @Import(name="displayName", required=true)
    private Output<String> displayName;

    /**
     * @return The friendly name for this named location.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }

    /**
     * An `ip` block as documented below, which configures an IP-based named location.
     * 
     * &gt; Exactly one of `ip` or `country` must be specified. Changing between these forces a new resource to be created.
     * 
     */
    @Import(name="ip")
    private @Nullable Output<NamedLocationIpArgs> ip;

    /**
     * @return An `ip` block as documented below, which configures an IP-based named location.
     * 
     * &gt; Exactly one of `ip` or `country` must be specified. Changing between these forces a new resource to be created.
     * 
     */
    public Optional<Output<NamedLocationIpArgs>> ip() {
        return Optional.ofNullable(this.ip);
    }

    private NamedLocationArgs() {}

    private NamedLocationArgs(NamedLocationArgs $) {
        this.country = $.country;
        this.displayName = $.displayName;
        this.ip = $.ip;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NamedLocationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NamedLocationArgs $;

        public Builder() {
            $ = new NamedLocationArgs();
        }

        public Builder(NamedLocationArgs defaults) {
            $ = new NamedLocationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param country A `country` block as documented below, which configures a country-based named location.
         * 
         * @return builder
         * 
         */
        public Builder country(@Nullable Output<NamedLocationCountryArgs> country) {
            $.country = country;
            return this;
        }

        /**
         * @param country A `country` block as documented below, which configures a country-based named location.
         * 
         * @return builder
         * 
         */
        public Builder country(NamedLocationCountryArgs country) {
            return country(Output.of(country));
        }

        /**
         * @param displayName The friendly name for this named location.
         * 
         * @return builder
         * 
         */
        public Builder displayName(Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName The friendly name for this named location.
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param ip An `ip` block as documented below, which configures an IP-based named location.
         * 
         * &gt; Exactly one of `ip` or `country` must be specified. Changing between these forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder ip(@Nullable Output<NamedLocationIpArgs> ip) {
            $.ip = ip;
            return this;
        }

        /**
         * @param ip An `ip` block as documented below, which configures an IP-based named location.
         * 
         * &gt; Exactly one of `ip` or `country` must be specified. Changing between these forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder ip(NamedLocationIpArgs ip) {
            return ip(Output.of(ip));
        }

        public NamedLocationArgs build() {
            $.displayName = Objects.requireNonNull($.displayName, "expected parameter 'displayName' to be non-null");
            return $;
        }
    }

}

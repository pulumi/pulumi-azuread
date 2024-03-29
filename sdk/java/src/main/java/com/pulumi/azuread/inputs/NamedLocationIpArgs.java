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


public final class NamedLocationIpArgs extends com.pulumi.resources.ResourceArgs {

    public static final NamedLocationIpArgs Empty = new NamedLocationIpArgs();

    /**
     * List of IP address ranges in IPv4 CIDR format (e.g. `1.2.3.4/32`) or any allowable IPv6 format from IETF RFC596. Each CIDR prefix must be `/8` or larger.
     * 
     */
    @Import(name="ipRanges", required=true)
    private Output<List<String>> ipRanges;

    /**
     * @return List of IP address ranges in IPv4 CIDR format (e.g. `1.2.3.4/32`) or any allowable IPv6 format from IETF RFC596. Each CIDR prefix must be `/8` or larger.
     * 
     */
    public Output<List<String>> ipRanges() {
        return this.ipRanges;
    }

    /**
     * Whether the named location is trusted. Defaults to `false`.
     * 
     */
    @Import(name="trusted")
    private @Nullable Output<Boolean> trusted;

    /**
     * @return Whether the named location is trusted. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> trusted() {
        return Optional.ofNullable(this.trusted);
    }

    private NamedLocationIpArgs() {}

    private NamedLocationIpArgs(NamedLocationIpArgs $) {
        this.ipRanges = $.ipRanges;
        this.trusted = $.trusted;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NamedLocationIpArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NamedLocationIpArgs $;

        public Builder() {
            $ = new NamedLocationIpArgs();
        }

        public Builder(NamedLocationIpArgs defaults) {
            $ = new NamedLocationIpArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ipRanges List of IP address ranges in IPv4 CIDR format (e.g. `1.2.3.4/32`) or any allowable IPv6 format from IETF RFC596. Each CIDR prefix must be `/8` or larger.
         * 
         * @return builder
         * 
         */
        public Builder ipRanges(Output<List<String>> ipRanges) {
            $.ipRanges = ipRanges;
            return this;
        }

        /**
         * @param ipRanges List of IP address ranges in IPv4 CIDR format (e.g. `1.2.3.4/32`) or any allowable IPv6 format from IETF RFC596. Each CIDR prefix must be `/8` or larger.
         * 
         * @return builder
         * 
         */
        public Builder ipRanges(List<String> ipRanges) {
            return ipRanges(Output.of(ipRanges));
        }

        /**
         * @param ipRanges List of IP address ranges in IPv4 CIDR format (e.g. `1.2.3.4/32`) or any allowable IPv6 format from IETF RFC596. Each CIDR prefix must be `/8` or larger.
         * 
         * @return builder
         * 
         */
        public Builder ipRanges(String... ipRanges) {
            return ipRanges(List.of(ipRanges));
        }

        /**
         * @param trusted Whether the named location is trusted. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder trusted(@Nullable Output<Boolean> trusted) {
            $.trusted = trusted;
            return this;
        }

        /**
         * @param trusted Whether the named location is trusted. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder trusted(Boolean trusted) {
            return trusted(Output.of(trusted));
        }

        public NamedLocationIpArgs build() {
            if ($.ipRanges == null) {
                throw new MissingRequiredPropertyException("NamedLocationIpArgs", "ipRanges");
            }
            return $;
        }
    }

}

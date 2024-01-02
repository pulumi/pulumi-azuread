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


public final class ApplicationOptionalClaimsAccessTokenArgs extends com.pulumi.resources.ResourceArgs {

    public static final ApplicationOptionalClaimsAccessTokenArgs Empty = new ApplicationOptionalClaimsAccessTokenArgs();

    /**
     * List of additional properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim. Possible values are: `cloud_displayname`, `dns_domain_and_sam_account_name`, `emit_as_roles`, `include_externally_authenticated_upn_without_hash`, `include_externally_authenticated_upn`, `max_size_limit`, `netbios_domain_and_sam_account_name`, `on_premise_security_identifier`, `sam_account_name`, and `use_guid`.
     * 
     */
    @Import(name="additionalProperties")
    private @Nullable Output<List<String>> additionalProperties;

    /**
     * @return List of additional properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim. Possible values are: `cloud_displayname`, `dns_domain_and_sam_account_name`, `emit_as_roles`, `include_externally_authenticated_upn_without_hash`, `include_externally_authenticated_upn`, `max_size_limit`, `netbios_domain_and_sam_account_name`, `on_premise_security_identifier`, `sam_account_name`, and `use_guid`.
     * 
     */
    public Optional<Output<List<String>>> additionalProperties() {
        return Optional.ofNullable(this.additionalProperties);
    }

    /**
     * Whether the claim specified by the client is necessary to ensure a smooth authorization experience.
     * 
     */
    @Import(name="essential")
    private @Nullable Output<Boolean> essential;

    /**
     * @return Whether the claim specified by the client is necessary to ensure a smooth authorization experience.
     * 
     */
    public Optional<Output<Boolean>> essential() {
        return Optional.ofNullable(this.essential);
    }

    /**
     * The name of the optional claim.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return The name of the optional claim.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * The source of the claim. If `source` is absent, the claim is a predefined optional claim. If `source` is `user`, the value of `name` is the extension property from the user object.
     * 
     */
    @Import(name="source")
    private @Nullable Output<String> source;

    /**
     * @return The source of the claim. If `source` is absent, the claim is a predefined optional claim. If `source` is `user`, the value of `name` is the extension property from the user object.
     * 
     */
    public Optional<Output<String>> source() {
        return Optional.ofNullable(this.source);
    }

    private ApplicationOptionalClaimsAccessTokenArgs() {}

    private ApplicationOptionalClaimsAccessTokenArgs(ApplicationOptionalClaimsAccessTokenArgs $) {
        this.additionalProperties = $.additionalProperties;
        this.essential = $.essential;
        this.name = $.name;
        this.source = $.source;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ApplicationOptionalClaimsAccessTokenArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ApplicationOptionalClaimsAccessTokenArgs $;

        public Builder() {
            $ = new ApplicationOptionalClaimsAccessTokenArgs();
        }

        public Builder(ApplicationOptionalClaimsAccessTokenArgs defaults) {
            $ = new ApplicationOptionalClaimsAccessTokenArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param additionalProperties List of additional properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim. Possible values are: `cloud_displayname`, `dns_domain_and_sam_account_name`, `emit_as_roles`, `include_externally_authenticated_upn_without_hash`, `include_externally_authenticated_upn`, `max_size_limit`, `netbios_domain_and_sam_account_name`, `on_premise_security_identifier`, `sam_account_name`, and `use_guid`.
         * 
         * @return builder
         * 
         */
        public Builder additionalProperties(@Nullable Output<List<String>> additionalProperties) {
            $.additionalProperties = additionalProperties;
            return this;
        }

        /**
         * @param additionalProperties List of additional properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim. Possible values are: `cloud_displayname`, `dns_domain_and_sam_account_name`, `emit_as_roles`, `include_externally_authenticated_upn_without_hash`, `include_externally_authenticated_upn`, `max_size_limit`, `netbios_domain_and_sam_account_name`, `on_premise_security_identifier`, `sam_account_name`, and `use_guid`.
         * 
         * @return builder
         * 
         */
        public Builder additionalProperties(List<String> additionalProperties) {
            return additionalProperties(Output.of(additionalProperties));
        }

        /**
         * @param additionalProperties List of additional properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim. Possible values are: `cloud_displayname`, `dns_domain_and_sam_account_name`, `emit_as_roles`, `include_externally_authenticated_upn_without_hash`, `include_externally_authenticated_upn`, `max_size_limit`, `netbios_domain_and_sam_account_name`, `on_premise_security_identifier`, `sam_account_name`, and `use_guid`.
         * 
         * @return builder
         * 
         */
        public Builder additionalProperties(String... additionalProperties) {
            return additionalProperties(List.of(additionalProperties));
        }

        /**
         * @param essential Whether the claim specified by the client is necessary to ensure a smooth authorization experience.
         * 
         * @return builder
         * 
         */
        public Builder essential(@Nullable Output<Boolean> essential) {
            $.essential = essential;
            return this;
        }

        /**
         * @param essential Whether the claim specified by the client is necessary to ensure a smooth authorization experience.
         * 
         * @return builder
         * 
         */
        public Builder essential(Boolean essential) {
            return essential(Output.of(essential));
        }

        /**
         * @param name The name of the optional claim.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the optional claim.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param source The source of the claim. If `source` is absent, the claim is a predefined optional claim. If `source` is `user`, the value of `name` is the extension property from the user object.
         * 
         * @return builder
         * 
         */
        public Builder source(@Nullable Output<String> source) {
            $.source = source;
            return this;
        }

        /**
         * @param source The source of the claim. If `source` is absent, the claim is a predefined optional claim. If `source` is `user`, the value of `name` is the extension property from the user object.
         * 
         * @return builder
         * 
         */
        public Builder source(String source) {
            return source(Output.of(source));
        }

        public ApplicationOptionalClaimsAccessTokenArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("ApplicationOptionalClaimsAccessTokenArgs", "name");
            }
            return $;
        }
    }

}

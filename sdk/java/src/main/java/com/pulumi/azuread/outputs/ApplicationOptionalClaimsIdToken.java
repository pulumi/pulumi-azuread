// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ApplicationOptionalClaimsIdToken {
    /**
     * @return List of additional properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim. Possible values are: `cloud_displayname`, `dns_domain_and_sam_account_name`, `emit_as_roles`, `include_externally_authenticated_upn_without_hash`, `include_externally_authenticated_upn`, `max_size_limit`, `netbios_domain_and_sam_account_name`, `on_premise_security_identifier`, `sam_account_name`, and `use_guid`.
     * 
     */
    private @Nullable List<String> additionalProperties;
    /**
     * @return Whether the claim specified by the client is necessary to ensure a smooth authorization experience.
     * 
     */
    private @Nullable Boolean essential;
    /**
     * @return The name of the optional claim.
     * 
     */
    private String name;
    /**
     * @return The source of the claim. If `source` is absent, the claim is a predefined optional claim. If `source` is `user`, the value of `name` is the extension property from the user object.
     * 
     */
    private @Nullable String source;

    private ApplicationOptionalClaimsIdToken() {}
    /**
     * @return List of additional properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim. Possible values are: `cloud_displayname`, `dns_domain_and_sam_account_name`, `emit_as_roles`, `include_externally_authenticated_upn_without_hash`, `include_externally_authenticated_upn`, `max_size_limit`, `netbios_domain_and_sam_account_name`, `on_premise_security_identifier`, `sam_account_name`, and `use_guid`.
     * 
     */
    public List<String> additionalProperties() {
        return this.additionalProperties == null ? List.of() : this.additionalProperties;
    }
    /**
     * @return Whether the claim specified by the client is necessary to ensure a smooth authorization experience.
     * 
     */
    public Optional<Boolean> essential() {
        return Optional.ofNullable(this.essential);
    }
    /**
     * @return The name of the optional claim.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return The source of the claim. If `source` is absent, the claim is a predefined optional claim. If `source` is `user`, the value of `name` is the extension property from the user object.
     * 
     */
    public Optional<String> source() {
        return Optional.ofNullable(this.source);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ApplicationOptionalClaimsIdToken defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> additionalProperties;
        private @Nullable Boolean essential;
        private String name;
        private @Nullable String source;
        public Builder() {}
        public Builder(ApplicationOptionalClaimsIdToken defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.additionalProperties = defaults.additionalProperties;
    	      this.essential = defaults.essential;
    	      this.name = defaults.name;
    	      this.source = defaults.source;
        }

        @CustomType.Setter
        public Builder additionalProperties(@Nullable List<String> additionalProperties) {
            this.additionalProperties = additionalProperties;
            return this;
        }
        public Builder additionalProperties(String... additionalProperties) {
            return additionalProperties(List.of(additionalProperties));
        }
        @CustomType.Setter
        public Builder essential(@Nullable Boolean essential) {
            this.essential = essential;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder source(@Nullable String source) {
            this.source = source;
            return this;
        }
        public ApplicationOptionalClaimsIdToken build() {
            final var o = new ApplicationOptionalClaimsIdToken();
            o.additionalProperties = additionalProperties;
            o.essential = essential;
            o.name = name;
            o.source = source;
            return o;
        }
    }
}

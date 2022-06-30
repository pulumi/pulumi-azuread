// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.azuread.outputs.ApplicationOptionalClaimsAccessToken;
import com.pulumi.azuread.outputs.ApplicationOptionalClaimsIdToken;
import com.pulumi.azuread.outputs.ApplicationOptionalClaimsSaml2Token;
import com.pulumi.core.annotations.CustomType;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class ApplicationOptionalClaims {
    /**
     * @return One or more `access_token` blocks as documented below.
     * 
     */
    private final @Nullable List<ApplicationOptionalClaimsAccessToken> accessTokens;
    /**
     * @return One or more `id_token` blocks as documented below.
     * 
     */
    private final @Nullable List<ApplicationOptionalClaimsIdToken> idTokens;
    /**
     * @return One or more `saml2_token` blocks as documented below.
     * 
     */
    private final @Nullable List<ApplicationOptionalClaimsSaml2Token> saml2Tokens;

    @CustomType.Constructor
    private ApplicationOptionalClaims(
        @CustomType.Parameter("accessTokens") @Nullable List<ApplicationOptionalClaimsAccessToken> accessTokens,
        @CustomType.Parameter("idTokens") @Nullable List<ApplicationOptionalClaimsIdToken> idTokens,
        @CustomType.Parameter("saml2Tokens") @Nullable List<ApplicationOptionalClaimsSaml2Token> saml2Tokens) {
        this.accessTokens = accessTokens;
        this.idTokens = idTokens;
        this.saml2Tokens = saml2Tokens;
    }

    /**
     * @return One or more `access_token` blocks as documented below.
     * 
     */
    public List<ApplicationOptionalClaimsAccessToken> accessTokens() {
        return this.accessTokens == null ? List.of() : this.accessTokens;
    }
    /**
     * @return One or more `id_token` blocks as documented below.
     * 
     */
    public List<ApplicationOptionalClaimsIdToken> idTokens() {
        return this.idTokens == null ? List.of() : this.idTokens;
    }
    /**
     * @return One or more `saml2_token` blocks as documented below.
     * 
     */
    public List<ApplicationOptionalClaimsSaml2Token> saml2Tokens() {
        return this.saml2Tokens == null ? List.of() : this.saml2Tokens;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ApplicationOptionalClaims defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable List<ApplicationOptionalClaimsAccessToken> accessTokens;
        private @Nullable List<ApplicationOptionalClaimsIdToken> idTokens;
        private @Nullable List<ApplicationOptionalClaimsSaml2Token> saml2Tokens;

        public Builder() {
    	      // Empty
        }

        public Builder(ApplicationOptionalClaims defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accessTokens = defaults.accessTokens;
    	      this.idTokens = defaults.idTokens;
    	      this.saml2Tokens = defaults.saml2Tokens;
        }

        public Builder accessTokens(@Nullable List<ApplicationOptionalClaimsAccessToken> accessTokens) {
            this.accessTokens = accessTokens;
            return this;
        }
        public Builder accessTokens(ApplicationOptionalClaimsAccessToken... accessTokens) {
            return accessTokens(List.of(accessTokens));
        }
        public Builder idTokens(@Nullable List<ApplicationOptionalClaimsIdToken> idTokens) {
            this.idTokens = idTokens;
            return this;
        }
        public Builder idTokens(ApplicationOptionalClaimsIdToken... idTokens) {
            return idTokens(List.of(idTokens));
        }
        public Builder saml2Tokens(@Nullable List<ApplicationOptionalClaimsSaml2Token> saml2Tokens) {
            this.saml2Tokens = saml2Tokens;
            return this;
        }
        public Builder saml2Tokens(ApplicationOptionalClaimsSaml2Token... saml2Tokens) {
            return saml2Tokens(List.of(saml2Tokens));
        }        public ApplicationOptionalClaims build() {
            return new ApplicationOptionalClaims(accessTokens, idTokens, saml2Tokens);
        }
    }
}

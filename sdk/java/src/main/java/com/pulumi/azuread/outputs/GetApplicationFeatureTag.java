// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetApplicationFeatureTag {
    /**
     * @return Whether this application represents a custom SAML application for linked service principals.
     * 
     */
    private @Nullable Boolean customSingleSignOn;
    /**
     * @return Whether this application represents an Enterprise Application for linked service principals.
     * 
     */
    private @Nullable Boolean enterprise;
    /**
     * @return Whether this application represents a gallery application for linked service principals.
     * 
     */
    private @Nullable Boolean gallery;
    /**
     * @return Whether this app is visible to users in My Apps and Office 365 Launcher.
     * 
     */
    private @Nullable Boolean hide;

    private GetApplicationFeatureTag() {}
    /**
     * @return Whether this application represents a custom SAML application for linked service principals.
     * 
     */
    public Optional<Boolean> customSingleSignOn() {
        return Optional.ofNullable(this.customSingleSignOn);
    }
    /**
     * @return Whether this application represents an Enterprise Application for linked service principals.
     * 
     */
    public Optional<Boolean> enterprise() {
        return Optional.ofNullable(this.enterprise);
    }
    /**
     * @return Whether this application represents a gallery application for linked service principals.
     * 
     */
    public Optional<Boolean> gallery() {
        return Optional.ofNullable(this.gallery);
    }
    /**
     * @return Whether this app is visible to users in My Apps and Office 365 Launcher.
     * 
     */
    public Optional<Boolean> hide() {
        return Optional.ofNullable(this.hide);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetApplicationFeatureTag defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean customSingleSignOn;
        private @Nullable Boolean enterprise;
        private @Nullable Boolean gallery;
        private @Nullable Boolean hide;
        public Builder() {}
        public Builder(GetApplicationFeatureTag defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.customSingleSignOn = defaults.customSingleSignOn;
    	      this.enterprise = defaults.enterprise;
    	      this.gallery = defaults.gallery;
    	      this.hide = defaults.hide;
        }

        @CustomType.Setter
        public Builder customSingleSignOn(@Nullable Boolean customSingleSignOn) {
            this.customSingleSignOn = customSingleSignOn;
            return this;
        }
        @CustomType.Setter
        public Builder enterprise(@Nullable Boolean enterprise) {
            this.enterprise = enterprise;
            return this;
        }
        @CustomType.Setter
        public Builder gallery(@Nullable Boolean gallery) {
            this.gallery = gallery;
            return this;
        }
        @CustomType.Setter
        public Builder hide(@Nullable Boolean hide) {
            this.hide = hide;
            return this;
        }
        public GetApplicationFeatureTag build() {
            final var _resultValue = new GetApplicationFeatureTag();
            _resultValue.customSingleSignOn = customSingleSignOn;
            _resultValue.enterprise = enterprise;
            _resultValue.gallery = gallery;
            _resultValue.hide = hide;
            return _resultValue;
        }
    }
}

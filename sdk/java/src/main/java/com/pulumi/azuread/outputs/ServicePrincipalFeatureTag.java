// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ServicePrincipalFeatureTag {
    /**
     * @return Whether this service principal represents a custom SAML application. Enabling this will assign the `WindowsAzureActiveDirectoryCustomSingleSignOnApplication` tag. Defaults to `false`.
     * 
     */
    private @Nullable Boolean customSingleSignOn;
    /**
     * @return Whether this service principal represents an Enterprise Application. Enabling this will assign the `WindowsAzureActiveDirectoryIntegratedApp` tag. Defaults to `false`.
     * 
     */
    private @Nullable Boolean enterprise;
    /**
     * @return Whether this service principal represents a gallery application. Enabling this will assign the `WindowsAzureActiveDirectoryGalleryApplicationNonPrimaryV1` tag. Defaults to `false`.
     * 
     */
    private @Nullable Boolean gallery;
    /**
     * @return Whether this app is invisible to users in My Apps and Office 365 Launcher. Enabling this will assign the `HideApp` tag. Defaults to `false`.
     * 
     */
    private @Nullable Boolean hide;

    private ServicePrincipalFeatureTag() {}
    /**
     * @return Whether this service principal represents a custom SAML application. Enabling this will assign the `WindowsAzureActiveDirectoryCustomSingleSignOnApplication` tag. Defaults to `false`.
     * 
     */
    public Optional<Boolean> customSingleSignOn() {
        return Optional.ofNullable(this.customSingleSignOn);
    }
    /**
     * @return Whether this service principal represents an Enterprise Application. Enabling this will assign the `WindowsAzureActiveDirectoryIntegratedApp` tag. Defaults to `false`.
     * 
     */
    public Optional<Boolean> enterprise() {
        return Optional.ofNullable(this.enterprise);
    }
    /**
     * @return Whether this service principal represents a gallery application. Enabling this will assign the `WindowsAzureActiveDirectoryGalleryApplicationNonPrimaryV1` tag. Defaults to `false`.
     * 
     */
    public Optional<Boolean> gallery() {
        return Optional.ofNullable(this.gallery);
    }
    /**
     * @return Whether this app is invisible to users in My Apps and Office 365 Launcher. Enabling this will assign the `HideApp` tag. Defaults to `false`.
     * 
     */
    public Optional<Boolean> hide() {
        return Optional.ofNullable(this.hide);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServicePrincipalFeatureTag defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean customSingleSignOn;
        private @Nullable Boolean enterprise;
        private @Nullable Boolean gallery;
        private @Nullable Boolean hide;
        public Builder() {}
        public Builder(ServicePrincipalFeatureTag defaults) {
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
        public ServicePrincipalFeatureTag build() {
            final var _resultValue = new ServicePrincipalFeatureTag();
            _resultValue.customSingleSignOn = customSingleSignOn;
            _resultValue.enterprise = enterprise;
            _resultValue.gallery = gallery;
            _resultValue.hide = hide;
            return _resultValue;
        }
    }
}

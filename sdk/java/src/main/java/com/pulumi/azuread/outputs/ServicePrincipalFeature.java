// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ServicePrincipalFeature {
    private @Nullable Boolean customSingleSignOnApp;
    private @Nullable Boolean enterpriseApplication;
    private @Nullable Boolean galleryApplication;
    private @Nullable Boolean visibleToUsers;

    private ServicePrincipalFeature() {}
    public Optional<Boolean> customSingleSignOnApp() {
        return Optional.ofNullable(this.customSingleSignOnApp);
    }
    public Optional<Boolean> enterpriseApplication() {
        return Optional.ofNullable(this.enterpriseApplication);
    }
    public Optional<Boolean> galleryApplication() {
        return Optional.ofNullable(this.galleryApplication);
    }
    public Optional<Boolean> visibleToUsers() {
        return Optional.ofNullable(this.visibleToUsers);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServicePrincipalFeature defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean customSingleSignOnApp;
        private @Nullable Boolean enterpriseApplication;
        private @Nullable Boolean galleryApplication;
        private @Nullable Boolean visibleToUsers;
        public Builder() {}
        public Builder(ServicePrincipalFeature defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.customSingleSignOnApp = defaults.customSingleSignOnApp;
    	      this.enterpriseApplication = defaults.enterpriseApplication;
    	      this.galleryApplication = defaults.galleryApplication;
    	      this.visibleToUsers = defaults.visibleToUsers;
        }

        @CustomType.Setter
        public Builder customSingleSignOnApp(@Nullable Boolean customSingleSignOnApp) {

            this.customSingleSignOnApp = customSingleSignOnApp;
            return this;
        }
        @CustomType.Setter
        public Builder enterpriseApplication(@Nullable Boolean enterpriseApplication) {

            this.enterpriseApplication = enterpriseApplication;
            return this;
        }
        @CustomType.Setter
        public Builder galleryApplication(@Nullable Boolean galleryApplication) {

            this.galleryApplication = galleryApplication;
            return this;
        }
        @CustomType.Setter
        public Builder visibleToUsers(@Nullable Boolean visibleToUsers) {

            this.visibleToUsers = visibleToUsers;
            return this;
        }
        public ServicePrincipalFeature build() {
            final var _resultValue = new ServicePrincipalFeature();
            _resultValue.customSingleSignOnApp = customSingleSignOnApp;
            _resultValue.enterpriseApplication = enterpriseApplication;
            _resultValue.galleryApplication = galleryApplication;
            _resultValue.visibleToUsers = visibleToUsers;
            return _resultValue;
        }
    }
}

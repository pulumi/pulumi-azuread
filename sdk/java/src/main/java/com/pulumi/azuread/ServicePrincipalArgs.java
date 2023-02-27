// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.inputs.ServicePrincipalFeatureArgs;
import com.pulumi.azuread.inputs.ServicePrincipalFeatureTagArgs;
import com.pulumi.azuread.inputs.ServicePrincipalSamlSingleSignOnArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServicePrincipalArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServicePrincipalArgs Empty = new ServicePrincipalArgs();

    /**
     * Whether or not the service principal account is enabled. Defaults to `true`.
     * 
     */
    @Import(name="accountEnabled")
    private @Nullable Output<Boolean> accountEnabled;

    /**
     * @return Whether or not the service principal account is enabled. Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> accountEnabled() {
        return Optional.ofNullable(this.accountEnabled);
    }

    /**
     * A set of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities.
     * 
     */
    @Import(name="alternativeNames")
    private @Nullable Output<List<String>> alternativeNames;

    /**
     * @return A set of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities.
     * 
     */
    public Optional<Output<List<String>>> alternativeNames() {
        return Optional.ofNullable(this.alternativeNames);
    }

    /**
     * Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application. Defaults to `false`.
     * 
     */
    @Import(name="appRoleAssignmentRequired")
    private @Nullable Output<Boolean> appRoleAssignmentRequired;

    /**
     * @return Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> appRoleAssignmentRequired() {
        return Optional.ofNullable(this.appRoleAssignmentRequired);
    }

    /**
     * The application ID (client ID) of the application for which to create a service principal.
     * 
     */
    @Import(name="applicationId", required=true)
    private Output<String> applicationId;

    /**
     * @return The application ID (client ID) of the application for which to create a service principal.
     * 
     */
    public Output<String> applicationId() {
        return this.applicationId;
    }

    /**
     * A description of the service principal provided for internal end-users.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A description of the service principal provided for internal end-users.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * A `feature_tags` block as described below. Cannot be used together with the `tags` property.
     * 
     */
    @Import(name="featureTags")
    private @Nullable Output<List<ServicePrincipalFeatureTagArgs>> featureTags;

    /**
     * @return A `feature_tags` block as described below. Cannot be used together with the `tags` property.
     * 
     */
    public Optional<Output<List<ServicePrincipalFeatureTagArgs>>> featureTags() {
        return Optional.ofNullable(this.featureTags);
    }

    /**
     * Block of features to configure for this service principal using tags
     * 
     * @deprecated
     * This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider
     * 
     */
    @Deprecated /* This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider */
    @Import(name="features")
    private @Nullable Output<List<ServicePrincipalFeatureArgs>> features;

    /**
     * @return Block of features to configure for this service principal using tags
     * 
     * @deprecated
     * This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider
     * 
     */
    @Deprecated /* This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider */
    public Optional<Output<List<ServicePrincipalFeatureArgs>>> features() {
        return Optional.ofNullable(this.features);
    }

    /**
     * The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on.
     * 
     */
    @Import(name="loginUrl")
    private @Nullable Output<String> loginUrl;

    /**
     * @return The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on.
     * 
     */
    public Optional<Output<String>> loginUrl() {
        return Optional.ofNullable(this.loginUrl);
    }

    /**
     * A free text field to capture information about the service principal, typically used for operational purposes.
     * 
     */
    @Import(name="notes")
    private @Nullable Output<String> notes;

    /**
     * @return A free text field to capture information about the service principal, typically used for operational purposes.
     * 
     */
    public Optional<Output<String>> notes() {
        return Optional.ofNullable(this.notes);
    }

    /**
     * A set of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
     * 
     */
    @Import(name="notificationEmailAddresses")
    private @Nullable Output<List<String>> notificationEmailAddresses;

    /**
     * @return A set of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
     * 
     */
    public Optional<Output<List<String>>> notificationEmailAddresses() {
        return Optional.ofNullable(this.notificationEmailAddresses);
    }

    /**
     * A set of object IDs of principals that will be granted ownership of the service principal. Supported object types are users or service principals. By default, no owners are assigned.
     * 
     */
    @Import(name="owners")
    private @Nullable Output<List<String>> owners;

    /**
     * @return A set of object IDs of principals that will be granted ownership of the service principal. Supported object types are users or service principals. By default, no owners are assigned.
     * 
     */
    public Optional<Output<List<String>>> owners() {
        return Optional.ofNullable(this.owners);
    }

    /**
     * The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps. Supported values are `oidc`, `password`, `saml` or `notSupported`. Omit this property or specify a blank string to unset.
     * 
     */
    @Import(name="preferredSingleSignOnMode")
    private @Nullable Output<String> preferredSingleSignOnMode;

    /**
     * @return The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps. Supported values are `oidc`, `password`, `saml` or `notSupported`. Omit this property or specify a blank string to unset.
     * 
     */
    public Optional<Output<String>> preferredSingleSignOnMode() {
        return Optional.ofNullable(this.preferredSingleSignOnMode);
    }

    /**
     * A `saml_single_sign_on` block as documented below.
     * 
     */
    @Import(name="samlSingleSignOn")
    private @Nullable Output<ServicePrincipalSamlSingleSignOnArgs> samlSingleSignOn;

    /**
     * @return A `saml_single_sign_on` block as documented below.
     * 
     */
    public Optional<Output<ServicePrincipalSamlSingleSignOnArgs>> samlSingleSignOn() {
        return Optional.ofNullable(this.samlSingleSignOn);
    }

    /**
     * A set of tags to apply to the service principal for configuring specific behaviours of the service principal. Note that these are not provided for use by practitioners. Cannot be used together with the `feature_tags` block.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return A set of tags to apply to the service principal for configuring specific behaviours of the service principal. Note that these are not provided for use by practitioners. Cannot be used together with the `feature_tags` block.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * When true, any existing service principal linked to the same application will be automatically imported. When false, an import error will be raised for any pre-existing service principal.
     * 
     */
    @Import(name="useExisting")
    private @Nullable Output<Boolean> useExisting;

    /**
     * @return When true, any existing service principal linked to the same application will be automatically imported. When false, an import error will be raised for any pre-existing service principal.
     * 
     */
    public Optional<Output<Boolean>> useExisting() {
        return Optional.ofNullable(this.useExisting);
    }

    private ServicePrincipalArgs() {}

    private ServicePrincipalArgs(ServicePrincipalArgs $) {
        this.accountEnabled = $.accountEnabled;
        this.alternativeNames = $.alternativeNames;
        this.appRoleAssignmentRequired = $.appRoleAssignmentRequired;
        this.applicationId = $.applicationId;
        this.description = $.description;
        this.featureTags = $.featureTags;
        this.features = $.features;
        this.loginUrl = $.loginUrl;
        this.notes = $.notes;
        this.notificationEmailAddresses = $.notificationEmailAddresses;
        this.owners = $.owners;
        this.preferredSingleSignOnMode = $.preferredSingleSignOnMode;
        this.samlSingleSignOn = $.samlSingleSignOn;
        this.tags = $.tags;
        this.useExisting = $.useExisting;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServicePrincipalArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServicePrincipalArgs $;

        public Builder() {
            $ = new ServicePrincipalArgs();
        }

        public Builder(ServicePrincipalArgs defaults) {
            $ = new ServicePrincipalArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accountEnabled Whether or not the service principal account is enabled. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder accountEnabled(@Nullable Output<Boolean> accountEnabled) {
            $.accountEnabled = accountEnabled;
            return this;
        }

        /**
         * @param accountEnabled Whether or not the service principal account is enabled. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder accountEnabled(Boolean accountEnabled) {
            return accountEnabled(Output.of(accountEnabled));
        }

        /**
         * @param alternativeNames A set of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities.
         * 
         * @return builder
         * 
         */
        public Builder alternativeNames(@Nullable Output<List<String>> alternativeNames) {
            $.alternativeNames = alternativeNames;
            return this;
        }

        /**
         * @param alternativeNames A set of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities.
         * 
         * @return builder
         * 
         */
        public Builder alternativeNames(List<String> alternativeNames) {
            return alternativeNames(Output.of(alternativeNames));
        }

        /**
         * @param alternativeNames A set of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities.
         * 
         * @return builder
         * 
         */
        public Builder alternativeNames(String... alternativeNames) {
            return alternativeNames(List.of(alternativeNames));
        }

        /**
         * @param appRoleAssignmentRequired Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder appRoleAssignmentRequired(@Nullable Output<Boolean> appRoleAssignmentRequired) {
            $.appRoleAssignmentRequired = appRoleAssignmentRequired;
            return this;
        }

        /**
         * @param appRoleAssignmentRequired Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder appRoleAssignmentRequired(Boolean appRoleAssignmentRequired) {
            return appRoleAssignmentRequired(Output.of(appRoleAssignmentRequired));
        }

        /**
         * @param applicationId The application ID (client ID) of the application for which to create a service principal.
         * 
         * @return builder
         * 
         */
        public Builder applicationId(Output<String> applicationId) {
            $.applicationId = applicationId;
            return this;
        }

        /**
         * @param applicationId The application ID (client ID) of the application for which to create a service principal.
         * 
         * @return builder
         * 
         */
        public Builder applicationId(String applicationId) {
            return applicationId(Output.of(applicationId));
        }

        /**
         * @param description A description of the service principal provided for internal end-users.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A description of the service principal provided for internal end-users.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param featureTags A `feature_tags` block as described below. Cannot be used together with the `tags` property.
         * 
         * @return builder
         * 
         */
        public Builder featureTags(@Nullable Output<List<ServicePrincipalFeatureTagArgs>> featureTags) {
            $.featureTags = featureTags;
            return this;
        }

        /**
         * @param featureTags A `feature_tags` block as described below. Cannot be used together with the `tags` property.
         * 
         * @return builder
         * 
         */
        public Builder featureTags(List<ServicePrincipalFeatureTagArgs> featureTags) {
            return featureTags(Output.of(featureTags));
        }

        /**
         * @param featureTags A `feature_tags` block as described below. Cannot be used together with the `tags` property.
         * 
         * @return builder
         * 
         */
        public Builder featureTags(ServicePrincipalFeatureTagArgs... featureTags) {
            return featureTags(List.of(featureTags));
        }

        /**
         * @param features Block of features to configure for this service principal using tags
         * 
         * @return builder
         * 
         * @deprecated
         * This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider
         * 
         */
        @Deprecated /* This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider */
        public Builder features(@Nullable Output<List<ServicePrincipalFeatureArgs>> features) {
            $.features = features;
            return this;
        }

        /**
         * @param features Block of features to configure for this service principal using tags
         * 
         * @return builder
         * 
         * @deprecated
         * This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider
         * 
         */
        @Deprecated /* This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider */
        public Builder features(List<ServicePrincipalFeatureArgs> features) {
            return features(Output.of(features));
        }

        /**
         * @param features Block of features to configure for this service principal using tags
         * 
         * @return builder
         * 
         * @deprecated
         * This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider
         * 
         */
        @Deprecated /* This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider */
        public Builder features(ServicePrincipalFeatureArgs... features) {
            return features(List.of(features));
        }

        /**
         * @param loginUrl The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on.
         * 
         * @return builder
         * 
         */
        public Builder loginUrl(@Nullable Output<String> loginUrl) {
            $.loginUrl = loginUrl;
            return this;
        }

        /**
         * @param loginUrl The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on.
         * 
         * @return builder
         * 
         */
        public Builder loginUrl(String loginUrl) {
            return loginUrl(Output.of(loginUrl));
        }

        /**
         * @param notes A free text field to capture information about the service principal, typically used for operational purposes.
         * 
         * @return builder
         * 
         */
        public Builder notes(@Nullable Output<String> notes) {
            $.notes = notes;
            return this;
        }

        /**
         * @param notes A free text field to capture information about the service principal, typically used for operational purposes.
         * 
         * @return builder
         * 
         */
        public Builder notes(String notes) {
            return notes(Output.of(notes));
        }

        /**
         * @param notificationEmailAddresses A set of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
         * 
         * @return builder
         * 
         */
        public Builder notificationEmailAddresses(@Nullable Output<List<String>> notificationEmailAddresses) {
            $.notificationEmailAddresses = notificationEmailAddresses;
            return this;
        }

        /**
         * @param notificationEmailAddresses A set of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
         * 
         * @return builder
         * 
         */
        public Builder notificationEmailAddresses(List<String> notificationEmailAddresses) {
            return notificationEmailAddresses(Output.of(notificationEmailAddresses));
        }

        /**
         * @param notificationEmailAddresses A set of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
         * 
         * @return builder
         * 
         */
        public Builder notificationEmailAddresses(String... notificationEmailAddresses) {
            return notificationEmailAddresses(List.of(notificationEmailAddresses));
        }

        /**
         * @param owners A set of object IDs of principals that will be granted ownership of the service principal. Supported object types are users or service principals. By default, no owners are assigned.
         * 
         * @return builder
         * 
         */
        public Builder owners(@Nullable Output<List<String>> owners) {
            $.owners = owners;
            return this;
        }

        /**
         * @param owners A set of object IDs of principals that will be granted ownership of the service principal. Supported object types are users or service principals. By default, no owners are assigned.
         * 
         * @return builder
         * 
         */
        public Builder owners(List<String> owners) {
            return owners(Output.of(owners));
        }

        /**
         * @param owners A set of object IDs of principals that will be granted ownership of the service principal. Supported object types are users or service principals. By default, no owners are assigned.
         * 
         * @return builder
         * 
         */
        public Builder owners(String... owners) {
            return owners(List.of(owners));
        }

        /**
         * @param preferredSingleSignOnMode The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps. Supported values are `oidc`, `password`, `saml` or `notSupported`. Omit this property or specify a blank string to unset.
         * 
         * @return builder
         * 
         */
        public Builder preferredSingleSignOnMode(@Nullable Output<String> preferredSingleSignOnMode) {
            $.preferredSingleSignOnMode = preferredSingleSignOnMode;
            return this;
        }

        /**
         * @param preferredSingleSignOnMode The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps. Supported values are `oidc`, `password`, `saml` or `notSupported`. Omit this property or specify a blank string to unset.
         * 
         * @return builder
         * 
         */
        public Builder preferredSingleSignOnMode(String preferredSingleSignOnMode) {
            return preferredSingleSignOnMode(Output.of(preferredSingleSignOnMode));
        }

        /**
         * @param samlSingleSignOn A `saml_single_sign_on` block as documented below.
         * 
         * @return builder
         * 
         */
        public Builder samlSingleSignOn(@Nullable Output<ServicePrincipalSamlSingleSignOnArgs> samlSingleSignOn) {
            $.samlSingleSignOn = samlSingleSignOn;
            return this;
        }

        /**
         * @param samlSingleSignOn A `saml_single_sign_on` block as documented below.
         * 
         * @return builder
         * 
         */
        public Builder samlSingleSignOn(ServicePrincipalSamlSingleSignOnArgs samlSingleSignOn) {
            return samlSingleSignOn(Output.of(samlSingleSignOn));
        }

        /**
         * @param tags A set of tags to apply to the service principal for configuring specific behaviours of the service principal. Note that these are not provided for use by practitioners. Cannot be used together with the `feature_tags` block.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A set of tags to apply to the service principal for configuring specific behaviours of the service principal. Note that these are not provided for use by practitioners. Cannot be used together with the `feature_tags` block.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags A set of tags to apply to the service principal for configuring specific behaviours of the service principal. Note that these are not provided for use by practitioners. Cannot be used together with the `feature_tags` block.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param useExisting When true, any existing service principal linked to the same application will be automatically imported. When false, an import error will be raised for any pre-existing service principal.
         * 
         * @return builder
         * 
         */
        public Builder useExisting(@Nullable Output<Boolean> useExisting) {
            $.useExisting = useExisting;
            return this;
        }

        /**
         * @param useExisting When true, any existing service principal linked to the same application will be automatically imported. When false, an import error will be raised for any pre-existing service principal.
         * 
         * @return builder
         * 
         */
        public Builder useExisting(Boolean useExisting) {
            return useExisting(Output.of(useExisting));
        }

        public ServicePrincipalArgs build() {
            $.applicationId = Objects.requireNonNull($.applicationId, "expected parameter 'applicationId' to be non-null");
            return $;
        }
    }

}

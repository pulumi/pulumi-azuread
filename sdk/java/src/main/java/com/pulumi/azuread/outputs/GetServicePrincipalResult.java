// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.azuread.outputs.GetServicePrincipalAppRole;
import com.pulumi.azuread.outputs.GetServicePrincipalFeature;
import com.pulumi.azuread.outputs.GetServicePrincipalFeatureTag;
import com.pulumi.azuread.outputs.GetServicePrincipalOauth2PermissionScope;
import com.pulumi.azuread.outputs.GetServicePrincipalSamlSingleSignOn;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetServicePrincipalResult {
    /**
     * @return Whether the service principal account is enabled.
     * 
     */
    private Boolean accountEnabled;
    /**
     * @return A list of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities.
     * 
     */
    private List<String> alternativeNames;
    /**
     * @return Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application.
     * 
     */
    private Boolean appRoleAssignmentRequired;
    /**
     * @return A mapping of app role values to app role IDs, as published by the associated application, intended to be useful when referencing app roles in other resources in your configuration.
     * 
     */
    private Map<String,String> appRoleIds;
    /**
     * @return A list of app roles published by the associated application, as documented below. For more information [official documentation](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).
     * 
     */
    private List<GetServicePrincipalAppRole> appRoles;
    /**
     * @return The tenant ID where the associated application is registered.
     * 
     */
    private String applicationTenantId;
    /**
     * @return The client ID of the application associated with this service principal.
     * 
     */
    private String clientId;
    /**
     * @return Permission help text that appears in the admin app assignment and consent experiences.
     * 
     */
    private String description;
    /**
     * @return Display name for the permission that appears in the admin consent and app assignment experiences.
     * 
     */
    private String displayName;
    private List<GetServicePrincipalFeatureTag> featureTags;
    /**
     * @return A `features` block as described below.
     * 
     * @deprecated
     * This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider
     * 
     */
    @Deprecated /* This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider */
    private List<GetServicePrincipalFeature> features;
    /**
     * @return Home page or landing page of the associated application.
     * 
     */
    private String homepageUrl;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps.
     * 
     */
    private String loginUrl;
    /**
     * @return The URL that will be used by Microsoft&#39;s authorization service to logout an user using OpenId Connect front-channel, back-channel or SAML logout protocols, taken from the associated application.
     * 
     */
    private String logoutUrl;
    /**
     * @return A free text field to capture information about the service principal, typically used for operational purposes.
     * 
     */
    private String notes;
    /**
     * @return A list of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
     * 
     */
    private List<String> notificationEmailAddresses;
    /**
     * @return A mapping of OAuth2.0 permission scope values to scope IDs, as exposed by the associated application, intended to be useful when referencing permission scopes in other resources in your configuration.
     * 
     */
    private Map<String,String> oauth2PermissionScopeIds;
    /**
     * @return A collection of OAuth 2.0 delegated permissions exposed by the associated application. Each permission is covered by an `oauth2_permission_scopes` block as documented below.
     * 
     */
    private List<GetServicePrincipalOauth2PermissionScope> oauth2PermissionScopes;
    /**
     * @return The object ID of the service principal.
     * 
     */
    private String objectId;
    /**
     * @return The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps.
     * 
     */
    private String preferredSingleSignOnMode;
    /**
     * @return A list of URLs where user tokens are sent for sign-in with the associated application, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent for the associated application.
     * 
     */
    private List<String> redirectUris;
    /**
     * @return The URL where the service exposes SAML metadata for federation.
     * 
     */
    private String samlMetadataUrl;
    /**
     * @return A `saml_single_sign_on` block as documented below.
     * 
     */
    private List<GetServicePrincipalSamlSingleSignOn> samlSingleSignOns;
    /**
     * @return A list of identifier URI(s), copied over from the associated application.
     * 
     */
    private List<String> servicePrincipalNames;
    /**
     * @return The Microsoft account types that are supported for the associated application. Possible values include `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`.
     * 
     */
    private String signInAudience;
    /**
     * @return A list of tags applied to the service principal.
     * 
     */
    private List<String> tags;
    /**
     * @return Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Possible values are `User` or `Admin`.
     * 
     */
    private String type;

    private GetServicePrincipalResult() {}
    /**
     * @return Whether the service principal account is enabled.
     * 
     */
    public Boolean accountEnabled() {
        return this.accountEnabled;
    }
    /**
     * @return A list of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities.
     * 
     */
    public List<String> alternativeNames() {
        return this.alternativeNames;
    }
    /**
     * @return Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application.
     * 
     */
    public Boolean appRoleAssignmentRequired() {
        return this.appRoleAssignmentRequired;
    }
    /**
     * @return A mapping of app role values to app role IDs, as published by the associated application, intended to be useful when referencing app roles in other resources in your configuration.
     * 
     */
    public Map<String,String> appRoleIds() {
        return this.appRoleIds;
    }
    /**
     * @return A list of app roles published by the associated application, as documented below. For more information [official documentation](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).
     * 
     */
    public List<GetServicePrincipalAppRole> appRoles() {
        return this.appRoles;
    }
    /**
     * @return The tenant ID where the associated application is registered.
     * 
     */
    public String applicationTenantId() {
        return this.applicationTenantId;
    }
    /**
     * @return The client ID of the application associated with this service principal.
     * 
     */
    public String clientId() {
        return this.clientId;
    }
    /**
     * @return Permission help text that appears in the admin app assignment and consent experiences.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return Display name for the permission that appears in the admin consent and app assignment experiences.
     * 
     */
    public String displayName() {
        return this.displayName;
    }
    public List<GetServicePrincipalFeatureTag> featureTags() {
        return this.featureTags;
    }
    /**
     * @return A `features` block as described below.
     * 
     * @deprecated
     * This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider
     * 
     */
    @Deprecated /* This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider */
    public List<GetServicePrincipalFeature> features() {
        return this.features;
    }
    /**
     * @return Home page or landing page of the associated application.
     * 
     */
    public String homepageUrl() {
        return this.homepageUrl;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps.
     * 
     */
    public String loginUrl() {
        return this.loginUrl;
    }
    /**
     * @return The URL that will be used by Microsoft&#39;s authorization service to logout an user using OpenId Connect front-channel, back-channel or SAML logout protocols, taken from the associated application.
     * 
     */
    public String logoutUrl() {
        return this.logoutUrl;
    }
    /**
     * @return A free text field to capture information about the service principal, typically used for operational purposes.
     * 
     */
    public String notes() {
        return this.notes;
    }
    /**
     * @return A list of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
     * 
     */
    public List<String> notificationEmailAddresses() {
        return this.notificationEmailAddresses;
    }
    /**
     * @return A mapping of OAuth2.0 permission scope values to scope IDs, as exposed by the associated application, intended to be useful when referencing permission scopes in other resources in your configuration.
     * 
     */
    public Map<String,String> oauth2PermissionScopeIds() {
        return this.oauth2PermissionScopeIds;
    }
    /**
     * @return A collection of OAuth 2.0 delegated permissions exposed by the associated application. Each permission is covered by an `oauth2_permission_scopes` block as documented below.
     * 
     */
    public List<GetServicePrincipalOauth2PermissionScope> oauth2PermissionScopes() {
        return this.oauth2PermissionScopes;
    }
    /**
     * @return The object ID of the service principal.
     * 
     */
    public String objectId() {
        return this.objectId;
    }
    /**
     * @return The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps.
     * 
     */
    public String preferredSingleSignOnMode() {
        return this.preferredSingleSignOnMode;
    }
    /**
     * @return A list of URLs where user tokens are sent for sign-in with the associated application, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent for the associated application.
     * 
     */
    public List<String> redirectUris() {
        return this.redirectUris;
    }
    /**
     * @return The URL where the service exposes SAML metadata for federation.
     * 
     */
    public String samlMetadataUrl() {
        return this.samlMetadataUrl;
    }
    /**
     * @return A `saml_single_sign_on` block as documented below.
     * 
     */
    public List<GetServicePrincipalSamlSingleSignOn> samlSingleSignOns() {
        return this.samlSingleSignOns;
    }
    /**
     * @return A list of identifier URI(s), copied over from the associated application.
     * 
     */
    public List<String> servicePrincipalNames() {
        return this.servicePrincipalNames;
    }
    /**
     * @return The Microsoft account types that are supported for the associated application. Possible values include `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`.
     * 
     */
    public String signInAudience() {
        return this.signInAudience;
    }
    /**
     * @return A list of tags applied to the service principal.
     * 
     */
    public List<String> tags() {
        return this.tags;
    }
    /**
     * @return Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Possible values are `User` or `Admin`.
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetServicePrincipalResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean accountEnabled;
        private List<String> alternativeNames;
        private Boolean appRoleAssignmentRequired;
        private Map<String,String> appRoleIds;
        private List<GetServicePrincipalAppRole> appRoles;
        private String applicationTenantId;
        private String clientId;
        private String description;
        private String displayName;
        private List<GetServicePrincipalFeatureTag> featureTags;
        private List<GetServicePrincipalFeature> features;
        private String homepageUrl;
        private String id;
        private String loginUrl;
        private String logoutUrl;
        private String notes;
        private List<String> notificationEmailAddresses;
        private Map<String,String> oauth2PermissionScopeIds;
        private List<GetServicePrincipalOauth2PermissionScope> oauth2PermissionScopes;
        private String objectId;
        private String preferredSingleSignOnMode;
        private List<String> redirectUris;
        private String samlMetadataUrl;
        private List<GetServicePrincipalSamlSingleSignOn> samlSingleSignOns;
        private List<String> servicePrincipalNames;
        private String signInAudience;
        private List<String> tags;
        private String type;
        public Builder() {}
        public Builder(GetServicePrincipalResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accountEnabled = defaults.accountEnabled;
    	      this.alternativeNames = defaults.alternativeNames;
    	      this.appRoleAssignmentRequired = defaults.appRoleAssignmentRequired;
    	      this.appRoleIds = defaults.appRoleIds;
    	      this.appRoles = defaults.appRoles;
    	      this.applicationTenantId = defaults.applicationTenantId;
    	      this.clientId = defaults.clientId;
    	      this.description = defaults.description;
    	      this.displayName = defaults.displayName;
    	      this.featureTags = defaults.featureTags;
    	      this.features = defaults.features;
    	      this.homepageUrl = defaults.homepageUrl;
    	      this.id = defaults.id;
    	      this.loginUrl = defaults.loginUrl;
    	      this.logoutUrl = defaults.logoutUrl;
    	      this.notes = defaults.notes;
    	      this.notificationEmailAddresses = defaults.notificationEmailAddresses;
    	      this.oauth2PermissionScopeIds = defaults.oauth2PermissionScopeIds;
    	      this.oauth2PermissionScopes = defaults.oauth2PermissionScopes;
    	      this.objectId = defaults.objectId;
    	      this.preferredSingleSignOnMode = defaults.preferredSingleSignOnMode;
    	      this.redirectUris = defaults.redirectUris;
    	      this.samlMetadataUrl = defaults.samlMetadataUrl;
    	      this.samlSingleSignOns = defaults.samlSingleSignOns;
    	      this.servicePrincipalNames = defaults.servicePrincipalNames;
    	      this.signInAudience = defaults.signInAudience;
    	      this.tags = defaults.tags;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder accountEnabled(Boolean accountEnabled) {
            if (accountEnabled == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalResult", "accountEnabled");
            }
            this.accountEnabled = accountEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder alternativeNames(List<String> alternativeNames) {
            if (alternativeNames == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalResult", "alternativeNames");
            }
            this.alternativeNames = alternativeNames;
            return this;
        }
        public Builder alternativeNames(String... alternativeNames) {
            return alternativeNames(List.of(alternativeNames));
        }
        @CustomType.Setter
        public Builder appRoleAssignmentRequired(Boolean appRoleAssignmentRequired) {
            if (appRoleAssignmentRequired == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalResult", "appRoleAssignmentRequired");
            }
            this.appRoleAssignmentRequired = appRoleAssignmentRequired;
            return this;
        }
        @CustomType.Setter
        public Builder appRoleIds(Map<String,String> appRoleIds) {
            if (appRoleIds == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalResult", "appRoleIds");
            }
            this.appRoleIds = appRoleIds;
            return this;
        }
        @CustomType.Setter
        public Builder appRoles(List<GetServicePrincipalAppRole> appRoles) {
            if (appRoles == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalResult", "appRoles");
            }
            this.appRoles = appRoles;
            return this;
        }
        public Builder appRoles(GetServicePrincipalAppRole... appRoles) {
            return appRoles(List.of(appRoles));
        }
        @CustomType.Setter
        public Builder applicationTenantId(String applicationTenantId) {
            if (applicationTenantId == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalResult", "applicationTenantId");
            }
            this.applicationTenantId = applicationTenantId;
            return this;
        }
        @CustomType.Setter
        public Builder clientId(String clientId) {
            if (clientId == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalResult", "clientId");
            }
            this.clientId = clientId;
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalResult", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder displayName(String displayName) {
            if (displayName == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalResult", "displayName");
            }
            this.displayName = displayName;
            return this;
        }
        @CustomType.Setter
        public Builder featureTags(List<GetServicePrincipalFeatureTag> featureTags) {
            if (featureTags == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalResult", "featureTags");
            }
            this.featureTags = featureTags;
            return this;
        }
        public Builder featureTags(GetServicePrincipalFeatureTag... featureTags) {
            return featureTags(List.of(featureTags));
        }
        @CustomType.Setter
        public Builder features(List<GetServicePrincipalFeature> features) {
            if (features == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalResult", "features");
            }
            this.features = features;
            return this;
        }
        public Builder features(GetServicePrincipalFeature... features) {
            return features(List.of(features));
        }
        @CustomType.Setter
        public Builder homepageUrl(String homepageUrl) {
            if (homepageUrl == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalResult", "homepageUrl");
            }
            this.homepageUrl = homepageUrl;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder loginUrl(String loginUrl) {
            if (loginUrl == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalResult", "loginUrl");
            }
            this.loginUrl = loginUrl;
            return this;
        }
        @CustomType.Setter
        public Builder logoutUrl(String logoutUrl) {
            if (logoutUrl == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalResult", "logoutUrl");
            }
            this.logoutUrl = logoutUrl;
            return this;
        }
        @CustomType.Setter
        public Builder notes(String notes) {
            if (notes == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalResult", "notes");
            }
            this.notes = notes;
            return this;
        }
        @CustomType.Setter
        public Builder notificationEmailAddresses(List<String> notificationEmailAddresses) {
            if (notificationEmailAddresses == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalResult", "notificationEmailAddresses");
            }
            this.notificationEmailAddresses = notificationEmailAddresses;
            return this;
        }
        public Builder notificationEmailAddresses(String... notificationEmailAddresses) {
            return notificationEmailAddresses(List.of(notificationEmailAddresses));
        }
        @CustomType.Setter
        public Builder oauth2PermissionScopeIds(Map<String,String> oauth2PermissionScopeIds) {
            if (oauth2PermissionScopeIds == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalResult", "oauth2PermissionScopeIds");
            }
            this.oauth2PermissionScopeIds = oauth2PermissionScopeIds;
            return this;
        }
        @CustomType.Setter
        public Builder oauth2PermissionScopes(List<GetServicePrincipalOauth2PermissionScope> oauth2PermissionScopes) {
            if (oauth2PermissionScopes == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalResult", "oauth2PermissionScopes");
            }
            this.oauth2PermissionScopes = oauth2PermissionScopes;
            return this;
        }
        public Builder oauth2PermissionScopes(GetServicePrincipalOauth2PermissionScope... oauth2PermissionScopes) {
            return oauth2PermissionScopes(List.of(oauth2PermissionScopes));
        }
        @CustomType.Setter
        public Builder objectId(String objectId) {
            if (objectId == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalResult", "objectId");
            }
            this.objectId = objectId;
            return this;
        }
        @CustomType.Setter
        public Builder preferredSingleSignOnMode(String preferredSingleSignOnMode) {
            if (preferredSingleSignOnMode == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalResult", "preferredSingleSignOnMode");
            }
            this.preferredSingleSignOnMode = preferredSingleSignOnMode;
            return this;
        }
        @CustomType.Setter
        public Builder redirectUris(List<String> redirectUris) {
            if (redirectUris == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalResult", "redirectUris");
            }
            this.redirectUris = redirectUris;
            return this;
        }
        public Builder redirectUris(String... redirectUris) {
            return redirectUris(List.of(redirectUris));
        }
        @CustomType.Setter
        public Builder samlMetadataUrl(String samlMetadataUrl) {
            if (samlMetadataUrl == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalResult", "samlMetadataUrl");
            }
            this.samlMetadataUrl = samlMetadataUrl;
            return this;
        }
        @CustomType.Setter
        public Builder samlSingleSignOns(List<GetServicePrincipalSamlSingleSignOn> samlSingleSignOns) {
            if (samlSingleSignOns == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalResult", "samlSingleSignOns");
            }
            this.samlSingleSignOns = samlSingleSignOns;
            return this;
        }
        public Builder samlSingleSignOns(GetServicePrincipalSamlSingleSignOn... samlSingleSignOns) {
            return samlSingleSignOns(List.of(samlSingleSignOns));
        }
        @CustomType.Setter
        public Builder servicePrincipalNames(List<String> servicePrincipalNames) {
            if (servicePrincipalNames == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalResult", "servicePrincipalNames");
            }
            this.servicePrincipalNames = servicePrincipalNames;
            return this;
        }
        public Builder servicePrincipalNames(String... servicePrincipalNames) {
            return servicePrincipalNames(List.of(servicePrincipalNames));
        }
        @CustomType.Setter
        public Builder signInAudience(String signInAudience) {
            if (signInAudience == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalResult", "signInAudience");
            }
            this.signInAudience = signInAudience;
            return this;
        }
        @CustomType.Setter
        public Builder tags(List<String> tags) {
            if (tags == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalResult", "tags");
            }
            this.tags = tags;
            return this;
        }
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }
        @CustomType.Setter
        public Builder type(String type) {
            if (type == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalResult", "type");
            }
            this.type = type;
            return this;
        }
        public GetServicePrincipalResult build() {
            final var _resultValue = new GetServicePrincipalResult();
            _resultValue.accountEnabled = accountEnabled;
            _resultValue.alternativeNames = alternativeNames;
            _resultValue.appRoleAssignmentRequired = appRoleAssignmentRequired;
            _resultValue.appRoleIds = appRoleIds;
            _resultValue.appRoles = appRoles;
            _resultValue.applicationTenantId = applicationTenantId;
            _resultValue.clientId = clientId;
            _resultValue.description = description;
            _resultValue.displayName = displayName;
            _resultValue.featureTags = featureTags;
            _resultValue.features = features;
            _resultValue.homepageUrl = homepageUrl;
            _resultValue.id = id;
            _resultValue.loginUrl = loginUrl;
            _resultValue.logoutUrl = logoutUrl;
            _resultValue.notes = notes;
            _resultValue.notificationEmailAddresses = notificationEmailAddresses;
            _resultValue.oauth2PermissionScopeIds = oauth2PermissionScopeIds;
            _resultValue.oauth2PermissionScopes = oauth2PermissionScopes;
            _resultValue.objectId = objectId;
            _resultValue.preferredSingleSignOnMode = preferredSingleSignOnMode;
            _resultValue.redirectUris = redirectUris;
            _resultValue.samlMetadataUrl = samlMetadataUrl;
            _resultValue.samlSingleSignOns = samlSingleSignOns;
            _resultValue.servicePrincipalNames = servicePrincipalNames;
            _resultValue.signInAudience = signInAudience;
            _resultValue.tags = tags;
            _resultValue.type = type;
            return _resultValue;
        }
    }
}

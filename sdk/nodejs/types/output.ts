// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";

export interface ApplicationApi {
    /**
     * A set of application IDs (client IDs), used for bundling consent if you have a solution that contains two parts: a client app and a custom web API app.
     */
    knownClientApplications?: string[];
    /**
     * Allows an application to use claims mapping without specifying a custom signing key. Defaults to `false`.
     */
    mappedClaimsEnabled?: boolean;
    /**
     * One or more `oauth2PermissionScope` blocks as documented below, to describe delegated permissions exposed by the web API represented by this application.
     */
    oauth2PermissionScopes?: outputs.ApplicationApiOauth2PermissionScope[];
    /**
     * The access token version expected by this resource. Must be one of `1` or `2`, and must be `2` when `signInAudience` is either `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount` Defaults to `1`.
     */
    requestedAccessTokenVersion?: number;
}

export interface ApplicationApiOauth2PermissionScope {
    /**
     * Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
     */
    adminConsentDescription?: string;
    /**
     * Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
     */
    adminConsentDisplayName?: string;
    /**
     * Determines if the permission scope is enabled. Defaults to `true`.
     */
    enabled?: boolean;
    /**
     * The unique identifier of the delegated permission. Must be a valid UUID.
     */
    id: string;
    /**
     * Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Defaults to `User`. Possible values are `User` or `Admin`.
     */
    type?: string;
    /**
     * Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
     */
    userConsentDescription?: string;
    /**
     * Display name for the delegated permission that appears in the end user consent experience.
     */
    userConsentDisplayName?: string;
    /**
     * The value that is used for the `scp` claim in OAuth 2.0 access tokens.
     */
    value?: string;
}

export interface ApplicationAppRole {
    /**
     * Specifies whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications (that are accessing this application in a standalone scenario) by setting to `Application`, or to both.
     */
    allowedMemberTypes: string[];
    /**
     * Description of the app role that appears when the role is being assigned and, if the role functions as an application permissions, during the consent experiences.
     */
    description: string;
    /**
     * Display name for the app role that appears during app role assignment and in consent experiences.
     */
    displayName: string;
    /**
     * Determines if the app role is enabled. Defaults to `true`.
     */
    enabled?: boolean;
    /**
     * The unique identifier of the app role. Must be a valid UUID.
     */
    id: string;
    /**
     * The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal.
     */
    value?: string;
}

export interface ApplicationFeatureTag {
    /**
     * Whether this application represents a custom SAML application for linked service principals. Enabling this will assign the `WindowsAzureActiveDirectoryCustomSingleSignOnApplication` tag. Defaults to `false`.
     */
    customSingleSignOn?: boolean;
    /**
     * Whether this application represents an Enterprise Application for linked service principals. Enabling this will assign the `WindowsAzureActiveDirectoryIntegratedApp` tag. Defaults to `false`.
     */
    enterprise?: boolean;
    /**
     * Whether this application represents a gallery application for linked service principals. Enabling this will assign the `WindowsAzureActiveDirectoryGalleryApplicationNonPrimaryV1` tag. Defaults to `false`.
     */
    gallery?: boolean;
    /**
     * Whether this app is invisible to users in My Apps and Office 365 Launcher. Enabling this will assign the `HideApp` tag. Defaults to `false`.
     */
    hide?: boolean;
}

export interface ApplicationOptionalClaims {
    /**
     * One or more `accessToken` blocks as documented below.
     */
    accessTokens?: outputs.ApplicationOptionalClaimsAccessToken[];
    /**
     * One or more `idToken` blocks as documented below.
     */
    idTokens?: outputs.ApplicationOptionalClaimsIdToken[];
    /**
     * One or more `saml2Token` blocks as documented below.
     */
    saml2Tokens?: outputs.ApplicationOptionalClaimsSaml2Token[];
}

export interface ApplicationOptionalClaimsAccessToken {
    /**
     * List of additional properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim.
     */
    additionalProperties?: string[];
    /**
     * Whether the claim specified by the client is necessary to ensure a smooth authorization experience.
     */
    essential?: boolean;
    /**
     * The name of the optional claim.
     */
    name: string;
    /**
     * The source of the claim. If `source` is absent, the claim is a predefined optional claim. If `source` is `user`, the value of `name` is the extension property from the user object.
     */
    source?: string;
}

export interface ApplicationOptionalClaimsIdToken {
    /**
     * List of additional properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim.
     */
    additionalProperties?: string[];
    /**
     * Whether the claim specified by the client is necessary to ensure a smooth authorization experience.
     */
    essential?: boolean;
    /**
     * The name of the optional claim.
     */
    name: string;
    /**
     * The source of the claim. If `source` is absent, the claim is a predefined optional claim. If `source` is `user`, the value of `name` is the extension property from the user object.
     */
    source?: string;
}

export interface ApplicationOptionalClaimsSaml2Token {
    /**
     * List of additional properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim.
     */
    additionalProperties?: string[];
    /**
     * Whether the claim specified by the client is necessary to ensure a smooth authorization experience.
     */
    essential?: boolean;
    /**
     * The name of the optional claim.
     */
    name: string;
    /**
     * The source of the claim. If `source` is absent, the claim is a predefined optional claim. If `source` is `user`, the value of `name` is the extension property from the user object.
     */
    source?: string;
}

export interface ApplicationPublicClient {
    /**
     * A set of URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent. Must be a valid `https` or `ms-appx-web` URL.
     */
    redirectUris?: string[];
}

export interface ApplicationRequiredResourceAccess {
    /**
     * A collection of `resourceAccess` blocks as documented below, describing OAuth2.0 permission scopes and app roles that the application requires from the specified resource.
     */
    resourceAccesses: outputs.ApplicationRequiredResourceAccessResourceAccess[];
    /**
     * The unique identifier for the resource that the application requires access to. This should be the Application ID of the target application.
     */
    resourceAppId: string;
}

export interface ApplicationRequiredResourceAccessResourceAccess {
    /**
     * The unique identifier for an app role or OAuth2 permission scope published by the resource application.
     */
    id: string;
    /**
     * Specifies whether the `id` property references an app role or an OAuth2 permission scope. Possible values are `Role` or `Scope`.
     */
    type: string;
}

export interface ApplicationSinglePageApplication {
    /**
     * A set of URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent. Must be a valid `https` URL.
     */
    redirectUris?: string[];
}

export interface ApplicationWeb {
    /**
     * Home page or landing page of the application.
     */
    homepageUrl?: string;
    /**
     * An `implicitGrant` block as documented above.
     */
    implicitGrant?: outputs.ApplicationWebImplicitGrant;
    /**
     * The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols.
     */
    logoutUrl?: string;
    /**
     * A set of URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent. Must be a valid `http` URL or a URN.
     */
    redirectUris?: string[];
}

export interface ApplicationWebImplicitGrant {
    /**
     * Whether this web application can request an access token using OAuth 2.0 implicit flow.
     */
    accessTokenIssuanceEnabled?: boolean;
    /**
     * Whether this web application can request an ID token using OAuth 2.0 implicit flow.
     */
    idTokenIssuanceEnabled?: boolean;
}

export interface ConditionalAccessPolicyConditions {
    /**
     * An `applications` block as documented below, which specifies applications and user actions included in and excluded from the policy.
     */
    applications: outputs.ConditionalAccessPolicyConditionsApplications;
    /**
     * A list of client application types included in the policy. Possible values are: `all`, `browser`, `mobileAppsAndDesktopClients`, `exchangeActiveSync`, `easSupported` and `other`.
     */
    clientAppTypes: string[];
    /**
     * A `devices` block as documented below, which describes devices to be included in and excluded from the policy. A `devices` block can be added to an existing policy, but removing the `devices` block forces a new resource to be created.
     */
    devices?: outputs.ConditionalAccessPolicyConditionsDevices;
    /**
     * A `locations` block as documented below, which specifies locations included in and excluded from the policy.
     */
    locations?: outputs.ConditionalAccessPolicyConditionsLocations;
    /**
     * A `platforms` block as documented below, which specifies platforms included in and excluded from the policy.
     */
    platforms?: outputs.ConditionalAccessPolicyConditionsPlatforms;
    /**
     * A list of sign-in risk levels included in the policy. Possible values are: `low`, `medium`, `high`, `hidden`, `none`, `unknownFutureValue`.
     */
    signInRiskLevels?: string[];
    /**
     * A list of user risk levels included in the policy. Possible values are: `low`, `medium`, `high`, `hidden`, `none`, `unknownFutureValue`.
     */
    userRiskLevels?: string[];
    /**
     * A `users` block as documented below, which specifies users, groups, and roles included in and excluded from the policy.
     */
    users: outputs.ConditionalAccessPolicyConditionsUsers;
}

export interface ConditionalAccessPolicyConditionsApplications {
    /**
     * A list of application IDs explicitly excluded from the policy. Can also be set to `Office365`.
     */
    excludedApplications?: string[];
    /**
     * A list of application IDs the policy applies to, unless explicitly excluded (in `excludedApplications`). Can also be set to `All`, `None` or `Office365`. Cannot be specified with `includedUserActions`. One of `includedApplications` or `includedUserActions` must be specified.
     */
    includedApplications?: string[];
    /**
     * A list of user actions to include. Supported values are `urn:user:registerdevice` and `urn:user:registersecurityinfo`. Cannot be specified with `includedApplications`. One of `includedApplications` or `includedUserActions` must be specified.
     */
    includedUserActions?: string[];
}

export interface ConditionalAccessPolicyConditionsDevices {
    /**
     * A `filter` block as described below. A `filter` block can be added to an existing policy, but removing the `filter` block forces a new resource to be created.
     */
    filter?: outputs.ConditionalAccessPolicyConditionsDevicesFilter;
}

export interface ConditionalAccessPolicyConditionsDevicesFilter {
    /**
     * Whether to include in, or exclude from, matching devices from the policy. Supported values are `include` or `exclude`.
     */
    mode: string;
    /**
     * Condition filter to match devices. For more information, see [official documentation](https://docs.microsoft.com/en-us/azure/active-directory/conditional-access/concept-condition-filters-for-devices#supported-operators-and-device-properties-for-filters).
     */
    rule: string;
}

export interface ConditionalAccessPolicyConditionsLocations {
    /**
     * A list of location IDs excluded from scope of policy. Can also be set to `AllTrusted`.
     */
    excludedLocations?: string[];
    /**
     * A list of location IDs in scope of policy unless explicitly excluded. Can also be set to `All`, or `AllTrusted`.
     */
    includedLocations: string[];
}

export interface ConditionalAccessPolicyConditionsPlatforms {
    /**
     * A list of platforms explicitly excluded from the policy. Possible values are: `all`, `android`, `iOS`, `linux`, `macOS`, `windows`, `windowsPhone` or `unknownFutureValue`.
     */
    excludedPlatforms?: string[];
    /**
     * A list of platforms the policy applies to, unless explicitly excluded. Possible values are: `all`, `android`, `iOS`, `linux`, `macOS`, `windows`, `windowsPhone` or `unknownFutureValue`.
     */
    includedPlatforms: string[];
}

export interface ConditionalAccessPolicyConditionsUsers {
    /**
     * A list of group IDs excluded from scope of policy.
     */
    excludedGroups?: string[];
    /**
     * A list of role IDs excluded from scope of policy.
     */
    excludedRoles?: string[];
    /**
     * A list of user IDs excluded from scope of policy and/or `GuestsOrExternalUsers`.
     */
    excludedUsers?: string[];
    /**
     * A list of group IDs in scope of policy unless explicitly excluded.
     */
    includedGroups?: string[];
    /**
     * A list of role IDs in scope of policy unless explicitly excluded.
     */
    includedRoles?: string[];
    /**
     * A list of user IDs in scope of policy unless explicitly excluded, or `None` or `All` or `GuestsOrExternalUsers`.
     */
    includedUsers?: string[];
}

export interface ConditionalAccessPolicyGrantControls {
    /**
     * List of built-in controls required by the policy. Possible values are: `block`, `mfa`, `approvedApplication`, `compliantApplication`, `compliantDevice`, `domainJoinedDevice`, `passwordChange` or `unknownFutureValue`.
     */
    builtInControls: string[];
    /**
     * List of custom controls IDs required by the policy.
     */
    customAuthenticationFactors?: string[];
    /**
     * Defines the relationship of the grant controls. Possible values are: `AND`, `OR`.
     */
    operator: string;
    /**
     * List of terms of use IDs required by the policy.
     */
    termsOfUses?: string[];
}

export interface ConditionalAccessPolicySessionControls {
    /**
     * Whether or not application enforced restrictions are enabled. Defaults to `false`.
     */
    applicationEnforcedRestrictionsEnabled?: boolean;
    /**
     * Enables cloud app security and specifies the cloud app security policy to use. Possible values are: `blockDownloads`, `mcasConfigured`, `monitorOnly` or `unknownFutureValue`.
     */
    cloudAppSecurityPolicy?: string;
    /**
     * Session control to define whether to persist cookies or not. Possible values are: `always` or `never`.
     */
    persistentBrowserMode?: string;
    /**
     * Number of days or hours to enforce sign-in frequency. Required when `signInFrequencyPeriod` is specified. Due to an API issue, removing this property forces a new resource to be created.
     */
    signInFrequency?: number;
    /**
     * The time period to enforce sign-in frequency. Possible values are: `hours` or `days`. Required when `signInFrequencyPeriod` is specified. Due to an API issue, removing this property forces a new resource to be created.
     */
    signInFrequencyPeriod?: string;
}

export interface CustomDirectoryRolePermission {
    /**
     * A set of tasks that can be performed on a resource. For more information, see the [Permissions Reference](https://docs.microsoft.com/en-us/azure/active-directory/roles/permissions-reference) documentation.
     */
    allowedResourceActions: string[];
}

export interface GetApplicationApi {
    /**
     * A set of application IDs (client IDs), used for bundling consent if you have a solution that contains two parts: a client app and a custom web API app.
     */
    knownClientApplications: string[];
    /**
     * Allows an application to use claims mapping without specifying a custom signing key.
     */
    mappedClaimsEnabled: boolean;
    /**
     * One or more `oauth2PermissionScope` blocks as documented below, to describe delegated permissions exposed by the web API represented by this application.
     */
    oauth2PermissionScopes: outputs.GetApplicationApiOauth2PermissionScope[];
    /**
     * The access token version expected by this resource. Possible values are `1` or `2`.
     */
    requestedAccessTokenVersion: number;
}

export interface GetApplicationApiOauth2PermissionScope {
    /**
     * Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
     */
    adminConsentDescription: string;
    /**
     * Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
     */
    adminConsentDisplayName: string;
    /**
     * Determines if the app role is enabled.
     */
    enabled: boolean;
    /**
     * The unique identifier for an app role or OAuth2 permission scope published by the resource application.
     */
    id: string;
    /**
     * Specifies whether the `id` property references an app role or an OAuth2 permission scope. Possible values are `Role` or `Scope`.
     */
    type: string;
    /**
     * Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
     */
    userConsentDescription: string;
    /**
     * Display name for the delegated permission that appears in the end user consent experience.
     */
    userConsentDisplayName: string;
    /**
     * The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal.
     */
    value: string;
}

export interface GetApplicationAppRole {
    /**
     * Specifies whether this app role definition can be assigned to users and groups, or to other applications (that are accessing this application in a standalone scenario). Possible values are `User` or `Application`, or both.
     */
    allowedMemberTypes: string[];
    /**
     * Description of the app role that appears when the role is being assigned and, if the role functions as an application permissions, during the consent experiences.
     */
    description: string;
    /**
     * Specifies the display name of the application.
     */
    displayName: string;
    /**
     * Determines if the app role is enabled.
     */
    enabled: boolean;
    /**
     * The unique identifier for an app role or OAuth2 permission scope published by the resource application.
     */
    id: string;
    /**
     * The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal.
     */
    value: string;
}

export interface GetApplicationFeatureTag {
    /**
     * Whether this application represents a custom SAML application for linked service principals.
     */
    customSingleSignOn?: boolean;
    /**
     * Whether this application represents an Enterprise Application for linked service principals.
     */
    enterprise?: boolean;
    /**
     * Whether this application represents a gallery application for linked service principals.
     */
    gallery?: boolean;
    /**
     * Whether this app is visible to users in My Apps and Office 365 Launcher.
     */
    hide?: boolean;
}

export interface GetApplicationOptionalClaim {
    /**
     * One or more `accessToken` blocks as documented below.
     */
    accessTokens?: outputs.GetApplicationOptionalClaimAccessToken[];
    /**
     * One or more `idToken` blocks as documented below.
     */
    idTokens?: outputs.GetApplicationOptionalClaimIdToken[];
    /**
     * One or more `saml2Token` blocks as documented below.
     */
    saml2Tokens?: outputs.GetApplicationOptionalClaimSaml2Token[];
}

export interface GetApplicationOptionalClaimAccessToken {
    /**
     * List of Additional Properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim.
     */
    additionalProperties?: string[];
    /**
     * Whether the claim specified by the client is necessary to ensure a smooth authorization experience.
     */
    essential?: boolean;
    /**
     * The name of the optional claim.
     */
    name: string;
    /**
     * The source of the claim. If `source` is absent, the claim is a predefined optional claim. If `source` is `user`, the value of `name` is the extension property from the user object.
     */
    source?: string;
}

export interface GetApplicationOptionalClaimIdToken {
    /**
     * List of Additional Properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim.
     */
    additionalProperties?: string[];
    /**
     * Whether the claim specified by the client is necessary to ensure a smooth authorization experience.
     */
    essential?: boolean;
    /**
     * The name of the optional claim.
     */
    name: string;
    /**
     * The source of the claim. If `source` is absent, the claim is a predefined optional claim. If `source` is `user`, the value of `name` is the extension property from the user object.
     */
    source?: string;
}

export interface GetApplicationOptionalClaimSaml2Token {
    /**
     * List of Additional Properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim.
     */
    additionalProperties?: string[];
    /**
     * Whether the claim specified by the client is necessary to ensure a smooth authorization experience.
     */
    essential?: boolean;
    /**
     * The name of the optional claim.
     */
    name: string;
    /**
     * The source of the claim. If `source` is absent, the claim is a predefined optional claim. If `source` is `user`, the value of `name` is the extension property from the user object.
     */
    source?: string;
}

export interface GetApplicationPublicClient {
    /**
     * A list of URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent.
     */
    redirectUris: string[];
}

export interface GetApplicationRequiredResourceAccess {
    /**
     * A collection of `resourceAccess` blocks as documented below, describing OAuth2.0 permission scopes and app roles that the application requires from the specified resource.
     */
    resourceAccesses: outputs.GetApplicationRequiredResourceAccessResourceAccess[];
    /**
     * The unique identifier for the resource that the application requires access to. This is the Application ID of the target application.
     */
    resourceAppId: string;
}

export interface GetApplicationRequiredResourceAccessResourceAccess {
    /**
     * The unique identifier for an app role or OAuth2 permission scope published by the resource application.
     */
    id: string;
    /**
     * Specifies whether the `id` property references an app role or an OAuth2 permission scope. Possible values are `Role` or `Scope`.
     */
    type: string;
}

export interface GetApplicationSinglePageApplication {
    /**
     * A list of URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent.
     */
    redirectUris: string[];
}

export interface GetApplicationWeb {
    /**
     * Home page or landing page of the application.
     */
    homepageUrl: string;
    /**
     * An `implicitGrant` block as documented above.
     */
    implicitGrants: outputs.GetApplicationWebImplicitGrant[];
    /**
     * The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols.
     */
    logoutUrl: string;
    /**
     * A list of URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent.
     */
    redirectUris: string[];
}

export interface GetApplicationWebImplicitGrant {
    /**
     * Whether this web application can request an access token using OAuth 2.0 implicit flow.
     */
    accessTokenIssuanceEnabled: boolean;
    /**
     * Whether this web application can request an ID token using OAuth 2.0 implicit flow.
     */
    idTokenIssuanceEnabled: boolean;
}

export interface GetDomainsDomain {
    /**
     * Set to `true` to only return domains whose DNS is managed by Microsoft 365. Defaults to `false`.
     */
    adminManaged: boolean;
    /**
     * The authentication type of the domain. Possible values include `Managed` or `Federated`.
     */
    authenticationType: string;
    /**
     * Whether this is the default domain that is used for user creation.
     */
    default: boolean;
    /**
     * The name of the domain.
     */
    domainName: string;
    /**
     * Whether this is the initial domain created by Azure Active Directory.
     */
    initial: boolean;
    /**
     * Whether the domain is a verified root domain (not a subdomain).
     */
    root: boolean;
    /**
     * A list of capabilities / services supported by the domain. Possible values include `Email`, `Sharepoint`, `EmailInternalRelayOnly`, `OfficeCommunicationsOnline`, `SharePointDefaultDomain`, `FullRedelegation`, `SharePointPublic`, `OrgIdAuthentication`, `Yammer` and `Intune`.
     */
    supportedServices: string[];
    /**
     * Whether the domain has completed domain ownership verification.
     */
    verified: boolean;
}

export interface GetGroupDynamicMembership {
    /**
     * Whether rule processing is "On" (true) or "Paused" (false).
     */
    enabled: boolean;
    /**
     * The rule that determines membership of this group.
     */
    rule: string;
}

export interface GetServicePrincipalAppRole {
    /**
     * Specifies whether this app role definition can be assigned to users and groups, or to other applications (that are accessing this application in daemon service scenarios). Possible values are: `User` and `Application`, or both.
     */
    allowedMemberTypes: string[];
    /**
     * Permission help text that appears in the admin app assignment and consent experiences.
     */
    description: string;
    /**
     * The display name of the application associated with this service principal.
     */
    displayName: string;
    /**
     * Determines if the permission scope is enabled.
     */
    enabled: boolean;
    /**
     * The unique identifier of the delegated permission. Must be a valid UUID.
     */
    id: string;
    /**
     * The value that is used for the `scp` claim in OAuth 2.0 access tokens.
     */
    value: string;
}

export interface GetServicePrincipalFeature {
    /**
     * Whether this service principal represents a custom SAML application.
     */
    customSingleSignOnApp: boolean;
    /**
     * Whether this service principal represents an Enterprise Application.
     */
    enterpriseApplication: boolean;
    /**
     * Whether this service principal represents a gallery application.
     */
    galleryApplication: boolean;
    /**
     * Whether this app is visible to users in My Apps and Office 365 Launcher.
     */
    visibleToUsers: boolean;
}

export interface GetServicePrincipalFeatureTag {
    customSingleSignOn: boolean;
    enterprise: boolean;
    gallery: boolean;
    hide: boolean;
}

export interface GetServicePrincipalOauth2PermissionScope {
    /**
     * Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
     */
    adminConsentDescription: string;
    /**
     * Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
     */
    adminConsentDisplayName: string;
    /**
     * Determines if the permission scope is enabled.
     */
    enabled: boolean;
    /**
     * The unique identifier of the delegated permission. Must be a valid UUID.
     */
    id: string;
    /**
     * Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Possible values are `User` or `Admin`.
     */
    type: string;
    /**
     * Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
     */
    userConsentDescription: string;
    /**
     * Display name for the delegated permission that appears in the end user consent experience.
     */
    userConsentDisplayName: string;
    /**
     * The value that is used for the `scp` claim in OAuth 2.0 access tokens.
     */
    value: string;
}

export interface GetServicePrincipalSamlSingleSignOn {
    /**
     * The relative URI the service provider would redirect to after completion of the single sign-on flow.
     */
    relayState: string;
}

export interface GetServicePrincipalsServicePrincipal {
    /**
     * Whether or not the service principal account is enabled.
     */
    accountEnabled: boolean;
    /**
     * Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application.
     */
    appRoleAssignmentRequired: boolean;
    /**
     * The application ID (client ID) of the application associated with this service principal.
     */
    applicationId: string;
    /**
     * The tenant ID where the associated application is registered.
     */
    applicationTenantId: string;
    /**
     * The display name of the application associated with this service principal.
     */
    displayName: string;
    /**
     * The object ID of the service principal.
     */
    objectId: string;
    /**
     * The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps.
     */
    preferredSingleSignOnMode: string;
    /**
     * The URL where the service exposes SAML metadata for federation.
     */
    samlMetadataUrl: string;
    /**
     * A list of identifier URI(s), copied over from the associated application.
     */
    servicePrincipalNames: string[];
    /**
     * The Microsoft account types that are supported for the associated application. Possible values include `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`.
     */
    signInAudience: string;
    /**
     * A list of tags applied to the service principal.
     */
    tags: string[];
    /**
     * Identifies whether the service principal represents an application or a managed identity. Possible values include `Application` or `ManagedIdentity`.
     */
    type: string;
}

export interface GetUsersUser {
    /**
     * Whether or not the account is enabled.
     */
    accountEnabled: boolean;
    /**
     * The display name of the user.
     */
    displayName: string;
    /**
     * The primary email address of the user.
     */
    mail: string;
    /**
     * The email alias of the user.
     */
    mailNickname: string;
    /**
     * The object ID of the user.
     */
    objectId: string;
    /**
     * The value used to associate an on-premises Active Directory user account with their Azure AD user object.
     */
    onpremisesImmutableId: string;
    /**
     * The on-premise SAM account name of the user.
     */
    onpremisesSamAccountName: string;
    /**
     * The on-premise user principal name of the user.
     */
    onpremisesUserPrincipalName: string;
    /**
     * The usage location of the user.
     */
    usageLocation: string;
    /**
     * The user principal name (UPN) of the user.
     */
    userPrincipalName: string;
}

export interface GroupDynamicMembership {
    /**
     * Whether rule processing is "On" (true) or "Paused" (false).
     */
    enabled: boolean;
    /**
     * The rule that determines membership of this group. For more information, see official documentation on [membership rules syntax](https://docs.microsoft.com/en-gb/azure/active-directory/enterprise-users/groups-dynamic-membership).
     */
    rule: string;
}

export interface InvitationMessage {
    /**
     * Email addresses of additional recipients the invitation message should be sent to. Only 1 additional recipient is currently supported by Azure.
     */
    additionalRecipients?: string;
    /**
     * Customized message body you want to send if you don't want to send the default message. Cannot be specified with `language`.
     */
    body?: string;
    /**
     * The language you want to send the default message in. The value specified must be in ISO 639 format. Defaults to `en-US`. Cannot be specified with `body`.
     */
    language?: string;
}

export interface NamedLocationCountry {
    /**
     * List of countries and/or regions in two-letter format specified by ISO 3166-2.
     */
    countriesAndRegions: string[];
    /**
     * Whether IP addresses that don't map to a country or region should be included in the named location. Defaults to `false`.
     */
    includeUnknownCountriesAndRegions?: boolean;
}

export interface NamedLocationIp {
    /**
     * List of IP address ranges in IPv4 CIDR format (e.g. 1.2.3.4/32) or any allowable IPv6 format from IETF RFC596.
     */
    ipRanges: string[];
    /**
     * Whether the named location is trusted. Defaults to `false`.
     */
    trusted?: boolean;
}

export interface ServicePrincipalAppRole {
    /**
     * Specifies whether this app role definition can be assigned to users and groups, or to other applications (that are accessing this application in a standalone scenario). Possible values are: `User` and `Application`, or both.
     */
    allowedMemberTypes: string[];
    /**
     * A description of the service principal provided for internal end-users.
     */
    description: string;
    /**
     * Display name for the app role that appears during app role assignment and in consent experiences.
     */
    displayName: string;
    /**
     * Specifies whether the permission scope is enabled.
     */
    enabled: boolean;
    /**
     * The unique identifier of the delegated permission.
     */
    id: string;
    /**
     * The value that is used for the `scp` claim in OAuth 2.0 access tokens.
     */
    value: string;
}

export interface ServicePrincipalFeature {
    customSingleSignOnApp?: boolean;
    enterpriseApplication?: boolean;
    galleryApplication?: boolean;
    visibleToUsers?: boolean;
}

export interface ServicePrincipalFeatureTag {
    /**
     * Whether this service principal represents a custom SAML application. Enabling this will assign the `WindowsAzureActiveDirectoryCustomSingleSignOnApplication` tag. Defaults to `false`.
     */
    customSingleSignOn?: boolean;
    /**
     * Whether this service principal represents an Enterprise Application. Enabling this will assign the `WindowsAzureActiveDirectoryIntegratedApp` tag. Defaults to `false`.
     */
    enterprise?: boolean;
    /**
     * Whether this service principal represents a gallery application. Enabling this will assign the `WindowsAzureActiveDirectoryGalleryApplicationNonPrimaryV1` tag. Defaults to `false`.
     */
    gallery?: boolean;
    /**
     * Whether this app is invisible to users in My Apps and Office 365 Launcher. Enabling this will assign the `HideApp` tag. Defaults to `false`.
     */
    hide?: boolean;
}

export interface ServicePrincipalOauth2PermissionScope {
    /**
     * Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
     */
    adminConsentDescription: string;
    /**
     * Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
     */
    adminConsentDisplayName: string;
    /**
     * Specifies whether the permission scope is enabled.
     */
    enabled: boolean;
    /**
     * The unique identifier of the delegated permission.
     */
    id: string;
    /**
     * Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Possible values are `User` or `Admin`.
     */
    type: string;
    /**
     * Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
     */
    userConsentDescription: string;
    /**
     * Display name for the delegated permission that appears in the end user consent experience.
     */
    userConsentDisplayName: string;
    /**
     * The value that is used for the `scp` claim in OAuth 2.0 access tokens.
     */
    value: string;
}

export interface ServicePrincipalSamlSingleSignOn {
    /**
     * The relative URI the service provider would redirect to after completion of the single sign-on flow.
     */
    relayState?: string;
}

export interface SynchronizationJobSchedule {
    /**
     * Date and time when this job will expire, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
     */
    expiration: string;
    /**
     * The interval between synchronization iterations ISO8601. E.g. PT40M run every 40 minutes.
     */
    interval: string;
    /**
     * State of the job.
     */
    state: string;
}

export interface SynchronizationSecretCredential {
    /**
     * The key of the secret.
     */
    key: string;
    /**
     * The value of the secret.
     */
    value: string;
}

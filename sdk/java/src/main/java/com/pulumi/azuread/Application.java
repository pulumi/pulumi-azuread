// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.ApplicationArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.ApplicationState;
import com.pulumi.azuread.outputs.ApplicationApi;
import com.pulumi.azuread.outputs.ApplicationAppRole;
import com.pulumi.azuread.outputs.ApplicationFeatureTag;
import com.pulumi.azuread.outputs.ApplicationOptionalClaims;
import com.pulumi.azuread.outputs.ApplicationPublicClient;
import com.pulumi.azuread.outputs.ApplicationRequiredResourceAccess;
import com.pulumi.azuread.outputs.ApplicationSinglePageApplication;
import com.pulumi.azuread.outputs.ApplicationWeb;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * *Create an application*
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuread.AzureadFunctions;
 * import com.pulumi.azuread.Application;
 * import com.pulumi.azuread.ApplicationArgs;
 * import com.pulumi.azuread.inputs.ApplicationApiArgs;
 * import com.pulumi.azuread.inputs.ApplicationAppRoleArgs;
 * import com.pulumi.azuread.inputs.ApplicationFeatureTagArgs;
 * import com.pulumi.azuread.inputs.ApplicationOptionalClaimsArgs;
 * import com.pulumi.azuread.inputs.ApplicationRequiredResourceAccessArgs;
 * import com.pulumi.azuread.inputs.ApplicationWebArgs;
 * import com.pulumi.azuread.inputs.ApplicationWebImplicitGrantArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var current = AzureadFunctions.getClientConfig();
 * 
 *         var example = new Application(&#34;example&#34;, ApplicationArgs.builder()        
 *             .displayName(&#34;example&#34;)
 *             .identifierUris(&#34;api://example-app&#34;)
 *             .logoImage(Base64.getEncoder().encodeToString(Files.readAllBytes(Paths.get(&#34;/path/to/logo.png&#34;))))
 *             .owners(current.applyValue(getClientConfigResult -&gt; getClientConfigResult.objectId()))
 *             .signInAudience(&#34;AzureADMultipleOrgs&#34;)
 *             .api(ApplicationApiArgs.builder()
 *                 .mappedClaimsEnabled(true)
 *                 .requestedAccessTokenVersion(2)
 *                 .knownClientApplications(                
 *                     azuread_application.known1().application_id(),
 *                     azuread_application.known2().application_id())
 *                 .oauth2PermissionScopes(                
 *                     ApplicationApiOauth2PermissionScopeArgs.builder()
 *                         .adminConsentDescription(&#34;Allow the application to access example on behalf of the signed-in user.&#34;)
 *                         .adminConsentDisplayName(&#34;Access example&#34;)
 *                         .enabled(true)
 *                         .id(&#34;96183846-204b-4b43-82e1-5d2222eb4b9b&#34;)
 *                         .type(&#34;User&#34;)
 *                         .userConsentDescription(&#34;Allow the application to access example on your behalf.&#34;)
 *                         .userConsentDisplayName(&#34;Access example&#34;)
 *                         .value(&#34;user_impersonation&#34;)
 *                         .build(),
 *                     ApplicationApiOauth2PermissionScopeArgs.builder()
 *                         .adminConsentDescription(&#34;Administer the example application&#34;)
 *                         .adminConsentDisplayName(&#34;Administer&#34;)
 *                         .enabled(true)
 *                         .id(&#34;be98fa3e-ab5b-4b11-83d9-04ba2b7946bc&#34;)
 *                         .type(&#34;Admin&#34;)
 *                         .value(&#34;administer&#34;)
 *                         .build())
 *                 .build())
 *             .appRoles(            
 *                 ApplicationAppRoleArgs.builder()
 *                     .allowedMemberTypes(                    
 *                         &#34;User&#34;,
 *                         &#34;Application&#34;)
 *                     .description(&#34;Admins can manage roles and perform all task actions&#34;)
 *                     .displayName(&#34;Admin&#34;)
 *                     .enabled(true)
 *                     .id(&#34;1b19509b-32b1-4e9f-b71d-4992aa991967&#34;)
 *                     .value(&#34;admin&#34;)
 *                     .build(),
 *                 ApplicationAppRoleArgs.builder()
 *                     .allowedMemberTypes(&#34;User&#34;)
 *                     .description(&#34;ReadOnly roles have limited query access&#34;)
 *                     .displayName(&#34;ReadOnly&#34;)
 *                     .enabled(true)
 *                     .id(&#34;497406e4-012a-4267-bf18-45a1cb148a01&#34;)
 *                     .value(&#34;User&#34;)
 *                     .build())
 *             .featureTags(ApplicationFeatureTagArgs.builder()
 *                 .enterprise(true)
 *                 .gallery(true)
 *                 .build())
 *             .optionalClaims(ApplicationOptionalClaimsArgs.builder()
 *                 .accessTokens(                
 *                     ApplicationOptionalClaimsAccessTokenArgs.builder()
 *                         .name(&#34;myclaim&#34;)
 *                         .build(),
 *                     ApplicationOptionalClaimsAccessTokenArgs.builder()
 *                         .name(&#34;otherclaim&#34;)
 *                         .build())
 *                 .idTokens(ApplicationOptionalClaimsIdTokenArgs.builder()
 *                     .name(&#34;userclaim&#34;)
 *                     .source(&#34;user&#34;)
 *                     .essential(true)
 *                     .additionalProperties(&#34;emit_as_roles&#34;)
 *                     .build())
 *                 .saml2Tokens(ApplicationOptionalClaimsSaml2TokenArgs.builder()
 *                     .name(&#34;samlexample&#34;)
 *                     .build())
 *                 .build())
 *             .requiredResourceAccesses(            
 *                 ApplicationRequiredResourceAccessArgs.builder()
 *                     .resourceAppId(&#34;00000003-0000-0000-c000-000000000000&#34;)
 *                     .resourceAccesses(                    
 *                         ApplicationRequiredResourceAccessResourceAccessArgs.builder()
 *                             .id(&#34;df021288-bdef-4463-88db-98f22de89214&#34;)
 *                             .type(&#34;Role&#34;)
 *                             .build(),
 *                         ApplicationRequiredResourceAccessResourceAccessArgs.builder()
 *                             .id(&#34;b4e74841-8e56-480b-be8b-910348b18b4c&#34;)
 *                             .type(&#34;Scope&#34;)
 *                             .build())
 *                     .build(),
 *                 ApplicationRequiredResourceAccessArgs.builder()
 *                     .resourceAppId(&#34;c5393580-f805-4401-95e8-94b7a6ef2fc2&#34;)
 *                     .resourceAccesses(ApplicationRequiredResourceAccessResourceAccessArgs.builder()
 *                         .id(&#34;594c1fb6-4f81-4475-ae41-0c394909246c&#34;)
 *                         .type(&#34;Role&#34;)
 *                         .build())
 *                     .build())
 *             .web(ApplicationWebArgs.builder()
 *                 .homepageUrl(&#34;https://app.example.net&#34;)
 *                 .logoutUrl(&#34;https://app.example.net/logout&#34;)
 *                 .redirectUris(&#34;https://app.example.net/account&#34;)
 *                 .implicitGrant(ApplicationWebImplicitGrantArgs.builder()
 *                     .accessTokenIssuanceEnabled(true)
 *                     .idTokenIssuanceEnabled(true)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * *Create application from a gallery template*
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuread.AzureadFunctions;
 * import com.pulumi.azuread.inputs.GetApplicationTemplateArgs;
 * import com.pulumi.azuread.Application;
 * import com.pulumi.azuread.ApplicationArgs;
 * import com.pulumi.azuread.ServicePrincipal;
 * import com.pulumi.azuread.ServicePrincipalArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var exampleApplicationTemplate = AzureadFunctions.getApplicationTemplate(GetApplicationTemplateArgs.builder()
 *             .displayName(&#34;Marketo&#34;)
 *             .build());
 * 
 *         var exampleApplication = new Application(&#34;exampleApplication&#34;, ApplicationArgs.builder()        
 *             .displayName(&#34;example&#34;)
 *             .templateId(exampleApplicationTemplate.applyValue(getApplicationTemplateResult -&gt; getApplicationTemplateResult.templateId()))
 *             .build());
 * 
 *         var exampleServicePrincipal = new ServicePrincipal(&#34;exampleServicePrincipal&#34;, ServicePrincipalArgs.builder()        
 *             .applicationId(exampleApplication.applicationId())
 *             .useExisting(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Applications can be imported using their object ID, e.g.
 * 
 * ```sh
 *  $ pulumi import azuread:index/application:Application test 00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuread:index/application:Application")
public class Application extends com.pulumi.resources.CustomResource {
    /**
     * An `api` block as documented below, which configures API related settings for this application.
     * 
     */
    @Export(name="api", type=ApplicationApi.class, parameters={})
    private Output</* @Nullable */ ApplicationApi> api;

    /**
     * @return An `api` block as documented below, which configures API related settings for this application.
     * 
     */
    public Output<Optional<ApplicationApi>> api() {
        return Codegen.optional(this.api);
    }
    /**
     * A mapping of app role values to app role IDs, intended to be useful when referencing app roles in other resources in your configuration.
     * 
     */
    @Export(name="appRoleIds", type=Map.class, parameters={String.class, String.class})
    private Output<Map<String,String>> appRoleIds;

    /**
     * @return A mapping of app role values to app role IDs, intended to be useful when referencing app roles in other resources in your configuration.
     * 
     */
    public Output<Map<String,String>> appRoleIds() {
        return this.appRoleIds;
    }
    /**
     * A collection of `app_role` blocks as documented below. For more information see [official documentation on Application Roles](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).
     * 
     */
    @Export(name="appRoles", type=List.class, parameters={ApplicationAppRole.class})
    private Output</* @Nullable */ List<ApplicationAppRole>> appRoles;

    /**
     * @return A collection of `app_role` blocks as documented below. For more information see [official documentation on Application Roles](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).
     * 
     */
    public Output<Optional<List<ApplicationAppRole>>> appRoles() {
        return Codegen.optional(this.appRoles);
    }
    /**
     * The Application ID (also called Client ID).
     * 
     */
    @Export(name="applicationId", type=String.class, parameters={})
    private Output<String> applicationId;

    /**
     * @return The Application ID (also called Client ID).
     * 
     */
    public Output<String> applicationId() {
        return this.applicationId;
    }
    /**
     * A description of the application, as shown to end users.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return A description of the application, as shown to end users.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Specifies whether this application supports device authentication without a user. Defaults to `false`.
     * 
     */
    @Export(name="deviceOnlyAuthEnabled", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> deviceOnlyAuthEnabled;

    /**
     * @return Specifies whether this application supports device authentication without a user. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> deviceOnlyAuthEnabled() {
        return Codegen.optional(this.deviceOnlyAuthEnabled);
    }
    /**
     * Whether Microsoft has disabled the registered application. If the application is disabled, this will be a string indicating the status/reason, e.g. `DisabledDueToViolationOfServicesAgreement`
     * 
     */
    @Export(name="disabledByMicrosoft", type=String.class, parameters={})
    private Output<String> disabledByMicrosoft;

    /**
     * @return Whether Microsoft has disabled the registered application. If the application is disabled, this will be a string indicating the status/reason, e.g. `DisabledDueToViolationOfServicesAgreement`
     * 
     */
    public Output<String> disabledByMicrosoft() {
        return this.disabledByMicrosoft;
    }
    /**
     * The display name for the application.
     * 
     */
    @Export(name="displayName", type=String.class, parameters={})
    private Output<String> displayName;

    /**
     * @return The display name for the application.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * Specifies whether the application is a public client. Appropriate for apps using token grant flows that don&#39;t use a redirect URI. Defaults to `false`.
     * 
     */
    @Export(name="fallbackPublicClientEnabled", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> fallbackPublicClientEnabled;

    /**
     * @return Specifies whether the application is a public client. Appropriate for apps using token grant flows that don&#39;t use a redirect URI. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> fallbackPublicClientEnabled() {
        return Codegen.optional(this.fallbackPublicClientEnabled);
    }
    /**
     * A `feature_tags` block as described below. Cannot be used together with the `tags` property.
     * 
     * &gt; **Features and Tags** Features are configured for an application using tags, and are provided as a shortcut to set the corresponding magic tag value for each feature. You cannot configure `feature_tags` and `tags` for an application at the same time, so if you need to assign additional custom tags it&#39;s recommended to use the `tags` property instead. Tag values also propagate to any linked service principals.
     * 
     */
    @Export(name="featureTags", type=List.class, parameters={ApplicationFeatureTag.class})
    private Output<List<ApplicationFeatureTag>> featureTags;

    /**
     * @return A `feature_tags` block as described below. Cannot be used together with the `tags` property.
     * 
     * &gt; **Features and Tags** Features are configured for an application using tags, and are provided as a shortcut to set the corresponding magic tag value for each feature. You cannot configure `feature_tags` and `tags` for an application at the same time, so if you need to assign additional custom tags it&#39;s recommended to use the `tags` property instead. Tag values also propagate to any linked service principals.
     * 
     */
    public Output<List<ApplicationFeatureTag>> featureTags() {
        return this.featureTags;
    }
    /**
     * Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
     * 
     */
    @Export(name="groupMembershipClaims", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> groupMembershipClaims;

    /**
     * @return Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
     * 
     */
    public Output<Optional<List<String>>> groupMembershipClaims() {
        return Codegen.optional(this.groupMembershipClaims);
    }
    /**
     * A set of user-defined URI(s) that uniquely identify an application within its Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
     * 
     */
    @Export(name="identifierUris", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> identifierUris;

    /**
     * @return A set of user-defined URI(s) that uniquely identify an application within its Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
     * 
     */
    public Output<Optional<List<String>>> identifierUris() {
        return Codegen.optional(this.identifierUris);
    }
    /**
     * A logo image to upload for the application, as a raw base64-encoded string. The image should be in gif, jpeg or png format. Note that once an image has been uploaded, it is not possible to remove it without replacing it with another image.
     * 
     */
    @Export(name="logoImage", type=String.class, parameters={})
    private Output</* @Nullable */ String> logoImage;

    /**
     * @return A logo image to upload for the application, as a raw base64-encoded string. The image should be in gif, jpeg or png format. Note that once an image has been uploaded, it is not possible to remove it without replacing it with another image.
     * 
     */
    public Output<Optional<String>> logoImage() {
        return Codegen.optional(this.logoImage);
    }
    /**
     * CDN URL to the application&#39;s logo, as uploaded with the `logo_image` property.
     * 
     */
    @Export(name="logoUrl", type=String.class, parameters={})
    private Output<String> logoUrl;

    /**
     * @return CDN URL to the application&#39;s logo, as uploaded with the `logo_image` property.
     * 
     */
    public Output<String> logoUrl() {
        return this.logoUrl;
    }
    /**
     * URL of the application&#39;s marketing page.
     * 
     */
    @Export(name="marketingUrl", type=String.class, parameters={})
    private Output</* @Nullable */ String> marketingUrl;

    /**
     * @return URL of the application&#39;s marketing page.
     * 
     */
    public Output<Optional<String>> marketingUrl() {
        return Codegen.optional(this.marketingUrl);
    }
    /**
     * User-specified notes relevant for the management of the application.
     * 
     */
    @Export(name="notes", type=String.class, parameters={})
    private Output</* @Nullable */ String> notes;

    /**
     * @return User-specified notes relevant for the management of the application.
     * 
     */
    public Output<Optional<String>> notes() {
        return Codegen.optional(this.notes);
    }
    /**
     * A mapping of OAuth2.0 permission scope values to scope IDs, intended to be useful when referencing permission scopes in other resources in your configuration.
     * 
     */
    @Export(name="oauth2PermissionScopeIds", type=Map.class, parameters={String.class, String.class})
    private Output<Map<String,String>> oauth2PermissionScopeIds;

    /**
     * @return A mapping of OAuth2.0 permission scope values to scope IDs, intended to be useful when referencing permission scopes in other resources in your configuration.
     * 
     */
    public Output<Map<String,String>> oauth2PermissionScopeIds() {
        return this.oauth2PermissionScopeIds;
    }
    /**
     * Specifies whether, as part of OAuth 2.0 token requests, Azure AD allows POST requests, as opposed to GET requests. Defaults to `false`, which specifies that only GET requests are allowed.
     * 
     */
    @Export(name="oauth2PostResponseRequired", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> oauth2PostResponseRequired;

    /**
     * @return Specifies whether, as part of OAuth 2.0 token requests, Azure AD allows POST requests, as opposed to GET requests. Defaults to `false`, which specifies that only GET requests are allowed.
     * 
     */
    public Output<Optional<Boolean>> oauth2PostResponseRequired() {
        return Codegen.optional(this.oauth2PostResponseRequired);
    }
    /**
     * The application&#39;s object ID.
     * 
     */
    @Export(name="objectId", type=String.class, parameters={})
    private Output<String> objectId;

    /**
     * @return The application&#39;s object ID.
     * 
     */
    public Output<String> objectId() {
        return this.objectId;
    }
    /**
     * An `optional_claims` block as documented below.
     * 
     */
    @Export(name="optionalClaims", type=ApplicationOptionalClaims.class, parameters={})
    private Output</* @Nullable */ ApplicationOptionalClaims> optionalClaims;

    /**
     * @return An `optional_claims` block as documented below.
     * 
     */
    public Output<Optional<ApplicationOptionalClaims>> optionalClaims() {
        return Codegen.optional(this.optionalClaims);
    }
    /**
     * A list of object IDs of principals that will be granted ownership of the application
     * 
     */
    @Export(name="owners", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> owners;

    /**
     * @return A list of object IDs of principals that will be granted ownership of the application
     * 
     */
    public Output<Optional<List<String>>> owners() {
        return Codegen.optional(this.owners);
    }
    /**
     * If `true`, will return an error if an existing application is found with the same name. Defaults to `false`.
     * 
     */
    @Export(name="preventDuplicateNames", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> preventDuplicateNames;

    /**
     * @return If `true`, will return an error if an existing application is found with the same name. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> preventDuplicateNames() {
        return Codegen.optional(this.preventDuplicateNames);
    }
    /**
     * URL of the application&#39;s privacy statement.
     * 
     */
    @Export(name="privacyStatementUrl", type=String.class, parameters={})
    private Output</* @Nullable */ String> privacyStatementUrl;

    /**
     * @return URL of the application&#39;s privacy statement.
     * 
     */
    public Output<Optional<String>> privacyStatementUrl() {
        return Codegen.optional(this.privacyStatementUrl);
    }
    /**
     * A `public_client` block as documented below, which configures non-web app or non-web API application settings, for example mobile or other public clients such as an installed application running on a desktop device.
     * 
     */
    @Export(name="publicClient", type=ApplicationPublicClient.class, parameters={})
    private Output</* @Nullable */ ApplicationPublicClient> publicClient;

    /**
     * @return A `public_client` block as documented below, which configures non-web app or non-web API application settings, for example mobile or other public clients such as an installed application running on a desktop device.
     * 
     */
    public Output<Optional<ApplicationPublicClient>> publicClient() {
        return Codegen.optional(this.publicClient);
    }
    /**
     * The verified publisher domain for the application.
     * 
     */
    @Export(name="publisherDomain", type=String.class, parameters={})
    private Output<String> publisherDomain;

    /**
     * @return The verified publisher domain for the application.
     * 
     */
    public Output<String> publisherDomain() {
        return this.publisherDomain;
    }
    /**
     * A collection of `required_resource_access` blocks as documented below.
     * 
     */
    @Export(name="requiredResourceAccesses", type=List.class, parameters={ApplicationRequiredResourceAccess.class})
    private Output</* @Nullable */ List<ApplicationRequiredResourceAccess>> requiredResourceAccesses;

    /**
     * @return A collection of `required_resource_access` blocks as documented below.
     * 
     */
    public Output<Optional<List<ApplicationRequiredResourceAccess>>> requiredResourceAccesses() {
        return Codegen.optional(this.requiredResourceAccesses);
    }
    /**
     * References application context information from a Service or Asset Management database.
     * 
     */
    @Export(name="serviceManagementReference", type=String.class, parameters={})
    private Output</* @Nullable */ String> serviceManagementReference;

    /**
     * @return References application context information from a Service or Asset Management database.
     * 
     */
    public Output<Optional<String>> serviceManagementReference() {
        return Codegen.optional(this.serviceManagementReference);
    }
    /**
     * The Microsoft account types that are supported for the current application. Must be one of `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`. Defaults to `AzureADMyOrg`.
     * 
     * &gt; **Changing `sign_in_audience` for existing applications** When updating an existing application to use a `sign_in_audience` value of `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`, your configuration may no longer be valid. Refer to [official documentation](https://docs.microsoft.com/en-gb/azure/active-directory/develop/supported-accounts-validation) to understand the differences in supported configurations. Where possible, the provider will attempt to validate your configuration and try to avoid applying unsupported settings to your application.
     * 
     */
    @Export(name="signInAudience", type=String.class, parameters={})
    private Output</* @Nullable */ String> signInAudience;

    /**
     * @return The Microsoft account types that are supported for the current application. Must be one of `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`. Defaults to `AzureADMyOrg`.
     * 
     * &gt; **Changing `sign_in_audience` for existing applications** When updating an existing application to use a `sign_in_audience` value of `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`, your configuration may no longer be valid. Refer to [official documentation](https://docs.microsoft.com/en-gb/azure/active-directory/develop/supported-accounts-validation) to understand the differences in supported configurations. Where possible, the provider will attempt to validate your configuration and try to avoid applying unsupported settings to your application.
     * 
     */
    public Output<Optional<String>> signInAudience() {
        return Codegen.optional(this.signInAudience);
    }
    /**
     * A `single_page_application` block as documented below, which configures single-page application (SPA) related settings for this application.
     * 
     */
    @Export(name="singlePageApplication", type=ApplicationSinglePageApplication.class, parameters={})
    private Output</* @Nullable */ ApplicationSinglePageApplication> singlePageApplication;

    /**
     * @return A `single_page_application` block as documented below, which configures single-page application (SPA) related settings for this application.
     * 
     */
    public Output<Optional<ApplicationSinglePageApplication>> singlePageApplication() {
        return Codegen.optional(this.singlePageApplication);
    }
    /**
     * URL of the application&#39;s support page.
     * 
     */
    @Export(name="supportUrl", type=String.class, parameters={})
    private Output</* @Nullable */ String> supportUrl;

    /**
     * @return URL of the application&#39;s support page.
     * 
     */
    public Output<Optional<String>> supportUrl() {
        return Codegen.optional(this.supportUrl);
    }
    /**
     * A set of tags to apply to the application for configuring specific behaviours of the application and linked service principals. Note that these are not provided for use by practitioners. Cannot be used together with the `feature_tags` block.
     * 
     * &gt; **Tags and Features** Azure Active Directory uses special tag values to configure the behavior of applications. These can be specified using either the `tags` property or with the `feature_tags` block. If you need to set any custom tag values not supported by the `feature_tags` block, it&#39;s recommended to use the `tags` property. Tag values also propagate to any linked service principals.
     * 
     */
    @Export(name="tags", type=List.class, parameters={String.class})
    private Output<List<String>> tags;

    /**
     * @return A set of tags to apply to the application for configuring specific behaviours of the application and linked service principals. Note that these are not provided for use by practitioners. Cannot be used together with the `feature_tags` block.
     * 
     * &gt; **Tags and Features** Azure Active Directory uses special tag values to configure the behavior of applications. These can be specified using either the `tags` property or with the `feature_tags` block. If you need to set any custom tag values not supported by the `feature_tags` block, it&#39;s recommended to use the `tags` property. Tag values also propagate to any linked service principals.
     * 
     */
    public Output<List<String>> tags() {
        return this.tags;
    }
    /**
     * Unique ID for a templated application in the Azure AD App Gallery, from which to create the application. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="templateId", type=String.class, parameters={})
    private Output<String> templateId;

    /**
     * @return Unique ID for a templated application in the Azure AD App Gallery, from which to create the application. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> templateId() {
        return this.templateId;
    }
    /**
     * URL of the application&#39;s terms of service statement.
     * 
     */
    @Export(name="termsOfServiceUrl", type=String.class, parameters={})
    private Output</* @Nullable */ String> termsOfServiceUrl;

    /**
     * @return URL of the application&#39;s terms of service statement.
     * 
     */
    public Output<Optional<String>> termsOfServiceUrl() {
        return Codegen.optional(this.termsOfServiceUrl);
    }
    /**
     * A `web` block as documented below, which configures web related settings for this application.
     * 
     * &gt; **Application Name Uniqueness** Application names are not unique within Azure Active Directory. Use the `prevent_duplicate_names` argument to check for existing applications if you want to avoid name collisions.
     * 
     */
    @Export(name="web", type=ApplicationWeb.class, parameters={})
    private Output</* @Nullable */ ApplicationWeb> web;

    /**
     * @return A `web` block as documented below, which configures web related settings for this application.
     * 
     * &gt; **Application Name Uniqueness** Application names are not unique within Azure Active Directory. Use the `prevent_duplicate_names` argument to check for existing applications if you want to avoid name collisions.
     * 
     */
    public Output<Optional<ApplicationWeb>> web() {
        return Codegen.optional(this.web);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Application(String name) {
        this(name, ApplicationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Application(String name, ApplicationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Application(String name, ApplicationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/application:Application", name, args == null ? ApplicationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Application(String name, Output<String> id, @Nullable ApplicationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/application:Application", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Application get(String name, Output<String> id, @Nullable ApplicationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Application(name, id, state, options);
    }
}

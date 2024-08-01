// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.ServicePrincipalArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.ServicePrincipalState;
import com.pulumi.azuread.outputs.ServicePrincipalAppRole;
import com.pulumi.azuread.outputs.ServicePrincipalFeature;
import com.pulumi.azuread.outputs.ServicePrincipalFeatureTag;
import com.pulumi.azuread.outputs.ServicePrincipalOauth2PermissionScope;
import com.pulumi.azuread.outputs.ServicePrincipalSamlSingleSignOn;
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
 * *Create a service principal for an application*
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuread.AzureadFunctions;
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
 *         final var current = AzureadFunctions.getClientConfig();
 * 
 *         var example = new Application("example", ApplicationArgs.builder()
 *             .displayName("example")
 *             .owners(current.applyValue(getClientConfigResult -> getClientConfigResult.objectId()))
 *             .build());
 * 
 *         var exampleServicePrincipal = new ServicePrincipal("exampleServicePrincipal", ServicePrincipalArgs.builder()
 *             .clientId(example.clientId())
 *             .appRoleAssignmentRequired(false)
 *             .owners(current.applyValue(getClientConfigResult -> getClientConfigResult.objectId()))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * *Create a service principal for an enterprise application*
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuread.AzureadFunctions;
 * import com.pulumi.azuread.Application;
 * import com.pulumi.azuread.ApplicationArgs;
 * import com.pulumi.azuread.ServicePrincipal;
 * import com.pulumi.azuread.ServicePrincipalArgs;
 * import com.pulumi.azuread.inputs.ServicePrincipalFeatureTagArgs;
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
 *         var example = new Application("example", ApplicationArgs.builder()
 *             .displayName("example")
 *             .owners(current.applyValue(getClientConfigResult -> getClientConfigResult.objectId()))
 *             .build());
 * 
 *         var exampleServicePrincipal = new ServicePrincipal("exampleServicePrincipal", ServicePrincipalArgs.builder()
 *             .clientId(example.clientId())
 *             .appRoleAssignmentRequired(false)
 *             .owners(current.applyValue(getClientConfigResult -> getClientConfigResult.objectId()))
 *             .featureTags(ServicePrincipalFeatureTagArgs.builder()
 *                 .enterprise(true)
 *                 .gallery(true)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * *Manage a service principal for a first-party Microsoft application*
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuread.AzureadFunctions;
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
 *         final var wellKnown = AzureadFunctions.getApplicationPublishedAppIds();
 * 
 *         var msgraph = new ServicePrincipal("msgraph", ServicePrincipalArgs.builder()
 *             .clientId(wellKnown.applyValue(getApplicationPublishedAppIdsResult -> getApplicationPublishedAppIdsResult.result().microsoftGraph()))
 *             .useExisting(true)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * *Create a service principal for an application created from a gallery template*
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
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
 *         final var example = AzureadFunctions.getApplicationTemplate(GetApplicationTemplateArgs.builder()
 *             .displayName("Marketo")
 *             .build());
 * 
 *         var exampleApplication = new Application("exampleApplication", ApplicationArgs.builder()
 *             .displayName("example")
 *             .templateId(example.applyValue(getApplicationTemplateResult -> getApplicationTemplateResult.templateId()))
 *             .build());
 * 
 *         var exampleServicePrincipal = new ServicePrincipal("exampleServicePrincipal", ServicePrincipalArgs.builder()
 *             .clientId(exampleApplication.clientId())
 *             .useExisting(true)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Service principals can be imported using their object ID, e.g.
 * 
 * ```sh
 * $ pulumi import azuread:index/servicePrincipal:ServicePrincipal example 00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuread:index/servicePrincipal:ServicePrincipal")
public class ServicePrincipal extends com.pulumi.resources.CustomResource {
    /**
     * Whether or not the service principal account is enabled. Defaults to `true`.
     * 
     */
    @Export(name="accountEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> accountEnabled;

    /**
     * @return Whether or not the service principal account is enabled. Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> accountEnabled() {
        return Codegen.optional(this.accountEnabled);
    }
    /**
     * A set of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities.
     * 
     */
    @Export(name="alternativeNames", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> alternativeNames;

    /**
     * @return A set of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities.
     * 
     */
    public Output<Optional<List<String>>> alternativeNames() {
        return Codegen.optional(this.alternativeNames);
    }
    /**
     * Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application. Defaults to `false`.
     * 
     */
    @Export(name="appRoleAssignmentRequired", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> appRoleAssignmentRequired;

    /**
     * @return Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> appRoleAssignmentRequired() {
        return Codegen.optional(this.appRoleAssignmentRequired);
    }
    /**
     * A mapping of app role values to app role IDs, as published by the associated application, intended to be useful when referencing app roles in other resources in your configuration.
     * 
     */
    @Export(name="appRoleIds", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> appRoleIds;

    /**
     * @return A mapping of app role values to app role IDs, as published by the associated application, intended to be useful when referencing app roles in other resources in your configuration.
     * 
     */
    public Output<Map<String,String>> appRoleIds() {
        return this.appRoleIds;
    }
    /**
     * A list of app roles published by the associated application, as documented below. For more information [official documentation](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).
     * 
     */
    @Export(name="appRoles", refs={List.class,ServicePrincipalAppRole.class}, tree="[0,1]")
    private Output<List<ServicePrincipalAppRole>> appRoles;

    /**
     * @return A list of app roles published by the associated application, as documented below. For more information [official documentation](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).
     * 
     */
    public Output<List<ServicePrincipalAppRole>> appRoles() {
        return this.appRoles;
    }
    /**
     * The application ID (client ID) of the application for which to create a service principal
     * 
     * @deprecated
     * The `application_id` property has been replaced with the `client_id` property and will be removed in version 3.0 of the AzureAD provider
     * 
     */
    @Deprecated /* The `application_id` property has been replaced with the `client_id` property and will be removed in version 3.0 of the AzureAD provider */
    @Export(name="applicationId", refs={String.class}, tree="[0]")
    private Output<String> applicationId;

    /**
     * @return The application ID (client ID) of the application for which to create a service principal
     * 
     */
    public Output<String> applicationId() {
        return this.applicationId;
    }
    /**
     * The tenant ID where the associated application is registered.
     * 
     */
    @Export(name="applicationTenantId", refs={String.class}, tree="[0]")
    private Output<String> applicationTenantId;

    /**
     * @return The tenant ID where the associated application is registered.
     * 
     */
    public Output<String> applicationTenantId() {
        return this.applicationTenantId;
    }
    /**
     * The client ID of the application for which to create a service principal.
     * 
     */
    @Export(name="clientId", refs={String.class}, tree="[0]")
    private Output<String> clientId;

    /**
     * @return The client ID of the application for which to create a service principal.
     * 
     */
    public Output<String> clientId() {
        return this.clientId;
    }
    /**
     * A description of the service principal provided for internal end-users.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A description of the service principal provided for internal end-users.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Display name for the app role that appears during app role assignment and in consent experiences.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    /**
     * @return Display name for the app role that appears during app role assignment and in consent experiences.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * A `feature_tags` block as described below. Cannot be used together with the `tags` property.
     * 
     * &gt; **Features and Tags** Features are configured for a service principal using tags, and are provided as a shortcut to set the corresponding magic tag value for each feature. You cannot configure `feature_tags` and `tags` for a service principal at the same time, so if you need to assign additional custom tags it&#39;s recommended to use the `tags` property instead. Any tags configured for the linked application will propagate to this service principal.
     * 
     */
    @Export(name="featureTags", refs={List.class,ServicePrincipalFeatureTag.class}, tree="[0,1]")
    private Output<List<ServicePrincipalFeatureTag>> featureTags;

    /**
     * @return A `feature_tags` block as described below. Cannot be used together with the `tags` property.
     * 
     * &gt; **Features and Tags** Features are configured for a service principal using tags, and are provided as a shortcut to set the corresponding magic tag value for each feature. You cannot configure `feature_tags` and `tags` for a service principal at the same time, so if you need to assign additional custom tags it&#39;s recommended to use the `tags` property instead. Any tags configured for the linked application will propagate to this service principal.
     * 
     */
    public Output<List<ServicePrincipalFeatureTag>> featureTags() {
        return this.featureTags;
    }
    /**
     * Block of features to configure for this service principal using tags
     * 
     * @deprecated
     * This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider
     * 
     */
    @Deprecated /* This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider */
    @Export(name="features", refs={List.class,ServicePrincipalFeature.class}, tree="[0,1]")
    private Output<List<ServicePrincipalFeature>> features;

    /**
     * @return Block of features to configure for this service principal using tags
     * 
     */
    public Output<List<ServicePrincipalFeature>> features() {
        return this.features;
    }
    /**
     * Home page or landing page of the associated application.
     * 
     */
    @Export(name="homepageUrl", refs={String.class}, tree="[0]")
    private Output<String> homepageUrl;

    /**
     * @return Home page or landing page of the associated application.
     * 
     */
    public Output<String> homepageUrl() {
        return this.homepageUrl;
    }
    /**
     * The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on.
     * 
     */
    @Export(name="loginUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> loginUrl;

    /**
     * @return The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on.
     * 
     */
    public Output<Optional<String>> loginUrl() {
        return Codegen.optional(this.loginUrl);
    }
    /**
     * The URL that will be used by Microsoft&#39;s authorization service to log out an user using OpenId Connect front-channel, back-channel or SAML logout protocols, taken from the associated application.
     * 
     */
    @Export(name="logoutUrl", refs={String.class}, tree="[0]")
    private Output<String> logoutUrl;

    /**
     * @return The URL that will be used by Microsoft&#39;s authorization service to log out an user using OpenId Connect front-channel, back-channel or SAML logout protocols, taken from the associated application.
     * 
     */
    public Output<String> logoutUrl() {
        return this.logoutUrl;
    }
    /**
     * A free text field to capture information about the service principal, typically used for operational purposes.
     * 
     */
    @Export(name="notes", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> notes;

    /**
     * @return A free text field to capture information about the service principal, typically used for operational purposes.
     * 
     */
    public Output<Optional<String>> notes() {
        return Codegen.optional(this.notes);
    }
    /**
     * A set of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
     * 
     */
    @Export(name="notificationEmailAddresses", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> notificationEmailAddresses;

    /**
     * @return A set of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
     * 
     */
    public Output<Optional<List<String>>> notificationEmailAddresses() {
        return Codegen.optional(this.notificationEmailAddresses);
    }
    /**
     * A mapping of OAuth2.0 permission scope values to scope IDs, as exposed by the associated application, intended to be useful when referencing permission scopes in other resources in your configuration.
     * 
     */
    @Export(name="oauth2PermissionScopeIds", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> oauth2PermissionScopeIds;

    /**
     * @return A mapping of OAuth2.0 permission scope values to scope IDs, as exposed by the associated application, intended to be useful when referencing permission scopes in other resources in your configuration.
     * 
     */
    public Output<Map<String,String>> oauth2PermissionScopeIds() {
        return this.oauth2PermissionScopeIds;
    }
    /**
     * A list of OAuth 2.0 delegated permission scopes exposed by the associated application, as documented below.
     * 
     */
    @Export(name="oauth2PermissionScopes", refs={List.class,ServicePrincipalOauth2PermissionScope.class}, tree="[0,1]")
    private Output<List<ServicePrincipalOauth2PermissionScope>> oauth2PermissionScopes;

    /**
     * @return A list of OAuth 2.0 delegated permission scopes exposed by the associated application, as documented below.
     * 
     */
    public Output<List<ServicePrincipalOauth2PermissionScope>> oauth2PermissionScopes() {
        return this.oauth2PermissionScopes;
    }
    /**
     * The object ID of the service principal.
     * 
     */
    @Export(name="objectId", refs={String.class}, tree="[0]")
    private Output<String> objectId;

    /**
     * @return The object ID of the service principal.
     * 
     */
    public Output<String> objectId() {
        return this.objectId;
    }
    /**
     * A list of object IDs of principals that will be granted ownership of the service principal
     * 
     */
    @Export(name="owners", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> owners;

    /**
     * @return A list of object IDs of principals that will be granted ownership of the service principal
     * 
     */
    public Output<Optional<List<String>>> owners() {
        return Codegen.optional(this.owners);
    }
    /**
     * The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps. Supported values are `oidc`, `password`, `saml` or `notSupported`. Omit this property or specify a blank string to unset.
     * 
     */
    @Export(name="preferredSingleSignOnMode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> preferredSingleSignOnMode;

    /**
     * @return The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps. Supported values are `oidc`, `password`, `saml` or `notSupported`. Omit this property or specify a blank string to unset.
     * 
     */
    public Output<Optional<String>> preferredSingleSignOnMode() {
        return Codegen.optional(this.preferredSingleSignOnMode);
    }
    /**
     * A list of URLs where user tokens are sent for sign-in with the associated application, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent for the associated application.
     * 
     */
    @Export(name="redirectUris", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> redirectUris;

    /**
     * @return A list of URLs where user tokens are sent for sign-in with the associated application, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent for the associated application.
     * 
     */
    public Output<List<String>> redirectUris() {
        return this.redirectUris;
    }
    /**
     * The URL where the service exposes SAML metadata for federation.
     * 
     */
    @Export(name="samlMetadataUrl", refs={String.class}, tree="[0]")
    private Output<String> samlMetadataUrl;

    /**
     * @return The URL where the service exposes SAML metadata for federation.
     * 
     */
    public Output<String> samlMetadataUrl() {
        return this.samlMetadataUrl;
    }
    /**
     * A `saml_single_sign_on` block as documented below.
     * 
     */
    @Export(name="samlSingleSignOn", refs={ServicePrincipalSamlSingleSignOn.class}, tree="[0]")
    private Output</* @Nullable */ ServicePrincipalSamlSingleSignOn> samlSingleSignOn;

    /**
     * @return A `saml_single_sign_on` block as documented below.
     * 
     */
    public Output<Optional<ServicePrincipalSamlSingleSignOn>> samlSingleSignOn() {
        return Codegen.optional(this.samlSingleSignOn);
    }
    /**
     * A list of identifier URI(s), copied over from the associated application.
     * 
     */
    @Export(name="servicePrincipalNames", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> servicePrincipalNames;

    /**
     * @return A list of identifier URI(s), copied over from the associated application.
     * 
     */
    public Output<List<String>> servicePrincipalNames() {
        return this.servicePrincipalNames;
    }
    /**
     * The Microsoft account types that are supported for the associated application. Possible values include `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`.
     * 
     */
    @Export(name="signInAudience", refs={String.class}, tree="[0]")
    private Output<String> signInAudience;

    /**
     * @return The Microsoft account types that are supported for the associated application. Possible values include `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`.
     * 
     */
    public Output<String> signInAudience() {
        return this.signInAudience;
    }
    /**
     * A set of tags to apply to the service principal for configuring specific behaviours of the service principal. Note that these are not provided for use by practitioners. Cannot be used together with the `feature_tags` block.
     * 
     * &gt; **Tags and Features** Azure Active Directory uses special tag values to configure the behavior of service principals. These can be specified using either the `tags` property or with the `feature_tags` block. If you need to set any custom tag values not supported by the `feature_tags` block, it&#39;s recommended to use the `tags` property. Tag values set for the linked application will also propagate to this service principal.
     * 
     */
    @Export(name="tags", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> tags;

    /**
     * @return A set of tags to apply to the service principal for configuring specific behaviours of the service principal. Note that these are not provided for use by practitioners. Cannot be used together with the `feature_tags` block.
     * 
     * &gt; **Tags and Features** Azure Active Directory uses special tag values to configure the behavior of service principals. These can be specified using either the `tags` property or with the `feature_tags` block. If you need to set any custom tag values not supported by the `feature_tags` block, it&#39;s recommended to use the `tags` property. Tag values set for the linked application will also propagate to this service principal.
     * 
     */
    public Output<List<String>> tags() {
        return this.tags;
    }
    /**
     * Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Possible values are `User` or `Admin`.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Possible values are `User` or `Admin`.
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * When true, the resource will return an existing service principal instead of failing with an error
     * 
     */
    @Export(name="useExisting", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> useExisting;

    /**
     * @return When true, the resource will return an existing service principal instead of failing with an error
     * 
     */
    public Output<Optional<Boolean>> useExisting() {
        return Codegen.optional(this.useExisting);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServicePrincipal(String name) {
        this(name, ServicePrincipalArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServicePrincipal(String name, @Nullable ServicePrincipalArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServicePrincipal(String name, @Nullable ServicePrincipalArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/servicePrincipal:ServicePrincipal", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()));
    }

    private ServicePrincipal(String name, Output<String> id, @Nullable ServicePrincipalState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/servicePrincipal:ServicePrincipal", name, state, makeResourceOptions(options, id));
    }

    private static ServicePrincipalArgs makeArgs(@Nullable ServicePrincipalArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ServicePrincipalArgs.Empty : args;
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
    public static ServicePrincipal get(String name, Output<String> id, @Nullable ServicePrincipalState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServicePrincipal(name, id, state, options);
    }
}

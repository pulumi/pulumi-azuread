// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Manages a service principal associated with an application within Azure Active Directory.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this resource.
 *
 * When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.All` or `Directory.ReadWrite.All`
 *
 * It is not currently possible to manage service principals whilst having only the `Application.ReadWrite.OwnedBy` role granted.
 *
 * When authenticated with a user principal, this resource requires one of the following directory roles: `Application Administrator` or `Global Administrator`
 *
 * ## Example Usage
 *
 * *Create a service principal for an application*
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const current = azuread.getClientConfig({});
 * const exampleApplication = new azuread.Application("exampleApplication", {
 *     displayName: "example",
 *     owners: [current.then(current => current.objectId)],
 * });
 * const exampleServicePrincipal = new azuread.ServicePrincipal("exampleServicePrincipal", {
 *     applicationId: exampleApplication.applicationId,
 *     appRoleAssignmentRequired: false,
 *     owners: [current.then(current => current.objectId)],
 * });
 * ```
 *
 * *Create a service principal for an enterprise application*
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const current = azuread.getClientConfig({});
 * const exampleApplication = new azuread.Application("exampleApplication", {
 *     displayName: "example",
 *     owners: [current.then(current => current.objectId)],
 * });
 * const exampleServicePrincipal = new azuread.ServicePrincipal("exampleServicePrincipal", {
 *     applicationId: exampleApplication.applicationId,
 *     appRoleAssignmentRequired: false,
 *     owners: [current.then(current => current.objectId)],
 *     featureTags: [{
 *         enterprise: true,
 *         gallery: true,
 *     }],
 * });
 * ```
 *
 * *Manage a service principal for a first-party Microsoft application*
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const wellKnown = azuread.getApplicationPublishedAppIds({});
 * const msgraph = new azuread.ServicePrincipal("msgraph", {
 *     applicationId: wellKnown.then(wellKnown => wellKnown.result?.MicrosoftGraph),
 *     useExisting: true,
 * });
 * ```
 *
 * *Create a service principal for an application created from a gallery template*
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const exampleApplicationTemplate = azuread.getApplicationTemplate({
 *     displayName: "Marketo",
 * });
 * const exampleApplication = new azuread.Application("exampleApplication", {
 *     displayName: "example",
 *     templateId: exampleApplicationTemplate.then(exampleApplicationTemplate => exampleApplicationTemplate.templateId),
 * });
 * const exampleServicePrincipal = new azuread.ServicePrincipal("exampleServicePrincipal", {
 *     applicationId: exampleApplication.applicationId,
 *     useExisting: true,
 * });
 * ```
 *
 * ## Import
 *
 * Service principals can be imported using their object ID, e.g.
 *
 * ```sh
 *  $ pulumi import azuread:index/servicePrincipal:ServicePrincipal test 00000000-0000-0000-0000-000000000000
 * ```
 */
export class ServicePrincipal extends pulumi.CustomResource {
    /**
     * Get an existing ServicePrincipal resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServicePrincipalState, opts?: pulumi.CustomResourceOptions): ServicePrincipal {
        return new ServicePrincipal(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/servicePrincipal:ServicePrincipal';

    /**
     * Returns true if the given object is an instance of ServicePrincipal.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServicePrincipal {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServicePrincipal.__pulumiType;
    }

    /**
     * Whether or not the service principal account is enabled. Defaults to `true`.
     */
    public readonly accountEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * A set of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities.
     */
    public readonly alternativeNames!: pulumi.Output<string[] | undefined>;
    /**
     * Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application. Defaults to `false`.
     */
    public readonly appRoleAssignmentRequired!: pulumi.Output<boolean | undefined>;
    /**
     * A mapping of app role values to app role IDs, as published by the associated application, intended to be useful when referencing app roles in other resources in your configuration.
     */
    public /*out*/ readonly appRoleIds!: pulumi.Output<{[key: string]: string}>;
    /**
     * A list of app roles published by the associated application, as documented below. For more information [official documentation](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).
     */
    public /*out*/ readonly appRoles!: pulumi.Output<outputs.ServicePrincipalAppRole[]>;
    /**
     * The application ID (client ID) of the application for which to create a service principal.
     */
    public readonly applicationId!: pulumi.Output<string>;
    /**
     * The tenant ID where the associated application is registered.
     */
    public /*out*/ readonly applicationTenantId!: pulumi.Output<string>;
    /**
     * A description of the service principal provided for internal end-users.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Display name for the app role that appears during app role assignment and in consent experiences.
     */
    public /*out*/ readonly displayName!: pulumi.Output<string>;
    /**
     * A `featureTags` block as described below. Cannot be used together with the `tags` property.
     */
    public readonly featureTags!: pulumi.Output<outputs.ServicePrincipalFeatureTag[]>;
    /**
     * Block of features to configure for this service principal using tags
     *
     * @deprecated This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider
     */
    public readonly features!: pulumi.Output<outputs.ServicePrincipalFeature[]>;
    /**
     * Home page or landing page of the associated application.
     */
    public /*out*/ readonly homepageUrl!: pulumi.Output<string>;
    /**
     * The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on.
     */
    public readonly loginUrl!: pulumi.Output<string | undefined>;
    /**
     * The URL that will be used by Microsoft's authorization service to log out an user using OpenId Connect front-channel, back-channel or SAML logout protocols, taken from the associated application.
     */
    public /*out*/ readonly logoutUrl!: pulumi.Output<string>;
    /**
     * A free text field to capture information about the service principal, typically used for operational purposes.
     */
    public readonly notes!: pulumi.Output<string | undefined>;
    /**
     * A set of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
     */
    public readonly notificationEmailAddresses!: pulumi.Output<string[] | undefined>;
    /**
     * A mapping of OAuth2.0 permission scope values to scope IDs, as exposed by the associated application, intended to be useful when referencing permission scopes in other resources in your configuration.
     */
    public /*out*/ readonly oauth2PermissionScopeIds!: pulumi.Output<{[key: string]: string}>;
    /**
     * A list of OAuth 2.0 delegated permission scopes exposed by the associated application, as documented below.
     */
    public /*out*/ readonly oauth2PermissionScopes!: pulumi.Output<outputs.ServicePrincipalOauth2PermissionScope[]>;
    /**
     * The object ID of the service principal.
     */
    public /*out*/ readonly objectId!: pulumi.Output<string>;
    /**
     * A set of object IDs of principals that will be granted ownership of the service principal. Supported object types are users or service principals. By default, no owners are assigned.
     */
    public readonly owners!: pulumi.Output<string[] | undefined>;
    /**
     * The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps. Supported values are `oidc`, `password`, `saml` or `notSupported`. Omit this property or specify a blank string to unset.
     */
    public readonly preferredSingleSignOnMode!: pulumi.Output<string | undefined>;
    /**
     * A list of URLs where user tokens are sent for sign-in with the associated application, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent for the associated application.
     */
    public /*out*/ readonly redirectUris!: pulumi.Output<string[]>;
    /**
     * The URL where the service exposes SAML metadata for federation.
     */
    public /*out*/ readonly samlMetadataUrl!: pulumi.Output<string>;
    /**
     * A `samlSingleSignOn` block as documented below.
     */
    public readonly samlSingleSignOn!: pulumi.Output<outputs.ServicePrincipalSamlSingleSignOn | undefined>;
    /**
     * A list of identifier URI(s), copied over from the associated application.
     */
    public /*out*/ readonly servicePrincipalNames!: pulumi.Output<string[]>;
    /**
     * The Microsoft account types that are supported for the associated application. Possible values include `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`.
     */
    public /*out*/ readonly signInAudience!: pulumi.Output<string>;
    /**
     * A set of tags to apply to the service principal. Cannot be used together with the `featureTags` block.
     */
    public readonly tags!: pulumi.Output<string[]>;
    /**
     * Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Possible values are `User` or `Admin`.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * When true, any existing service principal linked to the same application will be automatically imported. When false, an import error will be raised for any pre-existing service principal.
     */
    public readonly useExisting!: pulumi.Output<boolean | undefined>;

    /**
     * Create a ServicePrincipal resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServicePrincipalArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServicePrincipalArgs | ServicePrincipalState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServicePrincipalState | undefined;
            resourceInputs["accountEnabled"] = state ? state.accountEnabled : undefined;
            resourceInputs["alternativeNames"] = state ? state.alternativeNames : undefined;
            resourceInputs["appRoleAssignmentRequired"] = state ? state.appRoleAssignmentRequired : undefined;
            resourceInputs["appRoleIds"] = state ? state.appRoleIds : undefined;
            resourceInputs["appRoles"] = state ? state.appRoles : undefined;
            resourceInputs["applicationId"] = state ? state.applicationId : undefined;
            resourceInputs["applicationTenantId"] = state ? state.applicationTenantId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["featureTags"] = state ? state.featureTags : undefined;
            resourceInputs["features"] = state ? state.features : undefined;
            resourceInputs["homepageUrl"] = state ? state.homepageUrl : undefined;
            resourceInputs["loginUrl"] = state ? state.loginUrl : undefined;
            resourceInputs["logoutUrl"] = state ? state.logoutUrl : undefined;
            resourceInputs["notes"] = state ? state.notes : undefined;
            resourceInputs["notificationEmailAddresses"] = state ? state.notificationEmailAddresses : undefined;
            resourceInputs["oauth2PermissionScopeIds"] = state ? state.oauth2PermissionScopeIds : undefined;
            resourceInputs["oauth2PermissionScopes"] = state ? state.oauth2PermissionScopes : undefined;
            resourceInputs["objectId"] = state ? state.objectId : undefined;
            resourceInputs["owners"] = state ? state.owners : undefined;
            resourceInputs["preferredSingleSignOnMode"] = state ? state.preferredSingleSignOnMode : undefined;
            resourceInputs["redirectUris"] = state ? state.redirectUris : undefined;
            resourceInputs["samlMetadataUrl"] = state ? state.samlMetadataUrl : undefined;
            resourceInputs["samlSingleSignOn"] = state ? state.samlSingleSignOn : undefined;
            resourceInputs["servicePrincipalNames"] = state ? state.servicePrincipalNames : undefined;
            resourceInputs["signInAudience"] = state ? state.signInAudience : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["useExisting"] = state ? state.useExisting : undefined;
        } else {
            const args = argsOrState as ServicePrincipalArgs | undefined;
            if ((!args || args.applicationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationId'");
            }
            resourceInputs["accountEnabled"] = args ? args.accountEnabled : undefined;
            resourceInputs["alternativeNames"] = args ? args.alternativeNames : undefined;
            resourceInputs["appRoleAssignmentRequired"] = args ? args.appRoleAssignmentRequired : undefined;
            resourceInputs["applicationId"] = args ? args.applicationId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["featureTags"] = args ? args.featureTags : undefined;
            resourceInputs["features"] = args ? args.features : undefined;
            resourceInputs["loginUrl"] = args ? args.loginUrl : undefined;
            resourceInputs["notes"] = args ? args.notes : undefined;
            resourceInputs["notificationEmailAddresses"] = args ? args.notificationEmailAddresses : undefined;
            resourceInputs["owners"] = args ? args.owners : undefined;
            resourceInputs["preferredSingleSignOnMode"] = args ? args.preferredSingleSignOnMode : undefined;
            resourceInputs["samlSingleSignOn"] = args ? args.samlSingleSignOn : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["useExisting"] = args ? args.useExisting : undefined;
            resourceInputs["appRoleIds"] = undefined /*out*/;
            resourceInputs["appRoles"] = undefined /*out*/;
            resourceInputs["applicationTenantId"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["homepageUrl"] = undefined /*out*/;
            resourceInputs["logoutUrl"] = undefined /*out*/;
            resourceInputs["oauth2PermissionScopeIds"] = undefined /*out*/;
            resourceInputs["oauth2PermissionScopes"] = undefined /*out*/;
            resourceInputs["objectId"] = undefined /*out*/;
            resourceInputs["redirectUris"] = undefined /*out*/;
            resourceInputs["samlMetadataUrl"] = undefined /*out*/;
            resourceInputs["servicePrincipalNames"] = undefined /*out*/;
            resourceInputs["signInAudience"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ServicePrincipal.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServicePrincipal resources.
 */
export interface ServicePrincipalState {
    /**
     * Whether or not the service principal account is enabled. Defaults to `true`.
     */
    accountEnabled?: pulumi.Input<boolean>;
    /**
     * A set of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities.
     */
    alternativeNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application. Defaults to `false`.
     */
    appRoleAssignmentRequired?: pulumi.Input<boolean>;
    /**
     * A mapping of app role values to app role IDs, as published by the associated application, intended to be useful when referencing app roles in other resources in your configuration.
     */
    appRoleIds?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A list of app roles published by the associated application, as documented below. For more information [official documentation](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).
     */
    appRoles?: pulumi.Input<pulumi.Input<inputs.ServicePrincipalAppRole>[]>;
    /**
     * The application ID (client ID) of the application for which to create a service principal.
     */
    applicationId?: pulumi.Input<string>;
    /**
     * The tenant ID where the associated application is registered.
     */
    applicationTenantId?: pulumi.Input<string>;
    /**
     * A description of the service principal provided for internal end-users.
     */
    description?: pulumi.Input<string>;
    /**
     * Display name for the app role that appears during app role assignment and in consent experiences.
     */
    displayName?: pulumi.Input<string>;
    /**
     * A `featureTags` block as described below. Cannot be used together with the `tags` property.
     */
    featureTags?: pulumi.Input<pulumi.Input<inputs.ServicePrincipalFeatureTag>[]>;
    /**
     * Block of features to configure for this service principal using tags
     *
     * @deprecated This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider
     */
    features?: pulumi.Input<pulumi.Input<inputs.ServicePrincipalFeature>[]>;
    /**
     * Home page or landing page of the associated application.
     */
    homepageUrl?: pulumi.Input<string>;
    /**
     * The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on.
     */
    loginUrl?: pulumi.Input<string>;
    /**
     * The URL that will be used by Microsoft's authorization service to log out an user using OpenId Connect front-channel, back-channel or SAML logout protocols, taken from the associated application.
     */
    logoutUrl?: pulumi.Input<string>;
    /**
     * A free text field to capture information about the service principal, typically used for operational purposes.
     */
    notes?: pulumi.Input<string>;
    /**
     * A set of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
     */
    notificationEmailAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A mapping of OAuth2.0 permission scope values to scope IDs, as exposed by the associated application, intended to be useful when referencing permission scopes in other resources in your configuration.
     */
    oauth2PermissionScopeIds?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A list of OAuth 2.0 delegated permission scopes exposed by the associated application, as documented below.
     */
    oauth2PermissionScopes?: pulumi.Input<pulumi.Input<inputs.ServicePrincipalOauth2PermissionScope>[]>;
    /**
     * The object ID of the service principal.
     */
    objectId?: pulumi.Input<string>;
    /**
     * A set of object IDs of principals that will be granted ownership of the service principal. Supported object types are users or service principals. By default, no owners are assigned.
     */
    owners?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps. Supported values are `oidc`, `password`, `saml` or `notSupported`. Omit this property or specify a blank string to unset.
     */
    preferredSingleSignOnMode?: pulumi.Input<string>;
    /**
     * A list of URLs where user tokens are sent for sign-in with the associated application, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent for the associated application.
     */
    redirectUris?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The URL where the service exposes SAML metadata for federation.
     */
    samlMetadataUrl?: pulumi.Input<string>;
    /**
     * A `samlSingleSignOn` block as documented below.
     */
    samlSingleSignOn?: pulumi.Input<inputs.ServicePrincipalSamlSingleSignOn>;
    /**
     * A list of identifier URI(s), copied over from the associated application.
     */
    servicePrincipalNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Microsoft account types that are supported for the associated application. Possible values include `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`.
     */
    signInAudience?: pulumi.Input<string>;
    /**
     * A set of tags to apply to the service principal. Cannot be used together with the `featureTags` block.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Possible values are `User` or `Admin`.
     */
    type?: pulumi.Input<string>;
    /**
     * When true, any existing service principal linked to the same application will be automatically imported. When false, an import error will be raised for any pre-existing service principal.
     */
    useExisting?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a ServicePrincipal resource.
 */
export interface ServicePrincipalArgs {
    /**
     * Whether or not the service principal account is enabled. Defaults to `true`.
     */
    accountEnabled?: pulumi.Input<boolean>;
    /**
     * A set of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities.
     */
    alternativeNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application. Defaults to `false`.
     */
    appRoleAssignmentRequired?: pulumi.Input<boolean>;
    /**
     * The application ID (client ID) of the application for which to create a service principal.
     */
    applicationId: pulumi.Input<string>;
    /**
     * A description of the service principal provided for internal end-users.
     */
    description?: pulumi.Input<string>;
    /**
     * A `featureTags` block as described below. Cannot be used together with the `tags` property.
     */
    featureTags?: pulumi.Input<pulumi.Input<inputs.ServicePrincipalFeatureTag>[]>;
    /**
     * Block of features to configure for this service principal using tags
     *
     * @deprecated This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider
     */
    features?: pulumi.Input<pulumi.Input<inputs.ServicePrincipalFeature>[]>;
    /**
     * The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on.
     */
    loginUrl?: pulumi.Input<string>;
    /**
     * A free text field to capture information about the service principal, typically used for operational purposes.
     */
    notes?: pulumi.Input<string>;
    /**
     * A set of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
     */
    notificationEmailAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A set of object IDs of principals that will be granted ownership of the service principal. Supported object types are users or service principals. By default, no owners are assigned.
     */
    owners?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps. Supported values are `oidc`, `password`, `saml` or `notSupported`. Omit this property or specify a blank string to unset.
     */
    preferredSingleSignOnMode?: pulumi.Input<string>;
    /**
     * A `samlSingleSignOn` block as documented below.
     */
    samlSingleSignOn?: pulumi.Input<inputs.ServicePrincipalSamlSingleSignOn>;
    /**
     * A set of tags to apply to the service principal. Cannot be used together with the `featureTags` block.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * When true, any existing service principal linked to the same application will be automatically imported. When false, an import error will be raised for any pre-existing service principal.
     */
    useExisting?: pulumi.Input<boolean>;
}

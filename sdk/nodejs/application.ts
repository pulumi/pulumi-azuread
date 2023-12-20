// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * *Create an application*
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 * import * as fs from "fs";
 *
 * const current = azuread.getClientConfig({});
 * const example = new azuread.Application("example", {
 *     displayName: "example",
 *     identifierUris: ["api://example-app"],
 *     logoImage: Buffer.from(fs.readFileSync("/path/to/logo.png", 'binary')).toString('base64'),
 *     owners: [current.then(current => current.objectId)],
 *     signInAudience: "AzureADMultipleOrgs",
 *     api: {
 *         mappedClaimsEnabled: true,
 *         requestedAccessTokenVersion: 2,
 *         knownClientApplications: [
 *             azuread_application.known1.application_id,
 *             azuread_application.known2.application_id,
 *         ],
 *         oauth2PermissionScopes: [
 *             {
 *                 adminConsentDescription: "Allow the application to access example on behalf of the signed-in user.",
 *                 adminConsentDisplayName: "Access example",
 *                 enabled: true,
 *                 id: "96183846-204b-4b43-82e1-5d2222eb4b9b",
 *                 type: "User",
 *                 userConsentDescription: "Allow the application to access example on your behalf.",
 *                 userConsentDisplayName: "Access example",
 *                 value: "user_impersonation",
 *             },
 *             {
 *                 adminConsentDescription: "Administer the example application",
 *                 adminConsentDisplayName: "Administer",
 *                 enabled: true,
 *                 id: "be98fa3e-ab5b-4b11-83d9-04ba2b7946bc",
 *                 type: "Admin",
 *                 value: "administer",
 *             },
 *         ],
 *     },
 *     appRoles: [
 *         {
 *             allowedMemberTypes: [
 *                 "User",
 *                 "Application",
 *             ],
 *             description: "Admins can manage roles and perform all task actions",
 *             displayName: "Admin",
 *             enabled: true,
 *             id: "1b19509b-32b1-4e9f-b71d-4992aa991967",
 *             value: "admin",
 *         },
 *         {
 *             allowedMemberTypes: ["User"],
 *             description: "ReadOnly roles have limited query access",
 *             displayName: "ReadOnly",
 *             enabled: true,
 *             id: "497406e4-012a-4267-bf18-45a1cb148a01",
 *             value: "User",
 *         },
 *     ],
 *     featureTags: [{
 *         enterprise: true,
 *         gallery: true,
 *     }],
 *     optionalClaims: {
 *         accessTokens: [
 *             {
 *                 name: "myclaim",
 *             },
 *             {
 *                 name: "otherclaim",
 *             },
 *         ],
 *         idTokens: [{
 *             name: "userclaim",
 *             source: "user",
 *             essential: true,
 *             additionalProperties: ["emit_as_roles"],
 *         }],
 *         saml2Tokens: [{
 *             name: "samlexample",
 *         }],
 *     },
 *     requiredResourceAccesses: [
 *         {
 *             resourceAppId: "00000003-0000-0000-c000-000000000000",
 *             resourceAccesses: [
 *                 {
 *                     id: "df021288-bdef-4463-88db-98f22de89214",
 *                     type: "Role",
 *                 },
 *                 {
 *                     id: "b4e74841-8e56-480b-be8b-910348b18b4c",
 *                     type: "Scope",
 *                 },
 *             ],
 *         },
 *         {
 *             resourceAppId: "c5393580-f805-4401-95e8-94b7a6ef2fc2",
 *             resourceAccesses: [{
 *                 id: "594c1fb6-4f81-4475-ae41-0c394909246c",
 *                 type: "Role",
 *             }],
 *         },
 *     ],
 *     web: {
 *         homepageUrl: "https://app.example.net",
 *         logoutUrl: "https://app.example.net/logout",
 *         redirectUris: ["https://app.example.net/account"],
 *         implicitGrant: {
 *             accessTokenIssuanceEnabled: true,
 *             idTokenIssuanceEnabled: true,
 *         },
 *     },
 * });
 * ```
 *
 * *Create application from a gallery template*
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
 * Applications can be imported using the object ID of the application, in the following format.
 *
 * ```sh
 *  $ pulumi import azuread:index/application:Application example /applications/00000000-0000-0000-0000-000000000000
 * ```
 */
export class Application extends pulumi.CustomResource {
    /**
     * Get an existing Application resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApplicationState, opts?: pulumi.CustomResourceOptions): Application {
        return new Application(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/application:Application';

    /**
     * Returns true if the given object is an instance of Application.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Application {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Application.__pulumiType;
    }

    /**
     * An `api` block as documented below, which configures API related settings for this application.
     */
    public readonly api!: pulumi.Output<outputs.ApplicationApi | undefined>;
    /**
     * A mapping of app role values to app role IDs, intended to be useful when referencing app roles in other resources in your configuration.
     */
    public /*out*/ readonly appRoleIds!: pulumi.Output<{[key: string]: string}>;
    /**
     * A collection of `appRole` blocks as documented below. For more information see [official documentation on Application Roles](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).
     */
    public readonly appRoles!: pulumi.Output<outputs.ApplicationAppRole[] | undefined>;
    /**
     * The Application ID (also called Client ID)
     *
     * @deprecated The `application_id` attribute has been replaced by the `client_id` attribute and will be removed in version 3.0 of the AzureAD provider
     */
    public /*out*/ readonly applicationId!: pulumi.Output<string>;
    /**
     * The Client ID for the application.
     */
    public /*out*/ readonly clientId!: pulumi.Output<string>;
    /**
     * A description of the application, as shown to end users.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether this application supports device authentication without a user. Defaults to `false`.
     */
    public readonly deviceOnlyAuthEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Whether Microsoft has disabled the registered application. If the application is disabled, this will be a string indicating the status/reason, e.g. `DisabledDueToViolationOfServicesAgreement`
     */
    public /*out*/ readonly disabledByMicrosoft!: pulumi.Output<string>;
    /**
     * The display name for the application.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Specifies whether the application is a public client. Appropriate for apps using token grant flows that don't use a redirect URI. Defaults to `false`.
     */
    public readonly fallbackPublicClientEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * A `featureTags` block as described below. Cannot be used together with the `tags` property.
     *
     * > **Features and Tags** Features are configured for an application using tags, and are provided as a shortcut to set the corresponding magic tag value for each feature. You cannot configure `featureTags` and `tags` for an application at the same time, so if you need to assign additional custom tags it's recommended to use the `tags` property instead. Tag values also propagate to any linked service principals.
     */
    public readonly featureTags!: pulumi.Output<outputs.ApplicationFeatureTag[]>;
    /**
     * Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
     */
    public readonly groupMembershipClaims!: pulumi.Output<string[] | undefined>;
    /**
     * A set of user-defined URI(s) that uniquely identify an application within its Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
     */
    public readonly identifierUris!: pulumi.Output<string[] | undefined>;
    /**
     * A logo image to upload for the application, as a raw base64-encoded string. The image should be in gif, jpeg or png format. Note that once an image has been uploaded, it is not possible to remove it without replacing it with another image.
     */
    public readonly logoImage!: pulumi.Output<string | undefined>;
    /**
     * CDN URL to the application's logo, as uploaded with the `logoImage` property.
     */
    public /*out*/ readonly logoUrl!: pulumi.Output<string>;
    /**
     * URL of the application's marketing page.
     */
    public readonly marketingUrl!: pulumi.Output<string | undefined>;
    /**
     * User-specified notes relevant for the management of the application.
     */
    public readonly notes!: pulumi.Output<string | undefined>;
    /**
     * A mapping of OAuth2.0 permission scope values to scope IDs, intended to be useful when referencing permission scopes in other resources in your configuration.
     */
    public /*out*/ readonly oauth2PermissionScopeIds!: pulumi.Output<{[key: string]: string}>;
    /**
     * Specifies whether, as part of OAuth 2.0 token requests, Azure AD allows POST requests, as opposed to GET requests. Defaults to `false`, which specifies that only GET requests are allowed.
     */
    public readonly oauth2PostResponseRequired!: pulumi.Output<boolean | undefined>;
    /**
     * The application's object ID.
     */
    public /*out*/ readonly objectId!: pulumi.Output<string>;
    /**
     * An `optionalClaims` block as documented below.
     */
    public readonly optionalClaims!: pulumi.Output<outputs.ApplicationOptionalClaims | undefined>;
    /**
     * A list of object IDs of principals that will be granted ownership of the application
     */
    public readonly owners!: pulumi.Output<string[] | undefined>;
    /**
     * If `true`, will return an error if an existing application is found with the same name. Defaults to `false`.
     */
    public readonly preventDuplicateNames!: pulumi.Output<boolean | undefined>;
    /**
     * URL of the application's privacy statement.
     */
    public readonly privacyStatementUrl!: pulumi.Output<string | undefined>;
    /**
     * A `publicClient` block as documented below, which configures non-web app or non-web API application settings, for example mobile or other public clients such as an installed application running on a desktop device.
     */
    public readonly publicClient!: pulumi.Output<outputs.ApplicationPublicClient | undefined>;
    /**
     * The verified publisher domain for the application.
     */
    public /*out*/ readonly publisherDomain!: pulumi.Output<string>;
    /**
     * A collection of `requiredResourceAccess` blocks as documented below.
     */
    public readonly requiredResourceAccesses!: pulumi.Output<outputs.ApplicationRequiredResourceAccess[] | undefined>;
    /**
     * References application context information from a Service or Asset Management database.
     */
    public readonly serviceManagementReference!: pulumi.Output<string | undefined>;
    /**
     * The Microsoft account types that are supported for the current application. Must be one of `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`. Defaults to `AzureADMyOrg`.
     *
     * > **Changing `signInAudience` for existing applications** When updating an existing application to use a `signInAudience` value of `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`, your configuration may no longer be valid. Refer to [official documentation](https://docs.microsoft.com/en-gb/azure/active-directory/develop/supported-accounts-validation) to understand the differences in supported configurations. Where possible, the provider will attempt to validate your configuration and try to avoid applying unsupported settings to your application.
     */
    public readonly signInAudience!: pulumi.Output<string | undefined>;
    /**
     * A `singlePageApplication` block as documented below, which configures single-page application (SPA) related settings for this application.
     */
    public readonly singlePageApplication!: pulumi.Output<outputs.ApplicationSinglePageApplication | undefined>;
    /**
     * URL of the application's support page.
     */
    public readonly supportUrl!: pulumi.Output<string | undefined>;
    /**
     * A set of tags to apply to the application for configuring specific behaviours of the application and linked service principals. Note that these are not provided for use by practitioners. Cannot be used together with the `featureTags` block.
     *
     * > **Tags and Features** Azure Active Directory uses special tag values to configure the behavior of applications. These can be specified using either the `tags` property or with the `featureTags` block. If you need to set any custom tag values not supported by the `featureTags` block, it's recommended to use the `tags` property. Tag values also propagate to any linked service principals.
     */
    public readonly tags!: pulumi.Output<string[]>;
    /**
     * Unique ID for a templated application in the Azure AD App Gallery, from which to create the application. Changing this forces a new resource to be created.
     *
     * > **Tip for Gallery Applications** This resource can  be used to instantiate a gallery application, however it will also attempt to manage the properties of the resulting application. If this is not desired, consider using the azuread.ApplicationRegistration resource instead.
     */
    public readonly templateId!: pulumi.Output<string>;
    /**
     * URL of the application's terms of service statement.
     */
    public readonly termsOfServiceUrl!: pulumi.Output<string | undefined>;
    /**
     * A `web` block as documented below, which configures web related settings for this application.
     *
     * > **Application Name Uniqueness** Application names are not unique within Azure Active Directory. Use the `preventDuplicateNames` argument to check for existing applications if you want to avoid name collisions.
     */
    public readonly web!: pulumi.Output<outputs.ApplicationWeb | undefined>;

    /**
     * Create a Application resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationArgs | ApplicationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApplicationState | undefined;
            resourceInputs["api"] = state ? state.api : undefined;
            resourceInputs["appRoleIds"] = state ? state.appRoleIds : undefined;
            resourceInputs["appRoles"] = state ? state.appRoles : undefined;
            resourceInputs["applicationId"] = state ? state.applicationId : undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["deviceOnlyAuthEnabled"] = state ? state.deviceOnlyAuthEnabled : undefined;
            resourceInputs["disabledByMicrosoft"] = state ? state.disabledByMicrosoft : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["fallbackPublicClientEnabled"] = state ? state.fallbackPublicClientEnabled : undefined;
            resourceInputs["featureTags"] = state ? state.featureTags : undefined;
            resourceInputs["groupMembershipClaims"] = state ? state.groupMembershipClaims : undefined;
            resourceInputs["identifierUris"] = state ? state.identifierUris : undefined;
            resourceInputs["logoImage"] = state ? state.logoImage : undefined;
            resourceInputs["logoUrl"] = state ? state.logoUrl : undefined;
            resourceInputs["marketingUrl"] = state ? state.marketingUrl : undefined;
            resourceInputs["notes"] = state ? state.notes : undefined;
            resourceInputs["oauth2PermissionScopeIds"] = state ? state.oauth2PermissionScopeIds : undefined;
            resourceInputs["oauth2PostResponseRequired"] = state ? state.oauth2PostResponseRequired : undefined;
            resourceInputs["objectId"] = state ? state.objectId : undefined;
            resourceInputs["optionalClaims"] = state ? state.optionalClaims : undefined;
            resourceInputs["owners"] = state ? state.owners : undefined;
            resourceInputs["preventDuplicateNames"] = state ? state.preventDuplicateNames : undefined;
            resourceInputs["privacyStatementUrl"] = state ? state.privacyStatementUrl : undefined;
            resourceInputs["publicClient"] = state ? state.publicClient : undefined;
            resourceInputs["publisherDomain"] = state ? state.publisherDomain : undefined;
            resourceInputs["requiredResourceAccesses"] = state ? state.requiredResourceAccesses : undefined;
            resourceInputs["serviceManagementReference"] = state ? state.serviceManagementReference : undefined;
            resourceInputs["signInAudience"] = state ? state.signInAudience : undefined;
            resourceInputs["singlePageApplication"] = state ? state.singlePageApplication : undefined;
            resourceInputs["supportUrl"] = state ? state.supportUrl : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["templateId"] = state ? state.templateId : undefined;
            resourceInputs["termsOfServiceUrl"] = state ? state.termsOfServiceUrl : undefined;
            resourceInputs["web"] = state ? state.web : undefined;
        } else {
            const args = argsOrState as ApplicationArgs | undefined;
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            resourceInputs["api"] = args ? args.api : undefined;
            resourceInputs["appRoles"] = args ? args.appRoles : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["deviceOnlyAuthEnabled"] = args ? args.deviceOnlyAuthEnabled : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["fallbackPublicClientEnabled"] = args ? args.fallbackPublicClientEnabled : undefined;
            resourceInputs["featureTags"] = args ? args.featureTags : undefined;
            resourceInputs["groupMembershipClaims"] = args ? args.groupMembershipClaims : undefined;
            resourceInputs["identifierUris"] = args ? args.identifierUris : undefined;
            resourceInputs["logoImage"] = args ? args.logoImage : undefined;
            resourceInputs["marketingUrl"] = args ? args.marketingUrl : undefined;
            resourceInputs["notes"] = args ? args.notes : undefined;
            resourceInputs["oauth2PostResponseRequired"] = args ? args.oauth2PostResponseRequired : undefined;
            resourceInputs["optionalClaims"] = args ? args.optionalClaims : undefined;
            resourceInputs["owners"] = args ? args.owners : undefined;
            resourceInputs["preventDuplicateNames"] = args ? args.preventDuplicateNames : undefined;
            resourceInputs["privacyStatementUrl"] = args ? args.privacyStatementUrl : undefined;
            resourceInputs["publicClient"] = args ? args.publicClient : undefined;
            resourceInputs["requiredResourceAccesses"] = args ? args.requiredResourceAccesses : undefined;
            resourceInputs["serviceManagementReference"] = args ? args.serviceManagementReference : undefined;
            resourceInputs["signInAudience"] = args ? args.signInAudience : undefined;
            resourceInputs["singlePageApplication"] = args ? args.singlePageApplication : undefined;
            resourceInputs["supportUrl"] = args ? args.supportUrl : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["templateId"] = args ? args.templateId : undefined;
            resourceInputs["termsOfServiceUrl"] = args ? args.termsOfServiceUrl : undefined;
            resourceInputs["web"] = args ? args.web : undefined;
            resourceInputs["appRoleIds"] = undefined /*out*/;
            resourceInputs["applicationId"] = undefined /*out*/;
            resourceInputs["clientId"] = undefined /*out*/;
            resourceInputs["disabledByMicrosoft"] = undefined /*out*/;
            resourceInputs["logoUrl"] = undefined /*out*/;
            resourceInputs["oauth2PermissionScopeIds"] = undefined /*out*/;
            resourceInputs["objectId"] = undefined /*out*/;
            resourceInputs["publisherDomain"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Application.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Application resources.
 */
export interface ApplicationState {
    /**
     * An `api` block as documented below, which configures API related settings for this application.
     */
    api?: pulumi.Input<inputs.ApplicationApi>;
    /**
     * A mapping of app role values to app role IDs, intended to be useful when referencing app roles in other resources in your configuration.
     */
    appRoleIds?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A collection of `appRole` blocks as documented below. For more information see [official documentation on Application Roles](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).
     */
    appRoles?: pulumi.Input<pulumi.Input<inputs.ApplicationAppRole>[]>;
    /**
     * The Application ID (also called Client ID)
     *
     * @deprecated The `application_id` attribute has been replaced by the `client_id` attribute and will be removed in version 3.0 of the AzureAD provider
     */
    applicationId?: pulumi.Input<string>;
    /**
     * The Client ID for the application.
     */
    clientId?: pulumi.Input<string>;
    /**
     * A description of the application, as shown to end users.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies whether this application supports device authentication without a user. Defaults to `false`.
     */
    deviceOnlyAuthEnabled?: pulumi.Input<boolean>;
    /**
     * Whether Microsoft has disabled the registered application. If the application is disabled, this will be a string indicating the status/reason, e.g. `DisabledDueToViolationOfServicesAgreement`
     */
    disabledByMicrosoft?: pulumi.Input<string>;
    /**
     * The display name for the application.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Specifies whether the application is a public client. Appropriate for apps using token grant flows that don't use a redirect URI. Defaults to `false`.
     */
    fallbackPublicClientEnabled?: pulumi.Input<boolean>;
    /**
     * A `featureTags` block as described below. Cannot be used together with the `tags` property.
     *
     * > **Features and Tags** Features are configured for an application using tags, and are provided as a shortcut to set the corresponding magic tag value for each feature. You cannot configure `featureTags` and `tags` for an application at the same time, so if you need to assign additional custom tags it's recommended to use the `tags` property instead. Tag values also propagate to any linked service principals.
     */
    featureTags?: pulumi.Input<pulumi.Input<inputs.ApplicationFeatureTag>[]>;
    /**
     * Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
     */
    groupMembershipClaims?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A set of user-defined URI(s) that uniquely identify an application within its Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
     */
    identifierUris?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A logo image to upload for the application, as a raw base64-encoded string. The image should be in gif, jpeg or png format. Note that once an image has been uploaded, it is not possible to remove it without replacing it with another image.
     */
    logoImage?: pulumi.Input<string>;
    /**
     * CDN URL to the application's logo, as uploaded with the `logoImage` property.
     */
    logoUrl?: pulumi.Input<string>;
    /**
     * URL of the application's marketing page.
     */
    marketingUrl?: pulumi.Input<string>;
    /**
     * User-specified notes relevant for the management of the application.
     */
    notes?: pulumi.Input<string>;
    /**
     * A mapping of OAuth2.0 permission scope values to scope IDs, intended to be useful when referencing permission scopes in other resources in your configuration.
     */
    oauth2PermissionScopeIds?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies whether, as part of OAuth 2.0 token requests, Azure AD allows POST requests, as opposed to GET requests. Defaults to `false`, which specifies that only GET requests are allowed.
     */
    oauth2PostResponseRequired?: pulumi.Input<boolean>;
    /**
     * The application's object ID.
     */
    objectId?: pulumi.Input<string>;
    /**
     * An `optionalClaims` block as documented below.
     */
    optionalClaims?: pulumi.Input<inputs.ApplicationOptionalClaims>;
    /**
     * A list of object IDs of principals that will be granted ownership of the application
     */
    owners?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If `true`, will return an error if an existing application is found with the same name. Defaults to `false`.
     */
    preventDuplicateNames?: pulumi.Input<boolean>;
    /**
     * URL of the application's privacy statement.
     */
    privacyStatementUrl?: pulumi.Input<string>;
    /**
     * A `publicClient` block as documented below, which configures non-web app or non-web API application settings, for example mobile or other public clients such as an installed application running on a desktop device.
     */
    publicClient?: pulumi.Input<inputs.ApplicationPublicClient>;
    /**
     * The verified publisher domain for the application.
     */
    publisherDomain?: pulumi.Input<string>;
    /**
     * A collection of `requiredResourceAccess` blocks as documented below.
     */
    requiredResourceAccesses?: pulumi.Input<pulumi.Input<inputs.ApplicationRequiredResourceAccess>[]>;
    /**
     * References application context information from a Service or Asset Management database.
     */
    serviceManagementReference?: pulumi.Input<string>;
    /**
     * The Microsoft account types that are supported for the current application. Must be one of `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`. Defaults to `AzureADMyOrg`.
     *
     * > **Changing `signInAudience` for existing applications** When updating an existing application to use a `signInAudience` value of `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`, your configuration may no longer be valid. Refer to [official documentation](https://docs.microsoft.com/en-gb/azure/active-directory/develop/supported-accounts-validation) to understand the differences in supported configurations. Where possible, the provider will attempt to validate your configuration and try to avoid applying unsupported settings to your application.
     */
    signInAudience?: pulumi.Input<string>;
    /**
     * A `singlePageApplication` block as documented below, which configures single-page application (SPA) related settings for this application.
     */
    singlePageApplication?: pulumi.Input<inputs.ApplicationSinglePageApplication>;
    /**
     * URL of the application's support page.
     */
    supportUrl?: pulumi.Input<string>;
    /**
     * A set of tags to apply to the application for configuring specific behaviours of the application and linked service principals. Note that these are not provided for use by practitioners. Cannot be used together with the `featureTags` block.
     *
     * > **Tags and Features** Azure Active Directory uses special tag values to configure the behavior of applications. These can be specified using either the `tags` property or with the `featureTags` block. If you need to set any custom tag values not supported by the `featureTags` block, it's recommended to use the `tags` property. Tag values also propagate to any linked service principals.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Unique ID for a templated application in the Azure AD App Gallery, from which to create the application. Changing this forces a new resource to be created.
     *
     * > **Tip for Gallery Applications** This resource can  be used to instantiate a gallery application, however it will also attempt to manage the properties of the resulting application. If this is not desired, consider using the azuread.ApplicationRegistration resource instead.
     */
    templateId?: pulumi.Input<string>;
    /**
     * URL of the application's terms of service statement.
     */
    termsOfServiceUrl?: pulumi.Input<string>;
    /**
     * A `web` block as documented below, which configures web related settings for this application.
     *
     * > **Application Name Uniqueness** Application names are not unique within Azure Active Directory. Use the `preventDuplicateNames` argument to check for existing applications if you want to avoid name collisions.
     */
    web?: pulumi.Input<inputs.ApplicationWeb>;
}

/**
 * The set of arguments for constructing a Application resource.
 */
export interface ApplicationArgs {
    /**
     * An `api` block as documented below, which configures API related settings for this application.
     */
    api?: pulumi.Input<inputs.ApplicationApi>;
    /**
     * A collection of `appRole` blocks as documented below. For more information see [official documentation on Application Roles](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).
     */
    appRoles?: pulumi.Input<pulumi.Input<inputs.ApplicationAppRole>[]>;
    /**
     * A description of the application, as shown to end users.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies whether this application supports device authentication without a user. Defaults to `false`.
     */
    deviceOnlyAuthEnabled?: pulumi.Input<boolean>;
    /**
     * The display name for the application.
     */
    displayName: pulumi.Input<string>;
    /**
     * Specifies whether the application is a public client. Appropriate for apps using token grant flows that don't use a redirect URI. Defaults to `false`.
     */
    fallbackPublicClientEnabled?: pulumi.Input<boolean>;
    /**
     * A `featureTags` block as described below. Cannot be used together with the `tags` property.
     *
     * > **Features and Tags** Features are configured for an application using tags, and are provided as a shortcut to set the corresponding magic tag value for each feature. You cannot configure `featureTags` and `tags` for an application at the same time, so if you need to assign additional custom tags it's recommended to use the `tags` property instead. Tag values also propagate to any linked service principals.
     */
    featureTags?: pulumi.Input<pulumi.Input<inputs.ApplicationFeatureTag>[]>;
    /**
     * Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
     */
    groupMembershipClaims?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A set of user-defined URI(s) that uniquely identify an application within its Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
     */
    identifierUris?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A logo image to upload for the application, as a raw base64-encoded string. The image should be in gif, jpeg or png format. Note that once an image has been uploaded, it is not possible to remove it without replacing it with another image.
     */
    logoImage?: pulumi.Input<string>;
    /**
     * URL of the application's marketing page.
     */
    marketingUrl?: pulumi.Input<string>;
    /**
     * User-specified notes relevant for the management of the application.
     */
    notes?: pulumi.Input<string>;
    /**
     * Specifies whether, as part of OAuth 2.0 token requests, Azure AD allows POST requests, as opposed to GET requests. Defaults to `false`, which specifies that only GET requests are allowed.
     */
    oauth2PostResponseRequired?: pulumi.Input<boolean>;
    /**
     * An `optionalClaims` block as documented below.
     */
    optionalClaims?: pulumi.Input<inputs.ApplicationOptionalClaims>;
    /**
     * A list of object IDs of principals that will be granted ownership of the application
     */
    owners?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If `true`, will return an error if an existing application is found with the same name. Defaults to `false`.
     */
    preventDuplicateNames?: pulumi.Input<boolean>;
    /**
     * URL of the application's privacy statement.
     */
    privacyStatementUrl?: pulumi.Input<string>;
    /**
     * A `publicClient` block as documented below, which configures non-web app or non-web API application settings, for example mobile or other public clients such as an installed application running on a desktop device.
     */
    publicClient?: pulumi.Input<inputs.ApplicationPublicClient>;
    /**
     * A collection of `requiredResourceAccess` blocks as documented below.
     */
    requiredResourceAccesses?: pulumi.Input<pulumi.Input<inputs.ApplicationRequiredResourceAccess>[]>;
    /**
     * References application context information from a Service or Asset Management database.
     */
    serviceManagementReference?: pulumi.Input<string>;
    /**
     * The Microsoft account types that are supported for the current application. Must be one of `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`. Defaults to `AzureADMyOrg`.
     *
     * > **Changing `signInAudience` for existing applications** When updating an existing application to use a `signInAudience` value of `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`, your configuration may no longer be valid. Refer to [official documentation](https://docs.microsoft.com/en-gb/azure/active-directory/develop/supported-accounts-validation) to understand the differences in supported configurations. Where possible, the provider will attempt to validate your configuration and try to avoid applying unsupported settings to your application.
     */
    signInAudience?: pulumi.Input<string>;
    /**
     * A `singlePageApplication` block as documented below, which configures single-page application (SPA) related settings for this application.
     */
    singlePageApplication?: pulumi.Input<inputs.ApplicationSinglePageApplication>;
    /**
     * URL of the application's support page.
     */
    supportUrl?: pulumi.Input<string>;
    /**
     * A set of tags to apply to the application for configuring specific behaviours of the application and linked service principals. Note that these are not provided for use by practitioners. Cannot be used together with the `featureTags` block.
     *
     * > **Tags and Features** Azure Active Directory uses special tag values to configure the behavior of applications. These can be specified using either the `tags` property or with the `featureTags` block. If you need to set any custom tag values not supported by the `featureTags` block, it's recommended to use the `tags` property. Tag values also propagate to any linked service principals.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Unique ID for a templated application in the Azure AD App Gallery, from which to create the application. Changing this forces a new resource to be created.
     *
     * > **Tip for Gallery Applications** This resource can  be used to instantiate a gallery application, however it will also attempt to manage the properties of the resulting application. If this is not desired, consider using the azuread.ApplicationRegistration resource instead.
     */
    templateId?: pulumi.Input<string>;
    /**
     * URL of the application's terms of service statement.
     */
    termsOfServiceUrl?: pulumi.Input<string>;
    /**
     * A `web` block as documented below, which configures web related settings for this application.
     *
     * > **Application Name Uniqueness** Application names are not unique within Azure Active Directory. Use the `preventDuplicateNames` argument to check for existing applications if you want to avoid name collisions.
     */
    web?: pulumi.Input<inputs.ApplicationWeb>;
}

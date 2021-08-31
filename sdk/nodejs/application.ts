// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * ## Import
 *
 * Applications can be imported using their object ID, e.g.
 *
 * ```sh
 *  $ pulumi import azuread:index/application:Application test 00000000-0000-0000-0000-000000000000
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
     * The Application ID (also called Client ID).
     */
    public /*out*/ readonly applicationId!: pulumi.Output<string>;
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
     * Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
     */
    public readonly groupMembershipClaims!: pulumi.Output<string[] | undefined>;
    /**
     * A set of user-defined URI(s) that uniquely identify an application within its Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
     */
    public readonly identifierUris!: pulumi.Output<string[] | undefined>;
    /**
     * CDN URL to the application's logo.
     */
    public /*out*/ readonly logoUrl!: pulumi.Output<string>;
    /**
     * URL of the application's marketing page.
     */
    public readonly marketingUrl!: pulumi.Output<string | undefined>;
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
     * A set of object IDs of principals that will be granted ownership of the application. Supported object types are users or service principals. By default, no owners are assigned.
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
     * The Microsoft account types that are supported for the current application. Must be one of `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`. Defaults to `AzureADMyOrg`.
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
     * URL of the application's terms of service statement.
     */
    public readonly termsOfServiceUrl!: pulumi.Output<string | undefined>;
    /**
     * A `web` block as documented below, which configures web related settings for this application.
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
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApplicationState | undefined;
            inputs["api"] = state ? state.api : undefined;
            inputs["appRoleIds"] = state ? state.appRoleIds : undefined;
            inputs["appRoles"] = state ? state.appRoles : undefined;
            inputs["applicationId"] = state ? state.applicationId : undefined;
            inputs["deviceOnlyAuthEnabled"] = state ? state.deviceOnlyAuthEnabled : undefined;
            inputs["disabledByMicrosoft"] = state ? state.disabledByMicrosoft : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["fallbackPublicClientEnabled"] = state ? state.fallbackPublicClientEnabled : undefined;
            inputs["groupMembershipClaims"] = state ? state.groupMembershipClaims : undefined;
            inputs["identifierUris"] = state ? state.identifierUris : undefined;
            inputs["logoUrl"] = state ? state.logoUrl : undefined;
            inputs["marketingUrl"] = state ? state.marketingUrl : undefined;
            inputs["oauth2PermissionScopeIds"] = state ? state.oauth2PermissionScopeIds : undefined;
            inputs["oauth2PostResponseRequired"] = state ? state.oauth2PostResponseRequired : undefined;
            inputs["objectId"] = state ? state.objectId : undefined;
            inputs["optionalClaims"] = state ? state.optionalClaims : undefined;
            inputs["owners"] = state ? state.owners : undefined;
            inputs["preventDuplicateNames"] = state ? state.preventDuplicateNames : undefined;
            inputs["privacyStatementUrl"] = state ? state.privacyStatementUrl : undefined;
            inputs["publicClient"] = state ? state.publicClient : undefined;
            inputs["publisherDomain"] = state ? state.publisherDomain : undefined;
            inputs["requiredResourceAccesses"] = state ? state.requiredResourceAccesses : undefined;
            inputs["signInAudience"] = state ? state.signInAudience : undefined;
            inputs["singlePageApplication"] = state ? state.singlePageApplication : undefined;
            inputs["supportUrl"] = state ? state.supportUrl : undefined;
            inputs["termsOfServiceUrl"] = state ? state.termsOfServiceUrl : undefined;
            inputs["web"] = state ? state.web : undefined;
        } else {
            const args = argsOrState as ApplicationArgs | undefined;
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            inputs["api"] = args ? args.api : undefined;
            inputs["appRoles"] = args ? args.appRoles : undefined;
            inputs["deviceOnlyAuthEnabled"] = args ? args.deviceOnlyAuthEnabled : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["fallbackPublicClientEnabled"] = args ? args.fallbackPublicClientEnabled : undefined;
            inputs["groupMembershipClaims"] = args ? args.groupMembershipClaims : undefined;
            inputs["identifierUris"] = args ? args.identifierUris : undefined;
            inputs["marketingUrl"] = args ? args.marketingUrl : undefined;
            inputs["oauth2PostResponseRequired"] = args ? args.oauth2PostResponseRequired : undefined;
            inputs["optionalClaims"] = args ? args.optionalClaims : undefined;
            inputs["owners"] = args ? args.owners : undefined;
            inputs["preventDuplicateNames"] = args ? args.preventDuplicateNames : undefined;
            inputs["privacyStatementUrl"] = args ? args.privacyStatementUrl : undefined;
            inputs["publicClient"] = args ? args.publicClient : undefined;
            inputs["requiredResourceAccesses"] = args ? args.requiredResourceAccesses : undefined;
            inputs["signInAudience"] = args ? args.signInAudience : undefined;
            inputs["singlePageApplication"] = args ? args.singlePageApplication : undefined;
            inputs["supportUrl"] = args ? args.supportUrl : undefined;
            inputs["termsOfServiceUrl"] = args ? args.termsOfServiceUrl : undefined;
            inputs["web"] = args ? args.web : undefined;
            inputs["appRoleIds"] = undefined /*out*/;
            inputs["applicationId"] = undefined /*out*/;
            inputs["disabledByMicrosoft"] = undefined /*out*/;
            inputs["logoUrl"] = undefined /*out*/;
            inputs["oauth2PermissionScopeIds"] = undefined /*out*/;
            inputs["objectId"] = undefined /*out*/;
            inputs["publisherDomain"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Application.__pulumiType, name, inputs, opts);
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
     * The Application ID (also called Client ID).
     */
    applicationId?: pulumi.Input<string>;
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
     * Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
     */
    groupMembershipClaims?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A set of user-defined URI(s) that uniquely identify an application within its Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
     */
    identifierUris?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * CDN URL to the application's logo.
     */
    logoUrl?: pulumi.Input<string>;
    /**
     * URL of the application's marketing page.
     */
    marketingUrl?: pulumi.Input<string>;
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
     * A set of object IDs of principals that will be granted ownership of the application. Supported object types are users or service principals. By default, no owners are assigned.
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
     * The Microsoft account types that are supported for the current application. Must be one of `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`. Defaults to `AzureADMyOrg`.
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
     * URL of the application's terms of service statement.
     */
    termsOfServiceUrl?: pulumi.Input<string>;
    /**
     * A `web` block as documented below, which configures web related settings for this application.
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
     * Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
     */
    groupMembershipClaims?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A set of user-defined URI(s) that uniquely identify an application within its Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
     */
    identifierUris?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * URL of the application's marketing page.
     */
    marketingUrl?: pulumi.Input<string>;
    /**
     * Specifies whether, as part of OAuth 2.0 token requests, Azure AD allows POST requests, as opposed to GET requests. Defaults to `false`, which specifies that only GET requests are allowed.
     */
    oauth2PostResponseRequired?: pulumi.Input<boolean>;
    /**
     * An `optionalClaims` block as documented below.
     */
    optionalClaims?: pulumi.Input<inputs.ApplicationOptionalClaims>;
    /**
     * A set of object IDs of principals that will be granted ownership of the application. Supported object types are users or service principals. By default, no owners are assigned.
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
     * The Microsoft account types that are supported for the current application. Must be one of `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`. Defaults to `AzureADMyOrg`.
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
     * URL of the application's terms of service statement.
     */
    termsOfServiceUrl?: pulumi.Input<string>;
    /**
     * A `web` block as documented below, which configures web related settings for this application.
     */
    web?: pulumi.Input<inputs.ApplicationWeb>;
}

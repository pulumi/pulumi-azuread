// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Manages an Application within Azure Active Directory.
 *
 * > **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write owned by applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = new azuread.Application("example", {
 *     appRoles: [{
 *         allowedMemberTypes: [
 *             "User",
 *             "Application",
 *         ],
 *         description: "Admins can manage roles and perform all task actions",
 *         displayName: "Admin",
 *         isEnabled: true,
 *         value: "Admin",
 *     }],
 *     availableToOtherTenants: false,
 *     displayName: "example",
 *     homepage: "https://homepage",
 *     identifierUris: ["https://uri"],
 *     oauth2AllowImplicitFlow: true,
 *     oauth2Permissions: [
 *         {
 *             adminConsentDescription: "Allow the application to access example on behalf of the signed-in user.",
 *             adminConsentDisplayName: "Access example",
 *             isEnabled: true,
 *             type: "User",
 *             userConsentDescription: "Allow the application to access example on your behalf.",
 *             userConsentDisplayName: "Access example",
 *             value: "user_impersonation",
 *         },
 *         {
 *             adminConsentDescription: "Administer the example application",
 *             adminConsentDisplayName: "Administer",
 *             isEnabled: true,
 *             type: "Admin",
 *             value: "administer",
 *         },
 *     ],
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
 *             additionalProperties: ["emit_as_roles"],
 *             essential: true,
 *             name: "userclaim",
 *             source: "user",
 *         }],
 *     },
 *     owners: ["00000004-0000-0000-c000-000000000000"],
 *     replyUrls: ["https://replyurl"],
 *     requiredResourceAccesses: [
 *         {
 *             resourceAccesses: [
 *                 {
 *                     id: "...",
 *                     type: "Role",
 *                 },
 *                 {
 *                     id: "...",
 *                     type: "Scope",
 *                 },
 *                 {
 *                     id: "...",
 *                     type: "Scope",
 *                 },
 *             ],
 *             resourceAppId: "00000003-0000-0000-c000-000000000000",
 *         },
 *         {
 *             resourceAccesses: [{
 *                 id: "...",
 *                 type: "Scope",
 *             }],
 *             resourceAppId: "00000002-0000-0000-c000-000000000000",
 *         },
 *     ],
 *     type: "webapp/api",
 * });
 * ```
 *
 * ## Import
 *
 * Azure Active Directory Applications can be imported using the `object id`, e.g.
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
     * A collection of `appRole` blocks as documented below. For more information https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles
     */
    public readonly appRoles!: pulumi.Output<outputs.ApplicationAppRole[]>;
    /**
     * The Application ID (Client ID).
     */
    public /*out*/ readonly applicationId!: pulumi.Output<string>;
    /**
     * Is this Azure AD Application available to other tenants? Defaults to `false`.
     */
    public readonly availableToOtherTenants!: pulumi.Output<boolean | undefined>;
    /**
     * The display name for the application.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects. Defaults to `SecurityGroup`. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
     */
    public readonly groupMembershipClaims!: pulumi.Output<string | undefined>;
    /**
     * The URL to the application's home page.
     */
    public readonly homepage!: pulumi.Output<string>;
    /**
     * A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
     */
    public readonly identifierUris!: pulumi.Output<string[]>;
    /**
     * The URL of the logout page.
     */
    public readonly logoutUrl!: pulumi.Output<string | undefined>;
    /**
     * The name of the optional claim.
     *
     * @deprecated This property has been renamed to `display_name` and will be removed in version 2.0 of this provider.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Does this Azure AD Application allow OAuth2.0 implicit flow tokens? Defaults to `false`.
     */
    public readonly oauth2AllowImplicitFlow!: pulumi.Output<boolean | undefined>;
    /**
     * A collection of OAuth 2.0 permission scopes that the web API (resource) app exposes to client apps. Each permission is covered by `oauth2Permissions` blocks as documented below.
     */
    public readonly oauth2Permissions!: pulumi.Output<outputs.ApplicationOauth2Permission[]>;
    /**
     * The Application's Object ID.
     */
    public /*out*/ readonly objectId!: pulumi.Output<string>;
    /**
     * A collection of `accessToken` or `idToken` blocks as documented below which list the optional claims configured for each token type. For more information see https://docs.microsoft.com/en-us/azure/active-directory/develop/active-directory-optional-claims
     */
    public readonly optionalClaims!: pulumi.Output<outputs.ApplicationOptionalClaims | undefined>;
    /**
     * A list of Azure AD Object IDs that will be granted ownership of the application. Defaults to the Object ID of the caller creating the application. If a list is specified the caller Object ID will no longer be included unless explicitly added to the list.
     */
    public readonly owners!: pulumi.Output<string[]>;
    /**
     * If `true`, will return an error when an existing Application is found with the same name. Defaults to `false`.
     */
    public readonly preventDuplicateNames!: pulumi.Output<boolean | undefined>;
    /**
     * Is this Azure AD Application a public client? Defaults to `false`.
     */
    public readonly publicClient!: pulumi.Output<boolean>;
    /**
     * A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.
     */
    public readonly replyUrls!: pulumi.Output<string[]>;
    /**
     * A collection of `requiredResourceAccess` blocks as documented below.
     */
    public readonly requiredResourceAccesses!: pulumi.Output<outputs.ApplicationRequiredResourceAccess[] | undefined>;
    /**
     * Type of an application: `webapp/api` or `native`. Defaults to `webapp/api`. For `native` apps type `identifierUris` property can not not be set.
     *
     * @deprecated This property is deprecated and will be removed in version 2.0 of this provider.
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a Application resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ApplicationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationArgs | ApplicationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ApplicationState | undefined;
            inputs["appRoles"] = state ? state.appRoles : undefined;
            inputs["applicationId"] = state ? state.applicationId : undefined;
            inputs["availableToOtherTenants"] = state ? state.availableToOtherTenants : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["groupMembershipClaims"] = state ? state.groupMembershipClaims : undefined;
            inputs["homepage"] = state ? state.homepage : undefined;
            inputs["identifierUris"] = state ? state.identifierUris : undefined;
            inputs["logoutUrl"] = state ? state.logoutUrl : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["oauth2AllowImplicitFlow"] = state ? state.oauth2AllowImplicitFlow : undefined;
            inputs["oauth2Permissions"] = state ? state.oauth2Permissions : undefined;
            inputs["objectId"] = state ? state.objectId : undefined;
            inputs["optionalClaims"] = state ? state.optionalClaims : undefined;
            inputs["owners"] = state ? state.owners : undefined;
            inputs["preventDuplicateNames"] = state ? state.preventDuplicateNames : undefined;
            inputs["publicClient"] = state ? state.publicClient : undefined;
            inputs["replyUrls"] = state ? state.replyUrls : undefined;
            inputs["requiredResourceAccesses"] = state ? state.requiredResourceAccesses : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as ApplicationArgs | undefined;
            inputs["appRoles"] = args ? args.appRoles : undefined;
            inputs["availableToOtherTenants"] = args ? args.availableToOtherTenants : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["groupMembershipClaims"] = args ? args.groupMembershipClaims : undefined;
            inputs["homepage"] = args ? args.homepage : undefined;
            inputs["identifierUris"] = args ? args.identifierUris : undefined;
            inputs["logoutUrl"] = args ? args.logoutUrl : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["oauth2AllowImplicitFlow"] = args ? args.oauth2AllowImplicitFlow : undefined;
            inputs["oauth2Permissions"] = args ? args.oauth2Permissions : undefined;
            inputs["optionalClaims"] = args ? args.optionalClaims : undefined;
            inputs["owners"] = args ? args.owners : undefined;
            inputs["preventDuplicateNames"] = args ? args.preventDuplicateNames : undefined;
            inputs["publicClient"] = args ? args.publicClient : undefined;
            inputs["replyUrls"] = args ? args.replyUrls : undefined;
            inputs["requiredResourceAccesses"] = args ? args.requiredResourceAccesses : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["applicationId"] = undefined /*out*/;
            inputs["objectId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Application.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Application resources.
 */
export interface ApplicationState {
    /**
     * A collection of `appRole` blocks as documented below. For more information https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles
     */
    readonly appRoles?: pulumi.Input<pulumi.Input<inputs.ApplicationAppRole>[]>;
    /**
     * The Application ID (Client ID).
     */
    readonly applicationId?: pulumi.Input<string>;
    /**
     * Is this Azure AD Application available to other tenants? Defaults to `false`.
     */
    readonly availableToOtherTenants?: pulumi.Input<boolean>;
    /**
     * The display name for the application.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects. Defaults to `SecurityGroup`. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
     */
    readonly groupMembershipClaims?: pulumi.Input<string>;
    /**
     * The URL to the application's home page.
     */
    readonly homepage?: pulumi.Input<string>;
    /**
     * A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
     */
    readonly identifierUris?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The URL of the logout page.
     */
    readonly logoutUrl?: pulumi.Input<string>;
    /**
     * The name of the optional claim.
     *
     * @deprecated This property has been renamed to `display_name` and will be removed in version 2.0 of this provider.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Does this Azure AD Application allow OAuth2.0 implicit flow tokens? Defaults to `false`.
     */
    readonly oauth2AllowImplicitFlow?: pulumi.Input<boolean>;
    /**
     * A collection of OAuth 2.0 permission scopes that the web API (resource) app exposes to client apps. Each permission is covered by `oauth2Permissions` blocks as documented below.
     */
    readonly oauth2Permissions?: pulumi.Input<pulumi.Input<inputs.ApplicationOauth2Permission>[]>;
    /**
     * The Application's Object ID.
     */
    readonly objectId?: pulumi.Input<string>;
    /**
     * A collection of `accessToken` or `idToken` blocks as documented below which list the optional claims configured for each token type. For more information see https://docs.microsoft.com/en-us/azure/active-directory/develop/active-directory-optional-claims
     */
    readonly optionalClaims?: pulumi.Input<inputs.ApplicationOptionalClaims>;
    /**
     * A list of Azure AD Object IDs that will be granted ownership of the application. Defaults to the Object ID of the caller creating the application. If a list is specified the caller Object ID will no longer be included unless explicitly added to the list.
     */
    readonly owners?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If `true`, will return an error when an existing Application is found with the same name. Defaults to `false`.
     */
    readonly preventDuplicateNames?: pulumi.Input<boolean>;
    /**
     * Is this Azure AD Application a public client? Defaults to `false`.
     */
    readonly publicClient?: pulumi.Input<boolean>;
    /**
     * A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.
     */
    readonly replyUrls?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A collection of `requiredResourceAccess` blocks as documented below.
     */
    readonly requiredResourceAccesses?: pulumi.Input<pulumi.Input<inputs.ApplicationRequiredResourceAccess>[]>;
    /**
     * Type of an application: `webapp/api` or `native`. Defaults to `webapp/api`. For `native` apps type `identifierUris` property can not not be set.
     *
     * @deprecated This property is deprecated and will be removed in version 2.0 of this provider.
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Application resource.
 */
export interface ApplicationArgs {
    /**
     * A collection of `appRole` blocks as documented below. For more information https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles
     */
    readonly appRoles?: pulumi.Input<pulumi.Input<inputs.ApplicationAppRole>[]>;
    /**
     * Is this Azure AD Application available to other tenants? Defaults to `false`.
     */
    readonly availableToOtherTenants?: pulumi.Input<boolean>;
    /**
     * The display name for the application.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects. Defaults to `SecurityGroup`. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
     */
    readonly groupMembershipClaims?: pulumi.Input<string>;
    /**
     * The URL to the application's home page.
     */
    readonly homepage?: pulumi.Input<string>;
    /**
     * A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
     */
    readonly identifierUris?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The URL of the logout page.
     */
    readonly logoutUrl?: pulumi.Input<string>;
    /**
     * The name of the optional claim.
     *
     * @deprecated This property has been renamed to `display_name` and will be removed in version 2.0 of this provider.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Does this Azure AD Application allow OAuth2.0 implicit flow tokens? Defaults to `false`.
     */
    readonly oauth2AllowImplicitFlow?: pulumi.Input<boolean>;
    /**
     * A collection of OAuth 2.0 permission scopes that the web API (resource) app exposes to client apps. Each permission is covered by `oauth2Permissions` blocks as documented below.
     */
    readonly oauth2Permissions?: pulumi.Input<pulumi.Input<inputs.ApplicationOauth2Permission>[]>;
    /**
     * A collection of `accessToken` or `idToken` blocks as documented below which list the optional claims configured for each token type. For more information see https://docs.microsoft.com/en-us/azure/active-directory/develop/active-directory-optional-claims
     */
    readonly optionalClaims?: pulumi.Input<inputs.ApplicationOptionalClaims>;
    /**
     * A list of Azure AD Object IDs that will be granted ownership of the application. Defaults to the Object ID of the caller creating the application. If a list is specified the caller Object ID will no longer be included unless explicitly added to the list.
     */
    readonly owners?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If `true`, will return an error when an existing Application is found with the same name. Defaults to `false`.
     */
    readonly preventDuplicateNames?: pulumi.Input<boolean>;
    /**
     * Is this Azure AD Application a public client? Defaults to `false`.
     */
    readonly publicClient?: pulumi.Input<boolean>;
    /**
     * A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.
     */
    readonly replyUrls?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A collection of `requiredResourceAccess` blocks as documented below.
     */
    readonly requiredResourceAccesses?: pulumi.Input<pulumi.Input<inputs.ApplicationRequiredResourceAccess>[]>;
    /**
     * Type of an application: `webapp/api` or `native`. Defaults to `webapp/api`. For `native` apps type `identifierUris` property can not not be set.
     *
     * @deprecated This property is deprecated and will be removed in version 2.0 of this provider.
     */
    readonly type?: pulumi.Input<string>;
}

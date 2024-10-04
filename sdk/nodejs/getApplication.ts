// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to access information about an existing Application within Azure Active Directory.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this data source.
 *
 * When authenticated with a service principal, this data source requires one of the following application roles: `Application.Read.All` or `Directory.Read.All`
 *
 * When authenticated with a user principal, this data source does not require any additional roles.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = azuread.getApplication({
 *     displayName: "My First AzureAD Application",
 * });
 * export const applicationObjectId = example.then(example => example.objectId);
 * ```
 */
export function getApplication(args?: GetApplicationArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azuread:index/getApplication:getApplication", {
        "clientId": args.clientId,
        "displayName": args.displayName,
        "identifierUri": args.identifierUri,
        "objectId": args.objectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getApplication.
 */
export interface GetApplicationArgs {
    /**
     * Specifies the Client ID of the application.
     */
    clientId?: string;
    /**
     * Specifies the display name of the application.
     */
    displayName?: string;
    /**
     * Specifies any identifier URI of the application. See also the `identifierUris` attribute which contains a list of all identifier URIs for the application.
     *
     * > One of `clientId`, `displayName`, `objectId`, or `identifierUri` must be specified.
     */
    identifierUri?: string;
    /**
     * Specifies the Object ID of the application.
     */
    objectId?: string;
}

/**
 * A collection of values returned by getApplication.
 */
export interface GetApplicationResult {
    /**
     * An `api` block as documented below.
     */
    readonly apis: outputs.GetApplicationApi[];
    /**
     * A mapping of app role values to app role IDs, intended to be useful when referencing app roles in other resources in your configuration.
     */
    readonly appRoleIds: {[key: string]: string};
    /**
     * A collection of `appRole` blocks as documented below. For more information see [official documentation on Application Roles](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).
     */
    readonly appRoles: outputs.GetApplicationAppRole[];
    /**
     * The Client ID for the application.
     */
    readonly clientId: string;
    /**
     * Description of the app role that appears when the role is being assigned and, if the role functions as an application permissions, during the consent experiences.
     */
    readonly description: string;
    /**
     * Specifies whether this application supports device authentication without a user.
     */
    readonly deviceOnlyAuthEnabled: boolean;
    /**
     * Whether Microsoft has disabled the registered application. If the application is disabled, this will be a string indicating the status/reason, e.g. `DisabledDueToViolationOfServicesAgreement`
     */
    readonly disabledByMicrosoft: string;
    /**
     * Display name for the app role that appears during app role assignment and in consent experiences.
     */
    readonly displayName: string;
    /**
     * The fallback application type as public client, such as an installed application running on a mobile device.
     */
    readonly fallbackPublicClientEnabled: boolean;
    /**
     * A `features` block as described below.
     */
    readonly featureTags: outputs.GetApplicationFeatureTag[];
    /**
     * The `groups` claim issued in a user or OAuth 2.0 access token that the app expects.
     */
    readonly groupMembershipClaims: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly identifierUri: string;
    /**
     * A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
     */
    readonly identifierUris: string[];
    /**
     * CDN URL to the application's logo.
     */
    readonly logoUrl: string;
    /**
     * URL of the application's marketing page.
     */
    readonly marketingUrl: string;
    /**
     * User-specified notes relevant for the management of the application.
     */
    readonly notes: string;
    /**
     * A mapping of OAuth2.0 permission scope values to scope IDs, intended to be useful when referencing permission scopes in other resources in your configuration.
     */
    readonly oauth2PermissionScopeIds: {[key: string]: string};
    /**
     * Specifies whether, as part of OAuth 2.0 token requests, Azure AD allows POST requests, as opposed to GET requests. When `false`, only GET requests are allowed.
     */
    readonly oauth2PostResponseRequired: boolean;
    /**
     * The application's object ID.
     */
    readonly objectId: string;
    /**
     * An `optionalClaims` block as documented below.
     */
    readonly optionalClaims: outputs.GetApplicationOptionalClaim[];
    /**
     * A list of object IDs of principals that are assigned ownership of the application.
     */
    readonly owners: string[];
    /**
     * URL of the application's privacy statement.
     */
    readonly privacyStatementUrl: string;
    /**
     * A `publicClient` block as documented below.
     */
    readonly publicClients: outputs.GetApplicationPublicClient[];
    /**
     * The verified publisher domain for the application.
     */
    readonly publisherDomain: string;
    /**
     * A collection of `requiredResourceAccess` blocks as documented below.
     */
    readonly requiredResourceAccesses: outputs.GetApplicationRequiredResourceAccess[];
    /**
     * References application context information from a Service or Asset Management database.
     */
    readonly serviceManagementReference: string;
    /**
     * The Microsoft account types that are supported for the current application. One of `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`.
     */
    readonly signInAudience: string;
    /**
     * A `singlePageApplication` block as documented below.
     */
    readonly singlePageApplications: outputs.GetApplicationSinglePageApplication[];
    /**
     * URL of the application's support page.
     */
    readonly supportUrl: string;
    /**
     * A list of tags applied to the application.
     */
    readonly tags: string[];
    /**
     * URL of the application's terms of service statement.
     */
    readonly termsOfServiceUrl: string;
    /**
     * A `web` block as documented below.
     */
    readonly webs: outputs.GetApplicationWeb[];
}
/**
 * Use this data source to access information about an existing Application within Azure Active Directory.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this data source.
 *
 * When authenticated with a service principal, this data source requires one of the following application roles: `Application.Read.All` or `Directory.Read.All`
 *
 * When authenticated with a user principal, this data source does not require any additional roles.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = azuread.getApplication({
 *     displayName: "My First AzureAD Application",
 * });
 * export const applicationObjectId = example.then(example => example.objectId);
 * ```
 */
export function getApplicationOutput(args?: GetApplicationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetApplicationResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azuread:index/getApplication:getApplication", {
        "clientId": args.clientId,
        "displayName": args.displayName,
        "identifierUri": args.identifierUri,
        "objectId": args.objectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getApplication.
 */
export interface GetApplicationOutputArgs {
    /**
     * Specifies the Client ID of the application.
     */
    clientId?: pulumi.Input<string>;
    /**
     * Specifies the display name of the application.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Specifies any identifier URI of the application. See also the `identifierUris` attribute which contains a list of all identifier URIs for the application.
     *
     * > One of `clientId`, `displayName`, `objectId`, or `identifierUri` must be specified.
     */
    identifierUri?: pulumi.Input<string>;
    /**
     * Specifies the Object ID of the application.
     */
    objectId?: pulumi.Input<string>;
}

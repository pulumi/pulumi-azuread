// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages an application registration within Azure Active Directory.
 *
 * For a more comprehensive alternative, please see the azuread.Application resource. Please note that this resource should not be used together with the `azuread.Application` resource when managing the same application.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this resource.
 *
 * When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.OwnedBy` or `Application.ReadWrite.All`
 *
 * When authenticated with a user principal, this resource may require one of the following directory roles: `Application Administrator` or `Global Administrator`
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = new azuread.ApplicationRegistration("example", {
 *     displayName: "Example Application",
 *     description: "My example application",
 *     signInAudience: "AzureADMyOrg",
 *     homepageUrl: "https://app.hashitown.com/",
 *     logoutUrl: "https://app.hashitown.com/logout",
 *     marketingUrl: "https://hashitown.com/",
 *     privacyStatementUrl: "https://hashitown.com/privacy",
 *     supportUrl: "https://support.hashitown.com/",
 *     termsOfServiceUrl: "https://hashitown.com/terms",
 * });
 * ```
 *
 * ## Import
 *
 * Application Registrations can be imported using the object ID of the application, in the following format.
 *
 * ```sh
 * $ pulumi import azuread:index/applicationRegistration:ApplicationRegistration example /applications/00000000-0000-0000-0000-000000000000
 * ```
 */
export class ApplicationRegistration extends pulumi.CustomResource {
    /**
     * Get an existing ApplicationRegistration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApplicationRegistrationState, opts?: pulumi.CustomResourceOptions): ApplicationRegistration {
        return new ApplicationRegistration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/applicationRegistration:ApplicationRegistration';

    /**
     * Returns true if the given object is an instance of ApplicationRegistration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApplicationRegistration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApplicationRegistration.__pulumiType;
    }

    /**
     * The Client ID for the application, which is globally unique.
     */
    public /*out*/ readonly clientId!: pulumi.Output<string>;
    /**
     * A description of the application, as shown to end users.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Whether Microsoft has disabled the registered application. If the application is disabled, this will be a string indicating the status/reason, e.g. `DisabledDueToViolationOfServicesAgreement`
     */
    public /*out*/ readonly disabledByMicrosoft!: pulumi.Output<string>;
    /**
     * The display name for the application.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Configures the `groups` claim issued in a user or OAuth access token that the app expects. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
     */
    public readonly groupMembershipClaims!: pulumi.Output<string[] | undefined>;
    /**
     * Home page or landing page of the application.
     */
    public readonly homepageUrl!: pulumi.Output<string | undefined>;
    /**
     * Whether this web application can request an access token using OAuth implicit flow.
     */
    public readonly implicitAccessTokenIssuanceEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Whether this web application can request an ID token using OAuth implicit flow.
     */
    public readonly implicitIdTokenIssuanceEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols.
     */
    public readonly logoutUrl!: pulumi.Output<string | undefined>;
    /**
     * URL of the marketing page for the application.
     */
    public readonly marketingUrl!: pulumi.Output<string | undefined>;
    /**
     * User-specified notes relevant for the management of the application.
     */
    public readonly notes!: pulumi.Output<string | undefined>;
    /**
     * The object ID of the application within the tenant.
     */
    public /*out*/ readonly objectId!: pulumi.Output<string>;
    /**
     * URL of the privacy statement for the application.
     */
    public readonly privacyStatementUrl!: pulumi.Output<string | undefined>;
    /**
     * The verified publisher domain for the application.
     */
    public /*out*/ readonly publisherDomain!: pulumi.Output<string>;
    /**
     * The access token version expected by this resource. Must be one of `1` or `2`, and must be `2` when `signInAudience` is either `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount` Defaults to `2`.
     */
    public readonly requestedAccessTokenVersion!: pulumi.Output<number | undefined>;
    /**
     * References application context information from a Service or Asset Management database.
     */
    public readonly serviceManagementReference!: pulumi.Output<string | undefined>;
    /**
     * The Microsoft account types that are supported for the current application. Must be one of `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`. Defaults to `AzureADMyOrg`.
     */
    public readonly signInAudience!: pulumi.Output<string | undefined>;
    /**
     * URL of the support page for the application.
     */
    public readonly supportUrl!: pulumi.Output<string | undefined>;
    /**
     * URL of the terms of service statement for the application.
     */
    public readonly termsOfServiceUrl!: pulumi.Output<string | undefined>;

    /**
     * Create a ApplicationRegistration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationRegistrationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationRegistrationArgs | ApplicationRegistrationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApplicationRegistrationState | undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["disabledByMicrosoft"] = state ? state.disabledByMicrosoft : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["groupMembershipClaims"] = state ? state.groupMembershipClaims : undefined;
            resourceInputs["homepageUrl"] = state ? state.homepageUrl : undefined;
            resourceInputs["implicitAccessTokenIssuanceEnabled"] = state ? state.implicitAccessTokenIssuanceEnabled : undefined;
            resourceInputs["implicitIdTokenIssuanceEnabled"] = state ? state.implicitIdTokenIssuanceEnabled : undefined;
            resourceInputs["logoutUrl"] = state ? state.logoutUrl : undefined;
            resourceInputs["marketingUrl"] = state ? state.marketingUrl : undefined;
            resourceInputs["notes"] = state ? state.notes : undefined;
            resourceInputs["objectId"] = state ? state.objectId : undefined;
            resourceInputs["privacyStatementUrl"] = state ? state.privacyStatementUrl : undefined;
            resourceInputs["publisherDomain"] = state ? state.publisherDomain : undefined;
            resourceInputs["requestedAccessTokenVersion"] = state ? state.requestedAccessTokenVersion : undefined;
            resourceInputs["serviceManagementReference"] = state ? state.serviceManagementReference : undefined;
            resourceInputs["signInAudience"] = state ? state.signInAudience : undefined;
            resourceInputs["supportUrl"] = state ? state.supportUrl : undefined;
            resourceInputs["termsOfServiceUrl"] = state ? state.termsOfServiceUrl : undefined;
        } else {
            const args = argsOrState as ApplicationRegistrationArgs | undefined;
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["groupMembershipClaims"] = args ? args.groupMembershipClaims : undefined;
            resourceInputs["homepageUrl"] = args ? args.homepageUrl : undefined;
            resourceInputs["implicitAccessTokenIssuanceEnabled"] = args ? args.implicitAccessTokenIssuanceEnabled : undefined;
            resourceInputs["implicitIdTokenIssuanceEnabled"] = args ? args.implicitIdTokenIssuanceEnabled : undefined;
            resourceInputs["logoutUrl"] = args ? args.logoutUrl : undefined;
            resourceInputs["marketingUrl"] = args ? args.marketingUrl : undefined;
            resourceInputs["notes"] = args ? args.notes : undefined;
            resourceInputs["privacyStatementUrl"] = args ? args.privacyStatementUrl : undefined;
            resourceInputs["requestedAccessTokenVersion"] = args ? args.requestedAccessTokenVersion : undefined;
            resourceInputs["serviceManagementReference"] = args ? args.serviceManagementReference : undefined;
            resourceInputs["signInAudience"] = args ? args.signInAudience : undefined;
            resourceInputs["supportUrl"] = args ? args.supportUrl : undefined;
            resourceInputs["termsOfServiceUrl"] = args ? args.termsOfServiceUrl : undefined;
            resourceInputs["clientId"] = undefined /*out*/;
            resourceInputs["disabledByMicrosoft"] = undefined /*out*/;
            resourceInputs["objectId"] = undefined /*out*/;
            resourceInputs["publisherDomain"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ApplicationRegistration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApplicationRegistration resources.
 */
export interface ApplicationRegistrationState {
    /**
     * The Client ID for the application, which is globally unique.
     */
    clientId?: pulumi.Input<string>;
    /**
     * A description of the application, as shown to end users.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether Microsoft has disabled the registered application. If the application is disabled, this will be a string indicating the status/reason, e.g. `DisabledDueToViolationOfServicesAgreement`
     */
    disabledByMicrosoft?: pulumi.Input<string>;
    /**
     * The display name for the application.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Configures the `groups` claim issued in a user or OAuth access token that the app expects. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
     */
    groupMembershipClaims?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Home page or landing page of the application.
     */
    homepageUrl?: pulumi.Input<string>;
    /**
     * Whether this web application can request an access token using OAuth implicit flow.
     */
    implicitAccessTokenIssuanceEnabled?: pulumi.Input<boolean>;
    /**
     * Whether this web application can request an ID token using OAuth implicit flow.
     */
    implicitIdTokenIssuanceEnabled?: pulumi.Input<boolean>;
    /**
     * The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols.
     */
    logoutUrl?: pulumi.Input<string>;
    /**
     * URL of the marketing page for the application.
     */
    marketingUrl?: pulumi.Input<string>;
    /**
     * User-specified notes relevant for the management of the application.
     */
    notes?: pulumi.Input<string>;
    /**
     * The object ID of the application within the tenant.
     */
    objectId?: pulumi.Input<string>;
    /**
     * URL of the privacy statement for the application.
     */
    privacyStatementUrl?: pulumi.Input<string>;
    /**
     * The verified publisher domain for the application.
     */
    publisherDomain?: pulumi.Input<string>;
    /**
     * The access token version expected by this resource. Must be one of `1` or `2`, and must be `2` when `signInAudience` is either `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount` Defaults to `2`.
     */
    requestedAccessTokenVersion?: pulumi.Input<number>;
    /**
     * References application context information from a Service or Asset Management database.
     */
    serviceManagementReference?: pulumi.Input<string>;
    /**
     * The Microsoft account types that are supported for the current application. Must be one of `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`. Defaults to `AzureADMyOrg`.
     */
    signInAudience?: pulumi.Input<string>;
    /**
     * URL of the support page for the application.
     */
    supportUrl?: pulumi.Input<string>;
    /**
     * URL of the terms of service statement for the application.
     */
    termsOfServiceUrl?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ApplicationRegistration resource.
 */
export interface ApplicationRegistrationArgs {
    /**
     * A description of the application, as shown to end users.
     */
    description?: pulumi.Input<string>;
    /**
     * The display name for the application.
     */
    displayName: pulumi.Input<string>;
    /**
     * Configures the `groups` claim issued in a user or OAuth access token that the app expects. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
     */
    groupMembershipClaims?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Home page or landing page of the application.
     */
    homepageUrl?: pulumi.Input<string>;
    /**
     * Whether this web application can request an access token using OAuth implicit flow.
     */
    implicitAccessTokenIssuanceEnabled?: pulumi.Input<boolean>;
    /**
     * Whether this web application can request an ID token using OAuth implicit flow.
     */
    implicitIdTokenIssuanceEnabled?: pulumi.Input<boolean>;
    /**
     * The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols.
     */
    logoutUrl?: pulumi.Input<string>;
    /**
     * URL of the marketing page for the application.
     */
    marketingUrl?: pulumi.Input<string>;
    /**
     * User-specified notes relevant for the management of the application.
     */
    notes?: pulumi.Input<string>;
    /**
     * URL of the privacy statement for the application.
     */
    privacyStatementUrl?: pulumi.Input<string>;
    /**
     * The access token version expected by this resource. Must be one of `1` or `2`, and must be `2` when `signInAudience` is either `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount` Defaults to `2`.
     */
    requestedAccessTokenVersion?: pulumi.Input<number>;
    /**
     * References application context information from a Service or Asset Management database.
     */
    serviceManagementReference?: pulumi.Input<string>;
    /**
     * The Microsoft account types that are supported for the current application. Must be one of `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`. Defaults to `AzureADMyOrg`.
     */
    signInAudience?: pulumi.Input<string>;
    /**
     * URL of the support page for the application.
     */
    supportUrl?: pulumi.Input<string>;
    /**
     * URL of the terms of service statement for the application.
     */
    termsOfServiceUrl?: pulumi.Input<string>;
}

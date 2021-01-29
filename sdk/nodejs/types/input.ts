// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";

export interface ApplicationAppRole {
    /**
     * Specifies whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications (that are accessing this application in daemon service scenarios) by setting to `Application`, or to both.
     */
    allowedMemberTypes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Permission help text that appears in the admin app assignment and consent experiences.
     */
    description: pulumi.Input<string>;
    /**
     * Display name for the permission that appears in the admin consent and app assignment experiences.
     */
    displayName: pulumi.Input<string>;
    /**
     * The unique identifier of the permision. This attribute is computed and cannot be specified manually in this block. If you need to specify a custom `id`, it's recommended to use the azuread.ApplicationOAuth2Permission resource.
     */
    id?: pulumi.Input<string>;
    /**
     * Determines if the permission is enabled: defaults to `true`.
     */
    isEnabled?: pulumi.Input<boolean>;
    /**
     * The value of the scope claim that the resource application should expect in the OAuth 2.0 access token.
     */
    value?: pulumi.Input<string>;
}

export interface ApplicationOauth2Permission {
    /**
     * Permission help text that appears in the admin consent and app assignment experiences.
     */
    adminConsentDescription?: pulumi.Input<string>;
    /**
     * Display name for the permission that appears in the admin consent and app assignment experiences.
     */
    adminConsentDisplayName?: pulumi.Input<string>;
    /**
     * The unique identifier of the app role. This attribute is computed and cannot be specified manually in this block. If you need to specify a custom `id`, it's recommended to use the azuread.ApplicationAppRole resource.
     */
    id?: pulumi.Input<string>;
    /**
     * Determines if the app role is enabled: Defaults to `true`.
     */
    isEnabled?: pulumi.Input<boolean>;
    /**
     * Type of an application: `webapp/api` or `native`. Defaults to `webapp/api`. For `native` apps type `identifierUris` property can not not be set.
     */
    type?: pulumi.Input<string>;
    /**
     * Permission help text that appears in the end user consent experience.
     */
    userConsentDescription?: pulumi.Input<string>;
    /**
     * Display name for the permission that appears in the end user consent experience.
     */
    userConsentDisplayName?: pulumi.Input<string>;
    /**
     * Specifies the value of the roles claim that the application should expect in the authentication and access tokens.
     */
    value?: pulumi.Input<string>;
}

export interface ApplicationOptionalClaims {
    accessTokens?: pulumi.Input<pulumi.Input<inputs.ApplicationOptionalClaimsAccessToken>[]>;
    idTokens?: pulumi.Input<pulumi.Input<inputs.ApplicationOptionalClaimsIdToken>[]>;
}

export interface ApplicationOptionalClaimsAccessToken {
    /**
     * List of Additional Properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim.
     */
    additionalProperties?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether the claim specified by the client is necessary to ensure a smooth authorization experience.
     */
    essential?: pulumi.Input<boolean>;
    /**
     * The name of the optional claim.
     */
    name: pulumi.Input<string>;
    /**
     * The source of the claim. If `source` is absent, the claim is a predefined optional claim. If `source` is `user`, the value of `name` is the extension property from the user object.
     */
    source?: pulumi.Input<string>;
}

export interface ApplicationOptionalClaimsIdToken {
    /**
     * List of Additional Properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim.
     */
    additionalProperties?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether the claim specified by the client is necessary to ensure a smooth authorization experience.
     */
    essential?: pulumi.Input<boolean>;
    /**
     * The name of the optional claim.
     */
    name: pulumi.Input<string>;
    /**
     * The source of the claim. If `source` is absent, the claim is a predefined optional claim. If `source` is `user`, the value of `name` is the extension property from the user object.
     */
    source?: pulumi.Input<string>;
}

export interface ApplicationRequiredResourceAccess {
    /**
     * A collection of `resourceAccess` blocks as documented below.
     */
    resourceAccesses: pulumi.Input<pulumi.Input<inputs.ApplicationRequiredResourceAccessResourceAccess>[]>;
    /**
     * The unique identifier for the resource that the application requires access to. This should be equal to the appId declared on the target resource application.
     */
    resourceAppId: pulumi.Input<string>;
}

export interface ApplicationRequiredResourceAccessResourceAccess {
    /**
     * The unique identifier for one of the `OAuth2Permission` or `AppRole` instances that the resource application exposes.
     */
    id: pulumi.Input<string>;
    /**
     * Specifies whether the id property references an `OAuth2Permission` or an `AppRole`. Possible values are `Scope` or `Role`.
     */
    type: pulumi.Input<string>;
}

export interface GetApplicationOauth2Permission {
    /**
     * The description of the admin consent
     */
    adminConsentDescription?: string;
    /**
     * The display name of the admin consent
     */
    adminConsentDisplayName?: string;
    /**
     * The unique identifier for one of the `OAuth2Permission` or `AppRole` instances that the resource application exposes.
     */
    id?: string;
    /**
     * Is this permission enabled?
     */
    isEnabled?: boolean;
    /**
     * Specifies whether the id property references an `OAuth2Permission` or an `AppRole`.
     */
    type?: string;
    /**
     * The description of the user consent
     */
    userConsentDescription?: string;
    /**
     * The display name of the user consent
     */
    userConsentDisplayName?: string;
    /**
     * The name of this permission
     */
    value?: string;
}

export interface GetApplicationOptionalClaims {
    accessTokens?: inputs.GetApplicationOptionalClaimsAccessToken[];
    idTokens?: inputs.GetApplicationOptionalClaimsIdToken[];
}

export interface GetApplicationOptionalClaimsAccessToken {
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

export interface GetApplicationOptionalClaimsIdToken {
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

export interface GetServicePrincipalOauth2Permission {
    /**
     * The description of the admin consent
     */
    adminConsentDescription?: string;
    /**
     * The display name of the admin consent
     */
    adminConsentDisplayName?: string;
    /**
     * The unique identifier for one of the `OAuth2Permission`
     */
    id?: string;
    /**
     * Is this permission enabled?
     */
    isEnabled?: boolean;
    /**
     * The type of the permission
     */
    type?: string;
    /**
     * The description of the user consent
     */
    userConsentDescription?: string;
    /**
     * The display name of the user consent
     */
    userConsentDisplayName?: string;
    /**
     * The name of this permission
     */
    value?: string;
}

export interface ServicePrincipalAppRole {
    allowedMemberTypes?: pulumi.Input<pulumi.Input<string>[]>;
    description?: pulumi.Input<string>;
    /**
     * The Display Name of the Application associated with this Service Principal.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The unique identifier for one of the `OAuth2Permission`.
     */
    id?: pulumi.Input<string>;
    /**
     * Is this permission enabled?
     */
    isEnabled?: pulumi.Input<boolean>;
    /**
     * The name of this permission.
     */
    value?: pulumi.Input<string>;
}

export interface ServicePrincipalOauth2Permission {
    /**
     * The description of the admin consent.
     */
    adminConsentDescription?: pulumi.Input<string>;
    /**
     * The display name of the admin consent.
     */
    adminConsentDisplayName?: pulumi.Input<string>;
    /**
     * The unique identifier for one of the `OAuth2Permission`.
     */
    id?: pulumi.Input<string>;
    /**
     * Is this permission enabled?
     */
    isEnabled?: pulumi.Input<boolean>;
    /**
     * The type of the permission.
     */
    type?: pulumi.Input<string>;
    /**
     * The description of the user consent.
     */
    userConsentDescription?: pulumi.Input<string>;
    /**
     * The display name of the user consent.
     */
    userConsentDisplayName?: pulumi.Input<string>;
    /**
     * The name of this permission.
     */
    value?: pulumi.Input<string>;
}

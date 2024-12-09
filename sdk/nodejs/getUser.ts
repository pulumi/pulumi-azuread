// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Gets information about an Azure Active Directory user.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this data source.
 *
 * When authenticated with a service principal, this data source requires one of the following application roles: `User.Read.All` or `Directory.Read.All`
 *
 * When authenticated with a user principal, this data source does not require any additional roles.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = azuread.getUser({
 *     userPrincipalName: "user@example.com",
 * });
 * ```
 */
export function getUser(args?: GetUserArgs, opts?: pulumi.InvokeOptions): Promise<GetUserResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azuread:index/getUser:getUser", {
        "employeeId": args.employeeId,
        "mail": args.mail,
        "mailNickname": args.mailNickname,
        "objectId": args.objectId,
        "userPrincipalName": args.userPrincipalName,
    }, opts);
}

/**
 * A collection of arguments for invoking getUser.
 */
export interface GetUserArgs {
    /**
     * The employee identifier assigned to the user by the organisation.
     */
    employeeId?: string;
    /**
     * The SMTP address for the user.
     */
    mail?: string;
    /**
     * The email alias of the user.
     */
    mailNickname?: string;
    /**
     * The object ID of the user.
     */
    objectId?: string;
    /**
     * The user principal name (UPN) of the user.
     *
     * > One of `userPrincipalName`, `objectId`, `mail`, `mailNickname` or `employeeId` must be specified.
     */
    userPrincipalName?: string;
}

/**
 * A collection of values returned by getUser.
 */
export interface GetUserResult {
    /**
     * Whether or not the account is enabled.
     */
    readonly accountEnabled: boolean;
    /**
     * The age group of the user. Supported values are `Adult`, `NotAdult` and `Minor`.
     */
    readonly ageGroup: string;
    /**
     * A list of telephone numbers for the user.
     */
    readonly businessPhones: string[];
    /**
     * The city in which the user is located.
     */
    readonly city: string;
    /**
     * The company name which the user is associated. This property can be useful for describing the company that an external user comes from.
     */
    readonly companyName: string;
    /**
     * Whether consent has been obtained for minors. Supported values are `Granted`, `Denied` and `NotRequired`.
     */
    readonly consentProvidedForMinor: string;
    /**
     * The cost center associated with the user.
     */
    readonly costCenter: string;
    /**
     * The country/region in which the user is located, e.g. `US` or `UK`.
     */
    readonly country: string;
    /**
     * Indicates whether the user account was created as a regular school or work account (`null`), an external account (`Invitation`), a local account for an Azure Active Directory B2C tenant (`LocalAccount`) or self-service sign-up using email verification (`EmailVerified`).
     */
    readonly creationType: string;
    /**
     * The name for the department in which the user works.
     */
    readonly department: string;
    /**
     * The display name of the user.
     */
    readonly displayName: string;
    /**
     * The name of the division in which the user works.
     */
    readonly division: string;
    /**
     * The employee identifier assigned to the user by the organisation.
     */
    readonly employeeId: string;
    /**
     * Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor.
     */
    readonly employeeType: string;
    /**
     * For an external user invited to the tenant, this property represents the invited user's invitation status. Possible values are `PendingAcceptance` or `Accepted`.
     */
    readonly externalUserState: string;
    /**
     * The fax number of the user.
     */
    readonly faxNumber: string;
    /**
     * The given name (first name) of the user.
     */
    readonly givenName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of instant message voice over IP (VOIP) session initiation protocol (SIP) addresses for the user.
     */
    readonly imAddresses: string[];
    /**
     * The user’s job title.
     */
    readonly jobTitle: string;
    /**
     * The SMTP address for the user.
     */
    readonly mail: string;
    /**
     * The email alias of the user.
     */
    readonly mailNickname: string;
    /**
     * The object ID of the user's manager.
     */
    readonly managerId: string;
    /**
     * The primary cellular telephone number for the user.
     */
    readonly mobilePhone: string;
    /**
     * The object ID of the user.
     */
    readonly objectId: string;
    /**
     * The office location in the user's place of business.
     */
    readonly officeLocation: string;
    /**
     * The on-premises distinguished name (DN) of the user, synchronised from the on-premises directory when Azure AD Connect is used.
     */
    readonly onpremisesDistinguishedName: string;
    /**
     * The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.
     */
    readonly onpremisesDomainName: string;
    /**
     * The value used to associate an on-premise Active Directory user account with their Azure AD user object.
     */
    readonly onpremisesImmutableId: string;
    /**
     * The on-premise SAM account name of the user.
     */
    readonly onpremisesSamAccountName: string;
    /**
     * The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.
     */
    readonly onpremisesSecurityIdentifier: string;
    /**
     * Whether this user is synchronised from an on-premises directory (`true`), no longer synchronised (`false`), or has never been synchronised (`null`).
     */
    readonly onpremisesSyncEnabled: boolean;
    /**
     * The on-premise user principal name of the user.
     */
    readonly onpremisesUserPrincipalName: string;
    /**
     * A list of additional email addresses for the user.
     */
    readonly otherMails: string[];
    /**
     * The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code.
     */
    readonly postalCode: string;
    /**
     * The user's preferred language, in ISO 639-1 notation.
     */
    readonly preferredLanguage: string;
    /**
     * List of email addresses for the user that direct to the same mailbox.
     */
    readonly proxyAddresses: string[];
    /**
     * Whether or not the Outlook global address list should include this user.
     */
    readonly showInAddressList: boolean;
    /**
     * The state or province in the user's address.
     */
    readonly state: string;
    /**
     * The street address of the user's place of business.
     */
    readonly streetAddress: string;
    /**
     * The user's surname (family name or last name).
     */
    readonly surname: string;
    /**
     * The usage location of the user.
     */
    readonly usageLocation: string;
    /**
     * The user principal name (UPN) of the user.
     */
    readonly userPrincipalName: string;
    /**
     * The user type in the directory. Possible values are `Guest` or `Member`.
     */
    readonly userType: string;
}
/**
 * Gets information about an Azure Active Directory user.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this data source.
 *
 * When authenticated with a service principal, this data source requires one of the following application roles: `User.Read.All` or `Directory.Read.All`
 *
 * When authenticated with a user principal, this data source does not require any additional roles.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = azuread.getUser({
 *     userPrincipalName: "user@example.com",
 * });
 * ```
 */
export function getUserOutput(args?: GetUserOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetUserResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azuread:index/getUser:getUser", {
        "employeeId": args.employeeId,
        "mail": args.mail,
        "mailNickname": args.mailNickname,
        "objectId": args.objectId,
        "userPrincipalName": args.userPrincipalName,
    }, opts);
}

/**
 * A collection of arguments for invoking getUser.
 */
export interface GetUserOutputArgs {
    /**
     * The employee identifier assigned to the user by the organisation.
     */
    employeeId?: pulumi.Input<string>;
    /**
     * The SMTP address for the user.
     */
    mail?: pulumi.Input<string>;
    /**
     * The email alias of the user.
     */
    mailNickname?: pulumi.Input<string>;
    /**
     * The object ID of the user.
     */
    objectId?: pulumi.Input<string>;
    /**
     * The user principal name (UPN) of the user.
     *
     * > One of `userPrincipalName`, `objectId`, `mail`, `mailNickname` or `employeeId` must be specified.
     */
    userPrincipalName?: pulumi.Input<string>;
}

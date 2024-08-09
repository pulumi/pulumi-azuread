// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.UserArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.UserState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a user within Azure Active Directory.
 * 
 * ## API Permissions
 * 
 * The following API permissions are required in order to use this resource.
 * 
 * When authenticated with a service principal, this resource requires one of the following application roles: `User.ReadWrite.All` or `Directory.ReadWrite.All`
 * 
 * When authenticated with a user principal, this resource requires one of the following directory roles: `User Administrator` or `Global Administrator`
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuread.User;
 * import com.pulumi.azuread.UserArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App }{{@code
 *     public static void main(String[] args) }{{@code
 *         Pulumi.run(App::stack);
 *     }}{@code
 * 
 *     public static void stack(Context ctx) }{{@code
 *         var example = new User("example", UserArgs.builder()
 *             .userPrincipalName("jdoe}{@literal @}{@code example.com")
 *             .displayName("J. Doe")
 *             .mailNickname("jdoe")
 *             .password("SecretP}{@literal @}{@code sswd99!")
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Users can be imported using their object ID, e.g.
 * 
 * ```sh
 * $ pulumi import azuread:index/user:User my_user 00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuread:index/user:User")
public class User extends com.pulumi.resources.CustomResource {
    /**
     * A freeform field for the user to describe themselves
     * 
     */
    @Export(name="aboutMe", refs={String.class}, tree="[0]")
    private Output<String> aboutMe;

    /**
     * @return A freeform field for the user to describe themselves
     * 
     */
    public Output<String> aboutMe() {
        return this.aboutMe;
    }
    /**
     * Whether or not the account should be enabled.
     * 
     */
    @Export(name="accountEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> accountEnabled;

    /**
     * @return Whether or not the account should be enabled.
     * 
     */
    public Output<Optional<Boolean>> accountEnabled() {
        return Codegen.optional(this.accountEnabled);
    }
    /**
     * The age group of the user. Supported values are `Adult`, `NotAdult` and `Minor`. Omit this property or specify a blank string to unset.
     * 
     */
    @Export(name="ageGroup", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ageGroup;

    /**
     * @return The age group of the user. Supported values are `Adult`, `NotAdult` and `Minor`. Omit this property or specify a blank string to unset.
     * 
     */
    public Output<Optional<String>> ageGroup() {
        return Codegen.optional(this.ageGroup);
    }
    /**
     * A list of telephone numbers for the user. Only one number can be set for this property. Read-only for users synced with Azure AD Connect.
     * 
     */
    @Export(name="businessPhones", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> businessPhones;

    /**
     * @return A list of telephone numbers for the user. Only one number can be set for this property. Read-only for users synced with Azure AD Connect.
     * 
     */
    public Output<List<String>> businessPhones() {
        return this.businessPhones;
    }
    /**
     * The city in which the user is located.
     * 
     */
    @Export(name="city", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> city;

    /**
     * @return The city in which the user is located.
     * 
     */
    public Output<Optional<String>> city() {
        return Codegen.optional(this.city);
    }
    /**
     * The company name which the user is associated. This property can be useful for describing the company that an external user comes from.
     * 
     */
    @Export(name="companyName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> companyName;

    /**
     * @return The company name which the user is associated. This property can be useful for describing the company that an external user comes from.
     * 
     */
    public Output<Optional<String>> companyName() {
        return Codegen.optional(this.companyName);
    }
    /**
     * Whether consent has been obtained for minors. Supported values are `Granted`, `Denied` and `NotRequired`. Omit this property or specify a blank string to unset.
     * 
     */
    @Export(name="consentProvidedForMinor", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> consentProvidedForMinor;

    /**
     * @return Whether consent has been obtained for minors. Supported values are `Granted`, `Denied` and `NotRequired`. Omit this property or specify a blank string to unset.
     * 
     */
    public Output<Optional<String>> consentProvidedForMinor() {
        return Codegen.optional(this.consentProvidedForMinor);
    }
    /**
     * The cost center associated with the user.
     * 
     */
    @Export(name="costCenter", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> costCenter;

    /**
     * @return The cost center associated with the user.
     * 
     */
    public Output<Optional<String>> costCenter() {
        return Codegen.optional(this.costCenter);
    }
    /**
     * The country/region in which the user is located. Examples include: `NO`, `JP`, and `GB`.
     * 
     */
    @Export(name="country", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> country;

    /**
     * @return The country/region in which the user is located. Examples include: `NO`, `JP`, and `GB`.
     * 
     */
    public Output<Optional<String>> country() {
        return Codegen.optional(this.country);
    }
    /**
     * Indicates whether the user account was created as a regular school or work account (`null`), an external account (`Invitation`), a local account for an Azure Active Directory B2C tenant (`LocalAccount`) or self-service sign-up using email verification (`EmailVerified`).
     * 
     */
    @Export(name="creationType", refs={String.class}, tree="[0]")
    private Output<String> creationType;

    /**
     * @return Indicates whether the user account was created as a regular school or work account (`null`), an external account (`Invitation`), a local account for an Azure Active Directory B2C tenant (`LocalAccount`) or self-service sign-up using email verification (`EmailVerified`).
     * 
     */
    public Output<String> creationType() {
        return this.creationType;
    }
    /**
     * The name for the department in which the user works.
     * 
     */
    @Export(name="department", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> department;

    /**
     * @return The name for the department in which the user works.
     * 
     */
    public Output<Optional<String>> department() {
        return Codegen.optional(this.department);
    }
    /**
     * Whether the user&#39;s password is exempt from expiring. Defaults to `false`.
     * 
     */
    @Export(name="disablePasswordExpiration", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disablePasswordExpiration;

    /**
     * @return Whether the user&#39;s password is exempt from expiring. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> disablePasswordExpiration() {
        return Codegen.optional(this.disablePasswordExpiration);
    }
    /**
     * Whether the user is allowed weaker passwords than the default policy to be specified. Defaults to `false`.
     * 
     */
    @Export(name="disableStrongPassword", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disableStrongPassword;

    /**
     * @return Whether the user is allowed weaker passwords than the default policy to be specified. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> disableStrongPassword() {
        return Codegen.optional(this.disableStrongPassword);
    }
    /**
     * The name to display in the address book for the user.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    /**
     * @return The name to display in the address book for the user.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * The name of the division in which the user works.
     * 
     */
    @Export(name="division", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> division;

    /**
     * @return The name of the division in which the user works.
     * 
     */
    public Output<Optional<String>> division() {
        return Codegen.optional(this.division);
    }
    /**
     * The employee identifier assigned to the user by the organisation.
     * 
     */
    @Export(name="employeeId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> employeeId;

    /**
     * @return The employee identifier assigned to the user by the organisation.
     * 
     */
    public Output<Optional<String>> employeeId() {
        return Codegen.optional(this.employeeId);
    }
    /**
     * Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor.
     * 
     */
    @Export(name="employeeType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> employeeType;

    /**
     * @return Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor.
     * 
     */
    public Output<Optional<String>> employeeType() {
        return Codegen.optional(this.employeeType);
    }
    /**
     * For an external user invited to the tenant, this property represents the invited user&#39;s invitation status. Possible values are `PendingAcceptance` or `Accepted`.
     * 
     */
    @Export(name="externalUserState", refs={String.class}, tree="[0]")
    private Output<String> externalUserState;

    /**
     * @return For an external user invited to the tenant, this property represents the invited user&#39;s invitation status. Possible values are `PendingAcceptance` or `Accepted`.
     * 
     */
    public Output<String> externalUserState() {
        return this.externalUserState;
    }
    /**
     * The fax number of the user.
     * 
     */
    @Export(name="faxNumber", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> faxNumber;

    /**
     * @return The fax number of the user.
     * 
     */
    public Output<Optional<String>> faxNumber() {
        return Codegen.optional(this.faxNumber);
    }
    /**
     * Whether the user is forced to change the password during the next sign-in. Only takes effect when also changing the password. Defaults to `false`.
     * 
     */
    @Export(name="forcePasswordChange", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> forcePasswordChange;

    /**
     * @return Whether the user is forced to change the password during the next sign-in. Only takes effect when also changing the password. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> forcePasswordChange() {
        return Codegen.optional(this.forcePasswordChange);
    }
    /**
     * The given name (first name) of the user.
     * 
     */
    @Export(name="givenName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> givenName;

    /**
     * @return The given name (first name) of the user.
     * 
     */
    public Output<Optional<String>> givenName() {
        return Codegen.optional(this.givenName);
    }
    /**
     * A list of instant message voice over IP (VOIP) session initiation protocol (SIP) addresses for the user.
     * 
     */
    @Export(name="imAddresses", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> imAddresses;

    /**
     * @return A list of instant message voice over IP (VOIP) session initiation protocol (SIP) addresses for the user.
     * 
     */
    public Output<List<String>> imAddresses() {
        return this.imAddresses;
    }
    /**
     * The user’s job title.
     * 
     */
    @Export(name="jobTitle", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> jobTitle;

    /**
     * @return The user’s job title.
     * 
     */
    public Output<Optional<String>> jobTitle() {
        return Codegen.optional(this.jobTitle);
    }
    /**
     * The SMTP address for the user. This property cannot be unset once specified.
     * 
     */
    @Export(name="mail", refs={String.class}, tree="[0]")
    private Output<String> mail;

    /**
     * @return The SMTP address for the user. This property cannot be unset once specified.
     * 
     */
    public Output<String> mail() {
        return this.mail;
    }
    /**
     * The mail alias for the user. Defaults to the user name part of the user principal name (UPN).
     * 
     */
    @Export(name="mailNickname", refs={String.class}, tree="[0]")
    private Output<String> mailNickname;

    /**
     * @return The mail alias for the user. Defaults to the user name part of the user principal name (UPN).
     * 
     */
    public Output<String> mailNickname() {
        return this.mailNickname;
    }
    /**
     * The object ID of the user&#39;s manager.
     * 
     */
    @Export(name="managerId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> managerId;

    /**
     * @return The object ID of the user&#39;s manager.
     * 
     */
    public Output<Optional<String>> managerId() {
        return Codegen.optional(this.managerId);
    }
    /**
     * The primary cellular telephone number for the user.
     * 
     */
    @Export(name="mobilePhone", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> mobilePhone;

    /**
     * @return The primary cellular telephone number for the user.
     * 
     */
    public Output<Optional<String>> mobilePhone() {
        return Codegen.optional(this.mobilePhone);
    }
    /**
     * The object ID of the user.
     * 
     */
    @Export(name="objectId", refs={String.class}, tree="[0]")
    private Output<String> objectId;

    /**
     * @return The object ID of the user.
     * 
     */
    public Output<String> objectId() {
        return this.objectId;
    }
    /**
     * The office location in the user&#39;s place of business.
     * 
     */
    @Export(name="officeLocation", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> officeLocation;

    /**
     * @return The office location in the user&#39;s place of business.
     * 
     */
    public Output<Optional<String>> officeLocation() {
        return Codegen.optional(this.officeLocation);
    }
    /**
     * The on-premises distinguished name (DN) of the user, synchronised from the on-premises directory when Azure AD Connect is used.
     * 
     */
    @Export(name="onpremisesDistinguishedName", refs={String.class}, tree="[0]")
    private Output<String> onpremisesDistinguishedName;

    /**
     * @return The on-premises distinguished name (DN) of the user, synchronised from the on-premises directory when Azure AD Connect is used.
     * 
     */
    public Output<String> onpremisesDistinguishedName() {
        return this.onpremisesDistinguishedName;
    }
    /**
     * The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.
     * 
     */
    @Export(name="onpremisesDomainName", refs={String.class}, tree="[0]")
    private Output<String> onpremisesDomainName;

    /**
     * @return The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.
     * 
     */
    public Output<String> onpremisesDomainName() {
        return this.onpremisesDomainName;
    }
    /**
     * The value used to associate an on-premise Active Directory user account with their Azure AD user object. This must be specified if you are using a federated domain for the user&#39;s `user_principal_name` property when creating a new user account.
     * 
     */
    @Export(name="onpremisesImmutableId", refs={String.class}, tree="[0]")
    private Output<String> onpremisesImmutableId;

    /**
     * @return The value used to associate an on-premise Active Directory user account with their Azure AD user object. This must be specified if you are using a federated domain for the user&#39;s `user_principal_name` property when creating a new user account.
     * 
     */
    public Output<String> onpremisesImmutableId() {
        return this.onpremisesImmutableId;
    }
    /**
     * The on-premise SAM account name of the user.
     * 
     */
    @Export(name="onpremisesSamAccountName", refs={String.class}, tree="[0]")
    private Output<String> onpremisesSamAccountName;

    /**
     * @return The on-premise SAM account name of the user.
     * 
     */
    public Output<String> onpremisesSamAccountName() {
        return this.onpremisesSamAccountName;
    }
    /**
     * The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.
     * 
     */
    @Export(name="onpremisesSecurityIdentifier", refs={String.class}, tree="[0]")
    private Output<String> onpremisesSecurityIdentifier;

    /**
     * @return The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.
     * 
     */
    public Output<String> onpremisesSecurityIdentifier() {
        return this.onpremisesSecurityIdentifier;
    }
    /**
     * Whether this user is synchronised from an on-premises directory (`true`), no longer synchronised (`false`), or has never been synchronised (`null`).
     * 
     */
    @Export(name="onpremisesSyncEnabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> onpremisesSyncEnabled;

    /**
     * @return Whether this user is synchronised from an on-premises directory (`true`), no longer synchronised (`false`), or has never been synchronised (`null`).
     * 
     */
    public Output<Boolean> onpremisesSyncEnabled() {
        return this.onpremisesSyncEnabled;
    }
    /**
     * The on-premise user principal name of the user.
     * 
     */
    @Export(name="onpremisesUserPrincipalName", refs={String.class}, tree="[0]")
    private Output<String> onpremisesUserPrincipalName;

    /**
     * @return The on-premise user principal name of the user.
     * 
     */
    public Output<String> onpremisesUserPrincipalName() {
        return this.onpremisesUserPrincipalName;
    }
    /**
     * A list of additional email addresses for the user.
     * 
     */
    @Export(name="otherMails", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> otherMails;

    /**
     * @return A list of additional email addresses for the user.
     * 
     */
    public Output<Optional<List<String>>> otherMails() {
        return Codegen.optional(this.otherMails);
    }
    /**
     * The password for the user. The password must satisfy minimum requirements as specified by the password policy. The
     * maximum length is 256 characters. This property is required when creating a new user
     * 
     */
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output<String> password;

    /**
     * @return The password for the user. The password must satisfy minimum requirements as specified by the password policy. The
     * maximum length is 256 characters. This property is required when creating a new user
     * 
     */
    public Output<String> password() {
        return this.password;
    }
    /**
     * The postal code for the user&#39;s postal address. The postal code is specific to the user&#39;s country/region. In the United States of America, this attribute contains the ZIP code.
     * 
     */
    @Export(name="postalCode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> postalCode;

    /**
     * @return The postal code for the user&#39;s postal address. The postal code is specific to the user&#39;s country/region. In the United States of America, this attribute contains the ZIP code.
     * 
     */
    public Output<Optional<String>> postalCode() {
        return Codegen.optional(this.postalCode);
    }
    /**
     * The user&#39;s preferred language, in ISO 639-1 notation.
     * 
     */
    @Export(name="preferredLanguage", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> preferredLanguage;

    /**
     * @return The user&#39;s preferred language, in ISO 639-1 notation.
     * 
     */
    public Output<Optional<String>> preferredLanguage() {
        return Codegen.optional(this.preferredLanguage);
    }
    /**
     * List of email addresses for the user that direct to the same mailbox.
     * 
     */
    @Export(name="proxyAddresses", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> proxyAddresses;

    /**
     * @return List of email addresses for the user that direct to the same mailbox.
     * 
     */
    public Output<List<String>> proxyAddresses() {
        return this.proxyAddresses;
    }
    /**
     * Whether or not the Outlook global address list should include this user. Defaults to `true`.
     * 
     */
    @Export(name="showInAddressList", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> showInAddressList;

    /**
     * @return Whether or not the Outlook global address list should include this user. Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> showInAddressList() {
        return Codegen.optional(this.showInAddressList);
    }
    /**
     * The state or province in the user&#39;s address.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> state;

    /**
     * @return The state or province in the user&#39;s address.
     * 
     */
    public Output<Optional<String>> state() {
        return Codegen.optional(this.state);
    }
    /**
     * The street address of the user&#39;s place of business.
     * 
     */
    @Export(name="streetAddress", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> streetAddress;

    /**
     * @return The street address of the user&#39;s place of business.
     * 
     */
    public Output<Optional<String>> streetAddress() {
        return Codegen.optional(this.streetAddress);
    }
    /**
     * The user&#39;s surname (family name or last name).
     * 
     */
    @Export(name="surname", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> surname;

    /**
     * @return The user&#39;s surname (family name or last name).
     * 
     */
    public Output<Optional<String>> surname() {
        return Codegen.optional(this.surname);
    }
    /**
     * The usage location of the user. Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries. The usage location is a two letter country code (ISO standard 3166). Examples include: `NO`, `JP`, and `GB`. Cannot be reset to null once set.
     * 
     */
    @Export(name="usageLocation", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> usageLocation;

    /**
     * @return The usage location of the user. Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries. The usage location is a two letter country code (ISO standard 3166). Examples include: `NO`, `JP`, and `GB`. Cannot be reset to null once set.
     * 
     */
    public Output<Optional<String>> usageLocation() {
        return Codegen.optional(this.usageLocation);
    }
    /**
     * The user principal name (UPN) of the user.
     * 
     */
    @Export(name="userPrincipalName", refs={String.class}, tree="[0]")
    private Output<String> userPrincipalName;

    /**
     * @return The user principal name (UPN) of the user.
     * 
     */
    public Output<String> userPrincipalName() {
        return this.userPrincipalName;
    }
    /**
     * The user type in the directory. Possible values are `Guest` or `Member`.
     * 
     */
    @Export(name="userType", refs={String.class}, tree="[0]")
    private Output<String> userType;

    /**
     * @return The user type in the directory. Possible values are `Guest` or `Member`.
     * 
     */
    public Output<String> userType() {
        return this.userType;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public User(java.lang.String name) {
        this(name, UserArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public User(java.lang.String name, UserArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public User(java.lang.String name, UserArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/user:User", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private User(java.lang.String name, Output<java.lang.String> id, @Nullable UserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/user:User", name, state, makeResourceOptions(options, id), false);
    }

    private static UserArgs makeArgs(UserArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? UserArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "password"
            ))
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
    public static User get(java.lang.String name, Output<java.lang.String> id, @Nullable UserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new User(name, id, state, options);
    }
}

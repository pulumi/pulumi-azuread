// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD
{
    /// <summary>
    /// Manages a user within Azure Active Directory.
    /// 
    /// ## API Permissions
    /// 
    /// The following API permissions are required in order to use this resource.
    /// 
    /// When authenticated with a service principal, this resource requires one of the following application roles: `User.ReadWrite.All` or `Directory.ReadWrite.All`
    /// 
    /// When authenticated with a user principal, this resource requires one of the following directory roles: `User Administrator` or `Global Administrator`
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureAD = Pulumi.AzureAD;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new AzureAD.User("example", new()
    ///     {
    ///         DisplayName = "J. Doe",
    ///         MailNickname = "jdoe",
    ///         Password = "SecretP@sswd99!",
    ///         UserPrincipalName = "jdoe@hashicorp.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Users can be imported using their object ID, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azuread:index/user:User my_user 00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureADResourceType("azuread:index/user:User")]
    public partial class User : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A freeform field for the user to describe themselves
        /// </summary>
        [Output("aboutMe")]
        public Output<string> AboutMe { get; private set; } = null!;

        /// <summary>
        /// Whether or not the account should be enabled.
        /// </summary>
        [Output("accountEnabled")]
        public Output<bool?> AccountEnabled { get; private set; } = null!;

        /// <summary>
        /// The age group of the user. Supported values are `Adult`, `NotAdult` and `Minor`. Omit this property or specify a blank string to unset.
        /// </summary>
        [Output("ageGroup")]
        public Output<string?> AgeGroup { get; private set; } = null!;

        /// <summary>
        /// A list of telephone numbers for the user. Only one number can be set for this property. Read-only for users synced with Azure AD Connect.
        /// </summary>
        [Output("businessPhones")]
        public Output<ImmutableArray<string>> BusinessPhones { get; private set; } = null!;

        /// <summary>
        /// The city in which the user is located.
        /// </summary>
        [Output("city")]
        public Output<string?> City { get; private set; } = null!;

        /// <summary>
        /// The company name which the user is associated. This property can be useful for describing the company that an external user comes from.
        /// </summary>
        [Output("companyName")]
        public Output<string?> CompanyName { get; private set; } = null!;

        /// <summary>
        /// Whether consent has been obtained for minors. Supported values are `Granted`, `Denied` and `NotRequired`. Omit this property or specify a blank string to unset.
        /// </summary>
        [Output("consentProvidedForMinor")]
        public Output<string?> ConsentProvidedForMinor { get; private set; } = null!;

        /// <summary>
        /// The cost center associated with the user.
        /// </summary>
        [Output("costCenter")]
        public Output<string?> CostCenter { get; private set; } = null!;

        /// <summary>
        /// The country/region in which the user is located. Examples include: `NO`, `JP`, and `GB`.
        /// </summary>
        [Output("country")]
        public Output<string?> Country { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the user account was created as a regular school or work account (`null`), an external account (`Invitation`), a local account for an Azure Active Directory B2C tenant (`LocalAccount`) or self-service sign-up using email verification (`EmailVerified`).
        /// </summary>
        [Output("creationType")]
        public Output<string> CreationType { get; private set; } = null!;

        /// <summary>
        /// The name for the department in which the user works.
        /// </summary>
        [Output("department")]
        public Output<string?> Department { get; private set; } = null!;

        /// <summary>
        /// Whether the user's password is exempt from expiring. Defaults to `false`.
        /// </summary>
        [Output("disablePasswordExpiration")]
        public Output<bool?> DisablePasswordExpiration { get; private set; } = null!;

        /// <summary>
        /// Whether the user is allowed weaker passwords than the default policy to be specified. Defaults to `false`.
        /// </summary>
        [Output("disableStrongPassword")]
        public Output<bool?> DisableStrongPassword { get; private set; } = null!;

        /// <summary>
        /// The name to display in the address book for the user.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The name of the division in which the user works.
        /// </summary>
        [Output("division")]
        public Output<string?> Division { get; private set; } = null!;

        /// <summary>
        /// The employee identifier assigned to the user by the organisation.
        /// </summary>
        [Output("employeeId")]
        public Output<string?> EmployeeId { get; private set; } = null!;

        /// <summary>
        /// Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor.
        /// </summary>
        [Output("employeeType")]
        public Output<string?> EmployeeType { get; private set; } = null!;

        /// <summary>
        /// For an external user invited to the tenant, this property represents the invited user's invitation status. Possible values are `PendingAcceptance` or `Accepted`.
        /// </summary>
        [Output("externalUserState")]
        public Output<string> ExternalUserState { get; private set; } = null!;

        /// <summary>
        /// The fax number of the user.
        /// </summary>
        [Output("faxNumber")]
        public Output<string?> FaxNumber { get; private set; } = null!;

        /// <summary>
        /// Whether the user is forced to change the password during the next sign-in. Only takes effect when also changing the password. Defaults to `false`.
        /// </summary>
        [Output("forcePasswordChange")]
        public Output<bool?> ForcePasswordChange { get; private set; } = null!;

        /// <summary>
        /// The given name (first name) of the user.
        /// </summary>
        [Output("givenName")]
        public Output<string?> GivenName { get; private set; } = null!;

        /// <summary>
        /// A list of instant message voice over IP (VOIP) session initiation protocol (SIP) addresses for the user.
        /// </summary>
        [Output("imAddresses")]
        public Output<ImmutableArray<string>> ImAddresses { get; private set; } = null!;

        /// <summary>
        /// The user’s job title.
        /// </summary>
        [Output("jobTitle")]
        public Output<string?> JobTitle { get; private set; } = null!;

        /// <summary>
        /// The SMTP address for the user. This property cannot be unset once specified.
        /// </summary>
        [Output("mail")]
        public Output<string> Mail { get; private set; } = null!;

        /// <summary>
        /// The mail alias for the user. Defaults to the user name part of the user principal name (UPN).
        /// </summary>
        [Output("mailNickname")]
        public Output<string> MailNickname { get; private set; } = null!;

        /// <summary>
        /// The object ID of the user's manager.
        /// </summary>
        [Output("managerId")]
        public Output<string?> ManagerId { get; private set; } = null!;

        /// <summary>
        /// The primary cellular telephone number for the user.
        /// </summary>
        [Output("mobilePhone")]
        public Output<string?> MobilePhone { get; private set; } = null!;

        /// <summary>
        /// The object ID of the user.
        /// </summary>
        [Output("objectId")]
        public Output<string> ObjectId { get; private set; } = null!;

        /// <summary>
        /// The office location in the user's place of business.
        /// </summary>
        [Output("officeLocation")]
        public Output<string?> OfficeLocation { get; private set; } = null!;

        /// <summary>
        /// The on-premises distinguished name (DN) of the user, synchronised from the on-premises directory when Azure AD Connect is used.
        /// </summary>
        [Output("onpremisesDistinguishedName")]
        public Output<string> OnpremisesDistinguishedName { get; private set; } = null!;

        /// <summary>
        /// The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.
        /// </summary>
        [Output("onpremisesDomainName")]
        public Output<string> OnpremisesDomainName { get; private set; } = null!;

        /// <summary>
        /// The value used to associate an on-premise Active Directory user account with their Azure AD user object. This must be specified if you are using a federated domain for the user's `user_principal_name` property when creating a new user account.
        /// </summary>
        [Output("onpremisesImmutableId")]
        public Output<string> OnpremisesImmutableId { get; private set; } = null!;

        /// <summary>
        /// The on-premise SAM account name of the user.
        /// </summary>
        [Output("onpremisesSamAccountName")]
        public Output<string> OnpremisesSamAccountName { get; private set; } = null!;

        /// <summary>
        /// The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.
        /// </summary>
        [Output("onpremisesSecurityIdentifier")]
        public Output<string> OnpremisesSecurityIdentifier { get; private set; } = null!;

        /// <summary>
        /// Whether this user is synchronised from an on-premises directory (`true`), no longer synchronised (`false`), or has never been synchronised (`null`).
        /// </summary>
        [Output("onpremisesSyncEnabled")]
        public Output<bool> OnpremisesSyncEnabled { get; private set; } = null!;

        /// <summary>
        /// The on-premise user principal name of the user.
        /// </summary>
        [Output("onpremisesUserPrincipalName")]
        public Output<string> OnpremisesUserPrincipalName { get; private set; } = null!;

        /// <summary>
        /// A list of additional email addresses for the user.
        /// </summary>
        [Output("otherMails")]
        public Output<ImmutableArray<string>> OtherMails { get; private set; } = null!;

        /// <summary>
        /// The password for the user. The password must satisfy minimum requirements as specified by the password policy. The
        /// maximum length is 256 characters. This property is required when creating a new user
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code.
        /// </summary>
        [Output("postalCode")]
        public Output<string?> PostalCode { get; private set; } = null!;

        /// <summary>
        /// The user's preferred language, in ISO 639-1 notation.
        /// </summary>
        [Output("preferredLanguage")]
        public Output<string?> PreferredLanguage { get; private set; } = null!;

        /// <summary>
        /// List of email addresses for the user that direct to the same mailbox.
        /// </summary>
        [Output("proxyAddresses")]
        public Output<ImmutableArray<string>> ProxyAddresses { get; private set; } = null!;

        /// <summary>
        /// Whether or not the Outlook global address list should include this user. Defaults to `true`.
        /// </summary>
        [Output("showInAddressList")]
        public Output<bool?> ShowInAddressList { get; private set; } = null!;

        /// <summary>
        /// The state or province in the user's address.
        /// </summary>
        [Output("state")]
        public Output<string?> State { get; private set; } = null!;

        /// <summary>
        /// The street address of the user's place of business.
        /// </summary>
        [Output("streetAddress")]
        public Output<string?> StreetAddress { get; private set; } = null!;

        /// <summary>
        /// The user's surname (family name or last name).
        /// </summary>
        [Output("surname")]
        public Output<string?> Surname { get; private set; } = null!;

        /// <summary>
        /// The usage location of the user. Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries. The usage location is a two letter country code (ISO standard 3166). Examples include: `NO`, `JP`, and `GB`. Cannot be reset to null once set.
        /// </summary>
        [Output("usageLocation")]
        public Output<string?> UsageLocation { get; private set; } = null!;

        /// <summary>
        /// The user principal name (UPN) of the user.
        /// </summary>
        [Output("userPrincipalName")]
        public Output<string> UserPrincipalName { get; private set; } = null!;

        /// <summary>
        /// The user type in the directory. Possible values are `Guest` or `Member`.
        /// </summary>
        [Output("userType")]
        public Output<string> UserType { get; private set; } = null!;


        /// <summary>
        /// Create a User resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public User(string name, UserArgs args, CustomResourceOptions? options = null)
            : base("azuread:index/user:User", name, args ?? new UserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private User(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
            : base("azuread:index/user:User", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing User resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static User Get(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
        {
            return new User(name, id, state, options);
        }
    }

    public sealed class UserArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether or not the account should be enabled.
        /// </summary>
        [Input("accountEnabled")]
        public Input<bool>? AccountEnabled { get; set; }

        /// <summary>
        /// The age group of the user. Supported values are `Adult`, `NotAdult` and `Minor`. Omit this property or specify a blank string to unset.
        /// </summary>
        [Input("ageGroup")]
        public Input<string>? AgeGroup { get; set; }

        [Input("businessPhones")]
        private InputList<string>? _businessPhones;

        /// <summary>
        /// A list of telephone numbers for the user. Only one number can be set for this property. Read-only for users synced with Azure AD Connect.
        /// </summary>
        public InputList<string> BusinessPhones
        {
            get => _businessPhones ?? (_businessPhones = new InputList<string>());
            set => _businessPhones = value;
        }

        /// <summary>
        /// The city in which the user is located.
        /// </summary>
        [Input("city")]
        public Input<string>? City { get; set; }

        /// <summary>
        /// The company name which the user is associated. This property can be useful for describing the company that an external user comes from.
        /// </summary>
        [Input("companyName")]
        public Input<string>? CompanyName { get; set; }

        /// <summary>
        /// Whether consent has been obtained for minors. Supported values are `Granted`, `Denied` and `NotRequired`. Omit this property or specify a blank string to unset.
        /// </summary>
        [Input("consentProvidedForMinor")]
        public Input<string>? ConsentProvidedForMinor { get; set; }

        /// <summary>
        /// The cost center associated with the user.
        /// </summary>
        [Input("costCenter")]
        public Input<string>? CostCenter { get; set; }

        /// <summary>
        /// The country/region in which the user is located. Examples include: `NO`, `JP`, and `GB`.
        /// </summary>
        [Input("country")]
        public Input<string>? Country { get; set; }

        /// <summary>
        /// The name for the department in which the user works.
        /// </summary>
        [Input("department")]
        public Input<string>? Department { get; set; }

        /// <summary>
        /// Whether the user's password is exempt from expiring. Defaults to `false`.
        /// </summary>
        [Input("disablePasswordExpiration")]
        public Input<bool>? DisablePasswordExpiration { get; set; }

        /// <summary>
        /// Whether the user is allowed weaker passwords than the default policy to be specified. Defaults to `false`.
        /// </summary>
        [Input("disableStrongPassword")]
        public Input<bool>? DisableStrongPassword { get; set; }

        /// <summary>
        /// The name to display in the address book for the user.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// The name of the division in which the user works.
        /// </summary>
        [Input("division")]
        public Input<string>? Division { get; set; }

        /// <summary>
        /// The employee identifier assigned to the user by the organisation.
        /// </summary>
        [Input("employeeId")]
        public Input<string>? EmployeeId { get; set; }

        /// <summary>
        /// Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor.
        /// </summary>
        [Input("employeeType")]
        public Input<string>? EmployeeType { get; set; }

        /// <summary>
        /// The fax number of the user.
        /// </summary>
        [Input("faxNumber")]
        public Input<string>? FaxNumber { get; set; }

        /// <summary>
        /// Whether the user is forced to change the password during the next sign-in. Only takes effect when also changing the password. Defaults to `false`.
        /// </summary>
        [Input("forcePasswordChange")]
        public Input<bool>? ForcePasswordChange { get; set; }

        /// <summary>
        /// The given name (first name) of the user.
        /// </summary>
        [Input("givenName")]
        public Input<string>? GivenName { get; set; }

        /// <summary>
        /// The user’s job title.
        /// </summary>
        [Input("jobTitle")]
        public Input<string>? JobTitle { get; set; }

        /// <summary>
        /// The SMTP address for the user. This property cannot be unset once specified.
        /// </summary>
        [Input("mail")]
        public Input<string>? Mail { get; set; }

        /// <summary>
        /// The mail alias for the user. Defaults to the user name part of the user principal name (UPN).
        /// </summary>
        [Input("mailNickname")]
        public Input<string>? MailNickname { get; set; }

        /// <summary>
        /// The object ID of the user's manager.
        /// </summary>
        [Input("managerId")]
        public Input<string>? ManagerId { get; set; }

        /// <summary>
        /// The primary cellular telephone number for the user.
        /// </summary>
        [Input("mobilePhone")]
        public Input<string>? MobilePhone { get; set; }

        /// <summary>
        /// The office location in the user's place of business.
        /// </summary>
        [Input("officeLocation")]
        public Input<string>? OfficeLocation { get; set; }

        /// <summary>
        /// The value used to associate an on-premise Active Directory user account with their Azure AD user object. This must be specified if you are using a federated domain for the user's `user_principal_name` property when creating a new user account.
        /// </summary>
        [Input("onpremisesImmutableId")]
        public Input<string>? OnpremisesImmutableId { get; set; }

        [Input("otherMails")]
        private InputList<string>? _otherMails;

        /// <summary>
        /// A list of additional email addresses for the user.
        /// </summary>
        public InputList<string> OtherMails
        {
            get => _otherMails ?? (_otherMails = new InputList<string>());
            set => _otherMails = value;
        }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// The password for the user. The password must satisfy minimum requirements as specified by the password policy. The
        /// maximum length is 256 characters. This property is required when creating a new user
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code.
        /// </summary>
        [Input("postalCode")]
        public Input<string>? PostalCode { get; set; }

        /// <summary>
        /// The user's preferred language, in ISO 639-1 notation.
        /// </summary>
        [Input("preferredLanguage")]
        public Input<string>? PreferredLanguage { get; set; }

        /// <summary>
        /// Whether or not the Outlook global address list should include this user. Defaults to `true`.
        /// </summary>
        [Input("showInAddressList")]
        public Input<bool>? ShowInAddressList { get; set; }

        /// <summary>
        /// The state or province in the user's address.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The street address of the user's place of business.
        /// </summary>
        [Input("streetAddress")]
        public Input<string>? StreetAddress { get; set; }

        /// <summary>
        /// The user's surname (family name or last name).
        /// </summary>
        [Input("surname")]
        public Input<string>? Surname { get; set; }

        /// <summary>
        /// The usage location of the user. Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries. The usage location is a two letter country code (ISO standard 3166). Examples include: `NO`, `JP`, and `GB`. Cannot be reset to null once set.
        /// </summary>
        [Input("usageLocation")]
        public Input<string>? UsageLocation { get; set; }

        /// <summary>
        /// The user principal name (UPN) of the user.
        /// </summary>
        [Input("userPrincipalName", required: true)]
        public Input<string> UserPrincipalName { get; set; } = null!;

        public UserArgs()
        {
        }
        public static new UserArgs Empty => new UserArgs();
    }

    public sealed class UserState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A freeform field for the user to describe themselves
        /// </summary>
        [Input("aboutMe")]
        public Input<string>? AboutMe { get; set; }

        /// <summary>
        /// Whether or not the account should be enabled.
        /// </summary>
        [Input("accountEnabled")]
        public Input<bool>? AccountEnabled { get; set; }

        /// <summary>
        /// The age group of the user. Supported values are `Adult`, `NotAdult` and `Minor`. Omit this property or specify a blank string to unset.
        /// </summary>
        [Input("ageGroup")]
        public Input<string>? AgeGroup { get; set; }

        [Input("businessPhones")]
        private InputList<string>? _businessPhones;

        /// <summary>
        /// A list of telephone numbers for the user. Only one number can be set for this property. Read-only for users synced with Azure AD Connect.
        /// </summary>
        public InputList<string> BusinessPhones
        {
            get => _businessPhones ?? (_businessPhones = new InputList<string>());
            set => _businessPhones = value;
        }

        /// <summary>
        /// The city in which the user is located.
        /// </summary>
        [Input("city")]
        public Input<string>? City { get; set; }

        /// <summary>
        /// The company name which the user is associated. This property can be useful for describing the company that an external user comes from.
        /// </summary>
        [Input("companyName")]
        public Input<string>? CompanyName { get; set; }

        /// <summary>
        /// Whether consent has been obtained for minors. Supported values are `Granted`, `Denied` and `NotRequired`. Omit this property or specify a blank string to unset.
        /// </summary>
        [Input("consentProvidedForMinor")]
        public Input<string>? ConsentProvidedForMinor { get; set; }

        /// <summary>
        /// The cost center associated with the user.
        /// </summary>
        [Input("costCenter")]
        public Input<string>? CostCenter { get; set; }

        /// <summary>
        /// The country/region in which the user is located. Examples include: `NO`, `JP`, and `GB`.
        /// </summary>
        [Input("country")]
        public Input<string>? Country { get; set; }

        /// <summary>
        /// Indicates whether the user account was created as a regular school or work account (`null`), an external account (`Invitation`), a local account for an Azure Active Directory B2C tenant (`LocalAccount`) or self-service sign-up using email verification (`EmailVerified`).
        /// </summary>
        [Input("creationType")]
        public Input<string>? CreationType { get; set; }

        /// <summary>
        /// The name for the department in which the user works.
        /// </summary>
        [Input("department")]
        public Input<string>? Department { get; set; }

        /// <summary>
        /// Whether the user's password is exempt from expiring. Defaults to `false`.
        /// </summary>
        [Input("disablePasswordExpiration")]
        public Input<bool>? DisablePasswordExpiration { get; set; }

        /// <summary>
        /// Whether the user is allowed weaker passwords than the default policy to be specified. Defaults to `false`.
        /// </summary>
        [Input("disableStrongPassword")]
        public Input<bool>? DisableStrongPassword { get; set; }

        /// <summary>
        /// The name to display in the address book for the user.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The name of the division in which the user works.
        /// </summary>
        [Input("division")]
        public Input<string>? Division { get; set; }

        /// <summary>
        /// The employee identifier assigned to the user by the organisation.
        /// </summary>
        [Input("employeeId")]
        public Input<string>? EmployeeId { get; set; }

        /// <summary>
        /// Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor.
        /// </summary>
        [Input("employeeType")]
        public Input<string>? EmployeeType { get; set; }

        /// <summary>
        /// For an external user invited to the tenant, this property represents the invited user's invitation status. Possible values are `PendingAcceptance` or `Accepted`.
        /// </summary>
        [Input("externalUserState")]
        public Input<string>? ExternalUserState { get; set; }

        /// <summary>
        /// The fax number of the user.
        /// </summary>
        [Input("faxNumber")]
        public Input<string>? FaxNumber { get; set; }

        /// <summary>
        /// Whether the user is forced to change the password during the next sign-in. Only takes effect when also changing the password. Defaults to `false`.
        /// </summary>
        [Input("forcePasswordChange")]
        public Input<bool>? ForcePasswordChange { get; set; }

        /// <summary>
        /// The given name (first name) of the user.
        /// </summary>
        [Input("givenName")]
        public Input<string>? GivenName { get; set; }

        [Input("imAddresses")]
        private InputList<string>? _imAddresses;

        /// <summary>
        /// A list of instant message voice over IP (VOIP) session initiation protocol (SIP) addresses for the user.
        /// </summary>
        public InputList<string> ImAddresses
        {
            get => _imAddresses ?? (_imAddresses = new InputList<string>());
            set => _imAddresses = value;
        }

        /// <summary>
        /// The user’s job title.
        /// </summary>
        [Input("jobTitle")]
        public Input<string>? JobTitle { get; set; }

        /// <summary>
        /// The SMTP address for the user. This property cannot be unset once specified.
        /// </summary>
        [Input("mail")]
        public Input<string>? Mail { get; set; }

        /// <summary>
        /// The mail alias for the user. Defaults to the user name part of the user principal name (UPN).
        /// </summary>
        [Input("mailNickname")]
        public Input<string>? MailNickname { get; set; }

        /// <summary>
        /// The object ID of the user's manager.
        /// </summary>
        [Input("managerId")]
        public Input<string>? ManagerId { get; set; }

        /// <summary>
        /// The primary cellular telephone number for the user.
        /// </summary>
        [Input("mobilePhone")]
        public Input<string>? MobilePhone { get; set; }

        /// <summary>
        /// The object ID of the user.
        /// </summary>
        [Input("objectId")]
        public Input<string>? ObjectId { get; set; }

        /// <summary>
        /// The office location in the user's place of business.
        /// </summary>
        [Input("officeLocation")]
        public Input<string>? OfficeLocation { get; set; }

        /// <summary>
        /// The on-premises distinguished name (DN) of the user, synchronised from the on-premises directory when Azure AD Connect is used.
        /// </summary>
        [Input("onpremisesDistinguishedName")]
        public Input<string>? OnpremisesDistinguishedName { get; set; }

        /// <summary>
        /// The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.
        /// </summary>
        [Input("onpremisesDomainName")]
        public Input<string>? OnpremisesDomainName { get; set; }

        /// <summary>
        /// The value used to associate an on-premise Active Directory user account with their Azure AD user object. This must be specified if you are using a federated domain for the user's `user_principal_name` property when creating a new user account.
        /// </summary>
        [Input("onpremisesImmutableId")]
        public Input<string>? OnpremisesImmutableId { get; set; }

        /// <summary>
        /// The on-premise SAM account name of the user.
        /// </summary>
        [Input("onpremisesSamAccountName")]
        public Input<string>? OnpremisesSamAccountName { get; set; }

        /// <summary>
        /// The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.
        /// </summary>
        [Input("onpremisesSecurityIdentifier")]
        public Input<string>? OnpremisesSecurityIdentifier { get; set; }

        /// <summary>
        /// Whether this user is synchronised from an on-premises directory (`true`), no longer synchronised (`false`), or has never been synchronised (`null`).
        /// </summary>
        [Input("onpremisesSyncEnabled")]
        public Input<bool>? OnpremisesSyncEnabled { get; set; }

        /// <summary>
        /// The on-premise user principal name of the user.
        /// </summary>
        [Input("onpremisesUserPrincipalName")]
        public Input<string>? OnpremisesUserPrincipalName { get; set; }

        [Input("otherMails")]
        private InputList<string>? _otherMails;

        /// <summary>
        /// A list of additional email addresses for the user.
        /// </summary>
        public InputList<string> OtherMails
        {
            get => _otherMails ?? (_otherMails = new InputList<string>());
            set => _otherMails = value;
        }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// The password for the user. The password must satisfy minimum requirements as specified by the password policy. The
        /// maximum length is 256 characters. This property is required when creating a new user
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code.
        /// </summary>
        [Input("postalCode")]
        public Input<string>? PostalCode { get; set; }

        /// <summary>
        /// The user's preferred language, in ISO 639-1 notation.
        /// </summary>
        [Input("preferredLanguage")]
        public Input<string>? PreferredLanguage { get; set; }

        [Input("proxyAddresses")]
        private InputList<string>? _proxyAddresses;

        /// <summary>
        /// List of email addresses for the user that direct to the same mailbox.
        /// </summary>
        public InputList<string> ProxyAddresses
        {
            get => _proxyAddresses ?? (_proxyAddresses = new InputList<string>());
            set => _proxyAddresses = value;
        }

        /// <summary>
        /// Whether or not the Outlook global address list should include this user. Defaults to `true`.
        /// </summary>
        [Input("showInAddressList")]
        public Input<bool>? ShowInAddressList { get; set; }

        /// <summary>
        /// The state or province in the user's address.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The street address of the user's place of business.
        /// </summary>
        [Input("streetAddress")]
        public Input<string>? StreetAddress { get; set; }

        /// <summary>
        /// The user's surname (family name or last name).
        /// </summary>
        [Input("surname")]
        public Input<string>? Surname { get; set; }

        /// <summary>
        /// The usage location of the user. Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries. The usage location is a two letter country code (ISO standard 3166). Examples include: `NO`, `JP`, and `GB`. Cannot be reset to null once set.
        /// </summary>
        [Input("usageLocation")]
        public Input<string>? UsageLocation { get; set; }

        /// <summary>
        /// The user principal name (UPN) of the user.
        /// </summary>
        [Input("userPrincipalName")]
        public Input<string>? UserPrincipalName { get; set; }

        /// <summary>
        /// The user type in the directory. Possible values are `Guest` or `Member`.
        /// </summary>
        [Input("userType")]
        public Input<string>? UserType { get; set; }

        public UserState()
        {
        }
        public static new UserState Empty => new UserState();
    }
}

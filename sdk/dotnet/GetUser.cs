// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD
{
    public static class GetUser
    {
        /// <summary>
        /// Gets information about an Azure Active Directory user.
        /// 
        /// ## API Permissions
        /// 
        /// The following API permissions are required in order to use this data source.
        /// 
        /// When authenticated with a service principal, this data source requires one of the following application roles: `User.Read.All` or `Directory.Read.All`
        /// 
        /// When authenticated with a user principal, this data source does not require any additional roles.
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
        ///     var example = AzureAD.GetUser.Invoke(new()
        ///     {
        ///         UserPrincipalName = "user@example.com",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetUserResult> InvokeAsync(GetUserArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUserResult>("azuread:index/getUser:getUser", args ?? new GetUserArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about an Azure Active Directory user.
        /// 
        /// ## API Permissions
        /// 
        /// The following API permissions are required in order to use this data source.
        /// 
        /// When authenticated with a service principal, this data source requires one of the following application roles: `User.Read.All` or `Directory.Read.All`
        /// 
        /// When authenticated with a user principal, this data source does not require any additional roles.
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
        ///     var example = AzureAD.GetUser.Invoke(new()
        ///     {
        ///         UserPrincipalName = "user@example.com",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetUserResult> Invoke(GetUserInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUserResult>("azuread:index/getUser:getUser", args ?? new GetUserInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about an Azure Active Directory user.
        /// 
        /// ## API Permissions
        /// 
        /// The following API permissions are required in order to use this data source.
        /// 
        /// When authenticated with a service principal, this data source requires one of the following application roles: `User.Read.All` or `Directory.Read.All`
        /// 
        /// When authenticated with a user principal, this data source does not require any additional roles.
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
        ///     var example = AzureAD.GetUser.Invoke(new()
        ///     {
        ///         UserPrincipalName = "user@example.com",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetUserResult> Invoke(GetUserInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetUserResult>("azuread:index/getUser:getUser", args ?? new GetUserInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUserArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The employee identifier assigned to the user by the organisation.
        /// </summary>
        [Input("employeeId")]
        public string? EmployeeId { get; set; }

        /// <summary>
        /// The SMTP address for the user.
        /// </summary>
        [Input("mail")]
        public string? Mail { get; set; }

        /// <summary>
        /// The email alias of the user.
        /// </summary>
        [Input("mailNickname")]
        public string? MailNickname { get; set; }

        /// <summary>
        /// The object ID of the user.
        /// </summary>
        [Input("objectId")]
        public string? ObjectId { get; set; }

        /// <summary>
        /// The user principal name (UPN) of the user.
        /// 
        /// &gt; One of `user_principal_name`, `object_id`, `mail`, `mail_nickname` or `employee_id` must be specified.
        /// </summary>
        [Input("userPrincipalName")]
        public string? UserPrincipalName { get; set; }

        public GetUserArgs()
        {
        }
        public static new GetUserArgs Empty => new GetUserArgs();
    }

    public sealed class GetUserInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The employee identifier assigned to the user by the organisation.
        /// </summary>
        [Input("employeeId")]
        public Input<string>? EmployeeId { get; set; }

        /// <summary>
        /// The SMTP address for the user.
        /// </summary>
        [Input("mail")]
        public Input<string>? Mail { get; set; }

        /// <summary>
        /// The email alias of the user.
        /// </summary>
        [Input("mailNickname")]
        public Input<string>? MailNickname { get; set; }

        /// <summary>
        /// The object ID of the user.
        /// </summary>
        [Input("objectId")]
        public Input<string>? ObjectId { get; set; }

        /// <summary>
        /// The user principal name (UPN) of the user.
        /// 
        /// &gt; One of `user_principal_name`, `object_id`, `mail`, `mail_nickname` or `employee_id` must be specified.
        /// </summary>
        [Input("userPrincipalName")]
        public Input<string>? UserPrincipalName { get; set; }

        public GetUserInvokeArgs()
        {
        }
        public static new GetUserInvokeArgs Empty => new GetUserInvokeArgs();
    }


    [OutputType]
    public sealed class GetUserResult
    {
        /// <summary>
        /// Whether or not the account is enabled.
        /// </summary>
        public readonly bool AccountEnabled;
        /// <summary>
        /// The age group of the user. Supported values are `Adult`, `NotAdult` and `Minor`.
        /// </summary>
        public readonly string AgeGroup;
        /// <summary>
        /// A list of telephone numbers for the user.
        /// </summary>
        public readonly ImmutableArray<string> BusinessPhones;
        /// <summary>
        /// The city in which the user is located.
        /// </summary>
        public readonly string City;
        /// <summary>
        /// The company name which the user is associated. This property can be useful for describing the company that an external user comes from.
        /// </summary>
        public readonly string CompanyName;
        /// <summary>
        /// Whether consent has been obtained for minors. Supported values are `Granted`, `Denied` and `NotRequired`.
        /// </summary>
        public readonly string ConsentProvidedForMinor;
        /// <summary>
        /// The cost center associated with the user.
        /// </summary>
        public readonly string CostCenter;
        /// <summary>
        /// The country/region in which the user is located, e.g. `US` or `UK`.
        /// </summary>
        public readonly string Country;
        /// <summary>
        /// Indicates whether the user account was created as a regular school or work account (`null`), an external account (`Invitation`), a local account for an Azure Active Directory B2C tenant (`LocalAccount`) or self-service sign-up using email verification (`EmailVerified`).
        /// </summary>
        public readonly string CreationType;
        /// <summary>
        /// The name for the department in which the user works.
        /// </summary>
        public readonly string Department;
        /// <summary>
        /// The display name of the user.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The name of the division in which the user works.
        /// </summary>
        public readonly string Division;
        /// <summary>
        /// The hire date of the user, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
        /// </summary>
        public readonly string EmployeeHireDate;
        /// <summary>
        /// The employee identifier assigned to the user by the organisation.
        /// </summary>
        public readonly string EmployeeId;
        /// <summary>
        /// Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor.
        /// </summary>
        public readonly string EmployeeType;
        /// <summary>
        /// For an external user invited to the tenant, this property represents the invited user's invitation status. Possible values are `PendingAcceptance` or `Accepted`.
        /// </summary>
        public readonly string ExternalUserState;
        /// <summary>
        /// The fax number of the user.
        /// </summary>
        public readonly string FaxNumber;
        /// <summary>
        /// The given name (first name) of the user.
        /// </summary>
        public readonly string GivenName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of instant message voice over IP (VOIP) session initiation protocol (SIP) addresses for the user.
        /// </summary>
        public readonly ImmutableArray<string> ImAddresses;
        /// <summary>
        /// The user’s job title.
        /// </summary>
        public readonly string JobTitle;
        /// <summary>
        /// The SMTP address for the user.
        /// </summary>
        public readonly string Mail;
        /// <summary>
        /// The email alias of the user.
        /// </summary>
        public readonly string MailNickname;
        /// <summary>
        /// The object ID of the user's manager.
        /// </summary>
        public readonly string ManagerId;
        /// <summary>
        /// The primary cellular telephone number for the user.
        /// </summary>
        public readonly string MobilePhone;
        /// <summary>
        /// The object ID of the user.
        /// </summary>
        public readonly string ObjectId;
        /// <summary>
        /// The office location in the user's place of business.
        /// </summary>
        public readonly string OfficeLocation;
        /// <summary>
        /// The on-premises distinguished name (DN) of the user, synchronised from the on-premises directory when Azure AD Connect is used.
        /// </summary>
        public readonly string OnpremisesDistinguishedName;
        /// <summary>
        /// The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.
        /// </summary>
        public readonly string OnpremisesDomainName;
        /// <summary>
        /// The value used to associate an on-premise Active Directory user account with their Azure AD user object.
        /// </summary>
        public readonly string OnpremisesImmutableId;
        /// <summary>
        /// The on-premise SAM account name of the user.
        /// </summary>
        public readonly string OnpremisesSamAccountName;
        /// <summary>
        /// The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.
        /// </summary>
        public readonly string OnpremisesSecurityIdentifier;
        /// <summary>
        /// Whether this user is synchronised from an on-premises directory (`true`), no longer synchronised (`false`), or has never been synchronised (`null`).
        /// </summary>
        public readonly bool OnpremisesSyncEnabled;
        /// <summary>
        /// The on-premise user principal name of the user.
        /// </summary>
        public readonly string OnpremisesUserPrincipalName;
        /// <summary>
        /// A list of additional email addresses for the user.
        /// </summary>
        public readonly ImmutableArray<string> OtherMails;
        /// <summary>
        /// The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code.
        /// </summary>
        public readonly string PostalCode;
        /// <summary>
        /// The user's preferred language, in ISO 639-1 notation.
        /// </summary>
        public readonly string PreferredLanguage;
        /// <summary>
        /// List of email addresses for the user that direct to the same mailbox.
        /// </summary>
        public readonly ImmutableArray<string> ProxyAddresses;
        /// <summary>
        /// Whether or not the Outlook global address list should include this user.
        /// </summary>
        public readonly bool ShowInAddressList;
        /// <summary>
        /// The state or province in the user's address.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The street address of the user's place of business.
        /// </summary>
        public readonly string StreetAddress;
        /// <summary>
        /// The user's surname (family name or last name).
        /// </summary>
        public readonly string Surname;
        /// <summary>
        /// The usage location of the user.
        /// </summary>
        public readonly string UsageLocation;
        /// <summary>
        /// The user principal name (UPN) of the user.
        /// </summary>
        public readonly string UserPrincipalName;
        /// <summary>
        /// The user type in the directory. Possible values are `Guest` or `Member`.
        /// </summary>
        public readonly string UserType;

        [OutputConstructor]
        private GetUserResult(
            bool accountEnabled,

            string ageGroup,

            ImmutableArray<string> businessPhones,

            string city,

            string companyName,

            string consentProvidedForMinor,

            string costCenter,

            string country,

            string creationType,

            string department,

            string displayName,

            string division,

            string employeeHireDate,

            string employeeId,

            string employeeType,

            string externalUserState,

            string faxNumber,

            string givenName,

            string id,

            ImmutableArray<string> imAddresses,

            string jobTitle,

            string mail,

            string mailNickname,

            string managerId,

            string mobilePhone,

            string objectId,

            string officeLocation,

            string onpremisesDistinguishedName,

            string onpremisesDomainName,

            string onpremisesImmutableId,

            string onpremisesSamAccountName,

            string onpremisesSecurityIdentifier,

            bool onpremisesSyncEnabled,

            string onpremisesUserPrincipalName,

            ImmutableArray<string> otherMails,

            string postalCode,

            string preferredLanguage,

            ImmutableArray<string> proxyAddresses,

            bool showInAddressList,

            string state,

            string streetAddress,

            string surname,

            string usageLocation,

            string userPrincipalName,

            string userType)
        {
            AccountEnabled = accountEnabled;
            AgeGroup = ageGroup;
            BusinessPhones = businessPhones;
            City = city;
            CompanyName = companyName;
            ConsentProvidedForMinor = consentProvidedForMinor;
            CostCenter = costCenter;
            Country = country;
            CreationType = creationType;
            Department = department;
            DisplayName = displayName;
            Division = division;
            EmployeeHireDate = employeeHireDate;
            EmployeeId = employeeId;
            EmployeeType = employeeType;
            ExternalUserState = externalUserState;
            FaxNumber = faxNumber;
            GivenName = givenName;
            Id = id;
            ImAddresses = imAddresses;
            JobTitle = jobTitle;
            Mail = mail;
            MailNickname = mailNickname;
            ManagerId = managerId;
            MobilePhone = mobilePhone;
            ObjectId = objectId;
            OfficeLocation = officeLocation;
            OnpremisesDistinguishedName = onpremisesDistinguishedName;
            OnpremisesDomainName = onpremisesDomainName;
            OnpremisesImmutableId = onpremisesImmutableId;
            OnpremisesSamAccountName = onpremisesSamAccountName;
            OnpremisesSecurityIdentifier = onpremisesSecurityIdentifier;
            OnpremisesSyncEnabled = onpremisesSyncEnabled;
            OnpremisesUserPrincipalName = onpremisesUserPrincipalName;
            OtherMails = otherMails;
            PostalCode = postalCode;
            PreferredLanguage = preferredLanguage;
            ProxyAddresses = proxyAddresses;
            ShowInAddressList = showInAddressList;
            State = state;
            StreetAddress = streetAddress;
            Surname = surname;
            UsageLocation = usageLocation;
            UserPrincipalName = userPrincipalName;
            UserType = userType;
        }
    }
}

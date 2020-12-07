// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
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
        /// &gt; **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `Read directory data` within the `Windows Azure Active Directory` API.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AzureAD = Pulumi.AzureAD;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(AzureAD.GetUser.InvokeAsync(new AzureAD.GetUserArgs
        ///         {
        ///             UserPrincipalName = "user@hashicorp.com",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetUserResult> InvokeAsync(GetUserArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetUserResult>("azuread:index/getUser:getUser", args ?? new GetUserArgs(), options.WithVersion());
    }


    public sealed class GetUserArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The email alias of the Azure AD User.
        /// </summary>
        [Input("mailNickname")]
        public string? MailNickname { get; set; }

        /// <summary>
        /// Specifies the Object ID of the User within Azure Active Directory.
        /// </summary>
        [Input("objectId")]
        public string? ObjectId { get; set; }

        /// <summary>
        /// The User Principal Name of the Azure AD User.
        /// </summary>
        [Input("userPrincipalName")]
        public string? UserPrincipalName { get; set; }

        public GetUserArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetUserResult
    {
        /// <summary>
        /// `True` if the account is enabled; otherwise `False`.
        /// </summary>
        public readonly bool AccountEnabled;
        /// <summary>
        /// The city in which the user is located.
        /// </summary>
        public readonly string City;
        /// <summary>
        /// The company name which the user is associated. This property can be useful for describing the company that an external user comes from.
        /// </summary>
        public readonly string CompanyName;
        /// <summary>
        /// The country/region in which the user is located; for example, “US” or “UK”.
        /// </summary>
        public readonly string Country;
        /// <summary>
        /// The name for the department in which the user works.
        /// </summary>
        public readonly string Department;
        /// <summary>
        /// The Display Name of the Azure AD User.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The given name (first name) of the user.
        /// </summary>
        public readonly string GivenName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The value used to associate an on-premise Active Directory user account with their Azure AD user object.
        /// </summary>
        public readonly string ImmutableId;
        /// <summary>
        /// The user’s job title.
        /// </summary>
        public readonly string JobTitle;
        /// <summary>
        /// The primary email address of the Azure AD User.
        /// </summary>
        public readonly string Mail;
        /// <summary>
        /// The email alias of the Azure AD User.
        /// </summary>
        public readonly string MailNickname;
        /// <summary>
        /// The primary cellular telephone number for the user.
        /// </summary>
        public readonly string Mobile;
        public readonly string ObjectId;
        /// <summary>
        /// The on-premise SAM account name of the Azure AD User.
        /// </summary>
        public readonly string OnpremisesSamAccountName;
        /// <summary>
        /// The on-premise user principal name of the Azure AD User.
        /// </summary>
        public readonly string OnpremisesUserPrincipalName;
        /// <summary>
        /// The office location in the user's place of business.
        /// </summary>
        public readonly string PhysicalDeliveryOfficeName;
        /// <summary>
        /// The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code.
        /// </summary>
        public readonly string PostalCode;
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
        /// The usage location of the Azure AD User.
        /// </summary>
        public readonly string UsageLocation;
        /// <summary>
        /// The User Principal Name of the Azure AD User.
        /// </summary>
        public readonly string UserPrincipalName;

        [OutputConstructor]
        private GetUserResult(
            bool accountEnabled,

            string city,

            string companyName,

            string country,

            string department,

            string displayName,

            string givenName,

            string id,

            string immutableId,

            string jobTitle,

            string mail,

            string mailNickname,

            string mobile,

            string objectId,

            string onpremisesSamAccountName,

            string onpremisesUserPrincipalName,

            string physicalDeliveryOfficeName,

            string postalCode,

            string state,

            string streetAddress,

            string surname,

            string usageLocation,

            string userPrincipalName)
        {
            AccountEnabled = accountEnabled;
            City = city;
            CompanyName = companyName;
            Country = country;
            Department = department;
            DisplayName = displayName;
            GivenName = givenName;
            Id = id;
            ImmutableId = immutableId;
            JobTitle = jobTitle;
            Mail = mail;
            MailNickname = mailNickname;
            Mobile = mobile;
            ObjectId = objectId;
            OnpremisesSamAccountName = onpremisesSamAccountName;
            OnpremisesUserPrincipalName = onpremisesUserPrincipalName;
            PhysicalDeliveryOfficeName = physicalDeliveryOfficeName;
            PostalCode = postalCode;
            State = state;
            StreetAddress = streetAddress;
            Surname = surname;
            UsageLocation = usageLocation;
            UserPrincipalName = userPrincipalName;
        }
    }
}

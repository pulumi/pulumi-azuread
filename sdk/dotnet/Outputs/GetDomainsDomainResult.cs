// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Outputs
{

    [OutputType]
    public sealed class GetDomainsDomainResult
    {
        /// <summary>
        /// Set to `true` to only return domains whose DNS is managed by Microsoft 365. Defaults to `false`.
        /// </summary>
        public readonly bool AdminManaged;
        /// <summary>
        /// The authentication type of the domain. Possible values include `Managed` or `Federated`.
        /// </summary>
        public readonly string AuthenticationType;
        /// <summary>
        /// Whether this is the default domain that is used for user creation.
        /// </summary>
        public readonly bool Default;
        /// <summary>
        /// The name of the domain.
        /// </summary>
        public readonly string DomainName;
        /// <summary>
        /// Whether this is the initial domain created by Azure Active Directory.
        /// </summary>
        public readonly bool Initial;
        /// <summary>
        /// Whether the domain is a verified root domain (not a subdomain).
        /// </summary>
        public readonly bool Root;
        /// <summary>
        /// A list of capabilities / services supported by the domain. Possible values include `Email`, `Sharepoint`, `EmailInternalRelayOnly`, `OfficeCommunicationsOnline`, `SharePointDefaultDomain`, `FullRedelegation`, `SharePointPublic`, `OrgIdAuthentication`, `Yammer` and `Intune`.
        /// </summary>
        public readonly ImmutableArray<string> SupportedServices;
        /// <summary>
        /// Whether the domain has completed domain ownership verification.
        /// </summary>
        public readonly bool Verified;

        [OutputConstructor]
        private GetDomainsDomainResult(
            bool adminManaged,

            string authenticationType,

            bool @default,

            string domainName,

            bool initial,

            bool root,

            ImmutableArray<string> supportedServices,

            bool verified)
        {
            AdminManaged = adminManaged;
            AuthenticationType = authenticationType;
            Default = @default;
            DomainName = domainName;
            Initial = initial;
            Root = root;
            SupportedServices = supportedServices;
            Verified = verified;
        }
    }
}

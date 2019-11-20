// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to access information about an existing Domains within Azure Active Directory.
        /// 
        /// &gt; **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `Directory.Read.All` within the `Windows Azure Active Directory` API.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azuread/blob/master/website/docs/d/domains.html.markdown.
        /// </summary>
        public static Task<GetDomainsResult> GetDomains(GetDomainsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDomainsResult>("azuread:index/getDomains:getDomains", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetDomainsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set to `true` if unverified Azure AD Domains should be included. Defaults to `false`.
        /// </summary>
        [Input("includeUnverified")]
        public Input<bool>? IncludeUnverified { get; set; }

        /// <summary>
        /// Set to `true` to only return the default domain.
        /// </summary>
        [Input("onlyDefault")]
        public Input<bool>? OnlyDefault { get; set; }

        /// <summary>
        /// Set to `true` to only return the initial domain, which is your primary Azure Active Directory tenant domain. Defaults to `false`.
        /// </summary>
        [Input("onlyInitial")]
        public Input<bool>? OnlyInitial { get; set; }

        public GetDomainsArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetDomainsResult
    {
        /// <summary>
        /// One or more `domain` blocks as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDomainsDomainsResult> Domains;
        public readonly bool? IncludeUnverified;
        public readonly bool? OnlyDefault;
        public readonly bool? OnlyInitial;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetDomainsResult(
            ImmutableArray<Outputs.GetDomainsDomainsResult> domains,
            bool? includeUnverified,
            bool? onlyDefault,
            bool? onlyInitial,
            string id)
        {
            Domains = domains;
            IncludeUnverified = includeUnverified;
            OnlyDefault = onlyDefault;
            OnlyInitial = onlyInitial;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetDomainsDomainsResult
    {
        /// <summary>
        /// The authentication type of the domain (Managed or Federated).
        /// </summary>
        public readonly string AuthenticationType;
        /// <summary>
        /// The name of the domain.
        /// </summary>
        public readonly string DomainName;
        /// <summary>
        /// `True` if this is the default domain that is used for user creation.
        /// </summary>
        public readonly bool IsDefault;
        /// <summary>
        /// `True` if this is the initial domain created by Azure Activie Directory.
        /// </summary>
        public readonly bool IsInitial;
        /// <summary>
        /// `True` if the domain has completed domain ownership verification.
        /// </summary>
        public readonly bool IsVerified;

        [OutputConstructor]
        private GetDomainsDomainsResult(
            string authenticationType,
            string domainName,
            bool isDefault,
            bool isInitial,
            bool isVerified)
        {
            AuthenticationType = authenticationType;
            DomainName = domainName;
            IsDefault = isDefault;
            IsInitial = isInitial;
            IsVerified = isVerified;
        }
    }
    }
}

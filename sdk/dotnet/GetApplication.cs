// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD
{
    public static class GetApplication
    {
        /// <summary>
        /// Use this data source to access information about an existing Application within Azure Active Directory.
        /// 
        /// &gt; **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write all (or owned by) applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.
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
        ///         var example = Output.Create(AzureAD.GetApplication.InvokeAsync(new AzureAD.GetApplicationArgs
        ///         {
        ///             DisplayName = "My First AzureAD Application",
        ///         }));
        ///         this.AzureAdObjectId = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("azureAdObjectId")]
        ///     public Output&lt;string&gt; AzureAdObjectId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetApplicationResult> InvokeAsync(GetApplicationArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetApplicationResult>("azuread:index/getApplication:getApplication", args ?? new GetApplicationArgs(), options.WithVersion());
    }


    public sealed class GetApplicationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the Application ID of the Azure Active Directory Application.
        /// </summary>
        [Input("applicationId")]
        public string? ApplicationId { get; set; }

        /// <summary>
        /// Specifies the display name of the Application within Azure Active Directory.
        /// </summary>
        [Input("displayName")]
        public string? DisplayName { get; set; }

        /// <summary>
        /// The name of the optional claim.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        [Input("oauth2Permissions")]
        private List<Inputs.GetApplicationOauth2PermissionArgs>? _oauth2Permissions;

        /// <summary>
        /// A collection of OAuth 2.0 permission scopes that the web API (resource) app exposes to client apps. Each permission is covered by a `oauth2_permission` block as documented below.
        /// </summary>
        public List<Inputs.GetApplicationOauth2PermissionArgs> Oauth2Permissions
        {
            get => _oauth2Permissions ?? (_oauth2Permissions = new List<Inputs.GetApplicationOauth2PermissionArgs>());
            set => _oauth2Permissions = value;
        }

        /// <summary>
        /// Specifies the Object ID of the Application within Azure Active Directory.
        /// </summary>
        [Input("objectId")]
        public string? ObjectId { get; set; }

        /// <summary>
        /// A collection of `access_token` or `id_token` blocks as documented below which list the optional claims configured for each token type. For more information see https://docs.microsoft.com/en-us/azure/active-directory/develop/active-directory-optional-claims
        /// </summary>
        [Input("optionalClaims")]
        public Inputs.GetApplicationOptionalClaimsArgs? OptionalClaims { get; set; }

        public GetApplicationArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetApplicationResult
    {
        /// <summary>
        /// A collection of `app_role` blocks as documented below. For more information https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles
        /// </summary>
        public readonly ImmutableArray<Outputs.GetApplicationAppRoleResult> AppRoles;
        /// <summary>
        /// the Application ID of the Azure Active Directory Application.
        /// </summary>
        public readonly string ApplicationId;
        /// <summary>
        /// Is this Azure AD Application available to other tenants?
        /// </summary>
        public readonly bool AvailableToOtherTenants;
        /// <summary>
        /// Display name for the permission that appears in the admin consent and app assignment experiences.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The `groups` claim issued in a user or OAuth 2.0 access token that the app expects.
        /// </summary>
        public readonly string GroupMembershipClaims;
        public readonly string Homepage;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
        /// </summary>
        public readonly ImmutableArray<string> IdentifierUris;
        /// <summary>
        /// The URL of the logout page.
        /// </summary>
        public readonly string LogoutUrl;
        /// <summary>
        /// The name of the optional claim.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Does this Azure AD Application allow OAuth2.0 implicit flow tokens?
        /// </summary>
        public readonly bool Oauth2AllowImplicitFlow;
        /// <summary>
        /// A collection of OAuth 2.0 permission scopes that the web API (resource) app exposes to client apps. Each permission is covered by a `oauth2_permission` block as documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetApplicationOauth2PermissionResult> Oauth2Permissions;
        /// <summary>
        /// the Object ID of the Azure Active Directory Application.
        /// </summary>
        public readonly string ObjectId;
        /// <summary>
        /// A collection of `access_token` or `id_token` blocks as documented below which list the optional claims configured for each token type. For more information see https://docs.microsoft.com/en-us/azure/active-directory/develop/active-directory-optional-claims
        /// </summary>
        public readonly Outputs.GetApplicationOptionalClaimsResult? OptionalClaims;
        /// <summary>
        /// A list of User Object IDs that are assigned ownership of the application registration.
        /// </summary>
        public readonly ImmutableArray<string> Owners;
        /// <summary>
        /// A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.
        /// </summary>
        public readonly ImmutableArray<string> ReplyUrls;
        /// <summary>
        /// A collection of `required_resource_access` blocks as documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetApplicationRequiredResourceAccessResult> RequiredResourceAccesses;
        /// <summary>
        /// Specifies whether the id property references an `OAuth2Permission` or an `AppRole`.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetApplicationResult(
            ImmutableArray<Outputs.GetApplicationAppRoleResult> appRoles,

            string applicationId,

            bool availableToOtherTenants,

            string displayName,

            string groupMembershipClaims,

            string homepage,

            string id,

            ImmutableArray<string> identifierUris,

            string logoutUrl,

            string name,

            bool oauth2AllowImplicitFlow,

            ImmutableArray<Outputs.GetApplicationOauth2PermissionResult> oauth2Permissions,

            string objectId,

            Outputs.GetApplicationOptionalClaimsResult? optionalClaims,

            ImmutableArray<string> owners,

            ImmutableArray<string> replyUrls,

            ImmutableArray<Outputs.GetApplicationRequiredResourceAccessResult> requiredResourceAccesses,

            string type)
        {
            AppRoles = appRoles;
            ApplicationId = applicationId;
            AvailableToOtherTenants = availableToOtherTenants;
            DisplayName = displayName;
            GroupMembershipClaims = groupMembershipClaims;
            Homepage = homepage;
            Id = id;
            IdentifierUris = identifierUris;
            LogoutUrl = logoutUrl;
            Name = name;
            Oauth2AllowImplicitFlow = oauth2AllowImplicitFlow;
            Oauth2Permissions = oauth2Permissions;
            ObjectId = objectId;
            OptionalClaims = optionalClaims;
            Owners = owners;
            ReplyUrls = replyUrls;
            RequiredResourceAccesses = requiredResourceAccesses;
            Type = type;
        }
    }
}

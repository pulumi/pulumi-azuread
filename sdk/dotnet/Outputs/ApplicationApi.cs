// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Outputs
{

    [OutputType]
    public sealed class ApplicationApi
    {
        /// <summary>
        /// A set of application IDs (client IDs), used for bundling consent if you have a solution that contains two parts: a client app and a custom web API app.
        /// </summary>
        public readonly ImmutableArray<string> KnownClientApplications;
        /// <summary>
        /// Allows an application to use claims mapping without specifying a custom signing key. Defaults to `false`.
        /// </summary>
        public readonly bool? MappedClaimsEnabled;
        /// <summary>
        /// One or more `oauth2_permission_scope` blocks as documented below, to describe delegated permissions exposed by the web API represented by this application.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationApiOauth2PermissionScope> Oauth2PermissionScopes;
        /// <summary>
        /// The access token version expected by this resource. Must be one of `1` or `2`, and must be `2` when `sign_in_audience` is either `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount` Defaults to `1`.
        /// </summary>
        public readonly int? RequestedAccessTokenVersion;

        [OutputConstructor]
        private ApplicationApi(
            ImmutableArray<string> knownClientApplications,

            bool? mappedClaimsEnabled,

            ImmutableArray<Outputs.ApplicationApiOauth2PermissionScope> oauth2PermissionScopes,

            int? requestedAccessTokenVersion)
        {
            KnownClientApplications = knownClientApplications;
            MappedClaimsEnabled = mappedClaimsEnabled;
            Oauth2PermissionScopes = oauth2PermissionScopes;
            RequestedAccessTokenVersion = requestedAccessTokenVersion;
        }
    }
}

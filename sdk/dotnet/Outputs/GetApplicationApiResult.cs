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
    public sealed class GetApplicationApiResult
    {
        public readonly ImmutableArray<Outputs.GetApplicationApiOauth2PermissionScopeResult> Oauth2PermissionScopes;

        [OutputConstructor]
        private GetApplicationApiResult(ImmutableArray<Outputs.GetApplicationApiOauth2PermissionScopeResult> oauth2PermissionScopes)
        {
            Oauth2PermissionScopes = oauth2PermissionScopes;
        }
    }
}

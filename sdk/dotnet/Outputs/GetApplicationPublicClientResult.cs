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
    public sealed class GetApplicationPublicClientResult
    {
        /// <summary>
        /// A list of URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent.
        /// </summary>
        public readonly ImmutableArray<string> RedirectUris;

        [OutputConstructor]
        private GetApplicationPublicClientResult(ImmutableArray<string> redirectUris)
        {
            RedirectUris = redirectUris;
        }
    }
}

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
    public sealed class GetApplicationSinglePageApplicationResult
    {
        public readonly ImmutableArray<string> RedirectUris;

        [OutputConstructor]
        private GetApplicationSinglePageApplicationResult(ImmutableArray<string> redirectUris)
        {
            RedirectUris = redirectUris;
        }
    }
}

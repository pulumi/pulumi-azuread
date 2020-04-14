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
    public sealed class ApplicationRequiredResourceAccess
    {
        /// <summary>
        /// A collection of `resource_access` blocks as documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationRequiredResourceAccessResourceAccess> ResourceAccesses;
        /// <summary>
        /// The unique identifier for the resource that the application requires access to. This should be equal to the appId declared on the target resource application.
        /// </summary>
        public readonly string ResourceAppId;

        [OutputConstructor]
        private ApplicationRequiredResourceAccess(
            ImmutableArray<Outputs.ApplicationRequiredResourceAccessResourceAccess> resourceAccesses,

            string resourceAppId)
        {
            ResourceAccesses = resourceAccesses;
            ResourceAppId = resourceAppId;
        }
    }
}

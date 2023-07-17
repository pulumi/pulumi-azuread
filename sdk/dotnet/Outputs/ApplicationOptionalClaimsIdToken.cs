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
    public sealed class ApplicationOptionalClaimsIdToken
    {
        /// <summary>
        /// List of additional properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim.
        /// </summary>
        public readonly ImmutableArray<string> AdditionalProperties;
        /// <summary>
        /// Whether the claim specified by the client is necessary to ensure a smooth authorization experience.
        /// </summary>
        public readonly bool? Essential;
        /// <summary>
        /// The name of the optional claim.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The source of the claim. If `source` is absent, the claim is a predefined optional claim. If `source` is `user`, the value of `name` is the extension property from the user object.
        /// </summary>
        public readonly string? Source;

        [OutputConstructor]
        private ApplicationOptionalClaimsIdToken(
            ImmutableArray<string> additionalProperties,

            bool? essential,

            string name,

            string? source)
        {
            AdditionalProperties = additionalProperties;
            Essential = essential;
            Name = name;
            Source = source;
        }
    }
}

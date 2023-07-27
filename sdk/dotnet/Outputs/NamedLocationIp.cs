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
    public sealed class NamedLocationIp
    {
        /// <summary>
        /// List of IP address ranges in IPv4 CIDR format (e.g. `1.2.3.4/32`) or any allowable IPv6 format from IETF RFC596.
        /// </summary>
        public readonly ImmutableArray<string> IpRanges;
        /// <summary>
        /// Whether the named location is trusted. Defaults to `false`.
        /// </summary>
        public readonly bool? Trusted;

        [OutputConstructor]
        private NamedLocationIp(
            ImmutableArray<string> ipRanges,

            bool? trusted)
        {
            IpRanges = ipRanges;
            Trusted = trusted;
        }
    }
}

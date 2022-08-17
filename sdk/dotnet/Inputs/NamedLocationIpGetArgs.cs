// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class NamedLocationIpGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("ipRanges", required: true)]
        private InputList<string>? _ipRanges;

        /// <summary>
        /// List of IP address ranges in IPv4 CIDR format (e.g. 1.2.3.4/32) or any allowable IPv6 format from IETF RFC596.
        /// </summary>
        public InputList<string> IpRanges
        {
            get => _ipRanges ?? (_ipRanges = new InputList<string>());
            set => _ipRanges = value;
        }

        /// <summary>
        /// Whether the named location is trusted. Defaults to `false`.
        /// </summary>
        [Input("trusted")]
        public Input<bool>? Trusted { get; set; }

        public NamedLocationIpGetArgs()
        {
        }
        public static new NamedLocationIpGetArgs Empty => new NamedLocationIpGetArgs();
    }
}

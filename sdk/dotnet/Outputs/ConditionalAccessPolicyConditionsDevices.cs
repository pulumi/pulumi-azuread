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
    public sealed class ConditionalAccessPolicyConditionsDevices
    {
        /// <summary>
        /// A `filter` block as described below.
        /// </summary>
        public readonly Outputs.ConditionalAccessPolicyConditionsDevicesFilter? Filter;

        [OutputConstructor]
        private ConditionalAccessPolicyConditionsDevices(Outputs.ConditionalAccessPolicyConditionsDevicesFilter? filter)
        {
            Filter = filter;
        }
    }
}

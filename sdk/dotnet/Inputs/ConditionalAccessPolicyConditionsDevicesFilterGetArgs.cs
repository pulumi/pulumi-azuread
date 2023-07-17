// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class ConditionalAccessPolicyConditionsDevicesFilterGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to include in, or exclude from, matching devices from the policy. Supported values are `include` or `exclude`.
        /// </summary>
        [Input("mode", required: true)]
        public Input<string> Mode { get; set; } = null!;

        /// <summary>
        /// Condition filter to match devices. For more information, see [official documentation](https://docs.microsoft.com/en-us/azure/active-directory/conditional-access/concept-condition-filters-for-devices#supported-operators-and-device-properties-for-filters).
        /// </summary>
        [Input("rule", required: true)]
        public Input<string> Rule { get; set; } = null!;

        public ConditionalAccessPolicyConditionsDevicesFilterGetArgs()
        {
        }
        public static new ConditionalAccessPolicyConditionsDevicesFilterGetArgs Empty => new ConditionalAccessPolicyConditionsDevicesFilterGetArgs();
    }
}

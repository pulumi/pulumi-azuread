// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class SynchronizationJobScheduleGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("expiration")]
        public Input<string>? Expiration { get; set; }

        [Input("interval")]
        public Input<string>? Interval { get; set; }

        [Input("state")]
        public Input<string>? State { get; set; }

        public SynchronizationJobScheduleGetArgs()
        {
        }
        public static new SynchronizationJobScheduleGetArgs Empty => new SynchronizationJobScheduleGetArgs();
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class GroupDynamicMembershipGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        [Input("rule", required: true)]
        public Input<string> Rule { get; set; } = null!;

        public GroupDynamicMembershipGetArgs()
        {
        }
        public static new GroupDynamicMembershipGetArgs Empty => new GroupDynamicMembershipGetArgs();
    }
}

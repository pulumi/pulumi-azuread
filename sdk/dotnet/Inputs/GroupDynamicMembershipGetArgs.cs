// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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
        /// <summary>
        /// Whether rule processing is "On" (true) or "Paused" (false).
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// The rule that determines membership of this group. For more information, see official documentation on [membership rules syntax](https://docs.microsoft.com/en-gb/azure/active-directory/enterprise-users/groups-dynamic-membership).
        /// 
        /// &gt; **Dynamic Group Memberships** Remember to include `DynamicMembership` in the set of `types` for the group when configuring a dynamic membership rule. Dynamic membership is a premium feature which requires an Azure Active Directory P1 or P2 license.
        /// </summary>
        [Input("rule", required: true)]
        public Input<string> Rule { get; set; } = null!;

        public GroupDynamicMembershipGetArgs()
        {
        }
        public static new GroupDynamicMembershipGetArgs Empty => new GroupDynamicMembershipGetArgs();
    }
}

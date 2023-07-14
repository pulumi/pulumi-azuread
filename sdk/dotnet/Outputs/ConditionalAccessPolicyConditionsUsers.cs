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
    public sealed class ConditionalAccessPolicyConditionsUsers
    {
        public readonly ImmutableArray<string> ExcludedGroups;
        public readonly ImmutableArray<string> ExcludedRoles;
        public readonly ImmutableArray<string> ExcludedUsers;
        public readonly ImmutableArray<string> IncludedGroups;
        public readonly ImmutableArray<string> IncludedRoles;
        public readonly ImmutableArray<string> IncludedUsers;

        [OutputConstructor]
        private ConditionalAccessPolicyConditionsUsers(
            ImmutableArray<string> excludedGroups,

            ImmutableArray<string> excludedRoles,

            ImmutableArray<string> excludedUsers,

            ImmutableArray<string> includedGroups,

            ImmutableArray<string> includedRoles,

            ImmutableArray<string> includedUsers)
        {
            ExcludedGroups = excludedGroups;
            ExcludedRoles = excludedRoles;
            ExcludedUsers = excludedUsers;
            IncludedGroups = includedGroups;
            IncludedRoles = includedRoles;
            IncludedUsers = includedUsers;
        }
    }
}

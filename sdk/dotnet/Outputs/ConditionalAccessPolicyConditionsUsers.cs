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
        /// <summary>
        /// A list of group IDs excluded from scope of policy.
        /// </summary>
        public readonly ImmutableArray<string> ExcludedGroups;
        /// <summary>
        /// A list of role IDs excluded from scope of policy.
        /// </summary>
        public readonly ImmutableArray<string> ExcludedRoles;
        /// <summary>
        /// A list of user IDs excluded from scope of policy and/or `GuestsOrExternalUsers`.
        /// </summary>
        public readonly ImmutableArray<string> ExcludedUsers;
        /// <summary>
        /// A list of group IDs in scope of policy unless explicitly excluded.
        /// </summary>
        public readonly ImmutableArray<string> IncludedGroups;
        /// <summary>
        /// A list of role IDs in scope of policy unless explicitly excluded.
        /// </summary>
        public readonly ImmutableArray<string> IncludedRoles;
        /// <summary>
        /// A list of user IDs in scope of policy unless explicitly excluded, or `None` or `All` or `GuestsOrExternalUsers`.
        /// </summary>
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

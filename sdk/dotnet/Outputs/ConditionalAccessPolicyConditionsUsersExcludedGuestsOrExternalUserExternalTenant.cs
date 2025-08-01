// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Outputs
{

    [OutputType]
    public sealed class ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUserExternalTenant
    {
        /// <summary>
        /// A list tenant IDs. Can only be specified if `membership_kind` is `enumerated`.
        /// </summary>
        public readonly ImmutableArray<string> Members;
        /// <summary>
        /// The external tenant membership kind. Possible values are: `all`, `enumerated`, `unknownFutureValue`.
        /// </summary>
        public readonly string MembershipKind;

        [OutputConstructor]
        private ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUserExternalTenant(
            ImmutableArray<string> members,

            string membershipKind)
        {
            Members = members;
            MembershipKind = membershipKind;
        }
    }
}

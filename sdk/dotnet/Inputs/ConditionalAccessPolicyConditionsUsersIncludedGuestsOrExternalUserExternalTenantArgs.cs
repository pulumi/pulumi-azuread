// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUserExternalTenantArgs : global::Pulumi.ResourceArgs
    {
        [Input("members")]
        private InputList<string>? _members;

        /// <summary>
        /// A list tenant IDs. Can only be specified if `membership_kind` is `enumerated`.
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// The external tenant membership kind. Possible values are: `all`, `enumerated`, `unknownFutureValue`.
        /// </summary>
        [Input("membershipKind", required: true)]
        public Input<string> MembershipKind { get; set; } = null!;

        public ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUserExternalTenantArgs()
        {
        }
        public static new ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUserExternalTenantArgs Empty => new ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUserExternalTenantArgs();
    }
}

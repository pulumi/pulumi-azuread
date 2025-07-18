// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class AccessPackageAssignmentPolicyRequestorSettingsArgs : global::Pulumi.ResourceArgs
    {
        [Input("requestors")]
        private InputList<Inputs.AccessPackageAssignmentPolicyRequestorSettingsRequestorArgs>? _requestors;

        /// <summary>
        /// A block specifying the users who are allowed to request on this policy, as documented below.
        /// </summary>
        public InputList<Inputs.AccessPackageAssignmentPolicyRequestorSettingsRequestorArgs> Requestors
        {
            get => _requestors ?? (_requestors = new InputList<Inputs.AccessPackageAssignmentPolicyRequestorSettingsRequestorArgs>());
            set => _requestors = value;
        }

        /// <summary>
        /// Whether to accept requests using this policy. When `false`, no new requests can be made using this policy.
        /// </summary>
        [Input("requestsAccepted")]
        public Input<bool>? RequestsAccepted { get; set; }

        /// <summary>
        /// Specifies the scopes of the requestors. Valid values are `AllConfiguredConnectedOrganizationSubjects`, `AllExistingConnectedOrganizationSubjects`, `AllExistingDirectoryMemberUsers`, `AllExistingDirectorySubjects`, `AllExternalSubjects`, `NoSubjects`, `SpecificConnectedOrganizationSubjects`, or `SpecificDirectorySubjects`.
        /// </summary>
        [Input("scopeType")]
        public Input<string>? ScopeType { get; set; }

        public AccessPackageAssignmentPolicyRequestorSettingsArgs()
        {
        }
        public static new AccessPackageAssignmentPolicyRequestorSettingsArgs Empty => new AccessPackageAssignmentPolicyRequestorSettingsArgs();
    }
}

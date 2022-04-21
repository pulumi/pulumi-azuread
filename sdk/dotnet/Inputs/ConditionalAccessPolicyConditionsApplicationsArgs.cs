// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class ConditionalAccessPolicyConditionsApplicationsArgs : Pulumi.ResourceArgs
    {
        [Input("excludedApplications")]
        private InputList<string>? _excludedApplications;

        /// <summary>
        /// A list of application IDs explicitly excluded from the policy.
        /// </summary>
        public InputList<string> ExcludedApplications
        {
            get => _excludedApplications ?? (_excludedApplications = new InputList<string>());
            set => _excludedApplications = value;
        }

        [Input("includedApplications")]
        private InputList<string>? _includedApplications;

        /// <summary>
        /// A list of application IDs the policy applies to, unless explicitly excluded (in `excluded_applications`). Can also be set to `All`. Cannot be specified with `included_user_actions`. One of `included_applications` or `included_user_actions` must be specified.
        /// </summary>
        public InputList<string> IncludedApplications
        {
            get => _includedApplications ?? (_includedApplications = new InputList<string>());
            set => _includedApplications = value;
        }

        [Input("includedUserActions")]
        private InputList<string>? _includedUserActions;

        /// <summary>
        /// A list of user actions to include. Supported values are `urn:user:registerdevice` and `urn:user:registersecurityinfo`. Cannot be specified with `included_applications`. One of `included_applications` or `included_user_actions` must be specified.
        /// </summary>
        public InputList<string> IncludedUserActions
        {
            get => _includedUserActions ?? (_includedUserActions = new InputList<string>());
            set => _includedUserActions = value;
        }

        public ConditionalAccessPolicyConditionsApplicationsArgs()
        {
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class ConditionalAccessPolicyConditionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An `applications` block as documented below, which specifies applications and user actions included in and excluded from the policy.
        /// </summary>
        [Input("applications", required: true)]
        public Input<Inputs.ConditionalAccessPolicyConditionsApplicationsArgs> Applications { get; set; } = null!;

        [Input("clientAppTypes", required: true)]
        private InputList<string>? _clientAppTypes;

        /// <summary>
        /// A list of client application types included in the policy. Possible values are: `all`, `browser`, `mobileAppsAndDesktopClients`, `exchangeActiveSync`, `easSupported` and `other`.
        /// </summary>
        public InputList<string> ClientAppTypes
        {
            get => _clientAppTypes ?? (_clientAppTypes = new InputList<string>());
            set => _clientAppTypes = value;
        }

        /// <summary>
        /// A `devices` block as documented below, which describes devices to be included in and excluded from the policy. A `devices` block can be added to an existing policy, but removing the `devices` block forces a new resource to be created.
        /// </summary>
        [Input("devices")]
        public Input<Inputs.ConditionalAccessPolicyConditionsDevicesArgs>? Devices { get; set; }

        /// <summary>
        /// A `locations` block as documented below, which specifies locations included in and excluded from the policy.
        /// </summary>
        [Input("locations")]
        public Input<Inputs.ConditionalAccessPolicyConditionsLocationsArgs>? Locations { get; set; }

        /// <summary>
        /// A `platforms` block as documented below, which specifies platforms included in and excluded from the policy.
        /// </summary>
        [Input("platforms")]
        public Input<Inputs.ConditionalAccessPolicyConditionsPlatformsArgs>? Platforms { get; set; }

        [Input("signInRiskLevels")]
        private InputList<string>? _signInRiskLevels;

        /// <summary>
        /// A list of sign-in risk levels included in the policy. Possible values are: `low`, `medium`, `high`, `hidden`, `none`, `unknownFutureValue`.
        /// </summary>
        public InputList<string> SignInRiskLevels
        {
            get => _signInRiskLevels ?? (_signInRiskLevels = new InputList<string>());
            set => _signInRiskLevels = value;
        }

        [Input("userRiskLevels")]
        private InputList<string>? _userRiskLevels;

        /// <summary>
        /// A list of user risk levels included in the policy. Possible values are: `low`, `medium`, `high`, `hidden`, `none`, `unknownFutureValue`.
        /// </summary>
        public InputList<string> UserRiskLevels
        {
            get => _userRiskLevels ?? (_userRiskLevels = new InputList<string>());
            set => _userRiskLevels = value;
        }

        /// <summary>
        /// A `users` block as documented below, which specifies users, groups, and roles included in and excluded from the policy.
        /// </summary>
        [Input("users", required: true)]
        public Input<Inputs.ConditionalAccessPolicyConditionsUsersArgs> Users { get; set; } = null!;

        public ConditionalAccessPolicyConditionsArgs()
        {
        }
        public static new ConditionalAccessPolicyConditionsArgs Empty => new ConditionalAccessPolicyConditionsArgs();
    }
}

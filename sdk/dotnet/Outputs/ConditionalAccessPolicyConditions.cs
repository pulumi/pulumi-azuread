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
    public sealed class ConditionalAccessPolicyConditions
    {
        /// <summary>
        /// An `applications` block as documented below, which specifies applications and user actions included in and excluded from the policy.
        /// </summary>
        public readonly Outputs.ConditionalAccessPolicyConditionsApplications Applications;
        /// <summary>
        /// A list of client application types included in the policy. Possible values are: `all`, `browser`, `mobileAppsAndDesktopClients`, `exchangeActiveSync`, `easSupported` and `other`.
        /// </summary>
        public readonly ImmutableArray<string> ClientAppTypes;
        /// <summary>
        /// An `client_applications` block as documented below, which specifies service principals included in and excluded from the policy.
        /// </summary>
        public readonly Outputs.ConditionalAccessPolicyConditionsClientApplications? ClientApplications;
        /// <summary>
        /// A `devices` block as documented below, which describes devices to be included in and excluded from the policy. A `devices` block can be added to an existing policy, but removing the `devices` block forces a new resource to be created.
        /// </summary>
        public readonly Outputs.ConditionalAccessPolicyConditionsDevices? Devices;
        /// <summary>
        /// The insider risk level in the policy. Possible values are: `minor`, `moderate`, `elevated`, `unknownFutureValue`.
        /// </summary>
        public readonly string? InsiderRiskLevels;
        /// <summary>
        /// A `locations` block as documented below, which specifies locations included in and excluded from the policy.
        /// </summary>
        public readonly Outputs.ConditionalAccessPolicyConditionsLocations? Locations;
        /// <summary>
        /// A `platforms` block as documented below, which specifies platforms included in and excluded from the policy.
        /// </summary>
        public readonly Outputs.ConditionalAccessPolicyConditionsPlatforms? Platforms;
        /// <summary>
        /// A list of service principal sign-in risk levels included in the policy. Possible values are: `low`, `medium`, `high`, `none`, `unknownFutureValue`.
        /// </summary>
        public readonly ImmutableArray<string> ServicePrincipalRiskLevels;
        /// <summary>
        /// A list of user sign-in risk levels included in the policy. Possible values are: `low`, `medium`, `high`, `hidden`, `none`, `unknownFutureValue`.
        /// </summary>
        public readonly ImmutableArray<string> SignInRiskLevels;
        /// <summary>
        /// A list of user risk levels included in the policy. Possible values are: `low`, `medium`, `high`, `hidden`, `none`, `unknownFutureValue`.
        /// </summary>
        public readonly ImmutableArray<string> UserRiskLevels;
        /// <summary>
        /// A `users` block as documented below, which specifies users, groups, and roles included in and excluded from the policy.
        /// </summary>
        public readonly Outputs.ConditionalAccessPolicyConditionsUsers Users;

        [OutputConstructor]
        private ConditionalAccessPolicyConditions(
            Outputs.ConditionalAccessPolicyConditionsApplications applications,

            ImmutableArray<string> clientAppTypes,

            Outputs.ConditionalAccessPolicyConditionsClientApplications? clientApplications,

            Outputs.ConditionalAccessPolicyConditionsDevices? devices,

            string? insiderRiskLevels,

            Outputs.ConditionalAccessPolicyConditionsLocations? locations,

            Outputs.ConditionalAccessPolicyConditionsPlatforms? platforms,

            ImmutableArray<string> servicePrincipalRiskLevels,

            ImmutableArray<string> signInRiskLevels,

            ImmutableArray<string> userRiskLevels,

            Outputs.ConditionalAccessPolicyConditionsUsers users)
        {
            Applications = applications;
            ClientAppTypes = clientAppTypes;
            ClientApplications = clientApplications;
            Devices = devices;
            InsiderRiskLevels = insiderRiskLevels;
            Locations = locations;
            Platforms = platforms;
            ServicePrincipalRiskLevels = servicePrincipalRiskLevels;
            SignInRiskLevels = signInRiskLevels;
            UserRiskLevels = userRiskLevels;
            Users = users;
        }
    }
}

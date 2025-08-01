// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class ConditionalAccessPolicySessionControlsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether application enforced restrictions are enabled. Defaults to `false`.
        /// 
        /// &gt; Only Office 365, Exchange Online and Sharepoint Online support application enforced restrictions.
        /// </summary>
        [Input("applicationEnforcedRestrictionsEnabled")]
        public Input<bool>? ApplicationEnforcedRestrictionsEnabled { get; set; }

        /// <summary>
        /// Enables cloud app security and specifies the cloud app security policy to use. Possible values are: `blockDownloads`, `mcasConfigured`, `monitorOnly` or `unknownFutureValue`.
        /// </summary>
        [Input("cloudAppSecurityPolicy")]
        public Input<string>? CloudAppSecurityPolicy { get; set; }

        /// <summary>
        /// Disables [resilience defaults](https://learn.microsoft.com/en-us/azure/active-directory/conditional-access/resilience-defaults). Defaults to `false`.
        /// </summary>
        [Input("disableResilienceDefaults")]
        public Input<bool>? DisableResilienceDefaults { get; set; }

        /// <summary>
        /// Session control to define whether to persist cookies. Possible values are: `always` or `never`.
        /// </summary>
        [Input("persistentBrowserMode")]
        public Input<string>? PersistentBrowserMode { get; set; }

        /// <summary>
        /// Number of days or hours to enforce sign-in frequency. Required when `sign_in_frequency_period` is specified.
        /// </summary>
        [Input("signInFrequency")]
        public Input<int>? SignInFrequency { get; set; }

        /// <summary>
        /// Authentication type for enforcing sign-in frequency. Possible values are: `primaryAndSecondaryAuthentication` or `secondaryAuthentication`. Defaults to `primaryAndSecondaryAuthentication`.
        /// </summary>
        [Input("signInFrequencyAuthenticationType")]
        public Input<string>? SignInFrequencyAuthenticationType { get; set; }

        /// <summary>
        /// The interval to apply to sign-in frequency control. Possible values are: `timeBased` or `everyTime`. Defaults to `timeBased`.
        /// </summary>
        [Input("signInFrequencyInterval")]
        public Input<string>? SignInFrequencyInterval { get; set; }

        /// <summary>
        /// The time period to enforce sign-in frequency. Possible values are: `hours` or `days`. Required when `sign_in_frequency_period` is specified.
        /// </summary>
        [Input("signInFrequencyPeriod")]
        public Input<string>? SignInFrequencyPeriod { get; set; }

        public ConditionalAccessPolicySessionControlsGetArgs()
        {
        }
        public static new ConditionalAccessPolicySessionControlsGetArgs Empty => new ConditionalAccessPolicySessionControlsGetArgs();
    }
}

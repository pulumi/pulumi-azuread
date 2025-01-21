// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class ConditionalAccessPolicyGrantControlsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of an Authentication Strength Policy to use in this policy. When using a hard-coded ID, the UUID value should be prefixed with: `/policies/authenticationStrengthPolicies/`.
        /// </summary>
        [Input("authenticationStrengthPolicyId")]
        public Input<string>? AuthenticationStrengthPolicyId { get; set; }

        [Input("builtInControls")]
        private InputList<string>? _builtInControls;

        /// <summary>
        /// List of built-in controls required by the policy. Possible values are: `block`, `mfa`, `approvedApplication`, `compliantApplication`, `compliantDevice`, `domainJoinedDevice`, `passwordChange` or `unknownFutureValue`.
        /// </summary>
        public InputList<string> BuiltInControls
        {
            get => _builtInControls ?? (_builtInControls = new InputList<string>());
            set => _builtInControls = value;
        }

        [Input("customAuthenticationFactors")]
        private InputList<string>? _customAuthenticationFactors;

        /// <summary>
        /// List of custom controls IDs required by the policy.
        /// </summary>
        public InputList<string> CustomAuthenticationFactors
        {
            get => _customAuthenticationFactors ?? (_customAuthenticationFactors = new InputList<string>());
            set => _customAuthenticationFactors = value;
        }

        /// <summary>
        /// Defines the relationship of the grant controls. Possible values are: `AND`, `OR`.
        /// </summary>
        [Input("operator", required: true)]
        public Input<string> Operator { get; set; } = null!;

        [Input("termsOfUses")]
        private InputList<string>? _termsOfUses;

        /// <summary>
        /// List of terms of use IDs required by the policy.
        /// 
        /// &gt; At least one of `authentication_strength_policy_id`, `built_in_controls` or `terms_of_use` must be specified.
        /// </summary>
        public InputList<string> TermsOfUses
        {
            get => _termsOfUses ?? (_termsOfUses = new InputList<string>());
            set => _termsOfUses = value;
        }

        public ConditionalAccessPolicyGrantControlsGetArgs()
        {
        }
        public static new ConditionalAccessPolicyGrantControlsGetArgs Empty => new ConditionalAccessPolicyGrantControlsGetArgs();
    }
}

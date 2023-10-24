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
    public sealed class ConditionalAccessPolicyGrantControls
    {
        /// <summary>
        /// ID of an Authentication Strength Policy to use in this policy.
        /// </summary>
        public readonly string? AuthenticationStrengthPolicyId;
        /// <summary>
        /// List of built-in controls required by the policy. Possible values are: `block`, `mfa`, `approvedApplication`, `compliantApplication`, `compliantDevice`, `domainJoinedDevice`, `passwordChange` or `unknownFutureValue`.
        /// </summary>
        public readonly ImmutableArray<string> BuiltInControls;
        /// <summary>
        /// List of custom controls IDs required by the policy.
        /// </summary>
        public readonly ImmutableArray<string> CustomAuthenticationFactors;
        /// <summary>
        /// Defines the relationship of the grant controls. Possible values are: `AND`, `OR`.
        /// </summary>
        public readonly string Operator;
        /// <summary>
        /// List of terms of use IDs required by the policy.
        /// 
        /// &gt; At least one of `authentication_strength_policy_id`, `built_in_controls` or `terms_of_use` must be specified.
        /// </summary>
        public readonly ImmutableArray<string> TermsOfUses;

        [OutputConstructor]
        private ConditionalAccessPolicyGrantControls(
            string? authenticationStrengthPolicyId,

            ImmutableArray<string> builtInControls,

            ImmutableArray<string> customAuthenticationFactors,

            string @operator,

            ImmutableArray<string> termsOfUses)
        {
            AuthenticationStrengthPolicyId = authenticationStrengthPolicyId;
            BuiltInControls = builtInControls;
            CustomAuthenticationFactors = customAuthenticationFactors;
            Operator = @operator;
            TermsOfUses = termsOfUses;
        }
    }
}

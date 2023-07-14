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
        public readonly ImmutableArray<string> BuiltInControls;
        public readonly ImmutableArray<string> CustomAuthenticationFactors;
        public readonly string Operator;
        public readonly ImmutableArray<string> TermsOfUses;

        [OutputConstructor]
        private ConditionalAccessPolicyGrantControls(
            ImmutableArray<string> builtInControls,

            ImmutableArray<string> customAuthenticationFactors,

            string @operator,

            ImmutableArray<string> termsOfUses)
        {
            BuiltInControls = builtInControls;
            CustomAuthenticationFactors = customAuthenticationFactors;
            Operator = @operator;
            TermsOfUses = termsOfUses;
        }
    }
}

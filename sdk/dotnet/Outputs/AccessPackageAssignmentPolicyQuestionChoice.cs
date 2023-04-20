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
    public sealed class AccessPackageAssignmentPolicyQuestionChoice
    {
        /// <summary>
        /// The actual value of this choice.
        /// </summary>
        public readonly string ActualValue;
        /// <summary>
        /// A block describing the display text of this choice, as documented below.
        /// </summary>
        public readonly Outputs.AccessPackageAssignmentPolicyQuestionChoiceDisplayValue DisplayValue;

        [OutputConstructor]
        private AccessPackageAssignmentPolicyQuestionChoice(
            string actualValue,

            Outputs.AccessPackageAssignmentPolicyQuestionChoiceDisplayValue displayValue)
        {
            ActualValue = actualValue;
            DisplayValue = displayValue;
        }
    }
}

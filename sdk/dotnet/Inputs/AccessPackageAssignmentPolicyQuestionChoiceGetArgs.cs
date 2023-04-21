// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class AccessPackageAssignmentPolicyQuestionChoiceGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The actual value of this choice.
        /// </summary>
        [Input("actualValue", required: true)]
        public Input<string> ActualValue { get; set; } = null!;

        /// <summary>
        /// A block describing the display text of this choice, as documented below.
        /// </summary>
        [Input("displayValue", required: true)]
        public Input<Inputs.AccessPackageAssignmentPolicyQuestionChoiceDisplayValueGetArgs> DisplayValue { get; set; } = null!;

        public AccessPackageAssignmentPolicyQuestionChoiceGetArgs()
        {
        }
        public static new AccessPackageAssignmentPolicyQuestionChoiceGetArgs Empty => new AccessPackageAssignmentPolicyQuestionChoiceGetArgs();
    }
}
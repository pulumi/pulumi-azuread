// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class AccessPackageAssignmentPolicyQuestionChoiceArgs : global::Pulumi.ResourceArgs
    {
        [Input("actualValue", required: true)]
        public Input<string> ActualValue { get; set; } = null!;

        [Input("displayValue", required: true)]
        public Input<Inputs.AccessPackageAssignmentPolicyQuestionChoiceDisplayValueArgs> DisplayValue { get; set; } = null!;

        public AccessPackageAssignmentPolicyQuestionChoiceArgs()
        {
        }
        public static new AccessPackageAssignmentPolicyQuestionChoiceArgs Empty => new AccessPackageAssignmentPolicyQuestionChoiceArgs();
    }
}

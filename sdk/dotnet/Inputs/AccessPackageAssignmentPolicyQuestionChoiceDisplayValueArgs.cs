// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class AccessPackageAssignmentPolicyQuestionChoiceDisplayValueArgs : global::Pulumi.ResourceArgs
    {
        [Input("defaultText", required: true)]
        public Input<string> DefaultText { get; set; } = null!;

        [Input("localizedTexts")]
        private InputList<Inputs.AccessPackageAssignmentPolicyQuestionChoiceDisplayValueLocalizedTextArgs>? _localizedTexts;
        public InputList<Inputs.AccessPackageAssignmentPolicyQuestionChoiceDisplayValueLocalizedTextArgs> LocalizedTexts
        {
            get => _localizedTexts ?? (_localizedTexts = new InputList<Inputs.AccessPackageAssignmentPolicyQuestionChoiceDisplayValueLocalizedTextArgs>());
            set => _localizedTexts = value;
        }

        public AccessPackageAssignmentPolicyQuestionChoiceDisplayValueArgs()
        {
        }
        public static new AccessPackageAssignmentPolicyQuestionChoiceDisplayValueArgs Empty => new AccessPackageAssignmentPolicyQuestionChoiceDisplayValueArgs();
    }
}

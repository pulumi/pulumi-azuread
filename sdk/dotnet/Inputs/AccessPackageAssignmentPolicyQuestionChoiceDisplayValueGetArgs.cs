// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class AccessPackageAssignmentPolicyQuestionChoiceDisplayValueGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The default text of this question.`default_text` (Required) The default text of this question choice.
        /// </summary>
        [Input("defaultText", required: true)]
        public Input<string> DefaultText { get; set; } = null!;

        [Input("localizedTexts")]
        private InputList<Inputs.AccessPackageAssignmentPolicyQuestionChoiceDisplayValueLocalizedTextGetArgs>? _localizedTexts;

        /// <summary>
        /// One or more blocks describing localized text of this question, as documented below.
        /// 
        /// 
        /// `localized_text` (Optional) One or more blocks describing localized text of this question choice, as documented below.
        /// </summary>
        public InputList<Inputs.AccessPackageAssignmentPolicyQuestionChoiceDisplayValueLocalizedTextGetArgs> LocalizedTexts
        {
            get => _localizedTexts ?? (_localizedTexts = new InputList<Inputs.AccessPackageAssignmentPolicyQuestionChoiceDisplayValueLocalizedTextGetArgs>());
            set => _localizedTexts = value;
        }

        public AccessPackageAssignmentPolicyQuestionChoiceDisplayValueGetArgs()
        {
        }
        public static new AccessPackageAssignmentPolicyQuestionChoiceDisplayValueGetArgs Empty => new AccessPackageAssignmentPolicyQuestionChoiceDisplayValueGetArgs();
    }
}

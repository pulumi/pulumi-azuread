// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class AccessPackageAssignmentPolicyQuestionTextGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The default text of this question
        /// </summary>
        [Input("defaultText", required: true)]
        public Input<string> DefaultText { get; set; } = null!;

        [Input("localizedTexts")]
        private InputList<Inputs.AccessPackageAssignmentPolicyQuestionTextLocalizedTextGetArgs>? _localizedTexts;

        /// <summary>
        /// The localized text of this question
        /// </summary>
        public InputList<Inputs.AccessPackageAssignmentPolicyQuestionTextLocalizedTextGetArgs> LocalizedTexts
        {
            get => _localizedTexts ?? (_localizedTexts = new InputList<Inputs.AccessPackageAssignmentPolicyQuestionTextLocalizedTextGetArgs>());
            set => _localizedTexts = value;
        }

        public AccessPackageAssignmentPolicyQuestionTextGetArgs()
        {
        }
        public static new AccessPackageAssignmentPolicyQuestionTextGetArgs Empty => new AccessPackageAssignmentPolicyQuestionTextGetArgs();
    }
}

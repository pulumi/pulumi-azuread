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
    public sealed class AccessPackageAssignmentPolicyQuestionChoiceDisplayValueLocalizedText
    {
        /// <summary>
        /// The localized content of this question
        /// </summary>
        public readonly string Content;
        /// <summary>
        /// The language code of this question content
        /// </summary>
        public readonly string LanguageCode;

        [OutputConstructor]
        private AccessPackageAssignmentPolicyQuestionChoiceDisplayValueLocalizedText(
            string content,

            string languageCode)
        {
            Content = content;
            LanguageCode = languageCode;
        }
    }
}

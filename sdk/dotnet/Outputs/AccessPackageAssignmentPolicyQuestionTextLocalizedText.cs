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
    public sealed class AccessPackageAssignmentPolicyQuestionTextLocalizedText
    {
        /// <summary>
        /// The localized content of this question.`content` (Required) The localized content of this question choice.
        /// </summary>
        public readonly string Content;
        /// <summary>
        /// The ISO 639 language code for this question content.
        /// 
        /// 
        /// `language_code` (Required) The ISO 639 language code for this question choice content.
        /// </summary>
        public readonly string LanguageCode;

        [OutputConstructor]
        private AccessPackageAssignmentPolicyQuestionTextLocalizedText(
            string content,

            string languageCode)
        {
            Content = content;
            LanguageCode = languageCode;
        }
    }
}

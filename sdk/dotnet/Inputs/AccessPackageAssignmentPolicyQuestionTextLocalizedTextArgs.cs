// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class AccessPackageAssignmentPolicyQuestionTextLocalizedTextArgs : global::Pulumi.ResourceArgs
    {
        [Input("content", required: true)]
        public Input<string> Content { get; set; } = null!;

        [Input("languageCode", required: true)]
        public Input<string> LanguageCode { get; set; } = null!;

        public AccessPackageAssignmentPolicyQuestionTextLocalizedTextArgs()
        {
        }
        public static new AccessPackageAssignmentPolicyQuestionTextLocalizedTextArgs Empty => new AccessPackageAssignmentPolicyQuestionTextLocalizedTextArgs();
    }
}

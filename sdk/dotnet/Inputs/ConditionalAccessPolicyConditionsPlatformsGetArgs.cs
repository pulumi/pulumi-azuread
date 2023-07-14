// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class ConditionalAccessPolicyConditionsPlatformsGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("excludedPlatforms")]
        private InputList<string>? _excludedPlatforms;
        public InputList<string> ExcludedPlatforms
        {
            get => _excludedPlatforms ?? (_excludedPlatforms = new InputList<string>());
            set => _excludedPlatforms = value;
        }

        [Input("includedPlatforms", required: true)]
        private InputList<string>? _includedPlatforms;
        public InputList<string> IncludedPlatforms
        {
            get => _includedPlatforms ?? (_includedPlatforms = new InputList<string>());
            set => _includedPlatforms = value;
        }

        public ConditionalAccessPolicyConditionsPlatformsGetArgs()
        {
        }
        public static new ConditionalAccessPolicyConditionsPlatformsGetArgs Empty => new ConditionalAccessPolicyConditionsPlatformsGetArgs();
    }
}

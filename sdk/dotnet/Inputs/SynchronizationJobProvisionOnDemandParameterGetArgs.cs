// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class SynchronizationJobProvisionOnDemandParameterGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The identifier of the synchronization rule to be applied. This rule ID is defined in the schema for a given synchronization job or template.
        /// </summary>
        [Input("ruleId", required: true)]
        public Input<string> RuleId { get; set; } = null!;

        [Input("subjects", required: true)]
        private InputList<Inputs.SynchronizationJobProvisionOnDemandParameterSubjectGetArgs>? _subjects;

        /// <summary>
        /// One or more `subject` blocks as documented below.
        /// </summary>
        public InputList<Inputs.SynchronizationJobProvisionOnDemandParameterSubjectGetArgs> Subjects
        {
            get => _subjects ?? (_subjects = new InputList<Inputs.SynchronizationJobProvisionOnDemandParameterSubjectGetArgs>());
            set => _subjects = value;
        }

        public SynchronizationJobProvisionOnDemandParameterGetArgs()
        {
        }
        public static new SynchronizationJobProvisionOnDemandParameterGetArgs Empty => new SynchronizationJobProvisionOnDemandParameterGetArgs();
    }
}

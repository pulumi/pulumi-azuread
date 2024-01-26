// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class ApplicationAppRoleGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowedMemberTypes", required: true)]
        private InputList<string>? _allowedMemberTypes;

        /// <summary>
        /// Specifies whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications (that are accessing this application in a standalone scenario) by setting to `Application`, or to both.
        /// </summary>
        public InputList<string> AllowedMemberTypes
        {
            get => _allowedMemberTypes ?? (_allowedMemberTypes = new InputList<string>());
            set => _allowedMemberTypes = value;
        }

        /// <summary>
        /// Description of the app role that appears when the role is being assigned and, if the role functions as an application permissions, during the consent experiences.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// Display name for the app role that appears during app role assignment and in consent experiences.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// Determines if the app role is enabled. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The unique identifier of the app role. Must be a valid UUID.
        /// 
        /// &gt; **Tip: Generating a UUID for the `id` field** To generate a value for the `id` field in cases where the actual UUID is not important, you can use the `random_uuid` resource. See the application example in the provider repository.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("value")]
        public Input<string>? Value { get; set; }

        public ApplicationAppRoleGetArgs()
        {
        }
        public static new ApplicationAppRoleGetArgs Empty => new ApplicationAppRoleGetArgs();
    }
}

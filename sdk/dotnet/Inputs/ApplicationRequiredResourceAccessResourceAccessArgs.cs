// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class ApplicationRequiredResourceAccessResourceAccessArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique identifier for one of the `OAuth2Permission` or `AppRole` instances that the resource application exposes.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// Specifies whether the `id` property references an `OAuth2Permission` or an `AppRole`. Possible values are `Scope` or `Role`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ApplicationRequiredResourceAccessResourceAccessArgs()
        {
        }
    }
}

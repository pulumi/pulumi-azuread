// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class ApplicationRequiredResourceAccessGetArgs : Pulumi.ResourceArgs
    {
        [Input("resourceAccesses", required: true)]
        private InputList<Inputs.ApplicationRequiredResourceAccessResourceAccessGetArgs>? _resourceAccesses;

        /// <summary>
        /// A collection of `resource_access` blocks as documented below, describing OAuth2.0 permission scopes and app roles that the application requires from the specified resource.
        /// </summary>
        public InputList<Inputs.ApplicationRequiredResourceAccessResourceAccessGetArgs> ResourceAccesses
        {
            get => _resourceAccesses ?? (_resourceAccesses = new InputList<Inputs.ApplicationRequiredResourceAccessResourceAccessGetArgs>());
            set => _resourceAccesses = value;
        }

        /// <summary>
        /// The unique identifier for the resource that the application requires access to. This should be the Application ID of the target application.
        /// </summary>
        [Input("resourceAppId", required: true)]
        public Input<string> ResourceAppId { get; set; } = null!;

        public ApplicationRequiredResourceAccessGetArgs()
        {
        }
    }
}

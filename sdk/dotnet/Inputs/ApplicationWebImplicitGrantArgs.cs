// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class ApplicationWebImplicitGrantArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether this web application can request an access token using OAuth 2.0 implicit flow.
        /// </summary>
        [Input("accessTokenIssuanceEnabled")]
        public Input<bool>? AccessTokenIssuanceEnabled { get; set; }

        public ApplicationWebImplicitGrantArgs()
        {
        }
    }
}

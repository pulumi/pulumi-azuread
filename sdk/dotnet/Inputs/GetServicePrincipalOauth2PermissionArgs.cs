// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class GetServicePrincipalOauth2PermissionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The description of the admin consent
        /// </summary>
        [Input("adminConsentDescription", required: true)]
        public string AdminConsentDescription { get; set; } = null!;

        /// <summary>
        /// The display name of the admin consent
        /// </summary>
        [Input("adminConsentDisplayName", required: true)]
        public string AdminConsentDisplayName { get; set; } = null!;

        /// <summary>
        /// The unique identifier for one of the `OAuth2Permission`
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        /// <summary>
        /// Is this permission enabled?
        /// </summary>
        [Input("isEnabled", required: true)]
        public bool IsEnabled { get; set; }

        /// <summary>
        /// The type of the permission
        /// </summary>
        [Input("type", required: true)]
        public string Type { get; set; } = null!;

        /// <summary>
        /// The description of the user consent
        /// </summary>
        [Input("userConsentDescription", required: true)]
        public string UserConsentDescription { get; set; } = null!;

        /// <summary>
        /// The display name of the user consent
        /// </summary>
        [Input("userConsentDisplayName", required: true)]
        public string UserConsentDisplayName { get; set; } = null!;

        /// <summary>
        /// The name of this permission
        /// </summary>
        [Input("value", required: true)]
        public string Value { get; set; } = null!;

        public GetServicePrincipalOauth2PermissionArgs()
        {
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class ServicePrincipalOauth2PermissionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the admin consent.
        /// </summary>
        [Input("adminConsentDescription")]
        public Input<string>? AdminConsentDescription { get; set; }

        /// <summary>
        /// The display name of the admin consent.
        /// </summary>
        [Input("adminConsentDisplayName")]
        public Input<string>? AdminConsentDisplayName { get; set; }

        /// <summary>
        /// The unique identifier for one of the `OAuth2Permission`.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Is this permission enabled?
        /// </summary>
        [Input("isEnabled")]
        public Input<bool>? IsEnabled { get; set; }

        /// <summary>
        /// The type of the permission.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The description of the user consent.
        /// </summary>
        [Input("userConsentDescription")]
        public Input<string>? UserConsentDescription { get; set; }

        /// <summary>
        /// The display name of the user consent.
        /// </summary>
        [Input("userConsentDisplayName")]
        public Input<string>? UserConsentDisplayName { get; set; }

        /// <summary>
        /// The name of this permission.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public ServicePrincipalOauth2PermissionArgs()
        {
        }
    }
}

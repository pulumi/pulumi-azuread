// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class ServicePrincipalOauth2PermissionScopeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
        /// </summary>
        [Input("adminConsentDescription")]
        public Input<string>? AdminConsentDescription { get; set; }

        /// <summary>
        /// Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
        /// </summary>
        [Input("adminConsentDisplayName")]
        public Input<string>? AdminConsentDisplayName { get; set; }

        /// <summary>
        /// Specifies whether the permission scope is enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The unique identifier of the delegated permission.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Possible values are `User` or `Admin`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
        /// </summary>
        [Input("userConsentDescription")]
        public Input<string>? UserConsentDescription { get; set; }

        /// <summary>
        /// Display name for the delegated permission that appears in the end user consent experience.
        /// </summary>
        [Input("userConsentDisplayName")]
        public Input<string>? UserConsentDisplayName { get; set; }

        /// <summary>
        /// The value that is used for the `scp` claim in OAuth 2.0 access tokens.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public ServicePrincipalOauth2PermissionScopeArgs()
        {
        }
        public static new ServicePrincipalOauth2PermissionScopeArgs Empty => new ServicePrincipalOauth2PermissionScopeArgs();
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class ApplicationApiOauth2PermissionScopeGetArgs : global::Pulumi.ResourceArgs
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
        /// Determines if the permission scope is enabled. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Defaults to `User`. Possible values are `User` or `Admin`.
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

        [Input("value")]
        public Input<string>? Value { get; set; }

        public ApplicationApiOauth2PermissionScopeGetArgs()
        {
        }
        public static new ApplicationApiOauth2PermissionScopeGetArgs Empty => new ApplicationApiOauth2PermissionScopeGetArgs();
    }
}

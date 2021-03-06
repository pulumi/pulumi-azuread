// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Outputs
{

    [OutputType]
    public sealed class ServicePrincipalOauth2PermissionScope
    {
        /// <summary>
        /// The description of the admin consent.
        /// </summary>
        public readonly string? AdminConsentDescription;
        /// <summary>
        /// The display name of the admin consent.
        /// </summary>
        public readonly string? AdminConsentDisplayName;
        /// <summary>
        /// Is this permission enabled?
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// The unique identifier for one of the `OAuth2Permission`.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The type of the permission.
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// The description of the user consent.
        /// </summary>
        public readonly string? UserConsentDescription;
        /// <summary>
        /// The display name of the user consent.
        /// </summary>
        public readonly string? UserConsentDisplayName;
        /// <summary>
        /// The name of this permission.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private ServicePrincipalOauth2PermissionScope(
            string? adminConsentDescription,

            string? adminConsentDisplayName,

            bool? enabled,

            string? id,

            string? type,

            string? userConsentDescription,

            string? userConsentDisplayName,

            string? value)
        {
            AdminConsentDescription = adminConsentDescription;
            AdminConsentDisplayName = adminConsentDisplayName;
            Enabled = enabled;
            Id = id;
            Type = type;
            UserConsentDescription = userConsentDescription;
            UserConsentDisplayName = userConsentDisplayName;
            Value = value;
        }
    }
}

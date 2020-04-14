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
    public sealed class GetServicePrincipalOauth2PermissionResult
    {
        /// <summary>
        /// The description of the admin consent
        /// </summary>
        public readonly string AdminConsentDescription;
        /// <summary>
        /// The display name of the admin consent
        /// </summary>
        public readonly string AdminConsentDisplayName;
        /// <summary>
        /// The unique identifier of the `app_role`.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Determines if the app role is enabled.
        /// </summary>
        public readonly bool IsEnabled;
        /// <summary>
        /// The type of the permission
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The description of the user consent
        /// </summary>
        public readonly string UserConsentDescription;
        /// <summary>
        /// The display name of the user consent
        /// </summary>
        public readonly string UserConsentDisplayName;
        /// <summary>
        /// Specifies the value of the roles claim that the application should expect in the authentication and access tokens.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetServicePrincipalOauth2PermissionResult(
            string adminConsentDescription,

            string adminConsentDisplayName,

            string id,

            bool isEnabled,

            string type,

            string userConsentDescription,

            string userConsentDisplayName,

            string value)
        {
            AdminConsentDescription = adminConsentDescription;
            AdminConsentDisplayName = adminConsentDisplayName;
            Id = id;
            IsEnabled = isEnabled;
            Type = type;
            UserConsentDescription = userConsentDescription;
            UserConsentDisplayName = userConsentDisplayName;
            Value = value;
        }
    }
}

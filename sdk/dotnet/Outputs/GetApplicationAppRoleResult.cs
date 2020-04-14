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
    public sealed class GetApplicationAppRoleResult
    {
        /// <summary>
        /// Specifies whether this app role definition can be assigned to users and groups, or to other applications (that are accessing this application in daemon service scenarios). Possible values are: `User` and `Application`, or both.
        /// </summary>
        public readonly ImmutableArray<string> AllowedMemberTypes;
        /// <summary>
        /// Permission help text that appears in the admin app assignment and consent experiences.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Display name for the permission that appears in the admin consent and app assignment experiences.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The unique identifier of the `app_role`.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Determines if the app role is enabled.
        /// </summary>
        public readonly bool IsEnabled;
        /// <summary>
        /// Specifies the value of the roles claim that the application should expect in the authentication and access tokens.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetApplicationAppRoleResult(
            ImmutableArray<string> allowedMemberTypes,

            string description,

            string displayName,

            string id,

            bool isEnabled,

            string value)
        {
            AllowedMemberTypes = allowedMemberTypes;
            Description = description;
            DisplayName = displayName;
            Id = id;
            IsEnabled = isEnabled;
            Value = value;
        }
    }
}

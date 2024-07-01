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
    public sealed class ApplicationPassword
    {
        /// <summary>
        /// A display name for the password. Changing this field forces a new resource to be created.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        /// </summary>
        public readonly string? EndDate;
        /// <summary>
        /// (Required) The unique key ID for the generated password.
        /// </summary>
        public readonly string? KeyId;
        /// <summary>
        /// The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
        /// </summary>
        public readonly string? StartDate;
        /// <summary>
        /// (Required) The generated password for the application.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private ApplicationPassword(
            string displayName,

            string? endDate,

            string? keyId,

            string? startDate,

            string? value)
        {
            DisplayName = displayName;
            EndDate = endDate;
            KeyId = keyId;
            StartDate = startDate;
            Value = value;
        }
    }
}
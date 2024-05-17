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
    public sealed class GroupRoleManagementPolicyEligibleAssignmentRules
    {
        /// <summary>
        /// Must an assignment have an expiry date. `false` allows permanent assignment.
        /// </summary>
        public readonly bool? ExpirationRequired;
        /// <summary>
        /// The maximum length of time an assignment can be valid, as an ISO8601 duration. Permitted values: `P15D`, `P30D`, `P90D`, `P180D`, or `P365D`.
        /// 
        /// One of `expiration_required` or `expire_after` must be provided.
        /// </summary>
        public readonly string? ExpireAfter;

        [OutputConstructor]
        private GroupRoleManagementPolicyEligibleAssignmentRules(
            bool? expirationRequired,

            string? expireAfter)
        {
            ExpirationRequired = expirationRequired;
            ExpireAfter = expireAfter;
        }
    }
}
// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class GroupRoleManagementPolicyActiveAssignmentRulesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Must an assignment have an expiry date. `false` allows permanent assignment.
        /// </summary>
        [Input("expirationRequired")]
        public Input<bool>? ExpirationRequired { get; set; }

        /// <summary>
        /// The maximum length of time an assignment can be valid, as an ISO8601 duration. Permitted values: `P15D`, `P30D`, `P90D`, `P180D`, or `P365D`.
        /// </summary>
        [Input("expireAfter")]
        public Input<string>? ExpireAfter { get; set; }

        /// <summary>
        /// Is a justification required to create new assignments.
        /// </summary>
        [Input("requireJustification")]
        public Input<bool>? RequireJustification { get; set; }

        /// <summary>
        /// Is multi-factor authentication required to create new assignments.
        /// </summary>
        [Input("requireMultifactorAuthentication")]
        public Input<bool>? RequireMultifactorAuthentication { get; set; }

        /// <summary>
        /// Is ticket information required to create new assignments.
        /// 
        /// One of `expiration_required` or `expire_after` must be provided.
        /// </summary>
        [Input("requireTicketInfo")]
        public Input<bool>? RequireTicketInfo { get; set; }

        public GroupRoleManagementPolicyActiveAssignmentRulesArgs()
        {
        }
        public static new GroupRoleManagementPolicyActiveAssignmentRulesArgs Empty => new GroupRoleManagementPolicyActiveAssignmentRulesArgs();
    }
}

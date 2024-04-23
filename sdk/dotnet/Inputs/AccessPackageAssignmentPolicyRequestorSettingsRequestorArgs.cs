// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class AccessPackageAssignmentPolicyRequestorSettingsRequestorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// For a user in an approval stage, this property indicates whether the user is a backup fallback approver
        /// </summary>
        [Input("backup")]
        public Input<bool>? Backup { get; set; }

        /// <summary>
        /// The object ID of the subject
        /// </summary>
        [Input("objectId")]
        public Input<string>? ObjectId { get; set; }

        /// <summary>
        /// Type of users
        /// </summary>
        [Input("subjectType", required: true)]
        public Input<string> SubjectType { get; set; } = null!;

        public AccessPackageAssignmentPolicyRequestorSettingsRequestorArgs()
        {
        }
        public static new AccessPackageAssignmentPolicyRequestorSettingsRequestorArgs Empty => new AccessPackageAssignmentPolicyRequestorSettingsRequestorArgs();
    }
}

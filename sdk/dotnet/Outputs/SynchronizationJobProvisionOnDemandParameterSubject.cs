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
    public sealed class SynchronizationJobProvisionOnDemandParameterSubject
    {
        /// <summary>
        /// The identifier of an object to which a synchronization job is to be applied. Can be one of the following: (1) An onPremisesDistinguishedName for synchronization from Active Directory to Azure AD. (2) The user ID for synchronization from Azure AD to a third-party. (3) The Worker ID of the Workday worker for synchronization from Workday to either Active Directory or Azure AD.
        /// </summary>
        public readonly string ObjectId;
        /// <summary>
        /// The type of the object to which a synchronization job is to be applied. Can be one of the following: `user` for synchronizing between Active Directory and Azure AD, `User` for synchronizing a user between Azure AD and a third-party application, `Worker` for synchronization a user between Workday and either Active Directory or Azure AD, `Group` for synchronizing a group between Azure AD and a third-party application.
        /// </summary>
        public readonly string ObjectTypeName;

        [OutputConstructor]
        private SynchronizationJobProvisionOnDemandParameterSubject(
            string objectId,

            string objectTypeName)
        {
            ObjectId = objectId;
            ObjectTypeName = objectTypeName;
        }
    }
}

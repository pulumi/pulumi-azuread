// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Outputs
{

    [OutputType]
    public sealed class CustomDirectoryRolePermission
    {
        /// <summary>
        /// A set of tasks that can be performed on a resource. For more information, see the [Permissions Reference](https://docs.microsoft.com/en-us/azure/active-directory/roles/permissions-reference) documentation.
        /// </summary>
        public readonly ImmutableArray<string> AllowedResourceActions;

        [OutputConstructor]
        private CustomDirectoryRolePermission(ImmutableArray<string> allowedResourceActions)
        {
            AllowedResourceActions = allowedResourceActions;
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD
{
    public static class GetDirectoryRoles
    {
        /// <summary>
        /// Use this data source to access information about activated directory roles within Azure Active Directory.
        /// 
        /// ## API Permissions
        /// 
        /// The following API permissions are required in order to use this resource.
        /// 
        /// When authenticated with a service principal, this resource requires one of the following application roles: `RoleManagement.Read.Directory` or `Directory.Read.All`
        /// 
        /// When authenticated with a user principal, this data source does not require any additional roles.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AzureAD = Pulumi.AzureAD;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var current = AzureAD.GetDirectoryRoles.Invoke();
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["roles"] = current.Apply(getDirectoryRolesResult =&gt; getDirectoryRolesResult.ObjectIds),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDirectoryRolesResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDirectoryRolesResult>("azuread:index/getDirectoryRoles:getDirectoryRoles", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetDirectoryRolesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The object IDs of the roles.
        /// </summary>
        public readonly ImmutableArray<string> ObjectIds;
        /// <summary>
        /// A list of users. Each `role` object provides the attributes documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDirectoryRolesRoleResult> Roles;

        [OutputConstructor]
        private GetDirectoryRolesResult(
            string id,

            ImmutableArray<string> objectIds,

            ImmutableArray<Outputs.GetDirectoryRolesRoleResult> roles)
        {
            Id = id;
            ObjectIds = objectIds;
            Roles = roles;
        }
    }
}

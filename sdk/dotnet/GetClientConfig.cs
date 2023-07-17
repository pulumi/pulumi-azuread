// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD
{
    public static class GetClientConfig
    {
        /// <summary>
        /// Use this data source to access the configuration of the AzureAD provider.
        /// 
        /// ## API Permissions
        /// 
        /// No additional roles are required to use this data source.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AzureAD = Pulumi.AzureAD;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var current = AzureAD.GetClientConfig.Invoke();
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["objectId"] = current.Apply(getClientConfigResult =&gt; getClientConfigResult.ObjectId),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetClientConfigResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetClientConfigResult>("azuread:index/getClientConfig:getClientConfig", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetClientConfigResult
    {
        /// <summary>
        /// The client ID (application ID) linked to the authenticated principal, or the application used for delegated authentication.
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The object ID of the authenticated principal.
        /// </summary>
        public readonly string ObjectId;
        /// <summary>
        /// The tenant ID of the authenticated principal.
        /// </summary>
        public readonly string TenantId;

        [OutputConstructor]
        private GetClientConfigResult(
            string clientId,

            string id,

            string objectId,

            string tenantId)
        {
            ClientId = clientId;
            Id = id;
            ObjectId = objectId;
            TenantId = tenantId;
        }
    }
}

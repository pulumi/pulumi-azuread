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
        /// Use this data source to access the configuration of the AzureRM provider.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AzureAD = Pulumi.AzureAD;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var current = Output.Create(AzureAD.GetClientConfig.InvokeAsync());
        ///         this.AccountId = current.Apply(current =&gt; current.ClientId);
        ///     }
        /// 
        ///     [Output("accountId")]
        ///     public Output&lt;string&gt; AccountId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetClientConfigResult> InvokeAsync(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClientConfigResult>("azuread:index/getClientConfig:getClientConfig", InvokeArgs.Empty, options.WithVersion());
    }


    [OutputType]
    public sealed class GetClientConfigResult
    {
        public readonly string ClientId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string ObjectId;
        public readonly string SubscriptionId;
        public readonly string TenantId;

        [OutputConstructor]
        private GetClientConfigResult(
            string clientId,

            string id,

            string objectId,

            string subscriptionId,

            string tenantId)
        {
            ClientId = clientId;
            Id = id;
            ObjectId = objectId;
            SubscriptionId = subscriptionId;
            TenantId = tenantId;
        }
    }
}
